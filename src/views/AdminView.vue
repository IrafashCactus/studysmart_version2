<template>
  <div class="min-h-screen bg-slate-950 text-slate-100 py-8 px-4 sm:px-6 lg:px-8">
    <div class="max-w-6xl mx-auto space-y-8">
      <!-- Header -->
      <div class="flex flex-col md:flex-row md:items-center justify-between gap-4 border-b border-slate-800 pb-6">
        <div>
          <div class="flex items-center gap-2">
            <span class="text-3xl">⚙️</span>
            <h1 class="text-2xl font-black text-white">Admin & RAG Content Engine Control</h1>
          </div>
          <p class="text-sm text-slate-400 mt-1">
            Curriculum document ingestion, AI vectorization, human-in-the-loop verification, and RBAC user directory.
          </p>
        </div>

        <!-- Navigation Tabs -->
        <div class="flex bg-slate-900 border border-slate-800 p-1 rounded-2xl shrink-0 overflow-x-auto">
          <button
            v-for="tab in [
              { id: 'workbench', label: 'Verification Workbench', icon: '🔍' },
              { id: 'rag', label: 'RAG Pipeline & Documents', icon: '📄' },
              { id: 'users', label: 'User Directory & RBAC', icon: '👥' },
              { id: 'flags', label: 'Content Flags', icon: '🚩' },
            ]"
            :key="tab.id"
            @click="activeTab = tab.id"
            :class="activeTab === tab.id ? 'bg-emerald-500 text-slate-950 font-bold' : 'text-slate-400 hover:text-white'"
            class="px-3.5 py-2 rounded-xl text-xs font-semibold transition flex items-center gap-1.5 shrink-0"
          >
            <span>{{ tab.icon }}</span>
            <span>{{ tab.label }}</span>
          </button>
        </div>
      </div>

      <!-- TAB 1: Content Verification Workbench -->
      <div v-if="activeTab === 'workbench'" class="space-y-6">
        <div class="flex items-center justify-between">
          <div>
            <h2 class="text-lg font-bold text-white">Human-in-the-Loop Content Verification</h2>
            <p class="text-xs text-slate-400">Review AI-extracted questions and cards before publishing to live student Mind-Feeds.</p>
          </div>
          <span class="text-xs bg-amber-500/10 text-amber-400 border border-amber-500/30 px-3 py-1 rounded-full font-bold">
            {{ generatedItems.length }} Pending Review
          </span>
        </div>

        <div v-if="generatedItems.length === 0" class="text-center py-16 bg-slate-900 rounded-3xl border border-slate-800 p-8">
          <span class="text-4xl block mb-2">🎉</span>
          <h3 class="text-base font-bold text-white">All Generated Content Verified</h3>
          <p class="text-xs text-slate-400 mt-1">Upload new textbooks in the RAG Pipeline tab to generate more assessments.</p>
        </div>

        <div v-else class="grid grid-cols-1 gap-6">
          <div
            v-for="item in generatedItems"
            :key="item.id"
            class="bg-slate-900 border border-slate-800 rounded-3xl p-6 shadow-xl space-y-4"
          >
            <div class="flex items-start justify-between gap-4">
              <div>
                <span class="text-[10px] font-extrabold uppercase px-2.5 py-1 rounded-full bg-indigo-950 text-indigo-300 border border-indigo-800/40">
                  {{ item.contentType }}
                </span>
                <h3 class="text-base font-bold text-white mt-2">{{ item.title }}</h3>
                <p class="text-xs text-slate-400 mt-0.5">Source Doc: {{ item.document?.filename || 'Curriculum Syllabus' }}</p>
              </div>
              <span class="text-xs bg-amber-500/20 text-amber-300 px-2.5 py-1 rounded-full font-semibold capitalize border border-amber-500/30">
                {{ item.status }}
              </span>
            </div>

            <!-- Content Preview Box -->
            <div class="bg-slate-950 rounded-2xl p-4 border border-slate-800 text-xs font-mono text-slate-300 whitespace-pre-wrap leading-relaxed">
              {{ formatContentJSON(item.contentJson) }}
            </div>

            <div class="flex items-center justify-between pt-2 border-t border-slate-800/80">
              <span class="text-xs text-slate-500 italic">{{ item.reviewNotes }}</span>
              <div class="flex items-center gap-2">
                <button
                  @click="rejectContent(item.id)"
                  class="px-4 py-2 bg-rose-500/20 text-rose-300 hover:bg-rose-500/30 text-xs font-bold rounded-xl transition border border-rose-500/30"
                >
                  Reject
                </button>
                <button
                  @click="approveContent(item.id)"
                  class="px-5 py-2 bg-emerald-500 hover:bg-emerald-400 text-slate-950 text-xs font-extrabold rounded-xl transition shadow-md shadow-emerald-500/20"
                >
                  Approve & Publish to Question Bank →
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- TAB 2: RAG Pipeline & Document Ingestion -->
      <div v-else-if="activeTab === 'rag'" class="space-y-6">
        <!-- Upload Zone -->
        <div class="bg-slate-900 border-2 border-dashed border-slate-700 hover:border-emerald-500/60 rounded-3xl p-8 text-center space-y-4 transition">
          <span class="text-5xl block">📥</span>
          <div>
            <h3 class="text-base font-bold text-white">Ingest Raw Rwanda Curriculum Documents</h3>
            <p class="text-xs text-slate-400 max-w-md mx-auto mt-1">
              Supports REB textbooks, national past papers, syllabi, and teacher notes (.pdf, .docx).
            </p>
          </div>

          <div class="flex flex-wrap items-center justify-center gap-3 pt-2">
            <button
              @click="simulateUpload('REB_Chemistry_S5_Organic_2025.pdf', 3800000, 'S5', 'textbook')"
              class="px-4 py-2 bg-slate-800 hover:bg-slate-700 border border-slate-700 rounded-xl text-xs font-semibold text-slate-200 transition"
            >
              + Upload Sample Chemistry S5 PDF
            </button>
            <button
              @click="simulateUpload('REB_National_Math_Exam_2024.pdf', 2100000, 'S6', 'past_exam')"
              class="px-4 py-2 bg-slate-800 hover:bg-slate-700 border border-slate-700 rounded-xl text-xs font-semibold text-slate-200 transition"
            >
              + Upload National Math Exam PDF
            </button>
          </div>
        </div>

        <!-- Ingested Documents Table -->
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 shadow-xl space-y-4">
          <h3 class="text-base font-bold text-white">Ingested Curriculum Document Library</h3>
          <div class="overflow-x-auto">
            <table class="w-full text-left text-xs">
              <thead class="text-slate-400 border-b border-slate-800 uppercase tracking-wider">
                <tr>
                  <th class="pb-3 font-bold">Document Name</th>
                  <th class="pb-3 font-bold">Grade</th>
                  <th class="pb-3 font-bold">Category</th>
                  <th class="pb-3 font-bold">Chunks</th>
                  <th class="pb-3 font-bold">Vector Status</th>
                  <th class="pb-3 font-bold">Added Date</th>
                </tr>
              </thead>
              <tbody class="divide-y divide-slate-800/60 text-slate-300">
                <tr v-for="doc in documents" :key="doc.id" class="hover:bg-slate-800/40">
                  <td class="py-3 font-semibold text-white flex items-center gap-2">
                    <span>📄</span>
                    <span>{{ doc.filename }}</span>
                  </td>
                  <td class="py-3">{{ doc.gradeLevel || 'S4' }}</td>
                  <td class="py-3 capitalize">{{ doc.docCategory || 'textbook' }}</td>
                  <td class="py-3 font-mono text-emerald-400">{{ doc.chunkCount }}</td>
                  <td class="py-3">
                    <span class="px-2 py-0.5 rounded-full text-[10px] font-bold bg-emerald-950 text-emerald-300 border border-emerald-800/40">
                      {{ doc.processingStatus }}
                    </span>
                  </td>
                  <td class="py-3 text-slate-400">{{ formatDate(doc.createdAt) }}</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
      </div>

      <!-- TAB 3: User Directory & RBAC -->
      <div v-else-if="activeTab === 'users'" class="space-y-6">
        <div class="flex flex-col sm:flex-row sm:items-center justify-between gap-4">
          <div>
            <h2 class="text-lg font-bold text-white">User Directory & Permissions</h2>
            <p class="text-xs text-slate-400">Total accounts registered: {{ usersTotal }}</p>
          </div>

          <!-- Search & Filter Controls -->
          <div class="flex items-center gap-3">
            <input
              v-model="userSearch"
              @input="fetchUsers"
              type="text"
              placeholder="Search by name or email..."
              class="bg-slate-900 border border-slate-800 rounded-xl px-3.5 py-2 text-xs text-white placeholder-slate-500 focus:outline-none focus:border-emerald-500"
            />
            <select
              v-model="roleFilter"
              @change="fetchUsers"
              class="bg-slate-900 border border-slate-800 rounded-xl px-3 py-2 text-xs text-white focus:outline-none focus:border-emerald-500"
            >
              <option value="">All Roles</option>
              <option value="student">Student</option>
              <option value="teacher">Teacher</option>
              <option value="parent">Parent</option>
              <option value="super_admin">Admin</option>
            </select>
          </div>
        </div>

        <!-- Users Table -->
        <div class="bg-slate-900 border border-slate-800 rounded-3xl p-6 shadow-xl overflow-x-auto">
          <table class="w-full text-left text-xs">
            <thead class="text-slate-400 border-b border-slate-800 uppercase tracking-wider">
              <tr>
                <th class="pb-3 font-bold">User</th>
                <th class="pb-3 font-bold">Email</th>
                <th class="pb-3 font-bold">Role</th>
                <th class="pb-3 font-bold">Language</th>
                <th class="pb-3 font-bold">Status</th>
                <th class="pb-3 font-bold text-right">Actions</th>
              </tr>
            </thead>
            <tbody class="divide-y divide-slate-800/60 text-slate-300">
              <tr v-for="u in users" :key="u.id" class="hover:bg-slate-800/40">
                <td class="py-3 font-bold text-white flex items-center gap-2.5">
                  <img :src="u.avatarUrl" class="w-7 h-7 rounded-full bg-slate-800" />
                  <span>{{ u.fullName }}</span>
                </td>
                <td class="py-3 text-slate-400 font-mono">{{ u.email }}</td>
                <td class="py-3">
                  <span
                    class="px-2 py-0.5 rounded-full text-[10px] font-bold uppercase"
                    :class="{
                      'bg-emerald-950 text-emerald-300 border border-emerald-800/40': u.role === 'student',
                      'bg-indigo-950 text-indigo-300 border border-indigo-800/40': u.role === 'teacher',
                      'bg-amber-950 text-amber-300 border border-amber-800/40': u.role === 'parent',
                      'bg-rose-950 text-rose-300 border border-rose-800/40': u.role === 'super_admin',
                    }"
                  >
                    {{ u.role }}
                  </span>
                </td>
                <td class="py-3 uppercase">{{ u.languagePref || 'en' }}</td>
                <td class="py-3">
                  <span
                    class="px-2 py-0.5 rounded-full text-[10px] font-bold"
                    :class="u.isActive ? 'bg-emerald-950 text-emerald-400' : 'bg-rose-950 text-rose-400'"
                  >
                    {{ u.isActive ? 'Active' : 'Suspended' }}
                  </span>
                </td>
                <td class="py-3 text-right">
                  <button
                    @click="toggleUserStatus(u)"
                    class="px-2.5 py-1 rounded-lg text-[11px] font-bold transition border"
                    :class="u.isActive
                      ? 'bg-rose-500/10 text-rose-300 border-rose-500/30 hover:bg-rose-500/20'
                      : 'bg-emerald-500/10 text-emerald-300 border-emerald-500/30 hover:bg-emerald-500/20'"
                  >
                    {{ u.isActive ? 'Suspend' : 'Activate' }}
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>

      <!-- TAB 4: Content Flags -->
      <div v-else class="space-y-6">
        <div class="flex items-center justify-between">
          <div>
            <h2 class="text-lg font-bold text-white">Content Flagging Moderation Queue</h2>
            <p class="text-xs text-slate-400">Reports filed by students or teachers encountering syllabus errors or ambiguities.</p>
          </div>
        </div>

        <div v-if="flags.length === 0" class="text-center py-16 bg-slate-900 rounded-3xl border border-slate-800 p-8">
          <span class="text-4xl block mb-2">🛡️</span>
          <h3 class="text-base font-bold text-white">No Flagged Items</h3>
          <p class="text-xs text-slate-400 mt-1">Platform content is in good health.</p>
        </div>

        <div v-else class="space-y-4">
          <div
            v-for="flag in flags"
            :key="flag.id"
            class="bg-slate-900 border border-slate-800 rounded-3xl p-6 flex items-start justify-between gap-4"
          >
            <div>
              <span class="text-xs font-bold text-rose-400 uppercase tracking-wider">Reason: {{ flag.reason }}</span>
              <p class="text-sm text-slate-200 mt-1">{{ flag.details || 'Student reported question prompt ambiguity in Unit 2.' }}</p>
              <span class="text-xs text-slate-500 mt-2 block">Status: {{ flag.status }}</span>
            </div>
            <button
              v-if="flag.status === 'pending'"
              @click="resolveFlag(flag.id)"
              class="px-4 py-2 bg-emerald-500 hover:bg-emerald-400 text-slate-950 font-bold text-xs rounded-xl transition"
            >
              Resolve Flag
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import api from '../services/api'

