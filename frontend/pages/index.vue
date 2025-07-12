<template>
  <div>
    <Header />
    <main class="min-h-screen">
      <!-- Hero Section -->
      <section class="bg-gradient-to-br from-brand-light to-brand-primary py-20 px-4">
        <div class="container mx-auto text-center">
          <h1 class="text-5xl md:text-6xl font-bold text-brand-dark mb-6">
            {{ heroTitle }}
          </h1>
          <p class="text-xl text-brand-dark mb-8 max-w-2xl mx-auto">
            {{ heroDescription }}
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <NuxtLink 
              to="/about" 
              class="bg-brand-dark text-brand-primary px-8 py-3 rounded-lg hover:bg-gray-800 transition font-semibold"
            >
              {{ aboutButtonText }}
            </NuxtLink>
            <NuxtLink 
              to="/articles" 
              class="bg-brand-white text-brand-dark px-8 py-3 rounded-lg border-2 border-brand-dark hover:bg-gray-100 transition font-semibold"
            >
              {{ articlesButtonText }}
            </NuxtLink>
          </div>
        </div>
      </section>

      <!-- Introduction Section -->
      <section class="py-16 px-4 bg-brand-white">
        <div class="container mx-auto max-w-4xl">
          <div class="text-center mb-12">
            <h2 class="text-3xl font-bold text-brand-dark mb-4">{{ introTitle }}</h2>
            <p class="text-lg text-brand-gray">{{ introSubtitle }}</p>
          </div>
          <div class="grid md:grid-cols-2 gap-12 items-center">
            <div>
              <p class="text-brand-dark leading-relaxed mb-6">
                {{ introParagraph1 }}
              </p>
              <p class="text-brand-dark leading-relaxed mb-6">
                {{ introParagraph2 }}
              </p>
              <div class="flex flex-wrap gap-3">
                <span v-for="skill in skills" :key="skill" 
                      class="bg-brand-light text-brand-dark px-3 py-1 rounded-full text-sm font-medium border border-brand-primary">
                  {{ skill }}
                </span>
              </div>
            </div>
            <div class="bg-gray-200 h-80 rounded-lg flex items-center justify-center border-2 border-brand-primary">
              <span class="text-brand-gray text-lg">{{ photoPlaceholder }}</span>
            </div>
          </div>
        </div>
      </section>

      <!-- Latest Articles Section -->
      <section class="py-16 px-4 bg-gray-100">
        <div class="container mx-auto max-w-6xl">
          <div class="text-center mb-12">
            <h2 class="text-3xl font-bold text-brand-dark mb-4">Latest Articles</h2>
            <p class="text-lg text-brand-gray">Insights, tutorials, and thoughts on web development</p>
          </div>
          <div v-if="loading" class="text-center py-8">
            <div class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-brand-primary"></div>
            <p class="mt-2 text-brand-gray">Loading articles...</p>
          </div>
          <div v-else-if="articles.length > 0" class="grid md:grid-cols-2 lg:grid-cols-3 gap-8 mb-8">
            <article 
              v-for="article in articles.slice(0, 3)" 
              :key="article.id"
              class="bg-brand-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition border border-brand-gray"
            >
              <div class="p-6">
                <h3 class="text-xl font-semibold text-brand-dark mb-3 line-clamp-2">
                  {{ article.title }}
                </h3>
                <p class="text-brand-dark mb-4 line-clamp-3">
                  {{ article.content.substring(0, 150) }}{{ article.content.length > 150 ? '...' : '' }}
                </p>
                <div class="flex items-center justify-between">
                  <span class="text-sm text-brand-gray">
                    {{ formatDate(article.created_at) }}
                  </span>
                  <NuxtLink 
                    :to="`/articles/${article.id}`"
                    class="text-brand-primary hover:text-brand-light font-medium text-sm"
                  >
                    Read More →
                  </NuxtLink>
                </div>
              </div>
            </article>
          </div>
          <div v-else class="text-center py-8">
            <p class="text-brand-gray">No articles available yet. Check back soon!</p>
          </div>
          <div class="text-center">
            <NuxtLink 
              to="/articles" 
              class="bg-brand-primary text-brand-dark px-6 py-3 rounded-lg hover:bg-brand-light transition font-semibold"
            >
              View All Articles
            </NuxtLink>
          </div>
        </div>
      </section>

      <!-- Contact Preview Section -->
      <section class="py-16 px-4 bg-brand-dark text-brand-white">
        <div class="container mx-auto max-w-4xl text-center">
          <h2 class="text-3xl font-bold mb-4">Let's Work Together</h2>
          <p class="text-xl text-brand-gray mb-8 max-w-2xl mx-auto">
            Have a project in mind or want to collaborate? I'd love to hear from you and 
            discuss how we can bring your ideas to life.
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <NuxtLink 
              to="/contact" 
              class="bg-brand-primary text-brand-dark px-8 py-3 rounded-lg hover:bg-brand-light transition font-semibold"
            >
              Get In Touch
            </NuxtLink>
            <a 
              href="mailto:john@example.com" 
              class="bg-gray-800 text-brand-primary px-8 py-3 rounded-lg hover:bg-gray-700 transition font-semibold border border-brand-primary"
            >
              Send Email
            </a>
          </div>
          <div class="mt-8 flex justify-center space-x-6">
            <a href="https://github.com/johndoe" target="_blank" class="text-brand-gray hover:text-brand-primary transition">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M12 0c-6.626 0-12 5.373-12 12 0 5.302 3.438 9.8 8.207 11.387.599.111.793-.261.793-.577v-2.234c-3.338.726-4.033-1.416-4.033-1.416-.546-1.387-1.333-1.756-1.333-1.756-1.089-.745.083-.729.083-.729 1.205.084 1.839 1.237 1.839 1.237 1.07 1.834 2.807 1.304 3.492.997.107-.775.418-1.305.762-1.604-2.665-.305-5.467-1.334-5.467-5.931 0-1.311.469-2.381 1.236-3.221-.124-.303-.535-1.524.117-3.176 0 0 1.008-.322 3.301 1.23.957-.266 1.983-.399 3.003-.404 1.02.005 2.047.138 3.006.404 2.291-1.552 3.297-1.23 3.297-1.23.653 1.653.242 2.874.118 3.176.77.84 1.235 1.911 1.235 3.221 0 4.609-2.807 5.624-5.479 5.921.43.372.823 1.102.823 2.222v3.293c0 .319.192.694.801.576 4.765-1.589 8.199-6.086 8.199-11.386 0-6.627-5.373-12-12-12z"/>
              </svg>
            </a>
            <a href="https://linkedin.com/in/johndoe" target="_blank" class="text-brand-gray hover:text-brand-primary transition">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M20.447 20.452h-3.554v-5.569c0-1.328-.027-3.037-1.852-3.037-1.853 0-2.136 1.445-2.136 2.939v5.667H9.351V9h3.414v1.561h.046c.477-.9 1.637-1.85 3.37-1.85 3.601 0 4.267 2.37 4.267 5.455v6.286zM5.337 7.433c-1.144 0-2.063-.926-2.063-2.065 0-1.138.92-2.063 2.063-2.063 1.14 0 2.064.925 2.064 2.063 0 1.139-.925 2.065-2.064 2.065zm1.782 13.019H3.555V9h3.564v11.452zM22.225 0H1.771C.792 0 0 .774 0 1.729v20.542C0 23.227.792 24 1.771 24h20.451C23.2 24 24 23.227 24 22.271V1.729C24 .774 23.2 0 22.222 0h.003z"/>
              </svg>
            </a>
            <a href="https://twitter.com/johndoe" target="_blank" class="text-brand-gray hover:text-brand-primary transition">
              <svg class="w-6 h-6" fill="currentColor" viewBox="0 0 24 24">
                <path d="M23.953 4.57a10 10 0 01-2.825.775 4.958 4.958 0 002.163-2.723c-.951.555-2.005.959-3.127 1.184a4.92 4.92 0 00-8.384 4.482C7.69 8.095 4.067 6.13 1.64 3.162a4.822 4.822 0 00-.666 2.475c0 1.71.87 3.213 2.188 4.096a4.904 4.904 0 01-2.228-.616v.06a4.923 4.923 0 003.946 4.827 4.996 4.996 0 01-2.212.085 4.936 4.936 0 004.604 3.417 9.867 9.867 0 01-6.102 2.105c-.39 0-.779-.023-1.17-.067a13.995 13.995 0 007.557 2.209c9.053 0 13.998-7.496 13.998-13.985 0-.21 0-.42-.015-.63A9.935 9.935 0 0024 4.59z"/>
              </svg>
            </a>
          </div>
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useArticleStore } from '@/stores/article'

