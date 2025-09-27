<template>
  <div id="app">
    <!-- Navigation Header -->
    <header class="main-header">
      <nav class="navbar">
        <div class="container">
          <div class="navbar-content">
            <!-- Logo -->
            <div class="navbar-brand">
              <router-link to="/" class="brand-link">
                <h2>Mini-Ecommerce</h2>
              </router-link>
            </div>

            <!-- Navigation Menu -->
            <div class="navbar-menu">
              <ul class="nav-links">
                <li>
                  <router-link to="/" class="nav-link">
                    Home
                  </router-link>
                </li>
                <li>
                  <router-link to="/products" class="nav-link">
                    Products
                  </router-link>
                </li>
                <li v-if="authStore.isAdmin">
                  <router-link to="/admin" class="nav-link">
                    Admin
                  </router-link>
                </li>
              </ul>
            </div>

            <!-- User Actions -->
            <div class="navbar-actions">
              <!-- Cart Icon -->
              <router-link
                v-if="authStore.isAuthenticated"
                to="/cart"
                class="cart-link"
              >
                <el-badge :value="cartStore.totalItems" :hidden="cartStore.totalItems === 0">
                  <el-button circle>
                    <el-icon><ShoppingCart /></el-icon>
                  </el-button>
                </el-badge>
              </router-link>

              <!-- User Menu -->
              <div v-if="authStore.isAuthenticated" class="user-menu">
                <el-dropdown>
                  <el-button circle>
                    <el-icon><User /></el-icon>
                  </el-button>
                  <template #dropdown>
                    <el-dropdown-menu>
                      <el-dropdown-item @click="$router.push('/profile')">
                        <el-icon><User /></el-icon>
                        Profile
                      </el-dropdown-item>
                      <el-dropdown-item @click="$router.push('/orders')">
                        <el-icon><Document /></el-icon>
                        My Orders
                      </el-dropdown-item>
                      <el-dropdown-item @click="handleLogout" divided>
                        <el-icon><SwitchButton /></el-icon>
                        Logout
                      </el-dropdown-item>
                    </el-dropdown-menu>
                  </template>
                </el-dropdown>
              </div>

              <!-- Auth Buttons -->
              <div v-else class="auth-buttons">
                <el-button @click="$router.push('/login')">
                  Login
                </el-button>
                <el-button type="primary" @click="$router.push('/register')">
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

    <!-- Footer -->
    <footer class="main-footer">
      <div class="container">
        <div class="footer-content">
          <div class="footer-section">
            <h4>Mini-Ecommerce</h4>
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
          <p>&copy; 2024 Mini-Ecommerce. All rights reserved.</p>
        </div>
      </div>
    </footer>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from './stores/auth.store.ts'
import { useCartStore } from './stores/cart.store.ts'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ShoppingCart, User, Document, SwitchButton } from '@element-plus/icons-vue'

const router = useRouter()
const authStore = useAuthStore()
const cartStore = useCartStore()

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
})
</script>

<style scoped>
/* Header Styles */
.main-header {
  background: var(--white);
  box-shadow: var(--shadow-sm);
  position: sticky;
  top: 0;
  z-index: 1000;
}

.navbar {
  padding: 1rem 0;
}

.navbar-content {
  display: flex;
  align-items: center;
  justify-content: space-between;
}

.navbar-brand .brand-link {
  color: var(--primary-color);
  text-decoration: none;
}

.navbar-brand h2 {
  margin: 0;
  font-size: 1.5rem;
  font-weight: 700;
}

.navbar-menu {
  flex: 1;
  margin: 0 2rem;
}

.nav-links {
  display: flex;
  list-style: none;
  margin: 0;
  padding: 0;
  gap: 2rem;
}

.nav-link {
  color: var(--gray-700);
  text-decoration: none;
  font-weight: 500;
  transition: color var(--transition-fast);
}

.nav-link:hover,
.nav-link.router-link-active {
  color: var(--primary-color);
}

.navbar-actions {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.cart-link {
  color: var(--gray-700);
  transition: color var(--transition-fast);
}

.cart-link:hover {
  color: var(--primary-color);
}

.user-menu {
  display: flex;
  align-items: center;
}

.auth-buttons {
  display: flex;
  gap: 0.5rem;
}

/* Main Content */
.main-content {
  min-height: calc(100vh - 200px);
}

/* Footer Styles */
.main-footer {
  background: var(--gray-800);
  color: var(--white);
  padding: 3rem 0 1rem;
  margin-top: auto;
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
}

.footer-section ul {
  list-style: none;
  padding: 0;
  margin: 0;
}

.footer-section ul li {
  margin-bottom: 0.5rem;
}

.footer-section a {
  color: var(--gray-300);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.footer-section a:hover {
  color: var(--white);
}

.footer-bottom {
  text-align: center;
  padding-top: 2rem;
  border-top: 1px solid var(--gray-700);
  color: var(--gray-300);
}

/* Responsive Design */
@media (max-width: 768px) {
  .navbar-content {
    flex-direction: column;
    gap: 1rem;
  }

  .navbar-menu {
    margin: 0;
  }

  .nav-links {
    gap: 1rem;
  }

  .footer-content {
    grid-template-columns: 1fr;
    text-align: center;
  }
}
</style>