<script setup>
import {Content, useData} from 'vitepress'
import TitleWebGL from './components/TitleWebGL.vue'

// https://vitepress.dev/reference/runtime-api#usedata
const { site, frontmatter } = useData()
</script>

<template>
  <div class="flex flex-col justify-center items-center gap-8">
    <div class="min-h-72 w-screen flex flex-col justify-center items-center bg-neutral-50 dark:bg-neutral-950">
      <Suspense><TitleWebGL /></Suspense>
      <h1 class="font-bold text-2xl">{{ site.title }}</h1>
      <p>{{ site.description }}</p>
    </div>
    <div class="container md:max-w-md">
      <div class="prose max-w-none prose-neutral dark:prose-invert">
        <Content />
      </div>
    </div>
  </div>
</template>
