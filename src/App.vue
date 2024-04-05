<script setup>
import {defineAsyncComponent} from "vue";

const Dashboard = defineAsyncComponent(() => import('./components/dashboard/Dashboard.vue'))
const Portfolio = defineAsyncComponent(() => import('./components/portfolio/Portfolio.vue'))
</script>

<template>
  <div class="bg-neutral-200 text-stone-900 dark:bg-stone-800 dark:text-neutral-100 overflow-x-hidden">
    <Suspense>
      <Dashboard />
    </Suspense>
    <Suspense>
      <Portfolio />
    </Suspense>
  </div>
</template>

<style scoped>
</style>
