<script setup lang="ts">
import { onMounted, ref } from 'vue'

import ScheduleModal from '@/components/ScheduleModal.vue'
import SpeakerModal from '@/components/SpeakerModal.vue'
import { useScheduleStore } from '@/stores/schedules'
import { useSpeakerStore } from '@/stores/speakers'
import type { ApiSchedule, ScheduleSearchParams } from '@/models/models'
import dayjs from '@/utils/dayjs'

const scheduleStore = useScheduleStore()
const speakerStore = useSpeakerStore()

const showScheduleModal = ref(false)
const showSpeakerModal = ref(false)
const showAll = ref(false)

const inFuture = (date: Date) => {
  const d = dayjs(date)
  const today = dayjs()
  return d.isSameOrAfter(today)
}
const deleteSchedule = (id: number) => {
  if (window.confirm('Are your sure you want to delete this schedule?')) {
    scheduleStore.deleteSchedule(id).then(() => {
      getSchedules(true)
    })
  }
}
const addSchedule = () => {
  scheduleStore.scheduleEdit = null
  showScheduleModal.value = true
}
const editSchedule = (schedule: ApiSchedule) => {
  scheduleStore.scheduleEdit = schedule
  showScheduleModal.value = true
}

const toggleScheduleModal = (show: boolean) => {
  showScheduleModal.value = show
}

const toggleSpeakerModal = (show: boolean) => {
  showSpeakerModal.value = show
}
const addSpeaker = (schedule: ApiSchedule) => {
  speakerStore.setSpeakerAddSchedule(schedule)
  toggleSpeakerModal(true)
}
const editSpeaker = () => {
  showSpeakerModal.value = true
}

const speakerDeleted = () => {
  getSchedules(true)
  if (scheduleStore.scheduleEdit) {
    scheduleStore.scheduleEdit = scheduleStore.schedulesAdmin.find((schedule) => {
      return schedule.id === scheduleStore?.scheduleEdit?.id
    }) as ApiSchedule
  }
}

const getSchedules = (reloadHome: boolean = false) => {
  const params: ScheduleSearchParams = {}
  if (showAll.value === true) {
    const now = dayjs()
    params.f = now.subtract(100, 'year').format('YYYY-MM-DD')
    params.t = now.add(100, 'year').format('YYYY-MM-DD')
  }
  scheduleStore.getSchedulesAdmin(params).then(() => {
    if (scheduleStore.scheduleEdit) {
      scheduleStore.scheduleEdit = scheduleStore.schedulesAdmin.find((schedule) => {
        return schedule.id === scheduleStore?.scheduleEdit?.id
      }) as ApiSchedule
    }
  })
  if (reloadHome) {
    scheduleStore.getSchedulesHome()
  }
}
onMounted(() => {
  getSchedules()
})
</script>

