<script setup>
const {component} = defineProps(['component'])
</script>

<template>
  <div class=
           "w-screen md:container md:rounded-lg mx-auto my-auto justify-center min-h-24
              border bg-neutral-100 border-neutral-300 dark:bg-stone-700 dark:border-stone-900
            shadow-2xl shadow-neutral-300 dark:shadow-stone-900"
  >
    <component :is="component" />
  </div>
</template>

<style scoped>

</style>