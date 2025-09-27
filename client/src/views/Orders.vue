<template>
  <div class="orders-page">
    <div class="container">
      <div class="page-header">
        <h1>My Orders</h1>
        <p>Track and manage your orders</p>
      </div>

      <!-- Orders functionality temporarily disabled -->
      <div class="disabled-message">
        <el-empty description="Orders functionality is currently being updated">
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