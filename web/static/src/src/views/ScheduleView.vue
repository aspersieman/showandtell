<script setup lang="ts">
import { ref } from 'vue'
const state = ref(false)
const finished = ref(false)

const elapsed = () => {
  console.log('Finished!')
  state.value = false
  finished.value = true
}
</script>
<template>
  <main class="pt-8 pb-16 lg:pt-16 lg:pb-24 bg-white dark:bg-gray-900">
    <div class="flex justify-between px-4 mx-auto max-w-screen-xl">
      <div class="flex justify-between w-full border border-gray-300 rounded-md p-2 mx-2">
        <vue3-flip-countdown v-if="!state && !finished" deadline="2023-04-23 18:40:00" :show-labels="false"
          :show-days="false" :show-hours="false" :stop="state" @time-elapsed="elapsed" />
        <div class="flex justify-center mx-auto text-xl self-center justify-center" v-else-if="state && !finished">
          Timer Paused
        </div>
        <div class="flex justify-center mx-auto text-xl self-center justify-center text-amber-500" v-if="finished">
          Timer Finished!
        </div>
        <button v-if="!finished" type="button" @click="state = !state"
          class="text-gray-900 bg-white border border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
          {{ state ? 'Start' : 'Stop' }}
        </button>
      </div>
    </div>
  </main>
</template>
