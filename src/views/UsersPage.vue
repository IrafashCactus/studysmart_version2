<template>
  <div class="min-h-screen bg-gradient-to-br from-green-950 to-green-900 py-8 px-4">
    <div class="max-w-7xl mx-auto">
      <div class="text-center mb-8">
        <h1 class="text-4xl font-bold text-white mb-2">User Management System</h1>
        <p class="text-green-200">Manage your users with full CRUD operations</p>
      </div>
      
      <!-- Create Form -->
      <div class="bg-white rounded-xl shadow-2xl p-6 mb-8">
        <h2 class="text-xl font-bold text-green-950 mb-4">➕ Add New User</h2>
        
        <form @submit.prevent="createUser" class="space-y-4">
          <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
            <input id="newuser-name" name="name" v-model="newUser.name" type="text" placeholder="Full Name" required
              class="w-full px-4 py-2 border border-green-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500" />
            <input id="newuser-email" name="email" v-model="newUser.email" type="email" placeholder="Email Address" required
              class="w-full px-4 py-2 border border-green-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500" />
            <input id="newuser-age" name="age" v-model.number="newUser.age" type="number" placeholder="Age" required
              class="w-full px-4 py-2 border border-green-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500" />
          </div>
          <button type="submit" :disabled="isCreating"
            class="bg-green-500 hover:bg-green-600 text-white font-semibold py-2 px-6 rounded-lg transition">
            {{ isCreating ? 'Creating...' : '✨ Create User' }}
          </button>
        </form>
      </div>
      
      <!-- User List Section -->
      <div class="bg-white rounded-xl shadow-2xl p-6">
        <div class="flex justify-between items-center mb-6">
          <h2 class="text-2xl font-bold text-green-950">👥 User List ({{ filteredUsers.length }})</h2>
          <label for="search-users" class="sr-only">Search users</label>
          <input id="search-users" name="search" v-model="searchQuery" type="text" placeholder="🔍 Search users..."
            class="px-4 py-2 border border-green-200 rounded-lg w-64" aria-label="Search users" />
        </div>
        <!-- Loading/Error/Empty States -->
        <div v-if="loading" class="text-center py-12">Loading users...</div>
        <div v-else-if="error" class="text-center py-12 text-red-600">{{ error }}</div>
        <div v-else-if="filteredUsers.length === 0" class="text-center py-12 text-green-700">No users found</div>
        
        <!-- User Grid -->
        <div v-else class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          <div v-for="user in filteredUsers" :key="user.id" class="bg-gradient-to-br from-green-50 to-white rounded-lg border border-green-200 overflow-hidden hover:shadow-xl transition">
            <div class="bg-green-950 px-4 py-3">
              <h3 class="text-white font-bold">{{ user.name }}</h3>
            </div>
            <div class="p-4">
              <p>📧 {{ user.email }}</p>
              <p>🎂 Age: {{ user.age }}</p>
              <p>🆔 ID: {{ user.id }}</p>
            </div>
            <div class="flex border-t border-green-200">
              <button @click="openEditModal(user)" class="flex-1 py-2 bg-green-950 hover:bg-green-900 text-white">✏️ Edit</button>
              <button @click="confirmDelete(user)" class="flex-1 py-2 bg-white hover:bg-gray-200 text-black">🗑️ Delete</button>
            </div>
          </div>
        </div>
      </div>
    </div>
    
    <!-- Edit Modal -->
    <div v-if="showEditModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click.self="closeEditModal">
      <div class="bg-white rounded-xl max-w-md w-full p-6">
        <h2 class="text-xl font-bold mb-4">✏️ Edit User</h2>
        <label for="edit-name" class="sr-only">Edit name</label>
        <input id="edit-name" name="name" v-model="editForm.name" placeholder="Name" class="w-full px-4 py-2 border rounded-lg mb-3" />
        <label for="edit-email" class="sr-only">Edit email</label>
        <input id="edit-email" name="email" v-model="editForm.email" placeholder="Email" class="w-full px-4 py-2 border rounded-lg mb-3" />
        <label for="edit-age" class="sr-only">Edit age</label>
        <input id="edit-age" name="age" v-model.number="editForm.age" placeholder="Age" class="w-full px-4 py-2 border rounded-lg mb-3" />
        <div class="flex gap-3">
          <button @click="updateUser" class="flex-1 bg-green-500 text-white py-2 rounded-lg">Save</button>
          <button @click="closeEditModal" class="flex-1 bg-gray-200 py-2 rounded-lg">Cancel</button>
        </div>
      </div>
    </div>
    
    <!-- Delete Modal -->
    <div v-if="showDeleteModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50" @click.self="closeDeleteModal">
      <div class="bg-white rounded-xl max-w-md w-full p-6">
        <h2 class="text-xl font-bold text-red-600 mb-4">⚠️ Confirm Delete</h2>
        <p>Delete {{ userToDelete?.name }}?</p>
        <div class="flex gap-3 mt-4">
          <button @click="deleteUser" class="flex-1 bg-red-500 text-white py-2 rounded-lg">Delete</button>
          <button @click="closeDeleteModal" class="flex-1 bg-gray-200 py-2 rounded-lg">Cancel</button>
        </div>
      </div>
    </div>
    
    <!-- Toast -->
    <div v-if="successMessage" class="fixed bottom-4 right-4 bg-green-500 text-white px-6 py-3 rounded-lg">
      ✅ {{ successMessage }}
    </div>
  </div>
