<template>
  <div class="product-detail-page">
    <div class="container">
      <!-- Product detail functionality temporarily disabled -->
      <div class="disabled-message">
        <el-empty description="Product detail functionality is currently being updated">
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
.product-detail-page {
  padding: 2rem 0;
}

.product-details {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 3rem;
  background: var(--white);
  border-radius: var(--border-radius-lg);
  padding: 2rem;
  box-shadow: var(--shadow-sm);
}

.product-gallery {
  display: flex;
  flex-direction: column;
}

.main-image {
  width: 100%;
  height: 400px;
  border-radius: var(--border-radius-lg);
  overflow: hidden;
  margin-bottom: 1rem;
}

.main-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.product-info {
  padding: 1rem 0;
}

.product-title {
  font-size: 2rem;
  font-weight: 700;
  color: var(--gray-800);
  margin-bottom: 1rem;
}

.product-description {
  font-size: 1.125rem;
  color: var(--gray-600);
  line-height: 1.6;
  margin-bottom: 2rem;
}

.product-meta {
  margin-bottom: 2rem;
}

.product-price {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  background: var(--gray-50);
  border-radius: var(--border-radius-md);
}

.price {
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--primary-color);
}

.stock {
  font-size: 0.875rem;
  color: var(--gray-600);
}

.product-actions {
  display: flex;
  gap: 1rem;
  align-items: center;
}

/* Loading and Error States */
.loading-container,
.error-container,
.not-found {
  text-align: center;
  padding: 3rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .product-details {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .product-actions {
    flex-direction: column;
    align-items: stretch;
  }
}
</style>