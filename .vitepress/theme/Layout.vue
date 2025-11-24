<script setup>
import { useData } from 'vitepress'
import MeatballCanvas from '../../components/MeatballCanvas.vue'
import Links from '../../components/Links.vue';
import Projects from "../../components/Projects.vue";
import Employment from "../../components/Employment.vue";
import Education from '../../components/Education.vue';
import Skills from '../../components/Skills.vue';
import { TresCanvas } from "@tresjs/core";

// https://vitepress.dev/reference/runtime-api#usedata
const { site } = useData()

const rendererSettings = {
  antialias: false,
  alpha: true,
  clearAlpha: '0',
  powerPreference: 'low-power',
}
</script>

<template>
  <div class="flex flex-col gap-0 justify-center items-center overflow-x-hidden">
    <div
      class="p-4 pt-20 w-screen flex flex-col justify-center items-center text-center select-none pointer-events-none bg-white dark:bg-neutral-900">
      <div class="animate-slide-top w-[511px] h-[128px]">
        <TresCanvas v-bind="rendererSettings">
          <suspense>
            <MeatballCanvas />
          </suspense>
        </TresCanvas>
      </div>
      <h1 class="text-2xl">{{ site.title }}</h1>
      <p class="italic text-neutral-600 dark:text-neutral-400">{{ site.description }}</p>
    </div>
    <div class="container max-w-screen-sm p-4 text-justify">
      <div class="prose max-w-none prose-neutral dark:prose-invert">
        <Links />
        <content />
        <Projects />
        <Employment />
        <Education />
        <Skills />
      </div>
    </div>
  </div>
</template>
