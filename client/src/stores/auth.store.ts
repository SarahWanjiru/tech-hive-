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
      this.token = null
      this.user = null
      this.isAuthenticated = false
      this.error = null
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
    }
  }
})