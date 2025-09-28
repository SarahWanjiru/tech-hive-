<template>
  <div class="products-page">
    <div class="container">
      <!-- Page Header -->
      <div class="page-header">
        <h1>Our Products</h1>
        <p>Discover amazing products at great prices</p>
      </div>

      <!-- Main Content Layout -->
      <div class="products-layout">
        <!-- Filter Sidebar -->
        <aside class="filter-sidebar-container">
          <FilterSidebar
            :available-categories="availableCategories"
            :available-brands="availableBrands"
            :total-results="filteredProductsCount"
            @filter-change="handleFilterChange"
            @apply-filters="handleApplyFilters"
          />
        </aside>

        <!-- Products Content -->
        <main class="products-content">
          <!-- Search and Sort Controls -->
          <div class="content-header">
            <div class="search-summary">
              <AdvancedSearch
                v-model:query="searchQuery"
                @search="handleAdvancedProductSearch"
                placeholder="Search in products..."
              />
              <div class="results-summary">
                <span class="results-count">
                  {{ filteredProductsCount }} products found
                </span>
                <div v-if="hasActiveFilters" class="active-filters-summary">
                  <span>Filtered by:</span>
                  <div class="filter-tags">
                    <el-tag
                      v-for="filter in activeFilterTags"
                      :key="filter.key"
                      size="small"
                      closable
                      @close="removeFilter(filter.key)"
                    >
                      {{ filter.label }}
                    </el-tag>
                  </div>
                </div>
              </div>
            </div>

            <!-- Sort and View Controls -->
            <div class="controls-section">
              <div class="sort-controls">
                <el-select v-model="sortBy" @change="handleSortChange" placeholder="Sort by">
                  <el-option label="Featured" value="featured" />
                  <el-option label="Price: Low to High" value="price_asc" />
                  <el-option label="Price: High to Low" value="price_desc" />
                  <el-option label="Newest First" value="newest" />
                  <el-option label="Customer Rating" value="rating" />
                  <el-option label="Most Popular" value="popular" />
                </el-select>
              </div>

              <div class="view-controls">
                <el-radio-group v-model="viewMode" @change="handleViewChange" size="large">
                  <el-radio-button label="grid">
                    <el-icon><Grid /></el-icon>
                  </el-radio-button>
                  <el-radio-button label="list">
                    <el-icon><List /></el-icon>
                  </el-radio-button>
                </el-radio-group>
              </div>
            </div>
          </div>

          <!-- Loading State -->
          <div v-if="productStore.loading" class="loading-container">
            <div class="loading-skeleton">
              <div v-for="i in 6" :key="i" class="skeleton-card">
                <el-skeleton animated>
                  <template #template>
                    <div class="skeleton-image"></div>
                    <div class="skeleton-content">
                      <el-skeleton-item variant="text" style="width: 60%" />
                      <el-skeleton-item variant="text" style="width: 40%" />
                      <el-skeleton-item variant="rect" style="width: 100%; height: 40px" />
                    </div>
                  </template>
                </el-skeleton>
              </div>
            </div>
          </div>

          <!-- Error State -->
          <div v-else-if="productStore.error" class="error-container">
            <el-empty description="Failed to load products">
              <el-button type="primary" @click="loadProducts">Try Again</el-button>
            </el-empty>
          </div>

          <!-- Products Content -->
          <div v-else>
            <!-- Products Grid/List -->
            <div :class="['products-display', `view-${viewMode}`]">
              <div
                v-for="product in displayedProducts"
                :key="product.id"
                :class="['product-item', `view-${viewMode}`]"
                @click="navigateToProduct(product.id)"
              >
                <!-- Product Card View -->
                <div v-if="viewMode === 'grid'" class="product-card">
                  <div class="product-image-section">
                    <div class="product-image-container">
                      <img
                        :src="product.image_url || '/placeholder-product.svg'"
                        :alt="product.name"
                        class="product-image"
                        loading="lazy"
                        @error="handleImageError"
                      />

                      <!-- Image Overlay Actions -->
                      <div class="image-overlay">
                        <div class="overlay-actions">
                          <el-button
                            circle
                            size="small"
                            class="overlay-btn wishlist-btn"
                            @click.stop="toggleWishlist(product)"
                            :class="{ 'active': isInWishlist(product.id) }"
                          >
                            <el-icon><Heart /></el-icon>
                          </el-button>
                          <el-button
                            circle
                            size="small"
                            class="overlay-btn quick-view-btn"
                            @click.stop="quickView(product)"
                          >
                            <el-icon><View /></el-icon>
                          </el-button>
                        </div>
                      </div>

                      <!-- Stock Status Badge -->
                      <div v-if="product.stock === 0" class="stock-badge out-of-stock">
                        <el-icon><Close /></el-icon>
                        Out of Stock
                      </div>
                      <div v-else-if="product.stock <= 5" class="stock-badge low-stock">
                        <el-icon><Warning /></el-icon>
                        Only {{ product.stock }} left
                      </div>
                      <div v-else class="stock-badge in-stock">
                        <el-icon><Check /></el-icon>
                        In Stock
                      </div>

                      <!-- Featured Badge -->
                      <div v-if="product.featured" class="featured-badge">
                        <el-icon><Star /></el-icon>
                        Featured
                      </div>
                    </div>
                  </div>

                  <div class="product-info-section">
                    <div class="product-category">
                      <span class="category-tag">{{ product.category || 'General' }}</span>
                    </div>

                    <h3 class="product-name">{{ product.name }}</h3>

                    <div class="product-rating">
                      <el-rate
                        :model-value="4.5"
                        disabled
                        show-score
                        text-color="#ff9900"
                        score-template="{value}/5"
                        size="small"
                      />
                      <span class="rating-count">(24 reviews)</span>
                    </div>

                    <p class="product-description">{{ product.description }}</p>

                    <div class="price-section">
                      <div class="current-price">
                        <span class="currency">KES</span>
                        <span class="amount">{{ formatPrice(product.price) }}</span>
                      </div>
                      <div v-if="product.originalPrice" class="original-price">
                        <span class="currency">KES</span>
                        <span class="amount">{{ formatPrice(product.originalPrice) }}</span>
                      </div>
                      <div v-if="product.discount" class="discount-badge">
                        -{{ product.discount }}% OFF
                      </div>
                    </div>

                    <div class="product-actions">
                      <el-button
                        type="primary"
                        size="large"
                        class="add-to-cart-btn"
                        :disabled="product.stock === 0"
                        @click.stop="addToCart(product)"
                      >
                        <el-icon><ShoppingCart /></el-icon>
                        {{ product.stock === 0 ? 'Out of Stock' : 'Add to Cart' }}
                      </el-button>

                      <el-button
                        size="large"
                        class="buy-now-btn"
                        :disabled="product.stock === 0"
                        @click.stop="buyNow(product)"
                      >
                        Buy Now
                      </el-button>
                    </div>
                  </div>
                </div>

                <!-- Product List View -->
                <div v-else class="product-list-item">
                  <div class="list-image">
                    <img
                      :src="product.image_url || '/placeholder-product.svg'"
                      :alt="product.name"
                      @error="handleImageError"
                    />
                  </div>
                  <div class="list-content">
                    <div class="list-header">
                      <h4 class="list-name">{{ product.name }}</h4>
                      <div class="list-price">
                        <span class="currency">KES</span>
                        <span class="amount">{{ formatPrice(product.price) }}</span>
                      </div>
                    </div>
                    <p class="list-description">{{ product.description }}</p>
                    <div class="list-meta">
                      <span class="category">{{ product.category || 'General' }}</span>
                      <span class="stock">Stock: {{ product.stock }}</span>
                      <div class="rating">
                        <el-rate :model-value="4.5" disabled size="small" />
                        <span>(24)</span>
                      </div>
                    </div>
                    <div class="list-actions">
                      <el-button type="primary" @click.stop="addToCart(product)">
                        Add to Cart
                      </el-button>
                      <el-button @click.stop="buyNow(product)">Buy Now</el-button>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            <!-- No Products Found -->
            <div v-if="displayedProducts.length === 0" class="no-products">
              <el-empty description="No products found matching your criteria">
                <el-button type="primary" @click="clearAllFilters">
                  Clear All Filters
                </el-button>
              </el-empty>
            </div>

            <!-- Load More / Pagination -->
            <div v-if="hasMoreProducts" class="load-more-section">
              <el-button
                :loading="loadingMore"
                size="large"
                @click="loadMoreProducts"
                class="load-more-btn"
              >
                Load More Products
              </el-button>
            </div>
          </div>
        </main>
      </div>
    </div>
  </div>
