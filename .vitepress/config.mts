import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "My Golang Handbook",
  description: "Quick tips on how to get things done using the Go programming language ",
  ignoreDeadLinks: true,
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Start', link: '/exercises/0001-hello-world/README' },
    ],

    sidebar: [
      {
        text: 'Exercises',
        base: '/exercises',
        items: [
          { text: 'Hello World', link: '/0001-hello-world//' }
        ]
      }
    ],

    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
