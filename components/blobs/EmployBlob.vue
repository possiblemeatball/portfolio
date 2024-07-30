<script setup>
const props = defineProps(['company', 'url', 'title', 'start', 'end', 'description'])

</script>

<template>
<div class="p-2 flex flex-col gap-2">
  <span class="text-left">
    <a class="underline font-display font-bold text-lg" v-if="props.url" :href="props.url" target="_blank">{{ props.company }}</a>
    <span class="font-display font-bold text-lg" v-else>{{ props.company }}</span>
    <span class="ml-2 font-sans font-normal">{{ props.title }}</span>
    <span class="ml-2 text-sm text-neutral-600 dark:text-neutral-400">{{ props.start }} &mdash; <span v-if="props.end">{{ props.end }}</span> <span v-else>Present</span></span>
  </span>
  <span v-for="desc in props.description" class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">{{ desc }}</span>
</div>
</template>

<style scoped>

</style>