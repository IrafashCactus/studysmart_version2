<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 flex flex-col pb-16">
    <!-- Top Filter Bar -->
    <div class="sticky top-16 z-30 bg-slate-900/90 backdrop-blur-md border-b border-slate-800 px-4 py-3">
      <div class="max-w-3xl mx-auto flex items-center justify-between gap-3 overflow-x-auto no-scrollbar">
        <!-- Subject Pills -->
        <div class="flex items-center gap-2">
          <button
            @click="selectSubject('ALL')"
            :class="feedStore.activeSubject === 'ALL' && !feedStore.bookmarkedOnly
              ? 'bg-emerald-500 text-slate-950 font-bold shadow-md shadow-emerald-500/20'
              : 'bg-slate-800 text-slate-300 hover:bg-slate-700'"
            class="px-3.5 py-1.5 rounded-full text-xs font-semibold transition-all shrink-0"
          >
            🌟 All Subjects
          </button>
          <button
            v-for="sub in feedStore.subjects"
            :key="sub.id"
            @click="selectSubject(sub.code)"
            :class="feedStore.activeSubject === sub.code && !feedStore.bookmarkedOnly
              ? 'bg-emerald-500 text-slate-950 font-bold shadow-md shadow-emerald-500/20'
              : 'bg-slate-800 text-slate-300 hover:bg-slate-700'"
            class="px-3.5 py-1.5 rounded-full text-xs font-semibold transition-all shrink-0 flex items-center gap-1.5"
          >
            <span>{{ sub.name }}</span>
          </button>
        </div>

        <!-- Bookmarked Toggle Button -->
        <button
          @click="toggleBookmarkedView"
          :class="feedStore.bookmarkedOnly
            ? 'bg-amber-500 text-slate-950 font-bold'
            : 'bg-slate-800 text-slate-300 hover:bg-slate-700'"
          class="px-3 py-1.5 rounded-full text-xs font-semibold transition-all shrink-0 flex items-center gap-1.5 ml-auto border border-slate-700"
        >
          <span>⭐</span>
          <span>Saved ({{ feedStore.bookmarks.length }})</span>
        </button>
      </div>
    </div>

    <!-- Feed Content -->
    <div class="max-w-2xl w-full mx-auto px-4 py-6 flex-1 flex flex-col justify-center">
      <!-- Loading State -->
      <div v-if="feedStore.loading && displayCards.length === 0" class="text-center py-20">
        <div class="w-12 h-12 border-4 border-emerald-500 border-t-transparent rounded-full animate-spin mx-auto mb-4"></div>
        <p class="text-slate-400 text-sm">Curating personalized micro-feed...</p>
      </div>

      <!-- Empty State -->
      <div v-else-if="displayCards.length === 0" class="text-center py-16 bg-slate-900 rounded-2xl border border-slate-800 p-8">
        <span class="text-5xl mb-4 block">📖</span>
        <h3 class="text-lg font-bold text-white mb-2">No Cards in this Category</h3>
        <p class="text-slate-400 text-sm mb-6">Switch topic or clear filters to view more curriculum bites.</p>
        <button @click="selectSubject('ALL')" class="bg-emerald-500 text-slate-950 font-bold px-5 py-2.5 rounded-xl hover:bg-emerald-400 text-sm transition">
          Reset Topic Filter
        </button>
      </div>

      <!-- Active Micro-Card Container -->
      <div v-else class="space-y-4">
        <!-- Feed Card Progress Indicator -->
        <div class="flex items-center justify-between text-xs text-slate-400 px-1">
          <div class="flex items-center gap-2">
            <span class="inline-flex items-center px-2 py-0.5 rounded-full text-xs font-medium bg-emerald-950 text-emerald-400 border border-emerald-800/40">
              {{ currentCard.topic?.subject?.name || 'Curriculum Bite' }}
            </span>
            <span>• Unit {{ currentCard.topic?.unitNumber || 1 }}</span>
          </div>
          <span class="font-mono">{{ feedStore.currentIndex + 1 }} / {{ displayCards.length }}</span>
        </div>

        <!-- The Interactive Card -->
        <div class="relative bg-gradient-to-b from-slate-900 to-slate-900/90 rounded-3xl border border-slate-800 shadow-2xl p-6 sm:p-8 overflow-hidden transition-all">
          <!-- Ambient Glow by Subject -->
          <div class="absolute -top-24 -right-24 w-48 h-48 bg-emerald-500/10 rounded-full blur-3xl pointer-events-none"></div>

          <!-- Card Header -->
          <div class="flex items-start justify-between gap-4 mb-4">
            <div>
              <span class="text-xs uppercase tracking-wider font-extrabold text-emerald-400">
                {{ formatCardType(currentCard.cardType) }}
              </span>
              <h2 class="text-2xl font-black text-white mt-1 leading-snug">
                {{ currentCard.title }}
              </h2>
            </div>
            <!-- Bookmark Button -->
            <button
              @click="toggleBookmark(currentCard.id)"
              class="w-10 h-10 rounded-full flex items-center justify-center transition shrink-0"
              :class="currentCard.isBookmarked
                ? 'bg-amber-500/20 text-amber-400 border border-amber-500/30'
                : 'bg-slate-800 text-slate-400 hover:text-white'"
              title="Bookmark for dedicated review"
            >
              <span class="text-lg">{{ currentCard.isBookmarked ? '★' : '☆' }}</span>
            </button>
          </div>

          <!-- Flashcard Mode (Click to Flip) -->
          <div
            v-if="currentCard.cardType === 'flashcard'"
            @click="isFlipped = !isFlipped"
            class="my-6 cursor-pointer group"
          >
            <div
              class="rounded-2xl p-6 transition-all duration-300 border"
              :class="isFlipped
                ? 'bg-emerald-950/40 border-emerald-500/30 text-emerald-100 shadow-inner'
                : 'bg-slate-800/60 border-slate-700/60 hover:border-slate-600'"
            >
              <div class="flex items-center justify-between text-xs text-slate-400 mb-2">
                <span class="font-bold uppercase tracking-wider">{{ isFlipped ? 'Answer Side' : 'Question Side' }}</span>
                <span class="text-emerald-400 group-hover:underline">Tap to flip 🔄</span>
              </div>
              <p class="text-lg font-medium leading-relaxed">
                {{ isFlipped ? currentCard.back : currentCard.front }}
              </p>
            </div>
          </div>

          <!-- Standard Concept Body -->
          <div v-else class="my-6">
            <p class="text-base text-slate-200 leading-relaxed whitespace-pre-line">
              {{ currentCard.contentText }}
            </p>
          </div>

          <!-- Key Takeaway Banner -->
          <div v-if="currentCard.keyTakeaway" class="my-4 bg-emerald-500/10 border-l-4 border-emerald-500 rounded-r-xl p-3.5">
            <div class="flex items-start gap-2">
              <span class="text-emerald-400 text-sm">💡</span>
              <p class="text-xs text-emerald-200 font-medium leading-relaxed">
                <strong class="text-emerald-300">Key Takeaway:</strong> {{ currentCard.keyTakeaway }}
              </p>
            </div>
          </div>

          <!-- Inline Interactive Check (Embedded assessment) -->
          <div v-if="currentCard.questionPrompt" class="mt-6 pt-6 border-t border-slate-800">
            <div class="flex items-center gap-2 mb-3">
              <span class="w-2 h-2 rounded-full bg-emerald-400 animate-pulse"></span>
              <span class="text-xs font-bold uppercase tracking-wider text-slate-300">Quick Check (+5 XP)</span>
            </div>
            <p class="text-sm font-semibold text-white mb-4">
              {{ currentCard.questionPrompt }}
            </p>

            <!-- Options Grid -->
            <div class="grid grid-cols-1 gap-2.5">
              <button
                v-for="opt in parsedOptions"
                :key="opt"
                @click="checkAnswer(opt)"
                :disabled="hasAnswered"
                :class="getOptionClass(opt)"
                class="w-full text-left p-3.5 rounded-xl text-sm font-medium transition flex items-center justify-between border"
              >
                <span>{{ opt }}</span>
                <span v-if="hasAnswered && opt === currentCard.correctAnswer" class="text-emerald-400 font-bold">✓</span>
                <span v-else-if="hasAnswered && selectedOption === opt && opt !== currentCard.correctAnswer" class="text-rose-400 font-bold">✗</span>
              </button>
            </div>

            <!-- Explanation Feedback -->
            <div v-if="hasAnswered" class="mt-4 p-4 rounded-xl text-xs leading-relaxed transition"
              :class="isCorrect ? 'bg-emerald-950/60 border border-emerald-800/60 text-emerald-200' : 'bg-rose-950/60 border border-rose-800/60 text-rose-200'">
              <p class="font-bold mb-1 flex items-center gap-1.5">
                <span>{{ isCorrect ? '🎉 Correct!' : '⚠️ Review Note:' }}</span>
                <span v-if="isCorrect" class="text-emerald-400 font-extrabold">+5 XP Earned</span>
              </p>
              <p>{{ currentCard.explanation }}</p>
            </div>
          </div>

          <!-- Navigation Swipe Bar at bottom of card -->
          <div class="mt-8 pt-4 flex items-center justify-between border-t border-slate-800/80">
            <button
              @click="prevCard"
              :disabled="feedStore.currentIndex === 0"
              class="px-4 py-2 rounded-xl text-xs font-bold text-slate-400 hover:text-white disabled:opacity-30 transition flex items-center gap-1.5"
            >
              <span>← Previous</span>
            </button>

            <button
              @click="markCompleted"
              class="bg-emerald-500 hover:bg-emerald-400 text-slate-950 font-extrabold px-6 py-2.5 rounded-xl text-xs transition shadow-lg shadow-emerald-500/20 flex items-center gap-2"
            >
              <span>Next Bite</span>
              <span>→</span>
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, watch } from 'vue'
import { useFeedStore } from '../stores/feed'

