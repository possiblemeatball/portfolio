<script setup>
const props = defineProps(['blob'])
const blob = props.blob

</script>

<template>
  <div
    class="bg-neutral-300 border-neutral-400 text-neutral-800 dark:bg-neutral-600 dark:border-neutral-700 dark:text-neutral-200 flex flex-row justify-center items-center text-center gap-1 rounded-md border px-2 select-none pointer-events-none">
    {{ blob.skill }}
  </div>
</template>

<style scoped></style>