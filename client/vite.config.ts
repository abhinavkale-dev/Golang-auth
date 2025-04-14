import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      '/signup': 'http://localhost:8080',
      '/signin': 'http://localhost:8080',
      '/signout': 'http://localhost:8080'
    }
  }
})
