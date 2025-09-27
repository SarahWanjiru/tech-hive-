<template>
  <div class="register-page">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-md-6 col-lg-5">
          <div class="card">
            <div class="card-header text-center">
              <h3>Create Account</h3>
              <p class="text-muted">Join Mini-Ecommerce today</p>
            </div>

            <div class="card-body">
              <el-form
                ref="registerFormRef"
                :model="registerForm"
                :rules="registerRules"
                @submit.prevent="handleRegister"
                label-position="top"
              >
                <el-form-item label="Full Name" prop="name">
                  <el-input
                    v-model="registerForm.name"
                    placeholder="Enter your full name"
                    :disabled="authStore.isLoading"
                  />
                </el-form-item>

                <el-form-item label="Email" prop="email">
                  <el-input
                    v-model="registerForm.email"
                    type="email"
                    placeholder="Enter your email"
                    :disabled="authStore.isLoading"
                  />
                </el-form-item>

                <el-form-item label="Password" prop="password">
                  <el-input
                    v-model="registerForm.password"
                    type="password"
                    placeholder="Create a password"
                    show-password
                    :disabled="authStore.isLoading"
                  />
                </el-form-item>

                <el-form-item label="Confirm Password" prop="confirmPassword">
                  <el-input
                    v-model="registerForm.confirmPassword"
                    type="password"
                    placeholder="Confirm your password"
                    show-password
                    :disabled="authStore.isLoading"
                  />
                </el-form-item>

                <el-form-item label="Account Type" prop="role">
                  <el-radio-group v-model="registerForm.role">
                    <el-radio value="customer">Customer</el-radio>
                    <el-radio value="admin">Admin</el-radio>
                  </el-radio-group>
                </el-form-item>

                <el-form-item>
                  <el-button
                    type="primary"
                    native-type="submit"
                    class="w-100"
                    :loading="authStore.isLoading"
                    :disabled="authStore.isLoading"
                  >
                    {{ authStore.isLoading ? 'Creating account...' : 'Create Account' }}
                  </el-button>
                </el-form-item>

                <div v-if="authStore.error" class="alert alert-danger">
                  {{ authStore.error }}
                </div>
              </el-form>
            </div>

            <div class="card-footer text-center">
              <p class="mb-0">
                Already have an account?
                <el-link type="primary" @click="$router.push('/login')">
                  Sign in here
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
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth.store.ts'
import { ElMessage } from 'element-plus'

const router = useRouter()
const authStore = useAuthStore()

const registerFormRef = ref(null)

const registerForm = reactive({
  name: '',
  email: '',
  password: '',
  confirmPassword: '',
  role: 'customer'
})

const registerRules = {
  name: [
    { required: true, message: 'Please enter your full name', trigger: 'blur' },
    { min: 2, message: 'Name must be at least 2 characters', trigger: 'blur' }
  ],
  email: [
    { required: true, message: 'Please enter your email', trigger: 'blur' },
    { type: 'email', message: 'Please enter a valid email', trigger: 'blur' }
  ],
  password: [
    { required: true, message: 'Please enter a password', trigger: 'blur' },
    { min: 6, message: 'Password must be at least 6 characters', trigger: 'blur' }
  ],
  confirmPassword: [
    { required: true, message: 'Please confirm your password', trigger: 'blur' },
    {
      validator: (rule, value, callback) => {
        if (value !== registerForm.password) {
          callback(new Error('Passwords do not match'))
        } else {
          callback()
        }
      },
      trigger: 'blur'
    }
  ],
  role: [
    { required: true, message: 'Please select an account type', trigger: 'change' }
  ]
}

const handleRegister = async () => {
  try {
    await registerFormRef.value.validate()

    const response = await authStore.register({
      name: registerForm.name,
      email: registerForm.email,
      password: registerForm.password,
      role: registerForm.role
    })

    ElMessage.success('Account created successfully!')

    // Redirect to home page
    router.push('/')

  } catch (error) {
    console.error('Registration error:', error)
    // Error is already handled by the store
  }
}

onMounted(() => {
  // Clear any existing errors
  authStore.clearError()
})
</script>

<style scoped>
.register-page {
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