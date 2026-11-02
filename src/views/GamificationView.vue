<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-6xl mx-auto space-y-8">
      <!-- Top Banner: Streak, Level, & Focus Mode Sprint -->
      <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <!-- Streak Card -->
        <div class="bg-gradient-to-br from-amber-500/20 via-slate-900 to-slate-900 border border-amber-500/30 rounded-3xl p-6 shadow-xl relative overflow-hidden flex flex-col justify-between">
          <div class="absolute -right-6 -bottom-6 text-8xl opacity-10 select-none">🔥</div>
          <div>
            <div class="flex items-center justify-between mb-2">
              <span class="text-xs uppercase font-extrabold text-amber-400 tracking-wider">Daily Streak</span>
              <span class="text-xs bg-amber-500/20 text-amber-300 font-bold px-2.5 py-1 rounded-full border border-amber-500/30">
                Best: {{ gameStore.streak.longestStreak }} Days
              </span>
            </div>
            <div class="flex items-baseline gap-3 my-2">
              <span class="text-5xl font-black text-amber-400 tracking-tight flex items-center gap-1">
                <span class="animate-bounce">🔥</span>
                <span>{{ gameStore.streak.currentStreak }}</span>
              </span>
              <span class="text-slate-400 font-medium text-sm">Consecutive Days</span>
            </div>
            <p class="text-xs text-slate-400 mt-1">Study at least 1 card daily to maintain your active flame!</p>
          </div>

          <!-- Streak Freeze Pass -->
          <div class="mt-4 pt-4 border-t border-slate-800/80 flex items-center justify-between">
            <div class="flex items-center gap-1.5 text-xs text-slate-300">
              <span>❄️ Freeze Passes:</span>
              <strong class="text-cyan-400">{{ gameStore.streak.freezePasses }} Left</strong>
            </div>
            <button
              @click="useFreeze"
              :disabled="gameStore.streak.freezePasses <= 0"
              class="px-3 py-1.5 bg-cyan-500/20 text-cyan-300 hover:bg-cyan-500/30 text-xs font-bold rounded-lg transition border border-cyan-500/30 disabled:opacity-40"
            >
              Protect Streak
            </button>
          </div>
        </div>

        <!-- XP & Level Card -->
        <div class="bg-gradient-to-br from-emerald-500/20 via-slate-900 to-slate-900 border border-emerald-500/30 rounded-3xl p-6 shadow-xl relative overflow-hidden flex flex-col justify-between">
          <div class="absolute -right-6 -bottom-6 text-8xl opacity-10 select-none">⚡</div>
          <div>
            <div class="flex items-center justify-between mb-2">
              <span class="text-xs uppercase font-extrabold text-emerald-400 tracking-wider">Level Progression</span>
              <span class="text-xs bg-emerald-500/20 text-emerald-300 font-bold px-2.5 py-1 rounded-full border border-emerald-500/30">
                Level {{ gameStore.xp.currentLevel }}
              </span>
            </div>
            <div class="flex items-baseline gap-2 my-2">
              <span class="text-4xl font-black text-white tracking-tight">
                {{ gameStore.xp.totalXp.toLocaleString() }}
              </span>
              <span class="text-xs font-bold uppercase text-emerald-400">Total XP</span>
            </div>

            <!-- Progress Bar -->
            <div class="space-y-1.5 mt-3">
              <div class="flex justify-between text-xs text-slate-400 font-mono">
                <span>Progress</span>
                <span>{{ levelProgressPct }}%</span>
              </div>
              <div class="w-full h-3 bg-slate-800 rounded-full overflow-hidden p-0.5 border border-slate-700">
                <div
                  class="h-full bg-gradient-to-r from-emerald-500 to-cyan-400 rounded-full transition-all duration-500"
                  :style="{ width: `${levelProgressPct}%` }"
                ></div>
              </div>
              <p class="text-[11px] text-slate-400 text-right">{{ gameStore.xp.xpToNext }} XP to Level {{ gameStore.xp.currentLevel + 1 }}</p>
            </div>
          </div>

          <!-- Badges mini row -->
          <div class="mt-4 pt-3 border-t border-slate-800/80 flex items-center gap-2 overflow-x-auto">
            <span v-for="b in gameStore.badges" :key="b" class="text-[11px] bg-slate-800 px-2 py-0.5 rounded-md text-slate-300 shrink-0 border border-slate-700">
              🏅 {{ b }}
            </span>
          </div>
        </div>

        <!-- Pomodoro Focus Mode Timer Card -->
        <div class="bg-gradient-to-br from-indigo-500/20 via-slate-900 to-slate-900 border border-indigo-500/30 rounded-3xl p-6 shadow-xl relative overflow-hidden flex flex-col justify-between">
          <div class="absolute -right-6 -bottom-6 text-8xl opacity-10 select-none">⏱️</div>
          <div>
            <div class="flex items-center justify-between mb-2">
              <span class="text-xs uppercase font-extrabold text-indigo-400 tracking-wider">Pomodoro Focus Timer</span>
              <span class="text-xs font-bold text-amber-400 flex items-center gap-1">
                <span>+15 XP</span>
              </span>
            </div>
            <div class="text-center my-3">
              <span class="text-5xl font-mono font-black text-indigo-300 tracking-wider">
                {{ formatTimer(gameStore.focusTimeRemaining) }}
              </span>
              <p class="text-xs text-slate-400 mt-1">
                {{ gameStore.isFocusing ? 'Focus Mode Active • Notifications blocked' : '25-Minute Deep Study Block' }}
              </p>
            </div>
          </div>

          <!-- Timer Controls -->
          <div class="flex items-center gap-2 mt-4 pt-3 border-t border-slate-800/80">
            <button
              v-if="!gameStore.isFocusing"
              @click="gameStore.startFocusSession"
              class="flex-1 bg-indigo-600 hover:bg-indigo-500 text-white font-bold py-2 rounded-xl text-xs transition shadow-md shadow-indigo-600/30"
            >
              Start Focus Session
            </button>
            <button
              v-else
              @click="gameStore.pauseFocusSession"
              class="flex-1 bg-amber-600 hover:bg-amber-500 text-white font-bold py-2 rounded-xl text-xs transition"
            >
              Pause
            </button>
            <button
              @click="gameStore.resetFocusSession"
              class="px-3 py-2 bg-slate-800 hover:bg-slate-700 text-slate-300 text-xs font-semibold rounded-xl transition border border-slate-700"
            >
              Reset
            </button>
          </div>
        </div>
      </div>

      <!-- Main Body: Quests & Leaderboard Grid -->
      <div class="grid grid-cols-1 lg:grid-cols-2 gap-8">
        <!-- Quest Engine -->
        <div class="bg-slate-900/80 rounded-3xl border border-slate-800 p-6 sm:p-8 space-y-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <span class="text-2xl">🎯</span>
              <div>
                <h3 class="text-lg font-bold text-white">Daily Quests & Milestones</h3>
                <p class="text-xs text-slate-400">Complete tasks to unlock bonus XP and keep your streak</p>
              </div>
            </div>
          </div>

          <div class="space-y-4">
            <div
              v-for="quest in gameStore.quests"
              :key="quest.id"
              class="bg-slate-800/60 rounded-2xl p-4 border border-slate-700/60 flex items-center justify-between gap-4 transition hover:border-slate-600"
            >
              <div class="flex-1 space-y-1">
                <div class="flex items-center gap-2">
                  <span
                    class="text-[10px] font-extrabold uppercase px-2 py-0.5 rounded-full"
                    :class="quest.quest?.questType === 'weekly' ? 'bg-purple-900/60 text-purple-300 border border-purple-700/50' : 'bg-emerald-900/60 text-emerald-300 border border-emerald-700/50'"
                  >
                    {{ quest.quest?.questType || 'Daily' }}
                  </span>
                  <h4 class="text-sm font-bold text-slate-100">{{ quest.quest?.title }}</h4>
                </div>
                <p class="text-xs text-slate-400 leading-snug">{{ quest.quest?.description }}</p>

                <!-- Progress mini bar -->
                <div class="flex items-center gap-2 pt-2">
                  <div class="flex-1 h-1.5 bg-slate-700 rounded-full overflow-hidden">
                    <div
                      class="h-full bg-emerald-500 rounded-full"
                      :style="{ width: `${quest.isCompleted ? 100 : 35}%` }"
                    ></div>
                  </div>
                  <span class="text-[11px] font-mono text-slate-400">
                    {{ quest.isCompleted ? 'Done' : 'In Progress' }}
                  </span>
                </div>
              </div>

              <!-- Action / Claim Button -->
              <div class="shrink-0">
                <button
                  v-if="quest.isCompleted && !quest.claimed"
                  @click="claim(quest.questId || quest.id)"
                  class="bg-gradient-to-r from-amber-500 to-yellow-400 text-slate-950 font-black px-4 py-2 rounded-xl text-xs shadow-lg shadow-amber-500/20 hover:brightness-110 transition animate-pulse"
                >
                  Claim +{{ quest.quest?.xpReward }} XP
                </button>
                <span
                  v-else-if="quest.claimed"
                  class="text-xs font-bold text-emerald-400 px-3 py-1.5 bg-emerald-950/60 rounded-xl border border-emerald-800/40"
                >
                  ✓ Claimed
                </span>
                <span
                  v-else
                  class="text-xs font-semibold text-slate-400 bg-slate-800 px-3 py-1.5 rounded-xl border border-slate-700"
                >
                  +{{ quest.quest?.xpReward }} XP
                </span>
              </div>
            </div>
          </div>
        </div>

        <!-- Leaderboard -->
        <div class="bg-slate-900/80 rounded-3xl border border-slate-800 p-6 sm:p-8 space-y-6">
          <div class="flex items-center justify-between">
            <div class="flex items-center gap-3">
              <span class="text-2xl">🏆</span>
              <div>
                <h3 class="text-lg font-bold text-white">Competitive Leaderboard</h3>
                <p class="text-xs text-slate-400">Real-time ranks based on verified curriculum mastery</p>
              </div>
            </div>

            <!-- Scope Tabs -->
            <div class="flex bg-slate-800 rounded-xl p-1 border border-slate-700">
              <button
                v-for="s in ['classroom', 'school', 'national']"
                :key="s"
                @click="setScope(s)"
                :class="activeScope === s ? 'bg-emerald-500 text-slate-950 font-bold' : 'text-slate-400 hover:text-white'"
                class="px-2.5 py-1 text-xs rounded-lg capitalize transition"
              >
                {{ s }}
              </button>
            </div>
          </div>

          <!-- Ranking List -->
          <div class="space-y-2.5 max-h-[420px] overflow-y-auto pr-1">
            <div
              v-for="user in gameStore.leaderboard"
              :key="user.userId"
              class="flex items-center justify-between p-3.5 rounded-2xl transition border"
              :class="user.fullName.includes('Keza')
                ? 'bg-emerald-950/40 border-emerald-500/40 shadow-md shadow-emerald-900/20'
                : 'bg-slate-800/40 border-slate-800 hover:bg-slate-800/70'"
            >
              <div class="flex items-center gap-3.5">
                <!-- Rank Pill -->
                <div class="w-7 text-center font-black text-sm">
                  <span v-if="user.rank === 1" class="text-xl">🥇</span>
                  <span v-else-if="user.rank === 2" class="text-xl">🥈</span>
                  <span v-else-if="user.rank === 3" class="text-xl">🥉</span>
                  <span v-else class="text-slate-400">#{{ user.rank }}</span>
                </div>

                <!-- Avatar -->
                <img :src="user.avatarUrl" :alt="user.fullName" class="w-10 h-10 rounded-full bg-slate-800 border border-slate-700 object-cover" />

                <div>
                  <h4 class="text-sm font-bold text-white flex items-center gap-2">
                    <span>{{ user.fullName }}</span>
                    <span v-if="user.fullName.includes('Keza')" class="text-[10px] bg-emerald-500/20 text-emerald-400 px-1.5 py-0.2 rounded font-bold">You</span>
                  </h4>
                  <div class="flex items-center gap-2 text-xs text-slate-400">
                    <span>Lvl {{ user.currentLevel }}</span>
                    <span>•</span>
                    <span class="text-amber-400">🔥 {{ user.streak }}d</span>
                  </div>
                </div>
              </div>

              <div class="text-right">
                <span class="text-sm font-mono font-black text-emerald-400">{{ user.totalXp.toLocaleString() }}</span>
                <span class="text-[10px] text-slate-500 block uppercase font-bold">XP</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useGamificationStore } from '../stores/gamification'

const gameStore = useGamificationStore()
const activeScope = ref('national')

const levelProgressPct = computed(() => {
  const current = gameStore.xp.totalXp
  const toNext = gameStore.xp.xpToNext
  const totalInLevel = current + toNext
  if (totalInLevel <= 0) return 60
  return Math.min(100, Math.round((current / totalInLevel) * 100))
})

onMounted(async () => {
  await gameStore.fetchProfile()
  await gameStore.fetchLeaderboard(activeScope.value)
})

function setScope(scope) {
  activeScope.value = scope
  gameStore.fetchLeaderboard(scope)
}

function useFreeze() {
  gameStore.freezeStreak()
}

function claim(questId) {
  gameStore.claimQuest(questId)
}

function formatTimer(seconds) {
  const m = Math.floor(seconds / 60)
  const s = seconds % 60
  return `${m.toString().padStart(2, '0')}:${s.toString().padStart(2, '0')}`
}
</script>

