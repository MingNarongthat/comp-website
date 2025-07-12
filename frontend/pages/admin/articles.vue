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
});

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
  };
  showModal.value = true;
};

const editArticle = (article) => {
  isEditMode.value = true;
  currentArticle.value = { ...article };
  showModal.value = true;
};

const saveArticle = async () => {
  if (isEditMode.value) {
    await articleStore.updateArticle(currentArticle.value);
  } else {
    await articleStore.createArticle(currentArticle.value);
  }
  await fetchArticles();
  closeModal();
};

const deleteArticle = async (id) => {
  if (confirm('Are you sure you want to delete this article?')) {
    await articleStore.deleteArticle(id);
    await fetchArticles();
  }
};

const closeModal = () => {
  showModal.value = false;
};
</script>