</template>

<script>
import api from '../services/api'

export default {
  name: 'UsersPage',
  data() {
    return {
      users: [],
      loading: false,
      error: null,
      searchQuery: '',
      newUser: { name: '', email: '', age: null },
      isCreating: false,
      showEditModal: false,
      editForm: { id: null, name: '', email: '', age: null },
      showDeleteModal: false,
      userToDelete: null,
      successMessage: ''
    }
  },
  computed: {
    filteredUsers() {
      if (!this.searchQuery) return this.users
      const query = this.searchQuery.toLowerCase()
      return this.users.filter(u => u.name.toLowerCase().includes(query) || u.email.toLowerCase().includes(query))
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
        const payload = response.data || {}
        let users = []

        if (Array.isArray(payload)) {
          users = payload
        } else if (Array.isArray(payload.users)) {
          users = payload.users
        } else if (Array.isArray(payload.data)) {
          users = payload.data
        }

        this.users = users.map((user) => ({
          ...user,
          name: user.name || user.names || '',
          email: user.email || '',
          age: user.age || null
        }))
      } catch (err) {
        this.error = err.response?.data?.message || err.message || 'Unable to load users from API.'
      } finally {
        this.loading = false
      }
    },
    async createUser() {
      this.isCreating = true
      const newUser = { id: Date.now(), ...this.newUser }
      this.users.unshift(newUser)
      this.saveToLocalStorage()
      this.newUser = { name: '', email: '', age: null }
      this.showSuccess('User created!')
      this.isCreating = false
    },
    openEditModal(user) {
      this.editForm = { ...user }
      this.showEditModal = true
    },
    async updateUser() {
      const index = this.users.findIndex(u => u.id === this.editForm.id)
      if (index !== -1) {
        this.users[index] = { ...this.editForm }
        this.saveToLocalStorage()
        this.showSuccess('User updated!')
      }
      this.closeEditModal()
    },
    confirmDelete(user) {
      this.userToDelete = user
      this.showDeleteModal = true
    },
    async deleteUser() {
      this.users = this.users.filter(u => u.id !== this.userToDelete.id)
      this.saveToLocalStorage()
      this.showSuccess('User deleted!')
      this.closeDeleteModal()
    },
    closeEditModal() { this.showEditModal = false },
    closeDeleteModal() { this.showDeleteModal = false },
    saveToLocalStorage() { localStorage.setItem('users', JSON.stringify(this.users)) },
    showSuccess(msg) {
      this.successMessage = msg
      setTimeout(() => { this.successMessage = '' }, 3000)
    }
  }
}
</script>