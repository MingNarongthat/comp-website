import { defineStore } from 'pinia';

export const useArticleStore = defineStore('article', {
  state: () => ({
    articles: [],
  }),
  actions: {
    async fetchArticles() {
      const authStore = useAuthStore();
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/api/articles`, {
        headers: {
          Authorization: `Bearer ${authStore.token}`,
        },
      });
      if (!response.ok) {
        throw new Error('Failed to fetch articles');
      }
      const data = await response.json();
      this.articles = data;
      return data;
    },
    async fetchArticleById(id) {
      const authStore = useAuthStore();
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/api/articles/${id}`, {
        headers: {
          Authorization: `Bearer ${authStore.token}`,
        },
      });
      if (!response.ok) {
        throw new Error('Failed to fetch article');
      }
      const data = await response.json();
      return data;
    },
    async createArticle(article) {
      const authStore = useAuthStore();
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/api/articles`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify(article),
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to create article');
      }
      const data = await response.json();
      this.articles.push(data);
      return data;
    },
    async updateArticle(article) {
      const authStore = useAuthStore();
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/api/articles/${article.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
          Authorization: `Bearer ${authStore.token}`,
        },
        body: JSON.stringify(article),
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to update article');
      }
      const data = await response.json();
      const index = this.articles.findIndex((a) => a.id === data.id);
      if (index !== -1) {
        this.articles[index] = data;
      }
      return data;
    },
    async deleteArticle(id) {
      const authStore = useAuthStore();
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/api/articles/${id}`, {
        method: 'DELETE',
        headers: {
          Authorization: `Bearer ${authStore.token}`,
        },
      });
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Failed to delete article');
      }
      this.articles = this.articles.filter((article) => article.id !== id);
    },
  },
});
