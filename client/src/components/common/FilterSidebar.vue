<template>
  <div class="filter-sidebar">
    <div class="filter-header">
      <h3>
        <el-icon><Filter /></el-icon>
        Filters
      </h3>
      <el-button text size="small" @click="clearAllFilters">
        Clear All
      </el-button>
    </div>

    <el-collapse v-model="activePanels" class="filter-collapse">
      <!-- Categories Filter -->
      <el-collapse-item name="categories" title="Categories">
        <div class="filter-section">
          <el-checkbox-group v-model="filters.categories" @change="handleFilterChange">
            <div v-for="category in availableCategories" :key="category.id" class="category-item">
              <el-checkbox :label="category.id" size="large">
                <div class="category-info">
                  <span class="category-name">{{ category.name }}</span>
                  <span class="category-count">({{ category.count }})</span>
                </div>
              </el-checkbox>
            </div>
          </el-checkbox-group>
        </div>
      </el-collapse-item>

      <!-- Price Range Filter -->
      <el-collapse-item name="price" title="Price Range">
        <div class="filter-section">
          <div class="price-inputs">
            <el-input-number
              v-model="filters.priceRange[0]"
              :min="0"
              :max="filters.priceRange[1]"
              :precision="2"
              placeholder="Min price"
              @change="handlePriceChange"
              size="large"
            />
            <span class="price-separator">-</span>
            <el-input-number
              v-model="filters.priceRange[1]"
              :min="filters.priceRange[0]"
              :max="10000"
              :precision="2"
              placeholder="Max price"
              @change="handlePriceChange"
              size="large"
            />
          </div>

          <!-- Price Range Slider -->
          <div class="price-slider">
            <el-slider
              v-model="filters.priceRange"
              range
              :min="0"
              :max="10000"
              :step="10"
              @change="handlePriceChange"
              :format-tooltip="formatPrice"
            />
          </div>

          <!-- Quick Price Ranges -->
          <div class="price-ranges">
            <el-button
              v-for="range in priceRanges"
              :key="range.label"
              size="small"
              :type="isPriceRangeActive(range) ? 'primary' : 'default'"
              @click="setPriceRange(range)"
            >
              {{ range.label }}
            </el-button>
          </div>
        </div>
      </el-collapse-item>

      <!-- Brand Filter -->
      <el-collapse-item name="brands" title="Brands">
        <div class="filter-section">
          <el-input
            v-model="brandSearch"
            placeholder="Search brands..."
            size="large"
            clearable
          >
            <template #prefix>
              <el-icon><Search /></el-icon>
            </template>
          </el-input>

          <el-checkbox-group v-model="filters.brands" @change="handleFilterChange" class="brands-list">
            <div v-for="brand in filteredBrands" :key="brand.id" class="brand-item">
              <el-checkbox :label="brand.id" size="large">
                <div class="brand-info">
                  <span class="brand-name">{{ brand.name }}</span>
                  <span class="brand-count">({{ brand.count }})</span>
                </div>
              </el-checkbox>
            </div>
          </el-checkbox-group>
        </div>
      </el-collapse-item>

      <!-- Rating Filter -->
      <el-collapse-item name="rating" title="Customer Rating">
        <div class="filter-section">
          <el-radio-group v-model="filters.rating" @change="handleFilterChange">
            <div v-for="rating in ratingOptions" :key="rating.value" class="rating-option">
              <el-radio :label="rating.value" size="large">
                <div class="rating-info">
                  <el-rate :model-value="rating.value" disabled size="small" />
                  <span class="rating-label">{{ rating.label }}</span>
                </div>
              </el-radio>
            </div>
          </el-radio-group>
        </div>
      </el-collapse-item>

      <!-- Availability Filter -->
      <el-collapse-item name="availability" title="Availability">
        <div class="filter-section">
          <el-checkbox-group v-model="filters.availability" @change="handleFilterChange">
            <el-checkbox label="in_stock" size="large">
              <div class="availability-info">
                <el-icon><Check /></el-icon>
                <span>In Stock</span>
              </div>
            </el-checkbox>
            <el-checkbox label="out_of_stock" size="large">
              <div class="availability-info">
                <el-icon><Close /></el-icon>
                <span>Out of Stock</span>
              </div>
            </el-checkbox>
            <el-checkbox label="pre_order" size="large">
              <div class="availability-info">
                <el-icon><Clock /></el-icon>
                <span>Pre-order</span>
              </div>
            </el-checkbox>
          </el-checkbox-group>
        </div>
      </el-collapse-item>

      <!-- Special Features -->
      <el-collapse-item name="features" title="Special Features">
        <div class="filter-section">
          <el-checkbox-group v-model="filters.features" @change="handleFilterChange">
            <el-checkbox label="featured" size="large">
              <div class="feature-info">
                <el-icon><Star /></el-icon>
                <span>Featured Products</span>
              </div>
            </el-checkbox>
            <el-checkbox label="on_sale" size="large">
              <div class="feature-info">
                <el-icon><Discount /></el-icon>
                <span>On Sale</span>
              </div>
            </el-checkbox>
            <el-checkbox label="new_arrival" size="large">
              <div class="feature-info">
                <el-icon><Present /></el-icon>
                <span>New Arrival</span>
              </div>
            </el-checkbox>
            <el-checkbox label="free_shipping" size="large">
              <div class="feature-info">
                <el-icon><Van /></el-icon>
                <span>Free Shipping</span>
              </div>
            </el-checkbox>
          </el-checkbox-group>
        </div>
      </el-collapse-item>
    </el-collapse>

    <!-- Active Filters Summary -->
    <div v-if="hasActiveFilters" class="active-filters">
      <h4>Active Filters</h4>
      <div class="active-filters-list">
        <!-- Active Categories -->
        <el-tag
          v-for="categoryId in filters.categories"
          :key="`category-${categoryId}`"
          closable
          size="small"
          @close="removeCategoryFilter(categoryId)"
          class="active-filter-tag"
        >
          <el-icon><Grid /></el-icon>
          {{ getCategoryName(categoryId) }}
        </el-tag>

        <!-- Active Price Range -->
        <el-tag
          v-if="filters.priceRange[0] > 0 || filters.priceRange[1] < 10000"
          closable
          size="small"
          @close="clearPriceFilter"
          class="active-filter-tag"
        >
          <el-icon><Money /></el-icon>
          {{ formatPrice(filters.priceRange[0]) }} - {{ formatPrice(filters.priceRange[1]) }}
        </el-tag>

        <!-- Active Brands -->
        <el-tag
          v-for="brandId in filters.brands"
          :key="`brand-${brandId}`"
          closable
          size="small"
          @close="removeBrandFilter(brandId)"
          class="active-filter-tag"
        >
          <el-icon><Shop /></el-icon>
          {{ getBrandName(brandId) }}
        </el-tag>

        <!-- Active Rating -->
        <el-tag
          v-if="filters.rating"
          closable
          size="small"
          @close="clearRatingFilter"
          class="active-filter-tag"
        >
          <el-icon><Star /></el-icon>
          {{ filters.rating }} Stars & Up
        </el-tag>

        <!-- Active Features -->
        <el-tag
          v-for="feature in filters.features"
          :key="`feature-${feature}`"
          closable
          size="small"
          @close="removeFeatureFilter(feature)"
          class="active-filter-tag"
        >
          <el-icon><Check /></el-icon>
          {{ getFeatureLabel(feature) }}
        </el-tag>
      </div>
    </div>

    <!-- Filter Actions -->
    <div class="filter-actions">
      <el-button @click="applyFilters" type="primary" size="large" class="apply-btn">
        Apply Filters ({{ totalResults }})
      </el-button>
      <el-button @click="resetFilters" size="large" class="reset-btn">
        Reset
      </el-button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, watch } from 'vue'
