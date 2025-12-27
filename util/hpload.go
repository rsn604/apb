package main

import (
	"fmt"
	"strconv"
	"strings"
	//"time"

	"encoding/csv"
	//"encoding/json"
	"apb/appt"
	"os"
)

var databaseName = "BOLT"

func timeFormat(s string) string {
	if len(s) == 4 && s[1] == ':' {
		s = "0" + s
	}
	return s
}

func toInt(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}

func setYMD(y, m, d string) (int, int, int) {
	yy, _ := strconv.Atoi(strings.TrimSpace(y))
	mm, _ := strconv.Atoi(strings.TrimSpace(m))
	dd, _ := strconv.Atoi(strings.TrimSpace(d))
	return yy + 1900, mm + 1, dd + 1
}

func main() {

	var header = []string{"Sdate", "Description", "Location", "State", "Stime", "Etime", "ConDays", "alarm", "Note", "RepeatType", "Frequency", "days", "Month", "y1", "m1", "d1", "y2", "m2", "d2", "deleted"}

	if len(os.Args) < 3 {
		fmt.Println("Usage: hpload <load file> <DB name>")
		return
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	reader := csv.NewReader(fp)
	reader.Comma = ','
	reader.LazyQuotes = true
	reader.FieldsPerRecord = -1 // Nocheck fields count
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	manager := appt.GetManager(databaseName)
	err = manager.Connect(databaseName, os.Args[2])
	if err != nil {
		panic(err)
	}
	manager.Define()

	for _, record := range records {
		deleted := []string{}
		ap := appt.Appointment{}
		var item = make(map[string]string)
		for j, value := range record {
			if j >= len(header) {
				deleted = append(deleted, value)

			} else {
				item[header[j]] = value
			}
		}
		i := 0

		for {
			if i >= len(deleted) {
				break
			}
			deletedEntry := appt.DeletedEntry{}
			deletedEntry.RepeatSdateYear, deletedEntry.RepeatSdateMonth, deletedEntry.RepeatSdateDay = setYMD(deleted[i], deleted[i+1], deleted[i+2])
			deletedEntry.RepeatEdateYear, deletedEntry.RepeatEdateMonth, deletedEntry.RepeatEdateDay = setYMD(deleted[i], deleted[i+1], deleted[i+2])

			ap.DeletedEntry = append(ap.DeletedEntry, deletedEntry)
			i = i + 4
		}

		appt.UpdateSdate(&ap, item["Sdate"])
		//ap.Sdate = item["Sdate"]
		ap.Description = item["Description"]
		ap.Location = item["Location"]

		ap.ConDays = toInt(item["ConDays"])
		ap.Note = strings.Split(item["Note"], "\\r\\n")

		state := toInt(item["State"])
		if state == 20 {
			ap.State = appt.STATE_TODO
			ap.Priority = toInt(item["Stime"])
			dueDays := toInt(item["Etime"])
			if dueDays > 0 {
				ap.DueDays = dueDays - 1
			} else {
				ap.DueDays = 0
			}
			/*

			   TodoStatus
			*/

		}
		if state == 135 {
			ap.State = appt.STATE_APPT
			ap.Stime = timeFormat(item["Stime"])
			ap.Etime = timeFormat(item["Etime"])
		}

		// Repeat
		ap.RepeatType = toInt(item["RepeatType"])
		if ap.RepeatType != 0x00 {
			ap.Frequency = toInt(item["Frequency"])
			ap.Month = toInt(item["Month"])

			ap.RepeatSdateYear, ap.RepeatSdateMonth, ap.RepeatSdateDay = setYMD(item["y1"], item["m1"], item["d1"])
			ap.RepeatEdateYear, ap.RepeatEdateMonth, ap.RepeatEdateDay = setYMD(item["y2"], item["m2"], item["d2"])

			days := toInt(item["days"])
			if (days & 0x80) == 0 {
				ap.DayNumber = days
				if days > 0 {
					month := 0
					for i := 0; i < 12; i++ {
						if ((ap.Month >> i) & 0x01) != 0 {
							month = i
							break
						}
					}
					ap.DateOfYear = fmt.Sprintf("%02d/%02d", month+1, days)

				}
			} else {
				ap.DayPosition = days >> 8
				ap.DayOfWeek = (days & 0x007f) << 1
				if ap.DayOfWeek == 0x0080 {
					ap.DayOfWeek = 0x0001
				}
			}
		}

		manager.AddAppt(ap)
	}
	manager.Close()
}
