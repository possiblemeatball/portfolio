<script setup>
const props = defineProps(['title', 'url', 'description'])

</script>

<template>
<div class="p-2 flex flex-col gap-2">
  <span v-if="props.url">
    <a :href="props.url">{{ props.title }}</a>
  </span>
  <span v-else>
    {{ props.title }}
  </span>
  <span class="text-sm text-neutral-600 dark:text-neutral-400">{{ props.description }}</span>
</div>
</template>

<style scoped>

</style>