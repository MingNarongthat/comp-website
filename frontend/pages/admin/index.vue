<template>
  <NuxtLayout name="admin">
    <div class="container mx-auto p-4">
      <h1 class="text-3xl font-bold mb-6 text-brand-dark">Admin Dashboard</h1>
      <p class="text-brand-dark">Welcome to the admin panel. From here you can manage your articles and content.</p>
    </div>
  </NuxtLayout>
</template>

<script setup>
definePageMeta({
  middleware: ['auth'],
});
</script>
