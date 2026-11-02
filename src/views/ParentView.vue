<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-4xl mx-auto space-y-8">
      <!-- Header -->
      <div class="border-b border-slate-800 pb-6 flex items-center justify-between">
        <div>
          <div class="flex items-center gap-2">
            <span class="text-3xl">👨‍👩‍👧</span>
            <h1 class="text-2xl font-black text-white">Parent Accountability & Progress Digest</h1>
          </div>
          <p class="text-sm text-slate-400 mt-1">
            Monitoring learner autonomy and mastery progress for: <strong class="text-emerald-400">{{ digest.studentName || 'Keza Aline' }}</strong>
          </p>
        </div>
        <span class="text-xs bg-slate-900 border border-slate-800 text-slate-400 px-3.5 py-1.5 rounded-full font-mono">
          Weekly Snapshot
        </span>
      </div>

      <!-- Quick Metrics Summary -->
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-6">
        <!-- Streak Snapshot -->
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 text-center space-y-2">
          <span class="text-4xl block animate-bounce">🔥</span>
          <span class="text-3xl font-black text-amber-400">{{ digest.currentStreak || 7 }} Days</span>
          <p class="text-xs text-slate-400">Active Learning Streak</p>
        </div>

        <!-- Weekly Study Time -->
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 text-center space-y-2">
          <span class="text-4xl block">⏱️</span>
          <span class="text-3xl font-black text-indigo-400">{{ digest.totalStudyHours || 6.5 }} Hours</span>
          <p class="text-xs text-slate-400">Active Study Block Time</p>
        </div>

        <!-- Overall Concept Mastery -->
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 text-center space-y-2">
          <span class="text-4xl block">🎯</span>
          <span class="text-3xl font-black text-emerald-400">{{ digest.overallMasteryPct || 84.2 }}%</span>
          <p class="text-xs text-slate-400">Curriculum Concept Mastery</p>
        </div>
      </div>

      <!-- Productivity vs Screen-Time Ratio Gauge -->
      <div class="bg-gradient-to-br from-slate-900 to-slate-900/90 border border-slate-800 rounded-3xl p-6 sm:p-8 space-y-6 shadow-xl">
        <div class="flex items-center justify-between">
          <div>
            <h3 class="text-base font-bold text-white">Screen-Time Productivity Ratio</h3>
            <p class="text-xs text-slate-400 mt-0.5">Measures high-engagement micro-quizzes & flashcards vs passive idle browsing.</p>
          </div>
          <span class="text-lg font-black font-mono text-emerald-400">
            {{ digest.screenTimeProductivityPct || 91 }}% Productive
          </span>
        </div>

        <!-- Visual Stacked Bar -->
        <div class="space-y-2">
          <div class="w-full h-4 bg-slate-800 rounded-full overflow-hidden flex p-0.5 border border-slate-700">
            <div
              class="h-full bg-emerald-500 rounded-l-full transition-all duration-500"
              :style="{ width: `${digest.screenTimeProductivityPct || 91}%` }"
              title="Active Micro-Learning"
            ></div>
            <div
              class="h-full bg-slate-600 rounded-r-full transition-all duration-500"
              :style="{ width: `${100 - (digest.screenTimeProductivityPct || 91)}%` }"
              title="Passive Reading / Navigation"
            ></div>
          </div>

          <div class="flex justify-between text-xs text-slate-400 pt-1">
            <span class="flex items-center gap-1.5 text-emerald-400 font-semibold">
              <span class="w-2.5 h-2.5 rounded-full bg-emerald-500 inline-block"></span>
              Active Learning: {{ digest.screenTimeProductivityPct || 91 }}%
            </span>
            <span class="flex items-center gap-1.5 text-slate-400 font-semibold">
              <span class="w-2.5 h-2.5 rounded-full bg-slate-600 inline-block"></span>
              Passive / Navigation: {{ 100 - (digest.screenTimeProductivityPct || 91) }}%
            </span>
          </div>
        </div>
      </div>

      <!-- Strengths & Weakness Matrix -->
      <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
        <!-- Strengths -->
        <div class="bg-slate-900 border border-emerald-500/30 rounded-3xl p-6 space-y-4">
          <div class="flex items-center gap-2">
            <span class="text-emerald-400 text-lg">💪</span>
            <h3 class="text-sm font-bold text-white uppercase tracking-wider">Top Concept Strengths</h3>
          </div>
          <ul class="space-y-2.5 text-xs">
            <li
              v-for="s in (digest.strongSubjects || ['Mathematics (88.5%)', 'Entrepreneurship (92.0%)'])"
              :key="s"
              class="bg-emerald-950/40 border border-emerald-800/40 text-emerald-200 p-3 rounded-xl flex items-center justify-between"
            >
              <span>{{ s }}</span>
              <span class="text-emerald-400 font-bold">✓ Mastered</span>
            </li>
          </ul>
        </div>

        <!-- Areas for Reinforcement -->
        <div class="bg-slate-900 border border-amber-500/30 rounded-3xl p-6 space-y-4">
          <div class="flex items-center gap-2">
            <span class="text-amber-400 text-lg">🎯</span>
            <h3 class="text-sm font-bold text-white uppercase tracking-wider">Areas Requiring Reinforcement</h3>
          </div>
          <ul class="space-y-2.5 text-xs">
            <li
              v-for="w in (digest.needsReinforcement || ['Physics: Kinematics (64.0%)'])"
              :key="w"
              class="bg-amber-950/40 border border-amber-500/40 text-amber-200 p-3 rounded-xl flex items-center justify-between"
            >
              <span>{{ w }}</span>
              <span class="text-amber-400 font-bold">Reinforcement Scheduled</span>
            </li>
          </ul>
        </div>
      </div>

      <!-- Teacher Feedback Summary -->
      <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 space-y-2">
        <span class="text-xs uppercase font-bold text-slate-400 tracking-wider">Educator Guidance Summary</span>
        <p class="text-sm text-slate-200 italic leading-relaxed">
          "{{ digest.teacherNotes || 'Keza is performing well in Algebra. Extra review recommended on projectile vectors.' }}"
        </p>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const digest = ref({})

onMounted(async () => {
  try {
    const res = await api.get('/api/v1/parent/digest')
    if (res.data?.data) {
      digest.value = res.data.data
    }
  } catch (err) {
    console.warn('Digest load failed', err)
  }
})
</script>

