<template>
  <div>
    <Header />
    <main class="min-h-screen bg-gray-100">
      <!-- Hero Section -->
      <section class="bg-gradient-to-br from-brand-light to-brand-primary py-16">
        <div class="container mx-auto px-4 text-center">
          <h1 class="text-4xl md:text-5xl font-bold text-brand-dark mb-6">
            Ready to Transform Your Business?
          </h1>
          <p class="text-xl text-brand-dark mb-8 max-w-3xl mx-auto">
            Let's discuss how Kuro Haeksagon can help streamline your operations, 
            integrate your systems, and unlock the power of AI for your organization.
          </p>
        </div>
      </section>

      <!-- Contact Content -->
      <section class="py-16">
        <div class="container mx-auto px-4">
          <div class="max-w-6xl mx-auto">
            <!-- Why Choose Us Section -->
            <div class="grid md:grid-cols-3 gap-8 mb-16">
              <div class="text-center">
                <div class="w-16 h-16 bg-brand-primary rounded-lg flex items-center justify-center mx-auto mb-4">
                  <svg class="w-8 h-8 text-brand-dark" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
                  </svg>
                </div>
                <h3 class="text-lg font-semibold text-brand-dark mb-2">Expert Integration</h3>
                <p class="text-brand-gray text-sm">Seamless system connections for enhanced productivity</p>
              </div>
              <div class="text-center">
                <div class="w-16 h-16 bg-brand-primary rounded-lg flex items-center justify-center mx-auto mb-4">
                  <svg class="w-8 h-8 text-brand-dark" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zM9 17H7v-7h2v7zm4 0h-2V7h2v10zm4 0h-2v-4h2v4z"/>
                  </svg>
                </div>
                <h3 class="text-lg font-semibold text-brand-dark mb-2">ERP Excellence</h3>
                <p class="text-brand-gray text-sm">Specialized Odoo solutions for business optimization</p>
              </div>
              <div class="text-center">
                <div class="w-16 h-16 bg-brand-primary rounded-lg flex items-center justify-center mx-auto mb-4">
                  <svg class="w-8 h-8 text-brand-dark" fill="currentColor" viewBox="0 0 24 24">
                    <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
                  </svg>
                </div>
                <h3 class="text-lg font-semibold text-brand-dark mb-2">AI Innovation</h3>
                <p class="text-brand-gray text-sm">Cutting-edge AI solutions for competitive advantage</p>
              </div>
            </div>

            <div class="grid md:grid-cols-2 gap-12">
              <!-- Contact Form -->
              <div class="bg-brand-white p-8 rounded-lg shadow-lg border border-brand-gray">
                <h2 class="text-2xl font-semibold text-brand-dark mb-6">Start Your Project</h2>
            <form @submit.prevent="submitForm" class="space-y-6">
              <div>
                <label for="name" class="block text-sm font-medium text-brand-dark mb-2">Name *</label>
                <input 
                  id="name" 
                  v-model="form.name" 
                  type="text" 
                  required
                  class="w-full px-4 py-3 border border-brand-gray rounded-md focus:outline-none focus:ring-2 focus:ring-brand-primary focus:border-brand-primary"
                  placeholder="Your full name"
                />
              </div>
              
              <div>
                <label for="email" class="block text-sm font-medium text-brand-dark mb-2">Email *</label>
                <input 
                  id="email" 
                  v-model="form.email" 
                  type="email" 
                  required
                  class="w-full px-4 py-3 border border-brand-gray rounded-md focus:outline-none focus:ring-2 focus:ring-brand-primary focus:border-brand-primary"
                  placeholder="your.email@example.com"
                />
              </div>
              
              <div>
                <label for="subject" class="block text-sm font-medium text-brand-dark mb-2">Subject *</label>
                <input 
                  id="subject" 
                  v-model="form.subject" 
                  type="text" 
                  required
                  class="w-full px-4 py-3 border border-brand-gray rounded-md focus:outline-none focus:ring-2 focus:ring-brand-primary focus:border-brand-primary"
                  placeholder="What's this about?"
                />
              </div>
              
              <div>
                <label for="message" class="block text-sm font-medium text-brand-dark mb-2">Message *</label>
                <textarea 
                  id="message" 
                  v-model="form.message" 
                  rows="5" 
                  required
                  class="w-full px-4 py-3 border border-brand-gray rounded-md focus:outline-none focus:ring-2 focus:ring-brand-primary focus:border-brand-primary"
                  placeholder="Tell us about your project requirements, goals, or challenges..."
                ></textarea>
              </div>
              
              <button 
                type="submit" 
                :disabled="isSubmitting"
                class="w-full bg-brand-primary text-brand-dark py-3 px-6 rounded-md hover:bg-brand-light focus:outline-none focus:ring-2 focus:ring-brand-primary focus:ring-offset-2 transition duration-150 ease-in-out disabled:opacity-50 font-semibold"
              >
                {{ isSubmitting ? 'Sending...' : 'Send Message' }}
              </button>
              
              <p v-if="submitMessage" :class="submitSuccess ? 'text-green-600' : 'text-red-600'" class="text-sm text-center">
                {{ submitMessage }}
              </p>
            </form>
          </div>

          <!-- Contact Information -->
          <div class="space-y-8">
            <div class="bg-brand-white p-8 rounded-lg shadow-lg border border-brand-gray">
              <h2 class="text-2xl font-semibold text-brand-dark mb-6">Get In Touch</h2>
              <div class="space-y-4">
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-brand-light rounded-full flex items-center justify-center mr-4">
                    <svg class="w-4 h-4 text-brand-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 8l7.89 4.26a2 2 0 002.22 0L21 8M5 19h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v10a2 2 0 002 2z" />
                    </svg>
                  </div>
                  <div>
                    <p class="text-brand-gray">Email</p>
                    <p class="text-brand-dark font-medium">hello@kurohaeksagon.com</p>
                  </div>
                </div>
                
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-brand-light rounded-full flex items-center justify-center mr-4">
                    <svg class="w-4 h-4 text-brand-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 11-18 0 9 9 0 0118 0z" />
                    </svg>
                  </div>
                  <div>
                    <p class="text-brand-gray">Response Time</p>
                    <p class="text-brand-dark font-medium">Within 24 hours</p>
                  </div>
                </div>
                
                <div class="flex items-center">
                  <div class="w-8 h-8 bg-brand-light rounded-full flex items-center justify-center mr-4">
                    <svg class="w-4 h-4 text-brand-primary" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.657 16.657L13.414 20.9a1.998 1.998 0 01-2.827 0l-4.244-4.243a8 8 0 1111.314 0z" />
                      <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 11a3 3 0 11-6 0 3 3 0 016 0z" />
                    </svg>
                  </div>
                  <div>
                    <p class="text-brand-gray">Services</p>
                    <p class="text-brand-dark font-medium">AI • ERP • Integration</p>
                  </div>
                </div>
              </div>
            </div>

            <div class="bg-brand-white p-8 rounded-lg shadow-lg border border-brand-gray">
              <h2 class="text-2xl font-semibold text-brand-dark mb-6">Our Expertise</h2>
              <div class="space-y-4">
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-brand-primary rounded-full"></div>
                  <span class="text-brand-dark">Software Integration & API Development</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-brand-primary rounded-full"></div>
                  <span class="text-brand-dark">Odoo ERP Implementation & Customization</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-brand-primary rounded-full"></div>
                  <span class="text-brand-dark">AI & Machine Learning Solutions</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-brand-primary rounded-full"></div>
                  <span class="text-brand-dark">Cloud Platform Integration (GCP, AWS)</span>
                </div>
                <div class="flex items-center space-x-3">
                  <div class="w-2 h-2 bg-brand-primary rounded-full"></div>
                  <span class="text-brand-dark">N8N Workflow Automation</span>
                </div>
              </div>
            </div>

            <div class="bg-brand-light p-8 rounded-lg border border-brand-primary">
              <h3 class="text-lg font-semibold text-brand-dark mb-4">Why Choose Kuro Haeksagon?</h3>
              <p class="text-brand-dark mb-4">
                We deliver innovative technology solutions that drive real business value. Our team combines deep technical expertise with a practical understanding of business challenges.
              </p>
              <ul class="space-y-2 text-brand-dark text-sm">
                <li>• Proven track record in enterprise solutions</li>
                <li>• Dedicated support and maintenance</li>
                <li>• Rapid deployment and implementation</li>
                <li>• Scalable and future-ready architectures</li>
              </ul>
            </div>
          </div>
            </div>
          </div>
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

<script setup>
import { ref } from 'vue'

const form = ref({
  name: '',
  email: '',
  subject: '',
  message: ''
})

const isSubmitting = ref(false)
const submitMessage = ref('')
const submitSuccess = ref(false)

const submitForm = async () => {
  isSubmitting.value = true
  submitMessage.value = ''
  
  try {
    const config = useRuntimeConfig()
    const response = await fetch(`${config.public.apiBase}/api/contact`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(form.value)
    })

    const data = await response.json()

    if (response.ok) {
      submitMessage.value = data.message || 'Thank you for your message! Our team will get back to you within 24 hours.'
      submitSuccess.value = true
      
      // Reset form
      form.value = {
        name: '',
        email: '',
        subject: '',
        message: ''
      }
    } else {
      throw new Error(data.error || 'Failed to send message')
    }
  } catch (error) {
    submitMessage.value = error.message || 'There was an error sending your message. Please try again.'
    submitSuccess.value = false
  } finally {
    isSubmitting.value = false
  }
}

useSeoMeta({
  title: 'Contact Us - Kuro Haeksagon | AI & Software Solutions',
  description: 'Get in touch with Kuro Haeksagon for AI solutions, ERP implementation, and software development services. Contact our expert team today.'
})
</script>