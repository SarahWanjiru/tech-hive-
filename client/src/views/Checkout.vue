<template>
  <div class="checkout-page">
    <div class="container">
      <div class="page-header">
        <h1>Checkout</h1>
        <p>Complete your order</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="loading-container">
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
          <h3>Order Summary</h3>

          <div class="order-items">
            <div
              v-for="item in cartStore.items"
              :key="item.id"
              class="order-item"
            >
              <div class="item-image">
                <img
                  :src="item.product.image_url || '/placeholder-product.jpg'"
                  :alt="item.product.name"
                  @error="handleImageError"
                />
              </div>

              <div class="item-details">
                <h5>{{ item.product.name }}</h5>
                <p>{{ item.product.description }}</p>
                <div class="item-meta">
                  <span>Qty: {{ item.quantity }}</span>
                  <span>KSh {{ item.product.price }}</span>
                </div>
              </div>

              <div class="item-total">
                KSh {{ (item.product.price * item.quantity).toFixed(2) }}
              </div>
            </div>
          </div>

          <div class="order-totals">
            <div class="total-row">
              <span>Subtotal</span>
              <span>KSh {{ cartStore.totalAmount.toFixed(2) }}</span>
            </div>
            <div class="total-row">
              <span>Shipping</span>
              <span>FREE</span>
            </div>
            <div class="total-row">
              <span>Tax</span>
              <span>KSh {{ (cartStore.totalAmount * 0.1).toFixed(2) }}</span>
            </div>
            <div class="total-row final-total">
              <span>Total</span>
              <span>KSh {{ (cartStore.totalAmount * 1.1).toFixed(2) }}</span>
            </div>
          </div>
        </div>

        <!-- Checkout Form -->
        <div class="checkout-form">
          <el-form
            ref="checkoutForm"
            :model="checkoutData"
            :rules="rules"
            label-position="top"
            @submit.prevent="processCheckout"
          >
            <!-- Shipping Information -->
            <div class="form-section">
              <h4>Shipping Information</h4>

              <el-row :gutter="20">
                <el-col :span="24" :md="12">
                  <el-form-item label="Full Name" prop="fullName">
                    <el-input
                      v-model="checkoutData.fullName"
                      placeholder="Enter your full name"
                    />
                  </el-form-item>
                </el-col>

                <el-col :span="24" :md="12">
                  <el-form-item label="Email" prop="email">
                    <el-input
                      v-model="checkoutData.email"
                      placeholder="Enter your email"
                    />
                  </el-form-item>
                </el-col>
              </el-row>

              <el-form-item label="Shipping Address" prop="address">
                <el-input
                  v-model="checkoutData.address"
                  type="textarea"
                  :rows="3"
                  placeholder="Enter your complete shipping address"
                />
              </el-form-item>

              <el-row :gutter="20">
                <el-col :span="24" :md="12">
                  <el-form-item label="City" prop="city">
                    <el-input
                      v-model="checkoutData.city"
                      placeholder="Enter your city"
                    />
                  </el-form-item>
                </el-col>

                <el-col :span="24" :md="12">
                  <el-form-item label="Phone Number" prop="phone">
                    <el-input
                      v-model="checkoutData.phone"
                      placeholder="254XXXXXXXXX"
                    />
                  </el-form-item>
                </el-col>
              </el-row>
            </div>

            <!-- Payment Information -->
            <div class="form-section">
              <h4>Payment Information</h4>

              <div class="mpesa-section">
                <p class="mpesa-info">
                  You will receive an M-Pesa STK Push notification on your phone to complete the payment.
                </p>

                <el-form-item label="M-Pesa Phone Number" prop="mpesaPhone">
                  <el-input
                    v-model="checkoutData.mpesaPhone"
                    placeholder="254XXXXXXXXX"
                    @input="formatPhoneNumber"
                  />
                </el-form-item>
              </div>
            </div>

            <!-- Form Actions -->
            <div class="form-actions">
              <el-button
                type="primary"
                size="large"
                native-type="submit"
                :loading="processingPayment"
                class="w-100"
              >
                {{ processingPayment ? 'Processing...' : 'Place Order' }}
              </el-button>

              <el-button
                size="large"
                @click="$router.push('/cart')"
                class="w-100"
              >
                Back to Cart
              </el-button>
            </div>
          </el-form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart.store.ts'
