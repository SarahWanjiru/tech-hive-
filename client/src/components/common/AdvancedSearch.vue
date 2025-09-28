<template>
  <div class="advanced-search" ref="searchContainer">
    <div class="search-wrapper">
      <el-autocomplete
        v-model="searchQuery"
        :fetch-suggestions="querySearch"
        :placeholder="placeholder"
        class="search-autocomplete"
        @select="handleSelect"
        @focus="handleFocus"
        @blur="handleBlur"
        clearable
      >
        <template #prefix>
          <el-icon><Search /></el-icon>
        </template>
        <template #suffix>
          <el-button @click="handleSearch" type="primary" size="small" :loading="searching">
            Search
          </el-button>
        </template>

        <template #default="{ item }">
          <div class="suggestion-item">
            <div class="suggestion-icon">
              <el-icon v-if="item.type === 'product'"><Box /></el-icon>
              <el-icon v-else-if="item.type === 'category'"><Grid /></el-icon>
              <el-icon v-else><Search /></el-icon>
            </div>
            <div class="suggestion-content">
              <div class="suggestion-title">{{ item.title }}</div>
              <div class="suggestion-subtitle">{{ item.subtitle }}</div>
            </div>
            <div class="suggestion-badge" v-if="item.badge">
              {{ item.badge }}
            </div>
          </div>
        </template>
      </el-autocomplete>

      <!-- Search Suggestions Dropdown -->
      <div v-if="showSuggestions && suggestions.length > 0" class="search-suggestions">
        <div class="suggestions-header">
          <span>Popular searches</span>
          <el-button text size="small" @click="showSuggestions = false">
            <el-icon><Close /></el-icon>
          </el-button>
        </div>
        <div class="suggestions-list">
          <div
            v-for="suggestion in suggestions"
            :key="suggestion.id"
            class="suggestion-item"
            @click="selectSuggestion(suggestion)"
          >
            <div class="suggestion-icon">
              <el-icon v-if="suggestion.type === 'product'"><Box /></el-icon>
              <el-icon v-else-if="suggestion.type === 'category'"><Grid /></el-icon>
              <el-icon v-else><Search /></el-icon>
            </div>
            <div class="suggestion-content">
              <div class="suggestion-title">{{ suggestion.title }}</div>
              <div class="suggestion-subtitle">{{ suggestion.subtitle }}</div>
            </div>
          </div>
        </div>
      </div>

      <!-- Recent Searches -->
      <div v-if="showRecentSearches && recentSearches.length > 0" class="recent-searches">
        <div class="recent-header">
          <span>Recent searches</span>
          <el-button text size="small" @click="clearRecentSearches">
            Clear all
          </el-button>
        </div>
        <div class="recent-list">
          <el-tag
            v-for="search in recentSearches.slice(0, 5)"
            :key="search"
            size="small"
            closable
            @close="removeRecentSearch(search)"
            @click="selectRecentSearch(search)"
            class="recent-tag"
          >
            {{ search }}
          </el-tag>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { Search, Box, Grid, Close } from '@element-plus/icons-vue'

const router = useRouter()
const emit = defineEmits(['search'])

// Props
defineProps({
  placeholder: {
    type: String,
    default: 'Search products...'
  }
})

// Reactive data
const searchQuery = ref('')
const suggestions = ref([])
const recentSearches = ref([])
const showSuggestions = ref(false)
const showRecentSearches = ref(false)
const searching = ref(false)
const searchContainer = ref(null)

// Computed
const searchHistoryKey = computed(() => 'search_history')

// Methods
const querySearch = (queryString, callback) => {
  if (queryString) {
    // Simulate API call for suggestions
    const mockSuggestions = [
      {
        id: 1,
        type: 'product',
        title: `iPhone ${queryString}`,
        subtitle: 'Mobile Phone',
        badge: 'Product'
      },
      {
        id: 2,
        type: 'category',
        title: `${queryString} Electronics`,
        subtitle: 'Category',
        badge: 'Category'
      }
    ]

    // Filter suggestions based on query
    const filtered = mockSuggestions.filter(item =>
      item.title.toLowerCase().includes(queryString.toLowerCase())
    )

    callback(filtered)
  } else {
    callback([])
  }
}

const handleSelect = (item) => {
  searchQuery.value = item.title
  performSearch(item.title)
}

const handleSearch = async () => {
  if (!searchQuery.value.trim()) return

  searching.value = true

  try {
    // Add to recent searches
    addToRecentSearches(searchQuery.value)

    // Perform search
    await performSearch(searchQuery.value)

    // Hide suggestions
    showSuggestions.value = false
  } finally {
    searching.value = false
  }
}

