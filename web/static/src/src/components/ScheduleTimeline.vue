<script setup lang="ts">
import ProfileAvatar from 'vue-profile-avatar'
import { RouterLink } from 'vue-router'
import { useScheduleStore } from '@/stores/schedules'
import dayjs from '@/utils/dayjs'

const scheduleStore = useScheduleStore()

const isToday = (d: Date) => {
  const today = dayjs()
  const date = dayjs(d)
  return date.format('YYYY-MM-DD') === today.format('YYYY-MM-DD')
}
</script>

<template>
  <h1 v-if="scheduleStore.schedulesHome.length > 0"
    class="text-2xl font-semibold text-gray-900 dark:text-white text-center mb-1">
    Upcoming Sessions
  </h1>
  <div v-else>
    <h1 class="text-2xl font-semibold text-gray-900 dark:text-white text-center mb-1">
      No sessions yet
      <RouterLink
        class="align-middle text-white mx-1 bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:outline-none focus:ring-gray-300 font-medium rounded-lg text-sm px-4 py-2 text-center mr-3 md:mr-0 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-800"
        to="/schedules">Add a new Session
      </RouterLink>
    </h1>
  </div>
  <div v-for="(schedule, index) in scheduleStore.schedulesHome"
    class="p-5 pb-1 mb-4 border border-gray-100 rounded-lg bg-gray-50 dark:bg-gray-800 dark:border-gray-700" :key="index">
    <div class="relative w-full border-b pb-3">
      <span v-if="isToday(schedule.start_date_time)"
        class="bg-primary-100 text-primary-800 text-xs font-medium mr-1 px-2.5 py-0.5 rounded dark:bg-primary-900 dark:text-primary-300">Today</span>
      <RouterLink :to="`/schedule`">
        <time class="text-lg font-semibold text-gray-900 dark:text-white">{{ new
          Date(schedule.start_date_time).toDateString() }} @{{
    new Date(schedule.start_date_time).toLocaleTimeString()
  }}</time>
      </RouterLink>
      <RouterLink
        class="text-white float-right mx-1 bg-gray-700 hover:bg-gray-800 focus:ring-4 focus:outline-none focus:ring-gray-300 font-medium rounded-lg text-sm px-4 py-2 text-center mr-3 md:mr-0 dark:bg-gray-600 dark:hover:bg-gray-700 dark:focus:ring-gray-800"
        to="/schedule">View
      </RouterLink>
      <div class="flex mt-1">
        <svg xmlns="http://www.w3.org/2000/svg" height="20" viewBox="0 96 960 960" width="20">
          <path
            d="M180 976q-24 0-42-18t-18-42V296q0-24 18-42t42-18h65v-60h65v60h340v-60h65v60h65q24 0 42 18t18 42v620q0 24-18 42t-42 18H180Zm0-60h600V486H180v430Zm0-490h600V296H180v130Zm0 0V296v130Zm300 230q-17 0-28.5-11.5T440 616q0-17 11.5-28.5T480 576q17 0 28.5 11.5T520 616q0 17-11.5 28.5T480 656Zm-160 0q-17 0-28.5-11.5T280 616q0-17 11.5-28.5T320 576q17 0 28.5 11.5T360 616q0 17-11.5 28.5T320 656Zm320 0q-17 0-28.5-11.5T600 616q0-17 11.5-28.5T640 576q17 0 28.5 11.5T680 616q0 17-11.5 28.5T640 656ZM480 816q-17 0-28.5-11.5T440 776q0-17 11.5-28.5T480 736q17 0 28.5 11.5T520 776q0 17-11.5 28.5T480 816Zm-160 0q-17 0-28.5-11.5T280 776q0-17 11.5-28.5T320 736q17 0 28.5 11.5T360 776q0 17-11.5 28.5T320 816Zm320 0q-17 0-28.5-11.5T600 776q0-17 11.5-28.5T640 736q17 0 28.5 11.5T680 776q0 17-11.5 28.5T640 816Z" />
        </svg>
        {{ schedule.title }}
      </div>
    </div>
    <ol class="mt-1 divide-y divider-gray-200 dark:divide-gray-700">
      <li v-for="(speaker, speakerIndex) in schedule.speakers" :key="speakerIndex">
        <div class="items-center px-3 pt-2 pb-1 sm:flex">
          <ProfileAvatar class="w-12 h-12 mr-1 self-start float-left rounded-full sm:mb-0 small speaker-profile"
            :username="speaker.name">
          </ProfileAvatar>
          <div class="text-gray-600 dark:text-gray-400 w-full self-end">
            <div class="md:flex text-base font-normal md:justify-between">
              <span class="flex md:justify-start w-auto">
                <svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 96 960 960" width="24">
                  <path
                    d="M480 575q-66 0-108-42t-42-108q0-66 42-108t108-42q66 0 108 42t42 108q0 66-42 108t-108 42ZM160 896v-94q0-38 19-65t49-41q67-30 128.5-45T480 636q62 0 123 15.5t127.921 44.694q31.301 14.126 50.19 40.966Q800 764 800 802v94H160Zm60-60h520v-34q0-16-9.5-30.5T707 750q-64-31-117-42.5T480 696q-57 0-111 11.5T252 750q-14 7-23 21.5t-9 30.5v34Zm260-321q39 0 64.5-25.5T570 425q0-39-25.5-64.5T480 335q-39 0-64.5 25.5T390 425q0 39 25.5 64.5T480 515Zm0-90Zm0 411Z" />
                </svg>
                <span class="max-md:text-md max-sm:text-xs whitespace-nowrap mr-1">
                  <span class="font-medium text-gray-900 dark:text-white">Speaker:&nbsp;</span>{{ speaker.name }}</span>
              </span>
              <span class="flex md:justify-end w-auto">
                <svg xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 96 960 960" width="24">
                  <path
                    d="M250 566h460v-60H250v60Zm0 160h300v-60H250v60ZM141 896q-24 0-42-18.5T81 836V316q0-23 18-41.5t42-18.5h280l60 60h340q23 0 41.5 18.5T881 376v460q0 23-18.5 41.5T821 896H141Zm0-580v520h680V376H456l-60-60H141Zm0 0v520-520Z" />
                </svg>
                <span class="max-md:text-md max-sm:text-xs">
                  <span class="font-medium text-gray-900 dark:text-white">Topic:&nbsp;</span>{{ speaker.topic }}</span>
              </span>
            </div>
            <div class="speaker-description ml-1 max-md:text-md max-sm:text-xs font-normal mt-2 md:mt-1">
              <span class="font-medium text-gray-900 dark:text-white">Description:&nbsp;</span>{{ speaker.description }}
            </div>
            <div class="w-full h-1">&nbsp;</div>
            <span v-if="speaker.private"
              class="float-right inline-flex items-center text-xs font-normal text-gray-500 dark:text-gray-400">
              <svg aria-hidden="true" class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M3.707 2.293a1 1 0 00-1.414 1.414l14 14a1 1 0 001.414-1.414l-1.473-1.473A10.014 10.014 0 0019.542 10C18.268 5.943 14.478 3 10 3a9.958 9.958 0 00-4.512 1.074l-1.78-1.781zm4.261 4.26l1.514 1.515a2.003 2.003 0 012.45 2.45l1.514 1.514a4 4 0 00-5.478-5.478z"
                  clip-rule="evenodd"></path>
                <path
                  d="M12.454 16.697L9.75 13.992a4 4 0 01-3.742-3.741L2.335 6.578A9.98 9.98 0 00.458 10c1.274 4.057 5.065 7 9.542 7 .847 0 1.669-.105 2.454-.303z">
                </path>
              </svg>
              Private (EVS Only)
            </span>
            <span v-else
              class="float-right inline-flex items-center text-xs font-normal text-gray-500 dark:text-gray-400">
              <svg aria-hidden="true" class="w-3 h-3 mr-1" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M4.083 9h1.946c.089-1.546.383-2.97.837-4.118A6.004 6.004 0 004.083 9zM10 2a8 8 0 100 16 8 8 0 000-16zm0 2c-.076 0-.232.032-.465.262-.238.234-.497.623-.737 1.182-.389.907-.673 2.142-.766 3.556h3.936c-.093-1.414-.377-2.649-.766-3.556-.24-.56-.5-.948-.737-1.182C10.232 4.032 10.076 4 10 4zm3.971 5c-.089-1.546-.383-2.97-.837-4.118A6.004 6.004 0 0115.917 9h-1.946zm-2.003 2H8.032c.093 1.414.377 2.649.766 3.556.24.56.5.948.737 1.182.233.23.389.262.465.262.076 0 .232-.032.465-.262.238-.234.498-.623.737-1.182.389-.907.673-2.142.766-3.556zm1.166 4.118c.454-1.147.748-2.572.837-4.118h1.946a6.004 6.004 0 01-2.783 4.118zm-6.268 0C6.412 13.97 6.118 12.546 6.03 11H4.083a6.004 6.004 0 002.783 4.118z"
                  clip-rule="evenodd"></path>
              </svg>
              Public </span>&nbsp;
          </div>
        </div>
      </li>
    </ol>
  </div>
</template>

<style scoped>
.speaker-profile {
  min-width: 50px;
}

@media not all and (min-width: 640px) {
  .speaker-description {
    margin-left: calc(50px + 0.25rem);
  }
}
</style>
