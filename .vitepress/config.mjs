import { defineConfig } from 'vitepress'
import Tailwind from '@tailwindcss/vite'

// https://vitepress.dev/reference/site-config
export default defineConfig({
    lang: "en-US",
    title: "Matthew \"possiblemeatball\" Walker",
    description: "Software Engineer, Systems Administrator, Entrepreneur",
    cleanUrls: true,
    srcExclude: ["README.md"],
    themeConfig: {
        nav: {
            "Forgejo": "https://git.manthrowshat.net/meatball",
            "E-mail": "mailto:meatball@manthrowshat.net",
            "Discord": "https://discord.com/users/555553450583130117",
            "Blog": "https://manthrowshat.net/blog",
            "Tip Me!": "https://manthrowshat.net/tip",
        }
    },
    vite: {
        plugins: [Tailwind()]
    }
})