</template>

<script setup>
 import { ref, reactive, computed, onMounted, watch, nextTick } from 'vue'
 import { useRouter } from 'vue-router'
 import { useProductStore } from '../stores/product.store.ts'
 import { useCartStore } from '../stores/cart.store.ts'
 import { useAuthStore } from '../stores/auth.store.ts'
 import { authService } from '../services/auth.service'
 import { ElMessage } from 'element-plus'
 import AdvancedSearch from '../components/common/AdvancedSearch.vue'
 import FilterSidebar from '../components/common/FilterSidebar.vue'
 import { useTouchGestures, useSwipeNavigation } from '../composables/useTouchGestures.js'
 import {
   Search,
   Loading,
   Star,
   View,
   Close,
   Warning,
   Check,
   Box,
   Timer,
   ShoppingCart,
   Share,
   Grid,
   List
 } from '@element-plus/icons-vue'

const router = useRouter()
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

// New reactive data for enhanced features
const wishlist = ref(new Set())
const showQuickView = ref(false)
const selectedProduct = ref(null)

// New reactive data for enhanced layout
const searchQuery = ref('')
const sortBy = ref('featured')
const viewMode = ref('grid')
const loadingMore = ref(false)
const hasMoreProducts = ref(true)

// Mock data for brands (in a real app, this would come from the API)
const availableBrands = ref([
  { id: 1, name: 'Samsung', count: 15 },
  { id: 2, name: 'Apple', count: 8 },
  { id: 3, name: 'Huawei', count: 12 },
  { id: 4, name: 'OnePlus', count: 6 },
  { id: 5, name: 'Google', count: 4 }
])

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

