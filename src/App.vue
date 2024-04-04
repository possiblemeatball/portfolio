<script setup>
import TitleWebGL from './components/TitleWebGL.vue'

</script>

<template>
  <div class="w-screen h-screen flex flex-col justify-center">
    <Suspense>
      <TitleWebGL />
    </Suspense>
    <div class="animate-slide-bottom flex justify-center">meatball</div>
  </div>


  <router-view></router-view>
</template>

<style scoped>
</style>
