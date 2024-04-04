<script setup>
import {defineAsyncComponent} from "vue";

const TitleWebGL = defineAsyncComponent(() =>
    import('./TitleWebGL.vue')
)
</script>

<template>
  <div class=
           "w-screen h-screen
            grid grid-cols-1 lg:grid-cols-3
            justify-center content-center
            bg-neutral-100 dark:bg-stone-900"
  >
    <div class="will-change-transform my-auto mx-auto animate-fade-in lg:col-span-3">
      <Suspense>
        <TitleWebGL />
      </Suspense>
    </div>
    <div class="will-change-transform my-auto mx-auto animate-slide-bottom lg:col-span-3">meatball</div>
  </div>
</template>

<style scoped>

</style>