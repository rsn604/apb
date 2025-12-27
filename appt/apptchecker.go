package appt

import (
	//"fmt"
	"strings"
	"time"
	//"strconv"
	//"apb/appt"
	"regexp"
	"sort"
	//"log"
)

const (
	YYYYMM_FORMAT = "2006/01"
	MMDD_FORMAT   = "01/02"
	STATE_APPT    = 0x00
	STATE_TODO    = 0x02
	C_DAYS        = 7
)

var DATE_FORMAT = "2006/01/02"

var RepeatTBL = map[string]int{
	"NONE":    0x00,
	"DAILY":   0x02,
	"WEEKLY":  0x04,
	"MONTHLY": 0x08,
	"YEARLY":  0x10,
	"CUSTOM":  0x20,
}

var WeekTBL = map[string]int{
	"SUNDAY":    0x0001,
	"MONDAY":    0x0002,
	"TUESDAY":   0x0004,
	"WEDNESDAY": 0x0008,
	"THURSDAY":  0x0010,
	"FRIDAY":    0x0020,
	"SATURDAY":  0x0040,
}

var PositionTBL = map[string]int{
	"1ST":  0x0001,
	"2ND":  0x0002,
	"3RD":  0x0004,
	"4TH":  0x0008,
	"LAST": 0x0010,
}

var MonthTBL = map[string]int{
	"JANUARY":   0x0001,
	"FEBRUARY":  0x0002,
	"MARCH":     0x0004,
	"APRIL":     0x0008,
	"MAY":       0x0010,
	"JUNE":      0x0020,
	"JULY":      0x0040,
	"AUGUST":    0x0080,
	"SEPTEMBER": 0x0100,
	"OCTOBER":   0x0200,
	"NOVEMBER":  0x0400,
	"DECEMBER":  0x0800,
}

//

var ptnMD = regexp.MustCompile(`^(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])$`)
var ptnHHMM = regexp.MustCompile(`^(0[0-9]|1[0-9]|2[0-3]):([0-5][0-9])$`)

func CheckMD(s string) bool {
	return ptnMD.MatchString(s)
}

func CheckHHMM(s string) bool {
	return ptnHHMM.MatchString(s)
}

/*
var ptnYMD = regexp.MustCompile(`^[0-9]{4}/(0[1-9]|1[0-2])/(0[1-9]|[12][0-9]|3[01])$`)

	func CheckYMD(s string) bool {
		return ptnYMD.MatchString(s)
	}
*/
func CheckYMD(s string) bool {
	goLayout := formatLayout(DATE_FORMAT)
	_, err := time.Parse(goLayout, s)
	return err == nil
}

func formatLayout(format string) string {
	layout := format
	layout = replaceAll(layout, "YYYY", "2006")
	layout = replaceAll(layout, "YY", "06")
	layout = replaceAll(layout, "MM", "01")
	layout = replaceAll(layout, "DD", "02")
	layout = replaceAll(layout, "HH", "15")
	layout = replaceAll(layout, "mm", "04")
	layout = replaceAll(layout, "ss", "05")
	return layout
}

func replaceAll(s, old, new string) string {
	return strings.Replace(s, old, new, -1)
}

// ----------------------------------------------------
func GetDueDate(ap Appointment) string {
	t, _ := time.Parse(DATE_FORMAT, GetSdate(ap))
	t = t.AddDate(0, 0, ap.DueDays)
	return t.Format(DATE_FORMAT)
}

func UpdateDueDate(ap *Appointment, s string) {
	ts, _ := time.Parse(DATE_FORMAT, GetSdate(*ap))
	t, _ := time.Parse(DATE_FORMAT, s)
	ap.DueDays = GetDayDiff(t, ts)
}

func GetSdate(ap Appointment) string {
	t := time.Date(ap.SdateYear, time.Month(ap.SdateMonth), ap.SdateDay, 0, 0, 0, 0, time.Local)
	return t.Format(DATE_FORMAT)
}

func UpdateSdate(ap *Appointment, s string) {
	t, _ := time.Parse(DATE_FORMAT, s)
	ap.SdateYear = t.Year()
	ap.SdateMonth = int(t.Month())
	ap.SdateDay = t.Day()
}

