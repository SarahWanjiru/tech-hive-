<template>
  <div class="admin-page">
    <div class="container">
      <div class="page-header">
        <h1>Admin Dashboard</h1>
        <p>Manage your store</p>
      </div>

      <!-- Admin functionality temporarily disabled -->
      <div class="disabled-message">
        <el-empty description="Admin functionality is currently being updated">
          <el-button type="primary" @click="$router.push('/')">
            Go to Home
          </el-button>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useAuthStore } from '../stores/auth.store.ts'

const authStore = useAuthStore()
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