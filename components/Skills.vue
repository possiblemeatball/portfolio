<script setup>
const skills = {
  "Software Engineering, IT & Cybersecurity": {
    "Frameworks": {
      "Spring/Boot": 13,
      "Node": 12,
      "Express": 12,
      "React": 11,
      "Flask": 10,
      "Laravel": 10,
      "Vue": 8,
      "FastAPI": 6,
      "React Native": 6,
      "Next": 4,
      "Nuxt": 4,
      "Fiber": 3,
      "Kotlin Multiplatform": 3,
      "Flutter": 3,
      "Ruby on Rails": 3,
      "Rocket": 2,
    },
    "Languages": {
      "HTML/CSS/JS": 16,
      "Java": 16,
      "PHP": 14,
      "Scala": 12,
      "Kotlin/Native": 12,
      "C": 10,
      "Python": 10,
      "C++": 9,
      "Go": 6,
      "TypeScript": 6,
      "Rust": 4,
      "Swift": 4,
      "Dart": 3,
      "Ruby": 3,
    },
    "Libraries": {
      "jQuery": 15,
      "GLFW": 15,
      "SDL1/2/3": 14,
      "LWJGL": 14,
      "Bootstrap": 14,
      "TailwindCSS": 6,
      "LibHTTP": 4,
      "HTMX": 4,
      "alpine.js": 4,
    }
  },
  "Systems Administration & Networking": {
    "Automation and Scripting": {
      "Bash": 15,
      "Zsh": 9,
      "PowerShell": 6,
      "Ansible": 3,
    },
    "Cloud Service Providers": {
      "DigitalOcean": 14,
      "PayPal": 14,
      "Stripe": 11,
      "Amazon Web Services": 10,
      "Microsoft Azure": 9,
      "Google Cloud Platform": 8,
      "Vercel": 3,
    },
    "Databases & Filesystems": {
      "MySQL": 14,
      "ZFS": 12,
      "PostgreSQL": 8,
      "Redis": 7,
      "memcached": 7,
      "MongoDB": 6,
      "Btrfs": 6,
      "GraphQL": 5,
    },
    "Operating Systems": {
      "Debian GNU/Linux": 15,
      "Microsoft Windows": 15,
      "Ubuntu Linux": 14,
      "Arch Linux": 12,
      "Android": 11,
      "Apple macOS": 10,
      "Apple iOS/iPadOS/watchOS": 9,
      "Red Hat Enterprise Linux": 8,
      "Android wearOS": 8
    },
    "Server Protocols & Software Suites": {
      "Postfix": 15,
      "Dovecot": 15,
      "Subversion": 15,
      "NGINX": 14,
      "SpamAssassin": 14,
      "JetBrains": 12,
      "Git": 12,
      "OpenDKIM": 12,
      "Active Directory": 7,
      "OpenLDAP": 6,
      "Microsoft IIS": 5,
      "RspamD": 4,
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