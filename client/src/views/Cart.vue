<template>
  <div class="cart-page">
    <div class="container">
      <div class="page-header">
        <h1>Shopping Cart</h1>
        <p>Review and manage your items</p>
      </div>

      <!-- Loading State -->
      <div v-if="cartStore.loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading cart...</p>
      </div>

      <!-- Empty Cart -->
      <div v-else-if="cartStore.isEmpty" class="empty-cart">
        <el-empty description="Your cart is empty">
          <el-button type="primary" @click="$router.push('/products')">
            Continue Shopping
          </el-button>
        </el-empty>
      </div>

      <!-- Cart Items -->
      <div v-else class="cart-content">
        <div class="cart-items">
          <div
            v-for="item in cartStore.items"
            :key="item.id"
            class="cart-item"
          >
            <div class="item-image">
              <img
                :src="item.image_url || '/placeholder-product.jpg'"
                :alt="item.name"
                @error="handleImageError"
              />
            </div>

            <div class="item-info">
              <h4 class="item-name">{{ item.name }}</h4>
              <p class="item-description">{{ item.description }}</p>
              <div class="item-price">${{ item.price }}</div>
            </div>

            <div class="item-quantity">
              <el-input-number
                v-model="item.quantity"
                :min="1"
                :max="item.stock"
                @change="updateQuantity(item.id, $event)"
              />
            </div>

            <div class="item-total">
              <div class="total-price">${{ (item.price * item.quantity).toFixed(2) }}</div>
              <el-button
                type="danger"
                text
                @click="removeItem(item.id)"
              >
                <el-icon><Delete /></el-icon>
                Remove
              </el-button>
            </div>
          </div>
        </div>

        <!-- Cart Summary -->
        <div class="cart-summary">
          <el-card>
            <template #header>
              <div class="card-header">
                <h3>Cart Summary</h3>
              </div>
            </template>

            <div class="summary-details">
              <div class="summary-row">
                <span>Total Items:</span>
                <span>{{ cartStore.totalItems }}</span>
              </div>
              <div class="summary-row">
                <span>Subtotal:</span>
                <span>${{ cartStore.totalAmount.toFixed(2) }}</span>
              </div>
              <div class="summary-row">
                <span>Shipping:</span>
                <span>FREE</span>
              </div>
              <div class="summary-row total">
                <span>Total:</span>
                <span>${{ cartStore.totalAmount.toFixed(2) }}</span>
              </div>
            </div>

            <div class="summary-actions">
              <el-button
                type="primary"
                size="large"
                class="w-100"
                @click="$router.push('/checkout')"
              >
                Proceed to Checkout
              </el-button>
              <el-button
                size="large"
                @click="$router.push('/products')"
              >
                Continue Shopping
              </el-button>
            </div>
          </el-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useCartStore } from '../stores/cart.store.js'
import { useAuthStore } from '../stores/auth.store.js'
import { ElMessage } from 'element-plus'
import { Loading, Delete } from '@element-plus/icons-vue'

const cartStore = useCartStore()
const authStore = useAuthStore()

const loadCart = async () => {
  try {
    await cartStore.fetchCart()
  } catch (error) {
    console.error('Failed to load cart:', error)
  }
}

const updateQuantity = async (itemId, quantity) => {
  if (quantity <= 0) {
    await removeItem(itemId)
    return
  }

  try {
    await cartStore.updateCartItem(itemId, quantity)
    ElMessage.success('Cart updated!')
  } catch (error) {
    ElMessage.error('Failed to update cart')
    console.error('Update cart error:', error)
  }
}

const removeItem = async (itemId) => {
  try {
    await cartStore.removeFromCart(itemId)
    ElMessage.success('Item removed from cart!')
  } catch (error) {
    ElMessage.error('Failed to remove item')
    console.error('Remove item error:', error)
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(() => {
  loadCart()
})
</script>

<style scoped>
.cart-page {
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

.cart-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 2rem;
}

.cart-items {
  background: var(--white);
  border-radius: var(--border-radius-lg);
  box-shadow: var(--shadow-sm);
}

.cart-item {
  display: grid;
  grid-template-columns: 100px 1fr auto auto;
  gap: 1rem;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 1px solid var(--gray-200);
}

.cart-item:last-child {
  border-bottom: none;
}

.item-image {
  width: 80px;
  height: 80px;
  border-radius: var(--border-radius-md);
  overflow: hidden;
}

.item-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.item-info {
  flex: 1;
}

.item-name {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--gray-800);
  margin-bottom: 0.25rem;
}

.item-description {
  font-size: 0.875rem;
  color: var(--gray-600);
  margin-bottom: 0.5rem;
}

.item-price {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--primary-color);
}

.item-quantity {
  min-width: 120px;
}

.item-total {
  text-align: right;
}

.total-price {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--gray-800);
  margin-bottom: 0.5rem;
}

.cart-summary {
  position: sticky;
  top: 2rem;
  height: fit-content;
}

.summary-details {
  margin-bottom: 1.5rem;
}

.summary-row {
  display: flex;
  justify-content: space-between;
  margin-bottom: 0.75rem;
  font-size: 0.875rem;
}

.summary-row.total {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--primary-color);
  border-top: 1px solid var(--gray-200);
  padding-top: 0.75rem;
  margin-top: 1rem;
}

.summary-actions {
  display: flex;
  flex-direction: column;
  gap: 0.75rem;
}

.w-100 {
  width: 100%;
}

/* Empty Cart */
.empty-cart {
  text-align: center;
  padding: 4rem 2rem;
}

/* Loading */
.loading-container {
  text-align: center;
  padding: 3rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .cart-content {
    grid-template-columns: 1fr;
  }

  .cart-item {
    grid-template-columns: 80px 1fr;
    gap: 1rem;
  }

  .item-quantity,
  .item-total {
    grid-column: 1 / -1;
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-top: 1rem;
  }
}
</style>