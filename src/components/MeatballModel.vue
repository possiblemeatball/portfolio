<script setup>
import {useGLTF} from "@tresjs/cientos";
import * as THREE from "three";

const {scene} = await useGLTF("/meatball.glb", { draco: true })
scene.traverse((node) => {
  if (!node.isMesh) return
  node.material.wireframe = true
})

scene.scale.set(1, 1, 1)
scene.rotation.set(0, 1.57, 0)
scene.position.set(-2, 0, 3)
</script>

<template>
  <primitive :object="scene" />
</template>

<style scoped>

</style>