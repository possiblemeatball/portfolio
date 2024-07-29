<script setup>
import TEBlob from "./blobs/TEBlob.vue";

const blobs = [
  {
    icon: null,
    skill: "C"
  },
  {
    icon: null,
    skill: "Golang"
  },
  {
    icon: null,
    skill: "Java"
  },
  {
    icon: null,
    skill: "Kotlin"
  },
  {
    icon: null,
    skill: "C++"
  },
  {
    icon: null,
    skill: "Rust"
  },
  {
    icon: null,
    skill: "Python"
  },
  {
    icon: null,
    skill: "Swift"
  },
  {
    icon: null,
    skill: "Scala"
  },
  {
    icon: null,
    skill: "JavaScript"
  },
  {
    icon: null,
    skill: "Nodejs"
  },
  {
    icon: null,
    skill: "Express.js"
  },
  {
    icon: null,
    skill: "Vue.js"
  },
  {
    icon: null,
    skill: "React"
  },
  {
    icon: null,
    skill: "Angular"
  },
  {
    icon: null,
    skill: "FastAPI"
  },
  {
    icon: null,
    skill: "Flask"
  },
  {
    icon: null,
    skill: "Django"
  },
  {
    icon: null,
    skill: "Fiber"
  },
  {
    icon: null,
    skill: "TailwindCSS"
  },
  {
    icon: null,
    skill: "MySQL"
  },
  {
    icon: null,
    skill: "PostgreSQL"
  },
  {
    icon: null,
    skill: "MongoDB"
  },
  {
    icon: null,
    skill: "Redis/ValKey"
  },
  {
    icon: null,
    skill: "Postfix"
  },
  {
    icon: null,
    skill: "Dovecot"
  },
  {
    icon: null,
    skill: "SpamAssassin"
  },
  {
    icon: null,
    skill: "OpenDKIM"
  },
  {
    icon: null,
    skill: "ZFS"
  },
  {
    icon: null,
    skill: "KVM/QEMU"
  },
  {
    icon: null,
    skill: "Hyper-V"
  },
  {
    icon: null,
    skill: "Active Directory"
  },
  {
    icon: null,
    skill: "Android"
  },
  {
    icon: null,
    skill: "iOS"
  },
  {
    icon: null,
    skill: "MacOS"
  },
  {
    icon: null,
    skill: "Windows"
  },
  {
    icon: null,
    skill: "Arch Linux"
  },
  {
    icon: null,
    skill: "RHEL"
  }
]

</script>

<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-2 place-content-center">
    <t-e-blob v-for="blob in blobs" :skill="blob.skill" :icon="blob.icon" />
  </div>
</template>

<style scoped>

</style>