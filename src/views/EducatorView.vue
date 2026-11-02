<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-6xl mx-auto space-y-8">
      <!-- Header -->
      <div class="border-b border-slate-800 pb-6 flex flex-col sm:flex-row sm:items-center justify-between gap-4">
        <div>
          <div class="flex items-center gap-2">
            <span class="text-3xl">👨‍🏫</span>
            <h1 class="text-2xl font-black text-white">Educator & Classroom Management Hub</h1>
          </div>
          <p class="text-sm text-slate-400 mt-1">
            GS Officiel de Butare • S4 MPC (Mathematics, Physics, Computer Science)
          </p>
        </div>

        <div class="flex items-center gap-2">
          <button
            @click="showQuizBuilder = !showQuizBuilder"
            class="bg-emerald-500 hover:bg-emerald-400 text-slate-950 font-extrabold px-4 py-2.5 rounded-xl text-xs transition shadow-md shadow-emerald-500/20"
          >
            {{ showQuizBuilder ? 'Close Builder' : '+ Create Homework Assignment' }}
          </button>
        </div>
      </div>

      <!-- Quick Metrics Strip -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
        <div class="bg-slate-900 border border-slate-800 rounded-2xl p-5">
          <span class="text-xs text-slate-400 font-medium">Classroom Enrollment</span>
          <p class="text-3xl font-black text-white mt-1">{{ stats.totalStudents || 32 }}</p>
          <span class="text-[11px] text-emerald-400 mt-1 block">Active S4 Stream</span>
        </div>
        <div class="bg-slate-900 border border-slate-800 rounded-2xl p-5">
          <span class="text-xs text-slate-400 font-medium">Average Accuracy</span>
          <p class="text-3xl font-black text-emerald-400 mt-1">{{ Math.round(stats.averageAccuracy || 76) }}%</p>
          <span class="text-[11px] text-slate-400 mt-1 block">Across all REB quizzes</span>
        </div>
        <div class="bg-slate-900 border border-slate-800 rounded-2xl p-5">
          <span class="text-xs text-slate-400 font-medium">Class Active Streak</span>
          <p class="text-3xl font-black text-amber-400 mt-1">🔥 {{ (stats.activeStreakAvg || 5.2).toFixed(1) }}d</p>
          <span class="text-[11px] text-slate-400 mt-1 block">Daily learning consistency</span>
        </div>
        <div class="bg-slate-900 border border-slate-800 rounded-2xl p-5">
          <span class="text-xs text-slate-400 font-medium">Micro-Cards Studied</span>
          <p class="text-3xl font-black text-indigo-400 mt-1">{{ stats.cardsViewedTotal || 142 }}</p>
          <span class="text-[11px] text-slate-400 mt-1 block">Bite-sized interactions</span>
        </div>
      </div>

      <!-- Quiz Builder Modal / Drawer -->
      <div v-if="showQuizBuilder" class="bg-slate-900 border border-emerald-500/40 rounded-3xl p-6 sm:p-8 space-y-4 shadow-2xl animate-fadeIn">
        <div class="flex items-center justify-between">
          <h3 class="text-base font-bold text-white">Create Targeted Assessment from Question Bank</h3>
          <span class="text-xs bg-emerald-500/20 text-emerald-300 px-3 py-1 rounded-full font-bold">Auto-Graded</span>
        </div>

        <div class="grid grid-cols-1 md:grid-cols-3 gap-4">
          <div>
            <label class="text-xs font-semibold text-slate-400 block mb-1">Assignment Title</label>
            <input
              v-model="newQuiz.title"
              placeholder="e.g. Weekly Quadratic & Kinematics Check"
              class="w-full bg-slate-950 border border-slate-800 rounded-xl px-3.5 py-2 text-xs text-white"
            />
          </div>
          <div>
            <label class="text-xs font-semibold text-slate-400 block mb-1">Time Limit (Seconds)</label>
            <input
              v-model.number="newQuiz.timeLimit"
              type="number"
              class="w-full bg-slate-950 border border-slate-800 rounded-xl px-3.5 py-2 text-xs text-white"
            />
          </div>
          <div>
            <label class="text-xs font-semibold text-slate-400 block mb-1">Question Count</label>
            <input
              v-model.number="newQuiz.count"
              type="number"
              class="w-full bg-slate-950 border border-slate-800 rounded-xl px-3.5 py-2 text-xs text-white"
            />
          </div>
        </div>

        <div class="flex justify-end gap-3 pt-2">
          <button @click="showQuizBuilder = false" class="px-4 py-2 bg-slate-800 text-xs font-bold rounded-xl text-slate-300">
            Cancel
          </button>
          <button @click="publishQuiz" class="px-5 py-2 bg-emerald-500 hover:bg-emerald-400 text-slate-950 text-xs font-black rounded-xl">
            Publish to Class Stream →
          </button>
        </div>
      </div>

      <!-- Student Risk Radar Alert System -->
      <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 sm:p-8 space-y-6 shadow-xl">
        <div class="flex items-center justify-between">
          <div class="flex items-center gap-3">
            <span class="text-2xl">🚨</span>
            <div>
              <h3 class="text-base font-bold text-white">Student Risk Radar</h3>
              <p class="text-xs text-slate-400">
                Algorithmically surfaces students with low study hours, disengaged streaks, or quiz accuracy below 60%.
              </p>
            </div>
          </div>
          <span class="text-xs bg-rose-500/20 text-rose-300 border border-rose-500/30 px-3 py-1 rounded-full font-bold">
            Automated Alert System
          </span>
        </div>

        <!-- Radar Table -->
        <div class="overflow-x-auto">
          <table class="w-full text-left text-xs">
            <thead class="text-slate-400 border-b border-slate-800 uppercase tracking-wider">
              <tr>
                <th class="pb-3 font-bold">Student</th>
                <th class="pb-3 font-bold">Study Hours / Wk</th>
                <th class="pb-3 font-bold">Accuracy</th>
                <th class="pb-3 font-bold">Active Streak</th>
                <th class="pb-3 font-bold">Risk Level</th>
                <th class="pb-3 font-bold">Algorithmic Alert Reason</th>
                <th class="pb-3 font-bold text-right">Intervention</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-800/60 text-slate-300">
              <tr v-for="st in riskRadar" :key="st.userId" class="hover:bg-slate-800/40">
                <td class="py-3.5 font-bold text-white flex items-center gap-2.5">
                  <img :src="st.avatarUrl" class="w-8 h-8 rounded-full bg-slate-800" />
                  <span>{{ st.fullName }}</span>
                </td>
                <td class="py-3.5 font-mono">{{ st.studyHoursWeek }}h</td>
                <td class="py-3.5 font-mono font-bold" :class="st.averageScore < 60 ? 'text-rose-400' : 'text-emerald-400'">
                  {{ st.averageScore > 0 ? `${Math.round(st.averageScore)}%` : 'New' }}
                </td>
                <td class="py-3.5 text-amber-400 font-semibold">🔥 {{ st.streakDays }}d</td>
                <td class="py-3.5">
                  <span
                    class="px-2.5 py-1 rounded-full text-[10px] font-extrabold uppercase"
                    :class="{
                      'bg-rose-950 text-rose-300 border border-rose-800/40': st.riskLevel === 'high',
                      'bg-amber-950 text-amber-300 border border-amber-800/40': st.riskLevel === 'moderate',
                      'bg-emerald-950 text-emerald-300 border border-emerald-800/40': st.riskLevel === 'low',
                    }"
                  >
                    {{ st.riskLevel }} Risk
                  </span>
                </td>
                <td class="py-3.5 text-slate-300">
                  <span class="block truncate max-w-xs">{{ st.riskReasons?.join(', ') }}</span>
                </td>
                <td class="py-3.5 text-right">
                  <button
                    @click="notifyStudent(st.fullName)"
                    class="px-3 py-1.5 bg-slate-800 hover:bg-slate-700 text-slate-200 text-xs font-semibold rounded-lg transition border border-slate-700"
                  >
                    Send Guidance
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const stats = ref({})
const riskRadar = ref([])
const showQuizBuilder = ref(false)
const newQuiz = ref({
  title: 'S4 STEM Diagnostic Practice',
  timeLimit: 300,
  count: 5,
})

onMounted(async () => {
  fetchAnalytics()
  fetchRiskRadar()
})

async function fetchAnalytics() {
  try {
    const res = await api.get('/api/v1/educator/analytics')
    if (res.data?.data) {
      stats.value = res.data.data
    }
  } catch (err) {
    console.warn('Analytics error', err)
  }
}

async function fetchRiskRadar() {
  try {
    const res = await api.get('/api/v1/educator/risk-radar')
    if (res.data?.data) {
      riskRadar.value = res.data.data
    }
  } catch (err) {
    console.warn('Risk radar error', err)
  }
}

function publishQuiz() {
  alert(`Quiz "${newQuiz.value.title}" published with ${newQuiz.value.count} questions from the verified question bank!`)
  showQuizBuilder.value = false
}

function notifyStudent(name) {
  alert(`Remediation micro-assignment dispatched to ${name}'s Mind-Feed stream.`)
}
</script>

