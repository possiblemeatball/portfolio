<script setup>
const {title} = defineProps(['title'])
</script>

<template>
  <div class=
           "w-screen md:container md:rounded-lg mx-auto my-auto justify-center
              border bg-neutral-300 border-neutral-400 dark:bg-stone-700 dark:border-stone-600"
  >
    <h1 class="p-4 px-8 pb-2 font-semibold text-4xl">{{ title }}</h1>
    <div class="px-4 pb-8"><hr class="border-stone-700 dark:border-neutral-300" /></div>
  </div>
</template>

<style scoped>

</style>