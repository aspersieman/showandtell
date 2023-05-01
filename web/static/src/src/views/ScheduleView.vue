<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useRoute } from 'vue-router'

import StopWatch from '@/components/StopWatch.vue'
import { useScheduleStore } from '@/stores/schedules'
import type { ApiSchedule, ApiSpeaker } from '@/models/models'
import dayjs from '@/utils/dayjs'

const scheduleStore = useScheduleStore()

const route = useRoute()
const schedule = ref<ApiSchedule | null>(null)
const minutesIncrement = ref(0)

interface Speaker extends ApiSpeaker {
  startTime: string
  endTime: string
}
const speakers = ref<Speaker[]>([])

const getSchedule = (id: number) => {
  return scheduleStore.getSchedule(id)
}
const getTimeEnd = (startDate: Date | null | undefined, minutes: number | null): string => {
  if (!startDate || minutes === undefined) return ''
  minutes = minutes ? minutes : 0
  minutesIncrement.value += minutes
  return dayjs(startDate).add(minutesIncrement.value, 'minute').format('HH:mm')
}

onMounted(async () => {
  const scheduleId = parseInt(route.params.id[0])
  schedule.value = await getSchedule(scheduleId)

  schedule.value?.speakers.forEach((speaker) => {
    speakers.value.push({
      ...speaker,
      startTime: getTimeEnd(schedule.value?.start_date_time, 0),
      endTime: getTimeEnd(schedule.value?.start_date_time, speaker.minutes)
    })
  })
})
</script>
<template>
  <main class="pt-8 pb-16 lg:pt-16 lg:pb-24 mb-auto bg-white dark:bg-gray-900">
    <h1 class="text-3xl text-center font-extrabold leading-tight tracking-tight text-gray-900 dark:text-white">
      {{ schedule?.title }}
    </h1>
    <h1 class="text-1xl text-center font-medium leading-tight tracking-tight text-gray-900 dark:text-white">
      {{ schedule?.description }}
    </h1>
    <section class="bg-white dark:bg-gray-900 antialiased mt-6">
      <div class="max-w-screen-xl px-4 pb-8 mx-auto">
        <div class="max-w-3xl mx-auto text-center">
          <h2 class="text-2xl font-extrabold leading-tight tracking-tight text-gray-900 dark:text-white">
            Agenda
          </h2>
        </div>
        <div class="flow-root max-w-3xl mx-auto mt-8">
          <div class="-my-4 divide-y divide-gray-200 dark:divide-gray-700 sm:text-md md:text-lg lg:text-xl">
            <div v-for="(speaker, index) in speakers"
              class="grid sm:grid-cols-[95px_auto_150px] md:grid-cols-[125px_auto_150px] items-center gap-2 py-4 sm:gap-6"
              :key="index">
              <div
                class="justify-self-start font-normal whitespace-nowrap text-left text-gray-500 sm:text-right dark:text-gray-400">
                {{ speaker.startTime }} -
                {{ speaker.endTime }}
                <div class="text-sm text-left text-gray-500 dark:text-gray-400">
                  {{ speaker.name }}
                </div>
              </div>
              <h3 class="justify-self-start font-semibold text-gray-900 dark:text-white">
                <div>{{ speaker.topic }}</div>
              </h3>
              <p class="justify-self-end">
                <stop-watch :max="speaker.minutes" channel="speaker" :socket-id="speaker.id" />
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>
