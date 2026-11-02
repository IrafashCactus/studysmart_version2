import { createRouter, createWebHistory } from 'vue-router'
import auth from '../services/auth'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/HomePage.vue')
  },
  {
    path: '/users',
    name: 'Users',
    component: () => import('../views/UsersPage.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/table',
    name: 'Table',
    component: () => import('../views/TablePage.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: () => import('../views/LoginPage.vue')
  },
  {
    path: '/signup',
    name: 'Signup',
    component: () => import('../views/SignupPage.vue')
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: () => import('../views/DashboardPage.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/health',
    name: 'Health',
    component: () => import('../views/HealthView.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to) => {
  const loggedIn = auth.isLoggedIn()

  if (to.meta.requiresAuth && !loggedIn) {
    return { name: 'Login' }
  }

  if ((to.name === 'Login' || to.name === 'Signup') && loggedIn) {
    return { name: 'Home' }
  }

  return true
})

export default router
