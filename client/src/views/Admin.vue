<template>
  <div class="admin-page">
    <div class="container">
      <div class="page-header">
        <h1>Admin Dashboard</h1>
        <p>Manage your store</p>
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="loading-container">
        <el-icon class="loading"><Loading /></el-icon>
        <p>Loading dashboard...</p>
      </div>

      <!-- Dashboard Content -->
      <div v-else class="admin-content">
        <!-- Statistics Cards -->
        <div class="stats-grid">
          <div class="stat-card">
            <div class="stat-icon">
              <el-icon><Box /></el-icon>
            </div>
            <div class="stat-info">
              <h3>{{ stats.totalProducts }}</h3>
              <p>Total Products</p>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon" style="background: #67C23A;">
              <el-icon><ShoppingCart /></el-icon>
            </div>
            <div class="stat-info">
              <h3>{{ stats.totalOrders }}</h3>
              <p>Total Orders</p>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon" style="background: #E6A23C;">
              <el-icon><Money /></el-icon>
            </div>
            <div class="stat-info">
              <h3>KSh {{ stats.totalRevenue.toFixed(2) }}</h3>
              <p>Total Revenue</p>
            </div>
          </div>

          <div class="stat-card">
            <div class="stat-icon" style="background: #F56C6C;">
              <el-icon><User /></el-icon>
            </div>
            <div class="stat-info">
              <h3>{{ stats.totalUsers }}</h3>
              <p>Total Users</p>
            </div>
          </div>
        </div>

        <!-- Product Management Section -->
        <div class="product-management">
          <div class="section-header">
            <h3>Product Management</h3>
            <el-button
              type="primary"
              @click="showAddProductDialog = true"
            >
              <el-icon><Plus /></el-icon>
              Add New Product
            </el-button>
          </div>

          <!-- Products Table -->
          <el-table
            :data="products"
            style="width: 100%"
            v-loading="loadingProducts"
          >
            <el-table-column prop="name" label="Name" width="200">
              <template #default="scope">
                <div class="product-info">
                  <img
                    :src="scope.row.image_url || '/placeholder-product.svg'"
                    :alt="scope.row.name"
                    class="product-thumbnail"
                    @error="handleImageError"
                  />
                  <span>{{ scope.row.name }}</span>
                </div>
              </template>
            </el-table-column>

            <el-table-column prop="description" label="Description">
              <template #default="scope">
                <span class="description-cell">{{ scope.row.description }}</span>
              </template>
            </el-table-column>

            <el-table-column prop="price" label="Price" width="120">
              <template #default="scope">
                KSh {{ scope.row.price }}
              </template>
            </el-table-column>

            <el-table-column prop="stock" label="Stock" width="100">
              <template #default="scope">
                {{ scope.row.stock }}
              </template>
            </el-table-column>

            <el-table-column label="Actions" width="150">
              <template #default="scope">
                <el-button
                  size="small"
                  @click="editProduct(scope.row)"
                >
                  Edit
                </el-button>
                <el-button
                  size="small"
                  type="danger"
                  @click="deleteProduct(scope.row)"
                >
                  Delete
                </el-button>
              </template>
            </el-table-column>
          </el-table>
        </div>

        <!-- Recent Orders -->
        <div class="recent-section">
          <h3>Recent Orders</h3>
          <div class="recent-orders">
            <div
              v-for="order in recentOrders"
              :key="order.id"
              class="recent-order"
            >
              <div class="order-info">
                <h5>Order #{{ order.id }}</h5>
                <p>{{ order.user?.name || 'Unknown User' }}</p>
              </div>

              <div class="order-meta">
                <span>{{ formatDate(order.created_at) }}</span>
                <span class="order-total">KSh {{ order.total.toFixed(2) }}</span>
              </div>

              <div class="order-status">
                <el-tag
                  :type="getStatusType(order.status)"
                  size="small"
                >
                  {{ order.status.toUpperCase() }}
                </el-tag>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Product Management Dialog -->
      <el-dialog
        v-model="showAddProductDialog"
        :title="editingProduct ? 'Edit Product' : 'Add New Product'"
        width="600px"
        @close="resetProductForm"
      >
        <el-form
          ref="productForm"
          :model="productFormData"
          :rules="productRules"
          label-width="120px"
        >
          <el-form-item label="Name" prop="name">
            <el-input v-model="productFormData.name" placeholder="Product name" />
          </el-form-item>

          <el-form-item label="Description" prop="description">
            <el-input
              v-model="productFormData.description"
              type="textarea"
              :rows="3"
              placeholder="Product description"
            />
          </el-form-item>

          <el-form-item label="Price" prop="price">
            <el-input-number
              v-model="productFormData.price"
              :min="0"
              :precision="2"
              placeholder="0.00"
            />
          </el-form-item>

          <el-form-item label="Stock" prop="stock">
            <el-input-number
              v-model="productFormData.stock"
              :min="0"
              placeholder="0"
            />
          </el-form-item>

          <el-form-item label="Image URL" prop="image_url">
            <el-input
              v-model="productFormData.image_url"
              placeholder="https://example.com/image.jpg"
            />
          </el-form-item>
        </el-form>

        <template #footer>
          <span class="dialog-footer">
            <el-button @click="showAddProductDialog = false">Cancel</el-button>
            <el-button
              type="primary"
              @click="saveProduct"
              :loading="savingProduct"
            >
              {{ editingProduct ? 'Update' : 'Create' }} Product
            </el-button>
          </span>
        </template>
      </el-dialog>

      <!-- Seed Data Dialog -->
      <el-dialog
        v-model="showSeedDialog"
        title="Seed Sample Data"
        width="500px"
      >
        <p>This will add sample products and users to your store for testing.</p>

        <template #footer>
          <span class="dialog-footer">
            <el-button @click="showSeedDialog = false">Cancel</el-button>
            <el-button type="primary" @click="seedData" :loading="seeding">
              Seed Data
            </el-button>
          </span>
        </template>
      </el-dialog>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useAuthStore } from '../stores/auth.store.ts'
