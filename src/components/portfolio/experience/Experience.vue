<script setup>
import experience from './data'

</script>

<template>
  <div v-for="item in experience">
    <div v-if="item" class="p-4">
      <h1 v-if="item.title" class="px-2 pb-2 font-semibold text-4xl">{{ item.title }}</h1>
      <div v-if="item.title" class="pb-4"><hr class="border-stone-900 dark:border-neutral-100" /></div>
      <div v-if="item.text" class="px-4 pt-0 pb-2">{{ item.text }}</div>
    </div>
  </div>
</template>

<style scoped>

</style>