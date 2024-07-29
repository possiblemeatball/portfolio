<script setup>
import {ref} from "vue";
import {TresCanvas, useRenderLoop} from "@tresjs/core";
import {GLTFLoader} from "three/examples/jsm/loaders/GLTFLoader.js";
import {Color} from 'three'

const sceneRef = ref()
const loader = new GLTFLoader().setPath("/")
loader.load("meatball.glb", (gltf) => {
  const meatball = gltf.scene
  meatball.traverse((node) => {
    if (!node.isMesh) return
    node.material.wireframe = true

    if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
      node.material.color = new Color('white')
    }
  })

  meatball.scale.set(1, 1, 1)
  meatball.rotation.set(0, 1.57, 0)
  meatball.position.set(0, 0, 0)
  sceneRef.value.add(meatball)
})

const rotationSpeed = Math.random() * (3 - 0.5) + 0.5
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
          :position="[7, 0, 0]"
          :up="[0, 1, 0]"
          :look-at="[0, 0, 0]"
          :fov="10"
      />
      <TresAmbientLight
          :intensity="10"
          :color="0xffffff"
      />
      <TresScene ref="sceneRef" />
    </TresCanvas>
  </div>
</template>

<style scoped>

</style>