import { useProductStore } from '../stores/product.store.ts'
import { useOrderStore } from '../stores/order.store.ts'
import { ElMessage } from 'element-plus'
import { Loading, Box, ShoppingCart, Money, User, Plus, List, MagicStick } from '@element-plus/icons-vue'

const authStore = useAuthStore()
const productStore = useProductStore()
const orderStore = useOrderStore()

const loading = ref(false)
const seeding = ref(false)
const showSeedDialog = ref(false)
const showAddProductDialog = ref(false)
const loadingProducts = ref(false)
const savingProduct = ref(false)
const editingProduct = ref(null)

const productForm = ref()
const stats = reactive({
  totalProducts: 0,
  totalOrders: 0,
  totalRevenue: 0,
  totalUsers: 0
})

const recentOrders = ref([])
const products = ref([])

const productFormData = reactive({
  name: '',
  description: '',
  price: 0,
  stock: 0,
  image_url: ''
})

const productRules = {
  name: [
    { required: true, message: 'Please enter product name', trigger: 'blur' },
    { min: 2, message: 'Product name must be at least 2 characters', trigger: 'blur' }
  ],
  description: [
    { required: true, message: 'Please enter product description', trigger: 'blur' },
    { min: 10, message: 'Description must be at least 10 characters', trigger: 'blur' }
  ],
  price: [
    { required: true, message: 'Please enter product price', trigger: 'blur' },
    { type: 'number', min: 0, message: 'Price must be greater than 0', trigger: 'blur' }
  ],
  stock: [
    { required: true, message: 'Please enter stock quantity', trigger: 'blur' },
    { type: 'number', min: 0, message: 'Stock must be 0 or greater', trigger: 'blur' }
  ],
  image_url: [
    { type: 'url', message: 'Please enter a valid URL', trigger: 'blur' }
  ]
}

