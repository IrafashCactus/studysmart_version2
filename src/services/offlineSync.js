/**
 * Offline Sync Service using IndexedDB for Low-Connectivity / Offline-First Studying
 * Designed for Rwandan schools and rural learning centers.
 */

const DB_NAME = 'StudySmartOfflineDB'
const DB_VERSION = 1

class OfflineSyncService {
  constructor() {
    this.db = null
    this.isOnline = navigator.onLine
    this.initListeners()
  }

  initListeners() {
    window.addEventListener('online', () => {
      this.isOnline = true
      window.dispatchEvent(new CustomEvent('studySmartNetworkChange', { detail: { isOnline: true } }))
      this.syncQueuedInteractions()
    })

    window.addEventListener('offline', () => {
      this.isOnline = false
      window.dispatchEvent(new CustomEvent('studySmartNetworkChange', { detail: { isOnline: false } }))
    })
  }

  async getDB() {
    if (this.db) return this.db

    return new Promise((resolve, reject) => {
      const request = indexedDB.open(DB_NAME, DB_VERSION)

      request.onupgradeneeded = (event) => {
        const db = event.target.result

        // Store cached feed cards
        if (!db.objectStoreNames.contains('cachedCards')) {
          const cardStore = db.createObjectStore('cachedCards', { keyPath: 'id' })
          cardStore.createIndex('subject', 'topic.subject.code', { unique: false })
        }

        // Store offline pending telemetry interactions
        if (!db.objectStoreNames.contains('pendingQueue')) {
          db.createObjectStore('pendingQueue', { keyPath: 'queueId', autoIncrement: true })
        }

        // Store cached quizzes
        if (!db.objectStoreNames.contains('cachedQuizzes')) {
          db.createObjectStore('cachedQuizzes', { keyPath: 'id' })
        }
      }

      request.onsuccess = (event) => {
        this.db = event.target.result
        resolve(this.db)
      }

      request.onerror = (event) => {
        reject(event.target.error)
      }
    })
  }

  // Cache cards locally for offline studying
  async cacheCards(cards) {
    if (!cards || cards.length === 0) return
    const db = await this.getDB()
    const tx = db.transaction('cachedCards', 'readwrite')
    const store = tx.objectStore('cachedCards')

    for (const card of cards) {
      store.put(card)
    }

    return new Promise((resolve) => {
      tx.oncomplete = () => resolve(true)
    })
  }

  // Retrieve cached cards while offline
  async getCachedCards(subjectCode = 'ALL') {
    const db = await this.getDB()
    const tx = db.transaction('cachedCards', 'readonly')
    const store = tx.objectStore('cachedCards')

    return new Promise((resolve) => {
      const request = store.getAll()
      request.onsuccess = () => {
        let cards = request.result || []
        if (subjectCode !== 'ALL') {
          cards = cards.filter(c => c.topic?.subject?.code === subjectCode)
        }
        resolve(cards)
      }
      request.onerror = () => resolve([])
    })
  }

  // Queue interaction while offline
  async queueInteraction(interactionData) {
    const db = await this.getDB()
    const tx = db.transaction('pendingQueue', 'readwrite')
    const store = tx.objectStore('pendingQueue')

    store.add({
      ...interactionData,
      queuedAt: new Date().toISOString(),
    })

    return new Promise((resolve) => {
      tx.oncomplete = () => resolve(true)
    })
  }

  // Re-sync with server when connectivity returns
  async syncQueuedInteractions(apiClient) {
    if (!this.isOnline) return 0
    const db = await this.getDB()
    const tx = db.transaction('pendingQueue', 'readonly')
    const store = tx.objectStore('pendingQueue')

    const items = await new Promise((resolve) => {
      const req = store.getAll()
      req.onsuccess = () => resolve(req.result || [])
      req.onerror = () => resolve([])
    })

    if (items.length === 0) return 0

    // Push pending telemetry to server
    for (const item of items) {
      try {
        if (apiClient) {
          await apiClient.post('/api/v1/feed/interact', {
            cardId: item.cardId,
            action: item.action,
            durationSec: item.durationSec || 5,
          })
        }
      } catch (err) {
        console.warn('Sync item failed:', err)
      }
    }

    // Clear queue after sync
    const clearTx = db.transaction('pendingQueue', 'readwrite')
    clearTx.objectStore('pendingQueue').clear()

    return items.length
  }
}

export const offlineSync = new OfflineSyncService()
export default offlineSync

