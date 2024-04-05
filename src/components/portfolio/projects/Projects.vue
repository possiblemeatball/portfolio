<script setup>
const projects = [
  {
    title: 'Test Project',
    description: 'test'
  },
  {
    title: 'Test Project',
    description: 'test'
  },
  {
    title: 'Test Project',
    description: 'test'
  },
  {
    title: 'Test Project',
    description: 'test'
  },
]
</script>

<template>
  <div class="p-4">
    <h1 class="px-2 pb-8 font-semibold text-4xl">Project History</h1>
    <div class="grid grid-cols-1 md:grid-cols-2 gap-2">
      <div v-for="project in projects"
           class=
               "p-2 px-4
                rounded-lg bg-neutral-200 dark:bg-stone-800
                shadow shadow-neutral-300 dark:shadow-stone-800"
      >
        <h2 class="font-semibold">{{ project.title }}</h2>
        <h3 class="">{{ project.description }}</h3>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>