import {
  Filter,
  Search,
  Check,
  Close,
  Clock,
  Star,
  Discount,
  Present,
  Van,
  Money,
  Shop,
  Grid
} from '@element-plus/icons-vue'

const emit = defineEmits(['filter-change', 'apply-filters'])

// Props
const props = defineProps({
  availableCategories: {
    type: Array,
    default: () => []
  },
  availableBrands: {
    type: Array,
    default: () => []
  },
  totalResults: {
    type: Number,
    default: 0
  }
})

// Reactive data
const activePanels = ref(['categories', 'price'])
const brandSearch = ref('')

const filters = ref({
  categories: [],
  priceRange: [0, 10000],
  brands: [],
  rating: null,
  availability: [],
  features: []
})

// Computed
const filteredBrands = computed(() => {
  if (!brandSearch.value) return props.availableBrands
  return props.availableBrands.filter(brand =>
    brand.name.toLowerCase().includes(brandSearch.value.toLowerCase())
  )
})

const hasActiveFilters = computed(() => {
  return filters.value.categories.length > 0 ||
         filters.value.priceRange[0] > 0 ||
         filters.value.priceRange[1] < 10000 ||
         filters.value.brands.length > 0 ||
         filters.value.rating ||
         filters.value.availability.length > 0 ||
         filters.value.features.length > 0
})

