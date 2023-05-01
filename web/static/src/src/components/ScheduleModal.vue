<script setup lang="ts">
import { ref, watch } from 'vue'
import VueTailwindDatepicker from 'vue-tailwind-datepicker'
import { VueDraggableNext } from 'vue-draggable-next'

import { useScheduleStore } from '@/stores/schedules'
import { useSpeakerStore } from '@/stores/speakers'
import type { ApiSchedule, ApiSpeaker, ScheduleAdd, ScheduleUpdate } from '@/models/models'
import dayjs from '@/utils/dayjs'

const props = defineProps({
  showModal: Boolean
})
const emit = defineEmits<{
  (e: 'toggle', val: Boolean): void
  (e: 'added', val: Boolean): void
  (e: 'updated', val: Boolean): void
  (e: 'speakerDeleted', val: Boolean): void
  (e: 'speakerEdit', val: ApiSpeaker): void
}>()

const scheduleStore = useScheduleStore()
const speakerStore = useSpeakerStore()

const title = ref('')
const description = ref(
  'Please wait until the speaker has completed their presentation for any remarks/questions.'
)
const startDate = ref('')
const startTime = ref('15:00')
const endTime = ref('16:00')

const mode = ref('add')

const clear = () => {
  title.value = ''
  description.value = ''
  startDate.value = ''
  startTime.value = '15:00'
  endTime.value = '16:00'
  mode.value = 'add'
  scheduleStore.scheduleEdit = null
}

const close = () => {
  clear()
  emit('toggle', false)
}

const validateSave = () => {
  if (title.value === '') {
    alert('Title is required')
    return false
  }
  if (description.value === '') {
    alert('Description is required')
    return false
  }
  if (startDate.value === '') {
    alert('Start date is required')
    return false
  }
  if (startTime.value === '') {
    alert('Start time is required')
    return false
  }
  if (endTime.value === '') {
    alert('End time is required')
    return false
  }
  const start = new Date()
  start.setHours(parseInt(startTime.value.split(':')[0]))
  start.setMinutes(parseInt(startTime.value.split(':')[1]))
  const end = new Date()
  end.setHours(parseInt(endTime.value.split(':')[0]))
  end.setMinutes(parseInt(endTime.value.split(':')[1]))
  if (start > end) {
    alert('Start time must be before end time')
    return false
  }

  return true
}

const deleteSpeaker = (id: number) => {
  speakerStore
    .deleteSpeaker(id)
    .then(() => {
      emit('speakerDeleted', true)
    })
    .then(() => {
      if (scheduleStore.scheduleEdit) {
        scheduleStore.scheduleEdit.speakers = scheduleStore.scheduleEdit?.speakers.filter(
          (speaker) => speaker.id !== id
        )
      }
    })
}
const editSpeaker = (speaker: ApiSpeaker) => {
  speakerStore.setSpeakerAddSchedule(scheduleStore.scheduleEdit as ApiSchedule)
  speakerStore.speakerEdit = speaker
  emit('speakerEdit', speaker)
}
const orderChanged = async () => {
  let order = 1
  scheduleStore?.scheduleEdit?.speakers.forEach((speaker) => {
    speaker.order = order
    speakerStore.updateSpeaker(speaker)
    order++
  })

  emit('added', true)
}

const addSchedule = () => {
  if (validateSave()) {
    const start = new Date(startDate.value + ' ' + startTime.value)
    const end = new Date(startDate.value + ' ' + endTime.value)
    const schedule: ScheduleAdd = {
      title: title.value,
      description: description.value,
      start_date_time: start,
      end_date_time: end,
      speakers: null
    }
    scheduleStore.addSchedule(schedule).then(() => {
      clear()
      emit('toggle', false)
      emit('added', true)
    })
  }
}

const updateSchedule = () => {
  if (validateSave()) {
    const start = new Date(startDate.value + ' ' + startTime.value)
    const end = new Date(startDate.value + ' ' + endTime.value)
    const scheduleId = scheduleStore.scheduleEdit?.id
    console.log('updateSchedule: ', scheduleId)
    if (scheduleId) {
      const schedule: ScheduleUpdate = {
        id: scheduleId,
        title: title.value,
        description: description.value,
        start_date_time: start,
        end_date_time: end,
        speakers: null
      }
      scheduleStore.updateSchedule(schedule).then(() => {
        clear()
        emit('toggle', false)
        emit('added', true)
      })
    }
  }
}
const formatter = ref({
  date: 'YYYY-MM-DD',
  month: 'MMM'
})
const dDate = (date: Date) => {
  const yesterday = new Date()
  yesterday.setDate(yesterday.getDate() - 1)
  return date < yesterday
}

