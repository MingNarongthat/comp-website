<template>
  <div>
    <Header />
    <main class="min-h-screen bg-gray-100">
      <!-- Navigation Breadcrumb -->
      <section class="bg-brand-white shadow-sm border-b border-brand-gray">
        <div class="container mx-auto px-4 py-4">
          <nav class="flex items-center space-x-2 text-sm">
            <NuxtLink to="/" class="text-brand-primary hover:text-brand-light transition font-medium">
              Home
            </NuxtLink>
            <svg class="w-4 h-4 text-brand-gray" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
            </svg>
            <NuxtLink to="/articles" class="text-brand-primary hover:text-brand-light transition font-medium">
              Articles
            </NuxtLink>
            <svg class="w-4 h-4 text-brand-gray" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
            </svg>
            <span class="text-brand-dark">
              {{ article?.title || 'Loading...' }}
            </span>
          </nav>
        </div>
      </section>

      <!-- Article Content -->
      <section class="py-12">
        <div class="container mx-auto px-4">
          <div v-if="loading" class="max-w-4xl mx-auto text-center py-12">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-brand-primary"></div>
            <p class="mt-4 text-brand-gray">Loading article...</p>
          </div>

          <div v-else-if="!article" class="max-w-4xl mx-auto text-center py-12">
            <div class="mb-6">
              <svg class="mx-auto h-12 w-12 text-brand-gray" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
              </svg>
            </div>
            <h2 class="text-2xl font-bold text-brand-dark mb-4">Article Not Found</h2>
            <p class="text-brand-dark mb-8">
              Sorry, the article you're looking for doesn't exist or has been removed.
            </p>
            <div class="flex flex-col sm:flex-row gap-3 justify-center">
              <NuxtLink 
                to="/articles" 
                class="bg-brand-primary text-brand-dark px-6 py-3 rounded-lg hover:bg-brand-light transition font-semibold"
              >
                Browse Articles
              </NuxtLink>
              <NuxtLink 
                to="/" 
                class="bg-gray-200 text-brand-dark px-6 py-3 rounded-lg hover:bg-gray-300 transition font-semibold"
              >
                Back to Home
              </NuxtLink>
            </div>
          </div>

          <article v-else class="max-w-4xl mx-auto">
            <div class="bg-brand-white rounded-lg shadow-lg overflow-hidden border border-brand-gray">
              <!-- Article Header -->
              <div class="p-8 border-b border-brand-gray">
                <h1 class="text-4xl font-bold text-brand-dark mb-4">
                  {{ article.title }}
                </h1>
                <div class="flex items-center text-sm text-brand-gray space-x-4">
                  <span>
                    Published {{ formatDate(article.created_at) }}
                  </span>
                  <span>•</span>
                  <span>
                    {{ readingTime }} min read
                  </span>
                </div>
              </div>

              <!-- Article Images -->
              <div v-if="articleImages && articleImages.length > 0" class="p-8 pb-0">
                <div class="flex justify-center">
                  <div v-if="articleImages.length === 1" class="w-full max-w-2xl">
                    <img 
                      :src="getImageUrl(articleImages[0])" 
                      :alt="article.title"
                      class="article-image"
                    >
                  </div>
                  <div v-else class="grid grid-cols-1 md:grid-cols-2 gap-4 w-full max-w-4xl">
                    <img 
                      v-for="(image, index) in articleImages" 
                      :key="index"
                      :src="getImageUrl(image)" 
                      :alt="`${article.title} - Image ${index + 1}`"
                      class="article-image"
                    >
                  </div>
                </div>
              </div>

              <!-- Article Content -->
              <div class="p-8">
                <div class="prose prose-lg max-w-none">
                  <div class="whitespace-pre-wrap text-brand-dark leading-relaxed">
                    {{ article.content }}
                  </div>
                </div>
              </div>
            </div>

            <!-- Navigation Footer -->
            <div class="mt-12 flex flex-col sm:flex-row gap-4 justify-between">
              <NuxtLink 
                to="/articles" 
                class="inline-flex items-center text-brand-primary hover:text-brand-light transition font-medium"
              >
                <svg class="w-4 h-4 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 19l-7-7 7-7"></path>
                </svg>
                Back to Articles
              </NuxtLink>
              
              <div class="flex gap-3">
                <NuxtLink 
                  to="/contact" 
                  class="bg-brand-primary text-brand-dark px-4 py-2 rounded-lg hover:bg-brand-light transition text-sm font-semibold"
                >
                  Get in Touch
                </NuxtLink>
                <button 
                  @click="shareArticle"
                  class="bg-gray-200 text-brand-dark px-4 py-2 rounded-lg hover:bg-gray-300 transition text-sm font-semibold"
                >
                  Share Article
                </button>
              </div>
            </div>
          </article>
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()

// Fetch article data on server-side for better initial load
const { data: articleData } = await useFetch(`/api/articles/${route.params.id}`)
const article = ref(articleData.value)
const loading = ref(false)

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
}

const readingTime = computed(() => {
  if (!article.value?.content) return 0
  const wordsPerMinute = 200
  const wordCount = article.value.content.split(' ').length
  return Math.ceil(wordCount / wordsPerMinute)
})

const articleImages = computed(() => {
  if (!article.value?.images) return []
  try {
    return typeof article.value.images === 'string' 
      ? JSON.parse(article.value.images) 
      : article.value.images
  } catch {
    return []
  }
})

const shareArticle = () => {
  if (navigator.share) {
    navigator.share({
      title: article.value.title,
      text: article.value.content.substring(0, 150) + '...',
      url: window.location.href,
    })
  } else {
    // Fallback: copy to clipboard
    navigator.clipboard.writeText(window.location.href)
    // You could show a toast notification here
    alert('Link copied to clipboard!')
  }
}

const getImageUrl = (imagePath) => {
  if (imagePath.startsWith('http')) {
    return imagePath
  }
  // Use frontend proxy for images to avoid client-side CORS issues
  if (imagePath.startsWith('/uploads/')) {
    return `/api/proxy${imagePath}`
  }
  return imagePath
}

// Article data is already loaded via SSR

useSeoMeta({
  title: () => article.value ? `${article.value.title} - Personal Web` : 'Article - Personal Web',
  description: () => article.value ? article.value.content.substring(0, 160) + '...' : 'Read this article on Personal Web'
})
</script>

<style scoped>
/* Ensure images display properly */
img {
  display: block;
  max-width: 100%;
  height: auto;
}

.article-image {
  width: 100% !important;
  height: 16rem !important;
  object-fit: cover !important;
  border-radius: 0.5rem !important;
  box-shadow: 0 4px 6px -1px rgba(0, 0, 0, 0.1) !important;
}
</style>
