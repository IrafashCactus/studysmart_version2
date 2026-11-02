<template>
  <div class="min-h-screen bg-slate-950 text-white selection:bg-emerald-500 selection:text-slate-950">
    <!-- Hero Section -->
    <section class="relative isolate flex min-h-[85vh] items-center overflow-hidden border-b border-slate-900">
      <div class="absolute inset-0 bg-gradient-to-b from-emerald-950/20 via-slate-950 to-slate-950 pointer-events-none"></div>

      <div class="relative mx-auto grid w-full max-w-7xl gap-12 px-6 py-20 lg:grid-cols-[1.1fr_0.9fr] lg:items-center">
        <div class="max-w-3xl space-y-6">
          <div class="inline-flex items-center gap-2 rounded-full bg-emerald-500/10 border border-emerald-500/30 px-4 py-1.5 text-xs font-bold text-emerald-300">
            <span>🇷🇼 Rwanda REB Curriculum Grounded</span>
            <span>•</span>
            <span>Adaptive EdTech Ecosystem</span>
          </div>

          <h1 class="text-5xl font-black leading-tight sm:text-6xl lg:text-7xl tracking-tight text-white">
            Micro-learning. <br />
            <span class="bg-gradient-to-r from-emerald-400 via-teal-300 to-cyan-400 bg-clip-text text-transparent">
              Gamified & Adaptive.
            </span>
          </h1>

          <p class="max-w-2xl text-base sm:text-lg leading-relaxed text-slate-300">
            Bite-sized vertical swipe feeds, voice-to-voice AI tutoring, curriculum RAG ingestion, teacher risk radar, and parent accountability monitoring.
          </p>

          <!-- Core Call to Action Buttons -->
          <div class="flex flex-wrap items-center gap-3 pt-2">
            <router-link
              to="/feed"
              class="rounded-xl bg-emerald-500 hover:bg-emerald-400 px-6 py-3.5 text-sm font-black text-slate-950 shadow-xl shadow-emerald-500/25 transition transform active:scale-95 flex items-center gap-2"
            >
              <span>📚</span>
              <span>Launch Mind-Feed</span>
              <span>→</span>
            </router-link>

            <router-link
              to="/gamification"
              class="rounded-xl bg-slate-900 hover:bg-slate-800 border border-slate-800 px-5 py-3.5 text-sm font-bold text-slate-200 transition flex items-center gap-2"
            >
              <span>🔥</span>
              <span>Quests & XP (7d Streak)</span>
            </router-link>

            <router-link
              to="/tutor"
              class="rounded-xl bg-indigo-600/20 hover:bg-indigo-600/30 border border-indigo-500/30 px-5 py-3.5 text-sm font-bold text-indigo-300 transition flex items-center gap-2"
            >
              <span>🤖</span>
              <span>AI Problem Explainer</span>
            </router-link>
          </div>
        </div>

        <!-- Hero Feature Overview Dashboard Card -->
        <div class="rounded-3xl border border-slate-800 bg-slate-900/70 p-6 sm:p-8 shadow-2xl backdrop-blur-md space-y-6">
          <div class="flex items-center justify-between border-b border-slate-800 pb-4">
            <div class="flex items-center gap-2.5">
              <span class="w-3 h-3 rounded-full bg-emerald-400 animate-pulse"></span>
              <span class="text-xs font-bold uppercase tracking-wider text-slate-300">Active Ecosystem Modules</span>
            </div>
            <span class="text-xs font-mono text-emerald-400 font-bold">5 Core Engines Live</span>
          </div>

          <div class="grid grid-cols-1 sm:grid-cols-2 gap-3.5">
            <router-link
              to="/feed"
              class="p-4 rounded-2xl bg-slate-950/60 border border-slate-800 hover:border-emerald-500/40 transition group"
            >
              <span class="text-2xl mb-1 block group-hover:scale-110 transition-transform">📱</span>
              <h4 class="text-sm font-bold text-white group-hover:text-emerald-400">Mind-Feed Stream</h4>
              <p class="text-xs text-slate-400 mt-1">Bite-sized flashcards & inline checks</p>
            </router-link>

            <router-link
              to="/gamification"
              class="p-4 rounded-2xl bg-slate-950/60 border border-slate-800 hover:border-amber-500/40 transition group"
            >
              <span class="text-2xl mb-1 block group-hover:scale-110 transition-transform">🏆</span>
              <h4 class="text-sm font-bold text-white group-hover:text-amber-400">Gamification Hub</h4>
              <p class="text-xs text-slate-400 mt-1">Streaks, Pomodoro timer, Leaderboards</p>
            </router-link>

            <router-link
              to="/educator"
              class="p-4 rounded-2xl bg-slate-950/60 border border-slate-800 hover:border-indigo-500/40 transition group"
            >
              <span class="text-2xl mb-1 block group-hover:scale-110 transition-transform">📊</span>
              <h4 class="text-sm font-bold text-white group-hover:text-indigo-400">Educator Analytics</h4>
              <p class="text-xs text-slate-400 mt-1">Classroom stats & Student Risk Radar</p>
            </router-link>

            <router-link
              to="/admin"
              class="p-4 rounded-2xl bg-slate-950/60 border border-slate-800 hover:border-cyan-500/40 transition group"
            >
              <span class="text-2xl mb-1 block group-hover:scale-110 transition-transform">⚙️</span>
              <h4 class="text-sm font-bold text-white group-hover:text-cyan-400">RAG Workbench</h4>
              <p class="text-xs text-slate-400 mt-1">Ingest textbooks & verify AI quizzes</p>
            </router-link>
          </div>

          <div class="pt-2 border-t border-slate-800/80 flex items-center justify-between text-xs text-slate-400">
            <span>Offline-First Sync: <strong class="text-emerald-400">IndexedDB Ready</strong></span>
            <span>Languages: <strong class="text-white">EN • RW • FR</strong></span>
          </div>
        </div>
      </div>
    </section>

    <!-- Architecture Strip Section -->
    <section class="max-w-7xl mx-auto px-6 py-16 space-y-10">
      <div class="text-center max-w-2xl mx-auto space-y-2">
        <span class="text-xs font-extrabold uppercase tracking-widest text-emerald-400">Platform Capability Matrix</span>
        <h2 class="text-3xl font-black text-white">Full-Stack Scalability & Clean Architecture</h2>
        <p class="text-xs text-slate-400">Built with Go 1.22 (Gin + GORM), Vue 3 Composition API, PostgreSQL, and Docker containerization.</p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 space-y-3">
          <span class="text-3xl block">⚡</span>
          <h3 class="text-base font-bold text-white">Module 1: Student Engine</h3>
          <p class="text-xs text-slate-400 leading-relaxed">
            Eliminates passive learning through bite-sized cards, micro-quizzes, daily flame streaks, Pomodoro focus mode, and step-by-step AI problem explainer.
          </p>
          <router-link to="/feed" class="text-xs font-bold text-emerald-400 hover:underline inline-block pt-2">
            Explore Mind-Feed →
          </router-link>
        </div>

        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 space-y-3">
          <span class="text-3xl block">🔍</span>
          <h3 class="text-base font-bold text-white">Module 2: RAG Pipeline</h3>
          <p class="text-xs text-slate-400 leading-relaxed">
            Automates textbook & exam ingestion, chunk vectorization, LLM quiz generation, and a human-in-the-loop review workbench for teacher verification.
          </p>
          <router-link to="/admin" class="text-xs font-bold text-emerald-400 hover:underline inline-block pt-2">
            Open Verification Workbench →
          </router-link>
        </div>

        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 space-y-3">
          <span class="text-3xl block">📈</span>
          <h3 class="text-base font-bold text-white">Modules 3 & 4: Educator & Parent</h3>
          <p class="text-xs text-slate-400 leading-relaxed">
            High-density classroom analytics, automated student risk radar alerts, homework assignment builder, and weekly parent accountability digests.
          </p>
          <router-link to="/educator" class="text-xs font-bold text-emerald-400 hover:underline inline-block pt-2">
            View Risk Radar →
          </router-link>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useAuthStore } from '../stores/auth'
import { useGamificationStore } from '../stores/gamification'

const authStore = useAuthStore()
const gameStore = useGamificationStore()

onMounted(async () => {
  await authStore.initAuth()
  await gameStore.fetchProfile()
})
</script>
