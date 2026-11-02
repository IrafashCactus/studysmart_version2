<template>
  <div class="min-h-screen bg-gradient-to-br from-blue-950 to-blue-900 py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <div class="bg-white rounded-xl shadow-2xl p-6">
        <h1 class="text-3xl font-bold text-blue-900 mb-6">📊 Users Table</h1>

        <div v-if="loading" class="text-center py-12">Loading data...</div>
        <div v-else-if="error" class="text-center py-12 text-red-600">{{ error }}</div>

        <div v-else class="overflow-x-auto">
          <DataTable :value="users" responsiveLayout="scroll" class="min-w-full">
            <Column field="id" header="ID" style="min-width: 5rem" />
            <Column field="names" header="Name" style="min-width: 12rem" />

            <Column header="Avatar" style="min-width: 12rem">
              <template #body="{ data }">
                <div class="flex items-center gap-3">
                  <Avatar
                    :image="data.avatar"
                    :label="data.names ? data.names.charAt(0) : ''"
                    shape="circle"
                    size="large"
                  />
                </div>
              </template>
            </Column>

            <Column field="email" header="Email" style="min-width: 12rem" />
            <Column field="age" header="Age" style="min-width: 6rem" />
          </DataTable>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import api from '../services/api'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import Avatar from 'primevue/avatar'

export default {
  name: 'TablePage',
  components: {
    DataTable,
    Column,
    Avatar
  },
  data() {
    return {
      users: [],
      loading: false,
      error: null
    }
  },
  mounted() {
    this.fetchUsers()
  },
  methods: {
    async fetchUsers() {
      this.loading = true
      this.error = null

      try {
        const response = await api.get('/users')
        const payload = response.data
        this.users = payload.users || payload.data || payload || []
      } catch (err) {
        this.error = err.response?.data?.message || err.message || 'Unable to load users.'
      } finally {
        this.loading = false
      }
    }
  }
}
</script>

