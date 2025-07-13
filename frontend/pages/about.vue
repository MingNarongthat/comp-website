<template>
  <div>
    <Header />
    <main class="min-h-screen bg-gray-100">
      <!-- Hero Section -->
      <section class="bg-gradient-to-br from-brand-light to-brand-primary py-20">
        <div class="container mx-auto px-4 text-center">
          <h1 class="text-5xl md:text-6xl font-bold text-brand-dark mb-6">
            About Kuro Haeksagon
          </h1>
          <p class="text-xl text-brand-dark mb-8 max-w-3xl mx-auto">
            Transforming businesses through intelligent software solutions and innovative AI technologies
          </p>
        </div>
      </section>

      <!-- Company Overview -->
      <section class="py-16 px-4 bg-brand-white">
        <div class="container mx-auto max-w-6xl">
          <div class="grid md:grid-cols-2 gap-12 items-center">
            <div>
              <h2 class="text-3xl font-bold text-brand-dark mb-6">Who We Are</h2>
              <p class="text-brand-dark leading-relaxed mb-6">
                Kuro Haeksagon is a forward-thinking technology solutions provider committed to driving business transformation through intelligent software systems. Our core mission is to help organizations optimize their operations, unlock data-driven insights, and accelerate growth with practical, scalable tools.
              </p>
              <p class="text-brand-dark leading-relaxed mb-6">
                Founded with a vision to bridge the gap between complex technology and business needs, we specialize in creating seamless integrations, automated workflows, and AI-powered solutions that deliver real value to our clients.
              </p>
              <div class="flex flex-wrap gap-3">
                <span class="bg-brand-light text-brand-dark px-4 py-2 rounded-full text-sm font-medium border border-brand-primary">
                  Innovation
                </span>
                <span class="bg-brand-light text-brand-dark px-4 py-2 rounded-full text-sm font-medium border border-brand-primary">
                  Excellence
                </span>
                <span class="bg-brand-light text-brand-dark px-4 py-2 rounded-full text-sm font-medium border border-brand-primary">
                  Reliability
                </span>
                <span class="bg-brand-light text-brand-dark px-4 py-2 rounded-full text-sm font-medium border border-brand-primary">
                  Partnership
                </span>
              </div>
            </div>
            <div class="h-96 rounded-lg flex items-center justify-center border-2 border-brand-primary bg-brand-white">
              <img 
                src="~/assets/features.png" 
                alt="Kuro Haeksagon Solutions" 
                class="max-h-full max-w-full object-contain p-4"
              >
            </div>
          </div>
        </div>
      </section>

      <!-- Our Services -->
      <section class="py-16 px-4 bg-gray-100">
        <div class="container mx-auto max-w-6xl">
          <div class="text-center mb-12">
            <h2 class="text-3xl font-bold text-brand-dark mb-4">Our Expertise</h2>
            <p class="text-lg text-brand-gray max-w-2xl mx-auto">
              We specialize in three key areas that drive digital transformation and business growth
            </p>
          </div>
          
          <div class="grid md:grid-cols-3 gap-8">
            <!-- Software Integration -->
            <div class="bg-brand-white p-8 rounded-lg shadow-md border border-brand-gray hover:shadow-lg transition-shadow">
              <div class="w-16 h-16 bg-brand-primary rounded-lg flex items-center justify-center mb-6">
                <svg class="w-8 h-8 text-brand-dark" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2L2 7l10 5 10-5-10-5zM2 17l10 5 10-5M2 12l10 5 10-5"/>
                </svg>
              </div>
              <h3 class="text-xl font-bold text-brand-dark mb-4">Software Integration</h3>
              <p class="text-brand-dark leading-relaxed mb-4">
                Seamlessly connecting systems to ensure smooth data flow and enhanced productivity. We break down silos and create unified workflows.
              </p>
              <ul class="space-y-2 text-brand-dark text-sm">
                <li>• API Development & Integration</li>
                <li>• Data Migration & Synchronization</li>
                <li>• Legacy System Modernization</li>
                <li>• Workflow Automation</li>
              </ul>
            </div>

            <!-- ERP Modules -->
            <div class="bg-brand-white p-8 rounded-lg shadow-md border border-brand-gray hover:shadow-lg transition-shadow">
              <div class="w-16 h-16 bg-brand-primary rounded-lg flex items-center justify-center mb-6">
                <svg class="w-8 h-8 text-brand-dark" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M19 3H5c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zM9 17H7v-7h2v7zm4 0h-2V7h2v10zm4 0h-2v-4h2v4z"/>
                </svg>
              </div>
              <h3 class="text-xl font-bold text-brand-dark mb-4">ERP Modules</h3>
              <p class="text-brand-dark leading-relaxed mb-4">
                Streamlining enterprise resource planning for better control, planning, and operational efficiency with specialized Odoo solutions.
              </p>
              <ul class="space-y-2 text-brand-dark text-sm">
                <li>• Odoo Implementation & Customization</li>
                <li>• Business Process Optimization</li>
                <li>• Financial Management Systems</li>
                <li>• Inventory & Supply Chain</li>
              </ul>
            </div>

            <!-- AI Data Solutions -->
            <div class="bg-brand-white p-8 rounded-lg shadow-md border border-brand-gray hover:shadow-lg transition-shadow">
              <div class="w-16 h-16 bg-brand-primary rounded-lg flex items-center justify-center mb-6">
                <svg class="w-8 h-8 text-brand-dark" fill="currentColor" viewBox="0 0 24 24">
                  <path d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
                </svg>
              </div>
              <h3 class="text-xl font-bold text-brand-dark mb-4">AI Data Solutions</h3>
              <p class="text-brand-dark leading-relaxed mb-4">
                Harnessing artificial intelligence to turn data into actionable strategies that power innovation and competitive advantage.
              </p>
              <ul class="space-y-2 text-brand-dark text-sm">
                <li>• Machine Learning Implementation</li>
                <li>• Computer Vision Applications</li>
                <li>• Generative AI Integration</li>
                <li>• Data Analytics & Insights</li>
              </ul>
            </div>
          </div>
        </div>
      </section>

      <!-- Technology Stack -->
      <section class="py-16 px-4 bg-brand-white">
        <div class="container mx-auto max-w-6xl">
          <div class="text-center mb-12">
            <h2 class="text-3xl font-bold text-brand-dark mb-4">Technologies We Master</h2>
            <p class="text-lg text-brand-gray max-w-2xl mx-auto">
              Our technical expertise spans modern development frameworks, cloud platforms, and cutting-edge AI technologies
            </p>
          </div>
          
          <div class="grid md:grid-cols-4 gap-6">
            <!-- Cloud Platforms -->
            <div class="text-center">
              <h3 class="text-lg font-semibold text-brand-dark mb-4">Cloud Platforms</h3>
              <div class="space-y-3">
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Google Cloud Platform
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Amazon Web Services
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Cloud Solutions
                </div>
              </div>
            </div>

            <!-- Development -->
            <div class="text-center">
              <h3 class="text-lg font-semibold text-brand-dark mb-4">Development</h3>
              <div class="space-y-3">
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Software Development
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  N8N Automation
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Odoo
                </div>
              </div>
            </div>

            <!-- AI Technologies -->
            <div class="text-center">
              <h3 class="text-lg font-semibold text-brand-dark mb-4">AI Technologies</h3>
              <div class="space-y-3">
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Artificial Intelligence
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Generative AI
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Computer Vision
                </div>
              </div>
            </div>

            <!-- Industries -->
            <div class="text-center">
              <h3 class="text-lg font-semibold text-brand-dark mb-4">Industries</h3>
              <div class="space-y-3">
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Manufacturing
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  E-commerce
                </div>
                <div class="bg-brand-light text-brand-dark px-3 py-2 rounded-lg text-sm font-medium border border-brand-primary">
                  Healthcare
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>

      <!-- Our Approach -->
      <section class="py-16 px-4 bg-gray-100">
        <div class="container mx-auto max-w-6xl">
          <div class="text-center mb-12">
            <h2 class="text-3xl font-bold text-brand-dark mb-4">Our Approach</h2>
            <p class="text-lg text-brand-gray max-w-2xl mx-auto">
              We follow a proven methodology to ensure successful project delivery and long-term client satisfaction
            </p>
          </div>
          
          <div class="grid md:grid-cols-2 lg:grid-cols-4 gap-8">
            <div class="text-center">
              <div class="w-16 h-16 bg-brand-primary rounded-full flex items-center justify-center mx-auto mb-4">
                <span class="text-2xl font-bold text-brand-dark">1</span>
              </div>
              <h3 class="text-lg font-semibold text-brand-dark mb-3">Discovery</h3>
              <p class="text-brand-dark text-sm">
                We analyze your current systems, understand your challenges, and identify opportunities for improvement.
              </p>
            </div>
            
            <div class="text-center">
              <div class="w-16 h-16 bg-brand-primary rounded-full flex items-center justify-center mx-auto mb-4">
                <span class="text-2xl font-bold text-brand-dark">2</span>
              </div>
              <h3 class="text-lg font-semibold text-brand-dark mb-3">Strategy</h3>
              <p class="text-brand-dark text-sm">
                We develop a comprehensive roadmap with clear milestones, timelines, and success metrics.
              </p>
            </div>
            
            <div class="text-center">
              <div class="w-16 h-16 bg-brand-primary rounded-full flex items-center justify-center mx-auto mb-4">
                <span class="text-2xl font-bold text-brand-dark">3</span>
              </div>
              <h3 class="text-lg font-semibold text-brand-dark mb-3">Implementation</h3>
              <p class="text-brand-dark text-sm">
                Our expert team executes the solution with regular updates and transparent communication.
              </p>
            </div>
            
            <div class="text-center">
              <div class="w-16 h-16 bg-brand-primary rounded-full flex items-center justify-center mx-auto mb-4">
                <span class="text-2xl font-bold text-brand-dark">4</span>
              </div>
              <h3 class="text-lg font-semibold text-brand-dark mb-3">Support</h3>
              <p class="text-brand-dark text-sm">
                We provide ongoing maintenance, training, and optimization to ensure long-term success.
              </p>
            </div>
          </div>
        </div>
      </section>

      <!-- Call to Action -->
      <section class="py-16 px-4 bg-brand-dark text-brand-white">
        <div class="container mx-auto max-w-4xl text-center">
          <h2 class="text-3xl font-bold mb-6">Ready to Transform Your Business?</h2>
          <p class="text-xl text-brand-gray mb-8 max-w-2xl mx-auto">
            Let's discuss how Kuro Haeksagon can help streamline your operations, integrate your systems, and unlock the power of AI for your organization.
          </p>
          <div class="flex flex-col sm:flex-row gap-4 justify-center">
            <NuxtLink 
              to="/contact" 
              class="bg-brand-primary text-brand-dark px-8 py-3 rounded-lg hover:bg-brand-light transition font-semibold"
            >
              Start Your Project
            </NuxtLink>
            <a 
              href="mailto:hello@kurohaeksagon.com" 
              class="bg-gray-800 text-brand-primary px-8 py-3 rounded-lg hover:bg-gray-700 transition font-semibold border border-brand-primary"
            >
              Email Us
            </a>
          </div>
          
          <!-- Contact Info -->
          <div class="mt-12 grid md:grid-cols-3 gap-8 text-center">
            <div>
              <h3 class="font-semibold text-brand-primary mb-2">Email</h3>
              <p class="text-brand-gray">hello@kurohaeksagon.com</p>
            </div>
            <div>
              <h3 class="font-semibold text-brand-primary mb-2">Services</h3>
              <p class="text-brand-gray">AI Solutions • ERP • Integration</p>
            </div>
            <div>
              <h3 class="font-semibold text-brand-primary mb-2">Expertise</h3>
              <p class="text-brand-gray">Cloud • Automation • Innovation</p>
            </div>
          </div>
        </div>
      </section>
    </main>
    <Footer />
  </div>
</template>

<script setup>
useSeoMeta({
  title: 'About - Kuro Haeksagon | AI & Software Solutions',
  description: 'Learn about Kuro Haeksagon, a technology solutions provider specializing in software integration, ERP modules, and AI data solutions for business transformation.'
})
</script>