package types

type SchedulePostRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	StartDateTime string `json:"start_date_time"`
	EndDateTime   string `json:"end_date_time"`
}

type SchedulePutRequest struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	StartDateTime string `json:"start_date_time"`
	EndDateTime   string `json:"end_date_time"`
}

type SpeakerPostRequest struct {
	Name        string `json:"name"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Audience    string `json:"audience"`
	Order       int    `json:"order"`
	Minutes     int    `json:"minutes"`
	ScheduleId  int    `json:"schedule_id"`
	Private     bool   `json:"private"`
}

type SpeakerPutRequest struct {
	Name        string `json:"name"`
	Topic       string `json:"topic"`
	Description string `json:"description"`
	Audience    string `json:"audience"`
	Order       int    `json:"order"`
	Minutes     int    `json:"minutes"`
	ScheduleId  int    `json:"schedule_id"`
	Private     bool   `json:"private"`
}
