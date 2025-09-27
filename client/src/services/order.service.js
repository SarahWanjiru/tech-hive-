import { apiClient } from '../api/client.js'

export const orderService = {
  /**
   * Create new order from cart
   * @param {Object} orderData - Order data
   * @returns {Promise} API response
   */
  async createOrder(orderData) {
    try {
      const response = await apiClient.post('/v1/api/orders', {
        cart_id: orderData.cartId,
        shipping_address: orderData.shippingAddress
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to create order')
    }
  },

  /**
   * Get user's orders
   * @param {Object} params - Query parameters
   * @returns {Promise} API response
   */
  async getOrders(params = {}) {
    try {
      const response = await apiClient.get('/v1/api/orders', { params })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to fetch orders')
    }
  },

  /**
   * Get order by ID
   * @param {string|number} id - Order ID
   * @returns {Promise} API response
   */
  async getOrderById(id) {
    try {
      const response = await apiClient.get(`/v1/api/orders/${id}`)
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to fetch order')
    }
  },

  /**
   * Cancel order
   * @param {string|number} id - Order ID
   * @returns {Promise} API response
   */
  async cancelOrder(id) {
    try {
      const response = await apiClient.delete(`/v1/api/orders/${id}`)
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to cancel order')
    }
  },

  /**
   * Update order status (Admin only)
   * @param {string|number} id - Order ID
   * @param {string} status - New status
   * @returns {Promise} API response
   */
  async updateOrderStatus(id, status) {
    try {
      const response = await apiClient.put(`/v1/api/orders/${id}/status`, {
        status: status
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to update order status')
    }
  },

  /**
   * Get order statistics (Admin only)
   * @returns {Promise} API response
   */
  async getOrderStats() {
    try {
      const response = await apiClient.get('/v1/api/orders/stats')
      return response
    } catch (error) {
      // Return mock data for now
      return {
        data: {
          totalOrders: 0,
          pendingOrders: 0,
          completedOrders: 0,
          cancelledOrders: 0,
          totalRevenue: 0
        }
      }
    }
  }
}