const feedStore = useFeedStore()

const isFlipped = ref(false)
const selectedOption = ref(null)
const hasAnswered = ref(false)
const isCorrect = ref(false)

const displayCards = computed(() => {
  return feedStore.bookmarkedOnly ? feedStore.bookmarks : feedStore.cards
})

const currentCard = computed(() => {
  return displayCards.value[feedStore.currentIndex] || {}
})

const parsedOptions = computed(() => {
  if (!currentCard.value?.optionsJson) return []
  try {
    return JSON.parse(currentCard.value.optionsJson)
  } catch (e) {
    return []
  }
})

onMounted(async () => {
  await feedStore.fetchSubjects()
  await feedStore.fetchFeed(true)
  await feedStore.fetchBookmarks()
})

watch(() => feedStore.currentIndex, () => {
  isFlipped.value = false
  selectedOption.value = null
  hasAnswered.value = false
  isCorrect.value = false
})

function selectSubject(code) {
  feedStore.bookmarkedOnly = false
  feedStore.setSubject(code)
}

function toggleBookmarkedView() {
  feedStore.bookmarkedOnly = !feedStore.bookmarkedOnly
  if (feedStore.bookmarkedOnly) {
    feedStore.fetchBookmarks()
  }
}

function toggleBookmark(cardId) {
  feedStore.toggleBookmark(cardId)
}

