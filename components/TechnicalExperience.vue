<script setup>
import TEBlob from "../../../components/blobs/TEBlob.vue";

const blobs = [{
  skill: "Golang"
}, {
  skill: "Fiber"
}, {
  skill: "Java"
}, {
  skill: "Kotlin (Native)"
}, {
  skill: "Spring Boot"
}, {
  skill: "C"
}, {
  skill: "C++"
}, {
  skill: "LibHTTP"
}, {
  skill: "Rust"
}, {
  skill: "Rocket"
}, {
  skill: "Dart"
}, {
  skill: "Flutter"
}, {
  skill: "Python"
}, {
  skill: "FastAPI"
}, {
  skill: "Flask"
}, {
  skill: "Swift"
}, {
  skill: "Scala"
}, {
  skill: "PHP"
}, {
  skill: "Laravel"
}, {
  skill: "Ruby"
}, {
  skill: "Rails"
}, {
  skill: "HTML5"
}, {
  skill: "JavaScript"
}, {
  skill: "TypeScript"
}, {
  skill: "Node.js"
}, {
  skill: "Express.js"
}, {
  skill: "Vue.js"
}, {
  skill: "Nuxt"
}, {
  skill: "React (Native)"
}, {
  skill: "Next.js"
}, {
  skill: "jQuery"
}, {
  skill: "HTMX"
}, {
  skill: "alpine.js"
}, {
  skill: "CSS3"
}, {
  skill: "TailwindCSS"
}, {
  skill: "Bootstrap"
}, {
  skill: "Git"
}, {
  skill: "Subversion"
}, {
  skill: "Docker"
}, {
  skill: "Podman"
}, {
  skill: "Vulkan"
}, {
  skill: "Metal"
}, {
  skill: "OpenGL (ES)"
}, {
  skill: "GLFW"
}, {
  skill: "LWJGL"
}, {
  skill: "SDL2"
}, {
  skill: "MySQL"
}, {
  skill: "PostgreSQL"
}, {
  skill: "GraphQL"
}, {
  skill: "MongoDB"
}, {
  skill: "Redis"
}, {
  skill: "memcached"
}, {
  skill: "Ubuntu Linux"
}, {
  skill: "RHEL"
}, {
  skill: "Bash"
}, {
  skill: "Zsh"
}, {
  skill: "Postfix"
}, {
  skill: "Dovecot"
}, {
  skill: "SpamAssassin"
}, {
  skill: "OpenDKIM"
}, {
  skill: "NGINX"
}, {
  skill: "ZFS"
}, {
  skill: "KVM/QEMU"
}, {
  skill: "Windows"
}, {
  skill: "PowerShell"
}, {
  skill: "Active Directory"
}, {
  skill: "Microsoft IIS"
}, {
  skill: "Microsoft Azure"
}, {
  skill: "AWS"
}, {
  skill: "GCP"
}, {
  skill: "Vercel"
}, {
  skill: "DigitalOcean"
}, {
  skill: "VMware ESXi"
}, {
  skill: "Hyper-V"
}, {
  skill: "Proxmox"
}, {
  skill: "Stripe"
}, {
  skill: "PayPal"
}, {
  skill: "MacOS"
}, {
  skill: "iOS"
}, {
  skill: "Android"
}, {
  skill: "watchOS"
}, {
  skill: "iPadOS"
}, {
  skill: "Xcode"
}, {
  skill: "JetBrains"
},]

</script>

<template>
  <div class="grid grid-cols-2 md:grid-cols-4 gap-2 place-content-center font-display not-prose">
    <t-e-blob v-for="blob in blobs" :blob="blob" />
  </div>
</template>

<style scoped>

</style>