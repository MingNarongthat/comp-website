<template>
  <NuxtLayout name="admin">
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-brand-dark">Manage Trusted Companies</h1>
      <button
        @click="openCreateModal"
        class="px-4 py-2 mb-4 text-brand-dark bg-brand-primary rounded-md hover:bg-brand-light font-semibold"
      >
        Add Trusted Company
      </button>

      <div class="bg-brand-white shadow-md rounded-lg overflow-hidden border border-brand-gray">
        <table class="min-w-full divide-y divide-brand-gray">
          <thead class="bg-gray-100">
            <tr>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">
                Logo
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">
                Company Name
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">
                Website
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">
                Status
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">
                Order
              </th>
              <th class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider">
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-brand-white divide-y divide-brand-gray">
            <tr v-for="company in companies" :key="company.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <img :src="getLogoUrl(company.logo_url)" :alt="company.name" class="h-12 w-12 object-contain">
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-brand-dark">{{ company.name }}</div>
                <div class="text-sm text-brand-gray">{{ company.description }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <a v-if="company.website" :href="company.website" target="_blank" 
                   class="text-brand-primary hover:text-brand-light text-sm">
                  {{ company.website }}
                </a>
                <span v-else class="text-brand-gray text-sm">-</span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap">
                <span :class="company.is_active ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'" 
                      class="px-2 py-1 text-xs rounded-full">
                  {{ company.is_active ? 'Active' : 'Inactive' }}
                </span>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-sm text-brand-gray">
                {{ company.order }}
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button
                  @click="editCompany(company)"
                  class="text-brand-primary hover:text-brand-light mr-4 font-medium"
                >
                  Edit
                </button>
                <button
                  @click="deleteCompany(company.id)"
                  class="text-red-600 hover:text-red-800 font-medium"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Create/Edit Modal -->
      <div v-if="showModal" class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center p-4 z-50">
        <div class="bg-white rounded-lg max-w-lg w-full max-h-[90vh] overflow-y-auto">
          <div class="p-6">
            <div class="flex justify-between items-center mb-4">
              <h2 class="text-xl font-bold">{{ isEditing ? 'Edit Company' : 'Add Trusted Company' }}</h2>
              <button @click="closeModal" class="text-gray-500 hover:text-gray-700">
                <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6M6 6l12 12"></path>
                </svg>
              </button>
            </div>

            <form @submit.prevent="saveCompany">
              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">Company Name *</label>
                <input 
                  v-model="currentCompany.name" 
                  type="text" 
                  required
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="Enter company name"
                >
              </div>

              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">Company Logo *</label>
                <div v-if="currentCompany.logo_url" class="mb-3">
                  <img :src="getLogoUrl(currentCompany.logo_url)" alt="Current logo" class="h-16 w-auto object-contain border rounded">
                  <p class="text-xs text-gray-500 mt-1">Current logo</p>
                </div>
                <input 
                  ref="logoFileInput"
                  type="file" 
                  accept="image/*"
                  @change="handleFileUpload"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                >
                <p class="text-xs text-gray-500 mt-1">Upload an image file (max 5MB). Supported formats: JPG, PNG, GIF, SVG</p>
                <div v-if="uploading" class="mt-2">
                  <div class="flex items-center space-x-2">
                    <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-blue-600"></div>
                    <span class="text-sm text-gray-600">Uploading...</span>
                  </div>
                </div>
              </div>

              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">Website</label>
                <input 
                  v-model="currentCompany.website" 
                  type="url" 
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="https://company-website.com"
                >
              </div>

              <div class="mb-4">
                <label class="block text-sm font-medium text-gray-700 mb-2">Description</label>
                <textarea 
                  v-model="currentCompany.description" 
                  rows="3"
                  class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                  placeholder="Brief description of the company..."
                ></textarea>
              </div>

              <div class="grid grid-cols-2 gap-4 mb-6">
                <div>
                  <label class="block text-sm font-medium text-gray-700 mb-2">Display Order</label>
                  <input 
                    v-model.number="currentCompany.order" 
                    type="number" 
                    class="w-full px-3 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-500"
                    placeholder="0"
                  >
                </div>
                <div class="flex items-center">
                  <label class="flex items-center">
                    <input 
                      v-model="currentCompany.is_active" 
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
  </NuxtLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useAuthStore } from '@/stores/auth'

