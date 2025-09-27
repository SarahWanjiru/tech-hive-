<template>
  <div class="orders-page">
    <div class="container">
      <div class="page-header">
        <h1>My Orders</h1>
        <p>Track and manage your orders</p>
      </div>

      <!-- Loading State -->
      <div v-if="orderStore.loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading orders...</p>
      </div>

      <!-- Empty Orders -->
      <div v-else-if="orderStore.orders.length === 0" class="empty-orders">
        <el-empty description="No orders found">
          <el-button type="primary" @click="$router.push('/products')">
            Start Shopping
          </el-button>
        </el-empty>
      </div>

      <!-- Orders List -->
      <div v-else class="orders-list">
        <div
          v-for="order in orderStore.orders"
          :key="order.id"
          class="order-card"
          @click="$router.push(`/orders/${order.id}`)"
        >
          <div class="order-header">
            <div class="order-info">
              <h4>Order #{{ order.id }}</h4>
              <p class="order-date">{{ formatDate(order.created_at) }}</p>
            </div>
            <div class="order-status">
              <el-tag :type="getStatusType(order.status)">
                {{ getStatusText(order.status) }}
              </el-tag>
            </div>
          </div>

          <div class="order-items">
            <div
              v-for="item in order.items"
              :key="item.id"
              class="order-item"
            >
              <div class="item-image">
                <img
                  :src="item.image_url || '/placeholder-product.jpg'"
                  :alt="item.name"
                  @error="handleImageError"
                />
              </div>
              <div class="item-details">
                <h6>{{ item.name }}</h6>
                <p>{{ item.description }}</p>
                <div class="item-meta">
                  <span>Qty: {{ item.quantity }}</span>
                  <span>${{ item.price }} each</span>
                </div>
              </div>
              <div class="item-total">
                ${{ (item.price * item.quantity).toFixed(2) }}
              </div>
            </div>
          </div>

          <div class="order-footer">
            <div class="order-total">
              <span>Total: ${{ order.total_amount }}</span>
            </div>
            <div class="order-actions">
              <el-button
                v-if="order.status === 'pending'"
                type="danger"
                size="small"
                @click.stop="cancelOrder(order.id)"
              >
                Cancel Order
              </el-button>
              <el-button size="small" @click.stop="$router.push(`/orders/${order.id}`)">
                View Details
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useOrderStore } from '../stores/order.store.js'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const orderStore = useOrderStore()

const loadOrders = async () => {
  try {
    await orderStore.fetchOrders()
  } catch (error) {
    console.error('Failed to load orders:', error)
  }
}

const cancelOrder = async (orderId) => {
  try {
    await ElMessage.confirm(
      'Are you sure you want to cancel this order?',
      'Cancel Order',
      {
        confirmButtonText: 'Yes, Cancel',
        cancelButtonText: 'No',
        type: 'warning',
      }
    )

    await orderStore.cancelOrder(orderId)
    ElMessage.success('Order cancelled successfully!')

    // Reload orders
    loadOrders()
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to cancel order')
      console.error('Cancel order error:', error)
    }
  }
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusType = (status) => {
  const statusMap = {
    'pending': 'warning',
    'processing': 'info',
    'completed': 'success',
    'cancelled': 'danger'
  }
  return statusMap[status] || 'info'
}

const getStatusText = (status) => {
  const statusMap = {
    'pending': 'Pending',
    'processing': 'Processing',
    'completed': 'Completed',
    'cancelled': 'Cancelled'
  }
  return statusMap[status] || 'Unknown'
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(() => {
  loadOrders()
})
</script>

<style scoped>
.orders-page {
  padding: 2rem 0;
}

.page-header {
  text-align: center;
  margin-bottom: 2rem;
}

.page-header h1 {
  font-size: 2.5rem;
  color: var(--gray-800);
  margin-bottom: 0.5rem;
}

.page-header p {
  font-size: 1.125rem;
  color: var(--gray-600);
}

.orders-list {
  display: flex;
  flex-direction: column;
  gap: 1.5rem;
}

.order-card {
  background: var(--white);
  border-radius: var(--border-radius-lg);
  box-shadow: var(--shadow-sm);
  padding: 1.5rem;
  cursor: pointer;
  transition: all var(--transition-fast);
}

.order-card:hover {
  box-shadow: var(--shadow-md);
  transform: translateY(-2px);
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid var(--gray-200);
}

.order-info h4 {
  margin: 0 0 0.25rem 0;
  color: var(--gray-800);
}

.order-date {
  margin: 0;
  font-size: 0.875rem;
  color: var(--gray-600);
}

.order-items {
  margin-bottom: 1rem;
}

.order-item {
  display: flex;
  gap: 1rem;
  padding: 0.75rem 0;
  border-bottom: 1px solid var(--gray-100);
}

.order-item:last-child {
  border-bottom: none;
}

.item-image {
  width: 50px;
  height: 50px;
  border-radius: var(--border-radius-md);
  overflow: hidden;
  flex-shrink: 0;
}

.item-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-details {
  flex: 1;
}

.item-details h6 {
  margin: 0 0 0.25rem 0;
  font-size: 0.875rem;
  color: var(--gray-800);
}

.item-details p {
  margin: 0 0 0.25rem 0;
  font-size: 0.75rem;
  color: var(--gray-600);
}

.item-meta {
  display: flex;
  gap: 1rem;
  font-size: 0.75rem;
  color: var(--gray-600);
}

.item-total {
  font-weight: 600;
  color: var(--primary-color);
  font-size: 0.875rem;
}

.order-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-top: 1rem;
  border-top: 1px solid var(--gray-200);
}

.order-total {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--primary-color);
}

.order-actions {
  display: flex;
  gap: 0.5rem;
}

/* Loading and Empty States */
.loading-container,
.empty-orders {
  text-align: center;
  padding: 3rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .order-header {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }

  .order-footer {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .order-actions {
    justify-content: center;
  }
}
</style>