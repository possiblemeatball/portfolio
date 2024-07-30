<script setup>
import {useData} from 'vitepress'
import TitleWebGL from '../../components/TitleWebGL.vue'

// https://vitepress.dev/reference/runtime-api#usedata
const { site, frontmatter } = useData()
</script>

<template>
  <div class="flex flex-col justify-center items-center overflow-x-hidden">
    <div class="p-4 pt-24 pb-4 w-screen flex flex-col justify-center items-center text-center
                select-none pointer-events-none bg-white dark:bg-black">
      <Suspense>
        <TitleWebGL class="animate-slide-top" />
      </Suspense>
      <h1 class="font-display text-2xl">Matt "{{ site.title }}" Walker</h1>
      <p class="font-sans text-neutral-600 dark:text-neutral-400">{{ site.description }}</p>
    </div>
    <div class="container max-w-screen-sm p-8 pt-0 text-justify">
      <div class="prose prose-h3:font-display prose-p:font-sans max-w-none prose-neutral dark:prose-invert">
        <Content />
      </div>
    </div>
  </div>
</template>
