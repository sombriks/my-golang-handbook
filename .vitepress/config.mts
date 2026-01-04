import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "My Golang Handbook",
  description: "Quick tips on how to get things done using the Go programming language",
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
          { text: 'Hello World', link: '/0001-hello-world/README' },
          { text: 'Guess Number - Params', link: '/0002-guess-number-param/README' },
          { text: 'Guess Number - Interactive', link: '/0003-guess-number-interactive/README' },
          { text: 'Check Triangles', link: '/0004-check-triangles/README' },
          { text: 'TODO List - File', link: '/0005-todo-list-file/README' },
          { text: 'Count Letters', link: '/0006-count-letters/README' },
          { text: 'Tests', link: '/0007-tests/README' },
          { text: 'Go Routines', link: '/0008-go-routines/README' },
          { text: 'Go Channels', link: '/0009-go-channels/README' },
          { text: 'Classic Mutex', link: '/0010-classic-mutex/README' },
          { text: 'Reflection Basics', link: '/0011-reflection-basics/README' },
          { text: 'IO Over Networking', link: '/0012-io-over-networking/README' },
          { text: 'Databases', link: '/0013-databases/README' },
          { text: 'ORM With GORM', link: '/0014-orm-with-gorm/README' },
          { text: 'REST With HTMX', link: '/0015-rest-htmx/README' },
          { text: 'REST JSON', link: '/0016-rest-json/README' },
          { text: 'Sample gRPC', link: '/0017-sample-grpc/README' },
          { text: 'Events and Messages', link: '/0018-events-and-messages/README' }
        ],
      }
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/sombriks/my-golang-handbook' },
      { icon: 'bluesky', link: 'https://bsky.app/profile/sombriks.com.br' }
    ]
  }
})
