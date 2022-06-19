import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 3030,
  },
  // set dist directory to ../public
  build: {
    outDir: '../public',
  },
  plugins: [react(),
    // VitePluginLinaria()

]
})
