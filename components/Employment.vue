<script setup>
import EmployBlob from "./blobs/EmployBlob.vue";

const blobs = [
  {
    company: "Never Trust LLC",
    url: "https://nevertrustllc.com",
    title: "Sole Proprietor, Systems Administrator, Software Engineer",
    start: 2020,
    end: null,
    description: [
        "Provide technical support for small businesses with medium to large workloads and/or customer traffic",
        "Administer customer resources including RHEL-based PostgreSQL/MongoDB/Redis database and NGINX web servers",
        "Secure cloud infrastructure against threat actors, intrusion attempts, and certain application DDoS attacks"
    ]
  },
]
</script>

<template>
  <div class="flex flex-col not-prose">
    <employ-blob v-for="blob in blobs" :title="blob.title" :company="blob.company" :url="blob.url" :start="blob.start" :end="blob.end" :description="blob.description" />
  </div>
</template>

<style scoped>

</style>