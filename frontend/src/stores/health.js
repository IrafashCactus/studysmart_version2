import { defineStore } from 'pinia'
import { ref } from 'vue'
import apiClient from '@/services/api.js'

export const useHealthStore = defineStore('health', () => {
  const status = ref(null)
  const loading = ref(false)
  const error = ref(null)

  async function fetchHealth() {
    loading.value = true
    error.value = null
    try {
      const response = await apiClient.get('/api/v1/health')
      status.value = response
    } catch (err) {
      error.value = err.response?.data?.message || err.message || 'Failed to fetch health status'
      status.value = null
    } finally {
      loading.value = false
    }
  }

  return { status, loading, error, fetchHealth }
})
