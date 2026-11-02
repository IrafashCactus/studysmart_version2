<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-5xl mx-auto space-y-8">
      <!-- Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b border-slate-800 pb-6">
        <div>
          <div class="flex items-center gap-2">
            <span class="text-3xl">🤖</span>
            <h1 class="text-2xl font-black text-white">AI Interactive Tutor & Problem Explainer</h1>
          </div>
          <p class="text-sm text-slate-400 mt-1">
            Curriculum-grounded step-by-step breakdowns, voice dialogues, and adaptive remediation.
          </p>
        </div>

        <!-- Mode Toggle Pill -->
        <div class="flex bg-slate-900 border border-slate-800 p-1 rounded-2xl shrink-0">
          <button
            @click="activeMode = 'explainer'"
            :class="activeMode === 'explainer' ? 'bg-emerald-500 text-slate-950 font-bold' : 'text-slate-400 hover:text-white'"
            class="px-4 py-2 rounded-xl text-xs font-semibold transition flex items-center gap-1.5"
          >
            <span>📝</span>
            <span>Problem Explainer</span>
          </button>
          <button
            @click="activeMode = 'voice'"
            :class="activeMode === 'voice' ? 'bg-indigo-500 text-white font-bold' : 'text-slate-400 hover:text-white'"
            class="px-4 py-2 rounded-xl text-xs font-semibold transition flex items-center gap-1.5"
          >
            <span>🎙️</span>
            <span>Voice-to-Voice Prep</span>
          </button>
        </div>
      </div>

      <!-- Mode 1: Instant Problem Explainer -->
      <div v-if="activeMode === 'explainer'" class="space-y-6">
        <!-- Input Card -->
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 sm:p-8 shadow-xl space-y-4">
          <div class="flex items-center justify-between">
            <label class="text-xs font-extrabold uppercase tracking-wider text-emerald-400">
              Type or Paste Any STEM Question (Math, Physics, Bio, Chemistry)
            </label>
            <span class="text-xs text-slate-500">REB Syllabus Grounded</span>
          </div>

          <div class="relative">
            <textarea
              v-model="problemInput"
              rows="3"
              placeholder="e.g. Find the roots and nature of discriminant for 2x² - 4x + 2 = 0, or calculate projectile velocity at max height..."
              class="w-full bg-slate-950 border border-slate-800 rounded-2xl p-4 text-sm text-slate-100 placeholder-slate-500 focus:outline-none focus:border-emerald-500 transition"
            ></textarea>
          </div>

          <!-- Quick Sample Prompt Pills -->
          <div class="flex items-center gap-2 overflow-x-auto pb-1 text-xs text-slate-400">
            <span class="shrink-0 font-semibold text-slate-500">Quick Try:</span>
            <button
              @click="setSamplePrompt('Solve 2x² - 4x + 2 = 0 and determine discriminant Δ')"
              class="bg-slate-800 hover:bg-slate-700 px-3 py-1.5 rounded-lg text-slate-300 shrink-0 border border-slate-700/60 transition"
            >
              📐 Quadratic Roots & Δ
            </button>
            <button
              @click="setSamplePrompt('A projectile is launched at 30° with 20 m/s. What is velocity at maximum height?')"
              class="bg-slate-800 hover:bg-slate-700 px-3 py-1.5 rounded-lg text-slate-300 shrink-0 border border-slate-700/60 transition"
            >
              🚀 Projectile Peak Motion
            </button>
            <button
              @click="setSamplePrompt('Explain homozygous recessive vs dominant genotype with Punnett square')"
              class="bg-slate-800 hover:bg-slate-700 px-3 py-1.5 rounded-lg text-slate-300 shrink-0 border border-slate-700/60 transition"
            >
              🧬 Mendelian Genetics Cross
            </button>
          </div>

          <!-- Submit Button -->
          <div class="flex justify-end pt-2">
            <button
              @click="explainProblem"
              :disabled="loading || !problemInput.trim()"
              class="bg-emerald-500 hover:bg-emerald-400 text-slate-950 font-black px-6 py-3 rounded-xl text-xs transition shadow-lg shadow-emerald-500/20 disabled:opacity-50 flex items-center gap-2"
            >
              <span v-if="loading" class="animate-spin">⏳</span>
              <span>{{ loading ? 'Generating Step-by-Step Breakdown...' : 'Get AI Step-by-Step Breakdown →' }}</span>
            </button>
          </div>
        </div>

        <!-- Explanation Results -->
        <div v-if="explanation" class="bg-slate-900 border border-slate-800 rounded-3xl p-6 sm:p-8 shadow-2xl space-y-6">
          <div class="border-b border-slate-800 pb-4">
            <span class="text-xs uppercase font-extrabold text-emerald-400 tracking-wider">Concept Summary</span>
            <h3 class="text-xl font-black text-white mt-1">{{ explanation.summary }}</h3>
            <p class="text-xs text-slate-400 mt-2 bg-emerald-950/40 border border-emerald-800/40 p-3 rounded-xl text-emerald-200">
              💡 <strong>Core Principle:</strong> {{ explanation.keyConcept }}
            </p>
          </div>

          <!-- Step Cards -->
          <div class="space-y-4">
            <h4 class="text-xs uppercase font-bold text-slate-400 tracking-wider">Step-by-Step Mathematical Derivation</h4>
            <div
              v-for="step in explanation.steps"
              :key="step.stepNumber"
              class="bg-slate-800/50 border border-slate-800 rounded-2xl p-5 space-y-2 transition hover:border-slate-700"
            >
              <div class="flex items-center gap-3">
                <span class="w-6 h-6 rounded-full bg-emerald-500 text-slate-950 font-extrabold text-xs flex items-center justify-center">
                  {{ step.stepNumber }}
                </span>
                <h5 class="text-sm font-bold text-white">{{ step.title }}</h5>
              </div>
              <p class="text-sm text-slate-300 pl-9 leading-relaxed">{{ step.explanation }}</p>
              <div v-if="step.formula" class="ml-9 bg-slate-950 border border-slate-700/60 rounded-xl px-4 py-2 font-mono text-emerald-300 text-xs inline-block">
                {{ step.formula }}
              </div>
            </div>
          </div>

          <!-- Final Verified Answer -->
          <div class="bg-gradient-to-r from-emerald-950/60 to-slate-900 border border-emerald-500/40 rounded-2xl p-5 flex items-center justify-between">
            <div>
              <span class="text-xs font-bold text-emerald-400 uppercase tracking-wider block">Final Verified Solution</span>
              <span class="text-base font-black text-white font-mono mt-0.5 block">{{ explanation.finalAnswer }}</span>
            </div>
            <span class="text-2xl">✅</span>
          </div>
        </div>

        <!-- Adaptive Remediation Section -->
        <div v-if="recs.length > 0" class="bg-slate-900 border border-slate-800 rounded-3xl p-6 sm:p-8 space-y-4">
          <div class="flex items-center gap-2">
            <span class="text-xl">🎯</span>
            <h3 class="text-base font-bold text-white">Adaptive Remediation Queue</h3>
          </div>
          <p class="text-xs text-slate-400">
            Telemetry detected topics below 75% accuracy. Here are your high-priority review cards:
          </p>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-4 pt-2">
            <div
              v-for="rec in recs"
              :key="rec.weakTopicTitle"
              class="bg-slate-800/60 border border-amber-500/30 rounded-2xl p-4 space-y-3"
            >
              <div class="flex items-center justify-between">
                <h5 class="text-sm font-bold text-white">{{ rec.weakTopicTitle }}</h5>
                <span class="text-xs font-bold text-amber-400 bg-amber-500/10 px-2 py-0.5 rounded border border-amber-500/30">
                  {{ rec.masteryPct }}% Mastery
                </span>
              </div>
              <p class="text-xs text-slate-400">Recommended micro-cards injected back into your Mind-Feed stream.</p>
              <router-link
                to="/feed"
                class="inline-flex items-center gap-1.5 text-xs font-bold text-emerald-400 hover:text-emerald-300 transition"
              >
                <span>Study Remediation Cards</span>
                <span>→</span>
              </router-link>
            </div>
          </div>
        </div>
      </div>

      <!-- Mode 2: Voice-to-Voice AI Tutor Simulation -->
      <div v-else class="bg-slate-900 border border-slate-800 rounded-3xl p-6 sm:p-10 shadow-2xl space-y-8 text-center max-w-2xl mx-auto">
        <div>
          <span class="text-xs uppercase font-extrabold text-indigo-400 tracking-wider">Oral Exam & Dialogue Prep</span>
          <h2 class="text-2xl font-black text-white mt-1">Real-Time Voice AI Dialogue</h2>
          <p class="text-xs text-slate-400 mt-2 max-w-md mx-auto">
            Speak naturally in Kinyarwanda, English, or French. The AI evaluates pronunciation, terminology, and concept accuracy.
          </p>
        </div>

        <!-- Waveform Simulation -->
        <div class="h-32 bg-slate-950 rounded-2xl border border-slate-800 flex items-center justify-center gap-1.5 px-6">
          <div
            v-for="i in 28"
            :key="i"
            class="w-1.5 bg-indigo-500 rounded-full transition-all duration-150"
            :class="isSpeaking ? 'animate-pulse' : 'opacity-30'"
            :style="{ height: isSpeaking ? `${Math.floor(Math.sin(i) * 35 + 45)}px` : '12px' }"
          ></div>
        </div>

        <!-- Dialogue Bubble -->
        <div class="bg-slate-800/80 rounded-2xl p-5 border border-slate-700/60 text-left space-y-2">
          <div class="flex items-center gap-2 text-xs font-bold text-indigo-300">
            <span>🎙️ AI Tutor:</span>
          </div>
          <p class="text-sm text-slate-200 leading-relaxed">
            "Hello Keza! Today let's review your Physics Kinematics oral exam. Can you tell me what happens to horizontal velocity vx when an object moves through projectile motion without air resistance?"
          </p>
        </div>

        <!-- Push to Talk Button -->
        <div class="pt-4 flex flex-col items-center gap-3">
          <button
            @click="toggleVoice"
            :class="isSpeaking ? 'bg-rose-600 hover:bg-rose-500 scale-105' : 'bg-indigo-600 hover:bg-indigo-500'"
            class="w-20 h-20 rounded-full flex items-center justify-center text-3xl text-white shadow-2xl transition transform active:scale-95 border-4 border-slate-800"
          >
            <span>{{ isSpeaking ? '⏹️' : '🎙️' }}</span>
          </button>
          <span class="text-xs font-semibold text-slate-400">
            {{ isSpeaking ? 'Listening... Tap to conclude response' : 'Tap to start voice practice session' }}
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const activeMode = ref('explainer')
const problemInput = ref('')
const loading = ref(false)
const explanation = ref(null)
const recs = ref([])
const isSpeaking = ref(false)

onMounted(async () => {
  fetchAdaptive()
})

function setSamplePrompt(text) {
  problemInput.value = text
  explainProblem()
}

async function explainProblem() {
  if (!problemInput.value.trim()) return
  loading.value = true
  try {
    const res = await api.post('/api/v1/tutor/explain', {
      problemText: problemInput.value,
      subjectCode: 'STEM',
    })
    if (res.data?.data) {
      explanation.value = res.data.data
    }
  } catch (err) {
    console.error('Failed to get explanation', err)
  } finally {
    loading.value = false
  }
}

async function fetchAdaptive() {
  try {
    const res = await api.get('/api/v1/tutor/adaptive')
    if (res.data?.data) {
      recs.value = res.data.data
    }
  } catch (err) {
    console.warn('Adaptive recommendations unavailable', err)
  }
}

function toggleVoice() {
  isSpeaking.value = !isSpeaking.value
}
</script>

