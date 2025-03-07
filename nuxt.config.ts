// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: '2024-11-01',
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', '@nuxtjs/google-fonts', '@nuxt/icon'],
  runtimeConfig: {
    server_url: process.env.SERVER_URL,
    public: {
      server_url: process.env.SERVER_URL
    }
  },
  googleFonts: {
    families: {
      Inter: '200..900'
    }
  }
})