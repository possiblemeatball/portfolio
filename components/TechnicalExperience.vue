<script setup>
import TEBlob from "./blobs/TEBlob.vue";

const blobs = [
  {
    icon: null,
    skill: "C"
  },
  {
    icon: null,
    skill: "LibHTTP"
  },
  {
    icon: null,
    skill: "ImGui"
  },
  {
    icon: null,
    skill: "Golang"
  },
  {
    icon: null,
    skill: "Fiber"
  },
  {
    icon: null,
    skill: "Java"
  },
  {
    icon: null,
    skill: "Kotlin (Native)"
  },
  {
    icon: null,
    skill: "Spring Boot"
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
    skill: "egui"
  },
  {
    icon: null,
    skill: "Rocket"
  },
  {
    icon: null,
    skill: "Python"
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
    skill: "Swift"
  },
  {
    icon: null,
    skill: "Scala"
  },
  {
    icon: null,
    skill: "PHP"
  },
  {
    icon: null,
    skill: "Laravel"
  },
  {
    icon: null,
    skill: "JavaScript"
  },
  {
    icon: null,
    skill: "TypeScript"
  },
  {
    icon: null,
    skill: "jQuery"
  },
  {
    icon: null,
    skill: "Node.js"
  },
  {
    icon: null,
    skill: "Express.js"
  },
  {
    icon: null,
    skill: "Next.js"
  },
  {
    icon: null,
    skill: "Vue.js"
  },
  {
    icon: null,
    skill: "React (Native)"
  },
  {
    icon: null,
    skill: "AngularJS"
  },
  {
    icon: null,
    skill: "Ruby"
  },
  {
    icon: null,
    skill: "Rails"
  },
  {
    icon: null,
    skill: "TailwindCSS"
  },
  {
    icon: null,
    skill: "Git"
  },
  {
    icon: null,
    skill: "Subversion"
  },
  {
    icon: null,
    skill: "Docker"
  },
  {
    icon: null,
    skill: "Podman"
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
    skill: "GraphQL"
  },
  {
    icon: null,
    skill: "MongoDB"
  },
  {
    icon: null,
    skill: "Redis-compat"
  },
  {
    icon: null,
    skill: "memcached"
  },
  {
    icon: null,
    skill: "Arch Linux"
  },
  {
    icon: null,
    skill: "RHEL"
  },
  {
    icon: null,
    skill: "Bash"
  },
  {
    icon: null,
    skill: "zsh"
  },
  {
    icon: null,
    skill: "fish"
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
    skill: "NGINX"
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
    skill: "Windows"
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
    skill: "MacOS"
  },
  {
    icon: null,
    skill: "iOS"
  },
  {
    icon: null,
    skill: "Android"
  },
  {
    icon: null,
    skill: "watchOS"
  },
  {
    icon: null,
    skill: "iPadOS"
  },
  {
    icon: null,
    skill: "Xcode"
  },
  {
    icon: null,
    skill: "JetBrains"
  },
  {
    icon: null,
    skill: "AWS"
  },
  {
    icon: null,
    skill: "Microsoft Azure"
  },
  {
    icon: null,
    skill: "GCP"
  },
  {
    icon: null,
    skill: "DigitalOcean"
  },
]

</script>

<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-2 place-content-center font-display not-prose">
    <t-e-blob v-for="blob in blobs" :skill="blob.skill" :icon="blob.icon" />
  </div>
</template>

<style scoped>

</style>