// Enhanced computed properties
const filteredProductsCount = computed(() => {
  return productStore.filteredProducts?.length || 0
})

const displayedProducts = computed(() => {
  return productStore.filteredProducts || []
})

const hasActiveFilters = computed(() => {
  return filters.search ||
         filters.category ||
         filters.minPrice > 0 ||
         filters.maxPrice < 999999 ||
         !filters.inStock
})

const activeFilterTags = computed(() => {
  const tags = []

  if (filters.search) {
    tags.push({ key: 'search', label: `Search: ${filters.search}` })
  }

  if (filters.category) {
    tags.push({ key: 'category', label: `Category: ${filters.category}` })
  }

  if (filters.minPrice > 0 || filters.maxPrice < 999999) {
    tags.push({
      key: 'price',
      label: `Price: ${formatPrice(filters.minPrice)} - ${formatPrice(filters.maxPrice)}`
    })
  }

  if (!filters.inStock) {
    tags.push({ key: 'stock', label: 'Include out of stock' })
  }

  return tags
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
    // Check if user has account data (has registered before)
    const hasAccount = authService.getCurrentUser() !== null

    if (hasAccount) {
      // User has account but not logged in - redirect to login
      ElMessage.warning('Please login to add items to cart')
      router.push('/login')
    } else {
      // User doesn't have account - redirect to signup
      ElMessage.warning('Please create an account to add items to cart')
      router.push('/register')
    }
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
  event.target.src = '/placeholder-product.svg'
}

// Enhanced product card methods
const formatPrice = (price) => {
  return new Intl.NumberFormat('en-KE', {
    style: 'currency',
    currency: 'KES',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(price).replace('KES', '')
}

const isInWishlist = (productId) => {
  return wishlist.value.has(productId)
}

const toggleWishlist = (product) => {
  if (wishlist.value.has(product.id)) {
    wishlist.value.delete(product.id)
    ElMessage.success(`${product.name} removed from wishlist`)
  } else {
    wishlist.value.add(product.id)
    ElMessage.success(`${product.name} added to wishlist`)
  }
}

const quickView = (product) => {
  selectedProduct.value = product
  showQuickView.value = true
  // In a real app, this would open a modal with product details
  ElMessage.info(`Quick view: ${product.name}`)
}

const buyNow = async (product) => {
  if (!authStore.isAuthenticated) {
    // Check if user has account data (has registered before)
    const hasAccount = authService.getCurrentUser() !== null

    if (hasAccount) {
      // User has account but not logged in - redirect to login
      ElMessage.warning('Please login to purchase items')
      router.push('/login')
    } else {
      // User doesn't have account - redirect to signup
      ElMessage.warning('Please create an account to purchase items')
      router.push('/register')
    }
    return
  }

  try {
    // Add to cart and redirect to checkout
    await cartStore.addToCart(product.id, 1)
    ElMessage.success(`${product.name} added to cart!`)
    router.push('/checkout')
  } catch (error) {
    ElMessage.error('Failed to process purchase')
    console.error('Buy now error:', error)
  }
}

const compareProduct = (product) => {
  ElMessage.info(`Compare feature coming soon for: ${product.name}`)
}

const shareProduct = (product) => {
  if (navigator.share) {
    navigator.share({
      title: product.name,
      text: product.description,
      url: window.location.href + `/products/${product.id}`,
    })
  } else {
    // Fallback: copy to clipboard
    navigator.clipboard.writeText(window.location.href + `/products/${product.id}`)
    ElMessage.success('Product link copied to clipboard!')
  }
}

// Enhanced methods for new functionality
const handleAdvancedProductSearch = (query) => {
  searchQuery.value = query
  filters.search = query
  handleSearch()
}

const handleFilterChange = (newFilters) => {
  // Update local filters with sidebar filters
  Object.assign(filters, newFilters)
  handleSearch()
}

const handleApplyFilters = (appliedFilters) => {
  Object.assign(filters, appliedFilters)
  handleSearch()
}

const handleSortChange = (newSortBy) => {
  sortBy.value = newSortBy
  // Update filters with new sort
  filters.sortBy = newSortBy
  handleSearch()
}

const handleViewChange = (newViewMode) => {
  viewMode.value = newViewMode
}

const navigateToProduct = (productId) => {
  router.push(`/products/${productId}`)
}

const removeFilter = (filterKey) => {
  switch (filterKey) {
    case 'search':
      searchQuery.value = ''
      filters.search = ''
      break
    case 'category':
      filters.category = ''
      break
    case 'price':
      filters.minPrice = 0
      filters.maxPrice = 999999
      break
    case 'stock':
      filters.inStock = true
      break
  }
  handleSearch()
}

const clearAllFilters = () => {
  clearFilters()
  searchQuery.value = ''
  sortBy.value = 'featured'
  viewMode.value = 'grid'
}

const loadMoreProducts = async () => {
  loadingMore.value = true
  try {
    // In a real app, this would load more products from the API
    await new Promise(resolve => setTimeout(resolve, 1000)) // Simulate API call
    hasMoreProducts.value = false // For demo purposes
  } finally {
    loadingMore.value = false
  }
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

/* Products Layout */
.products-layout {
  display: grid;
  grid-template-columns: 280px 1fr;
  gap: 2rem;
  margin-top: 2rem;
}

.filter-sidebar-container {
  position: sticky;
  top: 2rem;
  height: fit-content;
}

.products-content {
  min-width: 0; /* Prevent grid overflow */
}

/* Content Header */
.content-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: 2rem;
  margin-bottom: 2rem;
  padding: 1.5rem;
  background: var(--white);
  border-radius: 16px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--gray-100);
}

.search-summary {
  flex: 1;
}

.results-summary {
  margin-top: 1rem;
}

.results-count {
  display: block;
  font-size: 1.125rem;
  font-weight: 600;
  color: var(--gray-800);
  margin-bottom: 0.5rem;
}

.active-filters-summary {
  display: flex;
  align-items: center;
  gap: 1rem;
  flex-wrap: wrap;
}

.active-filters-summary span:first-child {
  color: var(--gray-600);
  font-size: 0.875rem;
}

.filter-tags {
  display: flex;
  gap: 0.5rem;
  flex-wrap: wrap;
}

/* Controls Section */
.controls-section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  align-items: flex-end;
}

.sort-controls {
  min-width: 200px;
}

.view-controls {
  display: flex;
  border: 1px solid var(--gray-200);
  border-radius: 8px;
  overflow: hidden;
}

/* Products Display */
.products-display {
  display: grid;
  gap: 2rem;
}

.view-grid {
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
}

.view-list {
  grid-template-columns: 1fr;
}

/* Product List Item */
.product-list-item {
  display: grid;
  grid-template-columns: 200px 1fr;
  gap: 1.5rem;
  background: var(--white);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  border: 1px solid var(--gray-100);
  transition: all var(--transition-normal);
}

.product-list-item:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.12);
}

