import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import api, { getStoredAuth } from '../services/api'

export const useAuthStore = defineStore('auth', () => {
  const stored = getStoredAuth()
  const token = ref(stored.token || null)
  const user = ref(stored.user || null)
  const language = ref(localStorage.getItem('studySmartLang') || 'en')
  const loading = ref(false)
  const error = ref(null)

  const isLoggedIn = computed(() => !!token.value)
  const isStudent = computed(() => !user.value || user.value.role === 'student')
  const isTeacher = computed(() => user.value?.role === 'teacher')
  const isParent = computed(() => user.value?.role === 'parent')
  const isAdmin = computed(() => user.value?.role === 'super_admin' || user.value?.role === 'school_admin' || user.value?.role === 'admin')

  function saveAuthData(data) {
    token.value = data.token
    user.value = data.user
    localStorage.setItem('studySmartAuth', JSON.stringify({
      token: data.token,
      user: data.user,
      loggedInAt: Date.now(),
    }))
  }

  async function login(email, password) {
    loading.value = true
    error.value = null
    try {
      const res = await api.post('/api/v1/auth/login', { email, password })
      saveAuthData(res.data.data)
      return res.data.data
    } catch (err) {
      error.value = err.response?.data?.message || 'Login failed'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  async function demoLogin(role = 'student') {
    loading.value = true
    error.value = null
    try {
      const res = await api.post(`/api/v1/auth/demo?role=${role}`)
      saveAuthData(res.data.data)
      return res.data.data
    } catch (err) {
      error.value = err.response?.data?.message || 'Demo login failed'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  async function register(email, password, fullName, role = 'student', lang = 'en') {
    loading.value = true
    error.value = null
    try {
      const res = await api.post('/api/v1/auth/register', { email, password, fullName, role, language: lang })
      saveAuthData(res.data.data)
      return res.data.data
    } catch (err) {
      error.value = err.response?.data?.message || 'Registration failed'
      throw error.value
    } finally {
      loading.value = false
    }
  }

  function logout() {
    token.value = null
    user.value = null
    localStorage.removeItem('studySmartAuth')
  }

  function setLanguage(lang) {
    language.value = lang
    localStorage.setItem('studySmartLang', lang)
  }

  // Auto-login to Keza (demo student) on first app load if not logged in
  async function initAuth() {
    if (!token.value) {
      try {
        await demoLogin('student')
      } catch (e) {
        // backend might still be starting
      }
    }
  }

  return {
    token,
    user,
    language,
    loading,
    error,
    isLoggedIn,
    isStudent,
    isTeacher,
    isParent,
    isAdmin,
    login,
    demoLogin,
    register,
    logout,
    setLanguage,
    initAuth,
  }
})

