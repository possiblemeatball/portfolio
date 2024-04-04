<script setup>
import TitleWebGL from './components/TitleWebGL.vue'

</script>

<template>
  <div class="w-screen h-screen flex flex-col justify-center">
    <div class="flex-none mx-auto w-[640px] h-[256px]">
      <Suspense>
        <TitleWebGL />
      </Suspense>
    </div>
    <div class="flex justify-center">meatball</div>
  </div>


  <router-view></router-view>
</template>

<style scoped>
</style>