<template>
  <div>
    <schedule-modal :show-modal="showScheduleModal" @speaker-deleted="speakerDeleted()" @speaker-edit="editSpeaker()"
      @added="getSchedules(true)" @updated="getSchedules(true)" @toggle="toggleScheduleModal(false)" />
    <speaker-modal :show-modal="showSpeakerModal" @added="getSchedules(true)" @toggle="toggleSpeakerModal(false)" />
    <main class="pt-8 pb-16 lg:pt-16 lg:pb-24 bg-white dark:bg-gray-900">
      <div class="px-4 mx-auto max-w-screen-xl">
        <section
          class="flex mx-auto w-full max-w-2xl format format-sm sm:format-base lg:format-lg format-blue dark:format-invert mb-3">
          <button type="button" @click="addSchedule()"
            class="flex mx-1 text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-4 py-2 text-center mr-3 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800">
            <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 96 960 960" width="20">
              <path fill="white"
                d="M180 976q-24 0-42-18t-18-42V296q0-24 18-42t42-18h65v-60h65v60h340v-60h65v60h65q24 0 42 18t18 42v620q0 24-18 42t-42 18H180Zm0-60h600V486H180v430Zm0-490h600V296H180v130Zm0 0V296v130Zm300 230q-17 0-28.5-11.5T440 616q0-17 11.5-28.5T480 576q17 0 28.5 11.5T520 616q0 17-11.5 28.5T480 656Zm-160 0q-17 0-28.5-11.5T280 616q0-17 11.5-28.5T320 576q17 0 28.5 11.5T360 616q0 17-11.5 28.5T320 656Zm320 0q-17 0-28.5-11.5T600 616q0-17 11.5-28.5T640 576q17 0 28.5 11.5T680 616q0 17-11.5 28.5T640 656ZM480 816q-17 0-28.5-11.5T440 776q0-17 11.5-28.5T480 736q17 0 28.5 11.5T520 776q0 17-11.5 28.5T480 816Zm-160 0q-17 0-28.5-11.5T280 776q0-17 11.5-28.5T320 736q17 0 28.5 11.5T360 776q0 17-11.5 28.5T320 816Zm320 0q-17 0-28.5-11.5T600 776q0-17 11.5-28.5T640 736q17 0 28.5 11.5T680 776q0 17-11.5 28.5T640 816Z" />
            </svg>
            New Schedule
          </button>
          <div class="flex items-center">
            <input v-model="showAll" @change="getSchedules()" id="show-all" type="checkbox"
              class="ml-2 w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600" />
            <label for="show-all" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Show all</label>
          </div>
        </section>

        <section
          class="mx-auto w-full max-w-2xl format format-sm sm:format-base lg:format-lg format-blue dark:format-invert">
          <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
            <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
              <thead class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400">
                <tr>
                  <th scope="col" class="px-6 py-3">Title</th>
                  <th scope="col" class="px-6 py-3">Speakers</th>
                  <th scope="col" class="px-6 py-3">Start</th>
                  <th scope="col" class="px-6 py-3">End</th>
                  <th scope="col" class="px-6 py-3">&nbsp;</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="(schedule, index) in scheduleStore.schedulesAdmin"
                  class="bg-white border-b dark:bg-gray-800 dark:border-gray-700" :key="index">
                  <th scope="row" class="px-6 py-4 font-medium text-gray-900 dark:text-white">
                    {{ schedule.title }}
                  </th>
                  <td class="px-3 py-2 text-center">{{ schedule.speakers.length }}</td>
                  <td class="px-3 py-2 text-xs">
                    {{ new Date(schedule.start_date_time).toDateString() }} @{{
                      new Date(schedule.start_date_time).toLocaleTimeString()
                    }}
                  </td>
                  <td class="px-3 py-2 text-xs">
                    {{ new Date(schedule.end_date_time).toDateString() }} @{{
                      new Date(schedule.end_date_time).toLocaleTimeString()
                    }}
                  </td>
                  <td class="px-3 py-2">
                    <button @click="editSchedule(schedule)" :data-tooltip-target="`tooltip-edit-{index}`" type="button"
                      class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
                      <svg xmlns="http://www.w3.org/2000/svg" height="16" viewBox="0 96 960 960" width="16">
                        <path
                          d="M180 876h44l443-443-44-44-443 443v44Zm614-486L666 262l42-42q17-17 42-17t42 17l44 44q17 17 17 42t-17 42l-42 42Zm-42 42L248 936H120V808l504-504 128 128Zm-107-21-22-22 44 44-22-22Z" />
                      </svg>
                    </button>
                    <div :id="`tooltip-edit-{index}`" role="tooltip"
                      class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700">
                      Edit
                      <div class="tooltip-arrow" data-popper-arrow></div>
                    </div>
                    <button @click="deleteSchedule(schedule.id)" :data-tooltip-target="`tooltip-delete-{index}`"
                      type="button"
                      class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
                      <svg xmlns="http://www.w3.org/2000/svg" height="16" viewBox="0 96 960 960" width="16">
                        <path
                          d="M261 936q-24.75 0-42.375-17.625T201 876V306h-41v-60h188v-30h264v30h188v60h-41v570q0 24-18 42t-42 18H261Zm438-630H261v570h438V306ZM367 790h60V391h-60v399Zm166 0h60V391h-60v399ZM261 306v570-570Z" />
                      </svg>
                    </button>
                    <div :id="`tooltip-delete-{index}`" role="tooltip"
                      class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700">
                      Delete
                      <div class="tooltip-arrow" data-popper-arrow></div>
                    </div>
                    <button v-if="inFuture(schedule.start_date_time)" @click="addSpeaker(schedule)"
                      :data-tooltip-target="`tooltip-speaker-{index}`" type="button"
                      class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
                      <svg xmlns="http://www.w3.org/2000/svg" height="16" viewBox="0 96 960 960" width="16">
                        <path
                          d="M730 656V526H600v-60h130V336h60v130h130v60H790v130h-60Zm-370-81q-66 0-108-42t-42-108q0-66 42-108t108-42q66 0 108 42t42 108q0 66-42 108t-108 42ZM40 896v-94q0-35 17.5-63.5T108 696q75-33 133.338-46.5 58.339-13.5 118.5-13.5Q420 636 478 649.5 536 663 611 696q33 15 51 43t18 63v94H40Zm60-60h520v-34q0-16-9-30.5T587 750q-71-33-120-43.5T360 696q-58 0-107.5 10.5T132 750q-15 7-23.5 21.5T100 802v34Zm260-321q39 0 64.5-25.5T450 425q0-39-25.5-64.5T360 335q-39 0-64.5 25.5T270 425q0 39 25.5 64.5T360 515Zm0-90Zm0 411Z" />
                      </svg>
                    </button>
                    <div :id="`tooltip-speaker-{index}`" role="tooltip"
                      class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700">
                      Add speaker
                      <div class="tooltip-arrow" data-popper-arrow></div>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </section>
      </div>
    </main>
  </div>
</template>
