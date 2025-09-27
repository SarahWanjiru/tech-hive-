<template>
  <div class="products-page">
    <div class="container">
      <!-- Page Header -->
      <div class="page-header">
        <h1>Our Products</h1>
        <p>Discover amazing products at great prices</p>
      </div>

      <!-- Search and Filters -->
      <div class="filters-section">
        <el-card>
          <el-form :inline="true" :model="filters" @submit.prevent="handleSearch">
            <el-form-item label="Search">
              <el-input
                v-model="filters.search"
                placeholder="Search products..."
                clearable
                @input="handleSearchInput"
              />
            </el-form-item>

            <el-form-item label="Category">
              <el-select
                v-model="filters.category"
                placeholder="All Categories"
                clearable
                @change="handleSearch"
              >
                <el-option
                  v-for="category in availableCategories"
                  :key="category"
                  :label="category"
                  :value="category"
                />
              </el-select>
            </el-form-item>

            <el-form-item label="Min Price">
              <el-input-number
                v-model="filters.minPrice"
                :min="0"
                :precision="2"
                placeholder="0"
                @change="handleSearch"
              />
            </el-form-item>

            <el-form-item label="Max Price">
              <el-input-number
                v-model="filters.maxPrice"
                :min="0"
                :precision="2"
                placeholder="1000"
                @change="handleSearch"
              />
            </el-form-item>

            <el-form-item>
              <el-checkbox
                v-model="filters.inStock"
                @change="handleSearch"
              >
                In Stock Only
              </el-checkbox>
            </el-form-item>

            <el-form-item>
              <el-button type="primary" @click="handleSearch">
                <el-icon><Search /></el-icon>
                Search
              </el-button>
              <el-button @click="clearFilters">
                Clear
              </el-button>
            </el-form-item>
          </el-form>
        </el-card>
      </div>

      <!-- Loading State -->
      <div v-if="productStore.loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading products...</p>
      </div>

      <!-- Error State -->
      <div v-else-if="productStore.error" class="error-container">
        <p>{{ productStore.error }}</p>
        <el-button type="primary" @click="loadProducts">Try Again</el-button>
      </div>

      <!-- Products Grid -->
      <div v-else>
        <!-- Sort Controls -->
        <div class="sort-controls">
          <el-radio-group v-model="filters.sortBy" @change="handleSortChange">
            <el-radio label="name">Name</el-radio>
            <el-radio label="price">Price</el-radio>
            <el-radio label="stock">Stock</el-radio>
          </el-radio-group>

          <el-radio-group v-model="filters.sortOrder" @change="handleSortChange">
            <el-radio label="asc">Ascending</el-radio>
            <el-radio label="desc">Descending</el-radio>
          </el-radio-group>
        </div>

        <!-- Products Grid -->
        <div class="product-grid">
          <div
            v-for="product in productStore.filteredProducts"
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
              <div v-if="product.stock === 0" class="out-of-stock-badge">
                Out of Stock
              </div>
            </div>

            <div class="product-info">
              <h4 class="product-name">{{ product.name }}</h4>
              <p class="product-description">{{ product.description }}</p>

              <div class="product-meta">
                <span class="product-price">${{ product.price }}</span>
                <span class="product-stock">Stock: {{ product.stock }}</span>
              </div>

              <div class="product-actions">
                <el-button
                  type="primary"
                  :disabled="product.stock === 0"
                  @click.stop="addToCart(product)"
                >
                  {{ product.stock === 0 ? 'Out of Stock' : 'Add to Cart' }}
                </el-button>
              </div>
            </div>
          </div>
        </div>

        <!-- No Products Found -->
        <div v-if="productStore.filteredProducts.length === 0" class="no-products">
          <el-empty description="No products found">
            <el-button type="primary" @click="clearFilters">
              Clear Filters
            </el-button>
          </el-empty>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { useProductStore } from '../stores/product.store.ts'
import { useCartStore } from '../stores/cart.store.ts'
import { useAuthStore } from '../stores/auth.store.ts'
import { ElMessage } from 'element-plus'
import { Search, Loading } from '@element-plus/icons-vue'

const productStore = useProductStore()
const cartStore = useCartStore()
const authStore = useAuthStore()

const filters = reactive({
  search: '',
  category: '',
  minPrice: 0,
  maxPrice: 999999,
  inStock: true,
  sortBy: 'name',
  sortOrder: 'asc'
})

// Watch for filter changes and update store
watch(
  () => filters,
  (newFilters) => {
    productStore.updateFilters(newFilters)
  },
  { deep: true }
)

const availableCategories = computed(() => {
  return productStore.availableCategories
})

const loadProducts = async () => {
  try {
    await productStore.fetchProducts()
  } catch (error) {
    console.error('Failed to load products:', error)
  }
}

const handleSearchInput = () => {
  // Debounce search input
  clearTimeout(this.searchTimeout)
  this.searchTimeout = setTimeout(() => {
    handleSearch()
  }, 500)
}

const handleSearch = () => {
  productStore.updateFilters(filters)
}

const handleSortChange = () => {
  productStore.updateFilters(filters)
}

const clearFilters = () => {
  filters.search = ''
  filters.category = ''
  filters.minPrice = 0
  filters.maxPrice = 999999
  filters.inStock = true
  filters.sortBy = 'name'
  filters.sortOrder = 'asc'

  productStore.clearFilters()
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
  loadProducts()
})
</script>

<style scoped>
.products-page {
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

/* Filters Section */
.filters-section {
  margin-bottom: 2rem;
}

.sort-controls {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding: 1rem;
  background: var(--white);
  border-radius: var(--border-radius-lg);
  box-shadow: var(--shadow-sm);
}

/* Product Grid */
.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
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
  position: relative;
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

.out-of-stock-badge {
  position: absolute;
  top: 10px;
  right: 10px;
  background: var(--danger-color);
  color: var(--white);
  padding: 0.25rem 0.75rem;
  border-radius: var(--border-radius-sm);
  font-size: 0.75rem;
  font-weight: 600;
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

.product-meta {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.product-price {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--primary-color);
}

.product-stock {
  font-size: 0.875rem;
  color: var(--gray-600);
}

.product-actions {
  text-align: center;
}

/* Loading and Error States */
.loading-container,
.error-container {
  text-align: center;
  padding: 3rem;
}

.no-products {
  padding: 3rem 0;
}

/* Responsive Design */
@media (max-width: 768px) {
  .sort-controls {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .product-grid {
    grid-template-columns: 1fr;
  }
}
</style>