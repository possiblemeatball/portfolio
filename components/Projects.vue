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
    description: "Online services provider including e-mail, file upload, and search engine. Built using Golang, React, TailwindCSS, Postfix, Dovecot, and SearXNG."
  },
]
</script>

<template>
  <div class="flex flex-col">
    <project-blob v-for="blob in blobs" :title="blob.title" :url="blob.url" :description="blob.description" />
  </div>
</template>

<style scoped>

</style>