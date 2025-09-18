<script setup>
import {useData} from 'vitepress'
import MeatballWebGl from '../../components/MeatballWebGL.vue'
import Projects from "../../components/Projects.vue";
import Employment from "../../components/Employment.vue";
import TechnicalExperience from "../../components/TechnicalExperience.vue";

// https://vitepress.dev/reference/runtime-api#usedata
const {site} = useData()
</script>

<template>
  <div class="flex flex-col gap-0 justify-center items-center overflow-x-hidden">
    <div class="p-4 pt-20 w-screen flex flex-col justify-center items-center text-center
                select-none pointer-events-none bg-white dark:bg-neutral-900">
      <Suspense>
        <MeatballWebGl class="animate-slide-top"/>
      </Suspense>
      <h1 class="font-display text-2xl">{{ site.title }}</h1>
      <p class="font-sans text-neutral-600 dark:text-neutral-400">{{ site.description }}</p>
    </div>
    <div class="container max-w-screen-sm p-4 text-justify">
      <div class="prose prose-h3:font-display prose-p:font-sans max-w-none prose-neutral dark:prose-invert">
        <p>Over 12 years experience in software engineering, systems administration, computer networking, cybersecurity,
          and general information technology, 6 years experience with client communications, customer service and
          support technician work, and 5 years of entrepreneurship.</p>
        <p>Previously self-employed with <a href="https://nevertrustllc.com" target="_blank">Never Trust LLC</a> since
          2020 until 2025. We provided support to small-business owners in building a cloud-computing business network
          capable of sustaining heavy compute workloads and network traffic.</p>
        <p>Currently, seeking employment opportunities in related fields.</p>
        <h3>Projects</h3>
        <Projects/>
        <h3>Employment</h3>
        <Employment/>
        <h3>Education</h3>
        <div class="p-2 flex flex-col gap-2">
          <span class="text-left">
            <a class="font-display font-bold text-lg" href="https://miamioh.edu"
               target="_blank">Miami University Oxford</a>
            <span class="ml-2 font-sans font-normal">Oxford, OH</span>
            <span class="ml-2 text-sm text-neutral-600 dark:text-neutral-400">2018 &mdash; 2020</span>
          </span>
          <span class="font-mono indent-8 text-sm text-neutral-600 dark:text-neutral-400">Associate's degree, maj. Information Technology min. Media & Culture, Music Theory</span>
        </div>
        <h3>Technical Experience</h3>
        <TechnicalExperience/>
        <h3>Links</h3>
        <span class="flex flex-row justify-between pb-20">
          <a class="font-mono text-sm" href="https://github.com/possiblemeatball" target="_blank">GitHub</a>
          <a class="font-mono text-sm" href="https://youtube.com/@possiblemeatball" target="_blank">YouTube</a>
          <a class="font-mono text-sm" href="https://t.me/possiblemeatball" target="_blank">Telegram</a>
          <a class="font-mono text-sm" href="mailto:meatball@manthrowshat.net" target="_blank">E-mail</a>
          <a class="font-mono text-sm" href="https://manthrowshat.net/tip" target="_blank">Tip Me</a>
        </span>
      </div>
    </div>
  </div>
</template>
