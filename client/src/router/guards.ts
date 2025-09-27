import { useAuthStore } from '../stores/auth.store'
import type { NavigationGuard, RouteLocationNormalized } from 'vue-router'

/**
 * Authentication guard - requires user to be logged in
 */
export const requireAuth: NavigationGuard = (to, from, next) => {
  const authStore = useAuthStore()

  if (!authStore.isAuthenticated) {
    next({
      name: 'login',
      query: { redirect: to.fullPath }
    })
  } else {
    next()
  }
}

/**
 * Guest guard - requires user to NOT be logged in
 */
export const requireGuest: NavigationGuard = (to, from, next) => {
  const authStore = useAuthStore()

  if (authStore.isAuthenticated) {
    next({ name: 'home' })
  } else {
    next()
  }
}

/**
 * Admin guard - requires user to be admin
 */
export const requireAdmin: NavigationGuard = (to, from, next) => {
  const authStore = useAuthStore()

  if (!authStore.isAuthenticated) {
    next({
      name: 'login',
      query: { redirect: to.fullPath }
    })
  } else if (!authStore.isAdmin) {
    next({ name: 'home' })
  } else {
    next()
  }
}

/**
 * Initialize auth state before navigation
 */
export const initializeAuth: NavigationGuard = (to, from, next) => {
  const authStore = useAuthStore()
  authStore.initializeAuth()
  next()
}