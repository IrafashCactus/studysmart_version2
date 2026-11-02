import axios from 'axios'

const api = axios.create({
  baseURL: 'http://127.0.0.1:3658/m1/1301610-1300998-default/',
  headers: {
    'Content-Type': 'application/json',
  },
})

function getStoredToken() {
  try {
    const stored = localStorage.getItem('studySmartAuth')
    return stored ? JSON.parse(stored).token : null
  } catch (err) {
    return null
  }
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