<script setup>
import {ref} from "vue";
import {TresCanvas, useRenderLoop} from "@tresjs/core";
import {GLTFLoader} from "three/examples/jsm/loaders/GLTFLoader.js";
import {Color} from 'three'

const sceneRef = ref()
function setMeatballColor({ matches }) {
  const scene = sceneRef.value
  if (!scene) return
  scene.traverse((node) => {
    if (!node.isMesh) return
    if (matches) {
      node.material.color = new Color('white')
    } else {
      node.material.color = new Color('black')
    }
  })
}

const loader = new GLTFLoader().setPath("/")
loader.load("meatball.glb", (gltf) => {
  const meatball = gltf.scene
  meatball.traverse((node) => {
    if (!node.isMesh) return
    node.material.wireframe = true
  })

  meatball.scale.set(1, 1, 1)
  meatball.rotation.set(0, 1.57, 0)
  meatball.position.set(0, 0, 0.09)
  sceneRef.value.add(meatball)
  if (window.matchMedia) {
    setMeatballColor(window.matchMedia('(prefers-color-scheme: dark)'))
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', setMeatballColor)
  }
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
  <div class="w-[511px] h-[128px]">
    <TresCanvas
        alpha
        :antialias="false"
    >
      <TresPerspectiveCamera
          :position="[5, 0, 0]"
          :up="[0, 1, 0]"
          :look-at="[0, 0, 0]"
          :fov="13"
          :aspect="3080 / 1080"
          :near="0.005"
          :far="10000"
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