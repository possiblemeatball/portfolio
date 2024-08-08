<script setup>
import ProjectBlob from "./blobs/ProjectBlob.vue";

const blobs = [
  {
    title: "NeverTrustLLC.com",
    url: "https://nevertrustllc.com",
    description: "Corporate front-page website and customer portal. "
  },
  {
    title: "Rock Eaters",
    url: "https://rockeaters.net",
    description: "Online services provider including e-mail, file upload, and search engine. Currently under heavy development.",
  },
  {
    title: "ManThrowsHat.net",
    url: "https://manthrowshat.net",
    description: "Personal website that I have a little bit of fun on. I make blog posts here, as well as accept tips if you're feeling generous."
  },
  {
    title: "Free Discord Bot",
    url: "https://manthrowshat.net/projects/freediscordbot",
    description: "Free and open source Discord Bot with a front-end web control panel. Currently in initial development stages.",
    tags: ['in-dev']
  },
  {
    title: "Hot Grabba",
    url: "https://manthrowshat.net/projects/hot-grabba",
    description: "Android and iOS companion app for Amazon Flex drivers. Assisted drivers in meeting goals and expectations of their gig economy work. Defunct as of 2024.",
    tags: ['defunct']
  },
  {
    title: "XIV",
    url: "https://manthrowshat.net/projects/xiv",
    description: "An event-driven Minecraft client modification. For legal and copyright reasons, source code is available upon request, with certain redactions.",
    tags: ['abandoned', 'legacy']
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