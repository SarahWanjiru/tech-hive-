import { defineStore } from 'pinia'
import { cartService } from '../services/cart.service'
import type { Cart, CartItem } from '../types/api'

interface CartState {
  cart: Cart | null
  loading: boolean
  error: string | null
}

export const useCartStore = defineStore('cart', {
  state: (): CartState => ({
    cart: null,
    loading: false,
    error: null
  }),

  getters: {
    // Get all cart items
    items: (state: CartState): CartItem[] => state.cart?.items || [],

    // Check if cart is empty
    isEmpty: (state: CartState): boolean => !state.cart || !state.cart.items || state.cart.items.length === 0,

    // Get total number of items
    totalItems: (state: CartState): number => {
      return state.cart?.items?.reduce((total, item) => total + item.quantity, 0) || 0
    },

    // Get total amount
    totalAmount: (state: CartState): number => {
      return state.cart?.items?.reduce((total, item) => total + (item.price * item.quantity), 0) || 0
    },

    // Get item by product ID
    getItemByProductId: (state: CartState) => (productId: string): CartItem | undefined => {
      return state.cart?.items?.find(item => item.product_id === productId)
    }
  },

  actions: {
    /**
     * Fetch user's cart
     */
    async fetchCart(): Promise<Cart> {
      this.loading = true
      this.error = null

      try {
        const response = await cartService.getCart()
        this.cart = response.data
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch cart'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Add item to cart
     */
    async addToCart(productId: string, quantity: number): Promise<CartItem> {
      if (!cartService.validateCartData(productId, quantity)) {
        throw new Error('Invalid cart data')
      }

      this.loading = true
      this.error = null

      try {
        const response = await cartService.addToCart(productId, quantity)

        // Ensure cart exists, fetch if necessary
        if (!this.cart) {
          await this.fetchCart()
        }

        // Update or add item in cart
        if (this.cart && this.cart.items) {
          const existingItemIndex = this.cart.items.findIndex(item => item.product_id === productId)
          if (existingItemIndex !== -1) {
            // Update existing item quantity
            this.cart.items[existingItemIndex].quantity += quantity
          } else {
            // Add new item
            this.cart.items.push(response.data)
          }
          this.cart.total = this.totalAmount
        }

        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to add item to cart'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Update cart item quantity
     */
    async updateCartItem(itemId: number, quantity: number): Promise<CartItem> {
      if (quantity <= 0) {
        throw new Error('Quantity must be greater than 0')
      }

      this.loading = true
      this.error = null

      try {
        const response = await cartService.updateCartItem(itemId, quantity)

        // Update item in cart
        if (this.cart) {
          const itemIndex = this.cart.items.findIndex(item => item.id === itemId)
          if (itemIndex !== -1) {
            this.cart.items[itemIndex] = response.data
            this.cart.total = this.totalAmount
          }
        }

        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to update cart item'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Remove item from cart
     */
    async removeFromCart(itemId: number): Promise<void> {
      this.loading = true
      this.error = null

      try {
        await cartService.removeFromCart(itemId)

        // Remove item from cart
        if (this.cart) {
          this.cart.items = this.cart.items.filter(item => item.id !== itemId)
          this.cart.total = this.totalAmount
        }
      } catch (error: any) {
        this.error = error.message || 'Failed to remove item from cart'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Clear entire cart
     */
    async clearCart(): Promise<void> {
      this.loading = true
      this.error = null

      try {
        await cartService.clearCart()
        this.cart = null
      } catch (error: any) {
        this.error = error.message || 'Failed to clear cart'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Get cart item count
     */
    async getCartCount(): Promise<number> {
      try {
        // Fetch cart to get current count
        if (!this.cart) {
          await this.fetchCart()
        }
        return this.totalItems
      } catch (error) {
        console.error('Failed to get cart count:', error)
        return 0
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