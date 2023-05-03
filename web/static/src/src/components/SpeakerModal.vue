<script setup lang="ts">
import { ref, watch } from 'vue'
import { useSpeakerStore } from '@/stores/speakers'

export interface Props {
  showModal?: Boolean
  shedule: Object | null
}
const props = defineProps({
  showModal: Boolean
})
const emit = defineEmits<{
  (e: 'toggle', val: Boolean): void
  (e: 'added', val: Boolean): void
}>()

const speakerStore = useSpeakerStore()

const mode = ref('add')

const name = ref('')
const topic = ref('')
const description = ref('')
const audience = ref('')
const minutes = ref<number>(15)
const isPrivate = ref(true)

const clear = () => {
  name.value = ''
  topic.value = ''
  description.value = ''
  audience.value = ''
  minutes.value = 15
  isPrivate.value = true
}

const validateSave = () => {
  if (name.value === '') {
    alert('Name is required')
    return false
  }
  if (topic.value === '') {
    alert('Topic is required')
    return false
  }
  if (description.value === '') {
    alert('Description is required')
    return false
  }
  if (audience.value === '') {
    alert('Audience is required')
    return false
  }
  return true
}
const addSpeaker = () => {
  if (validateSave()) {
    speakerStore
      .addSpeaker({
        name: name.value,
        topic: topic.value,
        description: description.value,
        audience: audience.value,
        order: 1,
        minutes: parseInt(String(minutes.value)),
        schedule_id: speakerStore.speakerAddSchedule?.id,
        private: isPrivate.value
      })
      .then(() => {
        emit('toggle', false)
        emit('added', true)
        clear()
      })
  }
}
const updateSpeaker = () => {
  if (validateSave()) {
    const speakerId = speakerStore.speakerEdit?.id
    const order = speakerStore.speakerEdit?.order
    const scheduleId = speakerStore.speakerEdit?.schedule_id
    if (speakerId && order && scheduleId) {
      speakerStore
        .updateSpeaker({
          id: speakerId,
          name: name.value,
          topic: topic.value,
          description: description.value,
          audience: audience.value,
          order: order,
          minutes: parseInt(String(minutes.value)),
          schedule_id: scheduleId,
          private: isPrivate.value
        })
        .then(() => {
          emit('toggle', false)
          emit('added', true)
          clear()
        })
    }
  }
}
watch(
  () => props.showModal,
  (newValue) => {
    if (newValue && speakerStore.speakerEdit) {
      name.value = speakerStore.speakerEdit.name
      topic.value = speakerStore.speakerEdit.topic
      description.value = speakerStore.speakerEdit.description
      audience.value = speakerStore.speakerEdit.audience
      minutes.value = speakerStore.speakerEdit.minutes
      isPrivate.value = speakerStore.speakerEdit.private
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
          @click="emit('toggle', false)"
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
            Add a new speaker
          </h3>
          <h3 v-if="mode === 'edit'" class="mb-4 text-xl font-medium text-gray-900 dark:text-white">
            Edit speaker
          </h3>
          <h4 class="mb-4 text-md font-medium text-gray-900 dark:text-white">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              height="20"
              viewBox="0 96 960 960"
              width="20"
              class="float-left mr-1"
            >
              <path
                d="M180 976q-24 0-42-18t-18-42V296q0-24 18-42t42-18h65v-60h65v60h340v-60h65v60h65q24 0 42 18t18 42v620q0 24-18 42t-42 18H180Zm0-60h600V486H180v430Zm0-490h600V296H180v130Zm0 0V296v130Zm300 230q-17 0-28.5-11.5T440 616q0-17 11.5-28.5T480 576q17 0 28.5 11.5T520 616q0 17-11.5 28.5T480 656Zm-160 0q-17 0-28.5-11.5T280 616q0-17 11.5-28.5T320 576q17 0 28.5 11.5T360 616q0 17-11.5 28.5T320 656Zm320 0q-17 0-28.5-11.5T600 616q0-17 11.5-28.5T640 576q17 0 28.5 11.5T680 616q0 17-11.5 28.5T640 656ZM480 816q-17 0-28.5-11.5T440 776q0-17 11.5-28.5T480 736q17 0 28.5 11.5T520 776q0 17-11.5 28.5T480 816Zm-160 0q-17 0-28.5-11.5T280 776q0-17 11.5-28.5T320 736q17 0 28.5 11.5T360 776q0 17-11.5 28.5T320 816Zm320 0q-17 0-28.5-11.5T600 776q0-17 11.5-28.5T640 736q17 0 28.5 11.5T680 776q0 17-11.5 28.5T640 816Z"
              />
            </svg>
            {{ speakerStore.speakerAddSchedule.title }},&nbsp;{{
              new Date(speakerStore.speakerAddSchedule.start_date_time).toDateString()
            }}&nbsp;@&nbsp;{{
              new Date(speakerStore.speakerAddSchedule.start_date_time).toLocaleTimeString()
            }}
          </h4>
          <div class="space-y-2">
            <div>
              <label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Speaker name
              </label>
              <input
                type="text"
                v-model="name"
                id="name"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="Speaker name"
                required
              />
            </div>
            <div>
              <label
                for="topic"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Topic
              </label>
              <input
                type="text"
                v-model="topic"
                id="topic"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="Speaker topic"
                required
              />
            </div>
            <div>
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
            </div>
            <div>
              <label
                for="audience"
                class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                >Audience
              </label>
              <input
                type="text"
                v-model="audience"
                id="audience"
                class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-500 focus:border-primary-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
                placeholder="Intended audience"
                required
              />
            </div>
            <div>
              <div class="flex items-center">
                <input
                  v-model="isPrivate"
                  id="private"
                  type="checkbox"
                  class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
                />
                <label
                  for="private"
                  class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300"
                  >Private (Internal only)</label
                >
              </div>
            </div>
            <div>
              <div>
                <label
                  for="duration-range"
                  class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >Duration</label
                >
                <input
                  v-model="minutes"
                  id="duration-range"
                  type="range"
                  min="5"
                  max="60"
                  step="5"
                  class="w-full h-2 bg-gray-200 rounded-lg appearance-none cursor-pointer dark:bg-gray-700"
                />
                <label class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
                  >{{ minutes }} minutes</label
                >
              </div>
              <button
                v-if="speakerStore.speakerEdit !== null"
                @click="updateSpeaker()"
                type="submit"
                class="w-full text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:outline-none focus:ring-primary-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-800"
              >
                Update
              </button>
              <button
                v-else
                @click="addSpeaker()"
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
  </div>
</template>