definePageMeta({
  middleware: ['auth']
})

const authStore = useAuthStore()
const companies = ref([])
const showModal = ref(false)
const isEditing = ref(false)
const uploading = ref(false)
const logoFileInput = ref(null)
const currentCompany = ref({
  name: '',
  logo_url: '',
  website: '',
  description: '',
  is_active: true,
  order: 0
})

const fetchCompanies = async () => {
  try {
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/companies`, {
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    
    if (response.ok) {
      const data = await response.json()
      companies.value = data || []
    }
  } catch (error) {
    console.error('Error fetching companies:', error)
  }
}

const openCreateModal = () => {
  isEditing.value = false
  currentCompany.value = {
    name: '',
    logo_url: '',
    website: '',
    description: '',
    is_active: true,
    order: 0
  }
  showModal.value = true
}

const editCompany = (company) => {
  isEditing.value = true
  currentCompany.value = { ...company }
  showModal.value = true
}

const closeModal = () => {
  showModal.value = false
  isEditing.value = false
  uploading.value = false
  if (logoFileInput.value) {
    logoFileInput.value.value = ''
  }
  currentCompany.value = {
    name: '',
    logo_url: '',
    website: '',
    description: '',
    is_active: true,
    order: 0
  }
}

const saveCompany = async () => {
  // Validate required fields
  if (!currentCompany.value.name) {
    alert('Company name is required')
    return
  }
  
  if (!isEditing.value && !currentCompany.value.logo_url) {
    alert('Please upload a company logo')
    return
  }
  
  try {
    const config = useRuntimeConfig()
    const url = isEditing.value 
      ? `${config.public.apiBase}/api/companies/${currentCompany.value.id}`
      : `${config.public.apiBase}/api/companies`
    
    const method = isEditing.value ? 'PUT' : 'POST'
    
    const response = await fetch(url, {
      method,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${authStore.token}`
      },
      body: JSON.stringify(currentCompany.value)
    })
    
    if (response.ok) {
      await fetchCompanies()
      closeModal()
    } else {
      const error = await response.json()
      console.error('Error saving company:', error)
      alert('Error saving company: ' + (error.error || 'Unknown error'))
    }
  } catch (error) {
    console.error('Error saving company:', error)
    alert('Error saving company')
  }
}

const deleteCompany = async (id) => {
  if (!confirm('Are you sure you want to delete this company?')) {
    return
  }
  
  try {
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/companies/${id}`, {
      method: 'DELETE',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      }
    })
    
    if (response.ok) {
      await fetchCompanies()
    } else {
      alert('Error deleting company')
    }
  } catch (error) {
    console.error('Error deleting company:', error)
    alert('Error deleting company')
  }
}

const handleFileUpload = async (event) => {
  const file = event.target.files[0]
  if (!file) return

  uploading.value = true
  
  try {
    const formData = new FormData()
    formData.append('logo', file)
    
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/companies/upload-logo`, {
      method: 'POST',
      headers: {
        'Authorization': `Bearer ${authStore.token}`
      },
      body: formData
    })
    
    if (response.ok) {
      const data = await response.json()
      currentCompany.value.logo_url = data.logo_url
    } else {
      const error = await response.json()
      alert('Error uploading file: ' + (error.error || 'Unknown error'))
    }
  } catch (error) {
    console.error('Error uploading file:', error)
    alert('Error uploading file')
  } finally {
    uploading.value = false
  }
}

const getLogoUrl = (logoPath) => {
  const config = useRuntimeConfig()
  if (logoPath.startsWith('http')) {
    return logoPath // For existing URLs
  }
  return `${config.public.apiBase}${logoPath}`
}

onMounted(() => {
  fetchCompanies()
})
</script>