<script setup>
import {ref} from "vue";
import {TresCanvas, useRenderLoop} from "@tresjs/core";
import {useGLTF} from "@tresjs/cientos";
import {Color} from 'three'

const meatball = await useGLTF("/meatball.glb", { draco: true })
meatball.scene.traverse((node) => {
  if (!node.isMesh) return
  node.material.wireframe = true

  if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
    node.material.color = new Color('white')
  }
})

meatball.scene.scale.set(1, 1, 1)
meatball.scene.rotation.set(0, 1.57, 0)
meatball.scene.position.set(0, 0, 0)

const sceneRef = ref()
const rotationSpeed = Math.random() * (4 - 1) + 1
const random = Math.random()

if (random >= 0.9 || (random < 0.9 && random >= 0.5)) {
  useRenderLoop().onLoop(({delta}) => {
    sceneRef.value.rotation.y -= (delta * rotationSpeed)
  })
}

if (random >= 0.9 || (random < 0.5 && random >= 0)) {
  useRenderLoop().onLoop(({delta}) => {
    sceneRef.value.rotation.z -= (delta * rotationSpeed)
  })
}
</script>

<template>
  <div class="w-[360px] h-[140px]">
    <TresCanvas
        alpha
        :antialias="false"
    >
      <TresPerspectiveCamera
          :position="[8, 0, 0]"
          :look-at="[0, 0, 0]"
          :fov="10"
      />
      <TresAmbientLight
          :intensity="10"
          :color="0xffffff"
      />
      <TresScene ref="sceneRef">
        <Suspense>
          <primitive :object="meatball.scene" />
        </Suspense>
      </TresScene>
    </TresCanvas>
  </div>
</template>

<style scoped>

</style>