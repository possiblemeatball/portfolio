<script setup>
import ProjectBlob from "./blobs/ProjectBlob.vue";

const blobs = [
  {
    title: "ManThrowsHat.net",
    url: "https://manthrowshat.net",
    description: "Personal website that I have a little bit of fun on. I make blog posts here, as well as accept tips if you're feeling generous. Built using Golang with Fiber, TailwindCSS, and HTMX."
  },
  {
    title: "Rock Eaters",
    url: "https://rockeaters.net",
    description: "Online services provider including e-mail, file upload, and search engine. Built using React, TailwindCSS, Postfix, Dovecot, and SearXNG.",
    tags: ["rocks", "rocks", "rocks"]
  },
  {
    title: "Free Discord Bot",
    url: "https://github.com/possiblemeatball/freediscordbot",
    description: "Free and open source Discord Bot with a front-end web control panel. Currently in initial development stages.",
    tags: ['in-dev']
  },
  {
    title: "XIV",
    url: null,
    description: "An event-driven Minecraft client modification including ESP, Aim bot, and Anticheating avoidance. For legal and copyright reasons, source code is available upon request, with certain redactions.",
    tags: ['abandoned', 'defunct', "blarg"]
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