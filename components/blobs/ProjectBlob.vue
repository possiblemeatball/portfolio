<script setup>
const props = defineProps(['title', 'url', 'description'])

</script>

<template>
<div class="p-2 flex flex-col gap-2">
  <span class="font-display font-bold text-lg">
    <a class="underline" v-if="props.url" :href="props.url">{{ props.title }}</a>
    <span v-else>{{ props.title }}</span>
  </span>
  <span class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">{{ props.description }}</span>
</div>
</template>

<style scoped>

</style>