<template>
  <div>
    <div class="flex justify-between items-center mb-6">
      <div>
        <h1 class="text-3xl font-bold text-brand-dark">Content Management</h1>
        <p class="text-brand-gray mt-2">Manage all static content across your website</p>
      </div>
      <button 
        @click="openCreateModal" 
        class="bg-brand-primary text-brand-dark px-4 py-2 rounded-lg hover:bg-brand-light transition font-semibold"
      >
        + Add Content
      </button>
    </div>

    <!-- Content List -->
    <div class="bg-brand-white rounded-lg shadow border border-brand-gray">
      <div class="overflow-x-auto">
        <table class="w-full">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">Key</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">Title</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">Section</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">Type</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">Status</th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">Actions</th>
            </tr>
          </thead>
          <tbody class="bg-brand-white divide-y divide-brand-gray">
            <tr v-for="content in contents" :key="content.id">
              <td class="px-6 py-4 whitespace-nowrap text-sm font-mono text-brand-dark">
                {{ content.key }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-brand-dark">
                {{ content.title || '-' }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-brand-gray">
                {{ content.section }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-brand-gray">
                {{ content.type }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="content.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" 
                      class="px-2 py-1 text-xs rounded-full">
                  {{ content.is_active ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-brand-gray">
                <button 
                  @click="editContent(content)" 
                  class="text-brand-primary hover:text-brand-light mr-3 font-medium"
                >
                  Edit
                </button>
                <button 
                  @click="deleteContent(content.id)" 
                  class="text-red-600 hover:text-red-800 font-medium"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>

    <!-- Create/Edit Modal -->
    <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
      <div class="bg-white rounded-lg max-w-2xl w-full max-h-[90vh] overflow-y-auto">
        <div class="p-6">
          <div class="flex justify-between items-center mb-4">
            <h2 class="text-xl font-bold">{{ isEditing ? 'Edit Content' : 'Create Content' }}</h2>
            <button @click="closeModal" class="text-gray-500 hover:text-gray-700">
              <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
              </svg>
            </button>
          </div>

          <form @submit.prevent="saveContent">
            <div class="grid grid-cols-2 gap-4 mb-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Key *</label>
                <input 
                  v-model="currentContent.key" 
                  type="text" 
                  required
                  :disabled="isEditing"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="e.g., hero.title"
                >
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Section *</label>
                <select 
                  v-model="currentContent.section" 
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="">Select Section</option>
                  <option value="home">Home</option>
                  <option value="about">About</option>
                  <option value="contact">Contact</option>
                  <option value="general">General</option>
                </select>
              </div>
            </div>

            <div class="grid grid-cols-2 gap-4 mb-4">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Title</label>
                <input 
                  v-model="currentContent.title" 
                  type="text" 
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="Display title"
                >
              </div>
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Type</label>
                <select 
                  v-model="currentContent.type" 
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                  <option value="text">Text</option>
                  <option value="html">HTML</option>
                  <option value="markdown">Markdown</option>
                  <option value="json">JSON</option>
                </select>
              </div>
            </div>

            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">Content *</label>
              <textarea 
                v-model="currentContent.content" 
                required
                rows="6"
                class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                placeholder="Enter your content here..."
              ></textarea>
            </div>

            <div class="grid grid-cols-2 gap-4 mb-6">
              <div>
                <label class="block text-sm font-medium text-gray-700 mb-2">Order</label>
                <input 
                  v-model.number="currentContent.order" 
                  type="number" 
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="0"
                >
              </div>
              <div class="flex items-center">
                <label class="flex items-center">
                  <input 
                    v-model="currentContent.is_active" 
                    type="checkbox" 
                    class="mr-2 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded"
                  >
                  <span class="text-sm font-medium text-gray-700">Active</span>
                </label>
              </div>
            </div>

            <div class="flex justify-end space-x-3">
              <button 
                type="button" 
                @click="closeModal"
                class="px-4 py-2 text-gray-700 bg-gray-200 rounded-md hover:bg-gray-300 transition"
              >
                Cancel
              </button>
              <button 
                type="submit"
                class="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700 transition"
              >
                {{ isEditing ? 'Update' : 'Create' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'

definePageMeta({
  layout: 'admin'
})

const authStore = useAuthStore()
const contents = ref([])
const showModal = ref(false)
const isEditing = ref(false)
const currentContent = ref({
  key: '',
  title: '',
  content: '',
  type: 'text',
  section: '',
  order: 0,
  is_active: true
})

const fetchContents = async () => {
  try {
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/content`, {
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    
    if (response.ok) {
      const data = await response.json()
      contents.value = data || []
    }
  } catch (error) {
    console.error('Error fetching contents:', error)
  }
}

const openCreateModal = () => {
  isEditing.value = false
  currentContent.value = {
    key: '',
    title: '',
    content: '',
    type: 'text',
    section: '',
    order: 0,
    is_active: true
  }
  showModal.value = true
}

const editContent = (content) => {
  isEditing.value = true
  currentContent.value = { ...content }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  isEditing.value = false
  currentContent.value = {
    key: '',
    title: '',
    content: '',
    type: 'text',
    section: '',
    order: 0,
    is_active: true
  }
}

const saveContent = async () => {
  try {
    const config = useRuntimeConfig()
    const url = isEditing.value 
      ? `${config.public.apiBase}/api/content/${currentContent.value.id}`
      : `${config.public.apiBase}/api/content`
    
    const method = isEditing.value ? 'PUT' : 'POST'
    
    const response = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify(currentContent.value)
    })
    
    if (response.ok) {
      await fetchContents()
      closeModal()
    } else {
      const error = await response.json()
      console.error('Error saving content:', error)
      alert('Error saving content: ' + (error.error || 'Unknown error'))
    }
  } catch (error) {
    console.error('Error saving content:', error)
    alert('Error saving content')
  }
}

const deleteContent = async (id) => {
  if (!confirm('Are you sure you want to delete this content?')) {
    return
  }
  
  try {
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/content/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    
    if (response.ok) {
      await fetchContents()
    } else {
      alert('Error deleting content')
    }
  } catch (error) {
    console.error('Error deleting content:', error)
    alert('Error deleting content')
  }
}

onMounted(() => {
  fetchContents()
})
</script>