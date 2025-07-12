import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: null,
    user: null,
  }),
  getters: {
    isAuthenticated: (state) => !!state.token,
    isSuperadmin: (state) => state.user && state.user.role === 'superadmin',
  },
  actions: {
    async login(email, password) {
      const config = useRuntimeConfig();
      const response = await fetch(`${config.public.apiBase}/api/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ email, password }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Login failed');
      }

      const data = await response.json();
      this.token = data.token;
      // Decode token to get user info (role, email)
      const base64Url = this.token.split('.')[1];
      const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
      const decoded = JSON.parse(window.atob(base64));
      this.user = { email: decoded.email, role: decoded.role };

      localStorage.setItem('token', this.token);
      localStorage.setItem('user', JSON.stringify(this.user));
    },
    logout() {
      this.token = null;
      this.user = null;
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    },
    initialize() {
      if (process.client) {
        this.token = localStorage.getItem('token');
        const user = localStorage.getItem('user');
        if (user) {
          this.user = JSON.parse(user);
        }
      }
    },
  },
});
