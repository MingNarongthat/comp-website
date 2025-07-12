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
const article = ref(null)
const loading = ref(true)

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

const fetchArticle = async () => {
  try {
    const articleId = route.params.id
    const config = useRuntimeConfig()
    
    // Try to fetch without auth first (for public access)
    let response = await fetch(`${config.public.apiBase}/api/articles/${articleId}`)
    
    if (response.ok) {
      const data = await response.json()
      article.value = data
    } else {
      article.value = null
    }
  } catch (error) {
    console.log('Could not fetch article')
    article.value = null
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchArticle()
})

useSeoMeta({
  title: () => article.value ? `${article.value.title} - Personal Web` : 'Article - Personal Web',
  description: () => article.value ? article.value.content.substring(0, 160) + '...' : 'Read this article on Personal Web'
})
</script>
