interface ApiSpeaker {
  id: number
  created_at: Date
  updated_at: Date
  deleted_at: Date | null
  name: string
  topic: string
  description: string
  order: number
  minutes: number
  schedule_id: number
  private: boolean
}

interface SpeakerAdd {
  name: string
  topic: string
  description: string
  order: number
  minutes: number
  schedule_id: number
  private: boolean
}

interface SpeakerUpdate {
  id: number
  name: string
  topic: string
  description: string
  order: number
  minutes: number
  schedule_id: number
  private: boolean
}

interface ApiSchedule {
  id: number
  created_at: Date
  updated_at: Date
  deleted_at: Date | null
  title: string
  description: string
  start_date_time: Date
  end_date_time: Date
  speakers: ApiSpeaker[]
}

interface ScheduleAdd {
  title: string
  description: string
  start_date_time: Date
  end_date_time: Date
  speakers: SpeakerAdd[] | null
}

interface ScheduleUpdate {
  id: number
  title: string
  description: string
  start_date_time: Date
  end_date_time: Date
  speakers: SpeakerAdd[] | null
}

interface ScheduleSearchParams {
  f?: string // start_date_time From
  t?: string // start_date_time To
}

export type {
  ApiSpeaker,
  SpeakerAdd,
  SpeakerUpdate,
  ApiSchedule,
  ScheduleAdd,
  ScheduleUpdate,
  ScheduleSearchParams
}