const performSearch = async (query) => {
  emit('search', query)

  // Navigate to products page with search query
  router.push(`/products?search=${encodeURIComponent(query)}`)
}

const handleFocus = () => {
  if (searchQuery.value.trim()) {
    showSuggestions.value = true
  } else {
    showRecentSearches.value = true
  }
}

const handleBlur = () => {
  // Delay hiding to allow for clicks
  setTimeout(() => {
    showSuggestions.value = false
    showRecentSearches.value = false
  }, 200)
}

const selectSuggestion = (suggestion) => {
  searchQuery.value = suggestion.title
  performSearch(suggestion.title)
  showSuggestions.value = false
}

const selectRecentSearch = (search) => {
  searchQuery.value = search
  performSearch(search)
}

const addToRecentSearches = (search) => {
  if (!search.trim()) return

  const history = recentSearches.value.filter(s => s !== search)
  history.unshift(search)

  // Keep only last 10 searches
  recentSearches.value = history.slice(0, 10)

  // Save to localStorage
  localStorage.setItem(searchHistoryKey.value, JSON.stringify(recentSearches.value))
}

const removeRecentSearch = (search) => {
  recentSearches.value = recentSearches.value.filter(s => s !== search)
  localStorage.setItem(searchHistoryKey.value, JSON.stringify(recentSearches.value))
}

const clearRecentSearches = () => {
  recentSearches.value = []
  localStorage.removeItem(searchHistoryKey.value)
}

// Load recent searches on mount
onMounted(() => {
  const saved = localStorage.getItem(searchHistoryKey.value)
  if (saved) {
    try {
      recentSearches.value = JSON.parse(saved)
    } catch (e) {
      console.error('Failed to load search history:', e)
    }
  }

  // Click outside to close
  document.addEventListener('click', (e) => {
    if (searchContainer.value && !searchContainer.value.contains(e.target)) {
      showSuggestions.value = false
      showRecentSearches.value = false
    }
  })
})

onUnmounted(() => {
  // Event listener is removed automatically since it was added inline
})
</script>

<style scoped>
.advanced-search {
  position: relative;
  width: 100%;
}

.search-wrapper {
  position: relative;
}

.search-autocomplete {
  width: 100%;
}

.search-autocomplete :deep(.el-input__wrapper) {
  border-radius: 50px;
  border: 2px solid transparent;
  transition: all var(--transition-fast);
  padding-right: 120px;
}

.search-autocomplete :deep(.el-input__wrapper:hover) {
  border-color: var(--primary-color);
}

.search-autocomplete :deep(.el-input__wrapper.is-focus) {
  border-color: var(--primary-color);
  box-shadow: 0 0 0 3px rgba(64, 158, 255, 0.1);
}

.search-suggestions,
.recent-searches {
  position: absolute;
  top: 100%;
  left: 0;
  right: 0;
  background: var(--white);
  border: 1px solid var(--gray-200);
  border-radius: 12px;
  box-shadow: 0 8px 30px rgba(0, 0, 0, 0.15);
  z-index: 1000;
  margin-top: 4px;
  max-height: 400px;
  overflow: hidden;
}

.suggestions-header,
.recent-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid var(--gray-100);
  background: var(--gray-50);
}

.suggestions-list,
.recent-list {
  max-height: 300px;
  overflow-y: auto;
}

.suggestion-item {
  display: flex;
  align-items: center;
  gap: 1rem;
  padding: 1rem;
  cursor: pointer;
  transition: background-color var(--transition-fast);
  border-bottom: 1px solid var(--gray-100);
}

.suggestion-item:hover {
  background: var(--gray-50);
}

.suggestion-item:last-child {
  border-bottom: none;
}

.suggestion-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: var(--primary-color-light);
  border-radius: 8px;
  color: var(--primary-color);
}

.suggestion-content {
  flex: 1;
}

.suggestion-title {
  font-weight: 600;
  color: var(--gray-800);
  margin-bottom: 0.25rem;
}

.suggestion-subtitle {
  font-size: 0.875rem;
  color: var(--gray-600);
}

.suggestion-badge {
  background: var(--success-color);
  color: var(--white);
  padding: 0.25rem 0.5rem;
  border-radius: 12px;
  font-size: 0.75rem;
  font-weight: 600;
}

.recent-list {
  display: flex;
  flex-wrap: wrap;
  gap: 0.5rem;
  padding: 1rem;
}

.recent-tag {
  cursor: pointer;
  transition: all var(--transition-fast);
}

.recent-tag:hover {
  background: var(--primary-color-light);
  color: var(--primary-color);
}
</style>