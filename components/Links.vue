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
        <a v-for="title, link in links" :href="title" target="_blank"
            class="border px-1 no-underline border-neutral-700 text-neutral-500 hover:border-neutral-200 hover:text-white">
            {{ link }}
        </a>
    </span>
</template>

<style lang="css" scoped></style>