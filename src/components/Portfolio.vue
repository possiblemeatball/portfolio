<script setup>
import {defineAsyncComponent} from "vue";

const About = defineAsyncComponent(() => import('./about/About.vue'))
</script>

<template>
  <div class="min-h-lvh flex flex-col">
    <div class=
             "lg:container lg:rounded-lg mx-auto my-auto justify-center
              border bg-neutral-300 border-neutral-400 dark:bg-stone-700 dark:border-stone-600"
    >
      <About />
    </div>
  </div>
</template>

<style scoped>

</style>