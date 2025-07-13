<template>
  <div>
    <Header />
    <main class="min-h-screen bg-gray-100">
      <!-- Articles Content -->
      <section class="py-12">
        <div class="container mx-auto px-4">
          <div v-if="loading" class="text-center py-12">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-brand-primary"></div>
            <p class="mt-4 text-brand-gray">Loading articles...</p>
          </div>
          
          <div v-else-if="articles.length === 0" class="text-center py-12">
            <div class="max-w-md mx-auto">
              <div class="mb-6">
                <svg class="mx-auto h-12 w-12 text-brand-gray" fill="none" viewBox="0 0 24 24" stroke="currentColor">
                  <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z" />
                </svg>
              </div>
              <h3 class="text-lg font-medium text-brand-dark mb-2">No articles found</h3>
              <p class="text-brand-gray mb-6">
                Check back soon for new content, or explore other sections of the site.
              </p>
              <div class="flex flex-col sm:flex-row gap-3 justify-center">
                <NuxtLink 
                  to="/" 
                  class="bg-brand-primary text-brand-dark px-6 py-2 rounded-lg hover:bg-brand-light transition font-semibold"
                >
                  Back to Home
                </NuxtLink>
                <NuxtLink 
                  to="/about" 
                  class="bg-gray-200 text-brand-dark px-6 py-2 rounded-lg hover:bg-gray-300 transition font-semibold"
                >
                  Learn About Me
                </NuxtLink>
              </div>
            </div>
          </div>
          
          <div v-else>
            <!-- Articles Grid -->
            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
              <article 
                v-for="article in articles" 
                :key="article.id" 
                class="bg-brand-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-all duration-300 border border-brand-gray group"
                @mouseenter="startImageSlideshow(article.id)"
                @mouseleave="stopImageSlideshow(article.id)"
              >
                <!-- Article Image -->
                <div v-if="getArticleImages(article).length > 0" class="relative h-48 overflow-hidden">
                  <div 
                    v-for="(image, index) in getArticleImages(article)" 
                    :key="index"
                    class="absolute inset-0 transition-opacity duration-500"
                    :class="{ 'opacity-100': (currentImageIndex[article.id] || 0) === index, 'opacity-0': (currentImageIndex[article.id] || 0) !== index }"
                  >
                    <img 
                      :src="getImageUrl(image)" 
                      :alt="`${article.title} - Image ${index + 1}`"
                      class="w-full h-full object-cover"
                    >
                  </div>
                  <!-- Image indicator dots (for multiple images) -->
                  <div v-if="getArticleImages(article).length > 1" class="absolute bottom-2 left-1/2 transform -translate-x-1/2 flex space-x-1">
                    <div 
                      v-for="(image, index) in getArticleImages(article)" 
                      :key="index"
                      class="w-2 h-2 rounded-full transition-all duration-300"
                      :class="(currentImageIndex[article.id] || 0) === index ? 'bg-white' : 'bg-white/50'"
                    ></div>
                  </div>
                </div>

                <div class="p-6">
                  <h2 class="text-xl font-semibold text-brand-dark mb-3 line-clamp-2">
                    {{ article.title }}
                  </h2>
                  <p class="text-brand-dark mb-4 line-clamp-3">
                    {{ article.content.substring(0, 150) }}{{ article.content.length > 150 ? '...' : '' }}
                  </p>
                  <div class="flex items-center justify-between">
                    <span class="text-sm text-brand-gray">
                      {{ formatDate(article.created_at) }}
                    </span>
                    <NuxtLink 
                      :to="`/articles/${article.id}`"
                      class="inline-flex items-center text-brand-primary hover:text-brand-light font-medium text-sm transition"
                    >
                      Read More
                      <svg class="ml-1 w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5l7 7-7 7"></path>
                      </svg>
                    </NuxtLink>
                  </div>
                </div>
              </article>
            </div>

            <!-- Load More / Pagination could go here -->
            <div class="text-center mt-12">
              <p class="text-brand-gray">
                Showing {{ articles.length }} article{{ articles.length === 1 ? '' : 's' }}
              </p>
            </div>
          </div>
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'

// Fetch articles data on server-side for better initial load
const { data: articlesData } = await useFetch('/api/articles')

const articles = ref(articlesData.value || [])
const loading = ref(false)
const currentImageIndex = ref({})
const slideshowIntervals = ref({})

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
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

const getArticleImages = (article) => {
  if (!article?.images) return []
  try {
    return typeof article.images === 'string' 
      ? JSON.parse(article.images) 
      : article.images
  } catch {
    return []
  }
}

const initializeImageIndexes = () => {
  if (!articles.value) return
  articles.value.forEach(article => {
    if (article && article.id && !currentImageIndex.value[article.id]) {
      currentImageIndex.value[article.id] = 0
    }
  })
}

const startImageSlideshow = (articleId) => {
  if (!articles.value || !articleId) return
  const article = articles.value.find(a => a && a.id === articleId)
  if (!article) return
  
  const images = getArticleImages(article)
  if (images.length <= 1) return
  
  // Initialize index if not exists
  if (currentImageIndex.value[articleId] === undefined) {
    currentImageIndex.value[articleId] = 0
  }
  
  // Start slideshow
  slideshowIntervals.value[articleId] = setInterval(() => {
    currentImageIndex.value[articleId] = (currentImageIndex.value[articleId] + 1) % images.length
  }, 1500) // Change image every 1.5 seconds
}

const stopImageSlideshow = (articleId) => {
  if (!articleId) return
  if (slideshowIntervals.value[articleId]) {
    clearInterval(slideshowIntervals.value[articleId])
    delete slideshowIntervals.value[articleId]
    // Reset to first image
    if (currentImageIndex.value[articleId] !== undefined) {
      currentImageIndex.value[articleId] = 0
    }
  }
}

onMounted(() => {
  // Initialize image indexes for client-side interactivity
  initializeImageIndexes()
})

onUnmounted(() => {
  // Clear all slideshow intervals
  Object.values(slideshowIntervals.value).forEach(interval => clearInterval(interval))
})

useSeoMeta({
  title: 'Articles - Personal Web | John Doe',
  description: 'Read the latest articles and insights on web development, programming, and technology by John Doe.'
})
</script>
