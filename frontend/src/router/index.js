import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: () => import('@/views/HomeView.vue'),
    },
    {
      path: '/health',
      name: 'health',
      component: () => import('@/views/HealthView.vue'),
    },
  ],
})

export default router
