import {defineConfig} from 'vitepress'
import Tailwind from '@tailwindcss/vite'

// https://vitepress.dev/reference/site-config
export default defineConfig({
    lang: "en-US",
    title: "Matthew \"possiblemeatball\" Walker",
    description: "Software Engineer, Systems Administrator, Entrepreneur",
    cleanUrls: true,
    srcExclude: ["README.md"],
    vite: {
        plugins: [Tailwind()]
    }
})
