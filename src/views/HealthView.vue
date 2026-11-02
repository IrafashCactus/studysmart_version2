<script setup>
import { onMounted } from 'vue'
import { storeToRefs } from 'pinia'
import { useHealthStore } from '../stores/health'

const health = useHealthStore()
const { status, loading, error } = storeToRefs(health)

onMounted(health.fetchHealth)
</script>

<template>
  <section class="mx-auto max-w-3xl px-6 py-12">
    <h1 class="mb-6 text-3xl font-bold text-slate-900">System health</h1>
    <p v-if="loading" class="text-slate-600">Checking the API and database…</p>
    <p v-else-if="error" class="rounded-lg bg-red-50 p-4 text-red-700" role="alert">{{ error }}</p>
    <div v-else-if="status" class="space-y-3 rounded-lg border border-slate-200 bg-white p-6">
      <p><strong>API:</strong> {{ status.data?.api || 'unknown' }}</p>
      <p><strong>Database:</strong> {{ status.data?.database || 'unknown' }}</p>
      <p class="text-sm text-slate-500">{{ status.message }}</p>
    </div>
    <button class="mt-6 rounded-md bg-emerald-700 px-4 py-2 font-semibold text-white" @click="health.fetchHealth">
      Check again
    </button>
  </section>
</template>
