import { vitePreprocess } from '@sveltejs/vite-plugin-svelte'

export default {
  server: {
    port: 3030
  },
  preview: {
    port: 8080
  },
  host: "0.0.0.0",
  preprocess: vitePreprocess(),
}
