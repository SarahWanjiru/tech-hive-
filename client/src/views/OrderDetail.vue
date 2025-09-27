<template>
  <div class="order-detail-page">
    <div class="container">
      <!-- Order detail functionality temporarily disabled -->
      <div class="disabled-message">
        <el-empty description="Order detail functionality is currently being updated">
          <el-button type="primary" @click="$router.push('/')">
            Go to Home
          </el-button>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth.store.ts'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
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