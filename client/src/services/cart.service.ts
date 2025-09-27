import { apiClient } from '../api/client'
import type {
  Cart,
  CartItem,
  AddToCartModel,
  UpdateCartItemModel,
  GeneralResponse
} from '../types/api'

export class CartService {
  /**
   * Get user's cart
   */
  async getCart(): Promise<GeneralResponse<Cart>> {
    const response = await apiClient.get('/cart')
    return response.data
  }

  /**
   * Add item to cart
   */
  async addToCart(productId: string, quantity: number): Promise<GeneralResponse<CartItem>> {
    const cartData: AddToCartModel = {
      product_id: productId,
      quantity
    }
    const response = await apiClient.post('/cart/items', cartData)
    return response.data
  }

  /**
   * Update cart item quantity
   */
  async updateCartItem(itemId: number, quantity: number): Promise<GeneralResponse<CartItem>> {
    const updateData: UpdateCartItemModel = { quantity }
    const response = await apiClient.put(`/cart/items/${itemId}`, updateData)
    return response.data
  }

  /**
   * Remove item from cart
   */
  async removeFromCart(itemId: number): Promise<GeneralResponse<void>> {
    const response = await apiClient.delete(`/cart/items/${itemId}`)
    return response.data
  }

  /**
   * Clear entire cart
   */
  async clearCart(): Promise<GeneralResponse<void>> {
    const response = await apiClient.delete('/cart')
    return response.data
  }

  /**
   * Get cart item count
   */
  async getCartCount(): Promise<GeneralResponse<number>> {
    const response = await apiClient.get('/cart/count')
    return response.data
  }

  /**
   * Validate cart data
   */
  validateCartData(productId: string, quantity: number): boolean {
    return productId.length > 0 && quantity > 0
  }
}

export const cartService = new CartService()