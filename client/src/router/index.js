import { createRouter, createWebHistory } from 'vue-router'
import {
  requireAuth,
  requireGuest,
  requireAdmin,
  initializeAuth
} from './guards.js'

// Import views (we'll create these next)
import Home from '../views/Home.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'
import Products from '../views/Products.vue'
import ProductDetail from '../views/ProductDetail.vue'
import Cart from '../views/Cart.vue'
import Checkout from '../views/Checkout.vue'
import Orders from '../views/Orders.vue'
import OrderDetail from '../views/OrderDetail.vue'
import Profile from '../views/Profile.vue'
import Admin from '../views/Admin.vue'

const routes = [
  {
    path: '/',
    name: 'home',
    component: Home,
    meta: {
      title: 'Home - Mini-Ecommerce'
    }
  },
  {
    path: '/login',
    name: 'login',
    component: Login,
    meta: {
      title: 'Login - Mini-Ecommerce',
      requiresGuest: true
    },
    beforeEnter: requireGuest
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
    meta: {
      title: 'Register - Mini-Ecommerce',
      requiresGuest: true
    },
    beforeEnter: requireGuest
  },
  {
    path: '/products',
    name: 'products',
    component: Products,
    meta: {
      title: 'Products - Mini-Ecommerce'
    }
  },
  {
    path: '/products/:id',
    name: 'product-detail',
    component: ProductDetail,
    meta: {
      title: 'Product Details - Mini-Ecommerce'
    },
    props: true
  },
  {
    path: '/cart',
    name: 'cart',
    component: Cart,
    meta: {
      title: 'Shopping Cart - Mini-Ecommerce',
      requiresAuth: true
    },
    beforeEnter: requireAuth
  },
  {
    path: '/checkout',
    name: 'checkout',
    component: Checkout,
    meta: {
      title: 'Checkout - Mini-Ecommerce',
      requiresAuth: true
    },
    beforeEnter: requireAuth
  },
  {
    path: '/orders',
    name: 'orders',
    component: Orders,
    meta: {
      title: 'My Orders - Mini-Ecommerce',
      requiresAuth: true
    },
    beforeEnter: requireAuth
  },
  {
    path: '/orders/:id',
    name: 'order-detail',
    component: OrderDetail,
    meta: {
      title: 'Order Details - Mini-Ecommerce',
      requiresAuth: true
    },
    props: true,
    beforeEnter: requireAuth
  },
  {
    path: '/profile',
    name: 'profile',
    component: Profile,
    meta: {
      title: 'My Profile - Mini-Ecommerce',
      requiresAuth: true
    },
    beforeEnter: requireAuth
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    meta: {
      title: 'Admin Dashboard - Mini-Ecommerce',
      requiresAdmin: true
    },
    beforeEnter: requireAdmin
  },
  // Catch all route - must be last
  {
    path: '/:pathMatch(.*)*',
    name: 'not-found',
    component: () => import('../views/NotFound.vue'),
    meta: {
      title: 'Page Not Found - Mini-Ecommerce'
    }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  }
})

// Global navigation guards
router.beforeEach(initializeAuth)

// Set page title
router.afterEach((to) => {
  document.title = to.meta.title || 'Mini-Ecommerce'
})

export default router