const loadDashboardData = async () => {
  loading.value = true
  try {
    // Load products to get count
    await productStore.fetchProducts({ page: 1, limit: 1 })
    stats.totalProducts = productStore.totalCount

    // Load orders to get count and recent orders
    const ordersResponse = await orderStore.fetchAllOrders({ page: 1, limit: 5 })
    stats.totalOrders = ordersResponse.total || 0

    // Calculate revenue from orders
    const allOrdersResponse = await orderStore.fetchAllOrders({ page: 1, limit: 1000 })
    stats.totalRevenue = allOrdersResponse.data.reduce((sum, order) => sum + order.total, 0)

    // Set recent orders
    recentOrders.value = ordersResponse.data

    // For demo purposes, set users count
    stats.totalUsers = 2 // This would come from a users API in a real app

  } catch (error) {
    ElMessage.error('Failed to load dashboard data')
    console.error('Dashboard loading error:', error)
  } finally {
    loading.value = false
  }
}

const loadProducts = async () => {
  loadingProducts.value = true
  try {
    await productStore.fetchProducts({ page: 1, limit: 100 })
    products.value = productStore.products
  } catch (error) {
    ElMessage.error('Failed to load products')
    console.error('Load products error:', error)
  } finally {
    loadingProducts.value = false
  }
}

const resetProductForm = () => {
  productFormData.name = ''
  productFormData.description = ''
  productFormData.price = 0
  productFormData.stock = 0
  productFormData.image_url = ''
  editingProduct.value = null
}

const editProduct = (product) => {
  editingProduct.value = product
  productFormData.name = product.name
  productFormData.description = product.description
  productFormData.price = product.price
  productFormData.stock = product.stock
  productFormData.image_url = product.image_url || ''
  showAddProductDialog.value = true
}

const saveProduct = async () => {
  if (!productForm.value) return

  try {
    // Clear any previous validation errors
    productForm.value.clearValidate()

    // Validate form
    await productForm.value.validate()

    savingProduct.value = true

    const productData = {
      name: productFormData.name.trim(),
      description: productFormData.description.trim(),
      price: Number(productFormData.price),
      stock: Number(productFormData.stock),
      image_url: productFormData.image_url.trim() || undefined
    }

    if (editingProduct.value) {
      // Update existing product
      await productStore.updateProduct(editingProduct.value.id.toString(), productData)
      ElMessage.success('Product updated successfully')
    } else {
      // Create new product
      await productStore.createProduct(productData)
      ElMessage.success('Product created successfully')
    }

    showAddProductDialog.value = false
    resetProductForm()
    loadProducts()
    loadDashboardData()

  } catch (error) {
    console.error('Save product error:', error)
    // Only show error if it's not a validation error
    if (error.message !== 'validation failed') {
      ElMessage.error('Failed to save product: ' + (error.message || 'Unknown error'))
    }
  } finally {
    savingProduct.value = false
  }
}

