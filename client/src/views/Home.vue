<template>
  <div class="home-page">
    <!-- Header Section -->
    <header class="hero-section">
      <div class="container">
        <div class="hero-content">
          <div class="hero-text">
            <h1 class="hero-title">Welcome to Tech-Hive</h1>
            <p class="hero-subtitle">
              Discover the latest in technology with TechHive – your one-stop shop for phones, laptops, accessories, and more.
            </p>
            <div class="hero-actions">
              <el-button
                type="primary"
                size="large"
                @click="$router.push('/products')"
              >
                Shop Now
              </el-button>
              <el-button
                size="large"
                @click="$router.push('/register')"
              >
                Join Us
              </el-button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <!-- Features Section -->
    <section class="features-section">
      <div class="container">
        <div class="section-header">
          <h2>Why Choose Us?</h2>
          <p class="section-subtitle">Experience the best in online shopping</p>
        </div>

        <div class="features-grid">
          <div class="feature-card">
            <div class="feature-icon">
              🚚
            </div>
            <h3>Fast Delivery</h3>
            <p>Get your orders delivered quickly and safely to your doorstep.</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              🔒
            </div>
            <h3>Secure Payments</h3>
            <p>Your payment information is protected with industry-standard security.</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              🎧
            </div>
            <h3>24/7 Support</h3>
            <p>Our customer service team is always ready to help you.</p>
          </div>

          <div class="feature-card">
            <div class="feature-icon">
              💰
            </div>
            <h3>Best Prices</h3>
            <p>Competitive prices and great deals on all products.</p>
          </div>
        </div>
      </div>
    </section>

  </div>
</template>

<script setup>
import { useAuthStore } from '../stores/auth.store.ts'

const authStore = useAuthStore()
</script>

<style scoped>
.home-page {
  min-height: 100vh;
}

/* Hero Section */
.hero-section {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: var(--white);
  padding: 4rem 0;
  text-align: center;
}

.hero-content {
  max-width: 800px;
  margin: 0 auto;
}

.hero-title {
  font-size: 3rem;
  font-weight: 700;
  margin-bottom: 1rem;
}

.hero-subtitle {
  font-size: 1.25rem;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.hero-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}

/* Features Section */
.features-section {
  padding: 4rem 0;
  background: var(--white);
}

.section-header {
  text-align: center;
  margin-bottom: 3rem;
}

.section-header h2 {
  font-size: 2.5rem;
  color: var(--gray-800);
  margin-bottom: 1rem;
}

.section-subtitle {
  font-size: 1.125rem;
  color: var(--gray-600);
  margin-bottom: 0;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(250px, 1fr));
  gap: 2rem;
}

.feature-card {
  text-align: center;
  padding: 2rem;
  border-radius: var(--border-radius-lg);
  background: var(--gray-50);
  transition: transform var(--transition-fast);
}

.feature-card:hover {
  transform: translateY(-5px);
}

.feature-icon {
  width: 60px;
  height: 60px;
  margin: 0 auto 1rem;
  background: var(--primary-color);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--white);
  font-size: 24px;
}

.feature-card h3 {
  font-size: 1.25rem;
  margin-bottom: 1rem;
  color: var(--gray-800);
}


/* Responsive Design */
@media (max-width: 768px) {
  .hero-title {
    font-size: 2rem;
  }

  .hero-actions {
    flex-direction: column;
    align-items: center;
  }

  .features-grid {
    grid-template-columns: 1fr;
  }

}
</style>