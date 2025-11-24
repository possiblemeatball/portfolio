<script setup>
const jobs = [
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
  <h3>Employment</h3>
  <div class="flex flex-col not-prose">
    <div v-for="job in jobs" class="p-2 flex flex-col gap-2">
      <span class="text-left">
        <a v-if="job.url" :href="job.url" target="_blank" class="underline font-bold text-lg">
          {{ job.company }}
        </a>
        <span v-else class="font-bold text-lg">
          {{ job.company }}
          <span v-if="job.defunct" class="text-xs font-semibold ml-1">(defunct)</span>
        </span>
        <span class="ml-2">{{ job.title }}</span>
        <span class="ml-2 text-sm text-neutral-600 dark:text-neutral-400">
          {{ job.start }}
          <span v-if="job.end && (job.end !== job.start)">&mdash; {{ job.end }}</span>
          <span v-else-if="job.end"></span>
          <span v-else>&mdash; Present</span>
        </span>
      </span>
      <span v-for="desc in job.description" class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">
        {{ desc }}
      </span>
    </div>
  </div>
</template>

<style scoped></style>