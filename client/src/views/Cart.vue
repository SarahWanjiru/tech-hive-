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
                :src="item.product.image_url || '/placeholder-product.svg'"
                :alt="item.product.name"
                @error="handleImageError"
              />
            </div>

            <div class="item-info">
              <h4 class="item-name">{{ item.product.name }}</h4>
              <p class="item-description">{{ item.product.description }}</p>
              <p class="item-price">${{ item.product.price }}</p>
            </div>

            <div class="item-quantity">
              <el-input-number
                v-model="item.quantity"
                :min="1"
                :max="item.product.stock"
                @change="updateQuantity(item.id, $event)"
              />
            </div>

            <div class="item-total">
              <p class="total-price">${{ (item.product.price * item.quantity).toFixed(2) }}</p>
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

        <div class="cart-summary">
          <h3>Order Summary</h3>

          <div class="summary-details">
            <div class="summary-row">
              <span>Subtotal ({{ cartStore.totalItems }} items)</span>
              <span>${{ cartStore.totalAmount.toFixed(2) }}</span>
            </div>
            <div class="summary-row">
              <span>Shipping</span>
              <span>FREE</span>
            </div>
            <div class="summary-row">
              <span>Tax</span>
              <span>${{ (cartStore.totalAmount * 0.1).toFixed(2) }}</span>
            </div>
            <div class="summary-row total">
              <span>Total</span>
              <span>${{ (cartStore.totalAmount * 1.1).toFixed(2) }}</span>
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
              class="w-100"
              @click="$router.push('/products')"
            >
              Continue Shopping
            </el-button>
            <el-button
              type="danger"
              text
              @click="clearCart"
            >
              Clear Cart
            </el-button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useCartStore } from '../stores/cart.store.ts'
import { useAuthStore } from '../stores/auth.store.ts'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Loading, Delete } from '@element-plus/icons-vue'

const cartStore = useCartStore()
const authStore = useAuthStore()

const loadCart = async () => {
  try {
    await cartStore.fetchCart()
  } catch (error) {
    ElMessage.error('Failed to load cart')
    console.error('Cart loading error:', error)
  }
}

const updateQuantity = async (itemId, quantity) => {
  try {
    await cartStore.updateCartItem(itemId, quantity)
    ElMessage.success('Quantity updated')
  } catch (error) {
    ElMessage.error('Failed to update quantity')
    console.error('Update quantity error:', error)
    // Reload cart to revert changes
    loadCart()
  }
}

const removeItem = async (itemId) => {
  try {
    await ElMessageBox.confirm(
      'Are you sure you want to remove this item from your cart?',
      'Confirm Removal',
      {
        confirmButtonText: 'Remove',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
    )

    await cartStore.removeFromCart(itemId)
    ElMessage.success('Item removed from cart')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to remove item')
      console.error('Remove item error:', error)
    }
  }
}

const clearCart = async () => {
  try {
    await ElMessageBox.confirm(
      'Are you sure you want to clear your entire cart?',
      'Clear Cart',
      {
        confirmButtonText: 'Clear All',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
    )

    await cartStore.clearCart()
    ElMessage.success('Cart cleared')
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to clear cart')
      console.error('Clear cart error:', error)
    }
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(() => {
  if (authStore.isAuthenticated) {
    loadCart()
  } else {
    ElMessage.warning('Please login to view your cart')
  }
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