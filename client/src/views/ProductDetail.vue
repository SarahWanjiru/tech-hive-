<template>
  <div class="product-detail-page">
    <div class="container">
      <!-- Loading State -->
      <div v-if="productStore.loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading product details...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="productStore.error" class="error-container">
        <p>{{ productStore.error }}</p>
        <el-button type="primary" @click="loadProduct">Try Again</el-button>
      </div>

      <!-- Product Details -->
      <div v-else-if="productStore.currentProduct" class="product-details">
        <div class="product-gallery">
          <div class="main-image">
            <img
              :src="productStore.currentProduct.image_url || '/placeholder-product.jpg'"
              :alt="productStore.currentProduct.name"
              @error="handleImageError"
            />
          </div>
        </div>

        <div class="product-info">
          <h1 class="product-title">{{ productStore.currentProduct.name }}</h1>
          <p class="product-description">{{ productStore.currentProduct.description }}</p>

          <div class="product-meta">
            <div class="product-price">
              <span class="price">${{ productStore.currentProduct.price }}</span>
              <span class="stock">Stock: {{ productStore.currentProduct.stock }}</span>
            </div>
          </div>

          <div class="product-actions">
            <el-input-number
              v-model="quantity"
              :min="1"
              :max="productStore.currentProduct.stock"
              size="large"
            />

            <el-button
              type="primary"
              size="large"
              :disabled="productStore.currentProduct.stock === 0"
              @click="addToCart"
            >
              {{ productStore.currentProduct.stock === 0 ? 'Out of Stock' : 'Add to Cart' }}
            </el-button>
          </div>
        </div>
      </div>

      <!-- Product Not Found -->
      <div v-else class="not-found">
        <el-empty description="Product not found">
          <el-button type="primary" @click="$router.push('/products')">
            Browse Products
          </el-button>
        </el-empty>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useProductStore } from '../stores/product.store.js'
import { useCartStore } from '../stores/cart.store.js'
import { useAuthStore } from '../stores/auth.store.js'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const productStore = useProductStore()
const cartStore = useCartStore()
const authStore = useAuthStore()

const quantity = ref(1)

const loadProduct = async () => {
  try {
    await productStore.fetchProductById(route.params.id)
  } catch (error) {
    console.error('Failed to load product:', error)
  }
}

const addToCart = async () => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('Please login to add items to cart')
    return
  }

  try {
    await cartStore.addToCart(route.params.id, quantity.value)
    ElMessage.success('Product added to cart!')
  } catch (error) {
    ElMessage.error('Failed to add product to cart')
    console.error('Add to cart error:', error)
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(() => {
  loadProduct()
})
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