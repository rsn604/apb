package appt

type DeletedEntry struct {
	RepeatSdateYear  int `json:"repeatSdateYear"`
	RepeatSdateMonth int `json:"repeatSdateMonth"`
	RepeatSdateDay   int `json:"repeatSdateDay"`
	RepeatEdateYear  int `json:"repeatEdateYear"`
	RepeatEdateMonth int `json:"repeatEdateMonth"`
	RepeatEdateDay   int `json:"repeatEdateDay"`
}

type Repeat struct {
	Frequency        int            `json:"frequency"`
	DayNumber        int            `json:"dayNumber"`
	DayPosition      int            `json:"dayPosition"`
	DayOfWeek        int            `json:"dayOfWeek"`
	Month            int            `json:"month"`
	DateOfYear       string         `json:"dateOfYear"`
	RepeatSdateYear  int            `json:"repeatSdateYear"`
	RepeatSdateMonth int            `json:"repeatSdateMonth"`
	RepeatSdateDay   int            `json:"repeatSdateDay"`
	RepeatEdateYear  int            `json:"repeatEdateYear"`
	RepeatEdateMonth int            `json:"repeatEdateMonth"`
	RepeatEdateDay   int            `json:"repeatEdateDay"`
	DeletedEntry     []DeletedEntry `json:"deletedEntry"`
}

type Appointment struct {
	Id          int      `json:"id"`
	RepeatType  int      `json:"repeatType"`
	State       int      `json:"state"`
	Description string   `json:"description"`
	SdateYear   int      `json:"sdateYear"`
	SdateMonth  int      `json:"sdateMonth"`
	SdateDay    int      `json:"sdateDay"`
	DueDays     int      `json:"duedays"`
	TodoStatus  string   `json:"todostatus"`
	Stime       string   `json:"stime"`
	Etime       string   `json:"etime"`
	Location    string   `json:"location"`
	ConDays     int      `json:"conDays"`
	Priority    int      `json:"priority"`
	Note        []string `json:"note"`
	Repeat
}
