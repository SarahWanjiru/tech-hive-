<template>
  <div class="mobile-bottom-nav" v-show="visible">
    <div class="nav-container">
      <!-- Home -->
      <router-link to="/" class="nav-item" :class="{ 'active': isActiveRoute('/') }">
        <div class="nav-icon">
          <el-icon><House /></el-icon>
        </div>
        <div class="nav-label">Home</div>
      </router-link>

      <!-- Products -->
      <router-link to="/products" class="nav-item" :class="{ 'active': isActiveRoute('/products') }">
        <div class="nav-icon">
          <el-badge :value="totalProducts" :hidden="totalProducts === 0" class="product-badge">
            <el-icon><ShoppingBag /></el-icon>
          </el-badge>
        </div>
        <div class="nav-label">Products</div>
      </router-link>

      <!-- Search -->
      <div class="nav-item search-item" @click="openSearch">
        <div class="nav-icon search-icon">
          <el-icon><Search /></el-icon>
        </div>
        <div class="nav-label">Search</div>
      </div>

      <!-- Cart -->
      <router-link v-if="authStore.isAuthenticated" to="/cart" class="nav-item" :class="{ 'active': isActiveRoute('/cart') }">
        <div class="nav-icon">
          <el-badge :value="cartStore.totalItems" :hidden="cartStore.totalItems === 0" class="cart-badge">
            <el-icon><ShoppingCart /></el-icon>
          </el-badge>
        </div>
        <div class="nav-label">Cart</div>
      </router-link>

      <!-- Profile -->
      <router-link v-if="authStore.isAuthenticated" to="/profile" class="nav-item" :class="{ 'active': isActiveRoute('/profile') }">
        <div class="nav-icon">
          <el-avatar :size="24" class="nav-avatar">
            {{ authStore.userName.charAt(0).toUpperCase() }}
          </el-avatar>
        </div>
        <div class="nav-label">Profile</div>
      </router-link>

      <!-- Auth buttons for non-authenticated users -->
      <div v-if="!authStore.isAuthenticated" class="nav-item" @click="navigateToLogin">
        <div class="nav-icon">
          <el-icon><User /></el-icon>
        </div>
        <div class="nav-label">Login</div>
      </div>
    </div>

    <!-- Search Modal -->
    <div v-show="showSearchModal" class="search-modal" @click="closeSearchModal">
      <div class="search-modal-content" @click.stop>
        <div class="search-modal-header">
          <h3>Search Products</h3>
          <el-button circle size="small" @click="closeSearchModal">
            <el-icon><Close /></el-icon>
          </el-button>
        </div>
        <div class="search-modal-body">
          <AdvancedSearch
            ref="searchComponent"
            @search="handleMobileSearch"
            placeholder="Search for products..."
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../../stores/auth.store.ts'
import { useCartStore } from '../../stores/cart.store.ts'
import { useProductStore } from '../../stores/product.store.ts'
import AdvancedSearch from './AdvancedSearch.vue'
import {
  House,
  ShoppingBag,
  Search,
  ShoppingCart,
  User,
  Close
} from '@element-plus/icons-vue'

const route = useRoute()
const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()
const productStore = useProductStore()

// Props
defineProps({
  visible: {
    type: Boolean,
    default: true
  }
})

// Reactive data
const showSearchModal = ref(false)
const searchComponent = ref(null)

// Computed
const totalProducts = computed(() => {
  return productStore.filteredProducts?.length || 0
})

// Methods
const isActiveRoute = (path) => {
  if (path === '/') {
    return route.path === '/'
  }
  return route.path.startsWith(path)
}

const openSearch = () => {
  showSearchModal.value = true
}

const closeSearchModal = () => {
  showSearchModal.value = false
}

const handleMobileSearch = (query) => {
  closeSearchModal()
  router.push(`/products?search=${encodeURIComponent(query)}`)
}

const navigateToLogin = () => {
  router.push('/login')
}

// Handle escape key for search modal
const handleKeydown = (e) => {
  if (e.key === 'Escape' && showSearchModal.value) {
    closeSearchModal()
  }
}

onMounted(() => {
  document.addEventListener('keydown', handleKeydown)
})

onUnmounted(() => {
  document.removeEventListener('keydown', handleKeydown)
})
</script>

<style scoped>
.mobile-bottom-nav {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  background: var(--white);
  border-top: 1px solid var(--gray-200);
  box-shadow: 0 -4px 20px rgba(0, 0, 0, 0.1);
  z-index: 1000;
  padding-bottom: env(safe-area-inset-bottom);
}

.nav-container {
  display: flex;
  justify-content: space-around;
  align-items: center;
  padding: 0.75rem 0.5rem;
  max-width: 480px;
  margin: 0 auto;
}

.nav-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: 0.25rem;
  padding: 0.5rem;
  border-radius: 12px;
  text-decoration: none;
  color: var(--gray-600);
  transition: all var(--transition-fast);
  min-width: 60px;
  position: relative;
}

.nav-item.active {
  color: var(--primary-color);
  background: var(--primary-color-light);
}

.nav-item:hover:not(.active) {
  color: var(--gray-800);
  background: var(--gray-50);
}

.nav-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 24px;
  height: 24px;
  position: relative;
}

.nav-label {
  font-size: 0.75rem;
  font-weight: 500;
  line-height: 1;
}

.nav-avatar {
  background: linear-gradient(135deg, var(--primary-color), #667eea);
  color: var(--white);
  font-weight: 600;
  font-size: 0.75rem;
}

.search-item {
  cursor: pointer;
}

.search-icon {
  background: linear-gradient(135deg, var(--primary-color), #667eea);
  color: var(--white);
  border-radius: 6px;
}

/* Badges */
.product-badge :deep(.el-badge__content),
.cart-badge :deep(.el-badge__content) {
  background: var(--primary-color) !important;
  border-color: var(--white) !important;
  font-size: 0.625rem;
  height: 14px;
  line-height: 12px;
  padding: 0 4px;
  min-width: 14px;
}

/* Search Modal */
.search-modal {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 2000;
  padding: 1rem;
}

.search-modal-content {
  background: var(--white);
  border-radius: 16px;
  width: 100%;
  max-width: 500px;
  max-height: 80vh;
  overflow: hidden;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.3);
}

.search-modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1.5rem;
  border-bottom: 1px solid var(--gray-200);
}

.search-modal-header h3 {
  margin: 0;
  color: var(--gray-800);
  font-weight: 600;
}

.search-modal-body {
  padding: 1.5rem;
}

/* Desktop hide */
@media (min-width: 769px) {
  .mobile-bottom-nav {
    display: none;
  }
}

/* Safe area for devices with notches */
@supports (padding: max(0px)) {
  .mobile-bottom-nav {
    padding-bottom: max(0.75rem, env(safe-area-inset-bottom));
  }
}
</style>