import axios, { type AxiosInstance, type AxiosResponse, type AxiosError } from 'axios'

// Create axios instance with base configuration
const apiClient: AxiosInstance = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  timeout: parseInt(import.meta.env.VITE_API_TIMEOUT) || 10000,
  headers: {
    'Content-Type': 'application/json',
  },
})

// Request interceptor to add auth token
apiClient.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem('auth_token')
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error: AxiosError) => {
    return Promise.reject(error)
  }
)

// Response interceptor to handle common errors
apiClient.interceptors.response.use(
  (response: AxiosResponse) => {
    return response
  },
  (error: AxiosError) => {
    console.log('API Error:', {
      status: error.response?.status,
      url: error.config?.url,
      message: error.message
    })

    if (error.response?.status === 401) {
      // Only clear auth data and redirect if it's not an authentication endpoint
      // This prevents clearing auth data during login attempts
      const url = error.config?.url || ''
      if (!url.includes('/authentication') && !url.includes('/users')) {
        console.log('401 Unauthorized on protected endpoint - clearing auth data and redirecting to login')
        localStorage.removeItem('auth_token')
        localStorage.removeItem('user_data')
        window.location.href = '/login'
      } else {
        console.log('401 Unauthorized on auth endpoint - not clearing auth data')
      }
    }

    // Special handling for network errors that might be caused by CORS or server issues
    if (!error.response && error.message.includes('Network Error')) {
      console.log('Network error detected - not clearing auth data')
      // Don't clear auth data for network errors
      throw new Error('Network error - please check your connection or try refreshing the page')
    }

    // Handle network errors
    if (!error.response) {
      console.error('Network error:', error)
      throw new Error('Network error - please check your connection')
    }

    // Handle API errors
    const message = (error.response?.data as any)?.message || error.message || 'An error occurred'
    throw new Error(message)
  }
)

export { apiClient }
export default apiClient