import { defineStore } from 'pinia'
import { authService } from '../services/auth.service'
import type { User, LoginRequest, RegisterRequest, GeneralResponse } from '../types/api'

interface AuthState {
  user: User | null
  token: string | null
  isAuthenticated: boolean
  isLoading: boolean
  error: string | null
}

export const useAuthStore = defineStore('auth', {
  state: (): AuthState => ({
    user: null,
    token: null,
    isAuthenticated: false,
    isLoading: false,
    error: null
  }),

  getters: {
    isAdmin: (state: AuthState): boolean => state.user?.role === 'admin',
    userName: (state: AuthState): string => state.user?.name || '',
    userEmail: (state: AuthState): string => state.user?.email || ''
  },

  actions: {
    /**
      * Initialize auth state from localStorage
      */
    initializeAuth(): void {
      try {
        // First, clean up any temporary bypass data
        this.clearTemporaryBypassData()

        const token = authService.getToken()
        const user = authService.getCurrentUser()

        console.log('InitializeAuth - Token:', token ? 'present' : 'missing')
        console.log('InitializeAuth - User:', user ? 'present' : 'missing')

        if (token && user) {
          this.token = token
          this.user = user
          this.isAuthenticated = true

          // Debug logging for admin role
          console.log('Auth initialized successfully:', {
            user: user,
            role: user.role,
            isAdmin: user.role === 'admin'
          })
        } else {
          console.log('No valid auth data found - clearing state')
          this.clearAuthState()
        }
      } catch (error) {
        console.error('Error initializing auth:', error)
        this.clearAuthState()
      }
    },

    /**
     * Clear authentication state
     */
    clearAuthState(): void {
      this.token = null
      this.user = null
      this.isAuthenticated = false
      this.error = null
    },

    /**
     * Login user
     * @param credentials - User credentials
     */
    async login(credentials: LoginRequest): Promise<GeneralResponse> {
      this.isLoading = true
      this.error = null

      try {
        const response = await authService.login(credentials)

        this.token = response.data.token
        this.user = response.data.user
        this.isAuthenticated = true

        // Debug logging for role
        console.log('User logged in:', {
          user: response.data.user,
          role: response.data.user.role,
          isAdmin: response.data.user.role === 'admin'
        })

        // Auto-redirect based on role
        if (response.data.user.role === 'admin') {
          // Import router dynamically to avoid circular dependency
          const router = await import('../router')
          router.default.push('/admin')
        }

        return response
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.isLoading = false
      }
    },

    /**
     * Register new user
     * @param userData - User registration data
     */
    async register(userData: RegisterRequest): Promise<GeneralResponse> {
      this.isLoading = true
      this.error = null

      try {
        const response = await authService.register(userData)

        this.token = response.data.token
        this.user = response.data.user
        this.isAuthenticated = true

        // Auto-redirect based on role (for admin registrations)
        if (response.data.user.role === 'admin') {
          const router = await import('../router')
          router.default.push('/admin')
        }

        return response
      } catch (error: any) {
        this.error = error.message
        throw error
      } finally {
        this.isLoading = false
      }
    },

    /**
     * Logout user
     */
    logout(): void {
      authService.logout()
      this.clearAuthState()
    },

    /**
     * Clear error state
     */
    clearError(): void {
      this.error = null
    },

    /**
     * Update user profile
     * @param userData - Updated user data
     */
    updateProfile(userData: Partial<User>): void {
      if (this.user) {
        this.user = { ...this.user, ...userData }
        // Update localStorage as well
        localStorage.setItem('user_data', JSON.stringify(this.user))
      }
    },

    /**
     * Refresh authentication state
     */
    refreshAuth(): void {
      this.initializeAuth()
    },

    /**
     * Force set admin status (for debugging)
     */
    setAdminStatus(isAdmin: boolean): void {
      if (this.user) {
        this.user.role = isAdmin ? 'admin' : 'customer'
        localStorage.setItem('user_data', JSON.stringify(this.user))
      }
    },

    /**
      * Disable temporary bypass (for production)
      */
    disableTemporaryBypass(): void {
      // Clear any bypass data
      if (this.token === 'temp-admin-token') {
        this.clearAuthState()
        localStorage.removeItem('auth_token')
        localStorage.removeItem('user_data')
      }
    },

    /**
      * Clear any temporary bypass data from localStorage
      */
    clearTemporaryBypassData(): void {
      try {
        const token = localStorage.getItem('auth_token')
        const userData = localStorage.getItem('user_data')

        // Check if this is temporary bypass data
        if (token && userData) {
          const user = JSON.parse(userData)

          // If this is the temporary admin user from bypass, clear it
          if (user && user.email === 'admin@example.com' && user.name === 'Admin User' && token.startsWith('temp-admin-token')) {
            console.log('🧹 Clearing temporary bypass data from localStorage')
            localStorage.removeItem('auth_token')
            localStorage.removeItem('user_data')
          }
        }
      } catch (error) {
        console.error('Error clearing temporary bypass data:', error)
        // If there's any error parsing, clear the data to be safe
        localStorage.removeItem('auth_token')
        localStorage.removeItem('user_data')
      }
    }
  }
})