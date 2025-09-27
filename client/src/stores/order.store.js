import { defineStore } from 'pinia'
import { orderService } from '../services/order.service.js'

export const useOrderStore = defineStore('order', {
  state: () => ({
    orders: [],
    currentOrder: null,
    loading: false,
    error: null,
    pagination: {
      page: 1,
      limit: 10,
      total: 0,
      totalPages: 0
    },
    stats: {
      totalOrders: 0,
      pendingOrders: 0,
      completedOrders: 0,
      cancelledOrders: 0,
      totalRevenue: 0
    }
  }),

  getters: {
    recentOrders: (state) => {
      return state.orders.slice(0, 5)
    },

    pendingOrders: (state) => {
      return state.orders.filter(order => order.status === 'pending')
    },

    completedOrders: (state) => {
      return state.orders.filter(order => order.status === 'completed')
    },

    cancelledOrders: (state) => {
      return state.orders.filter(order => order.status === 'cancelled')
    }
  },

  actions: {
    /**
     * Fetch user's orders
     * @param {Object} params - Query parameters
     */
    async fetchOrders(params = {}) {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.getOrders(params)
        this.orders = response.data.orders || response.data || []

        if (response.data.pagination) {
          this.pagination = { ...this.pagination, ...response.data.pagination }
        }

        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Get order by ID
     * @param {string|number} id - Order ID
     */
    async fetchOrderById(id) {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.getOrderById(id)
        this.currentOrder = response.data
        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Create new order
     * @param {Object} orderData - Order data
     */
    async createOrder(orderData) {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.createOrder(orderData)

        // Add to orders list
        this.orders.unshift(response.data)

        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Cancel order
     * @param {string|number} id - Order ID
     */
    async cancelOrder(id) {
      try {
        const response = await orderService.cancelOrder(id)

        // Update order status in local state
        const orderIndex = this.orders.findIndex(order => order.id === id)
        if (orderIndex !== -1) {
          this.orders[orderIndex].status = 'cancelled'
        }

        // Update current order if it's the same
        if (this.currentOrder?.id === id) {
          this.currentOrder.status = 'cancelled'
        }

        return response
      } catch (error) {
        this.error = error.message
        throw error
      }
    },

    /**
     * Update order status (Admin only)
     * @param {string|number} id - Order ID
     * @param {string} status - New status
     */
    async updateOrderStatus(id, status) {
      try {
        const response = await orderService.updateOrderStatus(id, status)

        // Update order status in local state
        const orderIndex = this.orders.findIndex(order => order.id === id)
        if (orderIndex !== -1) {
          this.orders[orderIndex].status = status
        }

        // Update current order if it's the same
        if (this.currentOrder?.id === id) {
          this.currentOrder.status = status
        }

        return response
      } catch (error) {
        this.error = error.message
        throw error
      }
    },

    /**
     * Get order statistics (Admin only)
     */
    async fetchOrderStats() {
      try {
        const response = await orderService.getOrderStats()
        this.stats = { ...this.stats, ...response.data }
        return response
      } catch (error) {
        // Keep default stats on error
        console.warn('Failed to fetch order stats:', error.message)
        return { data: this.stats }
      }
    },

    /**
     * Clear current order
     */
    clearCurrentOrder() {
      this.currentOrder = null
    },

    /**
     * Clear error state
     */
    clearError() {
      this.error = null
    }
  }
})