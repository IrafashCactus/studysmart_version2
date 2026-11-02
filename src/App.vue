  <template>
  <div id="app">
    <Transition name="nav-slide">
      <Navbar v-if="loggedIn && !isAdminRoute" />
    </Transition>

    <div class="app-shell">
      <DashboardLayout v-if="loggedIn && isAdminRoute">
        <RouterView />
      </DashboardLayout>
      <main v-else class="min-w-0 flex-1">
        <RouterView />
      </main>
    </div>
  </div>
</template>

<script setup>
import { computed, onMounted, onUnmounted, ref } from 'vue'
import { RouterView, useRoute } from 'vue-router'
import DashboardLayout from './components/DashboardLayout.vue'
import Navbar from './components/Navbar.vue'
import auth from './services/auth'

const route = useRoute()
const loggedIn = ref(auth.isLoggedIn())
const isAdminRoute = computed(() => ['Dashboard', 'Users', 'Table'].includes(route.name))
const updateAuthStatus = () => {
  loggedIn.value = auth.isLoggedIn()
}

onMounted(() => window.addEventListener('studySmartAuthChange', updateAuthStatus))
onUnmounted(() => window.removeEventListener('studySmartAuthChange', updateAuthStatus))
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
