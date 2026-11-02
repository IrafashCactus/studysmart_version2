<script setup>
import { RouterView, RouterLink } from 'vue-router'
</script>

<template>
  <div class="min-h-screen bg-gray-50">
    <!-- Navigation -->
    <nav class="bg-white shadow-sm border-b border-gray-200">
      <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
        <div class="flex justify-between h-16 items-center">
          <div class="flex items-center space-x-2">
            <span class="text-2xl">📚</span>
            <span class="text-xl font-bold text-primary-700">StudySmart</span>
          </div>
          <div class="flex space-x-6">
            <RouterLink
              to="/"
              class="text-gray-600 hover:text-primary-600 font-medium transition-colors"
            >
              Home
            </RouterLink>
            <RouterLink
              to="/health"
              class="text-gray-600 hover:text-primary-600 font-medium transition-colors"
            >
              Health
            </RouterLink>
          </div>
        </div>
      </div>
    </nav>

    <!-- Main Content -->
    <main class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
      <RouterView />
    </main>
  </div>
</template>
