import { ref } from 'vue'
import { defineStore } from 'pinia'
import type { ApiSchedule, ApiSpeaker, SpeakerAdd, SpeakerUpdate } from '@/models/models'
import { getApiBaseUrl, headers } from '@/utils/api'

export const useSpeakerStore = defineStore('speakerStore', () => {
  const speakers = ref<ApiSpeaker[]>([])
  const speakerAddSchedule = ref<ApiSchedule>({} as ApiSchedule)
  const scheduleSpeakers = ref<ApiSpeaker[]>([])
  const speakerEdit = ref<ApiSpeaker | null>(null)

  function init() { }

  async function getSpeakers() {
    speakers.value = []
    await fetch(getApiBaseUrl() + '/api/v1/speakers')
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        const items = data.data
        items.forEach((speaker: ApiSpeaker) => {
          speakers.value.push(speaker)
        })
        speakers.value.sort((a: ApiSpeaker, b: ApiSpeaker) => {
          return a.order - b.order
        })
      })
  }

  async function getScheduleSpeakers(scheduleId: number) {
    scheduleSpeakers.value = []
    await fetch(getApiBaseUrl() + `/api/v1/schedule/${scheduleId}`)
      .then((response) => {
        return response.json()
      })
      .then((data) => {
        const items = data.data.speakers
        items.forEach((speaker: ApiSpeaker) => {
          scheduleSpeakers.value.push(speaker)
        })
      })
  }

  async function addSpeaker(speaker: SpeakerAdd) {
    const authHeaders = headers()
    const response = await fetch(getApiBaseUrl() + '/api/v1/speakers', {
      method: 'POST',
      headers: authHeaders,
      body: JSON.stringify(speaker)
    })
    const result = await response.json()
    return result
  }

  async function updateSpeaker(speaker: SpeakerUpdate) {
    const authHeaders = headers()
    const response = await fetch(getApiBaseUrl() + '/api/v1/speakers/' + speaker.id, {
      method: 'PUT',
      headers: authHeaders,
      body: JSON.stringify(speaker)
    })
    const result = await response.json()
    return result
  }

  async function deleteSpeaker(speakerId: number) {
    const authHeaders = headers()
    const response = await fetch(getApiBaseUrl() + '/api/v1/speakers/' + speakerId, {
      method: 'DELETE',
      headers: authHeaders
    })
    const result = await response.json()
    return result
  }

  function setSpeakerAddSchedule(schedule: ApiSchedule) {
    speakerAddSchedule.value = schedule
  }

  return {
    init,
    speakers,
    speakerEdit,
    getSpeakers,
    scheduleSpeakers,
    getScheduleSpeakers,
    addSpeaker,
    updateSpeaker,
    deleteSpeaker,
    speakerAddSchedule,
    setSpeakerAddSchedule
  }
})
