import { createRouter, createWebHistory } from 'vue-router'
import auth from '../services/auth'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('../views/HomePage.vue')
  },
  {
    path: '/feed',
    name: 'Feed',
    component: () => import('../views/FeedView.vue')
  },
  {
    path: '/gamification',
    name: 'Gamification',
    component: () => import('../views/GamificationView.vue')
  },
  {
    path: '/tutor',
    name: 'Tutor',
    component: () => import('../views/TutorView.vue')
  },
  {
    path: '/educator',
    name: 'Educator',
    component: () => import('../views/EducatorView.vue')
  },
  {
    path: '/parent',
    name: 'Parent',
    component: () => import('../views/ParentView.vue')
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('../views/AdminView.vue')
  },
  {
    path: '/dashboard',
    redirect: '/admin'
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
    path: '/health',
    name: 'Health',
    component: () => import('../views/HealthView.vue')
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
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
})

export default router
