<script setup>
const props = defineProps(['blob'])
const blob = props.blob

</script>

<template>
<div class="p-2 flex flex-col gap-2">
  <span class="font-display font-bold text-lg">
    <a class="underline" v-if="blob.url" :href="blob.url" target="_blank">{{ blob.title }}</a>
    <span v-else>{{ blob.title }}</span>
  </span>
  <span class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">{{ blob.description }}</span>
</div>
</template>

<style scoped>

</style>