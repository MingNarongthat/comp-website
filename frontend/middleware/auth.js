import { useAuthStore } from '@/stores/auth';

export default defineNuxtRouteMiddleware((to, from) => {
  const authStore = useAuthStore();
  authStore.initialize(); // Ensure store is initialized on client-side navigation

  if (!authStore.isAuthenticated && to.path.startsWith('/admin')) {
    return navigateTo('/admin/login');
  }
});
