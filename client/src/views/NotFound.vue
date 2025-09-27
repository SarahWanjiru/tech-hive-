<template>
  <div class="not-found-page">
    <div class="container">
      <div class="not-found-content">
        <div class="error-code">404</div>
        <h1>Page Not Found</h1>
        <p>The page you're looking for doesn't exist or has been moved.</p>
        <div class="not-found-actions">
          <el-button type="primary" @click="$router.push('/')">
            Go Home
          </el-button>
          <el-button @click="$router.go(-1)">
            Go Back
          </el-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
// Component logic here
</script>

<style scoped>
.not-found-page {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: var(--white);
  text-align: center;
}

.not-found-content {
  max-width: 500px;
}

.error-code {
  font-size: 8rem;
  font-weight: 700;
  margin-bottom: 1rem;
  opacity: 0.8;
}

.not-found-content h1 {
  font-size: 2.5rem;
  margin-bottom: 1rem;
  color: var(--white);
}

.not-found-content p {
  font-size: 1.125rem;
  margin-bottom: 2rem;
  opacity: 0.9;
}

.not-found-actions {
  display: flex;
  gap: 1rem;
  justify-content: center;
  flex-wrap: wrap;
}
</style>