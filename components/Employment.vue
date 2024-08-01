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
  {
    company: "Luxury Haven LLC",
    defunct: true,
    url: null,
    title: "Reverse Engineering, Software Engineer",
    start: 2014,
    end: 2014,
    description: [
      "Perform reverse engineering on PC game software using early tools such as IDA Pro, Fernflower, ILSpy",
      "Create debug software for various PC games, including Source engine games, allowing for extensive modifications to the game",
      "Administer internal server software, ensuring the reliability of our software for over 10,000 monthly active users"
    ]
  },
]
</script>

<template>
  <div class="flex flex-col not-prose">
    <employ-blob v-for="blob in blobs" :blob="blob" />
  </div>
</template>

<style scoped>

</style>