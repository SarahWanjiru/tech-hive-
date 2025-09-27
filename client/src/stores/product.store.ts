import { defineStore } from 'pinia'
import { productService } from '../services/product.service'
import type {
  Product,
  ProductCreateOrUpdateModel,
  ProductSearchModel,
  //ProductSearchResponse
} from '../types/api'

interface ProductState {
  products: Product[]
  currentProduct: Product | null
  loading: boolean
  error: string | null
  filters: ProductSearchModel
  totalCount: number
  availableCategories: string[]
}

export const useProductStore = defineStore('product', {
  state: (): ProductState => ({
    products: [],
    currentProduct: null,
    loading: false,
    error: null,
    filters: {
      page: 1,
      limit: 20,
      sort_by: 'name',
      sort_order: 'asc'
    },
    totalCount: 0,
    availableCategories: []
  }),

  getters: {
    // Get filtered products based on current filters
    filteredProducts: (state: ProductState): Product[] => {
      // Ensure products is an array before spreading
      const productsArray = Array.isArray(state.products) ? state.products : []
      let filtered = [...productsArray]

      // Apply search filter
      if (state.filters.name) {
        const searchTerm = state.filters.name.toLowerCase()
        filtered = filtered.filter(product =>
          product.name.toLowerCase().includes(searchTerm) ||
          product.description.toLowerCase().includes(searchTerm)
        )
      }

      // Apply price filters
      if (state.filters.min_price !== undefined) {
        filtered = filtered.filter(product => product.price >= state.filters.min_price!)
      }

      if (state.filters.max_price !== undefined) {
        filtered = filtered.filter(product => product.price <= state.filters.max_price!)
      }

      // Apply stock filter
      if (state.filters.in_stock) {
        filtered = filtered.filter(product => product.stock > 0)
      }

      // Apply sorting
      const sortBy = state.filters.sort_by || 'name'
      const sortOrder = state.filters.sort_order || 'asc'

      filtered.sort((a, b) => {
        let aValue: any, bValue: any

        switch (sortBy) {
          case 'price':
            aValue = a.price
            bValue = b.price
            break
          case 'stock':
            aValue = a.stock
            bValue = b.stock
            break
          case 'name':
          default:
            aValue = a.name.toLowerCase()
            bValue = b.name.toLowerCase()
            break
        }

        if (sortOrder === 'asc') {
          return aValue < bValue ? -1 : aValue > bValue ? 1 : 0
        } else {
          return aValue > bValue ? -1 : aValue < bValue ? 1 : 0
        }
      })

      return filtered
    },

    // Get unique categories from products
    categories: (state: ProductState): string[] => {
      const categories = new Set(state.products.map(p => p.name.split(' ')[0]))
      return Array.from(categories)
    }
  },

  actions: {
    /**
     * Fetch all products with optional parameters
     */
    async fetchProducts(params?: Partial<ProductSearchModel>): Promise<void> {
      this.loading = true
      this.error = null

      try {
        if (params) {
          this.filters = { ...this.filters, ...params }
        }

        const response = await productService.getProducts(this.filters)
        this.products = response.data.products
        this.totalCount = response.data.total_count
        this.availableCategories = await this.fetchCategories()
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch products'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Fetch single product by ID
     */
    async fetchProductById(id: string): Promise<Product> {
      this.loading = true
      this.error = null

      try {
        const response = await productService.getProductById(id)
        this.currentProduct = response.data
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to fetch product'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Create new product
     */
    async createProduct(productData: ProductCreateOrUpdateModel): Promise<Product> {
      this.loading = true
      this.error = null

      try {
        const response = await productService.createProduct(productData)
        this.products.unshift(response.data)
        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to create product'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Update existing product
     */
    async updateProduct(id: string, productData: ProductCreateOrUpdateModel): Promise<Product> {
      this.loading = true
      this.error = null

      try {
        const response = await productService.updateProduct(id, productData)

        // Update in products array
        const index = this.products.findIndex(p => p.id === parseInt(id))
        if (index !== -1) {
          this.products[index] = response.data
        }

        // Update current product if it's the same
        if (this.currentProduct?.id === parseInt(id)) {
          this.currentProduct = response.data
        }

        return response.data
      } catch (error: any) {
        this.error = error.message || 'Failed to update product'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Delete product
     */
    async deleteProduct(id: string): Promise<void> {
      this.loading = true
      this.error = null

      try {
        await productService.deleteProduct(id)

        // Remove from products array
        this.products = this.products.filter(p => p.id !== parseInt(id))

        // Clear current product if it's the same
        if (this.currentProduct?.id === parseInt(id)) {
          this.currentProduct = null
        }
      } catch (error: any) {
        this.error = error.message || 'Failed to delete product'
        throw error
      } finally {
        this.loading = false
      }
    },

    /**
     * Update filters and refetch products
     */
    async updateFilters(filters: Partial<ProductSearchModel>): Promise<void> {
      this.filters = { ...this.filters, ...filters, page: 1 }
      await this.fetchProducts()
    },

    /**
     * Clear all filters
     */
    clearFilters(): void {
      this.filters = {
        page: 1,
        limit: 20,
        sort_by: 'name',
        sort_order: 'asc'
      }
    },

    /**
     * Fetch available categories
     */
    async fetchCategories(): Promise<string[]> {
      try {
        const response = await productService.getCategories()
        return response.data
      } catch (error) {
        console.error('Failed to fetch categories:', error)
        return []
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