<script setup lang="ts">
import { onMounted, ref, onBeforeUnmount } from 'vue'
import StartSound from '@/assets/start.mp3'
import StopSound from '@/assets/stop.mp3'
import { getApiBaseUrl } from '@/utils/api'

const props = defineProps({
  max: { type: Number, default: 0 },
  channel: String,
  socketId: Number,
  showControls: { type: Boolean, default: false }
})

const emit = defineEmits<{
  (e: 'start', val: number): void
  (e: 'stop', val: Object): void
  (e: 'pause', val: Object): void
}>()

const time = ref('')
const isRunning = ref(false)
const startTime = ref(0)
const times = ref([0, 0, 0, 0])
const frameId = ref<number | null>(null)
const progress = ref(0)
const progressColour = ref('bg-blue-600')

const sendMessages = () => {
  if (isRunning.value && props.socketId) {
    sendMessage({
      socketId: props.socketId,
      time: time.value,
      progress: progress.value
    })
  }
}
let messagesInterval = setInterval(sendMessages, 1000)

interface SpeakerSocketData {
  socketId: number
  time: string
  progress: number
}
const eventsFromServer = ref<SpeakerSocketData[]>([])

const serverUrl = getApiBaseUrl('ws') + '/ws/' + props.socketId
const socket = new WebSocket(serverUrl)

const sendMessage = (data: SpeakerSocketData) => {
  socket.send(JSON.stringify(data))
}

const pause = () => {
  isRunning.value = !isRunning.value
  if (isRunning.value) {
    if (time.value === '00:00:00') {
      sound('start')
    }
    startTime.value = performance.now()
    calculate(performance.now())
    frameId.value = requestAnimationFrame(step)
  }
  emit('pause', { running: isRunning.value, time: time.value })
}

const reset = () => {
  startTime.value = 0
  isRunning.value = false
  times.value = [0, 0, 0, 0]
  time.value = formatTimes()
  progress.value = 0
  progressColour.value = 'bg-blue-600'
}

const pad0 = (value: number, count: number) => {
  let result = value.toString()
  while (result.length < count) {
    result = '0' + result
    --count
  }
  return result
}

const formatTimes = () => {
  const hours = pad0(times.value[0], 2)
  const minutes = pad0(times.value[1], 2)
  const seconds = pad0(times.value[2], 2)
  return `${hours}:${minutes}:${seconds}`
}

const step = (timestamp: number) => {
  if (!isRunning.value) return
  calculate(timestamp)
  startTime.value = timestamp
  time.value = formatTimes()
  frameId.value = requestAnimationFrame(step)
}

const calculate = (timestamp: number) => {
  const diff = timestamp - startTime.value
  times.value[3] += diff / 10
  if (times.value[3] >= 100) {
    times.value[3] -= 100
    times.value[2] += 1
  }
  if (times.value[2] >= 60) {
    times.value[2] -= 60
    times.value[1] += 1
  }
  if (times.value[1] >= 60) {
    times.value[1] -= 60
    times.value[0] += 1
  }
  const seconds = times.value[1] * 60 + times.value[2]
  const minutes = seconds / 60
  progress.value = (minutes / props.max) * 100
  if (progress.value >= 100 || minutes >= props.max) {
    progress.value = 100
  }
  if (progress.value <= 60) {
    progressColour.value = 'bg-blue-600'
  } else if (progress.value <= 80) {
    progressColour.value = 'bg-orange-300'
  } else if (progress.value < 100) {
    progressColour.value = 'bg-red-600'
  } else if (progress.value >= 100) {
    sound('stop')
    pause()
  }
}

interface SoundDictionary {
  name: string
  file: string
}
const sounds: SoundDictionary[] = [
  {
    name: 'start',
    file: StartSound
  },
  {
    name: 'stop',
    file: StopSound
  }
]
const sound = (soundType: string) => {
  const s = sounds.find((s) => s.name === soundType)
  if (s) {
    const audio = new Audio(s.file)
    audio.play()
  }
}

const stopwatchSocket = (socketId: number) => {
  socket.addEventListener('open', () => {
    console.log('stopwatchSocket connected : ' + socketId)
  })

  socket.addEventListener('message', function (event) {
    if (socketId) {
      let eventData: SpeakerSocketData = {
        socketId: socketId,
        time: '',
        progress: 0
      }
      try {
        eventData = JSON.parse(event.data)
      } catch (e) {
        console.log(e)
      }
      console.log('Message from server ', event.data, isRunning.value)
      if (eventData.socketId == socketId && !isRunning.value) {
        time.value = eventData.time
        progress.value = eventData.progress
        eventsFromServer.value.push({
          socketId: eventData?.socketId,
          time: new Date().toString(),
          progress: eventData.progress
        })
      }
    }
  })

  return { eventsFromServer }
}

onMounted(() => {
  reset()
  if (props.socketId) {
    stopwatchSocket(props.socketId)
  }
})
onBeforeUnmount(() => {
  clearInterval(messagesInterval)
})
</script>

<template>
  <div v-if="max > 0" class="w-full bg-gray-200 rounded-full h-1.5 mb-4 dark:bg-gray-700">
    <div
      :class="`${progressColour} h-1.5 rounded-full dark:bg-blue-500`"
      :style="`width: ${progress}%;`"
    ></div>
  </div>
  <div class="flex justify-between">
    <div class="font-mono time mr-1">{{ time }}</div>
    <button
      v-if="!isRunning && props.showControls"
      @click="pause()"
      type="button"
      class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
    >
      <svg xmlns="http://www.w3.org/2000/svg" height="16" viewBox="0 96 960 960" width="16">
        <path d="M320 853V293l440 280-440 280Zm60-280Zm0 171 269-171-269-171v342Z" />
      </svg>
    </button>
    <button
      v-if="isRunning && props.showControls"
      @click="pause()"
      type="button"
      class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
    >
      <svg xmlns="http://www.w3.org/2000/svg" height="16" viewBox="0 96 960 960" width="16">
        <path
          d="M525 856V296h235v560H525Zm-325 0V296h235v560H200Zm385-60h115V356H585v440Zm-325 0h115V356H260v440Zm0-440v440-440Zm325 0v440-440Z"
        />
      </svg>
    </button>
    <button
      v-if="props.showControls"
      @click="reset()"
      type="button"
      class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-xs px-1 py-1 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700"
    >
      <svg xmlns="http://www.w3.org/2000/svg" height="16" viewBox="0 96 960 960" width="16">
        <path
          d="M477 936q-149 0-253-105.5T120 575h60q0 125 86 213t211 88q127 0 215-89t88-216q0-124-89-209.5T477 276q-68 0-127.5 31T246 389h105v60H142V241h60v106q52-61 123.5-96T477 216q75 0 141 28t115.5 76.5Q783 369 811.5 434T840 574q0 75-28.5 141t-78 115Q684 879 618 907.5T477 936Zm128-197L451 587V373h60v189l137 134-43 43Z"
        />
      </svg>
    </button>
  </div>
</template>

<style scoped></style>
