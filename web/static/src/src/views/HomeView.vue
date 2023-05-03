<script setup lang="ts">
import ScheduleTimeline from '../components/ScheduleTimeline.vue'
</script>

<template>
  <main class="pt-8 pb-16 lg:pt-16 lg:pb-24 bg-white dark:bg-gray-900">
    <div class="flex justify-between px-4 mx-auto max-w-screen-xl">
      <section
        class="mx-auto w-full max-w-2xl format format-sm sm:format-base lg:format-lg format-blue dark:format-invert">
        <ScheduleTimeline />
      </section>
    </div>
  </main>
</template>
