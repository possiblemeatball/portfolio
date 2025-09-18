<script setup>
import ProjectBlob from "../../../components/blobs/ProjectBlob.vue";

const blobs = [
  {
    title: "Rock Eaters Network",
    url: "https://rockeaters.net",
    description: "Online services provider including e-mail, file upload, and search engine.",
  },
  {
    title: "ManThrowsHat.net",
    url: "https://manthrowshat.net",
    description: "Personal web page and blogsite where I post articles and abstract free-form ramblings"
  },
  {
    title: "Free Discord Bot",
    url: "https://manthrowshat.net/projects/freediscordbot",
    description: "Free and open source Discord Bot with a front-end web control panel.",
    tags: ['in-dev']
  },
  {
    title: "NeverTrustLLC.com",
    url: "https://nevertrustllc.com",
    description: "Corporate front-page website and customer portal.",
    tags: ['defunct']
  },
  {
    title: "Hot Grabba",
    url: "https://manthrowshat.net/projects/hot-grabba",
    description: "Android and iOS companion app for Amazon Flex drivers. Assisted drivers in meeting goals and expectations of their gig economy work. Defunct as of 2024.",
    tags: ['abandoned']
  },
  {
    title: "XIV",
    url: "https://manthrowshat.net/projects/xiv",
    description: "An event-driven Minecraft client modification. For legal and copyright reasons, source code is available upon request, with certain redactions.",
    tags: ['abandoned']
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