func GetRepeatSdate(ap Appointment) string {
	//@@@@
	if ap.RepeatSdateYear == 0 || ap.RepeatSdateYear == 1 {
		return ""
	}
	t := time.Date(ap.RepeatSdateYear, time.Month(ap.RepeatSdateMonth), ap.RepeatSdateDay, 0, 0, 0, 0, time.Local)
	return t.Format(DATE_FORMAT)
}

func UpdateRepeatSdate(ap *Appointment, s string) {
	t, _ := time.Parse(DATE_FORMAT, s)
	ap.RepeatSdateYear = t.Year()
	ap.RepeatSdateMonth = int(t.Month())
	ap.RepeatSdateDay = t.Day()
}

func GetRepeatEdate(ap Appointment) string {
	t := time.Date(ap.RepeatEdateYear, time.Month(ap.RepeatEdateMonth), ap.RepeatEdateDay, 0, 0, 0, 0, time.Local)
	return t.Format(DATE_FORMAT)
}

func UpdateRepeatEdate(ap *Appointment, s string) {
	t, _ := time.Parse(DATE_FORMAT, s)
	ap.RepeatEdateYear = t.Year()
	ap.RepeatEdateMonth = int(t.Month())
	ap.RepeatEdateDay = t.Day()
}

func GetDeletedRepeatSdate(ap DeletedEntry) string {
	t := time.Date(ap.RepeatSdateYear, time.Month(ap.RepeatSdateMonth), ap.RepeatSdateDay, 0, 0, 0, 0, time.Local)
	return t.Format(DATE_FORMAT)
}

func UpdateDeletedRepeatSdate(ap *DeletedEntry, s string) {
	t, _ := time.Parse(DATE_FORMAT, s)
	ap.RepeatSdateYear = t.Year()
	ap.RepeatSdateMonth = int(t.Month())
	ap.RepeatSdateDay = t.Day()
}

func GetDeletedRepeatEdate(ap DeletedEntry) string {
	t := time.Date(ap.RepeatEdateYear, time.Month(ap.RepeatEdateMonth), ap.RepeatEdateDay, 0, 0, 0, 0, time.Local)
	return t.Format(DATE_FORMAT)
}

func UpdateDeletedRepeatEdate(ap *DeletedEntry, s string) {
	t, _ := time.Parse(DATE_FORMAT, s)
	ap.RepeatEdateYear = t.Year()
	ap.RepeatEdateMonth = int(t.Month())
	ap.RepeatEdateDay = t.Day()
}

// ----------------------------------------------------
type Pair struct {
	Key   string
	Value int
}

func SortMap(m map[string]int) []string {
	var pairs []Pair
	for k, v := range m {
		pairs = append(pairs, Pair{Key: k, Value: v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value < pairs[j].Value
	})

	var keyList []string
	for i := 0; i < len(pairs); i++ {
		keyList = append(keyList, pairs[i].Key)
	}
	return keyList
}

// ----------------------------------------------------
func GetRepeatType(s string) int {
	return RepeatTBL[strings.ToUpper(s)]

}

func GetRepeatString(w int) string {
	for k, v := range RepeatTBL {
		if v == w {
			return k
		}
	}
	return ""
}

func GetWeek(s string) int {
	return WeekTBL[strings.ToUpper(s)]

}

func GetWeekString(w int) string {
	for k, v := range WeekTBL {
		if v == w {
			return k
		}
	}
	return ""
}

func GetPosition(s string) int {
	return PositionTBL[strings.ToUpper(s)]
}

func GetPositionString(w int) string {
	for k, v := range PositionTBL {
		if v == w {
			return k
		}
	}
	return ""
}

func GetMonth(s string) int {
	return MonthTBL[strings.ToUpper(s)]
}

func GetMonthString(w int) string {
	for k, v := range MonthTBL {
		if v == w {
			return k
		}
	}
	return ""
}

func GetWeekNumber(t time.Time) int {
	firstDayOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	d := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)

	daysSinceFirst := d.Sub(firstDayOfMonth).Hours() / 24
	return int(daysSinceFirst/7) + 1
}

func GetDayDiff(t time.Time, ts time.Time) int {
	d1 := time.Date(ts.Year(), ts.Month(), ts.Day(), 0, 0, 0, 0, time.Local)
	d2 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
	return int(d2.Sub(d1).Hours() / 24)
}

