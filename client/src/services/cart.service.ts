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
    const response = await apiClient.get('/v1/api/cart')
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
    const response = await apiClient.post('/v1/api/cart/items', cartData)
    return response.data
  }

  /**
   * Update cart item quantity
   */
  async updateCartItem(itemId: number, quantity: number): Promise<GeneralResponse<CartItem>> {
    const updateData: UpdateCartItemModel = { quantity }
    const response = await apiClient.put(`/v1/api/cart/items/${itemId}`, updateData)
    return response.data
  }

  /**
   * Remove item from cart
   */
  async removeFromCart(itemId: number): Promise<GeneralResponse<void>> {
    const response = await apiClient.delete(`/v1/api/cart/items/${itemId}`)
    return response.data
  }

  /**
   * Clear entire cart
   */
  async clearCart(): Promise<GeneralResponse<void>> {
    const response = await apiClient.delete('/v1/api/cart')
    return response.data
  }

  /**
   * Get cart item count
   */
  // Note: getCartCount endpoint not implemented in backend
  // async getCartCount(): Promise<GeneralResponse<number>> {
  //   const response = await apiClient.get('/v1/api/cart/count')
  //   return response.data
  // }

  /**
   * Validate cart data
   */
  validateCartData(productId: string, quantity: number): boolean {
    console.log('🔍 Validating cart data:', {
      productId,
      productIdType: typeof productId,
      productIdLength: productId?.length,
      quantity,
      quantityType: typeof quantity
    })

    const isValid = !!(productId && productId.length > 0 && quantity > 0)

    console.log('✅ Validation result:', isValid)

    return isValid
  }
}

export const cartService = new CartService()