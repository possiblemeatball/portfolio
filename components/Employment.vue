<script setup>
import EmployBlob from "blobs/EmployBlob.vue";

const blobs = [
  {
    company: "Never Trust LLC",
    url: "https://nevertrustllc.com",
    title: "Sole Proprietor, Systems Administrator, Software Engineer",
    start: 2020,
    end: 2025,
    description: [
      "Provide technical support for small businesses with medium to large workloads and/or customer traffic",
      "Administer customer resources including RHEL-based PostgreSQL/MongoDB/Redis database and NGINX web servers",
      "Secure cloud infrastructure against threat actors, intrusion attempts, and certain application DDoS attacks"
    ]
  },
  {
    company: "London Computer Systems",
    url: "https://lcs.com/",
    title: "Software Developer",
    start: 2023,
    end: 2023,
    description: [
      "Troubleshoot and identify programming defects before code is introduced",
      "Collaborate with technical mentors to write functional solutions to business problems",
      "Maintain existing application infrastructure, with documentation, test case creation, and regression preparation in Python."
    ]
  },
  {
    company: "Tri-County Church of the Nazarene",
    url: "https://tcn.church/",
    title: "Audio/Video Technical Engineer",
    start: 2017,
    end: 2020,
    description: [
      "Set up and maintain audio and video equipment for live events",
      "Ensure proper functionality of sound systems, projectors, and video displays",
      "Perform routine maintenance and repairs on AV equipment",
      "Test equipment to ensure seamless operation"
    ]
  },
]
</script>

<template>
  <div class="flex flex-col not-prose">
    <employ-blob v-for="blob in blobs" :blob="blob"/>
  </div>
</template>

<style scoped>

</style>