func GetFirstDayOfYear(ap Appointment) time.Time {
	ts, _ := time.Parse(DATE_FORMAT, GetRepeatSdate(ap))
	return ts
}

func GetFirstDayOfMonth(ap Appointment) time.Time {
	ts, _ := time.Parse(DATE_FORMAT, GetRepeatSdate(ap))
	if ap.DayNumber != 0 {
		return time.Date(ts.Year(), ts.Month(), ap.DayNumber, 0, 0, 0, 0, time.UTC)

	}

	for {
		if checkWeekNumber(ap, ts) && checkDayOfWeek(ap, ts) {
			break
		}
		ts = ts.AddDate(0, 0, 1)
	}
	return ts
}

func getFirstDayOfWeek(ap Appointment) time.Time {
	ts, _ := time.Parse(DATE_FORMAT, GetRepeatSdate(ap))
	for {
		if checkDayOfWeek(ap, ts) {
			break
		}
		ts = ts.AddDate(0, 0, 1)
	}
	return ts
}

// ----------------------------------------------------
func checkDayNumber(ap Appointment, t time.Time) bool {
	if ap.DayNumber == int(t.Day()) {
		return true
	} else {
		return false
	}
}

func checkWeekNumber(ap Appointment, t time.Time) bool {
	weekNumber := GetWeekNumber(t)

	if ((ap.DayPosition >> (weekNumber - 1)) & 0x01) != 0 {
		return true
	}

	tn := t.AddDate(0, 0, 7)
	if GetWeekNumber(tn) == 1 {
		if ((ap.DayPosition >> (weekNumber)) & 0x01) != 0 {
			return true
		}
	}
	return false
}

func checkDayOfWeek(ap Appointment, t time.Time) bool {
	if ((ap.DayOfWeek >> int(t.Weekday())) & 0x01) != 0 {
		return true
	}
	return false
}

func checkMonth(ap Appointment, t time.Time) bool {
	if ((ap.Month >> (int(t.Month()) - 1)) & 0x01) != 0 {
		return true
	}
	return false
}

// ----------------------------------------------------
func checkDuration(ap Appointment, t time.Time) bool {
	//@@@
	if GetRepeatSdate(ap) == "" && GetRepeatEdate(ap) == "" {
		//return true
		return false
	}

	if !CheckYMD(GetRepeatSdate(ap)) || !CheckYMD(GetRepeatEdate(ap)) {
		return false
	}

	ts := t.Format(DATE_FORMAT)
	if ts >= GetRepeatSdate(ap) && ts <= GetRepeatEdate(ap) {
		return true
	} else {
		return false
	}
}

// ----------------------------------------------------
func checkFrequency(ap Appointment, t time.Time, ts time.Time) bool {
	diff := 0
	frequency := ap.Frequency
	if frequency == 0 {
		return false
	}

	if ap.RepeatType == GetRepeatType("DAILY") {
		diff = GetDayDiff(t, ts)
	}

	if ap.RepeatType == GetRepeatType("WEEKLY") {
		diff = GetDayDiff(t, ts)
		if (diff % (frequency * C_DAYS)) == 0 {
			return true
		} else {
			return false
		}
		/*
			_, week_t := t.ISOWeek()
			_, week_ts := ts.ISOWeek()
			diff = week_t - week_ts
		*/
	}

	if ap.RepeatType == GetRepeatType("MONTHLY") {
		year1, month1, _ := t.Date()
		year2, month2, _ := ts.Date()
		diff = (year2-year1)*12 + int(month2) - int(month1)
	}

	if ap.RepeatType == GetRepeatType("YEARLY") {
		diff = t.Year() - ts.Year()
		if t.YearDay() < ts.YearDay() {
			diff--
		}
	}

	if (diff % frequency) == 0 {
		//log.Printf("log 00 ts:%s t:%s ap.RepeatType:%s diff:%d frequency:%d int(t.Sub(ts):%d\n", ts.Format(DATE_FORMAT), t.Format(DATE_FORMAT), ap.Description, diff, frequency, int(t.Sub(ts).Hours()))
		return true
	} else {
		return false
	}

	return true
}

// ----------------------------------------------
func checkYearPosition(ap Appointment, t time.Time) bool {
	if checkWeekNumber(ap, t) && checkDayOfWeek(ap, t) && checkMonth(ap, t) {
		return true
	}
	return false
}

