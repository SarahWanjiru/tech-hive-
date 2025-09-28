import { apiClient } from '../api/client'
import type { User, LoginRequest, RegisterRequest, GeneralResponse } from '../types/api'

export const authService = {
  /**
   * Register a new user
   * @param userData - User registration data
   * @returns API response
   */
  async register(userData: RegisterRequest): Promise<GeneralResponse> {
    try {
      const response = await apiClient.post<GeneralResponse>('/v1/api/users', {
        name: userData.name,
        email: userData.email,
        password: userData.password,
        role: userData.role || 'customer'
      })

      if (response.data?.data?.token) {
        this.setAuthData(response.data.data)
      }

      return response.data
    } catch (error: any) {
      throw new Error(error.response?.data?.message || 'Registration failed')
    }
  },

  /**
   * Login user
   * @param credentials - User login credentials
   * @returns API response
   */
  async login(credentials: LoginRequest): Promise<GeneralResponse> {
    try {
      const response = await apiClient.post<GeneralResponse>('/v1/api/authentication', {
        email: credentials.email,
        password: credentials.password
      })

      if (response.data?.data?.token) {
        this.setAuthData(response.data.data)
      }

      return response.data
    } catch (error: any) {
      throw new Error(error.response?.data?.message || 'Login failed')
    }
  },

  /**
   * Logout user
   */
  logout(): void {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
    window.location.href = '/login'
  },

  /**
   * Get current user data
   * @returns User data or null
   */
  getCurrentUser(): User | null {
    try {
      const userData = localStorage.getItem('user_data')
      return userData && userData !== 'undefined' ? JSON.parse(userData) : null
    } catch (error) {
      console.error('Error parsing user data:', error)
      localStorage.removeItem('user_data')
      return null
    }
  },

  /**
   * Get auth token
   * @returns Auth token or null
   */
  getToken(): string | null {
    return localStorage.getItem('auth_token')
  },

  /**
   * Check if user is authenticated
   * @returns Authentication status
   */
  isAuthenticated(): boolean {
    return !!this.getToken() && !!this.getCurrentUser()
  },

  /**
   * Check if current user is admin
   * @returns Admin status
   */
  isAdmin(): boolean {
    const user = this.getCurrentUser()
    return user?.role === 'admin'
  },

  /**
   * Store authentication data
   * @param authData - Authentication response data
   */
  setAuthData(authData: { token: string; user: User }): void {
    localStorage.setItem('auth_token', authData.token)
    localStorage.setItem('user_data', JSON.stringify(authData.user))
  },

  /**
   * Refresh user session
   * @returns API response
   */
  async refreshSession(): Promise<GeneralResponse | null> {
    try {
      const user = this.getCurrentUser()
      if (!user) return null

      // In a real app, you'd have a refresh token endpoint
      // For now, we'll just return the current user data
      return {
        code: 200,
        message: 'Success',
        data: { user }
      }
    } catch (error) {
      this.logout()
      throw error
    }
  }
}