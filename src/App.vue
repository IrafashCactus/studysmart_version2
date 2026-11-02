<template>
  <div id="app" class="bg-slate-950 text-slate-100 min-h-screen flex flex-col">
    <Navbar />
    <main class="min-w-0 flex-1">
      <RouterView />
    </main>
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
</script>

<style>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
}

body {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  background: #020617;
  color: #f8fafc;
}

#app {
  min-height: 100vh;
}
</style>
