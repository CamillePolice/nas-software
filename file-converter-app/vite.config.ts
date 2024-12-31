import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path' // Correctly import resolve

export default defineConfig({
  plugins: [vue()],
  resolve: {
    alias: {
      '@': resolve(__dirname, './src'), // Define alias for '@'
    },
  },
  server: {
    port: 5173, // Optional: Set a custom port
  },
})
