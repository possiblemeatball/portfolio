<script setup>
import WebGL from './components/WebGL.vue'

</script>

<template>
  <div class="webgl w-[648px] h-[480px] mx-auto">
    <Suspense>
      <WebGL />
    </Suspense>
  </div>

  <router-view></router-view>
</template>

<style scoped>
</style>
