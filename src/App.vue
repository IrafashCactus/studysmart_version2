  <template>
  <div id="app">
    <Transition name="nav-slide">
      <NavBar v-if="loggedIn && !isAdminRoute" />
    </Transition>

    <div class="app-shell">
      <DashboardLayout v-if="loggedIn && isAdminRoute">
        <router-view />
      </DashboardLayout>

      <main v-else class="min-w-0 flex-1">
        <router-view />
      </main>
    </div>
  </div>
</template>

<script>
import NavBar from './components/NavBar.vue'
import DashboardLayout from './components/DashboardLayout.vue'
import auth from './services/auth'

export default {
  name: 'App',
  components: {
    NavBar,
    DashboardLayout,
  },
  data() {
    return {
      loggedIn: auth.isLoggedIn(),
    }
  },
  computed: {
    isAdminRoute() {
      return ['Dashboard', 'Users', 'Table'].includes(this.$route.name)
    },
  },
  mounted() {
    window.addEventListener('studySmartAuthChange', this.updateAuthStatus)
  },
  beforeUnmount() {
    window.removeEventListener('studySmartAuthChange', this.updateAuthStatus)
  },
  methods: {
    updateAuthStatus() {
      this.loggedIn = auth.isLoggedIn()
    }
  }
}
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: #f8fafc;
}

#app {
  min-height: 100vh;
}

.app-shell {
  min-height: 100vh;
}
</style>
