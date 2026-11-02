<script setup>
// Home view — landing page
</script>

<template>
  <div class="flex flex-col items-center justify-center py-16 space-y-8">
    <div class="text-center space-y-4">
      <h1 class="text-5xl font-extrabold text-gray-900">
        Welcome to <span class="text-primary-600">StudySmart</span>
      </h1>
      <p class="text-lg text-gray-500 max-w-2xl">
        A modern full-stack application built with Vue 3, Go, and PostgreSQL.
        Containerized with Docker for seamless development and deployment.
      </p>
    </div>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 w-full max-w-4xl mt-8">
      <!-- Vue Card -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 text-center hover:shadow-md transition-shadow">
        <div class="text-4xl mb-3">⚡</div>
        <h3 class="text-lg font-semibold text-gray-800">Vue 3 Frontend</h3>
        <p class="text-sm text-gray-500 mt-2">
          Composition API, Vite, Tailwind CSS, Pinia & Vue Router
        </p>
      </div>

      <!-- Go Card -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 text-center hover:shadow-md transition-shadow">
        <div class="text-4xl mb-3">🚀</div>
        <h3 class="text-lg font-semibold text-gray-800">Go Backend</h3>
        <p class="text-sm text-gray-500 mt-2">
          Gin framework, Clean Architecture, GORM & PostgreSQL
        </p>
      </div>

      <!-- Docker Card -->
      <div class="bg-white rounded-xl shadow-sm border border-gray-200 p-6 text-center hover:shadow-md transition-shadow">
        <div class="text-4xl mb-3">🐳</div>
        <h3 class="text-lg font-semibold text-gray-800">Docker Ready</h3>
        <p class="text-sm text-gray-500 mt-2">
          Multi-stage builds, Docker Compose, hot reload in dev
        </p>
      </div>
    </div>

    <RouterLink
      to="/health"
      class="mt-6 inline-flex items-center px-6 py-3 bg-primary-600 text-white font-medium rounded-lg hover:bg-primary-700 transition-colors shadow-sm"
    >
      Check System Health →
    </RouterLink>
  </div>
</template>
