import { defineStore } from 'pinia'
import { cartService } from '../services/cart.service.js'

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: [],
    loading: false,
    error: null,
    lastUpdated: null
  }),

  getters: {
    totalItems: (state) => {
      return state.items.reduce((total, item) => total + item.quantity, 0)
    },

    totalAmount: (state) => {
      return state.items.reduce((total, item) => total + (item.price * item.quantity), 0)
    },

    isEmpty: (state) => {
      return state.items.length === 0
    },

    itemCount: (state) => {
      return state.items.length
    },

    hasItems: (state) => {
      return state.items.length > 0
    }
  },

  actions: {
    /**
     * Fetch user's cart
     */
    async fetchCart() {
      this.loading = true
      this.error = null

      try {
        const response = await cartService.getCart()
        this.items = response.data?.items || []
        this.lastUpdated = new Date()
        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Add item to cart
     * @param {string|number} productId - Product ID
     * @param {number} quantity - Quantity to add
     */
    async addToCart(productId, quantity = 1) {
      this.loading = true
      this.error = null

      try {
        const response = await cartService.addToCart(productId, quantity)

        // Refresh cart data
        await this.fetchCart()

        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Update cart item quantity
     * @param {string|number} itemId - Cart item ID
     * @param {number} quantity - New quantity
     */
    async updateCartItem(itemId, quantity) {
      this.loading = true
      this.error = null

      try {
        const response = await cartService.updateCartItem(itemId, quantity)

        // Update local state immediately for better UX
        const itemIndex = this.items.findIndex(item => item.id === itemId)
        if (itemIndex !== -1) {
          if (quantity <= 0) {
            this.items.splice(itemIndex, 1)
          } else {
            this.items[itemIndex].quantity = quantity
          }
        }

        this.lastUpdated = new Date()
        return response
      } catch (error) {
        this.error = error.message
        // Refresh cart to get correct state
        await this.fetchCart()
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Remove item from cart
     * @param {string|number} itemId - Cart item ID
     */
    async removeFromCart(itemId) {
      this.loading = true
      this.error = null

      try {
        const response = await cartService.removeFromCart(itemId)

        // Remove from local state immediately
        this.items = this.items.filter(item => item.id !== itemId)
        this.lastUpdated = new Date()

        return response
      } catch (error) {
        this.error = error.message
        // Refresh cart to get correct state
        await this.fetchCart()
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Clear entire cart
     */
    async clearCart() {
      this.loading = true
      this.error = null

      try {
        const response = await cartService.clearCart()

        // Clear local state
        this.items = []
        this.lastUpdated = new Date()

        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Get cart summary
     */
    async getCartSummary() {
      try {
        const summary = await cartService.getCartSummary()
        return summary
      } catch (error) {
        return {
          totalItems: this.totalItems,
          totalAmount: this.totalAmount,
          itemCount: this.itemCount
        }
      }
    },

    /**
     * Check if product is in cart
     * @param {string|number} productId - Product ID
     * @returns {Object|null} Cart item or null
     */
    getCartItem(productId) {
      return this.items.find(item => item.product_id === productId) || null
    },

    /**
     * Get item quantity in cart
     * @param {string|number} productId - Product ID
     * @returns {number} Quantity in cart
     */
    getItemQuantity(productId) {
      const item = this.getCartItem(productId)
      return item ? item.quantity : 0
    },

    /**
     * Clear error state
     */
    clearError() {
      this.error = null
    },

    /**
     * Refresh cart data
     */
    async refreshCart() {
      await this.fetchCart()
    }
  }
})