const activeTab = ref('workbench')

// Workbench
const generatedItems = ref([])
// Documents
const documents = ref([])
// Users
const users = ref([])
const usersTotal = ref(0)
const userSearch = ref('')
const roleFilter = ref('')
// Flags
const flags = ref([])

onMounted(() => {
  fetchWorkbench()
  fetchDocuments()
  fetchUsers()
  fetchFlags()
})

async function fetchWorkbench() {
  try {
    const res = await api.get('/api/v1/admin/generated?status=draft')
    if (res.data?.data) {
      generatedItems.value = res.data.data
    }
  } catch (err) {
    console.warn('Workbench error', err)
  }
}

async function approveContent(id) {
  try {
    await api.post(`/api/v1/admin/generated/${id}/approve`, { notes: 'Verified by reviewer' })
    fetchWorkbench()
  } catch (err) {
    console.error('Approve failed', err)
  }
}

async function rejectContent(id) {
  try {
    await api.post(`/api/v1/admin/generated/${id}/reject`, { notes: 'Does not meet syllabus criteria' })
    fetchWorkbench()
  } catch (err) {
    console.error('Reject failed', err)
  }
}

async function fetchDocuments() {
  try {
    const res = await api.get('/api/v1/admin/documents')
    if (res.data?.data) {
      documents.value = res.data.data
    }
  } catch (err) {
    console.warn('Docs error', err)
  }
}

