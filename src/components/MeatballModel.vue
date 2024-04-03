<script setup>
import {useGLTF} from "@tresjs/cientos";

const {scene} = await useGLTF("/sean.gltf", { draco: true })
scene.traverse((node) => {
  if (!node.isMesh) return
  node.material.wireframe = true
})

scene.scale.set(1, 1, 1)
scene.rotation.set(0, 1.57, 0)
scene.position.set(0, 0, 0)
</script>

<template>
  <primitive :object="scene" />
</template>

<style scoped>

</style>