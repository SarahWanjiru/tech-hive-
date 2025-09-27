import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'

// Import App component
import App from './App.vue'

// Import global styles
import './assets/styles/main.css'

// Import router
import router from './router/index.js'

// Import stores
import { useAuthStore } from './stores/auth.store.js'

const app = createApp(App)
const pinia = createPinia()

// Register Element Plus icons globally
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

// Use plugins
app.use(pinia)
app.use(router)
app.use(ElementPlus)

// Initialize auth state
const authStore = useAuthStore()
authStore.initializeAuth()

app.mount('#app')