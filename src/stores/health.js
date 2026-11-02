import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'

export const useHealthStore = defineStore('health', () => {
  const status = ref(null)
  const loading = ref(false)
  const error = ref('')

  async function fetchHealth() {
    loading.value = true
    error.value = ''
    try {
      const response = await api.get('/api/v1/health')
      status.value = response.data
    } catch (requestError) {
      status.value = null
      error.value = requestError.response?.data?.message || requestError.message || 'Health check failed'
    } finally {
      loading.value = false
    }
  }

  return { status, loading, error, fetchHealth }
})