function checkAnswer(opt) {
  if (hasAnswered.value) return
  selectedOption.value = opt
  hasAnswered.value = true
  isCorrect.value = (opt === currentCard.value.correctAnswer)

  if (isCorrect.value) {
    feedStore.recordInteraction(currentCard.value.id, 'complete', 8)
  }
}

function getOptionClass(opt) {
  if (!hasAnswered.value) {
    return 'bg-slate-800/60 border-slate-700/60 hover:bg-slate-800 hover:border-emerald-500/50 text-slate-200'
  }
  if (opt === currentCard.value.correctAnswer) {
    return 'bg-emerald-950/80 border-emerald-500 text-emerald-200 shadow-sm'
  }
  if (selectedOption.value === opt && opt !== currentCard.value.correctAnswer) {
    return 'bg-rose-950/80 border-rose-500 text-rose-200'
  }
  return 'bg-slate-900 border-slate-800 text-slate-500 opacity-60'
}

function markCompleted() {
  if (currentCard.value?.id) {
    feedStore.recordInteraction(currentCard.value.id, 'complete', 10)
  }
  feedStore.nextCard()
}

function prevCard() {
  feedStore.prevCard()
}

function formatCardType(type) {
  switch (type) {
    case 'flashcard': return '⚡ Flashcard (Flip)'
    case 'video_30s': return '🎬 30-Sec Summary Video'
    case 'summary': return '📋 Unit Summary Bite'
    default: return '💡 Micro-Concept'
  }
}
</script>

<style scoped>
.no-scrollbar::-webkit-scrollbar {
  display: none;
}
.no-scrollbar {
  -ms-overflow-style: none;
  scrollbar-width: none;
}
</style>

