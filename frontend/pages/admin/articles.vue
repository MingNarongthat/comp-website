<template>
  <NuxtLayout name="admin">
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-brand-dark">Manage Articles</h1>
      <button
        @click="openCreateModal"
        class="px-4 py-2 mb-4 text-brand-dark bg-brand-primary rounded-md hover:bg-brand-light font-semibold"
      >
        Create New Article
      </button>

      <div class="bg-brand-white shadow-md rounded-lg overflow-hidden border border-brand-gray">
        <table class="min-w-full divide-y divide-brand-gray">
          <thead class="bg-gray-100">
            <tr>
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider"
              >
                Title
              </th>
              <th
                scope="col"
                class="px-6 py-3 text-left text-xs font-medium text-brand-dark uppercase tracking-wider"
              >
                Actions
              </th>
            </tr>
          </thead>
          <tbody class="bg-brand-white divide-y divide-brand-gray">
            <tr v-for="article in articles" :key="article.id">
              <td class="px-6 py-4 whitespace-nowrap">
                <div class="text-sm font-medium text-brand-dark">{{ article.title }}</div>
              </td>
              <td class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium">
                <button
                  @click="editArticle(article)"
                  class="text-brand-primary hover:text-brand-light mr-4 font-medium"
                >
                  Edit
                </button>
                <button
                  @click="deleteArticle(article.id)"
                  class="text-red-600 hover:text-red-800 font-medium"
                >
                  Delete
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <!-- Create/Edit Article Modal -->
      <div
        v-if="showModal"
        class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full flex items-center justify-center"
      >
        <div class="relative p-8 border border-brand-gray w-full max-w-md md:max-w-lg lg:max-w-xl shadow-lg rounded-md bg-brand-white">
          <h3 class="text-2xl font-bold mb-4 text-brand-dark">{{ isEditMode ? 'Edit Article' : 'Create Article' }}</h3>
          <form @submit.prevent="saveArticle">
            <div class="mb-4">
              <label for="title" class="block text-sm font-medium text-brand-dark">Title</label>
              <input
                type="text"
                id="title"
                v-model="currentArticle.title"
                class="mt-1 block w-full border border-brand-gray rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-brand-primary focus:border-brand-primary"
              />
            </div>
            <div class="mb-4">
              <label for="content" class="block text-sm font-medium text-brand-dark">Content</label>
              <textarea
                id="content"
                v-model="currentArticle.content"
                rows="10"
                class="mt-1 block w-full border border-brand-gray rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-brand-primary focus:border-brand-primary"
              ></textarea>
            </div>
            <div class="mb-4">
              <label for="images" class="block text-sm font-medium text-brand-dark">Images (Multiple)</label>
              <input
                type="file"
                id="images"
                ref="imageInput"
                @change="handleImageSelect"
                multiple
                accept="image/*"
                class="mt-1 block w-full border border-brand-gray rounded-md shadow-sm py-2 px-3 focus:outline-none focus:ring-brand-primary focus:border-brand-primary"
              />
              <p class="text-xs text-gray-500 mt-1">Select multiple images (JPEG, PNG, GIF). Max 5MB per image.</p>
              
              <!-- Image Preview -->
              <div v-if="selectedImages.length > 0" class="mt-3">
                <h4 class="text-sm font-medium text-brand-dark mb-2">Selected Images:</h4>
                <div class="grid grid-cols-2 md:grid-cols-3 gap-2">
                  <div v-for="(image, index) in selectedImages" :key="index" class="relative">
                    <img :src="image.preview" :alt="`Preview ${index + 1}`" class="w-full h-20 object-cover rounded border">
                    <button
                      @click="removeImage(index)"
                      type="button"
                      class="absolute -top-1 -right-1 bg-red-500 text-white rounded-full w-5 h-5 flex items-center justify-center text-xs hover:bg-red-600"
                    >
                      ×
                    </button>
                  </div>
                </div>
              </div>

              <!-- Existing Images (for edit mode) -->
              <div v-if="currentArticle.images && currentArticle.images.length > 0" class="mt-3">
                <h4 class="text-sm font-medium text-brand-dark mb-2">Current Images:</h4>
                <div class="grid grid-cols-2 md:grid-cols-3 gap-2">
                  <div v-for="(imageUrl, index) in currentArticle.images" :key="index" class="relative">
                    <img :src="getImageUrl(imageUrl)" :alt="`Current ${index + 1}`" class="w-full h-20 object-cover rounded border">
                    <button
                      @click="removeCurrentImage(index)"
                      type="button"
                      class="absolute -top-1 -right-1 bg-red-500 text-white rounded-full w-5 h-5 flex items-center justify-center text-xs hover:bg-red-600"
                    >
                      ×
                    </button>
                  </div>
                </div>
              </div>
            </div>
            <div class="flex justify-end">
              <button
                type="button"
                @click="closeModal"
                class="px-4 py-2 mr-2 text-brand-dark bg-gray-200 rounded-md hover:bg-gray-300 font-semibold"
              >
                Cancel
              </button>
              <button
                type="submit"
                class="px-4 py-2 text-brand-dark bg-brand-primary rounded-md hover:bg-brand-light font-semibold"
              >
                {{ isEditMode ? 'Update' : 'Create' }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { useArticleStore } from '@/stores/article';
import { useAuthStore } from '@/stores/auth';

definePageMeta({
  middleware: ['auth', 'superadmin'],
});

const articleStore = useArticleStore();
const articles = ref([]);
const showModal = ref(false);
const isEditMode = ref(false);
const currentArticle = ref({
  id: null,
  title: '',
  content: '',
  images: [],
});

const selectedImages = ref([]);
const imageInput = ref(null);

onMounted(async () => {
  await fetchArticles();
});

const fetchArticles = async () => {
  articles.value = await articleStore.fetchArticles();
};

const openCreateModal = () => {
  isEditMode.value = false;
  currentArticle.value = {
    id: null,
    title: '',
    content: '',
    images: [],
  };
  selectedImages.value = [];
  showModal.value = true;
};

const editArticle = (article) => {
  isEditMode.value = true;
  currentArticle.value = { 
    ...article,
    images: article.images ? (typeof article.images === 'string' ? JSON.parse(article.images) : article.images) : []
  };
  selectedImages.value = [];
  showModal.value = true;
};

const saveArticle = async () => {
  try {
    // Upload new images if any
    let uploadedImageUrls = [];
    if (selectedImages.value.length > 0) {
      uploadedImageUrls = await uploadImages();
    }

    // Combine existing images with newly uploaded ones
    const allImages = [...(currentArticle.value.images || []), ...uploadedImageUrls];
    const articleData = {
      ...currentArticle.value,
      images: allImages
    };

    if (isEditMode.value) {
      await articleStore.updateArticle(articleData);
    } else {
      await articleStore.createArticle(articleData);
    }
    await fetchArticles();
    closeModal();
  } catch (error) {
    console.error('Error saving article:', error);
    alert('Failed to save article. Please try again.');
  }
};

const deleteArticle = async (id) => {
  if (confirm('Are you sure you want to delete this article?')) {
    await articleStore.deleteArticle(id);
    await fetchArticles();
  }
};

const closeModal = () => {
  showModal.value = false;
  selectedImages.value = [];
  if (imageInput.value) {
    imageInput.value.value = '';
  }
};

const handleImageSelect = (event) => {
  const files = Array.from(event.target.files);
  selectedImages.value = [];
  
  files.forEach(file => {
    if (file.size > 5 * 1024 * 1024) {
      alert(`File ${file.name} is too large. Maximum size is 5MB.`);
      return;
    }
    
    const reader = new FileReader();
    reader.onload = (e) => {
      selectedImages.value.push({
        file: file,
        preview: e.target.result
      });
    };
    reader.readAsDataURL(file);
  });
};

const removeImage = (index) => {
  selectedImages.value.splice(index, 1);
};

const removeCurrentImage = (index) => {
  currentArticle.value.images.splice(index, 1);
};

const uploadImages = async () => {
  const formData = new FormData();
  selectedImages.value.forEach(image => {
    formData.append('images', image.file);
  });

  const config = useRuntimeConfig();
  const authStore = useAuthStore();
  
  const response = await fetch(`${config.public.apiBase}/api/articles/upload-images`, {
    method: 'POST',
    headers: {
      'Authorization': `Bearer ${authStore.token}`
    },
    body: formData
  });

  if (!response.ok) {
    throw new Error('Failed to upload images');
  }

  const data = await response.json();
  return data.images;
};

const getImageUrl = (imagePath) => {
  const config = useRuntimeConfig();
  if (imagePath.startsWith('http')) {
    return imagePath;
  }
  return `${config.public.apiBase}${imagePath}`;
};

</script>
