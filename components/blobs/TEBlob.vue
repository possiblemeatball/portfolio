<script setup>
const props = defineProps(['blob'])
const blob = props.blob

</script>

<template>
<div class="bg-neutral-300 text-neutral-800 dark:bg-neutral-700 dark:text-neutral-200 rounded-md px-2
            flex flex-row justify-center items-center text-center gap-1 select-none pointer-events-none">
  {{ blob.skill }}
</div>
</template>

<style scoped>

</style>