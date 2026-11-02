import api from './api'

const STORAGE_KEY = 'studySmartAuth'

function getStoredAuth() {
  try {
    return JSON.parse(localStorage.getItem(STORAGE_KEY)) || {}
  } catch (err) {
    return {}
  }
}

function saveAuth(authData) {
  localStorage.setItem(STORAGE_KEY, JSON.stringify(authData))
  window.dispatchEvent(new Event('studySmartAuthChange'))
}

function login(credentials) {
  return api.post('/login', credentials).then(response => {
    if (response.status !== 200 && response.status !== 201) {
      throw new Error('Invalid login response')
    }

    const data = response.data || {}
    const token = data.token || data.accessToken || data.sessionToken || 'authenticated'

    if (!token) {
      throw new Error('Login response did not include a token')
    }

    saveAuth({
      token,
      username: credentials.username,
      loggedInAt: Date.now(),
      user: data.user || null,
    })

    return data
  })
}

function logout() {
  localStorage.removeItem(STORAGE_KEY)
  window.dispatchEvent(new Event('studySmartAuthChange'))
}

function getToken() {
  return getStoredAuth().token || null
}

function isLoggedIn() {
  return Boolean(getToken())
}

function getUser() {
  const authData = getStoredAuth()
  return authData.user || { username: authData.username || 'Admin' }
}

export default {
  login,
  logout,
  getToken,
  isLoggedIn,
  getUser,
}