func checkDayPosition(ap Appointment, t time.Time) bool {
	if checkWeekNumber(ap, t) && checkDayOfWeek(ap, t) {
		return true
	}
	return false
}

// ----------------------------------------------
func CheckMMDD(s string) (string, bool) {
	mmdd := strings.Split(s, "/")
	if len(mmdd) != 2 {
		return "", false
	}
	if len(mmdd[0]) > 2 || len(mmdd[1]) > 2 {
		return "", false
	}

	if len(mmdd[0]) == 1 {
		mmdd[0] = "0" + mmdd[0]
	}
	if len(mmdd[1]) == 1 {
		mmdd[1] = "0" + mmdd[1]
	}

	if !CheckMD(mmdd[0] + "/" + mmdd[1]) {
		return "", false
	}

	return mmdd[0] + "/" + mmdd[1], true
}

func checkDayOfYear(ap Appointment, t time.Time) bool {
	mmdd, flag := CheckMMDD(ap.DateOfYear)
	if flag == false {
		return false
	}
	if t.Format(MMDD_FORMAT) == mmdd {
		return true
	}
	return false
}

// ----------------------------------------------------
// Customized
// ----------------------------------------------------
func checkCustomized(ap Appointment, t time.Time) bool {
	if !checkDuration(ap, t) {
		return false
	}

	if ap.DayNumber != 0 {
		if !checkDayNumber(ap, t) {
			return false
		}
	} else if !checkDayPosition(ap, t) {
		return false
	}

	if !checkMonth(ap, t) {
		return false
	}
	return true
}

// ----------------------------------------------------
// Yearly
// ----------------------------------------------------
func checkYearly(ap Appointment, t time.Time) bool {
	if !checkDuration(ap, t) {
		return false
	}

	if len(ap.DateOfYear) > 0 {
		if !checkDayOfYear(ap, t) {
			return false
		}
	} else {
		if !checkYearPosition(ap, t) {
			return false
		}
	}

	ts := GetFirstDayOfYear(ap)
	if !checkFrequency(ap, t, ts) {
		return false
	}

	return true
}

// ----------------------------------------------------
// Monthly
// ----------------------------------------------------
func checkMonthly(ap Appointment, t time.Time) bool {
	if !checkDuration(ap, t) {
		return false
	}
	if ap.DayNumber != 0 {
		if !checkDayNumber(ap, t) {
			return false
		}
	} else {
		if !checkDayPosition(ap, t) {
			return false
		}
	}

	ts := GetFirstDayOfMonth(ap)
	if !checkFrequency(ap, t, ts) {
		return false
	}

	return true
}

// ----------------------------------------------------
// Weekly
// ----------------------------------------------------
func checkWeekly(ap Appointment, t time.Time) bool {
	if !checkDuration(ap, t) {
		return false
	}

	if !checkDayOfWeek(ap, t) {
		return false
	}

	ts := getFirstDayOfWeek(ap)
	if !checkFrequency(ap, t, ts) {
		return false
	}

	return true
}

// ----------------------------------------------------
// Daily
// ----------------------------------------------------
func checkDaily(ap Appointment, t time.Time, databaseName, connectString string) bool {
	if !checkDuration(ap, t) {
		return false
	}

	manager := GetManager(databaseName)
	manager.Connect(databaseName, connectString)
	u := manager.GetApptByID(ap.Id)
	manager.Close()

	ts, _ := time.Parse(DATE_FORMAT, GetSdate(u))
	if t.Format(DATE_FORMAT) < GetSdate(u) {
		return false
	}

	if !checkFrequency(ap, t, ts) {
		return false
	}

	return true
}

// ----------------------------------------------------
func isDeletedEntry(ap Appointment, t time.Time) bool {
	for _, repeat := range ap.DeletedEntry {
		/*
			if t.Format(DATE_FORMAT) >= repeat.RepeatSdate  && t.Format(DATE_FORMAT) <= repeat.RepeatEdate {
				return true
			}
		*/
		//@@@@@@
		if t.Format(DATE_FORMAT) >= GetDeletedRepeatSdate(repeat) && t.Format(DATE_FORMAT) <= GetDeletedRepeatEdate(repeat) {
			return true
		}
	}

	return false
}

