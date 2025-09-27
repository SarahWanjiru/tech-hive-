import { useAuthStore } from '../stores/auth.store.js'

/**
 * Authentication guard - requires user to be logged in
 */
export function requireAuth(to, from, next) {
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
export function requireGuest(to, from, next) {
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
export function requireAdmin(to, from, next) {
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
export function initializeAuth(to, from, next) {
  const authStore = useAuthStore()
  authStore.initializeAuth()
  next()
}