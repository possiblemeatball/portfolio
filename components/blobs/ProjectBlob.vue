<script setup>
const props = defineProps(['blob'])
const blob = props.blob

</script>

<template>
<div class="p-2 flex flex-col gap-2">
  <span class="font-display font-bold text-lg">
    <a class="underline" v-if="blob.url" :href="blob.url" target="_blank">{{ blob.title }}</a>
    <span v-else>{{ blob.title }}</span>
    <span class="ml-1 px-1 border rounded-md text-sm font-semibold
                 bg-neutral-300 border-neutral-400 text-neutral-800
                 dark:bg-neutral-700 dark:border-neutral-800 dark:text-neutral-200" v-if="blob.tags" v-for="tag in blob.tags">{{ tag }}</span>
  </span>
  <span class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">{{ blob.description }}</span>
</div>
</template>

<style scoped>

</style>