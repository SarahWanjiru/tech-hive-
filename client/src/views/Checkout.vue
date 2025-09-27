<template>
  <div class="checkout-page">
    <div class="container">
      <div class="page-header">
        <h1>Checkout</h1>
        <p>Complete your order</p>
      </div>

      <!-- Checkout functionality temporarily disabled -->
      <div class="disabled-message">
        <el-empty description="Checkout functionality is currently being updated">
          <el-button type="primary" @click="$router.push('/')">
            Go to Home
          </el-button>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth.store.ts'

const router = useRouter()
const authStore = useAuthStore()
</script>

<style scoped>
.checkout-page {
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

.checkout-content {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 2rem;
}

.order-summary,
.checkout-form {
  height: fit-content;
}

.order-items {
  margin-bottom: 1.5rem;
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

.order-totals {
  border-top: 1px solid var(--gray-200);
  padding-top: 1rem;
}

.total-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.5rem;
  font-size: 0.875rem;
}

.total-row.final-total {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--primary-color);
  border-top: 1px solid var(--gray-200);
  padding-top: 0.75rem;
  margin-top: 1rem;
}

.form-section {
  margin-bottom: 2rem;
}

.form-section h4 {
  margin-bottom: 1rem;
  color: var(--gray-800);
  font-size: 1.125rem;
}

.mpesa-section {
  margin-top: 1rem;
  padding: 1rem;
  background: var(--gray-50);
  border-radius: var(--border-radius-md);
}

.form-actions {
  margin-top: 2rem;
  text-align: center;
}

.w-100 {
  width: 100%;
}

/* Loading and Empty States */
.loading-container,
.empty-cart {
  text-align: center;
  padding: 3rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .checkout-content {
    grid-template-columns: 1fr;
  }

  .order-item {
    flex-direction: column;
    text-align: center;
  }

  .item-total {
    text-align: center;
    margin-top: 0.5rem;
  }
}
</style>