<template>
  <div>
    <div class="rounded-3xl bg-white p-8 shadow-xl ring-1 ring-slate-200">
      <p class="text-sm font-semibold uppercase tracking-[0.3em] text-emerald-600">Admin dashboard</p>
      <h1 class="mt-4 text-3xl font-bold text-slate-900">Welcome to your management center</h1>
      <p class="mt-3 text-slate-600">Use the sidebar to jump into user management or the data table. The dashboard view keeps the interface focused and fast.</p>
      <div class="mt-8 grid gap-4 sm:grid-cols-2">
        <router-link to="/users" class="rounded-2xl border border-slate-500  px-6 py-5 text-left transition hover:border-emerald-300 hover:bg-white row-auto">
          <p class="text-sm font-semibold text-slate-900 bg-green-500 ">Users</p>
          <p class="mt-2 text-sm text-slate-600">Manage your user accounts, roles, and active sessions.</p>
        </router-link>
        <router-link to="/table" class="rounded-2xl border border-slate-200 bg-slate-50 px-6 py-5 text-left transition hover:border-emerald-300 hover:bg-white">
          <p class="text-sm font-semibold text-slate-900">Table</p>
          <p class="mt-2 text-sm text-slate-600">Access the structured data table for quick analytics and edits.</p>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DashboardPage',
}
</script>

<style scoped>
</style>
