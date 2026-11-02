<script setup>
import { onMounted } from 'vue'
import { useHealthStore } from '@/stores/health.js'

const healthStore = useHealthStore()

onMounted(() => {
  healthStore.fetchHealth()
})

function refresh() {
  healthStore.fetchHealth()
}
</script>

<template>
  <div class="max-w-2xl mx-auto py-8">
    <div class="flex items-center justify-between mb-6">
      <h1 class="text-3xl font-bold text-gray-900">System Health</h1>
      <button
        @click="refresh"
        :disabled="healthStore.loading"
        class="px-4 py-2 bg-primary-600 text-white text-sm font-medium rounded-lg hover:bg-primary-700 disabled:opacity-50 transition-colors"
      >
        {{ healthStore.loading ? 'Checking...' : 'Refresh' }}
      </button>
    </div>

    <!-- Loading State -->
    <div v-if="healthStore.loading && !healthStore.status" class="text-center py-12">
      <div class="animate-spin inline-block w-8 h-8 border-4 border-primary-200 border-t-primary-600 rounded-full"></div>
      <p class="mt-4 text-gray-500">Checking system health...</p>
    </div>

    <!-- Error State -->
    <div
      v-else-if="healthStore.error"
      class="bg-red-50 border border-red-200 rounded-xl p-6"
    >
      <div class="flex items-center space-x-3">
        <span class="text-2xl">❌</span>
        <div>
          <h3 class="font-semibold text-red-800">Connection Failed</h3>
          <p class="text-sm text-red-600 mt-1">{{ healthStore.error }}</p>
        </div>
      </div>
    </div>

    <!-- Success State -->
    <div
      v-else-if="healthStore.status"
      class="bg-white border border-gray-200 rounded-xl shadow-sm overflow-hidden"
    >
      <!-- Overall Status -->
      <div class="px-6 py-4 border-b border-gray-100 flex items-center justify-between">
        <span class="font-semibold text-gray-700">Overall Status</span>
        <span
          class="inline-flex items-center px-3 py-1 rounded-full text-sm font-medium"
          :class="healthStore.status.success
            ? 'bg-green-100 text-green-800'
            : 'bg-red-100 text-red-800'"
        >
          {{ healthStore.status.success ? '✅ Healthy' : '⚠️ Degraded' }}
        </span>
      </div>

      <!-- Service Details -->
      <div class="divide-y divide-gray-100">
        <div
          v-for="(value, key) in healthStore.status.data"
          :key="key"
          class="px-6 py-3 flex items-center justify-between"
        >
          <span class="text-gray-600 capitalize">{{ key }}</span>
          <span
            class="inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium"
            :class="value === 'up'
              ? 'bg-green-100 text-green-700'
              : 'bg-red-100 text-red-700'"
          >
            {{ value === 'up' ? '● Up' : '● Down' }}
          </span>
        </div>
      </div>

      <!-- Message -->
      <div class="px-6 py-3 bg-gray-50 text-sm text-gray-500">
        {{ healthStore.status.message }}
      </div>
    </div>
  </div>
</template>
