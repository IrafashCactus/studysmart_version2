import { defineStore } from 'pinia'
import { ref } from 'vue'
import api from '../services/api'
import { useGamificationStore } from './gamification'
import offlineSync from '../services/offlineSync'

export const useFeedStore = defineStore('feed', () => {
  const cards = ref([])
  const currentIndex = ref(0)
  const subjects = ref([])
  const activeSubject = ref('ALL')
  const bookmarkedOnly = ref(false)
  const bookmarks = ref([])
  const loading = ref(false)
  const page = ref(1)
  const hasMore = ref(true)
  const isOffline = ref(!navigator.onLine)

  window.addEventListener('studySmartNetworkChange', (e) => {
    isOffline.value = !e.detail.isOnline
    if (e.detail.isOnline) {
      offlineSync.syncQueuedInteractions(api)
    }
  })

  async function fetchSubjects() {
    try {
      const res = await api.get('/api/v1/feed/subjects')
      if (res.data?.data) {
        subjects.value = res.data.data
      }
    } catch (err) {
      console.warn('Failed to load subjects', err)
    }
  }

  async function fetchFeed(reset = false) {
    if (reset) {
      page.value = 1
      cards.value = []
      currentIndex.value = 0
    }
    loading.value = true
    try {
      const subjectParam = activeSubject.value === 'ALL' ? '' : activeSubject.value
      const res = await api.get(`/api/v1/feed?subject=${subjectParam}&page=${page.value}&perPage=10`)
      if (res.data?.data) {
        const newCards = res.data.data.cards || []
        cards.value = reset ? newCards : [...cards.value, ...newCards]
        hasMore.value = res.data.data.hasMore
        // Cache to IndexedDB for offline capability
        offlineSync.cacheCards(cards.value)
      }
    } catch (err) {
      console.warn('Online feed request failed, falling back to local IndexedDB cache', err)
      const cached = await offlineSync.getCachedCards(activeSubject.value)
      if (cached.length > 0) {
        cards.value = cached
        isOffline.value = true
      }
    } finally {
      loading.value = false
    }
  }

  async function fetchBookmarks() {
    try {
      const res = await api.get('/api/v1/feed/bookmarks')
      if (res.data?.data) {
        bookmarks.value = res.data.data
      }
    } catch (err) {
      console.warn('Failed to fetch bookmarks', err)
    }
  }

  async function toggleBookmark(cardId) {
    try {
      const res = await api.post(`/api/v1/feed/bookmarks/${cardId}`)
      const isBookmarked = res.data?.data?.isBookmarked
      const card = cards.value.find(c => c.id === cardId)
      if (card) {
        card.isBookmarked = isBookmarked
      }
      fetchBookmarks()
      return isBookmarked
    } catch (err) {
      console.error('Failed to toggle bookmark', err)
    }
  }

  async function recordInteraction(cardId, action, durationSec = 5) {
    try {
      const res = await api.post('/api/v1/feed/interact', {
        cardId,
        action,
        durationSec,
      })
      const xpEarned = res.data?.data?.xpEarned || 0
      if (xpEarned > 0) {
        const gameStore = useGamificationStore()
        gameStore.addLocalXP(xpEarned)
      }
      return xpEarned
    } catch (err) {
      console.warn('Network offline, queuing interaction locally in IndexedDB', err)
      await offlineSync.queueInteraction({ cardId, action, durationSec })
      if (action === 'complete') {
        const gameStore = useGamificationStore()
        gameStore.addLocalXP(5)
        return 5
      }
      return 0
    }
  }

  function setSubject(code) {
    activeSubject.value = code
    fetchFeed(true)
  }

  function nextCard() {
    if (currentIndex.value < cards.value.length - 1) {
      currentIndex.value++
      const current = cards.value[currentIndex.value]
      if (current) {
        recordInteraction(current.id, 'view', 3)
      }
    } else if (hasMore.value && !loading.value) {
      page.value++
      fetchFeed(false)
    }
  }

  function prevCard() {
    if (currentIndex.value > 0) {
      currentIndex.value--
    }
  }

  return {
    cards,
    currentIndex,
    subjects,
    activeSubject,
    bookmarkedOnly,
    bookmarks,
    loading,
    hasMore,
    isOffline,
    fetchSubjects,
    fetchFeed,
    fetchBookmarks,
    toggleBookmark,
    recordInteraction,
    setSubject,
    nextCard,
    prevCard,
  }
})
