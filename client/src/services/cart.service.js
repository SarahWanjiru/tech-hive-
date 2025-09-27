import { apiClient } from '../api/client.js'

export const cartService = {
  /**
   * Get user's cart
   * @returns {Promise} API response
   */
  async getCart() {
    try {
      const response = await apiClient.get('/v1/api/cart')
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to fetch cart')
    }
  },

  /**
   * Add item to cart
   * @param {string|number} productId - Product ID
   * @param {number} quantity - Quantity to add
   * @returns {Promise} API response
   */
  async addToCart(productId, quantity = 1) {
    try {
      const response = await apiClient.post('/v1/api/cart/items', {
        product_id: productId,
        quantity: quantity
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to add item to cart')
    }
  },

  /**
   * Update cart item quantity
   * @param {string|number} itemId - Cart item ID
   * @param {number} quantity - New quantity
   * @returns {Promise} API response
   */
  async updateCartItem(itemId, quantity) {
    try {
      const response = await apiClient.put(`/v1/api/cart/items/${itemId}`, {
        quantity: quantity
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to update cart item')
    }
  },

  /**
   * Remove item from cart
   * @param {string|number} itemId - Cart item ID
   * @returns {Promise} API response
   */
  async removeFromCart(itemId) {
    try {
      const response = await apiClient.delete(`/v1/api/cart/items/${itemId}`)
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to remove item from cart')
    }
  },

  /**
   * Clear entire cart
   * @returns {Promise} API response
   */
  async clearCart() {
    try {
      const response = await apiClient.delete('/v1/api/cart')
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to clear cart')
    }
  },

  /**
   * Get cart summary (total items and amount)
   * @returns {Promise} Cart summary
   */
  async getCartSummary() {
    try {
      const cart = await this.getCart()
      const items = cart.data?.items || []
      const totalItems = items.reduce((sum, item) => sum + item.quantity, 0)
      const totalAmount = items.reduce((sum, item) => sum + (item.price * item.quantity), 0)

      return {
        totalItems,
        totalAmount,
        itemCount: items.length
      }
    } catch (error) {
      return {
        totalItems: 0,
        totalAmount: 0,
        itemCount: 0
      }
    }
  }
}