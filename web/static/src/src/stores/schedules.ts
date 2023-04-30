import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { ApiSchedule, ScheduleAdd, ScheduleUpdate } from '@/models/models'
import { getApiBaseUrl, query } from '@/utils/api'

export const useScheduleStore = defineStore('scheduleStore', () => {
  const schedulesHome = ref<ApiSchedule[]>([])
  const schedulesAdmin = ref<ApiSchedule[]>([])
  const scheduleEdit = ref<ApiSchedule | null>(null)

  function init() {
    getSchedulesHome()
  }

  async function getSchedules(params: Object = {}) {
    const response = await fetch(query('/api/v1/schedules', params))
    const data = await response.json()
    const items = data?.data
    items.sort((a: ApiSchedule, b: ApiSchedule) => {
      return new Date(a.start_date_time).getTime() - new Date(b.start_date_time).getTime()
    })
    return items
  }

  async function getSchedulesHome(params: Object = {}) {
    const items: ApiSchedule[] = await getSchedules(params)
    schedulesHome.value = items
  }

  async function getSchedulesAdmin(params: Object = {}) {
    const items: ApiSchedule[] = await getSchedules(params)
    schedulesAdmin.value = items
  }

  async function addSchedule(schedule: ScheduleAdd) {
    const response = await fetch(getApiBaseUrl() + '/api/v1/schedules', {
      method: 'POST',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(schedule)
    })
    const result = await response.json()
    return result
  }

  async function updateSchedule(schedule: ScheduleUpdate) {
    const response = await fetch(getApiBaseUrl() + '/api/v1/schedules/' + schedule.id, {
      method: 'PUT',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(schedule)
    })
    const result = await response.json()
    return result
  }

  async function deleteSchedule(scheduleId: number) {
    const response = await fetch(getApiBaseUrl() + '/api/v1/schedules/' + scheduleId, {
      method: 'DELETE',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json'
      }
    })
    const result = await response.json()
    return result
  }

  return {
    init,
    schedulesHome,
    schedulesAdmin,
    scheduleEdit,
    getSchedules,
    getSchedulesAdmin,
    getSchedulesHome,
    addSchedule,
    updateSchedule,
    deleteSchedule
  }
})
