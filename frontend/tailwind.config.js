/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./components/**/*.{js,vue,ts}",
    "./layouts/**/*.vue",
    "./pages/**/*.vue",
    "./plugins/**/*.{js,ts}",
    "./nuxt.config.{js,ts}",
    "./app.vue"
  ],
  theme: {
    extend: {
      colors: {
        // Brand Primary Colors
        'brand': {
          'primary': '#ffd600',    // Yellow primary
          'dark': '#030d0f',       // Dark primary
          'light': '#ffe58b',      // Light yellow
          'gray': '#bdbdbd',       // Brand gray
          'white': '#ffffff'       // Pure white
        },
        // Override default colors with brand colors
        'yellow': {
          '50': '#fffdf0',
          '100': '#fffbeb',
          '200': '#fff5c7',
          '300': '#ffeda3',
          '400': '#ffe58b',
          '500': '#ffd600',  // Brand primary
          '600': '#e6c000',
          '700': '#ccaa00',
          '800': '#b39400',
          '900': '#997e00'
        },
        'gray': {
          '50': '#fafafa',
          '100': '#f5f5f5',
          '200': '#eeeeee',
          '300': '#e0e0e0',
          '400': '#bdbdbd',  // Brand gray
          '500': '#9e9e9e',
          '600': '#757575',
          '700': '#616161',
          '800': '#424242',
          '900': '#030d0f'   // Brand dark
        },
        'slate': {
          '50': '#f8fafc',
          '100': '#f1f5f9',
          '200': '#e2e8f0',
          '300': '#cbd5e1',
          '400': '#94a3b8',
          '500': '#64748b',
          '600': '#475569',
          '700': '#334155',
          '800': '#1e293b',
          '900': '#030d0f'   // Brand dark
        }
      }
    },
  },
  plugins: [],
}