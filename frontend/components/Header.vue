<template>
  <header class="w-full bg-brand-white shadow py-4 px-8 flex items-center justify-between">
    <NuxtLink to="/" class="flex items-center space-x-3 hover:opacity-80 transition-opacity">
      <img src="~/assets/KH logo.png" alt="KH Logo" class="h-10 w-auto">
      <!-- <div class="text-2xl font-bold text-brand-dark">Personal Web</div> -->
    </NuxtLink>
    <nav class="flex items-center space-x-4">
      <NuxtLink to="/" class="mx-2 text-brand-dark hover:text-brand-primary font-medium">Home</NuxtLink>
      <NuxtLink to="/about" class="mx-2 text-brand-dark hover:text-brand-primary font-medium">About</NuxtLink>
      <NuxtLink to="/articles" class="mx-2 text-brand-dark hover:text-brand-primary font-medium">Articles</NuxtLink>
      <NuxtLink to="/contact" class="mx-2 text-brand-dark hover:text-brand-primary font-medium">Contact</NuxtLink>
      <!-- <NuxtLink 
        to="/admin/login" 
        class="bg-brand-primary text-brand-dark px-4 py-2 rounded-md hover:bg-brand-light transition duration-200 ml-4 font-semibold"
      >
        Login
      </NuxtLink> -->
    </nav>
  </header>
</template> 
