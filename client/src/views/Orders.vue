<template>
  <div class="orders-page">
    <div class="container">
      <div class="page-header">
        <h1>My Orders</h1>
        <p>Track and manage your orders</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading orders...</p>
      </div>

      <!-- Empty Orders -->
      <div v-else-if="orders.length === 0" class="empty-orders">
        <el-empty description="You haven't placed any orders yet">
          <el-button type="primary" @click="$router.push('/products')">
            Start Shopping
          </el-button>
        </el-empty>
      </div>

      <!-- Orders List -->
      <div v-else class="orders-list">
        <div
          v-for="order in orders"
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
              <el-tag
                :type="getStatusType(order.status)"
                size="large"
              >
                {{ order.status.toUpperCase() }}
              </el-tag>
            </div>
          </div>

          <div class="order-items">
            <div
              v-for="item in order.order_items"
              :key="item.id"
              class="order-item"
            >
              <div class="item-image">
                <img
                  :src="item.product.image_url || '/placeholder-product.svg'"
                  :alt="item.product.name"
                  @error="handleImageError"
                />
              </div>

              <div class="item-details">
                <h6>{{ item.product.name }}</h6>
                <p>{{ item.product.description }}</p>
                <div class="item-meta">
                  <span>Qty: {{ item.quantity }}</span>
                  <span>${{ item.product.price }}</span>
                </div>
              </div>

              <div class="item-total">
                ${{ (item.product.price * item.quantity).toFixed(2) }}
              </div>
            </div>
          </div>

          <div class="order-footer">
            <div class="order-total">
              Total: ${{ order.total.toFixed(2) }}
            </div>

            <div class="order-actions">
              <el-button
                v-if="order.status === 'pending' && order.payment?.status === 'pending'"
                type="primary"
                size="small"
                @click.stop="retryPayment(order)"
              >
                Retry Payment
              </el-button>

              <el-button
                v-if="order.status === 'pending'"
                type="danger"
                size="small"
                @click.stop="cancelOrder(order)"
              >
                Cancel Order
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth.store.ts'
import { useOrderStore } from '../stores/order.store.ts'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const authStore = useAuthStore()
const orderStore = useOrderStore()

const loading = ref(false)
const orders = ref([])

const loadOrders = async () => {
  loading.value = true
  try {
    const response = await orderStore.fetchOrders()
    orders.value = response.data || []
  } catch (error) {
    ElMessage.error('Failed to load orders')
    console.error('Load orders error:', error)
    orders.value = [] // Ensure orders is always an array even on error
  } finally {
    loading.value = false
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
    'confirmed': 'info',
    'shipped': 'primary',
    'delivered': 'success',
    'cancelled': 'danger'
  }
  return statusMap[status] || 'info'
}

const retryPayment = async (order) => {
  try {
    const { value: phoneNumber } = await ElMessageBox.prompt(
      'Enter your M-Pesa phone number',
      'Retry Payment',
      {
        confirmButtonText: 'Send STK Push',
        cancelButtonText: 'Cancel',
        inputPattern: /^254[0-9]{9}$/,
        inputErrorMessage: 'Please enter a valid M-Pesa number (254XXXXXXXXX)',
        inputValue: order.payment?.phone_number || ''
      }
    )

    const paymentData = {
      order_id: order.id,
      phone_number: phoneNumber,
      amount: order.total
    }

    await orderStore.processMpesaPayment(paymentData)
    ElMessage.success('STK Push sent to your phone')

    // Reload orders to update status
    loadOrders()

  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to retry payment')
      console.error('Retry payment error:', error)
    }
  }
}

const cancelOrder = async (order) => {
  try {
    await ElMessageBox.confirm(
      'Are you sure you want to cancel this order?',
      'Cancel Order',
      {
        confirmButtonText: 'Cancel Order',
        cancelButtonText: 'Keep Order',
        type: 'warning',
      }
    )

    await orderStore.updateOrderStatus(order.id, 'cancelled')
    ElMessage.success('Order cancelled')

    // Reload orders to update status
    loadOrders()

  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to cancel order')
      console.error('Cancel order error:', error)
    }
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.svg'
}

onMounted(() => {
  if (authStore.isAuthenticated) {
    loadOrders()
  } else {
    ElMessage.warning('Please login to view your orders')
  }
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