.list-image {
  width: 100%;
  height: 200px;
  border-radius: 12px;
  overflow: hidden;
}

.list-image img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}

.list-content {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.list-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 1rem;
}

.list-name {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--gray-800);
  margin: 0;
  line-height: 1.3;
}

.list-price {
  display: flex;
  align-items: baseline;
  gap: 0.25rem;
}

.list-price .amount {
  font-size: 1.5rem;
  font-weight: 800;
  color: var(--primary-color);
}

.list-description {
  color: var(--gray-600);
  line-height: 1.6;
  margin-bottom: 1rem;
}

.list-meta {
  display: flex;
  gap: 1.5rem;
  margin-bottom: 1.5rem;
  font-size: 0.875rem;
  color: var(--gray-600);
}

.list-meta .category {
  background: var(--primary-color-light);
  color: var(--primary-color);
  padding: 0.25rem 0.75rem;
  border-radius: 12px;
  font-weight: 600;
}

.list-meta .rating {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.list-actions {
  display: flex;
  gap: 1rem;
}

/* Load More Section */
.load-more-section {
  text-align: center;
  margin-top: 3rem;
  padding-top: 2rem;
  border-top: 1px solid var(--gray-200);
}

.load-more-btn {
  min-width: 200px;
  height: 48px;
  border-radius: 8px;
  font-weight: 600;
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

/* Enhanced Product Grid */
.product-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 2rem;
  padding: 1rem 0;
}

.product-card {
  background: var(--white);
  border-radius: 16px;
  overflow: hidden;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.08);
  transition: all var(--transition-normal);
  cursor: pointer;
  position: relative;
  border: 1px solid var(--gray-100);
}