const deleteProduct = async (product) => {
  try {
    await ElMessageBox.confirm(
      `Are you sure you want to delete "${product.name}"?`,
      'Confirm Deletion',
      {
        confirmButtonText: 'Delete',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
    )

    await productStore.deleteProduct(product.id.toString())
    ElMessage.success('Product deleted successfully')
    loadProducts()
    loadDashboardData()

  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error('Failed to delete product')
      console.error('Delete product error:', error)
    }
  }
}

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString('en-US', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const getStatusType = (status) => {
  const statusMap = {
    'pending': 'warning',
    'confirmed': 'info',
    'shipped': 'primary',
    'delivered': 'success',
    'cancelled': 'danger'
  }
  return statusMap[status] || 'info'
}

const seedData = async () => {
  seeding.value = true
  try {
    // Call the seed all endpoint to seed both users and products
    const response = await fetch('/v1/api/seed/all', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      }
    })

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`)
    }

    const result = await response.json()

    if (result.code === 200) {
      ElMessage.success('Sample data seeded successfully!')
      showSeedDialog.value = false

      // Reload dashboard data to show updated counts
      loadDashboardData()
    } else {
      throw new Error(result.message || 'Failed to seed data')
    }

  } catch (error) {
    ElMessage.error('Failed to seed data: ' + error.message)
    console.error('Seed data error:', error)
  } finally {
    seeding.value = false
  }
}

const handleImageError = (event) => {
  event.target.src = '/placeholder-product.svg'
}

onMounted(() => {
  if (authStore.isAuthenticated && authStore.user?.role === 'admin') {
    loadDashboardData()
    loadProducts()
  } else {
    ElMessage.warning('Admin access required')
  }
})
</script>

<style scoped>
.admin-page {
  padding: 2rem 0;
}

.page-header {
  text-align: center;
  margin-bottom: 2rem;
}

.page-header h1 {
  font-size: 2.5rem;
  color: var(--gray-800);
  margin-bottom: 0.5rem;
}

.page-header p {
  font-size: 1.125rem;
  color: var(--gray-600);
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.stat-card {
  display: flex;
  align-items: center;
  gap: 1rem;
}

.stat-icon {
  width: 50px;
  height: 50px;
  background: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--white);
  font-size: 20px;
}

.stat-info h3 {
  margin: 0;
  font-size: 2rem;
  color: var(--gray-800);
}

.stat-info p {
  margin: 0;
  color: var(--gray-600);
  font-size: 0.875rem;
}

.admin-actions,
.recent-section {
  margin-bottom: 2rem;
}

.action-buttons {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
  gap: 1rem;
}

.recent-orders {
  max-height: 400px;
  overflow-y: auto;
}

.recent-order {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 1rem;
  border-bottom: 1px solid var(--gray-200);
}

.recent-order:last-child {
  border-bottom: none;
}

.order-info h5 {
  margin: 0 0 0.25rem 0;
  color: var(--gray-800);
}

.order-info p {
  margin: 0;
  font-size: 0.875rem;
  color: var(--gray-600);
}

.order-total {
  font-weight: 600;
  color: var(--primary-color);
}

/* Product Management Styles */
.product-management {
  margin-bottom: 2rem;
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
}

.section-header h3 {
  margin: 0;
  color: var(--gray-800);
}

.product-info {
  display: flex;
  align-items: center;
  gap: 0.75rem;
}

.product-thumbnail {
  width: 40px;
  height: 40px;
  border-radius: var(--border-radius-sm);
  object-fit: cover;
  flex-shrink: 0;
}

.description-cell {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  max-width: 300px;
}

/* Loading and Empty States */
.loading-container,
.empty-state {
  text-align: center;
  padding: 2rem;
}

/* Form Input Styles */
.el-form-item {
  margin-bottom: 1.5rem;
}

.el-input-number {
  width: 100%;
}

.el-input-number .el-input__inner {
  text-align: left;
}

/* Ensure form inputs are interactive */
.el-input__inner,
.el-textarea__inner,
.el-input-number__decrease,
.el-input-number__increase {
  pointer-events: auto !important;
  user-select: text;
}

/* Fix for input field typing issues */
.el-form-item.is-error .el-input__inner,
.el-form-item.is-error .el-textarea__inner {
  border-color: #f56c6c;
}

.el-form-item.is-success .el-input__inner,
.el-form-item.is-success .el-textarea__inner {
  border-color: #67c23a;
}

/* Responsive Design */
@media (max-width: 768px) {
  .stats-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .action-buttons {
    grid-template-columns: 1fr;
  }

  .recent-order {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.5rem;
  }

  .el-dialog {
    width: 95% !important;
    margin: 2rem auto;
  }

  .el-form-item {
    margin-bottom: 1rem;
  }
}
</style>