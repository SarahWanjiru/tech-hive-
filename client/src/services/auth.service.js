import { apiClient } from '../api/client.js'

export const authService = {
  /**
   * Register a new user
   * @param {Object} userData - User registration data
   * @returns {Promise} API response
   */
  async register(userData) {
    try {
      const response = await apiClient.post('/v1/api/users', {
        name: userData.name,
        email: userData.email,
        password: userData.password,
        role: userData.role || 'customer'
      })

      if (response.data?.token) {
        this.setAuthData(response.data)
      }

      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Registration failed')
    }
  },

  /**
   * Login user
   * @param {Object} credentials - User login credentials
   * @returns {Promise} API response
   */
  async login(credentials) {
    try {
      const response = await apiClient.post('/v1/api/authentication', {
        username: credentials.email, // Backend expects email in username field
        password: credentials.password
      })

      if (response.data?.token) {
        this.setAuthData(response.data)
      }

      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Login failed')
    }
  },

  /**
   * Logout user
   */
  logout() {
    localStorage.removeItem('auth_token')
    localStorage.removeItem('user_data')
    window.location.href = '/login'
  },

  /**
   * Get current user data
   * @returns {Object|null} User data or null
   */
  getCurrentUser() {
    const userData = localStorage.getItem('user_data')
    return userData ? JSON.parse(userData) : null
  },

  /**
   * Get auth token
   * @returns {string|null} Auth token or null
   */
  getToken() {
    return localStorage.getItem('auth_token')
  },

  /**
   * Check if user is authenticated
   * @returns {boolean} Authentication status
   */
  isAuthenticated() {
    return !!this.getToken() && !!this.getCurrentUser()
  },

  /**
   * Check if current user is admin
   * @returns {boolean} Admin status
   */
  isAdmin() {
    const user = this.getCurrentUser()
    return user?.role === 'admin'
  },

  /**
   * Store authentication data
   * @param {Object} authData - Authentication response data
   */
  setAuthData(authData) {
    localStorage.setItem('auth_token', authData.token)
    localStorage.setItem('user_data', JSON.stringify(authData.user))
  },

  /**
   * Refresh user session
   * @returns {Promise} API response
   */
  async refreshSession() {
    try {
      const user = this.getCurrentUser()
      if (!user) return null

      // In a real app, you'd have a refresh token endpoint
      // For now, we'll just return the current user data
      return { data: { user } }
    } catch (error) {
      this.logout()
      throw error
    }
  }
}