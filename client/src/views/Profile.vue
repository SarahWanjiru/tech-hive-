<template>
  <div class="profile-page">
    <div class="container">
      <div class="page-header">
        <h1>My Profile</h1>
        <p>Manage your account information</p>
      </div>

      <div class="profile-content">
        <div class="profile-info">
          <el-card>
            <template #header>
              <div class="card-header">
                <h3>Profile Information</h3>
              </div>
            </template>

            <div class="user-info">
              <div class="info-item">
                <label>Name:</label>
                <span>{{ authStore.user?.name }}</span>
              </div>
              <div class="info-item">
                <label>Email:</label>
                <span>{{ authStore.user?.email }}</span>
              </div>
              <div class="info-item">
                <label>Role:</label>
                <span>{{ authStore.user?.role }}</span>
              </div>
              <div class="info-item">
                <label>Member Since:</label>
                <span>{{ formatDate(authStore.user?.created_at) }}</span>
              </div>
            </div>
          </el-card>
        </div>

        <div class="profile-actions">
          <el-card>
            <template #header>
              <div class="card-header">
                <h3>Account Actions</h3>
              </div>
            </template>

            <div class="action-buttons">
              <el-button type="primary" @click="$router.push('/orders')">
                View My Orders
              </el-button>
              <el-button @click="handleLogout">
                Logout
              </el-button>
            </div>
          </el-card>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useAuthStore } from '../stores/auth.store.ts'
import { useRouter } from 'vue-router'
import { ElMessage, ElMessageBox } from 'element-plus'

const authStore = useAuthStore()
const router = useRouter()

const formatDate = (dateString) => {
  if (!dateString) return 'N/A'
  return new Date(dateString).toLocaleDateString('en-US', {
    year: 'numeric',
    month: 'long',
    day: 'numeric'
  })
}

const handleLogout = async () => {
  try {
    await ElMessageBox.confirm(
      'Are you sure you want to logout?',
      'Confirm Logout',
      {
        confirmButtonText: 'Logout',
        cancelButtonText: 'Cancel',
        type: 'warning',
      }
    )

    authStore.logout()
    ElMessage.success('Logged out successfully!')

    router.push('/')
  } catch {
    // User cancelled logout
  }
}
</script>

<style scoped>
.profile-page {
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

.profile-content {
  display: grid;
  grid-template-columns: 2fr 1fr;
  gap: 2rem;
}

.user-info {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.info-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0.75rem 0;
  border-bottom: 1px solid var(--gray-200);
}

.info-item:last-child {
  border-bottom: none;
}

.info-item label {
  font-weight: 600;
  color: var(--gray-700);
}

.info-item span {
  color: var(--gray-800);
}

.action-buttons {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

/* Responsive Design */
@media (max-width: 768px) {
  .profile-content {
    grid-template-columns: 1fr;
  }

  .info-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 0.25rem;
  }
}
</style>