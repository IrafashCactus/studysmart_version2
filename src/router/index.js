import { createRouter, createWebHistory } from 'vue-router'
import HomePage from '../views/HomePage.vue'
import UsersPage from '../views/UsersPage.vue'
import TablePage from '../views/TablePage.vue'
import LoginPage from '../views/LoginPage.vue'
import SignupPage from '../views/SignupPage.vue'
import auth from '../services/auth'
import DashboardPage from '../views/DashboardPage.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HomePage
  },
  {
    path: '/users',
    name: 'Users',
    component: UsersPage,
    meta: { requiresAuth: true }
  },
  {
    path: '/table',
    name: 'Table',
    component: TablePage,
    meta: { requiresAuth: true }
  },
  {
    path: '/login',
    name: 'Login',
    component: LoginPage
  },
  {
    path: '/signup',
    name: 'Signup',
    component: SignupPage
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: DashboardPage,
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to, from, next) => {
  const loggedIn = auth.isLoggedIn()

  if (to.meta.requiresAuth && !loggedIn) {
    return next({ name: 'Login' })
  }

  if ((to.name === 'Login' || to.name === 'Signup') && loggedIn) {
    return next({ name: 'Home' })
  }

  return next()
})

export default router