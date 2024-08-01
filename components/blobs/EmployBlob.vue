<script setup>
const props = defineProps(['blob'])
const blob = props.blob

</script>

<template>
<div class="p-2 flex flex-col gap-2">
  <span class="text-left">
    <a class="underline font-display font-bold text-lg" v-if="blob.url" :href="blob.url" target="_blank">{{ blob.company }}</a>
    <span class="font-display font-bold text-lg" v-else>{{ blob.company }}<span v-if="blob.defunct" class="text-xs font-display font-semibold ml-1">(defunct)</span></span>
    <span class="ml-2 font-sans font-normal">{{ blob.title }}</span>
    <span class="ml-2 text-sm text-neutral-600 dark:text-neutral-400">{{ blob.start }} <span v-if="blob.end && (blob.end !== blob.start)">&mdash; {{ blob.end }}</span><span v-else-if="blob.end"></span><span v-else>&mdash; Present</span></span>
  </span>
  <span v-for="desc in blob.description" class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">{{ desc }}</span>
</div>
</template>

<style scoped>

</style>