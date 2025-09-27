import { apiClient } from '../api/client'
import type {
  Product,
  ProductCreateOrUpdateModel,
  ProductSearchModel,
  ProductSearchResponse,
  GeneralResponse
} from '../types/api'

export class ProductService {
  /**
   * Get all products with optional filtering
   */
  async getProducts(params?: ProductSearchModel): Promise<GeneralResponse<ProductSearchResponse>> {
    const queryParams = new URLSearchParams()

    if (params) {
      Object.entries(params).forEach(([key, value]) => {
        if (value !== undefined && value !== null && value !== '') {
          queryParams.append(key, value.toString())
        }
      })
    }

    const response = await apiClient.get(`/v1/api/product?${queryParams.toString()}`)
    return response.data
  }

  /**
   * Get product by ID
   */
  async getProductById(id: string): Promise<GeneralResponse<Product>> {
    const response = await apiClient.get(`/v1/api/product/${id}`)
    return response.data
  }

  /**
   * Create new product
   */
  async createProduct(productData: ProductCreateOrUpdateModel): Promise<GeneralResponse<Product>> {
    const response = await apiClient.post('/v1/api/product', productData)
    return response.data
  }

  /**
   * Update existing product
   */
  async updateProduct(id: string, productData: ProductCreateOrUpdateModel): Promise<GeneralResponse<Product>> {
    const response = await apiClient.put(`/v1/api/product/${id}`, productData)
    return response.data
  }

  /**
   * Delete product
   */
  async deleteProduct(id: string): Promise<GeneralResponse<void>> {
    const response = await apiClient.delete(`/v1/api/product/${id}`)
    return response.data
  }

  /**
   * Search products
   */
  async searchProducts(query: string): Promise<GeneralResponse<Product[]>> {
    const response = await apiClient.post(`/v1/api/product/search`, { name: query })
    return response.data
  }

  /**
   * Get product categories
   */
  async getCategories(): Promise<GeneralResponse<string[]>> {
    // For now, return empty array as categories endpoint may not exist
    return { code: 200, message: 'Success', data: [] }
  }

  /**
   * Validate phone number format
   */
  validatePhoneNumber(phone: string): boolean {
    const phoneRegex = /^254[0-9]{9}$/
    return phoneRegex.test(phone.replace(/[^0-9]/g, ''))
  }

  /**
   * Format phone number to standard format
   */
  formatPhoneNumber(phone: string): string {
    const cleaned = phone.replace(/[^0-9]/g, '')
    if (cleaned.startsWith('0')) {
      return '254' + cleaned.substring(1)
    }
    return cleaned.startsWith('254') ? cleaned : '254' + cleaned
  }
}

export const productService = new ProductService()