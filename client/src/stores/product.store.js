import { defineStore } from 'pinia'
import { productService } from '../services/product.service.js'

export const useProductStore = defineStore('product', {
  state: () => ({
    products: [],
    currentProduct: null,
    categories: [],
    loading: false,
    error: null,
    pagination: {
      page: 1,
      limit: 10,
      total: 0,
      totalPages: 0
    },
    filters: {
      search: '',
      category: '',
      minPrice: 0,
      maxPrice: 999999,
      inStock: true,
      sortBy: 'name',
      sortOrder: 'asc'
    }
  }),

  getters: {
    filteredProducts: (state) => {
      let filtered = [...state.products]

      // Apply search filter
      if (state.filters.search) {
        filtered = filtered.filter(product =>
          product.name.toLowerCase().includes(state.filters.search.toLowerCase()) ||
          product.description.toLowerCase().includes(state.filters.search.toLowerCase())
        )
      }

      // Apply category filter
      if (state.filters.category) {
        filtered = filtered.filter(product => product.category === state.filters.category)
      }

      // Apply price filters
      filtered = filtered.filter(product =>
        product.price >= state.filters.minPrice &&
        product.price <= state.filters.maxPrice
      )

      // Apply stock filter
      if (state.filters.inStock) {
        filtered = filtered.filter(product => product.stock > 0)
      }

      // Apply sorting
      filtered.sort((a, b) => {
        let aValue = a[state.filters.sortBy]
        let bValue = b[state.filters.sortBy]

        if (state.filters.sortBy === 'price') {
          aValue = parseFloat(aValue)
          bValue = parseFloat(bValue)
        }

        if (state.filters.sortOrder === 'asc') {
          return aValue > bValue ? 1 : -1
        } else {
          return aValue < bValue ? 1 : -1
        }
      })

      return filtered
    },

    availableCategories: (state) => {
      const categories = [...new Set(state.products.map(p => p.category).filter(Boolean))]
      return categories
    }
  },

  actions: {
    /**
     * Fetch all products
     * @param {Object} params - Query parameters
     */
    async fetchProducts(params = {}) {
      this.loading = true
      this.error = null

      try {
        const response = await productService.getProducts(params)
        this.products = response.data.products || response.data || []

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
     * Search products with filters
     * @param {Object} searchCriteria - Search parameters
     */
    async searchProducts(searchCriteria = {}) {
      this.loading = true
      this.error = null

      try {
        const response = await productService.searchProducts(searchCriteria)
        this.products = response.data.products || response.data || []

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
     * Get product by ID
     * @param {string|number} id - Product ID
     */
    async fetchProductById(id) {
      this.loading = true
      this.error = null

      try {
        const response = await productService.getProductById(id)
        this.currentProduct = response.data
        return response
      } catch (error) {
        this.error = error.message
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Create new product (Admin only)
     * @param {Object} productData - Product data
     */
    async createProduct(productData) {
      try {
        const response = await productService.createProduct(productData)
        this.products.unshift(response.data)
        return response
      } catch (error) {
        this.error = error.message
        throw error
      }
    },

    /**
     * Update product (Admin only)
     * @param {string|number} id - Product ID
     * @param {Object} productData - Updated product data
     */
    async updateProduct(id, productData) {
      try {
        const response = await productService.updateProduct(id, productData)

        // Update product in the list
        const index = this.products.findIndex(p => p.id === id)
        if (index !== -1) {
          this.products[index] = { ...this.products[index], ...response.data }
        }

        // Update current product if it's the same
        if (this.currentProduct?.id === id) {
          this.currentProduct = { ...this.currentProduct, ...response.data }
        }

        return response
      } catch (error) {
        this.error = error.message
        throw error
      }
    },

    /**
     * Delete product (Admin only)
     * @param {string|number} id - Product ID
     */
    async deleteProduct(id) {
      try {
        await productService.deleteProduct(id)

        // Remove from products list
        this.products = this.products.filter(p => p.id !== id)

        // Clear current product if it's the deleted one
        if (this.currentProduct?.id === id) {
          this.currentProduct = null
        }
      } catch (error) {
        this.error = error.message
        throw error
      }
    },

    /**
     * Update filters
     * @param {Object} newFilters - New filter values
     */
    updateFilters(newFilters) {
      this.filters = { ...this.filters, ...newFilters }
    },

    /**
     * Clear all filters
     */
    clearFilters() {
      this.filters = {
        search: '',
        category: '',
        minPrice: 0,
        maxPrice: 999999,
        inStock: true,
        sortBy: 'name',
        sortOrder: 'asc'
      }
    },

    /**
     * Clear error state
     */
    clearError() {
      this.error = null
    }
  }
})