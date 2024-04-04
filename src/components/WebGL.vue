<script setup>
import {TresCanvas, useRenderLoop, useTresContext} from "@tresjs/core";
import MeatballModel from "./MeatballModel.vue";
import {ref} from "vue";

const sceneRef = ref()
const randSpeed = 0.02;
const { onLoop } = useRenderLoop()

onLoop(({delta, elapsed, clock}) => {
  //sceneRef.value.rotation.y += randSpeed
  sceneRef.value.rotation.z -= randSpeed
})
</script>

<template>
    <TresCanvas
        alpha
        power-preference="high-performance"
        :antialias="false"
    >
      <TresPerspectiveCamera
          :position="[12, 0, 0]"
          :look-at="[0, 0, 0]"
          :aspect="innerWidth / innerHeight"
          :fov="13"
          :near="0.005"
          :far="10000"
      />
      <TresAmbientLight
          :intensity="10"
          :color="0xffffff"
      />
      <TresScene ref="sceneRef">
        <Suspense>
          <MeatballModel />
        </Suspense>
      </TresScene>
    </TresCanvas>
</template>

<style scoped>
</style>