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
      <div class="animate-slide-bottom">meatball</div>
    </div>
    <div class=
             "w-screen h-screen py-4 md:p-4
              grid grid-cols-1 md:grid-cols-2 gap-4 md:gap-96 place-content-center md:place-content-around z-20">
      <div v-for="n in 4"
           class=
               "p-2 px-4
              md:rounded-lg bg-neutral-200 dark:bg-stone-800
              shadow shadow-neutral-300 dark:shadow-stone-800"
      >
        <h2 class="font-semibold">test</h2>
        <h3>{test</h3>
      </div>
    </div>
  </div>
</template>

<style scoped>

</style>