import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'

export const useGamificationStore = defineStore('gamification', () => {
  const xp = ref({ totalXp: 840, currentLevel: 4, xpToNext: 220 })
  const streak = ref({ currentStreak: 7, longestStreak: 12, freezePasses: 2 })
  const quests = ref([])
  const badges = ref([])
  const leaderboard = ref([])
  const loading = ref(false)
  const isFocusing = ref(false)
  const focusTimeRemaining = ref(25 * 60) // 25 mins in seconds
  let focusInterval = null

  async function fetchProfile() {
    loading.value = true
    try {
      const res = await api.get('/api/v1/gamification/profile')
      if (res.data?.data) {
        xp.value = res.data.data.xp || xp.value
        streak.value = res.data.data.streak || streak.value
        quests.value = res.data.data.quests || []
        badges.value = res.data.data.recentBadges || []
      }
    } catch (err) {
      console.warn('Could not load profile from backend, using current state', err)
    } finally {
      loading.value = false
    }
  }

  async function freezeStreak() {
    try {
      const res = await api.post('/api/v1/gamification/streak/freeze')
      if (res.data?.data) {
        streak.value = res.data.data
      }
      return true
    } catch (err) {
      console.error('Streak freeze failed', err)
      return false
    }
  }

  async function claimQuest(questId) {
    try {
      const res = await api.post(`/api/v1/gamification/quests/${questId}/claim`)
      const reward = res.data?.data?.xpReward || 0
      // Update local quest state
      const target = quests.value.find(q => q.questId === questId || q.id === questId)
      if (target) {
        target.claimed = true
      }
      // Add XP locally
      xp.value.totalXp += reward
      return reward
    } catch (err) {
      console.error('Failed to claim quest', err)
      return 0
    }
  }

  async function fetchLeaderboard(scope = 'national') {
    try {
      const res = await api.get(`/api/v1/gamification/leaderboard?scope=${scope}`)
      if (res.data?.data) {
        leaderboard.value = res.data.data
      }
    } catch (err) {
      console.warn('Failed to load leaderboard', err)
    }
  }

  function startFocusSession() {
    isFocusing.value = true
    if (focusInterval) clearInterval(focusInterval)
    focusInterval = setInterval(() => {
      if (focusTimeRemaining.value > 0) {
        focusTimeRemaining.value--
      } else {
        completeFocusSession(25)
      }
    }, 1000)
  }

  function pauseFocusSession() {
    isFocusing.value = false
    if (focusInterval) {
      clearInterval(focusInterval)
      focusInterval = null
    }
  }

  function resetFocusSession() {
    pauseFocusSession()
    focusTimeRemaining.value = 25 * 60
  }

  async function completeFocusSession(durationMinutes = 25) {
    pauseFocusSession()
    focusTimeRemaining.value = 25 * 60
    try {
      const res = await api.post('/api/v1/gamification/focus', { durationMinutes })
      const bonus = res.data?.data?.bonusXp || 15
      xp.value.totalXp += bonus
      return bonus
    } catch (err) {
      xp.value.totalXp += 15
      return 15
    }
  }

  function addLocalXP(amount) {
    xp.value.totalXp += amount
  }

  return {
    xp,
    streak,
    quests,
    badges,
    leaderboard,
    loading,
    isFocusing,
    focusTimeRemaining,
    fetchProfile,
    freezeStreak,
    claimQuest,
    fetchLeaderboard,
    startFocusSession,
    pauseFocusSession,
    resetFocusSession,
    completeFocusSession,
    addLocalXP,
  }
})

