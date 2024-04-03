<script setup>
import {TresCanvas} from "@tresjs/core";
import MeatballModel from "./MeatballModel.vue";
</script>

<template>
    <TresCanvas
        alpha
        preset="realistic"
        power-preference="high-performance"
        :antialias="false"
    >
      <TresPerspectiveCamera
        :position="[4, 0, 0]"
        :up="[0, 1, 0]"
        :look-at="[0, 0, 0]"
      />
      <TresAmbientLight
          :intensity="1"
          :color="0xffffff"
      />
      <Suspense>
        <MeatballModel />
      </Suspense>
    </TresCanvas>
</template>

<style scoped>
</style>