import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  base: './',
  plugins: [react()],
  preview: {
    port: 8081,
    host: '0.0.0.0',
    allowedHosts: ['www.alchemyrecipes.online'] // Tambahkan ini
  }
})
