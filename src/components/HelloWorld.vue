<script setup>
import { ref } from 'vue'

const count = ref(0)
</script>

<template>
  <h1 class="font-semibold text-5xl text-center py-6">possiblemeatball</h1>
  <hr>

  <div class="container mx-auto text-center my-4 p-4 bg-neutral-400 dark:bg-stone-900 w-1/3">
    <button class="rounded px-8 py-2 mx-auto bg-natural-300 text-stone-800 dark:bg-stone-700 dark:text-neutral-300" type="button" @click="count++">count is {{ count }}</button>
    <p>
      Edit
      <code>components/HelloWorld.vue</code> to test HMR
    </p>
  </div>
</template>

<style scoped>
</style>
