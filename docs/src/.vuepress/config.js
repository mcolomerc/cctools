const { description } = require('../../package')
import { searchPlugin } from '@vuepress/plugin-search'
import { backToTopPlugin } from '@vuepress/plugin-back-to-top'
import { defaultTheme } from '@vuepress/theme-default'
import { nprogressPlugin } from '@vuepress/plugin-nprogress'
import { gitPlugin } from '@vuepress/plugin-git'
import { prismjsPlugin } from '@vuepress/plugin-prismjs'
import { clipboardPlugin } from 'vuepress-plugin-clipboard'

export default {
  logo: '/logo.png',
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#title
   */
  title: '>_ cctools',
  /**
   * Ref：https://v1.vuepress.vuejs.org/config/#description
   */
  description: description,

  base: "/cctools/",

  /**
   * Extra tags to be injected to the page HTML `<head>`
   *
   * ref：https://v1.vuepress.vuejs.org/config/#head
   */
  head: [
    ['meta', { name: 'theme-color', content: '#3eaf7c' }],
    ['meta', { name: 'apple-mobile-web-app-capable', content: 'yes' }],
    ['meta', { name: 'apple-mobile-web-app-status-bar-style', content: 'black' }]
  ],
  theme: defaultTheme({
    // set config here
    navbar: [
      {
        text: 'Guide',
        link: '/guide/',
      },
      {
        text: 'Configurations',
        link: '/config/'
      },
      {
        text: 'Commands',
        link: '/commands/'
      },
      {
        text: 'GitHub',
        link: 'https://github.com/mcolomerc/cctools'
      }
    ],
    
  }),
  plugins: [
    searchPlugin({
      // options
    }),
    backToTopPlugin(),
    nprogressPlugin(),
    gitPlugin({
      // options
      contributors: true,
      lastUpdated: true,
    }),
    prismjsPlugin({
      // options
      preloadLanguages: ['bash', 'json', 'yaml', 'properties', 'markdown', 'shell'],
    }),
    clipboardPlugin({
      staticIcon: true,
      align: 'top',
    }),
  ],
  repo: 'mcolomerc/cctools',
  docsRepo: 'https://github.com/mcolomerc/cctools',
  docsBranch: 'main',
  docsDir: 'docs',
  editLinkPattern: ':repo/-/edit/:branch/:path',
  themePlugins: {
    nprogresst: true,
    git: true,
    backToTop: true,
  },

}
