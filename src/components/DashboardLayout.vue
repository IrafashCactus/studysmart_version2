<template>
  <div class="dashboard-layout">
    <header class="dashboard-header">
      <div class="mx-auto flex max-w-7xl items-center justify-between px-4 py-4 sm:px-6 lg:px-8">
        <div class="flex items-center gap-3">
          <router-link to="/" class="flex items-center gap-3 text-slate-950">
            <span
              class="grid h-10 w-10 place-items-center rounded-lg bg-emerald-700 shadow text-lg font-black text-white">S</span>
            <span class="text-lg font-bold text-emerald-800">StudySmart</span>
          </router-link>

        </div>

        <div class="flex items-center gap-3">
          <div class="text-right text-sm">
            <p class="font-semibold text-slate-900">{{ displayName }}</p>
            <p class="text-slate-500">Administrator</p>
          </div>
          <button
            class="flex items-center gap-2 rounded-md border bg-white px-3 py-2 text-sm font-semibold text-slate-700 shadow-sm transition hover:scale-105"
            @click="toggleProfile">
            <span class="grid h-8 w-8 place-items-center rounded-full bg-emerald-700 text-xs font-bold text-white">{{
              initials }}</span>
            Admin Profile
            <span class="pi pi-angle-down text-xs"></span>
          </button>
        </div>
      </div>

      <div v-if="profileOpen" class="profile-dropdown">
        <div class="mx-auto max-w-7xl px-4 sm:px-6 lg:px-8">
          <div class="rounded-2xl border border-slate-200 bg-white p-4 shadow-lg">
            <div class="flex items-center justify-between gap-4">
              <div>
                <p class="text-sm font-semibold text-slate-900">{{ displayName }}</p>
                <p class="text-xs text-slate-500">Platform administrator</p>
              </div>
              <button class="text-sm font-semibold text-red-600 hover:text-red-800" @click="logout">Logout</button>
            </div>
          </div>
        </div>
      </div>
    </header>

    <Transition name="banner-drop">
      <section v-if="isDashboard" class="dashboard-banner">
        <div
          class="mx-auto flex max-w-7xl flex-col gap-3 px-4 py-5 sm:px-6 lg:flex-row lg:items-center lg:justify-between lg:px-8">
          <div>
            <p class="text-xs font-semibold uppercase tracking-[0.25em] text-emerald-100">Admin workspace</p>
            <h1 class="mt-2 text-3xl font-bold text-white">Welcome, Admin</h1>
          </div>
          <!-- <p class="max-w-2xl text-sm text-emerald-100">Your command center is ready. Jump into user management and data tables with a clean, fixed admin layout.</p> -->
        </div>
      </section>
    </Transition>

    <aside class="admin-sidebar">
      <div class="mb-6">
        <p class="text-xs font-semibold uppercase tracking-[0.2em] text-slate-400">Navigation</p>
        <h2 class="mt-2 text-lg font-bold text-slate-900">Management</h2>
      </div>
      <nav class="space-y-2">
        <div class="bg-green-500"> <router-link to="/users"
            class="block rounded-md px-3 py-3 text-sm font-semibold text-slate-600 transition hover:text-emerald-700"
            active-class="bg-emerald-50 text-emerald-700">Users</router-link></div>
        <div class="bg-green-500"> <router-link to="/table"
            class="block rounded-md px-3 py-3 text-sm font-semibold text-slate-600 transition  hover:text-emerald-700"
            active-class="bg-emerald-50 text-emerald-700">Table</router-link></div>
      </nav>
    </aside>

    <main class="dashboard-main" :style="mainStyles">
      <div class="mx-auto max-w-7xl px-4 py-6 sm:px-6 lg:px-8">
        <slot />
      </div>
    </main>
  </div>
</template>

<script>
import auth from '../services/auth'

export default {
  name: 'DashboardLayout',
  data() {
    return {
      profileOpen: false,
      user: auth.getUser(),
    }
  },
  computed: {
    displayName() {
      return this.user?.name || this.user?.username || 'Admin'
    },
    initials() {
      return this.displayName
        .split(' ')
        .map((part) => part[0])
        .slice(0, 2)
        .join('')
        .toUpperCase()
    },
    isDashboard() {
      return this.$route.name === 'Dashboard'
    },
    mainStyles() {
      return {
        paddingTop: this.isDashboard ? '8rem' : '4rem',
        marginLeft: '18rem',
      }
    },
  },
  methods: {
    toggleProfile() {
      this.profileOpen = !this.profileOpen
    },
    logout() {
      auth.logout()
      this.$router.push('/')
    },
  },
}
</script>

<style scoped>
.dashboard-header {
  position: fixed;
  inset: 0 auto auto 0;
  z-index: 50;
  width: 100%;
  border-bottom: 1px solid #e2e8f0;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(12px);
}

.profile-dropdown {
  position: fixed;
  top: 5.5rem;
  left: 0;
  right: 0;
  z-index: 40;
}

.dashboard-banner {
  position: fixed;
  top: 4rem;
  left: 0;
  right: 0;
  z-index: 40;
  border-bottom: 1px solid rgba(16, 185, 129, 0.12);
  background: linear-gradient(90deg, #065f46 0%, #047857 100%);
}

.admin-sidebar {
  position: fixed;
  top: 4rem;
  bottom: 0;
  left: 0;
  width: 18rem;
  padding: 1.5rem;
  background: linear-gradient(180deg, #ecfdf5 0%, #ffffff 100%);
  border-right: 1px solid rgba(16, 185, 129, 0.06);
  box-shadow: 0 12px 30px rgba(6, 95, 70, 0.06);
  z-index: 20;
}

.dashboard-main {
  min-height: 100vh;
  background: linear-gradient(180deg, #f8fafc 0%, #f0fdf4 100%);
}

@media (max-width: 1024px) {
  .admin-sidebar {
    display: none;
  }

  .dashboard-main {
    margin-left: 0 !important;
  }
}
</style>
