<script setup>
import { defineAsyncComponent } from "vue"
import Container from './Container.vue'
const About = defineAsyncComponent(() => import("./about/About.vue"))
const Experience = defineAsyncComponent(() => import("./experience/Experience.vue"))
const Projects = defineAsyncComponent(() => import("./projects/Projects.vue"))

</script>

<template>
  <div class="min-h-lvh flex flex-col">
    <Suspense>
      <Container :component="About" />
    </Suspense>
    <Suspense>
      <Container :component="Experience" />
    </Suspense>
    <Suspense>
      <Container :component="Projects" />
    </Suspense>
  </div>
</template>

<style scoped>

</style>