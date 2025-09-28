import { defineStore } from 'pinia'
import { orderService } from '../services/order.service'
import type { Order, CreateOrderModel, UpdateOrderStatusModel } from '../types/api'

interface OrderState {
  orders: Order[]
  currentOrder: Order | null
  loading: boolean
  error: string | null
  stats: {
    totalOrders: number
    pendingOrders: number
    completedOrders: number
    totalRevenue: number
  }
}

export const useOrderStore = defineStore('order', {
  state: (): OrderState => ({
    orders: [],
    currentOrder: null,
    loading: false,
    error: null,
    stats: {
      totalOrders: 0,
      pendingOrders: 0,
      completedOrders: 0,
      totalRevenue: 0
    }
  }),

  getters: {
    // Get recent orders (last 5)
    recentOrders: (state: OrderState): Order[] => {
      return state.orders
        .sort((a, b) => new Date(b.created_at).getTime() - new Date(a.created_at).getTime())
        .slice(0, 5)
    },

    // Get orders by status
    getOrdersByStatus: (state: OrderState) => (status: string): Order[] => {
      return state.orders.filter(order => order.status === status)
    },

    // Get pending orders count
    pendingOrdersCount: (state: OrderState): number => {
      return state.orders.filter(order => order.status === 'pending').length
    }
  },

  actions: {
    /**
     * Fetch user's orders
     */
    async fetchOrders(params?: { page?: number; limit?: number }): Promise<Order[]> {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.getOrders(params)
        this.orders = response.data
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch orders'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Fetch all orders (admin only)
     */
    async fetchAllOrders(params?: { page?: number; limit?: number }): Promise<{ data: Order[], total: number }> {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.getOrders(params)
        return {
          data: response.data,
          total: response.data.length
        }
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch orders'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Fetch single order by ID
     */
    async fetchOrderById(id: string): Promise<Order> {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.getOrderById(id)
        this.currentOrder = response.data
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch order'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Create new order
     */
    async createOrder(orderData: CreateOrderModel): Promise<Order> {
      if (!orderService.validateOrderData(orderData)) {
        throw new Error('Invalid order data')
      }

      this.loading = true
      this.error = null

      try {
        const response = await orderService.createOrder(orderData)
        this.orders.unshift(response.data)
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to create order'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Update order status (admin only)
     */
    async updateOrderStatus(id: string, status: string): Promise<Order> {
      this.loading = true
      this.error = null

      try {
        const statusData: UpdateOrderStatusModel = { status }
        const response = await orderService.updateOrderStatus(id, statusData)

        // Update in orders array
        const index = this.orders.findIndex(order => order.id === parseInt(id))
        if (index !== -1) {
          this.orders[index] = response.data
        }

        // Update current order if it's the same
        if (this.currentOrder?.id === parseInt(id)) {
          this.currentOrder = response.data
        }

        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to update order status'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Cancel order
     */
    async cancelOrder(id: string): Promise<Order> {
      this.loading = true
      this.error = null

      try {
        const response = await orderService.cancelOrder(id)

        // Update in orders array
        const index = this.orders.findIndex(order => order.id === parseInt(id))
        if (index !== -1) {
          this.orders[index] = response.data
        }

        // Update current order if it's the same
        if (this.currentOrder?.id === parseInt(id)) {
          this.currentOrder = response.data
        }

        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to cancel order'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Fetch order statistics (admin only)
     */
    async fetchOrderStats(): Promise<void> {
      try {
        const response = await orderService.getOrderStats()
        this.stats = response.data
      } catch (error: any) {
        console.error('Failed to fetch order stats:', error)
        // Don't throw error for stats, just log it
      }
    },

    /**
     * Process M-Pesa payment
     */
    async processMpesaPayment(paymentData: {
      order_id: number
      phone_number: string
      amount: number
    }): Promise<{ data: any }> {
      this.loading = true
      this.error = null

      try {
        // Use the payment service to process M-Pesa payment
        const { paymentService } = await import('../services/payment.service')

        const response = await paymentService.initiateSTKPush({
          phone_number: paymentData.phone_number,
          amount: paymentData.amount,
          order_id: paymentData.order_id
        })

        return response
      } catch (error: any) {
        this.error = error.message || 'Failed to process M-Pesa payment'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Clear error state
     */
    clearError(): void {
      this.error = null
    }
  }
})