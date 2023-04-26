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
  <main class="pt-8 pb-16 lg:pt-16 lg:pb-24 mb-auto bg-white dark:bg-gray-900">
    <div class="flex justify-between px-4 mx-auto max-w-screen-xl">
      <div class="flex justify-between w-full border border-gray-300 rounded-md p-2 mx-2">
        <div if="state && !finished" class="flex justify-center text-xl self-center">Countdown</div>
        <vue3-flip-countdown v-if="!state && !finished" deadline="2023-04-24 08:30:00" :show-labels="false"
          :show-days="false" :show-hours="false" :stop="state" @time-elapsed="elapsed" />
        <div class="flex justify-center mx-auto text-xl self-center" v-else-if="state && !finished">
          Timer Paused
        </div>
        <div class="flex justify-center mx-auto text-xl self-center text-amber-500" v-if="finished">
          Timer Finished!
        </div>
        <button v-if="!finished" type="button" @click="state = !state"
          class="text-gray-900 bg-white border flex self-center justify-center border-gray-300 focus:outline-none hover:bg-gray-100 focus:ring-4 focus:ring-gray-200 font-medium rounded-lg text-sm px-5 py-2.5 mr-2 mb-2 dark:bg-gray-800 dark:text-white dark:border-gray-600 dark:hover:bg-gray-700 dark:hover:border-gray-600 dark:focus:ring-gray-700">
          {{ state ? 'Start' : 'Stop' }}
        </button>
      </div>
    </div>
    <section class="bg-white dark:bg-gray-900 antialiased">
      <div class="max-w-screen-xl px-4 py-8 mx-auto lg:px-6 sm:py-16 lg:py-24">
        <div class="max-w-3xl mx-auto text-center">
          <h2 class="text-4xl font-extrabold leading-tight tracking-tight text-gray-900 dark:text-white">
            Agenda
          </h2>

          <div class="mt-4">
            <a href="#" title=""
              class="inline-flex items-center text-lg font-medium text-primary-600 hover:underline dark:text-primary-500">
              Learn more about our agenda
              <svg aria-hidden="true" class="w-5 h-5 ml-2" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"
                fill="currentColor">
                <path fill-rule="evenodd"
                  d="M10.293 3.293a1 1 0 011.414 0l6 6a1 1 0 010 1.414l-6 6a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-4.293-4.293a1 1 0 010-1.414z"
                  clip-rule="evenodd" />
              </svg>
            </a>
          </div>
        </div>

        <div class="flow-root max-w-3xl mx-auto mt-8 sm:mt-12 lg:mt-16">
          <div class="-my-4 divide-y divide-gray-200 dark:divide-gray-700">
            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                08:00 - 09:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Opening remarks</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                09:00 - 10:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Bergside LLC: Controlling the video traffic flows</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                10:00 - 11:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Flowbite - An Open Framework for Forensic Watermarking</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                11:00 - 12:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Coffee Break</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                12:00 - 13:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Scaling your brand from €0 to multimillion euros</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                13:00 - 14:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Updates from the Open Source Multimedia community</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                14:00 - 15:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Exploring the balance between customer acquisition and retention</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                15:00 - 16:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Flowbite - An Open Framework for Forensic Watermarking</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                16:00 - 14:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Scaling your brand from €0 to multimillion euros</a>
              </h3>
            </div>

            <div class="flex flex-col gap-2 py-4 sm:gap-6 sm:flex-row sm:items-center">
              <p class="w-32 text-lg font-normal text-gray-500 sm:text-right dark:text-gray-400 shrink-0">
                17:00 - 14:00
              </p>
              <h3 class="text-lg font-semibold text-gray-900 dark:text-white">
                <a href="#" class="hover:underline">Drinks & networking</a>
              </h3>
            </div>
          </div>
        </div>
      </div>
    </section>
  </main>
</template>
