<script setup>
const certs = [
    {
        organization: "Miami University Oxford",
        url: "https://miamioh.edu",
        location: "Oxford, OH",
        start: 2018,
        end: 2020,
        description: ["Associate's degree", "Major in Information Technology", "Minor in Media & Culture, Music Theory"]
    }
]
</script>

<template>
    <h2>Education & Certifications</h2>
    <div class="flex flex-col not-prose">
        <div v-for="cert in certs" class="p-2 flex flex-col gap-2">
            <span class="text-left">
                <a v-if="cert.url" class="underline font-bold text-lg" :href="cert.url" target="_blank">
                    {{ cert.organization }}
                </a>
                <span v-else class="font-bold text-lg">{{ cert.organization }}</span>
                <span v-if="cert.location" class="ml-2">{{ cert.location }}</span>
                <span class="ml-2 text-sm text-neutral-600 dark:text-neutral-400">
                    {{ cert.start }}
                    <span v-if="cert.end && (cert.end !== cert.start)">&mdash; {{ cert.end }}</span>
                    <span v-else-if="cert.end"></span>
                    <span v-else>&mdash; Present</span>
                </span>
            </span>
            <span v-for="desc in cert.description"
                class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">
                {{ desc }}
            </span>
        </div>
    </div>
</template>

<style lang="css" scoped></style>