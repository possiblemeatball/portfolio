<script setup>
const projects = [
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
    description: "Free and open source Discord Bot with a front-end web control panel"
  },
  {
    title: "NeverTrustLLC.com",
    url: "https://nevertrustllc.com",
    description: "Corporate front-page website and customer portal"
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
    <div class="p-2 flex flex-col gap-2" v-for="project in projects">
      <span class="font-display font-bold text-lg">
        <a class="underline" v-if="project.url" :href="project.url" target="_blank">{{ project.title }}</a>
        <span v-else>{{ project.title }}</span>
        <span
          class="ml-1 px-1 border rounded-md text-sm font-semibold bg-neutral-300 border-neutral-400 text-neutral-800 dark:bg-neutral-600 dark:border-neutral-700 dark:text-neutral-200"
          v-if="project.tags" v-for="tag in project.tags">{{ tag }}</span>
      </span>
      <span class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">{{ project.description }}</span>
    </div>
  </div>
</template>

<style scoped></style>