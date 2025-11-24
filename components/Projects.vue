<script setup>
import ProjectBlob from "./blobs/ProjectBlob.vue";

const blobs = [
  {
    title: "Rock Eaters Network",
    url: "https://rockeaters.net",
    description: "Online services provider including e-mail, file upload, and search engine",
  },
  {
    title: "libnomnom",
    url: "https://manthrowshat.net/projects/libnomnom",
    description: "Library for documenting, utilizing, and abstracting the API endpoints for Freddy's Steakburgers",
  },
  {
    title: "Free Discord Bot",
    url: "https://manthrowshat.net/projects/freediscordbot",
    description: "Free and open source Discord Bot with a front-end web control panel",
    tags: ['in-dev']
  },
  {
    title: "NeverTrustLLC.com",
    url: "https://nevertrustllc.com",
    description: "Corporate front-page website and customer portal",
    tags: ['defunct']
  },
  {
    title: "ManThrowsHat.net",
    url: "https://manthrowshat.net",
    description: "Personal project and blogsite page where I wax poetic about what's rattling around my brain"
  },
  {
    title: "Hot Grabba",
    url: "https://manthrowshat.net/projects/hot-grabba",
    description: "Android and iOS companion app for Amazon Flex drivers. Assisted drivers in meeting goals and expectations of their gig economy work. Defunct as of 2024",
    tags: ['abandoned']
  },
  {
    title: "XIV",
    url: "https://manthrowshat.net/projects/xiv",
    description: "An event-driven Minecraft client modification for Minecraft version 1.8",
    tags: ['abandoned']
  },
]
</script>

<template>
  <h3>Projects</h3>
  <div class="flex flex-col not-prose">
    <project-blob v-for="blob in blobs" :blob="blob" />
  </div>
</template>

<style scoped></style>