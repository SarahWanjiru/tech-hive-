<template>
  <div id="app">
    <!-- Enhanced Navigation Header -->
    <header class="main-header">
      <!-- Top Bar -->
      <div class="top-bar">
        <div class="container">
          <div class="top-bar-content">
            <div class="top-links">
              <a href="#" class="top-link">Help Center</a>
              <a href="#" class="top-link">Track Order</a>
              <a href="#" class="top-link">Returns</a>
            </div>
            <div class="top-actions">
              <span v-if="authStore.isAuthenticated" class="welcome-text">
                Welcome back, {{ authStore.userName }}!
              </span>
              <div v-else class="currency-selector">
                <el-select v-model="currency" placeholder="Currency" size="small">
                  <el-option label="KES" value="kes" />
                  <el-option label="USD" value="usd" />
                </el-select>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Main Navigation -->
      <nav class="navbar">
        <div class="container">
          <div class="navbar-content">
            <!-- Logo -->
            <div class="navbar-brand">
              <router-link to="/" class="brand-link">
                <img src="/favicon.ico" alt="Logo" class="brand-logo" />
                <span class="brand-text">TechHive</span>
              </router-link>
            </div>

            <!-- Advanced Search Bar -->
            <div class="search-section" role="search" aria-label="Product search">
              <AdvancedSearch @search="handleAdvancedSearch" />
            </div>

            <!-- Navigation Menu & Actions -->
            <div class="navbar-right">
              <!-- Desktop Navigation -->
              <div class="navbar-menu desktop-menu">
                <ul class="nav-links">
                  <!-- Customer Navigation -->
                  <template v-if="!authStore.isAuthenticated || !authStore.isAdmin">
                    <li class="nav-item has-dropdown">
                      <a href="#" class="nav-link" @click.prevent="toggleCategories">
                        <el-icon><Grid /></el-icon>
                        Categories
                        <el-icon class="dropdown-arrow"><ArrowDown /></el-icon>
                      </a>
                      <!-- Mega Menu -->
                      <div v-show="showCategories" class="mega-menu">
                        <div class="mega-menu-content">
                          <div class="category-section">
                            <h4>Electronics</h4>
                            <ul>
                              <li><router-link to="/products?category=phones">Phones</router-link></li>
                              <li><router-link to="/products?category=laptops">Laptops</router-link></li>
                              <li><router-link to="/products?category=tablets">Tablets</router-link></li>
                              <li><router-link to="/products?category=headphones">Headphones</router-link></li>
                            </ul>
                          </div>
                          <div class="category-section">
                            <h4>Fashion</h4>
                            <ul>
                              <li><router-link to="/products?category=clothing">Clothing</router-link></li>
                              <li><router-link to="/products?category=shoes">Shoes</router-link></li>
                              <li><router-link to="/products?category=accessories">Accessories</router-link></li>
                            </ul>
                          </div>
                          <div class="category-section">
                            <h4>Home & Garden</h4>
                            <ul>
                              <li><router-link to="/products?category=furniture">Furniture</router-link></li>
                              <li><router-link to="/products?category=appliances">Appliances</router-link></li>
                              <li><router-link to="/products?category=decor">Home Decor</router-link></li>
                            </ul>
                          </div>
                        </div>
                      </div>
                    </li>
                    <li>
                      <router-link to="/products" class="nav-link">
                        All Products
                      </router-link>
                    </li>
                    <li>
                      <router-link to="/products?featured=true" class="nav-link featured-link">
                        <el-icon><Star /></el-icon>
                        Featured
                      </router-link>
                    </li>
                  </template>

                  <!-- Admin Navigation -->
                  <template v-if="authStore.isAuthenticated && authStore.isAdmin">
                    <li>
                      <router-link to="/admin" class="nav-link admin-link">
                        <el-icon><Setting /></el-icon>
                        Admin Dashboard
                      </router-link>
                    </li>
                    <li>
                      <router-link to="/products" class="nav-link">
                        Browse Products
                      </router-link>
                    </li>
                  </template>
                </ul>
              </div>

              <!-- User Actions -->
              <div class="navbar-actions">
                <!-- Cart Icon -->
                <router-link
                  v-if="authStore.isAuthenticated"
                  to="/cart"
                  class="cart-link"
                  :class="{ 'has-items': cartStore.totalItems > 0 }"
                >
                  <el-badge :value="cartStore.totalItems" :hidden="cartStore.totalItems === 0" class="cart-badge">
                    <el-button circle size="large" class="cart-button">
                      <el-icon size="20"><ShoppingCart /></el-icon>
                    </el-button>
                  </el-badge>
                </router-link>

                <!-- User Menu -->
                  <div v-if="authStore.isAuthenticated" class="user-menu">
                    <el-dropdown @visible-change="handleUserMenuToggle">
                      <el-button circle size="large" class="user-button">
                        <el-avatar :size="32" :src="userAvatar" class="user-avatar">
                          {{ authStore.userName.charAt(0).toUpperCase() }}
                        </el-avatar>
                      </el-button>
                      <template #dropdown>
                        <el-dropdown-menu class="user-dropdown">
                          <div class="user-info">
                            <strong>{{ authStore.userName }}</strong>
                            <small>{{ authStore.userEmail }}</small>
                            <el-tag v-if="authStore.isAdmin" size="small" type="danger">Admin</el-tag>
                            <el-tag v-else size="small" type="primary">Customer</el-tag>
                          </div>
 
                          <!-- Customer Menu Items -->
                          <template v-if="!authStore.isAdmin">
                            <el-dropdown-item @click="$router.push('/profile')" icon="User">
                              My Profile
                            </el-dropdown-item>
                            <el-dropdown-item @click="$router.push('/orders')" icon="Document">
                              My Orders
                            </el-dropdown-item>
                            <el-dropdown-item @click="$router.push('/wishlist')" icon="Star">
                              Wishlist
                            </el-dropdown-item>
                          </template>
 
                          <!-- Admin Menu Items -->
                          <template v-if="authStore.isAdmin">
                            <el-dropdown-item @click="$router.push('/admin')" icon="Setting">
                              Admin Dashboard
                            </el-dropdown-item>
                            <el-dropdown-item @click="$router.push('/profile')" icon="User">
                              My Profile
                            </el-dropdown-item>
                          </template>
 
                          <el-dropdown-item @click="handleLogout" divided icon="SwitchButton">
                            Logout
                          </el-dropdown-item>
                        </el-dropdown-menu>
                      </template>
                    </el-dropdown>
                  </div>

                <!-- Auth Buttons -->
                <div v-else class="auth-buttons">
                  <el-button size="large" @click="$router.push('/login')" text>
                    Sign In
                  </el-button>
                  <el-button size="large" type="primary" @click="$router.push('/register')">
                    Sign Up
                  </el-button>
                </div>

                <!-- Mobile Menu Toggle -->
                <el-button
                  class="mobile-menu-toggle"
                  circle
                  @click="toggleMobileMenu"
                >
                  <el-icon><Menu /></el-icon>
                </el-button>
              </div>
            </div>
          </div>

          <!-- Mobile Menu -->
          <div v-show="showMobileMenu" class="mobile-menu">
            <div class="mobile-menu-content">
              <ul class="mobile-nav-links">
                <li>
                  <router-link to="/" class="mobile-nav-link" @click="closeMobileMenu">
                    <el-icon><Home /></el-icon>
                    Home
                  </router-link>
                </li>

                <!-- Customer Navigation -->
                <template v-if="!authStore.isAuthenticated || !authStore.isAdmin">
                  <li>
                    <router-link to="/products" class="mobile-nav-link" @click="closeMobileMenu">
                      <el-icon><Grid /></el-icon>
                      Products
                    </router-link>
                  </li>
                  <li>
                    <router-link to="/products?featured=true" class="mobile-nav-link" @click="closeMobileMenu">
                      <el-icon><Star /></el-icon>
                      Featured
                    </router-link>
                  </li>
                </template>

                <!-- Admin Navigation -->
                <template v-if="authStore.isAuthenticated && authStore.isAdmin">
                  <li>
                    <router-link to="/admin" class="mobile-nav-link" @click="closeMobileMenu">
                      <el-icon><Setting /></el-icon>
                      Admin Dashboard
                    </router-link>
                  </li>
                  <li>
                    <router-link to="/products" class="mobile-nav-link" @click="closeMobileMenu">
                      <el-icon><Grid /></el-icon>
                      Browse Products
                    </router-link>
                  </li>
                </template>
              </ul>

              <div v-if="authStore.isAuthenticated" class="mobile-user-section">
                <div class="mobile-user-info">
                  <el-avatar :size="40" class="mobile-user-avatar">
                    {{ authStore.userName.charAt(0).toUpperCase() }}
                  </el-avatar>
                  <div>
                    <strong>{{ authStore.userName }}</strong>
                    <div>{{ authStore.userEmail }}</div>
                    <el-tag v-if="authStore.isAdmin" size="small" type="danger">Admin</el-tag>
                    <el-tag v-else size="small" type="primary">Customer</el-tag>
                  </div>
                </div>
                <div class="mobile-user-actions">
                  <!-- Customer Actions -->
                  <template v-if="!authStore.isAdmin">
                    <el-button @click="navigateToProfile" text>
                      <el-icon><User /></el-icon>
                      Profile
                    </el-button>
                    <el-button @click="navigateToOrders" text>
                      <el-icon><Document /></el-icon>
                      Orders
                    </el-button>
                  </template>

                  <!-- Admin Actions -->
                  <template v-if="authStore.isAdmin">
                    <el-button @click="$router.push('/admin')" text>
                      <el-icon><Setting /></el-icon>
                      Admin Dashboard
                    </el-button>
                    <el-button @click="navigateToProfile" text>
                      <el-icon><User /></el-icon>
                      Profile
                    </el-button>
                  </template>

                  <el-button @click="handleLogout" text>
                    <el-icon><SwitchButton /></el-icon>
                    Logout
                  </el-button>
                </div>
              </div>

              <div v-else class="mobile-auth-section">
                <el-button @click="navigateToLogin" size="large" plain>
                  Sign In
                </el-button>
                <el-button @click="navigateToRegister" size="large" type="primary">
                  Sign Up
                </el-button>
              </div>
            </div>
          </div>
        </div>
      </nav>
    </header>

    <!-- Main Content -->
    <main class="main-content">
      <router-view />
    </main>

    <!-- Mobile Bottom Navigation -->
    <MobileBottomNav />

    <!-- Footer -->
    <footer class="main-footer">
      <div class="container">
        <div class="footer-content">
          <div class="footer-section">
            <h4>TechHive</h4>
            <p>Your trusted online shopping destination</p>
          </div>

          <div class="footer-section">
            <h4>Quick Links</h4>
            <ul>
              <li><router-link to="/products">Products</router-link></li>
              <li><router-link to="/cart">Cart</router-link></li>
              <li><router-link to="/orders">Orders</router-link></li>
            </ul>
          </div>

          <div class="footer-section">
            <h4>Support</h4>
            <ul>
              <li><a href="#">Help Center</a></li>
              <li><a href="#">Contact Us</a></li>
              <li><a href="#">Returns</a></li>
            </ul>
          </div>
        </div>

        <div class="footer-bottom">
          <p>&copy; 2024 TechHive. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
 import { onMounted, ref, computed } from 'vue'
 import { useRouter } from 'vue-router'
 import { useAuthStore } from './stores/auth.store.ts'
 import { useCartStore } from './stores/cart.store.ts'
 import { ElMessage, ElMessageBox } from 'element-plus'
 import AdvancedSearch from './components/common/AdvancedSearch.vue'
 import MobileBottomNav from './components/common/MobileBottomNav.vue'
 import {
   ShoppingCart,
   User,
   Document,
   SwitchButton,
   Search,
   Grid,
   ArrowDown,
   Star,
   Setting,
   Menu
 } from '@element-plus/icons-vue'

 const router = useRouter()
 const authStore = useAuthStore()
 const cartStore = useCartStore()

 // Reactive data for enhanced navigation
 const showCategories = ref(false)
 const showMobileMenu = ref(false)
 const currency = ref('kes')
 const userAvatar = ref('')

 // Computed properties
 const userInitial = computed(() => {
   return authStore.userName ? authStore.userName.charAt(0).toUpperCase() : 'U'
 })

