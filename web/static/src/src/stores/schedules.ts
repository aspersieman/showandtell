import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { ApiSchedule } from '@/models/models'
import { getApiBaseUrl } from '@/utils/api'

export const useScheduleStore = defineStore('scheduleStore', () => {
  const schedules = ref<ApiSchedule[]>([])
  const scheduleList = ref<ApiSchedule[]>([])
  function init() {
    getSchedules()
  }

  async function getSchedules() {
    schedules.value = []
    await fetch(getApiBaseUrl() + '/api/v1/schedules')
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        const items = data.data
        items.forEach((schedule: ApiSchedule) => {
          schedules.value.push(schedule)
        })
      })
  }

  async function getScheduleList() {
    scheduleList.value = []
    await fetch(getApiBaseUrl() + '/api/v1/schedules?b=0&e=100')
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        const items = data.data
        items.forEach((schedule: ApiSchedule) => {
          scheduleList.value.push(schedule)
        })
      })
  }

  return { init, schedules, getSchedules, scheduleList, getScheduleList }
})
