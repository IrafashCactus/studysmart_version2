<template>
  <div class="min-h-screen bg-gradient-to-br from-emerald-950 to-emerald-900 flex items-center justify-center py-12 px-4">
    <div class="max-w-md w-full bg-white rounded-xl shadow-2xl p-8">
      <div class="text-center mb-8">
        <h2 class="text-3xl font-bold text-emerald-900">Create Account</h2>
        <p class="text-emerald-600 mt-2">Join StudySmart today</p>
      </div>
      
      <form @submit.prevent="handleSignup">
        <div class="mb-4">
          <label for="signup-name" class="block text-emerald-900 font-medium mb-2">Full Name</label>
          <input id="signup-name" name="name" v-model="name" type="text" autocomplete="name" required
            class="w-full px-4 py-2 border border-emerald-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500" />
        </div>
        
        <div class="mb-4">
          <label for="signup-email" class="block text-emerald-900 font-medium mb-2">Email</label>
          <input id="signup-email" name="email" v-model="email" type="email" autocomplete="email" required
            class="w-full px-4 py-2 border border-emerald-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500" />
        </div>
        
        <div class="mb-4">
          <label for="signup-password" class="block text-emerald-900 font-medium mb-2">Password</label>
          <input id="signup-password" name="password" v-model="password" type="password" autocomplete="new-password" required
            class="w-full px-4 py-2 border border-emerald-200 rounded-lg focus:outline-none focus:ring-2 focus:ring-emerald-500" />
        </div>
        
        <button type="submit" :disabled="loading"
          class="w-full bg-emerald-500 hover:bg-emerald-600 text-white font-semibold py-2 rounded-lg transition">
          {{ loading ? 'Creating Account...' : 'Sign Up' }}
        </button>
        
        <p class="text-center mt-4 text-emerald-600">
          Already have an account? 
          <router-link to="/login" class="text-emerald-700 font-semibold hover:underline">Login</router-link>
        </p>
      </form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'SignupPage',
  data() {
    return {
      name: '',
      email: '',
      password: '',
      loading: false
    }
  },
  methods: {
    async handleSignup() {
      this.loading = true
      setTimeout(() => {
        this.$router.push('/login')
        this.loading = false
      }, 1000)
    }
  }
}
</script>