// ----------------------------------------------------
func GetApptToday(databaseName, connectString string, t time.Time) []Appointment {
	return getObjectToday(databaseName, connectString, t, STATE_APPT)
}

func GetTodoToday(databaseName, connectString string, t time.Time) []Appointment {
	return getObjectToday(databaseName, connectString, t, STATE_TODO)
}

func getObjectToday(databaseName, connectString string, t time.Time, state int) []Appointment {
	var s []Appointment
	manager := GetManager(databaseName)
	manager.Connect(databaseName, connectString)
	appts := manager.GetAppts()
	manager.Close()
	/*
		log.Printf("Time:%s\n", t.Format(DATE_FORMAT))
		for _, appt := range appts {
			log.Printf("Description:%s Sdate:%s\n", Description, Sdate)
		}
		log.Printf("------------\n")
	*/

	for _, ap := range appts {
		if ap.State != state {
			continue
		}

		if isDeletedEntry(ap, t) {
			continue
		}

		if ap.State == STATE_TODO {
			today := time.Now().Format(DATE_FORMAT)

			//@@@@@@
			if GetSdate(ap) <= today && ap.RepeatType == GetRepeatType("NONE") {
				s = append(s, ap)
				continue
			}
			if t.Format(DATE_FORMAT) <= today {
				s = append(s, ap)
				continue
			}

		}

		if ap.RepeatType == GetRepeatType("NONE") {
			if ap.State == STATE_APPT {
				if GetSdate(ap) == t.Format(DATE_FORMAT) {
					s = append(s, ap)
				} else if ap.ConDays > 0 {
					ts, _ := time.Parse(DATE_FORMAT, GetSdate(ap))
					diff := GetDayDiff(t, ts)
					if diff > 0 && ap.ConDays > diff {
						s = append(s, ap)
					}
				}
				continue
			}
			/*
				if ap.State == STATE_TODO{
					s = append(s, ap)
					continue
				}
			*/
		}

		if ap.RepeatType == GetRepeatType("DAILY") {
			if checkDaily(ap, t, databaseName, connectString) {
				s = append(s, ap)
			}
			continue
		}

		if ap.RepeatType == GetRepeatType("WEEKLY") {
			if checkWeekly(ap, t) {
				s = append(s, ap)
			} else {
				ts := t
				for i := 0; i < ap.ConDays-1; i++ {
					ts = ts.AddDate(0, 0, -1)
					if checkWeekly(ap, ts) {
						s = append(s, ap)
						break
					}
				}
			}
			continue
		}

		if ap.RepeatType == GetRepeatType("MONTHLY") {
			if checkMonthly(ap, t) {
				s = append(s, ap)
			} else {
				ts := t
				for i := 0; i < ap.ConDays-1; i++ {
					ts = ts.AddDate(0, 0, -1)
					if checkMonthly(ap, ts) {
						s = append(s, ap)
						break
					}
				}
			}
			continue
		}

		if ap.RepeatType == GetRepeatType("YEARLY") {
			if checkYearly(ap, t) {
				s = append(s, ap)
			} else {
				ts := t
				for i := 0; i < ap.ConDays-1; i++ {
					ts = ts.AddDate(0, 0, -1)
					if checkYearly(ap, ts) {
						s = append(s, ap)
						break
					}
				}
			}
			continue
		}

		if ap.RepeatType == GetRepeatType("CUSTOM") {
			if checkCustomized(ap, t) {
				s = append(s, ap)
			} else {
				ts := t
				for i := 0; i < ap.ConDays-1; i++ {
					ts = ts.AddDate(0, 0, -1)
					if checkCustomized(ap, ts) {
						s = append(s, ap)
						break
					}
				}
			}
			continue
		}
	}
	return s
}

// ----------------------------------------------------
func GetNextTodos(databaseName, connectString string, currentTime time.Time) map[string]Appointment {
	return getTargetObjects(databaseName, connectString, currentTime, STATE_TODO, true)
}

func GetPreviousTodos(databaseName, connectString string, currentTime time.Time) map[string]Appointment {
	return getTargetObjects(databaseName, connectString, currentTime, STATE_TODO, false)
}

