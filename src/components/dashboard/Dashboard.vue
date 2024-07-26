<script setup>
import {defineAsyncComponent} from "vue";

const TitleWebGL = defineAsyncComponent(() =>
    import('./TitleWebGL.vue')
)
</script>

<template>
  <div class=
           "w-screen h-screen relative
            bg-neutral-50 dark:bg-stone-950"
  >
    <div class="absolute text-center left-1/2 -translate-x-1/2 md:top-1/2 md:-translate-y-1/2 z-10">
      <div class="animate-slide-top">
        <Suspense>
          <TitleWebGL/>
        </Suspense>
      </div>
      <div class="animate-slide-bottom">website is under construction. check back later</div>
    </div>
  </div>
</template>

<style scoped>

</style>