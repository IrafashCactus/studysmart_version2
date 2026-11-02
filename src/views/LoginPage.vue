<template>
  <div class="min-h-screen bg-gradient-to-b from-emerald-950 via-emerald-900 to-emerald-800 text-white">
    <div class="mx-auto max-w-7xl px-6 py-8">
      <!-- Hero Section -->
      <section class="grid gap-10 lg:grid-cols-[1.2fr_0.8fr] lg:items-center">
        <div class="space-y-6">
          <div class="inline-flex items-center gap-2 rounded-full bg-white/10 px-4 py-2 text-sm">
            <span class="inline-block h-2 w-2 rounded-full bg-emerald-400"></span>
            Trusted by students and educators worldwide
          </div>
          <h1 class="text-4xl font-extrabold tracking-tight sm:text-5xl">
            Learn smarter, perform better.
          </h1>
          <p class="max-w-xl text-base leading-7 text-emerald-200">
            A modern learning experience designed to help students study efficiently with personalized plans, progress tracking, and interactive tools.
          </p>
          <div class="flex flex-wrap gap-3">
            <router-link to="/login" class="rounded-full bg-emerald-400 px-6 py-3 text-sm font-semibold text-emerald-950 transition hover:bg-emerald-300">
              Get Started
            </router-link>
            <router-link to="/users" class="rounded-full border border-white/20 bg-white/10 px-6 py-3 text-sm font-semibold transition hover:bg-white/20">
              View Users
            </router-link>
          </div>
        </div>

        <div class="rounded-[2rem] bg-white/10 p-8 shadow-2xl ring-1 ring-white/10 backdrop-blur-xl">
          <template v-if="loggedIn">
            <div class="space-y-4">
              <p class="text-sm uppercase tracking-[0.3em] text-emerald-300">Dashboard</p>
              <h3 class="text-2xl font-semibold">Welcome back to StudySmart.</h3>
              <p class="text-emerald-200">Your dashboard is ready. Navigate to the table or user management pages to continue.</p>
            </div>
            <div class="mt-8 grid gap-4 sm:grid-cols-2">
              <router-link to="/table" class="rounded-3xl bg-emerald-500 px-5 py-6 text-center text-white shadow-xl transition hover:bg-emerald-600">
                <p class="text-lg font-semibold">Table View</p>
                <p class="text-sm text-emerald-100 mt-2">Open the PrimeVue data table.</p>
              </router-link>
              <router-link to="/users" class="rounded-3xl bg-white/10 px-5 py-6 text-center text-white shadow-xl transition hover:bg-white/20">
                <p class="text-lg font-semibold">Users Page</p>
                <p class="text-sm text-emerald-100 mt-2">Manage users and view your list.</p>
              </router-link>
            </div>
          </template>
          <template v-else>
            <div class="space-y-4 mb-8">
              <p class="text-sm uppercase tracking-[0.3em] text-emerald-300">Sign in</p>
              <h3 class="text-2xl font-semibold">Secure login for your study dashboard.</h3>
              <p class="text-emerald-200">Log in to access your table, user management, and progress tools.</p>
            </div>
            <LoginForm />
          </template>
        </div>
      </section>

      <!-- Features Section -->
      <section class="mt-12 rounded-[2rem] bg-white/5 p-6 shadow-2xl ring-1 ring-white/10 backdrop-blur-xl">
        <div class="flex flex-col gap-4 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <p class="text-sm uppercase tracking-[0.3em] text-emerald-300">Featured tools</p>
            <h3 class="text-2xl font-semibold">Explore the core features</h3>
          </div>
          <p class="max-w-2xl text-emerald-200">Switch between Users and Table to see different data views.</p>
        </div>

        <div class="mt-6">
          <div v-if="loadingFeatures" class="rounded-3xl bg-white/10 p-6 text-center text-emerald-200">
            Loading features...
          </div>
          <div v-else-if="featuresError" class="rounded-3xl bg-red-500/10 p-6 text-center text-red-200">
            {{ featuresError }}
          </div>
          <div v-else class="grid gap-6 md:grid-cols-2 lg:grid-cols-4">
            <div v-for="item in features" :key="item.id" class="rounded-3xl bg-white/10 p-6 shadow-xl ring-1 ring-white/10 transition hover:-translate-y-1 hover:bg-white/15">
              <h4 class="text-lg font-semibold">{{ item.title }}</h4>
              <p class="mt-3 text-sm leading-6 text-emerald-200">{{ item.description }}</p>
            </div>
          </div>
        </div>
      </section>
    </div>
  </div>
</template>

<script>
import auth from '../services/auth'
import api from '../services/api'
import LoginForm from '../components/Loginform.vue'

export default {
  name: 'HomePage',
  components: {
    LoginForm,
  },
  data() {
    return {
      loggedIn: auth.isLoggedIn(),
      features: [],
      loadingFeatures: false,
      featuresError: null,
    }
  },
  mounted() {
    window.addEventListener('studySmartAuthChange', this.updateAuthStatus)
    this.loadHomeItems()
  },
  beforeUnmount() {
    window.removeEventListener('studySmartAuthChange', this.updateAuthStatus)
  },
  methods: {
    async loadHomeItems() {
      this.loadingFeatures = true
      this.featuresError = null
      try {
        const response = await api.get('homeitems')
        const items = response?.data?.data
        if (Array.isArray(items)) {
          this.features = items.map((item, index) => ({
            id: item.id ?? index,
            title: item.title || 'No title',
            description: item.description || 'No description available.',
          }))
        } else {
          this.featuresError = 'Unexpected response from home items API.'
        }
      } catch (error) {
        this.featuresError = 'Unable to load feature items.'
        console.error('HomePage loadHomeItems error:', error)
      } finally {
        this.loadingFeatures = false
      }
    },
    updateAuthStatus() {
      this.loggedIn = auth.isLoggedIn()
    }
  }
}
</script>