<template>
  <div class="min-h-screen bg-slate-950 text-white">
    <section class="relative isolate flex min-h-screen items-center overflow-hidden">
      <img src="../assets/hero.png" alt="" class="absolute inset-0 h-full w-full object-cover opacity-35" />
      <div class="absolute inset-0 bg-slate-950/70"></div>

      <div class="relative mx-auto grid w-full max-w-7xl gap-10 px-6 py-16 lg:grid-cols-[1.05fr_0.95fr] lg:items-center">
        <div class="max-w-3xl">
          <p class="mb-5 inline-flex rounded-md bg-emerald-400/15 px-4 py-2 text-sm font-semibold text-emerald-200 ring-1 ring-emerald-300/20">
            StudySmart learning workspace
          </p>
          <h1 class="text-5xl font-black leading-tight text-white sm:text-6xl lg:text-7xl">
            Learn smarter. Manage sharper.
          </h1>
          <p class="mt-6 max-w-2xl text-lg leading-8 text-slate-200">
            A focused student-first entry point that stays quiet until you sign in, then opens the tools you need when you ask for the dashboard.
          </p>
          <div class="mt-8 flex flex-wrap gap-3">
            <button v-if="!loggedIn" type="button" class="rounded-md bg-emerald-400 px-6 py-3 text-sm font-bold text-slate-950 shadow-lg shadow-emerald-950/30 transition hover:bg-emerald-300" @click="showLoginModal = true">
              Login to Continue
            </button>
            <router-link v-else to="/dashboard" class="rounded-md bg-emerald-400 px-6 py-3 text-sm font-bold text-slate-950 shadow-lg shadow-emerald-950/30 transition hover:bg-emerald-300">
              Open Dashboard
            </router-link>
            <router-link to="/health" class="rounded-md border border-white/30 px-6 py-3 text-sm font-semibold text-white transition hover:bg-white/10">
              Check system health
            </router-link>
          </div>
        </div>

        <div class="hidden rounded-lg border border-white/10 bg-white/10 p-6 shadow-2xl backdrop-blur md:block">
          <div class="grid gap-4">
            <div class="rounded-md bg-slate-950/70 p-4">
              <p class="text-sm text-slate-300">Current mode</p>
              <p class="mt-1 text-2xl font-bold">{{ loggedIn ? 'You got it!' : 'Splash only' }}</p>
            </div>
            <div class="grid grid-cols-2 gap-4">
              <div class="rounded-md bg-white/10 p-4">
                <p class="text-sm text-slate-300">Student entry</p>
                <p class="mt-2 text-3xl font-black text-emerald-300">Clean</p>
              </div>
              <div class="rounded-md bg-white/10 p-4">
                <p class="text-sm text-slate-300">Admin tools</p>
                <p class="mt-2 text-3xl font-black text-cyan-300">Ready</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <template v-if="loggedIn">
      <section id="about" class="bg-white px-6 py-16 text-slate-900">
        <div class="mx-auto max-w-7xl">
          <p class="text-sm font-semibold uppercase tracking-[0.25em] text-emerald-600">About Us</p>
          <h2 class="mt-3 text-3xl font-bold">A calmer way to move from study mode to admin mode.</h2>
          <p class="mt-4 max-w-3xl leading-7 text-slate-600">StudySmart keeps students focused first, then reveals the management surface only after the dashboard is opened.</p>
        </div>
      </section>

      <section id="services" class="bg-slate-100 px-6 py-16 text-slate-900">
        <div class="mx-auto max-w-7xl">
          <div class="flex flex-col gap-4 lg:flex-row lg:items-end lg:justify-between">
            <div>
              <p class="text-sm font-semibold uppercase tracking-[0.25em] text-emerald-600">Services</p>
              <h2 class="mt-3 text-3xl font-bold">Core management tools</h2>
            </div>
            <router-link to="/dashboard" class="inline-flex rounded-md bg-slate-950 px-5 py-3 text-sm font-semibold text-white transition hover:bg-emerald-700">
              Go to Dashboard
            </router-link>
          </div>

          <div class="mt-8">
            <div v-if="loadingFeatures" class="rounded-md bg-white p-6 text-center text-slate-500">
              Loading services...
            </div>
            <div v-else-if="featuresError" class="rounded-md bg-red-50 p-6 text-center text-red-600">
              {{ featuresError }}
            </div>
            <div v-else class="grid gap-5 md:grid-cols-2 lg:grid-cols-4">
              <div v-for="item in (features.length ? features : serviceItems)" :key="item.id" class="rounded-md bg-white p-6 shadow-sm ring-1 ring-slate-200">
                <h3 class="text-lg font-bold text-slate-900">{{ item.title }}</h3>
                <p class="mt-3 text-sm leading-6 text-slate-600">{{ item.description }}</p>
              </div>
            </div>
          </div>
        </div>
      </section>
    </template>

    <Transition name="modal-fade">
      <div v-if="showLoginModal" class="fixed inset-0 z-50 flex items-center justify-center bg-slate-950/75 p-4 backdrop-blur" @click.self="showLoginModal = false">
        <div class="relative w-full max-w-lg rounded-lg bg-white p-2 shadow-2xl">
          <button type="button" class="absolute right-4 top-4 z-10 grid h-9 w-9 place-items-center rounded-md text-slate-500 transition hover:bg-slate-100 hover:text-slate-900" aria-label="Close login modal" @click="showLoginModal = false">
            <span class="pi pi-times"></span>
          </button>
          <LoginForm />
        </div>
      </div>
    </Transition>
  </div>
</template>

<script>
import auth from '../services/auth'
import api from '../services/api'
import LoginForm from '../components/Loginform.vue'

export default {
  name: 'HomePage',
  components: { LoginForm },
  data() {
    return {
      loggedIn: auth.isLoggedIn(),
      showLoginModal: false,
      features: [],
      loadingFeatures: false,
      featuresError: null,
      serviceItems: [
        { id: 'users', title: 'User Management', description: 'Create, review, edit, and remove user records in a focused way.' },
        { id: 'table', title: 'Data Table', description: 'View and filter user rows with responsive, simple controls.' },
        { id: 'dashboard', title: 'Dashboard', description: 'Open the management command center with analytics and recent actions.' },
        { id: 'focus', title: 'Focused Entry', description: 'The splash page stays calm until you choose to sign in.' },
      ],
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
        const res = await api.get('homeitems')
        const items = res?.data?.data ?? res?.data ?? []
        if (Array.isArray(items)) {
          this.features = items.map((item, index) => ({
            id: item.id ?? index,
            title: item.title || item.name || 'No title',
            description: item.description || item.summary || '',
          }))
        } else {
          this.features = []
          this.featuresError = 'Unexpected response shape from homeitems API.'
        }
      } catch (err) {
        this.featuresError = 'Unable to load feature items.'
        console.error('HomePage loadHomeItems error:', err)
      } finally {
        this.loadingFeatures = false
      }
    },
    updateAuthStatus() {
      this.loggedIn = auth.isLoggedIn()
      if (this.loggedIn) this.showLoginModal = false
    }
  }
}
</script>

<style scoped>
.modal-fade-enter-active,
.modal-fade-leave-active {
  transition: opacity 0.25s ease;
}

.modal-fade-enter-from,
.modal-fade-leave-to {
  opacity: 0;
}
</style>
