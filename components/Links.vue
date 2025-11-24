<script setup>
const links = {
    "Forgejo": "https://git.manthrowshat.net/meatball",
    "GitHub": "https://github.com/possiblemeatball",
    "Twitch": "https://www.twitch.tv/possiblemeatball",
    "E-mail": "mailto:meatball@manthrowshat.net",
    "Tip Me": "https://manthrowshat.net/tip",
}
</script>

<template>
    <span class="flex flex-row justify-between font-mono text-sm">
        <a v-for="link, title in links" :href="link" target="_blank" class="border px-1 no-underline">
            {{ title }}
        </a>
    </span>
</template>

<style lang="tailwindcss" scoped>
</style>