export interface ApiSpeaker {
  id: number
  created_at: Date
  updated_at: Date
  deleted_at: Date | null
  name: string
  topic: string
  description: string
  schedule_id: number
}

export interface ApiSchedule {
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