const articleStore = useArticleStore()
const articles = ref([])
const loading = ref(true)
const { getContent } = useContent()

// CMS Content variables
const heroTitle = ref('Welcome to My Personal Web')
const heroDescription = ref('A passionate full-stack developer sharing insights, experiences, and knowledge through modern web technologies and innovative solutions.')
const aboutButtonText = ref('Learn More About Me')
const articlesButtonText = ref('Read My Articles')
const introTitle = ref('Hello, I\'m John Doe')
const introSubtitle = ref('Full-Stack Developer & Tech Enthusiast')
const introParagraph1 = ref('With over 5 years of experience in modern web development, I specialize in creating scalable applications using cutting-edge technologies like Go, Vue.js, and cloud platforms.')
const introParagraph2 = ref('I\'m passionate about sharing knowledge, building innovative solutions, and helping fellow developers grow in their careers through practical insights and real-world examples.')
const photoPlaceholder = ref('Developer Photo')
const skills = ref(['Go', 'Vue.js', 'PostgreSQL', 'Docker'])

const formatDate = (dateString) => {
  const date = new Date(dateString)
  return date.toLocaleDateString('en-US', { 
    year: 'numeric', 
    month: 'long', 
    day: 'numeric' 
  })
}

const fetchLatestArticles = async () => {
  try {
    // For the home page, we'll fetch public articles without auth
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/articles`)
    if (response.ok) {
      const data = await response.json()
      articles.value = data || []
    }
  } catch (error) {
    console.log('Could not fetch articles for home page preview')
    articles.value = []
  } finally {
    loading.value = false
  }
}

const loadContent = async () => {
  // Load CMS content with fallbacks
  heroTitle.value = await getContent('home.hero.title', heroTitle.value)
  heroDescription.value = await getContent('home.hero.description', heroDescription.value)
  aboutButtonText.value = await getContent('home.hero.about_button', aboutButtonText.value)
  articlesButtonText.value = await getContent('home.hero.articles_button', articlesButtonText.value)
  
  introTitle.value = await getContent('home.intro.title', introTitle.value)
  introSubtitle.value = await getContent('home.intro.subtitle', introSubtitle.value)
  introParagraph1.value = await getContent('home.intro.paragraph1', introParagraph1.value)
  introParagraph2.value = await getContent('home.intro.paragraph2', introParagraph2.value)
  photoPlaceholder.value = await getContent('home.intro.photo_placeholder', photoPlaceholder.value)
  
  // Load skills as JSON or comma-separated string
  const skillsContent = await getContent('home.intro.skills', '')
  if (skillsContent) {
    try {
      // Try to parse as JSON first
      skills.value = JSON.parse(skillsContent)
    } catch {
      // Fall back to comma-separated string
      skills.value = skillsContent.split(',').map(s => s.trim()).filter(s => s)
    }
  }
}

onMounted(() => {
  fetchLatestArticles()
  loadContent()
})

useSeoMeta({
  title: 'Personal Web - John Doe | Full-Stack Developer',
  description: 'Welcome to John Doe\'s personal website. A passionate full-stack developer sharing insights, tutorials, and experiences in modern web development.'
})
</script> 