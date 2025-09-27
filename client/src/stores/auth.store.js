import { defineStore } from 'pinia'
import { authService } from '../services/auth.service.js'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    user: null,
    token: null,
    isAuthenticated: false,
    isLoading: false,
    error: null
  }),

  getters: {
    isAdmin: (state) => state.user?.role === 'admin',
    userName: (state) => state.user?.name || '',
    userEmail: (state) => state.user?.email || ''
  },

  actions: {
    /**
     * Initialize auth state from localStorage
     */
    initializeAuth() {
      const token = authService.getToken()
      const user = authService.getCurrentUser()

      if (token && user) {
        this.token = token
        this.user = user
        this.isAuthenticated = true
      }
    },

    /**
     * Login user
     * @param {Object} credentials - User credentials
     */
    async login(credentials) {
      this.isLoading = true
      this.error = null

      try {
        const response = await authService.login(credentials)

        this.token = response.data.token
        this.user = response.data.user
        this.isAuthenticated = true

        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.isLoading = false
      }
    },

    /**
     * Register new user
     * @param {Object} userData - User registration data
     */
    async register(userData) {
      this.isLoading = true
      this.error = null

      try {
        const response = await authService.register(userData)

        this.token = response.data.token
        this.user = response.data.user
        this.isAuthenticated = true

        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.isLoading = false
      }
    },

    /**
     * Logout user
     */
    logout() {
      authService.logout()
      this.token = null
      this.user = null
      this.isAuthenticated = false
      this.error = null
    },

    /**
     * Clear error state
     */
    clearError() {
      this.error = null
    },

    /**
     * Update user profile
     * @param {Object} userData - Updated user data
     */
    updateProfile(userData) {
      this.user = { ...this.user, ...userData }
      // Update localStorage as well
      localStorage.setItem('user_data', JSON.stringify(this.user))
    }
  }
})