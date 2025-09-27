<template>
  <div class="checkout-page">
    <div class="container">
      <div class="page-header">
        <h1>Checkout</h1>
        <p>Complete your order</p>
      </div>

      <!-- Loading State -->
      <div v-if="orderStore.loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Processing your order...</p>
      </div>

      <!-- Empty Cart -->
      <div v-else-if="cartStore.isEmpty" class="empty-cart">
        <el-empty description="Your cart is empty">
          <el-button type="primary" @click="$router.push('/products')">
            Continue Shopping
          </el-button>
        </el-empty>
      </div>

      <!-- Checkout Form -->
      <div v-else class="checkout-content">
        <!-- Order Summary -->
        <div class="order-summary">
          <el-card>
            <template #header>
              <div class="card-header">
                <h3>Order Summary</h3>
              </div>
            </template>

            <div class="order-items">
              <div
                v-for="item in cartStore.items"
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

            <div class="order-totals">
              <div class="total-row">
                <span>Subtotal:</span>
                <span>${{ cartStore.totalAmount.toFixed(2) }}</span>
              </div>
              <div class="total-row">
                <span>Shipping:</span>
                <span>FREE</span>
              </div>
              <div class="total-row final-total">
                <span>Total:</span>
                <span>${{ cartStore.totalAmount.toFixed(2) }}</span>
              </div>
            </div>
          </el-card>
        </div>

        <!-- Shipping & Payment -->
        <div class="checkout-form">
          <el-card>
            <template #header>
              <div class="card-header">
                <h3>Shipping & Payment</h3>
              </div>
            </template>

            <el-form
              ref="checkoutFormRef"
              :model="checkoutData"
              :rules="checkoutRules"
              @submit.prevent="processCheckout"
            >
              <!-- Shipping Address -->
              <div class="form-section">
                <h4>Shipping Address</h4>
                <el-form-item label="Full Address" prop="shippingAddress">
                  <el-input
                    v-model="checkoutData.shippingAddress"
                    type="textarea"
                    :rows="3"
                    placeholder="Enter your complete shipping address"
                  />
                </el-form-item>
              </div>

              <!-- Payment Method -->
              <div class="form-section">
                <h4>Payment Method</h4>
                <el-radio-group v-model="checkoutData.paymentMethod">
                  <el-radio value="mpesa">M-Pesa</el-radio>
                  <el-radio value="card" disabled>Card (Coming Soon)</el-radio>
                </el-radio-group>

                <!-- M-Pesa Payment -->
                <div v-if="checkoutData.paymentMethod === 'mpesa'" class="mpesa-section">
                  <el-form-item label="M-Pesa Phone Number" prop="phoneNumber">
                    <el-input
                      v-model="checkoutData.phoneNumber"
                      placeholder="254712345678"
                    />
                  </el-form-item>
                </div>
              </div>

              <!-- Place Order -->
              <div class="form-actions">
                <el-button
                  type="primary"
                  size="large"
                  native-type="submit"
                  class="w-100"
                  :loading="orderStore.loading"
                >
                  {{ orderStore.loading ? 'Processing...' : 'Place Order' }}
                </el-button>
              </div>
            </el-form>
          </el-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart.store.js'
import { useOrderStore } from '../stores/order.store.js'
import { useAuthStore } from '../stores/auth.store.js'
import { usePaymentStore } from '../stores/payment.store.js'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const router = useRouter()
const cartStore = useCartStore()
const orderStore = useOrderStore()
const authStore = useAuthStore()

const checkoutFormRef = ref(null)

const checkoutData = reactive({
  shippingAddress: '',
  paymentMethod: 'mpesa',
  phoneNumber: ''
})

const checkoutRules = {
  shippingAddress: [
    { required: true, message: 'Please enter shipping address', trigger: 'blur' }
  ],
  phoneNumber: [
    { required: true, message: 'Please enter M-Pesa phone number', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value && !paymentStore.validatePhoneNumber(value)) {
          callback(new Error('Please enter a valid Kenyan phone number'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ]
}

const processCheckout = async () => {
  try {
    await checkoutFormRef.value.validate()

    // Create order
    const orderResponse = await orderStore.createOrder({
      cartId: 1, // This should come from cart store
      shippingAddress: checkoutData.shippingAddress
    })

    // Process payment
    if (checkoutData.paymentMethod === 'mpesa') {
      const paymentResponse = await paymentStore.initiateSTKPush({
        orderId: orderResponse.data.id,
        phoneNumber: paymentStore.formatPhoneNumber(checkoutData.phoneNumber),
        amount: cartStore.totalAmount
      })

      ElMessage.success('STK Push sent to your phone!')
    }

    // Clear cart
    await cartStore.clearCart()

    // Redirect to orders page
    router.push('/orders')

  } catch (error) {
    console.error('Checkout error:', error)
    ElMessage.error('Failed to process order')
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(() => {
  // Ensure cart is loaded
  if (cartStore.isEmpty) {
    cartStore.fetchCart().catch(console.error)
  }
})
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