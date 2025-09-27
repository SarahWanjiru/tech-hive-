import { apiClient } from '../api/client.js'

export const productService = {
  /**
   * Get all products with optional filters
   * @param {Object} params - Query parameters
   * @returns {Promise} API response
   */
  async getProducts(params = {}) {
    try {
      const response = await apiClient.get('/v1/api/product', { params })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to fetch products')
    }
  },

  /**
   * Search products with advanced filters
   * @param {Object} searchCriteria - Search parameters
   * @returns {Promise} API response
   */
  async searchProducts(searchCriteria) {
    try {
      const response = await apiClient.post('/v1/api/product/search', {
        name: searchCriteria.name || '',
        min_price: searchCriteria.minPrice || 0,
        max_price: searchCriteria.maxPrice || 999999,
        in_stock: searchCriteria.inStock !== false,
        page: searchCriteria.page || 1,
        limit: searchCriteria.limit || 10,
        sort_by: searchCriteria.sortBy || 'name',
        sort_order: searchCriteria.sortOrder || 'asc',
        category: searchCriteria.category || ''
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Search failed')
    }
  },

  /**
   * Get product by ID
   * @param {string|number} id - Product ID
   * @returns {Promise} API response
   */
  async getProductById(id) {
    try {
      const response = await apiClient.get(`/v1/api/product/${id}`)
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to fetch product')
    }
  },

  /**
   * Create new product (Admin only)
   * @param {Object} productData - Product data
   * @returns {Promise} API response
   */
  async createProduct(productData) {
    try {
      const response = await apiClient.post('/v1/api/product', {
        name: productData.name,
        description: productData.description,
        price: parseFloat(productData.price),
        stock: parseInt(productData.stock),
        image_url: productData.imageUrl || '',
        category: productData.category || ''
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to create product')
    }
  },

  /**
   * Update product (Admin only)
   * @param {string|number} id - Product ID
   * @param {Object} productData - Updated product data
   * @returns {Promise} API response
   */
  async updateProduct(id, productData) {
    try {
      const response = await apiClient.put(`/v1/api/product/${id}`, {
        name: productData.name,
        description: productData.description,
        price: parseFloat(productData.price),
        stock: parseInt(productData.stock),
        image_url: productData.imageUrl || '',
        category: productData.category || ''
      })
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to update product')
    }
  },

  /**
   * Delete product (Admin only)
   * @param {string|number} id - Product ID
   * @returns {Promise} API response
   */
  async deleteProduct(id) {
    try {
      const response = await apiClient.delete(`/v1/api/product/${id}`)
      return response
    } catch (error) {
      throw new Error(error.response?.data?.message || 'Failed to delete product')
    }
  },

  /**
   * Get product categories
   * @returns {Promise} API response
   */
  async getCategories() {
    try {
      // This would be a new endpoint in the backend
      const response = await apiClient.get('/v1/api/product/categories')
      return response
    } catch (error) {
      // For now, return mock categories
      return {
        data: {
          categories: [
            'Electronics',
            'Clothing',
            'Books',
            'Home & Garden',
            'Sports',
            'Beauty'
          ]
        }
      }
    }
  }
}