async function simulateUpload(filename, size, grade, category) {
  try {
    await api.post('/api/v1/admin/documents', {
      filename,
      fileSizeBytes: size,
      gradeLevel: grade,
      category,
    })
    fetchDocuments()
    fetchWorkbench()
  } catch (err) {
    console.error('Upload failed', err)
  }
}

async function fetchUsers() {
  try {
    const res = await api.get(`/api/v1/admin/users?search=${userSearch.value}&role=${roleFilter.value}`)
    if (res.data?.data) {
      users.value = res.data.data.users || []
      usersTotal.value = res.data.data.total || 0
    }
  } catch (err) {
    console.warn('Users error', err)
  }
}

async function toggleUserStatus(user) {
  try {
    await api.patch(`/api/v1/admin/users/${user.id}/status`, {
      isActive: !user.isActive,
    })
    fetchUsers()
  } catch (err) {
    console.error('Toggle status failed', err)
  }
}

async function fetchFlags() {
  try {
    const res = await api.get('/api/v1/admin/flags')
    if (res.data?.data) {
      flags.value = res.data.data
    }
  } catch (err) {
    console.warn('Flags error', err)
  }
}

async function resolveFlag(id) {
  try {
    await api.patch(`/api/v1/admin/flags/${id}`, { resolutionNote: 'Reviewed and confirmed correct' })
    fetchFlags()
  } catch (err) {
    console.error('Resolve flag failed', err)
  }
}

function formatContentJSON(jsonStr) {
  try {
    return JSON.stringify(JSON.parse(jsonStr), null, 2)
  } catch (e) {
    return jsonStr
  }
}

function formatDate(dateStr) {
  if (!dateStr) return 'Today'
  return new Date(dateStr).toLocaleDateString()
}
</script>

