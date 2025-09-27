<template>
  <div class="order-detail-page">
    <div class="container">
      <!-- Loading State -->
      <div v-if="orderStore.loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading order details...</p>
      </div>

      <!-- Order Details -->
      <div v-else-if="orderStore.currentOrder" class="order-details">
        <div class="order-header">
          <h1>Order #{{ orderStore.currentOrder.id }}</h1>
          <el-tag :type="getStatusType(orderStore.currentOrder.status)">
            {{ getStatusText(orderStore.currentOrder.status) }}
          </el-tag>
        </div>

        <div class="order-info-cards">
          <el-card>
            <template #header>
              <h3>Order Information</h3>
            </template>
            <div class="order-info">
              <p><strong>Order Date:</strong> {{ formatDate(orderStore.currentOrder.created_at) }}</p>
              <p><strong>Shipping Address:</strong> {{ orderStore.currentOrder.shipping_address }}</p>
              <p><strong>Total Amount:</strong> ${{ orderStore.currentOrder.total_amount }}</p>
            </div>
          </el-card>

          <el-card>
            <template #header>
              <h3>Order Items</h3>
            </template>
            <div class="order-items">
              <div
                v-for="item in orderStore.currentOrder.items"
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
                  <h5>{{ item.name }}</h5>
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
          </el-card>
        </div>

        <div class="order-actions">
          <el-button
            v-if="orderStore.currentOrder.status === 'pending'"
            type="danger"
            @click="cancelOrder"
          >
            Cancel Order
          </el-button>
          <el-button @click="$router.push('/orders')">
            Back to Orders
          </el-button>
        </div>
      </div>

      <!-- Order Not Found -->
      <div v-else class="not-found">
        <el-empty description="Order not found">
          <el-button type="primary" @click="$router.push('/orders')">
            View My Orders
          </el-button>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useOrderStore } from '../stores/order.store.js'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const orderStore = useOrderStore()

const loadOrder = async () => {
  try {
    await orderStore.fetchOrderById(route.params.id)
  } catch (error) {
    console.error('Failed to load order:', error)
  }
}

const cancelOrder = async () => {
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

    await orderStore.cancelOrder(route.params.id)
    ElMessage.success('Order cancelled successfully!')

    // Reload order
    loadOrder()
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
    month: 'long',
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
  loadOrder()
})
</script>

<style scoped>
.order-detail-page {
  padding: 2rem 0;
}

.order-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 1px solid var(--gray-200);
}

.order-header h1 {
  margin: 0;
  color: var(--gray-800);
}

.order-info-cards {
  display: grid;
  grid-template-columns: 1fr 2fr;
  gap: 2rem;
  margin-bottom: 2rem;
}

.order-info {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.order-items {
  max-height: 400px;
  overflow-y: auto;
}

.order-item {
  display: flex;
  gap: 1rem;
  padding: 1rem;
  border-bottom: 1px solid var(--gray-200);
}

.order-item:last-child {
  border-bottom: none;
}

.item-image {
  width: 60px;
  height: 60px;
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

.item-details h5 {
  margin: 0 0 0.25rem 0;
  font-size: 1rem;
  color: var(--gray-800);
}

.item-details p {
  margin: 0 0 0.5rem 0;
  font-size: 0.875rem;
  color: var(--gray-600);
}

.item-meta {
  display: flex;
  gap: 1rem;
  font-size: 0.875rem;
  color: var(--gray-600);
}

.item-total {
  font-weight: 600;
  color: var(--primary-color);
  text-align: right;
}

.order-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
}

/* Loading and Not Found States */
.loading-container,
.not-found {
  text-align: center;
  padding: 3rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .order-info-cards {
    grid-template-columns: 1fr;
  }

  .order-header {
    flex-direction: column;
    gap: 1rem;
    align-items: flex-start;
  }

  .order-actions {
    flex-direction: column;
  }
}
</style>