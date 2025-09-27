<template>
  <div class="admin-page">
    <div class="container">
      <div class="page-header">
        <h1>Admin Dashboard</h1>
        <p>Manage your store</p>
      </div>

      <!-- Stats Cards -->
      <div class="stats-grid">
        <el-card>
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Document /></el-icon>
            </div>
            <div class="stat-info">
              <h3>{{ orderStore.stats.totalOrders }}</h3>
              <p>Total Orders</p>
            </div>
          </div>
        </el-card>

        <el-card>
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Timer /></el-icon>
            </div>
            <div class="stat-info">
              <h3>{{ orderStore.stats.pendingOrders }}</h3>
              <p>Pending Orders</p>
            </div>
          </div>
        </el-card>

        <el-card>
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Check /></el-icon>
            </div>
            <div class="stat-info">
              <h3>{{ orderStore.stats.completedOrders }}</h3>
              <p>Completed Orders</p>
            </div>
          </div>
        </el-card>

        <el-card>
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Money /></el-icon>
            </div>
            <div class="stat-info">
              <h3>${{ orderStore.stats.totalRevenue }}</h3>
              <p>Total Revenue</p>
            </div>
          </div>
        </el-card>
      </div>

      <!-- Admin Actions -->
      <div class="admin-actions">
        <el-card>
          <template #header>
            <h3>Management Actions</h3>
          </template>

          <div class="action-buttons">
            <el-button type="primary" @click="$router.push('/products')">
              <el-icon><Box /></el-icon>
              Manage Products
            </el-button>
            <el-button type="success" @click="$router.push('/orders')">
              <el-icon><Document /></el-icon>
              Manage Orders
            </el-button>
            <el-button type="info" @click="seedDatabase">
              <el-icon><DataBoard /></el-icon>
              Seed Database
            </el-button>
          </div>
        </el-card>
      </div>

      <!-- Recent Orders -->
      <div class="recent-section">
        <el-card>
          <template #header>
            <h3>Recent Orders</h3>
          </template>

          <div v-if="orderStore.loading" class="loading-container">
            <el-icon class="loading"><Loading /></el-icon>
            <p>Loading orders...</p>
          </div>

          <div v-else-if="orderStore.orders.length === 0" class="empty-state">
            <p>No orders yet</p>
          </div>

          <div v-else class="recent-orders">
            <div
              v-for="order in orderStore.recentOrders"
              :key="order.id"
              class="recent-order"
            >
              <div class="order-info">
                <h5>Order #{{ order.id }}</h5>
                <p>{{ formatDate(order.created_at) }}</p>
              </div>
              <div class="order-status">
                <el-tag :type="getStatusType(order.status)">
                  {{ getStatusText(order.status) }}
                </el-tag>
              </div>
              <div class="order-total">
                ${{ order.total_amount }}
              </div>
            </div>
          </div>
        </el-card>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useOrderStore } from '../stores/order.store.js'
import { useAuthStore } from '../stores/auth.store.js'
import { ElMessage } from 'element-plus'
import { Document, Timer, Check, Money, Box, DataBoard, Loading } from '@element-plus/icons-vue'

const orderStore = useOrderStore()
const authStore = useAuthStore()

const loadDashboardData = async () => {
  try {
    await Promise.all([
      orderStore.fetchOrders({ limit: 5 }),
      orderStore.fetchOrderStats()
    ])
  } catch (error) {
    console.error('Failed to load dashboard data:', error)
  }
}

const seedDatabase = async () => {
  try {
    await ElMessage.confirm(
      'This will add sample data to the database. Continue?',
      'Seed Database',
      {
        confirmButtonText: 'Yes, Seed',
        cancelButtonText: 'Cancel',
        type: 'info',
      }
    )

    // Call seed endpoint
    const response = await fetch('http://localhost:9999/v1/api/seed/all', {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })

    if (response.ok) {
      ElMessage.success('Database seeded successfully!')
      // Reload dashboard data
      loadDashboardData()
    } else {
      throw new Error('Failed to seed database')
    }
  } catch (error) {
    ElMessage.error('Failed to seed database')
    console.error('Seed database error:', error)
  }
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
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

onMounted(() => {
  loadDashboardData()
})
</script>

<style scoped>
.admin-page {
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

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.stat-icon {
  width: 50px;
  height: 50px;
  background: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--white);
  font-size: 20px;
}

.stat-info h3 {
  margin: 0;
  font-size: 2rem;
  color: var(--gray-800);
}

.stat-info p {
  margin: 0;
  color: var(--gray-600);
  font-size: 0.875rem;
}

.admin-actions,
.recent-section {
  margin-bottom: 2rem;
}

.action-buttons {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.recent-orders {
  max-height: 400px;
  overflow-y: auto;
}

.recent-order {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid var(--gray-200);
}

.recent-order:last-child {
  border-bottom: none;
}

.order-info h5 {
  margin: 0 0 0.25rem 0;
  color: var(--gray-800);
}

.order-info p {
  margin: 0;
  font-size: 0.875rem;
  color: var(--gray-600);
}

.order-total {
  font-weight: 600;
  color: var(--primary-color);
}

/* Loading and Empty States */
.loading-container,
.empty-state {
  text-align: center;
  padding: 2rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .action-buttons {
    grid-template-columns: 1fr;
  }

  .recent-order {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }
}
</style>