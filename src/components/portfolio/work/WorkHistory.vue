<script setup>
const jobHistory = [
  {
    title: 'Owner and Sole Proprietor',
    company: 'Never Trust LLC',
    time_period: { start: Date.UTC(2020, 9) },
    duties: [
      '',
    ]
  },
  {
    title: 'Package Handler',
    company: 'FedEx Corporation',
    time_period: { start: Date.UTC(2019, 8), end: Date.UTC(2020, 10) },
    duties: [
      '',
    ]
  },
  {
    title: 'Head Kitchen Supervisor (Seasonal)',
    company: 'Cedar Fair Entertainment Company',
    time_period: { start: Date.UTC(2017, 5), end: Date.UTC(2019, 0) },
    duties: [
      '',
    ]
  },
  {
    title: 'Kitchen Staff',
    company: 'Yum! Brands Inc',
    time_period: { start: Date.UTC(2016, 4), end: Date.UTC(2017, 3) },
    duties: [
      '',
    ]
  },
]
</script>

<template>
  <div class="p-4">
    <h1 class="px-2 pb-8 font-semibold text-4xl">Work History</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
      <div v-for="job in jobHistory"
           class=
               "p-2 px-4
              rounded-lg bg-neutral-200 dark:bg-stone-800
              shadow shadow-neutral-300 dark:shadow-stone-800"
      >
        <h2 class="font-semibold">{{ job.title }}</h2>
        <h3>{{ job.company }}</h3>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>