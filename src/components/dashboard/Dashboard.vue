<script setup>
import {defineAsyncComponent} from "vue";

const TitleWebGL = defineAsyncComponent(() =>
    import('./TitleWebGL.vue')
)
</script>

<template>
  <div class=
           "w-screen h-screen flex flex-col items-center justify-center text-center
            bg-neutral-50 dark:bg-stone-950"
  >
    <div class="text-center">
      <div class="animate-slide-top flex flex-row gap-8 items-center justify-center">
        <img class="w-[64px]" src="../../assets/images/under_construction3.gif" alt="under construction ">
        <Suspense>
          <TitleWebGL/>
        </Suspense>
        <img class="w-[64px]" src="../../assets/images/under_construction3.gif" alt="under construction ">
      </div>
      <div class="animate-slide-bottom flex flex-row items-center justify-center">
        <img class="w-[64px]" src="../../assets/images/under_construction2.gif" alt="under construction ">
        <img class="h-[64px]" src="../../assets/images/under_construction.gif" alt="under construction ">
        <img class="w-[64px]" src="../../assets/images/under_construction2.gif" alt="under construction ">
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>