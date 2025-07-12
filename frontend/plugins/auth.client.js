import { useAuthStore } from '@/stores/auth';

export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook('app:created', () => {
    const authStore = useAuthStore();
    authStore.initialize();
  });
});
