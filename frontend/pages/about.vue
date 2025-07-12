<template>
  <div>
    <Header />
    <main class="container mx-auto px-4 py-12 bg-gray-100 min-h-screen">
      <div class="max-w-4xl mx-auto">
        <h1 class="text-4xl font-bold text-brand-dark mb-8 text-center">About Me</h1>
        
        <div class="grid md:grid-cols-2 gap-12 items-center mb-16">
          <div class="space-y-6">
            <h2 class="text-2xl font-semibold text-brand-dark">Hello, I'm John Doe</h2>
            <p class="text-brand-dark leading-relaxed">
              I'm a passionate full-stack developer with over 5 years of experience creating 
              modern web applications. I specialize in building scalable, user-friendly 
              solutions using cutting-edge technologies.
            </p>
            <p class="text-brand-dark leading-relaxed">
              When I'm not coding, you can find me exploring new technologies, contributing 
              to open-source projects, or sharing my knowledge through blog posts and tutorials.
            </p>
          </div>
          <div class="bg-gray-200 h-80 rounded-lg flex items-center justify-center border-2 border-brand-primary">
            <span class="text-brand-gray text-lg">Profile Photo</span>
          </div>
        </div>

        <div class="mb-16">
          <h2 class="text-3xl font-semibold text-brand-dark mb-8 text-center">My Skills</h2>
          <div class="grid md:grid-cols-3 gap-8">
            <div class="bg-brand-white p-6 rounded-lg shadow-md border border-brand-gray">
              <h3 class="text-xl font-semibold text-brand-primary mb-4">Frontend</h3>
              <ul class="space-y-2 text-brand-dark">
                <li>• Vue.js & Nuxt.js</li>
                <li>• React & Next.js</li>
                <li>• TypeScript</li>
                <li>• Tailwind CSS</li>
              </ul>
            </div>
            <div class="bg-brand-white p-6 rounded-lg shadow-md border border-brand-gray">
              <h3 class="text-xl font-semibold text-brand-primary mb-4">Backend</h3>
              <ul class="space-y-2 text-brand-dark">
                <li>• Go (Golang)</li>
                <li>• Node.js</li>
                <li>• Python</li>
                <li>• RESTful APIs</li>
              </ul>
            </div>
            <div class="bg-brand-white p-6 rounded-lg shadow-md border border-brand-gray">
              <h3 class="text-xl font-semibold text-brand-primary mb-4">Database & Tools</h3>
              <ul class="space-y-2 text-brand-dark">
                <li>• PostgreSQL</li>
                <li>• MongoDB</li>
                <li>• Docker</li>
                <li>• Git & GitHub</li>
              </ul>
            </div>
          </div>
        </div>

        <div class="mb-16">
          <h2 class="text-3xl font-semibold text-brand-dark mb-8 text-center">My Journey</h2>
          <div class="space-y-8">
            <div class="bg-brand-white p-6 rounded-lg shadow-md border border-brand-gray">
              <div class="flex items-center mb-4">
                <div class="w-3 h-3 bg-brand-primary rounded-full mr-4"></div>
                <h3 class="text-xl font-semibold text-brand-dark">Senior Full Stack Developer</h3>
                <span class="ml-auto text-brand-gray">2022 - Present</span>
              </div>
              <p class="text-brand-dark">
                Leading development of enterprise web applications using modern tech stacks. 
                Mentoring junior developers and architecting scalable solutions.
              </p>
            </div>
            <div class="bg-brand-white p-6 rounded-lg shadow-md border border-brand-gray">
              <div class="flex items-center mb-4">
                <div class="w-3 h-3 bg-brand-primary rounded-full mr-4"></div>
                <h3 class="text-xl font-semibold text-brand-dark">Software Engineer</h3>
                <span class="ml-auto text-brand-gray">2020 - 2022</span>
              </div>
              <p class="text-brand-dark">
                Developed and maintained web applications for various clients. 
                Specialized in React, Node.js, and cloud deployment strategies.
              </p>
            </div>
            <div class="bg-brand-white p-6 rounded-lg shadow-md border border-brand-gray">
              <div class="flex items-center mb-4">
                <div class="w-3 h-3 bg-brand-primary rounded-full mr-4"></div>
                <h3 class="text-xl font-semibold text-brand-dark">Junior Developer</h3>
                <span class="ml-auto text-brand-gray">2019 - 2020</span>
              </div>
              <p class="text-brand-dark">
                Started my professional journey building responsive websites and 
                learning best practices in software development.
              </p>
            </div>
          </div>
        </div>

        <div class="text-center">
          <h2 class="text-3xl font-semibold text-brand-dark mb-6">Let's Connect</h2>
          <p class="text-brand-dark mb-8">
            I'm always interested in discussing new opportunities and interesting projects.
          </p>
          <div class="flex justify-center space-x-6">
            <a href="mailto:john@example.com" class="bg-brand-primary text-brand-dark px-6 py-3 rounded-lg hover:bg-brand-light transition font-semibold">
              Email Me
            </a>
            <a href="/contact" class="bg-gray-200 text-brand-dark px-6 py-3 rounded-lg hover:bg-gray-300 transition font-semibold">
              Contact Form
            </a>
          </div>
        </div>
      </div>
    </main>
    <Footer />
  </div>
</template>

<script setup>
useSeoMeta({
  title: 'About - Personal Web',
  description: 'Learn more about John Doe, a passionate full-stack developer with expertise in modern web technologies.'
})
</script>