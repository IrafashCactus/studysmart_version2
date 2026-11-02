<template>
  <nav class="sticky top-0 z-40 border-b bg-white/95 shadow-sm backdrop-blur transition-all" style="border-color: rgba(16,185,129,0.12)">
    <div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
      <router-link to="/" class="flex items-center gap-3 text-slate-950">
        <span class="grid h-10 w-10 place-items-center rounded-lg bg-emerald-700 shadow-md text-lg font-black text-white">S</span>
        <span class="text-lg font-bold text-emerald-800">StudySmart</span>
      </router-link>

      <div class="hidden items-center gap-1 md:flex">
        <router-link to="/" class="rounded-md px-3 py-2 text-sm font-medium text-slate-600 transition hover:bg-emerald-50 hover:text-emerald-700">Home</router-link>
        <router-link to="/#about" class="rounded-md px-3 py-2 text-sm font-medium text-slate-600 transition hover:bg-emerald-50 hover:text-emerald-700">About Us</router-link>
        <router-link to="/#services" class="rounded-md px-3 py-2 text-sm font-medium text-slate-600 transition hover:bg-emerald-50 hover:text-emerald-700">Services</router-link>
        <router-link to="/dashboard" class="rounded-md px-3 py-2 text-sm font-medium text-slate-600 transition hover:bg-emerald-50 hover:text-emerald-700">Dashboard</router-link>
      </div>

      <button type="button" class="rounded-md border bg-emerald-600 px-3 py-2 text-sm font-semibold text-white shadow-sm transition hover:bg-emerald-500" @click="logout">
        Logout
      </button>
    </div>
  </nav>
</template>

<script>
import auth from '../services/auth'

export default {
  name: 'NavBar',
  methods: {
    logout() {
      auth.logout()
      this.$router.push('/')
    },
  },
}
</script>

<style scoped>
</style>
