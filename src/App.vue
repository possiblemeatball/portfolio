<script setup>
import {defineAsyncComponent} from "vue";

const Dashboard = defineAsyncComponent(() => import('./components/dashboard/Dashboard.vue'))
</script>

<template>
  <div class="bg-neutral-200 text-stone-900 dark:bg-stone-800 dark:text-neutral-100 overflow-x-hidden relative">
    <Suspense>
      <Dashboard />
    </Suspense>
  </div>
</template>

<style scoped>
</style>