.product-card:hover {
  transform: translateY(-8px);
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15);
  border-color: var(--primary-color);
}

/* Product Image Section */
.product-image-section {
  position: relative;
  overflow: hidden;
}

.product-image-container {
  position: relative;
  width: 100%;
  height: 240px;
  background: var(--gray-50);
}

.product-image {
  width: 100%;
  height: 100%;
  object-fit: cover;
  transition: transform var(--transition-normal);
}

.product-card:hover .product-image {
  transform: scale(1.1);
}

/* Image Overlay */
.image-overlay {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.3);
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity var(--transition-fast);
}

.product-card:hover .image-overlay {
  opacity: 1;
}

.overlay-actions {
  display: flex;
  gap: 0.5rem;
}

.overlay-btn {
  background: var(--white);
  border: none;
  color: var(--gray-700);
  transition: all var(--transition-fast);
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.overlay-btn:hover {
  background: var(--primary-color);
  color: var(--white);
  transform: scale(1.1);
}

.wishlist-btn.active {
  background: var(--danger-color);
  color: var(--white);
}

/* Stock Status Badges */
.stock-badge {
  position: absolute;
  top: 12px;
  left: 12px;
  display: flex;
  align-items: center;
  gap: 0.25rem;
  padding: 0.5rem 0.75rem;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 600;
  z-index: 2;
}

.stock-badge.out-of-stock {
  background: var(--danger-color);
  color: var(--white);
}

.stock-badge.low-stock {
  background: var(--warning-color);
  color: var(--white);
}

.stock-badge.in-stock {
  background: var(--success-color);
  color: var(--white);
}

/* Featured Badge */
.featured-badge {
  position: absolute;
  top: 12px;
  right: 12px;
  background: linear-gradient(135deg, #ff6b6b, #ee5a24);
  color: var(--white);
  padding: 0.5rem 0.75rem;
  border-radius: 20px;
  font-size: 0.75rem;
  font-weight: 600;
  display: flex;
  align-items: center;
  gap: 0.25rem;
  z-index: 2;
}

/* Product Info Section */
.product-info-section {
  padding: 1.5rem;
}

.product-category {
  margin-bottom: 0.75rem;
}

.category-tag {
  background: var(--primary-color-light);
  color: var(--primary-color);
  padding: 0.25rem 0.75rem;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.product-name {
  font-size: 1.25rem;
  font-weight: 700;
  color: var(--gray-800);
  margin-bottom: 0.75rem;
  line-height: 1.3;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Product Rating */
.product-rating {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  margin-bottom: 0.75rem;
}

.rating-count {
  font-size: 0.875rem;
  color: var(--gray-500);
}

/* Product Description */
.product-description {
  color: var(--gray-600);
  font-size: 0.875rem;
  line-height: 1.5;
  margin-bottom: 1rem;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

/* Price Section */
.price-section {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin-bottom: 1rem;
  flex-wrap: wrap;
}

.current-price {
  display: flex;
  align-items: baseline;
  gap: 0.25rem;
}

.currency {
  font-size: 0.875rem;
  color: var(--gray-600);
  font-weight: 500;
}

.amount {
  font-size: 1.5rem;
  font-weight: 800;
  color: var(--primary-color);
}

.original-price {
  display: flex;
  align-items: baseline;
  gap: 0.25rem;
  text-decoration: line-through;
  color: var(--gray-500);
}

.original-price .amount {
  font-size: 1.125rem;
  font-weight: 500;
}

.discount-badge {
  background: var(--danger-color);
  color: var(--white);
  padding: 0.25rem 0.5rem;
  border-radius: 6px;
  font-size: 0.75rem;
  font-weight: 700;
}

/* Product Meta */
.product-meta {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
  padding: 1rem;
  background: var(--gray-50);
  border-radius: 8px;
}

.meta-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.875rem;
  color: var(--gray-600);
}

/* Product Actions */
.product-actions {
  display: flex;
  gap: 0.75rem;
  margin-bottom: 1rem;
}

.add-to-cart-btn {
  flex: 2;
  height: 48px;
  border-radius: 8px;
  font-weight: 600;
  transition: all var(--transition-fast);
}

.buy-now-btn {
  flex: 1;
  height: 48px;
  border-radius: 8px;
  font-weight: 600;
  background: var(--gray-800);
  border-color: var(--gray-800);
  transition: all var(--transition-fast);
}

.buy-now-btn:hover {
  background: var(--gray-900);
  border-color: var(--gray-900);
}

/* Additional Actions */
.additional-actions {
  display: flex;
  justify-content: space-between;
  padding-top: 0.75rem;
  border-top: 1px solid var(--gray-100);
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

/* Loading Skeleton */
.loading-skeleton {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(320px, 1fr));
  gap: 2rem;
}

.skeleton-card {
  background: var(--white);
  border-radius: 16px;
  overflow: hidden;
  border: 1px solid var(--gray-100);
}

.skeleton-image {
  width: 100%;
  height: 240px;
  background: var(--gray-100);
}

.skeleton-content {
  padding: 1.5rem;
}

/* Responsive Design */
@media (max-width: 1200px) {
  .products-layout {
    grid-template-columns: 250px 1fr;
    gap: 1.5rem;
  }

  .view-grid {
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 1.5rem;
  }

  .loading-skeleton {
    grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
    gap: 1.5rem;
  }
}

@media (max-width: 992px) {
  .products-layout {
    grid-template-columns: 1fr;
    gap: 2rem;
  }

  .filter-sidebar-container {
    position: static;
  }

  .content-header {
    flex-direction: column;
    gap: 1.5rem;
    align-items: stretch;
  }

  .controls-section {
    flex-direction: row;
    justify-content: space-between;
    align-items: center;
  }
}

@media (max-width: 768px) {
  .content-header {
    padding: 1rem;
  }

  .controls-section {
    flex-direction: column;
    gap: 1rem;
    align-items: stretch;
  }

  .sort-controls {
    min-width: auto;
  }

  .view-controls {
    align-self: center;
  }

  .view-grid {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .loading-skeleton {
    grid-template-columns: 1fr;
    gap: 1.5rem;
  }

  .product-list-item {
    grid-template-columns: 150px 1fr;
    gap: 1rem;
    padding: 1rem;
  }

  .list-image {
    height: 150px;
  }

  .list-actions {
    flex-direction: column;
    gap: 0.5rem;
  }

  .product-image-container {
    height: 200px;
  }

  .product-info-section {
    padding: 1rem;
  }

  .product-actions {
    flex-direction: column;
  }

  .add-to-cart-btn,
  .buy-now-btn {
    flex: 1;
  }

  .additional-actions {
    flex-direction: column;
    gap: 0.5rem;
  }

  .price-section {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }

  .current-price .amount {
    font-size: 1.25rem;
  }
}

@media (max-width: 480px) {
  .products-page {
    padding: 1rem 0;
  }

  .page-header h1 {
    font-size: 2rem;
  }

  .content-header {
    padding: 0.75rem;
  }

  .view-grid {
    gap: 1rem;
  }

  .loading-skeleton {
    gap: 1rem;
  }

  .product-card {
    border-radius: 12px;
  }

  .product-list-item {
    grid-template-columns: 1fr;
    gap: 1rem;
    text-align: center;
  }

  .list-image {
    width: 100%;
    height: 200px;
    margin: 0 auto;
  }

  .list-header {
    flex-direction: column;
    gap: 0.5rem;
    align-items: center;
  }

  .list-meta {
    justify-content: center;
    flex-wrap: wrap;
  }

  .product-image-container {
    height: 180px;
  }

  .product-info-section {
    padding: 0.75rem;
  }

  .product-name {
    font-size: 1.125rem;
  }

  .product-actions {
    gap: 0.5rem;
  }

  .add-to-cart-btn,
  .buy-now-btn {
    height: 40px;
    font-size: 0.875rem;
  }

  .overlay-actions {
    flex-direction: column;
    gap: 0.25rem;
  }

  .stock-badge,
  .featured-badge {
    font-size: 0.625rem;
    padding: 0.25rem 0.5rem;
  }
}
</style>