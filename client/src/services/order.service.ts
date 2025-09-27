import { apiClient } from '../api/client'
import type {
  Order,
  CreateOrderModel,
  UpdateOrderStatusModel,
  GeneralResponse
} from '../types/api'

export class OrderService {
  /**
   * Get user's orders
   */
  async getOrders(params?: { page?: number; limit?: number }): Promise<GeneralResponse<Order[]>> {
    const queryParams = new URLSearchParams()
    if (params) {
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined) {
          queryParams.append(key, value.toString())
        }
      })
    }

    const response = await apiClient.get(`/orders?${queryParams.toString()}`)
    return response.data
  }

  /**
   * Get order by ID
   */
  async getOrderById(id: string): Promise<GeneralResponse<Order>> {
    const response = await apiClient.get(`/orders/${id}`)
    return response.data
  }

  /**
   * Create new order
   */
  async createOrder(orderData: CreateOrderModel): Promise<GeneralResponse<Order>> {
    const response = await apiClient.post('/orders', orderData)
    return response.data
  }

  /**
   * Update order status (admin only)
   */
  async updateOrderStatus(id: string, statusData: UpdateOrderStatusModel): Promise<GeneralResponse<Order>> {
    const response = await apiClient.put(`/orders/${id}/status`, statusData)
    return response.data
  }

  /**
   * Cancel order
   */
  async cancelOrder(id: string): Promise<GeneralResponse<Order>> {
    const response = await apiClient.put(`/orders/${id}/cancel`)
    return response.data
  }

  /**
   * Get order statistics (admin only)
   */
  async getOrderStats(): Promise<GeneralResponse<{
    totalOrders: number
    pendingOrders: number
    completedOrders: number
    totalRevenue: number
  }>> {
    const response = await apiClient.get('/orders/stats')
    return response.data
  }

  /**
   * Validate order data
   */
  validateOrderData(orderData: CreateOrderModel): boolean {
    return orderData.cart_id > 0 && orderData.shipping_address.trim().length > 0
  }
}

export const orderService = new OrderService()