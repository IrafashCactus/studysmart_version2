<template>
  <nav class="sticky top-0 z-40 border-b border-slate-800 bg-slate-950/95 backdrop-blur-md shadow-lg transition-all text-slate-100">
    <div class="mx-auto flex h-16 max-w-7xl items-center justify-between px-4 sm:px-6 lg:px-8">
      <!-- Brand Logo -->
      <router-link to="/" class="flex items-center gap-3">
        <span class="grid h-10 w-10 place-items-center rounded-xl bg-gradient-to-tr from-emerald-600 to-teal-400 shadow-lg shadow-emerald-500/20 text-lg font-black text-slate-950">
          S
        </span>
        <div class="flex flex-col">
          <span class="text-base font-black tracking-tight text-white flex items-center gap-1.5">
            <span>{{ t('appTitle') }}</span>
            <span class="text-[10px] bg-emerald-500/20 text-emerald-400 border border-emerald-500/30 px-1.5 py-0.2 rounded font-extrabold uppercase">REB</span>
          </span>
          <span class="text-[10px] text-slate-400 font-medium">{{ t('appSubtitle') }}</span>
        </div>
      </router-link>

      <!-- Center Navigation Links -->
      <div class="hidden lg:flex items-center gap-1">
        <router-link
          to="/feed"
          active-class="bg-slate-800 text-emerald-400 font-bold border-b-2 border-emerald-500"
          class="rounded-xl px-3 py-2 text-xs font-semibold text-slate-300 transition hover:bg-slate-800/60 hover:text-white flex items-center gap-1.5"
        >
          <span>📚</span>
          <span>Mind-Feed</span>
        </router-link>

        <router-link
          to="/gamification"
          active-class="bg-slate-800 text-emerald-400 font-bold border-b-2 border-emerald-500"
          class="rounded-xl px-3 py-2 text-xs font-semibold text-slate-300 transition hover:bg-slate-800/60 hover:text-white flex items-center gap-1.5"
        >
          <span>🏆</span>
          <span>{{ t('dailyQuests') }}</span>
        </router-link>

        <router-link
          to="/tutor"
          active-class="bg-slate-800 text-emerald-400 font-bold border-b-2 border-emerald-500"
          class="rounded-xl px-3 py-2 text-xs font-semibold text-slate-300 transition hover:bg-slate-800/60 hover:text-white flex items-center gap-1.5"
        >
          <span>🤖</span>
          <span>{{ t('aiTutor') }}</span>
        </router-link>

        <router-link
          to="/educator"
          active-class="bg-slate-800 text-emerald-400 font-bold border-b-2 border-emerald-500"
          class="rounded-xl px-3 py-2 text-xs font-semibold text-slate-300 transition hover:bg-slate-800/60 hover:text-white flex items-center gap-1.5"
        >
          <span>👨‍🏫</span>
          <span>{{ t('educatorHub') }}</span>
        </router-link>

        <router-link
          to="/parent"
          active-class="bg-slate-800 text-emerald-400 font-bold border-b-2 border-emerald-500"
          class="rounded-xl px-3 py-2 text-xs font-semibold text-slate-300 transition hover:bg-slate-800/60 hover:text-white flex items-center gap-1.5"
        >
          <span>👨‍👩‍👧</span>
          <span>{{ t('parentDigest') }}</span>
        </router-link>

        <router-link
          to="/admin"
          active-class="bg-slate-800 text-emerald-400 font-bold border-b-2 border-emerald-500"
          class="rounded-xl px-3 py-2 text-xs font-semibold text-slate-300 transition hover:bg-slate-800/60 hover:text-white flex items-center gap-1.5"
        >
          <span>⚙️</span>
          <span>{{ t('adminWorkbench') }}</span>
        </router-link>

        <router-link
          to="/health"
          active-class="bg-slate-800 text-emerald-400 font-bold"
          class="rounded-xl px-2.5 py-2 text-xs font-semibold text-slate-400 transition hover:bg-slate-800/60 hover:text-white"
        >
          Health
        </router-link>
      </div>

      <!-- Right Hand Controls (Network Sync Status, Streak, Level, Role Switcher, Lang) -->
      <div class="flex items-center gap-2.5">
        <!-- Offline/Online Status Indicator Badge -->
        <span
          class="hidden sm:inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-[10px] font-bold border"
          :class="isOnline
            ? 'bg-emerald-950/60 text-emerald-400 border-emerald-800/40'
            : 'bg-amber-950/60 text-amber-400 border-amber-800/40'"
          :title="isOnline ? 'Connected to live cloud API' : 'IndexedDB offline storage active'"
        >
          <span class="w-1.5 h-1.5 rounded-full" :class="isOnline ? 'bg-emerald-400 animate-pulse' : 'bg-amber-400'"></span>
          <span>{{ isOnline ? 'Online' : 'Offline Mode' }}</span>
        </span>

        <!-- Streak Flame Counter -->
        <router-link
          to="/gamification"
          class="flex items-center gap-1.5 bg-amber-500/10 border border-amber-500/30 px-3 py-1.5 rounded-full text-xs font-black text-amber-400 hover:bg-amber-500/20 transition shadow-sm"
          title="Daily Learning Streak"
        >
          <span class="animate-bounce">🔥</span>
          <span>{{ gameStore.streak.currentStreak }}d</span>
        </router-link>

        <!-- XP Level Pill -->
        <router-link
          to="/gamification"
          class="hidden sm:flex items-center gap-1.5 bg-emerald-500/10 border border-emerald-500/30 px-3 py-1.5 rounded-full text-xs font-bold text-emerald-400 hover:bg-emerald-500/20 transition shadow-sm"
          title="Current Level"
        >
          <span>⚡</span>
          <span>Lvl {{ gameStore.xp.currentLevel }}</span>
        </router-link>

        <!-- Quick Role Switcher -->
        <div class="relative">
          <select
            v-model="selectedRole"
            @change="switchRole"
            class="bg-slate-900 border border-slate-700 rounded-xl px-2 py-1.5 text-xs font-bold text-slate-200 focus:outline-none focus:border-emerald-500 cursor-pointer"
            title="Fast switch between multi-role demo accounts"
          >
            <option value="student">🎓 Keza (Student)</option>
            <option value="teacher">👨‍🏫 Jean (Teacher)</option>
            <option value="parent">👨‍👩‍👧 David (Parent)</option>
            <option value="admin">🛡️ Admin</option>
          </select>
        </div>

        <!-- Language Selector -->
        <select
          v-model="selectedLang"
          @change="changeLang"
          class="bg-slate-900 border border-slate-800 rounded-xl px-2 py-1.5 text-xs font-semibold text-slate-300 focus:outline-none"
        >
          <option value="en">🇬🇧 EN</option>
          <option value="rw">🇷🇼 RW</option>
          <option value="fr">🇫🇷 FR</option>
        </select>
      </div>
    </div>
  </nav>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useGamificationStore } from '../stores/gamification'
import { useI18n } from '../locales/i18n'

const router = useRouter()
const authStore = useAuthStore()
const gameStore = useGamificationStore()
const { t } = useI18n()

const selectedRole = ref(authStore.user?.role || 'student')
const selectedLang = ref(localStorage.getItem('studySmartLang') || 'en')
const isOnline = ref(navigator.onLine)

onMounted(async () => {
  await authStore.initAuth()
  await gameStore.fetchProfile()

  window.addEventListener('online', () => { isOnline.value = true })
  window.addEventListener('offline', () => { isOnline.value = false })
})

async function switchRole() {
  await authStore.demoLogin(selectedRole.value)
  await gameStore.fetchProfile()

  switch (selectedRole.value) {
    case 'student':
      router.push('/feed')
      break
    case 'teacher':
      router.push('/educator')
      break
    case 'parent':
      router.push('/parent')
      break
    case 'admin':
      router.push('/admin')
      break
  }
}

function changeLang() {
  authStore.setLanguage(selectedLang.value)
  window.location.reload()
}
</script>