watch(
  () => props.showModal,
  (newValue) => {
    if (newValue && scheduleStore.scheduleEdit) {
      title.value = scheduleStore.scheduleEdit.title
      description.value = scheduleStore.scheduleEdit.description
      startDate.value = dayjs(scheduleStore.scheduleEdit.start_date_time).format('YYYY-MM-DD')
      startTime.value = dayjs(scheduleStore.scheduleEdit.start_date_time).format('HH:mm')
      endTime.value = dayjs(scheduleStore.scheduleEdit.end_date_time).format('HH:mm')
      mode.value = 'edit'
    }
  }
)
</script>
<template>
  <div
    v-if="props.showModal"
    id="schedule-add-modal"
    tabindex="-1"
    class="fixed top-0 left-0 right-0 z-50 w-full p-4 overflow-x-hidden overflow-y-auto md:inset-0 h-[calc(100%)] max-h-full bg-gray-400/50"
  >
    <div class="relative w-full max-w-md max-h-full m-auto">
      <!-- Modal content -->
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
        <button
          @click="close()"
          type="button"
          class="absolute top-3 right-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm p-1.5 ml-auto inline-flex items-center dark:hover:bg-gray-800 dark:hover:text-white"
        >
          <svg
            aria-hidden="true"
            class="w-5 h-5"
            fill="currentColor"
            viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              fill-rule="evenodd"
              d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
              clip-rule="evenodd"
            ></path>
          </svg>
          <span class="sr-only">Close modal</span>
        </button>
        <div class="px-6 py-6 lg:px-8">
          <h3 v-if="mode === 'add'" class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
            Add a new schedule
          </h3>
          <h3 v-if="mode === 'edit'" class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
            Edit schedule
          </h3>
          <div class="space-y-2">
            <div>
              <label
                for="title"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Title
              </label>
              <input
                type="text"
                v-model="title"
                id="title"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="Title"
                required
              />
            </div>
            <div>
              <label
                for="description"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Description
              </label>
              <textarea
                id="description"
                v-model="description"
                rows="4"
                class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="Description"
              ></textarea>
            </div>
            <div>
              <label for="" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Start date</label
              >
              <div class="block max-w-sm">
                <vue-tailwind-datepicker
                  v-model="startDate"
                  :formatter="formatter"
                  :disable-date="dDate"
                  input-classes="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                  overlay
                  placeholder="Start date (YYYY-MM-DD)"
                  as-single
                />
              </div>
            </div>
            <div>
              <label
                for="email"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Time</label
              >
              <div class="flex justify-between">
                <input
                  v-model="startTime"
                  type="time"
                  class="time-picker bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                />
                <small class="flex self-center align-middle">To</small>
                <input
                  v-model="endTime"
                  type="time"
                  class="time-picker bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                />
              </div>
            </div>
            <div v-if="scheduleStore.scheduleEdit?.speakers">
              <div class="relative overflow-x-auto shadow-md sm:rounded-lg">
                <table class="w-full text-sm text-left text-gray-500 dark:text-gray-400">
                  <caption
                    class="p-3 text-md font-semibold text-left text-gray-900 bg-white dark:text-white dark:bg-gray-800"
                  >
                    Speakers
                  </caption>
                  <thead
                    class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
                  >
                    <tr>
                      <th scope="col" class="px-3 py-1">Name</th>
                      <th scope="col" class="px-3 py-1">Topic</th>
                      <th scope="col" class="px-3 py-1">Order</th>
                      <th scope="col" class="px-3 py-1">&nbsp;</th>
                    </tr>
                  </thead>
                  <vue-draggable-next
                    tag="tbody"
                    @change="orderChanged"
                    v-model="scheduleStore.scheduleEdit.speakers"
                  >
                    <tr
                      v-for="(speaker, index) in scheduleStore.scheduleEdit?.speakers"
                      :key="index"
                      class="handle bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
                    >
                      <td
                        scope="row"
                        class="px-3 py-1 font-medium text-gray-900 whitespace-nowrap dark:text-white"
                      >
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          height="16"
                          viewBox="0 96 960 960"
                          width="16"
                          class="drag-icon absolute"
                        >
                          <path
                            d="M349.911 896Q321 896 300.5 875.411q-20.5-20.588-20.5-49.5Q280 797 300.589 776.5q20.588-20.5 49.5-20.5Q379 756 399.5 776.589q20.5 20.588 20.5 49.5Q420 855 399.411 875.5q-20.588 20.5-49.5 20.5Zm260 0Q581 896 560.5 875.411q-20.5-20.588-20.5-49.5Q540 797 560.589 776.5q20.588-20.5 49.5-20.5Q639 756 659.5 776.589q20.5 20.588 20.5 49.5Q680 855 659.411 875.5q-20.588 20.5-49.5 20.5Zm-260-250Q321 646 300.5 625.411q-20.5-20.588-20.5-49.5Q280 547 300.589 526.5q20.588-20.5 49.5-20.5Q379 506 399.5 526.589q20.5 20.588 20.5 49.5Q420 605 399.411 625.5q-20.588 20.5-49.5 20.5Zm260 0Q581 646 560.5 625.411q-20.5-20.588-20.5-49.5Q540 547 560.589 526.5q20.588-20.5 49.5-20.5Q639 506 659.5 526.589q20.5 20.588 20.5 49.5Q680 605 659.411 625.5q-20.588 20.5-49.5 20.5Zm-260-250Q321 396 300.5 375.411q-20.5-20.588-20.5-49.5Q280 297 300.589 276.5q20.588-20.5 49.5-20.5Q379 256 399.5 276.589q20.5 20.588 20.5 49.5Q420 355 399.411 375.5q-20.588 20.5-49.5 20.5Zm260 0Q581 396 560.5 375.411q-20.5-20.588-20.5-49.5Q540 297 560.589 276.5q20.588-20.5 49.5-20.5Q639 256 659.5 276.589q20.5 20.588 20.5 49.5Q680 355 659.411 375.5q-20.588 20.5-49.5 20.5Z"
                          />
                        </svg>
                        {{ speaker.name }}
                      </td>
                      <td class="px-3 py-1">{{ speaker.topic }}</td>
                      <td class="px-3 py-1">{{ speaker.order }}</td>
                      <td class="px-3 py-1 text-right">
                        <button
                          @click="editSpeaker(speaker)"
                          :data-tooltip-target="`tooltip-edit-{index}`"
                          type="button"
                          class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
                        >
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            height="16"
                            viewBox="0 96 960 960"
                            width="16"
                          >
                            <path
                              d="M180 876h44l443-443-44-44-443 443v44Zm614-486L666 262l42-42q17-17 42-17t42 17l44 44q17 17 17 42t-17 42l-42 42Zm-42 42L248 936H120V808l504-504 128 128Zm-107-21-22-22 44 44-22-22Z"
                            />
                          </svg>
                        </button>
                        <div
                          :id="`tooltip-edit-{index}`"
                          role="tooltip"
                          class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700"
                        >
                          Edit
                          <div class="tooltip-arrow" data-popper-arrow></div>
                        </div>
                        <button
                          @click="deleteSpeaker(speaker.id)"
                          :data-tooltip-target="`tooltip-delete-{index}`"
                          type="button"
                          class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
                        >
                          <svg
                            xmlns="http://www.w3.org/2000/svg"
                            height="16"
                            viewBox="0 96 960 960"
                            width="16"
                          >
                            <path
                              d="M261 936q-24.75 0-42.375-17.625T201 876V306h-41v-60h188v-30h264v30h188v60h-41v570q0 24-18 42t-42 18H261Zm438-630H261v570h438V306ZM367 790h60V391h-60v399Zm166 0h60V391h-60v399ZM261 306v570-570Z"
                            />
                          </svg>
                        </button>
                        <div
                          :id="`tooltip-delete-{index}`"
                          role="tooltip"
                          class="absolute z-10 invisible inline-block px-3 py-2 text-sm font-medium text-white transition-opacity duration-300 bg-gray-900 rounded-lg shadow-sm opacity-0 tooltip dark:bg-gray-700"
                        >
                          Delete
                          <div class="tooltip-arrow" data-popper-arrow></div>
                        </div>
                      </td>
                    </tr>
                  </vue-draggable-next>
                </table>
              </div>
            </div>
            <button
              v-if="scheduleStore.scheduleEdit !== null"
              @click="updateSchedule()"
              type="submit"
              class="w-full text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Update
            </button>
            <button
              v-else
              @click="addSchedule()"
              type="submit"
              class="w-full text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
            >
              Save
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.drag-icon {
  cursor: move;
  left: -2px;
}
</style>
