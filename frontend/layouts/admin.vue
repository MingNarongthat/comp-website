<template>
  <div class="min-h-screen bg-gray-100">
    <header class="w-full bg-brand-dark text-brand-white shadow-lg">
      <div class="px-8 py-4 flex items-center justify-between">
        <!-- Logo / Home Link -->
        <div class="flex items-center space-x-4">
          <NuxtLink 
            to="/" 
            class="flex items-center space-x-3 hover:text-brand-primary transition group"
            title="Back to Home"
          >
            <div class="bg-brand-primary p-2 rounded-lg group-hover:bg-brand-light transition">
              <svg class="w-6 h-6 text-brand-dark" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"></path>
              </svg>
            </div>
            <div>
              <div class="text-xl font-bold">Personal Web</div>
              <div class="text-xs text-brand-gray">Admin Panel</div>
            </div>
          </NuxtLink>
        </div>

        <!-- Navigation -->
        <nav class="flex items-center space-x-6">
          <NuxtLink 
            to="/admin" 
            class="flex items-center space-x-2 px-3 py-2 rounded-md hover:bg-gray-800 transition"
            :class="{ 'bg-gray-800 text-brand-primary': $route.path === '/admin' }"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"></path>
            </svg>
            <span>Dashboard</span>
          </NuxtLink>
          
          <NuxtLink 
            to="/admin/articles" 
            class="flex items-center space-x-2 px-3 py-2 rounded-md hover:bg-gray-800 transition"
            :class="{ 'bg-gray-800 text-brand-primary': $route.path.startsWith('/admin/articles') }"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z"></path>
            </svg>
            <span>Articles</span>
          </NuxtLink>

          <NuxtLink 
            to="/admin/companies" 
            class="flex items-center space-x-2 px-3 py-2 rounded-md hover:bg-gray-800 transition"
            :class="{ 'bg-gray-800 text-brand-primary': $route.path.startsWith('/admin/companies') }"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4"></path>
            </svg>
            <span>Companies</span>
          </NuxtLink>


          <div class="h-6 w-px bg-brand-gray"></div>

          <button 
            @click="logout" 
            class="flex items-center space-x-2 px-3 py-2 rounded-md hover:bg-red-600 hover:text-white transition"
            title="Logout"
          >
            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"></path>
            </svg>
            <span>Logout</span>
          </button>
        </nav>
      </div>
    </header>
    
    <main class="p-8">
      <slot />
    </main>
  </div>
</template>

<script setup>
import { useAuthStore } from '@/stores/auth';
import { useRouter } from 'vue-router';

const authStore = useAuthStore();
const router = useRouter();

const logout = () => {
  authStore.logout();
  router.push('/admin/login');
};
</script>
