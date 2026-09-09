import { defineConfig } from 'vitepress'

export default defineConfig({
  title: 'Za-go',
  description: 'Unofficial Zalo API for Golang',
  themeConfig: {
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Quickstart', link: '/guide/quickstart' },
      { text: 'API', link: '/api/overview' }
    ],
    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Quickstart', link: '/guide/quickstart' }
        ]
      },
      {
        text: 'API',
        items: [
          { text: 'Overview', link: '/api/overview' }
        ]
      }
    ]
  }
})
