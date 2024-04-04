<script setup>
import {defineAsyncComponent} from "vue";

const Dashboard = defineAsyncComponent(() => import('./components/dashboard/Dashboard.vue'))
const Portfolio = defineAsyncComponent(() => import('./components/portfolio/Portfolio.vue'))
</script>

<template>
  <Dashboard />
  <Portfolio />
</template>

<style scoped>
</style>
