import axios from 'axios'

const api = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080',
  headers: {
    'Content-Type': 'application/json',
  },
})

export function getStoredAuth() {
  try {
    return JSON.parse(localStorage.getItem('studySmartAuth')) || {}
  } catch (err) {
    return {}
  }
}

export function getStoredToken() {
  const auth = getStoredAuth()
  return auth.token || null
}

api.interceptors.request.use(
  (config) => {
    const token = getStoredToken()
    if (token) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => Promise.reject(error)
)

export default api
