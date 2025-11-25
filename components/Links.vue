<script setup>
import { useData } from 'vitepress'
const { site } = useData()
</script>

<template>
    <span class="pt-2 flex flex-row gap-4 font-mono text-sm">
        <a v-for="link, text in site.themeConfig.nav" :href="link" target="_blank" class="border px-1 no-underline">
            {{ text }}
        </a>
    </span>
</template>

<style lang="tailwindcss" scoped>
</style>