const priceRanges = computed(() => [
  { label: 'Under KES 1,000', range: [0, 1000] },
  { label: 'KES 1,000 - 5,000', range: [1000, 5000] },
  { label: 'KES 5,000 - 10,000', range: [5000, 10000] },
  { label: 'Over KES 10,000', range: [10000, 100000] }
])

const ratingOptions = computed(() => [
  { value: 4, label: '4 Stars & Up' },
  { value: 3, label: '3 Stars & Up' },
  { value: 2, label: '2 Stars & Up' },
  { value: 1, label: '1 Star & Up' }
])

// Methods
const handleFilterChange = () => {
  emit('filter-change', filters.value)
}

const handlePriceChange = () => {
  handleFilterChange()
}

const formatPrice = (value) => {
  return new Intl.NumberFormat('en-KE', {
    style: 'currency',
    currency: 'KES',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
  }).format(value).replace('KES', 'KES ')
}

const setPriceRange = (range) => {
  filters.value.priceRange = [...range.range]
  handlePriceChange()
}

const isPriceRangeActive = (range) => {
  return filters.value.priceRange[0] === range.range[0] &&
         filters.value.priceRange[1] === range.range[1]
}

const getCategoryName = (categoryId) => {
  const category = props.availableCategories.find(c => c.id === categoryId)
  return category ? category.name : ''
}

const getBrandName = (brandId) => {
  const brand = props.availableBrands.find(b => b.id === brandId)
  return brand ? brand.name : ''
}

const getFeatureLabel = (feature) => {
  const labels = {
    featured: 'Featured',
    on_sale: 'On Sale',
    new_arrival: 'New Arrival',
    free_shipping: 'Free Shipping'
  }
  return labels[feature] || feature
}

const removeCategoryFilter = (categoryId) => {
  filters.value.categories = filters.value.categories.filter(id => id !== categoryId)
  handleFilterChange()
}

const removeBrandFilter = (brandId) => {
  filters.value.brands = filters.value.brands.filter(id => id !== brandId)
  handleFilterChange()
}

const removeFeatureFilter = (feature) => {
  filters.value.features = filters.value.features.filter(f => f !== feature)
  handleFilterChange()
}

const clearPriceFilter = () => {
  filters.value.priceRange = [0, 10000]
  handlePriceChange()
}

const clearRatingFilter = () => {
  filters.value.rating = null
  handleFilterChange()
}

const clearAllFilters = () => {
  filters.value = {
    categories: [],
    priceRange: [0, 10000],
    brands: [],
    rating: null,
    availability: [],
    features: []
  }
  brandSearch.value = ''
  emit('filter-change', filters.value)
}

const resetFilters = () => {
  clearAllFilters()
}

const applyFilters = () => {
  emit('apply-filters', filters.value)
}

// Watch for external filter changes
watch(() => props.totalResults, (newCount) => {
  // Update UI when results count changes
})
</script>

<style scoped>
.filter-sidebar {
  background: var(--white);
  border-radius: 16px;
  padding: 1.5rem;
  box-shadow: 0 8px 32px rgba(0, 0, 0, 0.12), 0 2px 8px rgba(0, 0, 0, 0.08);
  border: 2px solid var(--gray-200);
  height: fit-content;
  position: sticky;
  top: 2rem;
  transition: all var(--transition-normal);
}

.filter-sidebar:hover {
  box-shadow: 0 12px 40px rgba(0, 0, 0, 0.15), 0 4px 12px rgba(0, 0, 0, 0.1);
  border-color: var(--primary-color);
}