import { useAuthStore } from '../stores/auth.store.ts'
import { useOrderStore } from '../stores/order.store.ts'
import { ElMessage, ElMessageBox } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const router = useRouter()
const cartStore = useCartStore()
const authStore = useAuthStore()
const orderStore = useOrderStore()

const loading = ref(false)
const processingPayment = ref(false)
const checkoutForm = ref()

const checkoutData = reactive({
  fullName: '',
  email: '',
  address: '',
  city: '',
  phone: '',
  mpesaPhone: ''
})

const rules = {
  fullName: [
    { required: true, message: 'Please enter your full name', trigger: 'blur' }
  ],
  email: [
    { required: true, message: 'Please enter your email', trigger: 'blur' },
    { type: 'email', message: 'Please enter a valid email', trigger: 'blur' }
  ],
  address: [
    { required: true, message: 'Please enter your shipping address', trigger: 'blur' }
  ],
  city: [
    { required: true, message: 'Please enter your city', trigger: 'blur' }
  ],
  phone: [
    { required: true, message: 'Please enter your phone number', trigger: 'blur' }
  ],
  mpesaPhone: [
    { required: true, message: 'Please enter your M-Pesa phone number', trigger: 'blur' }
  ]
}

const totalAmount = computed(() => {
  return cartStore.totalAmount * 1.1 // Including tax
})

const formatPhoneNumber = () => {
  // Format phone number to start with 254
  let phone = checkoutData.mpesaPhone.replace(/\D/g, '')
  if (phone.startsWith('0')) {
    phone = '254' + phone.substring(1)
  } else if (!phone.startsWith('254')) {
    phone = '254' + phone
  }
  checkoutData.mpesaPhone = phone
}

const processCheckout = async () => {
  if (!checkoutForm.value) return

  try {
    await checkoutForm.value.validate()

    if (cartStore.isEmpty) {
      ElMessage.error('Your cart is empty')
      return
    }

    processingPayment.value = true

    // Create order
    const orderData = {
      cart_id: cartStore.cart?.id,
      shipping_address: `${checkoutData.fullName}, ${checkoutData.address}, ${checkoutData.city}`
    }

    const orderResponse = await orderStore.createOrder(orderData)

    // Process M-Pesa payment
    try {
      const paymentData = {
        order_id: orderResponse.data.id,
        phone_number: checkoutData.mpesaPhone,
        amount: totalAmount.value
      }

      const paymentResponse = await orderStore.processMpesaPayment(paymentData)

      ElMessage.success('STK Push sent to your phone. Please enter your M-Pesa PIN to complete the payment.')

      // Redirect to orders page after successful payment initiation
      setTimeout(() => {
        router.push('/orders')
      }, 2000)

    } catch (paymentError) {
      console.error('Payment error:', paymentError)
      ElMessage.warning('Order created but payment failed. You can retry payment from your orders page.')
      router.push('/orders')
    }

  } catch (error) {
    console.error('Checkout error:', error)
    ElMessage.error('Failed to process checkout. Please try again.')
  } finally {
    processingPayment.value = false
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(async () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('Please login to checkout')
    router.push('/login')
    return
  }

  if (cartStore.isEmpty) {
    ElMessage.info('Your cart is empty')
    router.push('/products')
    return
  }

  // Pre-fill form with user data if available
  if (authStore.user) {
    checkoutData.fullName = authStore.user.name || ''
    checkoutData.email = authStore.user.email || ''
  }

  // Load cart if not already loaded
  if (!cartStore.cart) {
    try {
      await cartStore.fetchCart()
    } catch (error) {
      ElMessage.error('Failed to load cart')
      router.push('/cart')
    }
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