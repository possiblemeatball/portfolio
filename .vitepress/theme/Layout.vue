<script setup>
import {useData} from 'vitepress'
import TitleWebGL from './components/TitleWebGL.vue'

// https://vitepress.dev/reference/runtime-api#usedata
const { site, frontmatter } = useData()
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-8">
    <div class="min-h-72 w-screen flex flex-col justify-center items-center
                select-none pointer-events-none bg-neutral-50 dark:bg-neutral-950">
      <Suspense>
        <TitleWebGL />
      </Suspense>
      <h1 class="font-display text-2xl">Matt "possiblemeatball" Walker</h1>
      <p class="font-sans text-neutral-600 dark:text-neutral-400">{{ site.description }}</p>
    </div>
    <div class="container md:max-w-md">
      <div class="font-serif prose max-w-none prose-neutral dark:prose-invert">
        <Content />
      </div>
    </div>
  </div>
</template>