func GetNextAppts(databaseName, connectString string, currentTime time.Time) map[string]Appointment {
	return getTargetObjects(databaseName, connectString, currentTime, STATE_APPT, true)
}

func GetPreviousAppts(databaseName, connectString string, currentTime time.Time) map[string]Appointment {
	return getTargetObjects(databaseName, connectString, currentTime, STATE_APPT, false)
}

func getTargetObjects(databaseName, connectString string, currentTime time.Time, state int, next bool) map[string]Appointment {
	t := currentTime
	s := map[string]Appointment{}

	manager := GetManager(databaseName)
	manager.Connect(databaseName, connectString)
	appts := manager.GetAppts()
	manager.Close()

	for _, ap := range appts {
		if ap.State != state {
			continue
		}
		/*
			if isDeletedEntry(ap, t) {
				continue
			}
		*/
		for {

			if isDeletedEntry(ap, t) {
				goto NEXTENTRY
			}
			if ap.RepeatType == GetRepeatType("NONE") {
				//@@@@@@
				if GetSdate(ap) > t.Format(DATE_FORMAT) {
					goto NEXTENTRY
				}

				if GetSdate(ap) == t.Format(DATE_FORMAT) {
					s[t.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
					goto NEXTAPPT
				}

				if ap.ConDays > 0 {
					ts := t
					for i := 0; i < ap.ConDays-1; i++ {
						ts = ts.AddDate(0, 0, -1)
						if GetSdate(ap) == ts.Format(DATE_FORMAT) {
							s[ts.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
							break
						}
					}
				}
				goto NEXTAPPT
			}

			if !CheckYMD(GetRepeatSdate(ap)) || !CheckYMD(GetRepeatEdate(ap)) {
				goto NEXTAPPT
			}

			if ap.RepeatType == GetRepeatType("DAILY") {
				if checkDaily(ap, t, databaseName, connectString) {
					s[t.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
					goto NEXTAPPT
				}
			}

			if ap.RepeatType == GetRepeatType("WEEKLY") {
				if checkWeekly(ap, t) {
					s[t.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
					goto NEXTAPPT

				} else {
					ts := t
					for i := 0; i < ap.ConDays-1; i++ {
						ts = ts.AddDate(0, 0, -1)
						if checkWeekly(ap, ts) {
							s[ts.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
							goto NEXTAPPT
						}
					}
				}
			}

			if ap.RepeatType == GetRepeatType("MONTHLY") {
				if checkMonthly(ap, t) {
					s[t.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
					goto NEXTAPPT
				} else {
					ts := t
					for i := 0; i < ap.ConDays-1; i++ {
						ts = ts.AddDate(0, 0, -1)
						if checkMonthly(ap, ts) {
							s[ts.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
							goto NEXTAPPT
						}
					}
				}
			}

			if ap.RepeatType == GetRepeatType("YEARLY") {
				if checkYearly(ap, t) {
					s[t.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
					goto NEXTAPPT
				} else {
					ts := t
					for i := 0; i < ap.ConDays-1; i++ {
						ts = ts.AddDate(0, 0, -1)
						if checkYearly(ap, ts) {
							s[ts.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
							goto NEXTAPPT
						}
					}
				}
			}

			if ap.RepeatType == GetRepeatType("CUSTOM") {
				if checkCustomized(ap, t) {
					s[t.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
					goto NEXTAPPT
				} else {
					ts := t
					for i := 0; i < ap.ConDays-1; i++ {
						ts = ts.AddDate(0, 0, -1)
						if checkCustomized(ap, ts) {
							s[ts.Format(DATE_FORMAT)+" "+ap.Stime+" "+ap.Description] = ap
							goto NEXTAPPT
						}
					}
				}
			}
		NEXTENTRY:
			if next {
				t = t.AddDate(0, 0, 1)
				if ap.RepeatType != GetRepeatType("NONE") && GetRepeatEdate(ap) < t.Format(DATE_FORMAT) {
					break
				}
			} else {
				t = t.AddDate(0, 0, -1)
				//if ap.RepeatType != GetRepeatType("NONE") && GetRepeatSdate(ap) > t.Format(DATE_FORMAT){
				if GetSdate(ap) > t.Format(DATE_FORMAT) {
					break
				}
			}
		}
	NEXTAPPT:
		t = currentTime
	}
	return s
}
