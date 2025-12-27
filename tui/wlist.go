package tui

import (
	"apb/appt"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strconv"
	"strings"
)

func WlistPanel() *taps.Panel {
	var help = `Weekly List

Display Weekly Appts.

(1) View and Change Appt
Change Week(next)  : CTRL-N (Right Arrow),
            (prev) : CTRL-P (Left Arrow)
Return to Today    : CTRL-T
Add New Appt       : F2 -> Go to "Detail"
                   : Select field
                     -> Go to "Detail"
Delete Appt        : F3
Find Appt          : "<F>" F4
View Calendar      : "<G>" F5
View 6-Month       : "<6>" F6
Month              : "<M>" F7
                     -> Go to "Monthly List"

"ESC" "<Q>" F12 : Quit.

`
	var styleMatrix = [][]string{
		//	{"WEEKDAY", "yellow, underline", "default"},
		{"WEEKDAY", "yellow", "default"},
		{"DAYTIME", "white", "default"},
		{"PFKEY", "white", "default"},
		{"label", "aqua", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"list", "white", "default"},
		{"list_focus", "black", "aqua"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 9999
EndY = 9999
ExitKey = ["F1", "F4", "F5", "F6", "F7", "F10", "F12", "Ctrl-T", "Ctrl-P", "Ctrl-N"]

# -------------------------------------------------
[[Field]]	
Name = "L_YM"
X = 0
Y = 0
FieldLen = 8
Style = "WEEKDAY"
FieldType = "label"

[[Field]]	
Name = "L_WEEKDAY"
X = 8
Y = 0
FieldLen = 9
Cols = 7
Style = "WEEKDAY"
FieldType = "label"

[[Field]]	
Name = "L_DAYTIME"
X = 0
Y = 1
FieldLen = 7
Rows = 9997
Style = "DAYTIME"
FieldType = "label"

[[Field]]	
Name = "L_WLIST01"
X = 7
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST01"
X = 8
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

[[Field]]	
Name = "L_WLIST02"
X = 16
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST02"
X = 17
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

[[Field]]	
Name = "L_WLIST03"
X = 25
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST03"
X = 26
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

[[Field]]	
Name = "L_WLIST04"
X = 34
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST04"
X = 35
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

[[Field]]	
Name = "L_WLIST05"
X = 43
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST05"
X = 44
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

[[Field]]	
Name = "L_WLIST06"
X = 52
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST06"
X = 53
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

[[Field]]	
Name = "L_WLIST07"
X = 61
Y = 0
FieldLen=0
Rows = 9997
Rect = true
Style = "list"
FieldType = "label"

[[Field]]
Name = "S_WLIST07"
X = 62
Y = 1
FieldLen = 8
Rows = 9997
Style = "list, list_focus"
FieldType = "select"

# -------------------------------------------------
[[Field]]
Name = "H"
Data = "<H>"
X = 1
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "A"
Data = "Add"
X = 5
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "label"

[[Field]]	
Name = "D"
Data = "Del"
X = 9
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "label"

[[Field]]	
Name = "F"
Data = "<F>"
X = 13
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "G"
Data = "<G>"
X = 17
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "6"
Data = "<6>"
X = 21
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "M"
Data = "<M>"
X = 25
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

#[[Field]]	
#Name = "T"
#Data = "<T>"
#X = 33
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

[[Field]]	
Name = "Q"
Data = "<Q>"
X = 37
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "L01"
Data = " F1  F2  F3  F4  F5  F6  F7  F8  F10 F12"
X = 0
Y = 9999
FieldLen = 41
Style = "PFKEY"
FieldType = "label"
`
	return taps.NewPanel(doc, styleMatrix, help)
}

// -------------------------------------------------
type Wlist struct {
	panel *taps.Panel
}

func setWeekDay(t time.Time, delDays int) []string {
	var listData []string
	for i := 0; i < C_DAYS-delDays; i++ {
		listData = append(listData, fmt.Sprintf("%s %02d", t.Weekday().String()[:3], t.Day()))
		t = t.AddDate(0, 0, 1)
	}
	return listData
}

func setDayTime(startList int) []string {
	var listData []string
	for i := startList; i < C_HOURS; i++ {
		listData = append(listData, fmt.Sprintf(" %02d:00", i))
	}
	return listData
}

var apptTBL map[string][]appt.Appointment

// ---------------------------------------
// Main
// ---------------------------------------
func (m *Wlist) doFormat(t time.Time, currentTime time.Time, databaseName, connectString string, delDays, startList int) {
	m.panel.Store(fmt.Sprintf("%04d/%02d", t.Year(), int(t.Month())), "L_YM")
	m.panel.StoreList(setWeekDay(t, delDays), "L_WEEKDAY")
	m.panel.StoreList(setDayTime(startList), "L_DAYTIME")

	for i := 0; i < C_DAYS-delDays; i++ {
		appts := apptTBL[t.Format(appt.DATE_FORMAT)]
		if appts == nil {
			appts = appt.GetApptToday(databaseName, connectString, t)
			apptTBL[t.Format(appt.DATE_FORMAT)] = appts
		}
		var listData []string
		for j := 0; j < C_HOURS; j++ {
			// apptlist.go
			hourAppts := getApptHour(appts, j)
			if hourAppts == nil {
				listData = append(listData, " ")
				continue
			}
			s := setApptID(hourAppts[0].Description, 10, hourAppts[0].Id)
			listData = append(listData, s)
		}

		m.panel.StoreList(listData[startList:], fmt.Sprintf("S_WLIST%02d", i+1))
		t = t.AddDate(0, 0, 1)
	}

	if m.panel.SelectFocus == 0 {
		n := m.panel.GetFirstListName(fmt.Sprintf("S_WLIST%02d", int(currentTime.Weekday()+1)))
		m.panel.SelectFocus = m.panel.GetFieldNumber(n)
	}
	m.panel.Say()
}

func (m *Wlist) resetTime(t time.Time) time.Time {
	return t.AddDate(0, 0, -int(t.Weekday()))
}

func (m *Wlist) Run(common *Common) string {
	if m.panel == nil {
		m.panel = WlistPanel()
	}
	apptTBL = map[string][]appt.Appointment{}
	t := m.resetTime(common.getCurrentTime())

	delDays := 0
	if common.cols < 70 {
		delDays++
	}
	if common.cols < 60 {
		delDays++
	}
	if common.cols < 51 {
		delDays++
	}

	m.panel.AddExitKey("S_WLIST01", "Left")

	m.panel.AddExitKey(fmt.Sprintf("S_WLIST%02d", C_DAYS-delDays), "Right")

	for i := 1; i <= C_DAYS-delDays; i++ {
		f := m.panel.GetDataField(fmt.Sprintf("L_WLIST%02d", i))
		if common.rows > C_HOURS+2 {
			f.Rows = C_HOURS
		}
		m.panel.AddExitKey(fmt.Sprintf("S_WLIST%02d", i), "Ctrl-T")
		m.panel.AddExitKey(fmt.Sprintf("S_WLIST%02d", i), "Delete")
		m.panel.AddExitKey(fmt.Sprintf("S_WLIST%02d", i), "F2")
		m.panel.AddExitKey(fmt.Sprintf("S_WLIST%02d", i), "F3")

		m.panel.AddExitKey(m.panel.GetListFieldName(fmt.Sprintf("S_WLIST%02d", i), 0), "Up")
		m.panel.AddExitKey(m.panel.GetListFieldName(fmt.Sprintf("S_WLIST%02d", i), taps.GetFieldY(9997)-1), "Down")
	}

	for i := C_DAYS - delDays + 1; i <= C_DAYS; i++ {
		m.panel.SetDisabled(fmt.Sprintf("L_WLIST%02d", i))
		m.panel.SetDisabled(fmt.Sprintf("S_WLIST%02d", i))
	}

	calendar := &Calendar{}
	godate := &GoDate{}
	detail := &Detail{}
	mlist := &Mlist{}
	delappt := &DelAppt{}
	find := &Find{}
	help := &Help{}

	startList := 0
	for {
		m.doFormat(t, common.getCurrentTime(), common.databaseName, common.connectString, delDays, startList)
		k, n := m.panel.Read()
		if k == tcell.KeyEscape {
			break
		}

		if k == tcell.KeyRight {
			t = t.AddDate(0, 0, 1)
			continue
		}

		if k == tcell.KeyLeft {
			t = t.AddDate(0, 0, -1)
			continue
		}

		if k == tcell.KeyDown {
			startList++
			continue
		}

		if k == tcell.KeyUp {
			if startList > 0 {
				startList--
			}
			continue
		}

		if k == tcell.KeyCtrlT {
			common.setCurrentTime()
			t = m.resetTime(common.currentTime)

			m.panel.SelectFocus = 0
			continue
		}

		if k == tcell.KeyCtrlP {
			t = t.AddDate(0, 0, -7)
			continue
		}

		if k == tcell.KeyCtrlN {
			t = t.AddDate(0, 0, 7)
			continue
		}

		if n == "H" || k == tcell.KeyF1 {
			help.Run(m.panel.GetHelp())
			continue
		}

		if (k == tcell.KeyF2 || k == tcell.KeyEnter) && strings.HasPrefix(n, "S_WLIST") {
			dayTime := setDayTime(startList)
			common.startTime = strings.TrimSpace(dayTime[m.panel.GetListFocus(n)])
			listNum, _ := strconv.Atoi(n[len("S_WLIST") : len("S_WLIST")+2])
			saveTime := t
			common.currentTime = t.AddDate(0, 0, listNum-1)
			common.apptID = extractApptID(m.panel.Field[m.panel.SelectFocus].Data)
			delete(apptTBL, common.currentTime.Format(appt.DATE_FORMAT))

			rs := detail.Run(common)
			if rs == "Q" {
				//break
			}

			t = saveTime
			continue
		}

		if k == tcell.KeyF3 || k == tcell.KeyDelete {
			common.apptID = extractApptID(m.panel.Field[m.panel.SelectFocus].Data)
			if common.apptID >= 0 {
				rs := delappt.Run(common, "Delete item ?")
				if rs == "OK" {
				}

				listNum, _ := strconv.Atoi(n[len("S_WLIST") : len("S_WLIST")+2])
				ts := t.AddDate(0, 0, listNum-1)
				delete(apptTBL, ts.Format(appt.DATE_FORMAT))
			}
			continue
		}

		if n == "F" || k == tcell.KeyF4 {
			rs := find.Run(common, "Find Appointment")
			if len(rs) > 0 {
				kk := strings.Split(rs, " ")
				t, _ = time.Parse(appt.DATE_FORMAT, kk[0])
			}

			continue
		}

		if n == "G" || k == tcell.KeyF5 {
			rs := godate.Run(common)
			if appt.CheckYMD(rs) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, rs)
			}
			t = m.resetTime(common.currentTime)
			continue
		}

		if n == "6" || k == tcell.KeyF6 {
			rs := calendar.Run(common)
			if appt.CheckYMD(rs) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, rs)
			}
			t = m.resetTime(common.currentTime)
			continue
		}

		if n == "M" || k == tcell.KeyF7 {
			rs := mlist.Run(common)
			if rs == "Q" {
				break
			}
			t = m.resetTime(common.currentTime)
			continue
		}

		if n == "T" || k == tcell.KeyF10 {
			return ""
		}

		if n == "Q" || k == tcell.KeyF12 {
			return "Q"
		}

	}
	return ""
}
