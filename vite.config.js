import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  base: './', // atau bisa juga '/' untuk root
  plugins: [react()],
})
