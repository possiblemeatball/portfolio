<script setup>
import {defineAsyncComponent} from "vue";

const TitleWebGL = defineAsyncComponent(() =>
    import('./TitleWebGL.vue')
)
</script>

<template>
  <div class=
           "w-screen h-screen
            grid justify-center place-content-center
            bg-neutral-50 dark:bg-stone-950"
  >
    <div>
      <div class="animate-slide-top">
        <Suspense>
          <TitleWebGL/>
        </Suspense>
      </div>
      <div class="animate-slide-bottom text-center">meatball</div>
    </div>
  </div>
</template>

<style scoped>

</style>