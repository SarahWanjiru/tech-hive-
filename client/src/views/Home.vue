<template>
  <div class="home-page">
    <!-- Header Section -->
    <header class="hero-section">
      <div class="container">
        <div class="hero-content">
          <div class="hero-text">
            <h1 class="hero-title">Welcome to Mini-Ecommerce</h1>
            <p class="hero-subtitle">
              Discover amazing products at great prices. Shop with confidence and enjoy fast, secure delivery.
            </p>
            <div class="hero-actions">
              <el-button
                type="primary"
                size="large"
                @click="$router.push('/products')"
              >
                Shop Now
              </el-button>
              <el-button
                size="large"
                @click="$router.push('/register')"
              >
                Join Us
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Features Section -->
    <section class="features-section">
      <div class="container">
        <div class="section-header">
          <h2>Why Choose Us?</h2>
          <p class="section-subtitle">Experience the best in online shopping</p>
        </div>

        <div class="features-grid">
          <div class="feature-card">
            <div class="feature-icon">
              🚚
            </div>
            <h3>Fast Delivery</h3>
            <p>Get your orders delivered quickly and safely to your doorstep.</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              🔒
            </div>
            <h3>Secure Payments</h3>
            <p>Your payment information is protected with industry-standard security.</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              🎧
            </div>
            <h3>24/7 Support</h3>
            <p>Our customer service team is always ready to help you.</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              💰
            </div>
            <h3>Best Prices</h3>
            <p>Competitive prices and great deals on all products.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Featured Products Section -->
    <section class="featured-products-section">
      <div class="container">
        <div class="section-header">
          <h2>Featured Products</h2>
          <p class="section-subtitle">Check out our most popular items</p>
        </div>

        <div v-if="productStore.loading" class="text-center">
          <el-icon class="loading"><Loading /></el-icon>
          <p>Loading products...</p>
        </div>

        <div v-else-if="productStore.error" class="text-center">
          <p class="text-danger">{{ productStore.error }}</p>
          <el-button @click="loadFeaturedProducts">Try Again</el-button>
        </div>

        <div v-else class="product-grid">
          <div
            v-for="product in featuredProducts"
            :key="product.id"
            class="product-card"
            @click="$router.push(`/products/${product.id}`)"
          >
            <div class="product-image">
              <img
                :src="product.image_url || '/placeholder-product.jpg'"
                :alt="product.name"
                @error="handleImageError"
              />
            </div>
            <div class="product-info">
              <h4 class="product-name">{{ product.name }}</h4>
              <p class="product-description">{{ product.description }}</p>
              <div class="product-price">
                <span class="price">${{ product.price }}</span>
                <span v-if="product.stock <= 5" class="stock-warning">
                  Only {{ product.stock }} left!
                </span>
              </div>
              <el-button
                type="primary"
                class="add-to-cart-btn"
                @click.stop="addToCart(product)"
              >
                Add to Cart
              </el-button>
            </div>
          </div>
        </div>

        <div v-if="!productStore.loading && featuredProducts.length > 0" class="text-center">
          <el-button
            type="primary"
            size="large"
            @click="$router.push('/products')"
          >
            View All Products
          </el-button>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useProductStore } from '../stores/product.store.js'
import { useCartStore } from '../stores/cart.store.js'
import { useAuthStore } from '../stores/auth.store.js'
import { ElMessage } from 'element-plus'
import { Loading } from '@element-plus/icons-vue'

const productStore = useProductStore()
const cartStore = useCartStore()
const authStore = useAuthStore()

const featuredProducts = computed(() => {
  return productStore.products.slice(0, 8) // Show first 8 products as featured
})

const loadFeaturedProducts = async () => {
  try {
    await productStore.fetchProducts({ limit: 8 })
  } catch (error) {
    console.error('Failed to load featured products:', error)
  }
}

const addToCart = async (product) => {
  if (!authStore.isAuthenticated) {
    ElMessage.warning('Please login to add items to cart')
    return
  }

  try {
    await cartStore.addToCart(product.id, 1)
    ElMessage.success(`${product.name} added to cart!`)
  } catch (error) {
    ElMessage.error('Failed to add item to cart')
    console.error('Add to cart error:', error)
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.jpg'
}

onMounted(() => {
  loadFeaturedProducts()
})
</script>

<style scoped>
.home-page {
  min-height: 100vh;
}

/* Hero Section */
.hero-section {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: var(--white);
  padding: 4rem 0;
  text-align: center;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero-title {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 1rem;
}

.hero-subtitle {
  font-size: 1.25rem;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

/* Features Section */
.features-section {
  padding: 4rem 0;
  background: var(--white);
}

.section-header {
  text-align: center;
  margin-bottom: 3rem;
}

.section-header h2 {
  font-size: 2.5rem;
  color: var(--gray-800);
  margin-bottom: 1rem;
}

.section-subtitle {
  font-size: 1.125rem;
  color: var(--gray-600);
  margin-bottom: 0;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.feature-card {
  text-align: center;
  padding: 2rem;
  border-radius: var(--border-radius-lg);
  background: var(--gray-50);
  transition: transform var(--transition-fast);
}

.feature-card:hover {
  transform: translateY(-5px);
}

.feature-icon {
  width: 60px;
  height: 60px;
  margin: 0 auto 1rem;
  background: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--white);
  font-size: 24px;
}

.feature-card h3 {
  font-size: 1.25rem;
  margin-bottom: 1rem;
  color: var(--gray-800);
}

/* Featured Products Section */
.featured-products-section {
  padding: 4rem 0;
  background: var(--gray-50);
}

.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.product-card {
  background: var(--white);
  border-radius: var(--border-radius-lg);
  overflow: hidden;
  box-shadow: var(--shadow-sm);
  transition: all var(--transition-fast);
  cursor: pointer;
}

.product-card:hover {
  transform: translateY(-5px);
  box-shadow: var(--shadow-lg);
}

.product-image {
  width: 100%;
  height: 200px;
  overflow: hidden;
}

.product-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform var(--transition-fast);
}

.product-card:hover .product-image img {
  transform: scale(1.05);
}

.product-info {
  padding: 1.5rem;
}

.product-name {
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--gray-800);
  margin-bottom: 0.5rem;
}

.product-description {
  color: var(--gray-600);
  font-size: 0.875rem;
  margin-bottom: 1rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.product-price {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.price {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--primary-color);
}

.stock-warning {
  font-size: 0.75rem;
  color: var(--warning-color);
  background: var(--warning-color);
  color: var(--white);
  padding: 0.25rem 0.5rem;
  border-radius: var(--border-radius-sm);
}

.add-to-cart-btn {
  width: 100%;
}

/* Responsive Design */
@media (max-width: 768px) {
  .hero-title {
    font-size: 2rem;
  }

  .hero-actions {
    flex-direction: column;
    align-items: center;
  }

  .features-grid {
    grid-template-columns: 1fr;
  }

  .product-grid {
    grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  }
}
</style>