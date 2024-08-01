<script setup>
import ProjectBlob from "./blobs/ProjectBlob.vue";

const blobs = [
  {
    title: "ManThrowsHat.net",
    url: "https://manthrowshat.net",
    description: "Personal website that I have a little bit of fun on. I make blog posts here, as well as accept tips if you're feeling generous. Built using Golang with Fiber, TailwindCSS, and HTMX. Currently under heavy development."
  },
  {
    title: "Rock Eaters",
    url: "https://rockeaters.net",
    description: "Online services provider including e-mail, file upload, and search engine. Built using React, TailwindCSS, Postfix, Dovecot, and SearXNG. Currently under heavy development.",
  },
  {
    title: "Free Discord Bot",
    url: "https://github.com/possiblemeatball/freediscordbot",
    description: "Free and open source Discord Bot with a front-end web control panel. Currently in initial development stages.",
    tags: ['in-dev']
  },
  {
    title: "Hot Grabba",
    description: "Android and iOS companion app for Amazon Flex drivers. Assisted drivers in meeting goals and expectations of their gig economy work. Defunct as of 2024.",
    tags: ['defunct']
  },
  {
    title: "XIV",
    url: null,
    description: "An event-driven Minecraft client modification. Heavy LWJGL and OpenGL usage for game render modifications. For legal and copyright reasons, source code is available upon request, with certain redactions.",
    tags: ['abandoned', 'defunct']
  },
]
</script>

<template>
  <div class="flex flex-col not-prose">
    <project-blob v-for="blob in blobs" :blob="blob" />
  </div>
</template>

<style scoped>

</style>