<template>
  <div class="flex flex-col items-center justify-center min-h-screen bg-gray-100">
    <div class="w-full max-w-md p-8 space-y-6 bg-brand-white rounded shadow-md border border-brand-gray">
      <h2 class="text-2xl font-bold text-center text-brand-dark">Admin Login</h2>
      <form @submit.prevent="login" class="space-y-4">
        <div>
          <label for="email" class="block text-sm font-medium text-brand-dark">Email</label>
          <input
            type="email"
            id="email"
            v-model="email"
            required
            class="block w-full px-3 py-2 mt-1 border border-brand-gray rounded-md shadow-sm focus:outline-none focus:ring-brand-primary focus:border-brand-primary sm:text-sm"
          />
        </div>
        <div>
          <label for="password" class="block text-sm font-medium text-brand-dark">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            required
            class="block w-full px-3 py-2 mt-1 border border-brand-gray rounded-md shadow-sm focus:outline-none focus:ring-brand-primary focus:border-brand-primary sm:text-sm"
          />
        </div>
        <button
          type="submit"
          class="w-full px-4 py-2 text-brand-dark bg-brand-primary rounded-md hover:bg-brand-light focus:outline-none focus:ring-2 focus:ring-brand-primary focus:ring-offset-2 font-semibold"
        >
          Login
        </button>
        <p v-if="error" class="text-red-500 text-sm text-center">{{ error }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '@/stores/auth';

const email = ref('');
const password = ref('');
const error = ref(null);
const router = useRouter();
const authStore = useAuthStore();

const login = async () => {
  error.value = null;
  try {
    await authStore.login(email.value, password.value);
    router.push('/admin');
  } catch (err) {
    error.value = err.message || 'Login failed';
  }
};
</script>