// Navigation methods
const handleAdvancedSearch = (query) => {
  // This method receives search queries from the AdvancedSearch component
  console.log('Advanced search for:', query)
  // The AdvancedSearch component already handles navigation
}

const toggleCategories = () => {
  showCategories.value = !showCategories.value
}

const toggleMobileMenu = () => {
  showMobileMenu.value = !showMobileMenu.value
}

const closeMobileMenu = () => {
  showMobileMenu.value = false
}

const handleUserMenuToggle = (visible) => {
  // Handle user menu visibility changes
  console.log('User menu visibility:', visible)
}

// Navigation helpers
const navigateToProfile = () => {
  closeMobileMenu()
  router.push('/profile')
}

const navigateToOrders = () => {
  closeMobileMenu()
  router.push('/orders')
}

const navigateToLogin = () => {
  closeMobileMenu()
  router.push('/login')
}

const navigateToRegister = () => {
  closeMobileMenu()
  router.push('/register')
}

const handleLogout = async () => {
  try {
    await ElMessageBox.confirm(
      'Are you sure you want to logout?',
      'Confirm Logout',
      {
        confirmButtonText: 'Logout',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
    )

    authStore.logout()
    ElMessage.success('Logged out successfully!')
    closeMobileMenu()

    // Redirect to home page
    router.push('/')
  } catch {
    // User cancelled logout
  }
}

onMounted(() => {
  // Initialize auth state
  authStore.initializeAuth()

  // Load cart if user is authenticated
  if (authStore.isAuthenticated) {
    cartStore.fetchCart().catch(console.error)
  }

  // Set user avatar (placeholder logic)
  userAvatar.value = '' // In a real app, this would be the user's profile picture URL
})
</script>

<style scoped>
 /* Enhanced Header Styles */
 .main-header {
   background: var(--white);
   box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
   position: sticky;
   top: 0;
   z-index: 1000;
 }

 /* Top Bar */
 .top-bar {
   background: var(--gray-100);
   border-bottom: 1px solid var(--gray-200);
   padding: 0.5rem 0;
   font-size: 0.875rem;
 }

 .top-bar-content {
   display: flex;
   justify-content: space-between;
   align-items: center;
 }

 .top-links {
   display: flex;
   gap: 1.5rem;
 }

 .top-link {
   color: var(--gray-600);
   text-decoration: none;
   transition: color var(--transition-fast);
 }

 .top-link:hover {
   color: var(--primary-color);
 }

 .welcome-text {
   color: var(--gray-700);
   font-weight: 500;
 }

 /* Main Navigation */
 .navbar {
   padding: 1rem 0;
   background: var(--white);
 }

 .navbar-content {
   display: flex;
   align-items: center;
   gap: 2rem;
 }

 /* Brand */
 .navbar-brand {
   flex-shrink: 0;
 }

 .brand-link {
   display: flex;
   align-items: center;
   gap: 0.5rem;
   text-decoration: none;
   color: var(--gray-800);
 }

 .brand-logo {
   width: 32px;
   height: 32px;
 }

 .brand-text {
   font-size: 1.75rem;
   font-weight: 800;
   background: linear-gradient(135deg, var(--primary-color), #667eea);
   -webkit-background-clip: text;
   -webkit-text-fill-color: transparent;
   background-clip: text;
 }

 /* Search Section */
 .search-section {
   flex: 1;
   max-width: 600px;
   margin: 0 2rem;
 }

 .search-container {
   position: relative;
 }

 .search-input {
   width: 100%;
 }

 .search-input :deep(.el-input__wrapper) {
   border-radius: 50px;
   box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
   border: 2px solid transparent;
   transition: all var(--transition-fast);
 }

 .search-input :deep(.el-input__wrapper:hover) {
   border-color: var(--primary-color);
   box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
 }

 .search-input :deep(.el-input__wrapper.is-focus) {
   border-color: var(--primary-color);
   box-shadow: 0 0 0 3px rgba(64, 158, 255, 0.1);
 }

 /* Navigation Menu */
 .navbar-right {
   display: flex;
   align-items: center;
   gap: 1.5rem;
 }

 .desktop-menu {
   display: flex;
   align-items: center;
 }

 .nav-links {
   display: flex;
   list-style: none;
   margin: 0;
   padding: 0;
   gap: 0;
 }

 .nav-item {
   position: relative;
 }

 .nav-link {
   display: flex;
   align-items: center;
   gap: 0.5rem;
   padding: 0.75rem 1rem;
   color: var(--gray-700);
   text-decoration: none;
   font-weight: 500;
   border-radius: 8px;
   transition: all var(--transition-fast);
   position: relative;
 }

 .nav-link:hover {
   color: var(--primary-color);
   background: var(--gray-50);
 }

 .nav-link.router-link-active {
   color: var(--primary-color);
   background: var(--primary-color-light);
 }

 .dropdown-arrow {
   font-size: 0.75rem;
   transition: transform var(--transition-fast);
 }

 .nav-item:hover .dropdown-arrow {
   transform: rotate(180deg);
 }

 .featured-link {
   position: relative;
 }

 .featured-link::after {
   content: '';
   position: absolute;
   top: 0.5rem;
   right: 0.5rem;
   width: 8px;
   height: 8px;
   background: var(--danger-color);
   border-radius: 50%;
 }

 .admin-link {
   background: linear-gradient(135deg, #667eea, #764ba2);
   color: var(--white) !important;
 }

 .admin-link:hover {
   background: linear-gradient(135deg, #764ba2, #667eea);
   color: var(--white) !important;
 }

 /* Mega Menu */
 .mega-menu {
   position: absolute;
   top: 100%;
   left: 0;
   right: 0;
   background: var(--white);
   border: 1px solid var(--gray-200);
   border-radius: 12px;
   box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
   padding: 2rem;
   margin-top: 0.5rem;
   z-index: 1000;
   animation: fadeInUp 0.3s ease-out;
 }

 .mega-menu-content {
   display: grid;
   grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
   gap: 2rem;
 }

 .category-section h4 {
   color: var(--gray-800);
   margin-bottom: 1rem;
   font-weight: 600;
   border-bottom: 2px solid var(--primary-color);
   padding-bottom: 0.5rem;
 }

 .category-section ul {
   list-style: none;
   padding: 0;
   margin: 0;
 }

 .category-section ul li {
   margin-bottom: 0.5rem;
 }

 .category-section a {
   color: var(--gray-600);
   text-decoration: none;
   padding: 0.5rem 0;
   display: block;
   border-radius: 6px;
   transition: all var(--transition-fast);
 }

 .category-section a:hover {
   color: var(--primary-color);
   background: var(--gray-50);
   padding-left: 0.5rem;
 }

 /* User Actions */
 .navbar-actions {
   display: flex;
   align-items: center;
   gap: 1rem;
 }

 .cart-link {
   position: relative;
   transition: transform var(--transition-fast);
 }

 .cart-link:hover {
   transform: scale(1.05);
 }

 .cart-link.has-items::after {
   content: '';
   position: absolute;
   top: -2px;
   right: -2px;
   width: 12px;
   height: 12px;
   background: var(--success-color);
   border-radius: 50%;
   border: 2px solid var(--white);
 }

 .cart-badge :deep(.el-badge__content) {
   background: var(--primary-color) !important;
 }

 .cart-button {
   border: 2px solid var(--gray-200);
   transition: all var(--transition-fast);
 }

 .cart-button:hover {
   border-color: var(--primary-color);
   background: var(--primary-color-light);
 }

 .user-menu {
   position: relative;
 }

 .user-button {
   border: 2px solid var(--gray-200);
   transition: all var(--transition-fast);
 }

 .user-button:hover {
   border-color: var(--primary-color);
   background: var(--primary-color-light);
 }

 .user-avatar {
   background: linear-gradient(135deg, var(--primary-color), #667eea);
   color: var(--white);
   font-weight: 600;
 }

 .user-dropdown {
   min-width: 200px;
 }

 .user-info {
   padding: 1rem;
   border-bottom: 1px solid var(--gray-200);
   text-align: center;
   background: var(--gray-50);
 }

 .user-info strong {
   display: block;
   color: var(--gray-800);
   margin-bottom: 0.25rem;
 }

 .user-info small {
   color: var(--gray-600);
 }

 .auth-buttons {
   display: flex;
   gap: 0.5rem;
 }

 .mobile-menu-toggle {
   display: none;
   border: 2px solid var(--gray-200);
   transition: all var(--transition-fast);
 }

 .mobile-menu-toggle:hover {
   border-color: var(--primary-color);
   background: var(--primary-color-light);
 }

 /* Mobile Menu */
 .mobile-menu {
   position: absolute;
   top: 100%;
   left: 0;
   right: 0;
   background: var(--white);
   border: 1px solid var(--gray-200);
   border-radius: 0 0 12px 12px;
   box-shadow: 0 10px 40px rgba(0, 0, 0, 0.15);
   padding: 1rem 0;
   z-index: 1000;
   animation: fadeInDown 0.3s ease-out;
 }

 .mobile-menu-content {
   max-height: 70vh;
   overflow-y: auto;
 }

 .mobile-nav-links {
   list-style: none;
   padding: 0;
   margin: 0;
 }

 .mobile-nav-links li {
   border-bottom: 1px solid var(--gray-100);
 }

 .mobile-nav-link {
   display: flex;
   align-items: center;
   gap: 1rem;
   padding: 1rem 1.5rem;
   color: var(--gray-700);
   text-decoration: none;
   transition: all var(--transition-fast);
 }

 .mobile-nav-link:hover {
   color: var(--primary-color);
   background: var(--gray-50);
 }

 .mobile-nav-link.router-link-active {
   color: var(--primary-color);
   background: var(--primary-color-light);
 }

 .mobile-user-section {
   padding: 1.5rem;
   border-top: 1px solid var(--gray-200);
   background: var(--gray-50);
 }

 .mobile-user-info {
   display: flex;
   align-items: center;
   gap: 1rem;
   margin-bottom: 1rem;
 }

 .mobile-user-avatar {
   background: linear-gradient(135deg, var(--primary-color), #667eea);
   color: var(--white);
   font-weight: 600;
 }

 .mobile-user-actions {
   display: flex;
   flex-direction: column;
   gap: 0.5rem;
 }

 .mobile-auth-section {
   padding: 1.5rem;
   display: flex;
   flex-direction: column;
   gap: 1rem;
 }

 /* Animations */
 @keyframes fadeInUp {
   from {
     opacity: 0;
     transform: translateY(20px);
   }
   to {
     opacity: 1;
     transform: translateY(0);
   }
 }

 @keyframes fadeInDown {
   from {
     opacity: 0;
     transform: translateY(-20px);
   }
   to {
     opacity: 1;
     transform: translateY(0);
   }
 }

/* Main Content */
.main-content {
  min-height: calc(100vh - 200px);
}

/* Footer Styles */
.main-footer {
  background: linear-gradient(135deg, #2d3748, #1a202c);
  color: var(--white);
  padding: 3rem 0 1rem;
  margin-top: auto;
  position: relative;
  box-shadow: 0 -4px 20px rgba(0, 0, 0, 0.1);
}

/* Ensure text visibility if background becomes white */
.main-footer:not([style*="background: white"]):not([style*="background: #fff"]) {
  color: var(--white);
}

.main-footer[style*="background: white"],
.main-footer[style*="background: #fff"] {
  background: var(--white) !important;
  color: var(--gray-800) !important;
}

.main-footer[style*="background: white"] .footer-section h4,
.main-footer[style*="background: #fff"] .footer-section h4 {
  color: var(--gray-800) !important;
}

.main-footer[style*="background: white"] .footer-section a,
.main-footer[style*="background: #fff"] .footer-section a {
  color: var(--gray-600) !important;
}

.main-footer[style*="background: white"] .footer-section a:hover,
.main-footer[style*="background: #fff"] .footer-section a:hover {
  color: var(--primary-color) !important;
  background: var(--gray-100) !important;
}

.main-footer[style*="background: white"] .footer-bottom,
.main-footer[style*="background: #fff"] .footer-bottom {
  color: var(--gray-500) !important;
  border-top-color: var(--gray-200) !important;
}

.main-footer::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, var(--primary-color), #667eea, var(--primary-color));
}

.footer-content {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
  margin-bottom: 2rem;
}

.footer-section h4 {
  color: var(--white);
  margin-bottom: 1rem;
  font-weight: 700;
  font-size: 1.125rem;
  position: relative;
  padding-bottom: 0.5rem;
}

.footer-section h4::after {
  content: '';
  position: absolute;
  bottom: 0;
  left: 0;
  width: 30px;
  height: 2px;
  background: var(--primary-color);
  border-radius: 2px;
}

.footer-section ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer-section ul li {
  margin-bottom: 0.75rem;
}

.footer-section a {
  color: #cbd5e0;
  text-decoration: none;
  transition: all var(--transition-fast);
  padding: 0.25rem 0;
  border-radius: 4px;
  display: inline-block;
}

.footer-section a:hover {
  color: var(--white);
  background: rgba(255, 255, 255, 0.1);
  padding-left: 0.5rem;
  transform: translateX(2px);
}

.footer-bottom {
  text-align: center;
  padding-top: 2rem;
  border-top: 1px solid rgba(255, 255, 255, 0.1);
  color: #a0aec0;
  font-size: 0.875rem;
}

/* Responsive Design */
@media (max-width: 1024px) {
  .search-section {
    margin: 0 1rem;
  }

  .desktop-menu {
    display: none;
  }

  .mobile-menu-toggle {
    display: flex;
  }

  .mega-menu {
    position: fixed;
    left: 1rem;
    right: 1rem;
    border-radius: 12px;
  }
}

@media (max-width: 768px) {
  .top-bar {
    display: none;
  }

  .navbar-content {
    gap: 1rem;
  }

  .search-section {
    display: none;
  }

  .navbar-right {
    gap: 0.5rem;
  }

  .brand-text {
    font-size: 1.5rem;
  }

  .cart-button,
  .user-button {
    width: 40px;
    height: 40px;
  }

  .auth-buttons {
    flex-direction: column;
    gap: 0.25rem;
  }

  .footer-content {
    grid-template-columns: 1fr;
    text-align: center;
  }
}

@media (max-width: 480px) {
  .navbar {
    padding: 0.75rem 0;
  }

  .brand-text {
    display: none;
  }

  .navbar-actions {
    gap: 0.25rem;
  }

  .cart-button,
  .user-button {
    width: 36px;
    height: 36px;
  }

  .mobile-menu-content {
    padding: 1rem;
  }

  .mobile-user-info {
    flex-direction: column;
    text-align: center;
    gap: 0.5rem;
  }
}
</style>