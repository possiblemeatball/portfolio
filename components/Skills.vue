<script setup>
const skills = {
  "Software Engineering, IT & Cybersecurity": {
    "Frameworks": {
      "Fiber": 3,
      "Kotlin Multiplatform": 3,
      "Spring/Boot": 10,
      "Rocket": 2,
      "Flutter": 5,
      "Flask": 10,
      "FastAPI": 6,
      "Laravel": 10,
      "Ruby on Rails": 3,
      "Node": 12,
      "Express": 12,
      "React": 11,
      "React Native": 6,
      "Vue": 8,
      "Nuxt": 4,
      "Next": 4,
    },
    "Languages": {
      "Golang": 6,
      "Java": 15,
      "Kotlin/Native": 12,
      "C": 10,
      "C++": 9,
      "Rust": 5,
      "Dart": 5,
      "Python": 10,
      "Swift": 4,
      "Scala": 12,
      "PHP": 14,
      "Ruby": 3,
      "HTML/CSS/JS": 15,
      "TypeScript": 6,
    },
    "Libraries": {
      "LibHTTP": 4,
      "jQuery": 15,
      "HTMX": 4,
      "alpine.js": 4,
      "TailwindCSS": 6,
      "Bootstrap": 14,
      "LWJGL": 14,
      "GLFW": 15,
      "SDL1/2/3": 14,
    }
  },
  "Systems Administration & Networking": {
    "Automation and Scripting": {
      "Ansible": 3,
      "Bash": 15,
      "PowerShell": 6,
      "Zsh": 9
    },
    "Cloud Service Providers": {
      "Amazon Web Services": 10,
      "Google Cloud Platform": 8,
      "Microsoft Azure": 9,
      "DigitalOcean": 14,
      "Vercel": 3,
      "Stripe": 11,
      "PayPal": 14,
    },
    "Databases & Filesystems": {
      "MySQL": 14,
      "PostgreSQL": 8,
      "Redis": 7,
      "memcached": 7,
      "MongoDB": 6,
      "GraphQL": 5,
      "ZFS": 12,
      "Btrfs": 6,
    },
    "Operating Systems": {
      "Debian GNU/Linux": 15,
      "Ubuntu Linux": 14,
      "Arch Linux": 12,
      "Red Hat Enterprise Linux": 8,
      "Microsoft Windows": 15,
      "Apple macOS": 10,
      "Apple iOS/iPadOS/watchOS": 9,
      "Android": 11,
      "Android wearOS": 8
    },
    "Server Protocols & Software Suites": {
      "Postfix": 15,
      "Dovecot": 15,
      "NGINX": 14,
      "SpamAssassin": 14,
      "OpenDKIM": 12,
      "Active Directory": 7,
      "OpenLDAP": 6,
      "Microsoft IIS": 5,
      "RspamD": 4,
      "Git": 12,
      "Subversion": 15,
      "JetBrains": 12,
    },
    "Virtualization Frameworks": {
      "KVM/QEMU": 10,
      "Docker": 10,
      "Proxmox/PVE": 9,
      "VMWare ESX": 6,
      "Hyper-V": 6,
      "Podman": 5
    },
  },
  "Client Communications & Customer Support": {
    "Active Listening": {
      "Hear what they say": 0, "Ask questions to confirm & understand": 0, "Respond with empathy & action": 0
    },
    "Basic Human Emotions": {
      "Empathy": 0, "Patience": 0, "Tenacity": 0, "Kindness": 0, "Respect": 0
    }
  }
}

</script>

<template>
  <h2>Technical Experience</h2>
  <div class="flex flex-col gap-4 select-none pointer-events-none">
    <div v-for="subCategory, title in skills">
      <h3>{{ title }}</h3>
      <div v-for="skill, subTitle in subCategory">
        <h4>{{ subTitle }}</h4>
        <div v-for="experience, name in skill" class="inline-block mr-2 font-mono">
          <span class="border px-1">
            {{ name }}
            <sub v-if="experience > 0" class="top-0">{{ experience }} years</sub>
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped></style>