.filter-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding: 1rem 1.25rem;
  background: linear-gradient(135deg, var(--primary-color), #667eea);
  border-radius: 12px;
  border-bottom: 3px solid var(--primary-color-dark);
}

.filter-header h3 {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  margin: 0;
  color: var(--white);
  font-weight: 700;
  font-size: 1.125rem;
  text-shadow: 0 1px 2px rgba(0, 0, 0, 0.1);
}

.filter-header .el-button {
  color: var(--white);
  border-color: rgba(255, 255, 255, 0.3);
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
}

.filter-header .el-button:hover {
  background: rgba(255, 255, 255, 0.2);
  border-color: var(--white);
  color: var(--white);
}

.filter-collapse {
  border: none;
}

.filter-collapse :deep(.el-collapse-item__header) {
  font-weight: 700;
  color: var(--gray-800);
  padding: 1.25rem;
  border-bottom: 2px solid var(--gray-200);
  background: var(--gray-50);
  border-radius: 8px;
  margin-bottom: 0.5rem;
  transition: all var(--transition-fast);
}

.filter-collapse :deep(.el-collapse-item__header:hover) {
  background: var(--primary-color-light);
  color: var(--primary-color);
  border-color: var(--primary-color);
}

.filter-collapse :deep(.el-collapse-item__content) {
  padding: 1.5rem;
  background: var(--white);
  border-radius: 0 0 8px 8px;
  border: 1px solid var(--gray-200);
  border-top: none;
}

.filter-section {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.category-item,
.brand-item {
  margin-bottom: 0.5rem;
}

.category-info,
.brand-info {
  display: flex;
  justify-content: space-between;
  width: 100%;
  font-size: 0.875rem;
}

.category-name,
.brand-name {
  font-weight: 500;
}

.category-count,
.brand-count {
  color: var(--gray-500);
}

.price-inputs {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.price-separator {
  color: var(--gray-500);
  font-weight: 600;
}

.price-ranges {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
}

.price-ranges .el-button {
  font-size: 0.75rem;
}

.rating-option {
  margin-bottom: 0.75rem;
}

.rating-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.rating-label {
  font-size: 0.875rem;
  color: var(--gray-700);
}

.availability-info,
.feature-info {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  font-size: 0.875rem;
  color: var(--gray-700);
}

.brands-list {
  margin-top: 1rem;
}

.active-filters {
  margin-top: 2rem;
  padding: 1.5rem;
  background: linear-gradient(135deg, #fff3cd, #ffeaa7);
  border: 2px solid #ffc107;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(255, 193, 7, 0.2);
}

.active-filters h4 {
  margin: 0 0 1rem 0;
  color: #856404;
  font-size: 1.125rem;
  font-weight: 700;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.active-filters-list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  margin-bottom: 1.5rem;
}

.active-filter-tag {
  display: flex;
  align-items: center;
  gap: 0.25rem;
  background: var(--primary-color-light);
  color: var(--primary-color);
  border: 1px solid var(--primary-color);
}

.filter-actions {
  display: flex;
  flex-direction: column;
  gap: 1rem;
  padding: 1.5rem;
  background: linear-gradient(135deg, var(--gray-50), #f8f9fa);
  border: 2px solid var(--gray-200);
  border-radius: 12px;
  margin-top: 1.5rem;
}

.apply-btn {
  background: linear-gradient(135deg, var(--primary-color), #667eea);
  border: none;
  font-weight: 700;
  font-size: 1rem;
  padding: 1rem;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(64, 158, 255, 0.3);
  transition: all var(--transition-fast);
}

.apply-btn:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 16px rgba(64, 158, 255, 0.4);
}

.reset-btn {
  border: 2px solid var(--gray-300);
  color: var(--gray-700);
  font-weight: 600;
  padding: 0.875rem;
  border-radius: 8px;
  background: var(--white);
  transition: all var(--transition-fast);
}

.reset-btn:hover {
  border-color: var(--primary-color);
  color: var(--primary-color);
  background: var(--primary-color-light);
  transform: translateY(-1px);
}

/* Responsive */
@media (max-width: 768px) {
  .filter-sidebar {
    position: static;
    margin-bottom: 2rem;
  }

  .filter-actions {
    flex-direction: row;
  }

  .apply-btn,
  .reset-btn {
    flex: 1;
  }
}
</style>