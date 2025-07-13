<template>
  <NuxtLayout name="admin">
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-brand-dark">Admin Dashboard</h1>
      <p class="text-brand-dark">Welcome to the admin panel. From here you can manage your articles.</p>
      
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6 mt-8">
        <NuxtLink to="/admin/articles" class="block p-6 bg-brand-white rounded-lg shadow border border-brand-gray hover:shadow-md transition">
          <h3 class="text-xl font-semibold text-brand-dark mb-2">Articles</h3>
          <p class="text-brand-gray">Create, edit, and manage your articles</p>
        </NuxtLink>
        
        <NuxtLink to="/admin/companies" class="block p-6 bg-brand-white rounded-lg shadow border border-brand-gray hover:shadow-md transition">
          <h3 class="text-xl font-semibold text-brand-dark mb-2">Trusted Companies</h3>
          <p class="text-brand-gray">Manage client logos and trusted partners</p>
        </NuxtLink>
      </div>
    </div>
  </NuxtLayout>
</template>

<script setup>
definePageMeta({
  middleware: ['auth'],
});
</script>
