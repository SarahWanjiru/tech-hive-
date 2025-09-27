<template>
  <div class="login-page">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-md-6 col-lg-4">
          <div class="card">
            <div class="card-header text-center">
              <h3>Welcome Back</h3>
              <p class="text-muted">Sign in to your account</p>
            </div>

            <div class="card-body">
              <el-form
                ref="loginFormRef"
                :model="loginForm"
                :rules="loginRules"
                @submit.prevent="handleLogin"
                label-position="top"
              >
                <el-form-item label="Email" prop="email">
                  <el-input
                    v-model="loginForm.email"
                    type="email"
                    placeholder="Enter your email"
                    :disabled="authStore.isLoading"
                  />
                </el-form-item>

                <el-form-item label="Password" prop="password">
                  <el-input
                    v-model="loginForm.password"
                    type="password"
                    placeholder="Enter your password"
                    show-password
                    :disabled="authStore.isLoading"
                  />
                </el-form-item>

                <el-form-item>
                  <el-button
                    type="primary"
                    native-type="submit"
                    class="w-100"
                    :loading="authStore.isLoading"
                    :disabled="authStore.isLoading"
                  >
                    {{ authStore.isLoading ? 'Signing in...' : 'Sign In' }}
                  </el-button>
                </el-form-item>

                <div v-if="authStore.error" class="alert alert-danger">
                  {{ authStore.error }}
                </div>
              </el-form>
            </div>

            <div class="card-footer text-center">
              <p class="mb-0">
                Don't have an account?
                <el-link type="primary" @click="$router.push('/register')">
                  Sign up here
                </el-link>
              </p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth.store.js'
import { ElMessage } from 'element-plus'

const router = useRouter()
const route = useRoute()
const authStore = useAuthStore()

const loginFormRef = ref(null)

const loginForm = reactive({
  email: '',
  password: ''
})

const loginRules = {
  email: [
    { required: true, message: 'Please enter your email', trigger: 'blur' },
    { type: 'email', message: 'Please enter a valid email', trigger: 'blur' }
  ],
  password: [
    { required: true, message: 'Please enter your password', trigger: 'blur' },
    { min: 6, message: 'Password must be at least 6 characters', trigger: 'blur' }
  ]
}

const handleLogin = async () => {
  try {
    await loginFormRef.value.validate()

    const response = await authStore.login({
      email: loginForm.email,
      password: loginForm.password
    })

    ElMessage.success('Login successful!')

    // Redirect to intended page or home
    const redirectPath = route.query.redirect || '/'
    router.push(redirectPath)

  } catch (error) {
    console.error('Login error:', error)
    // Error is already handled by the store
  }
}

onMounted(() => {
  // Clear any existing errors
  authStore.clearError()
})
</script>

<style scoped>
.login-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.card {
  border: none;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.1);
}

.card-header {
  background: var(--white);
  border-bottom: 1px solid var(--gray-200);
}

.card-body {
  padding: 2rem;
}

.card-footer {
  background: var(--gray-50);
  border-top: 1px solid var(--gray-200);
}

.w-100 {
  width: 100%;
}

.alert {
  padding: 0.75rem;
  border-radius: var(--border-radius-md);
  margin-bottom: var(--spacing-md);
}

.alert-danger {
  background-color: #fef2f2;
  border: 1px solid #fecaca;
  color: #dc2626;
}
</style>