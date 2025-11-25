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
            "GitHub": "https://github.com/possiblemeatball",
            "Twitch": "https://www.twitch.tv/possiblemeatball",
            "E-mail": "mailto:meatball@manthrowshat.net",
            "Tip Me": "https://manthrowshat.net/tip",
        }
    },
    vite: {
        plugins: [Tailwind()]
    }
})
