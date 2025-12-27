package tui

import (
	"fmt"
	"time"

	"apb/appt"
	"github.com/rsn604/taps"

	"github.com/gdamore/tcell/v2"
	"strings"
)

func CalendarPanel() *taps.Panel {

	var styleMatrix = [][]string{
		{"APPT_H01", "yellow", "default"},
		{"CAL_YYYYMM", "yellow", "default"},
		{"PFKEY", "white", "default"},
		{"label", "aqua", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"list", "white", "default"},
		{"list_focus", "black", "aqua"},
		{"CAL", "white", "default"},
		{"CAL_HOLYDAY", "red", "default"},
		{"CAL_TODAY", "green, underline", "default"},
		{"CAL_FOCUS", "black", "aqua"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 9999
EndY = 9999
#Rect = true
ExitKey = ["F1", "F3", "F4", "F5", "F7", "F8",  "F10", "F12", "Ctrl-P", "Ctrl-N"]

# -------------------------------------------------
#[[Field]]	
#Name = "APPT_H01"
#X = 1
#Y = 1
#FieldLen = 40
#Style = "APPT_H01"
#FieldType = "label"

# -------------------------------------------------
[[Field]]	
Name = "CAL01_YYYYMM"
X = 3
Y = 0
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 2
Y = 1
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL01"
X = 2
Y = 2
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Name = "CAL02_YYYYMM"
X = 26
Y = 0
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 25
Y = 1
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL02"
X = 25
Y = 2
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Name = "CAL03_YYYYMM"
X = 49
Y = 0
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 48
Y = 1
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL03"
X = 48
Y = 2
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Name = "CAL04_YYYYMM"
X = 3
Y = 9
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 2
Y = 10
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL04"
X = 2
Y = 11
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Name = "CAL05_YYYYMM"
X = 26
Y = 9
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 25
Y = 10
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL05"
X = 25
Y = 11
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Name = "CAL06_YYYYMM"
X = 49
Y = 9
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 48
Y = 10
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL06"
X = 48
Y = 11
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
#[[Field]]
#Name = "H"
#Data = "<H>"
#X = 1
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

[[Field]]	
Name = "-"
Data = "<->"
X = 9
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "+"
Data = "<+>"
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
Name = "M"
Data = "<M>"
X = 25
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "W"
Data = "<W>"
X = 29
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

#[[Field]]	
#Name = "A"
#Data = "<A>"
#X = 33
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

[[Field]]	
Name = "T"
Data = "<T>"
X = 37
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "Q"
Data = "<Q>"
X = 41
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "L01"
Data = " F1  F2  F3  F4  F5  F6  F7  F8  F9  F10 F12"
X = 0
Y = 9999
#FieldLen = 41
Style = "PFKEY"
FieldType = "label"

`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type Calendar struct {
	panel *taps.Panel
}

func GetMonthCalendar(t time.Time) ([]string, int) {
	var listData []string
	today := 0
	firstOfMonth := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	firstWeekday := firstOfMonth.Weekday()
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	daysInMonth := lastOfMonth.Day()

	for i := 0; i < int(firstWeekday); i++ {
		listData = append(listData, "")
		today++
	}
	for i := 1; i <= daysInMonth; i++ {
		if i == t.Day() {
			today = today + i - 1
		}
		listData = append(listData, fmt.Sprintf("%2d", i))
	}
	return listData, today
}

func (m *Calendar) setDayFocus(name string, day int, flag bool) {
	pos := m.panel.GetFieldNumber(m.panel.GetFirstListName(name))
	for i := pos; i < len(m.panel.Field); i++ {
		if strings.HasPrefix(m.panel.Field[i].Name, name) {
			if ((i - pos) % C_DAYS) == 0 {
				m.panel.ResetFieldStyle(m.panel.Field[i].Name, "CAL_HOLYDAY, CAL_FOCUS")
				m.panel.Field[i].Say()
			}
			if (i == day+pos) && flag {
				m.panel.SelectFocus = i
			}
		} else {
			break
		}
	}
}

// ---------------------------------------
// Main
// ---------------------------------------
func (m *Calendar) doFormat(common *Common) {
	t := common.getCurrentTime()
	flag := true
	for i := 1; i <= 6; i++ {
		listData, today := GetMonthCalendar(t)
		name := fmt.Sprintf("CAL%02d", i)
		m.panel.StoreList(listData, name)
		m.panel.Store(fmt.Sprintf("%04d/%02d  %s", t.Year(), int(t.Month()), t.Month()), name+"_YYYYMM")
		//@@@@
		t = nextMonth(t)
		if i > 1 {
			flag = false
		}
		m.setDayFocus(name, today, flag)
	}
	m.panel.Say()
}

func (m *Calendar) Run(common *Common) string {
	if m.panel == nil {
		m.panel = CalendarPanel()
	}
	godate := &GoDate{}
	mlist := &Mlist{}
	wlist := &Wlist{}
	todolist := &ToDoList{}

	for {
		m.doFormat(common)
		k, n := m.panel.Read()
		if k == tcell.KeyEscape {
			break
		}

		if k == tcell.KeyEnter && len(n) > 3 && n[:4] == "CAL0" {
			dd := strings.TrimSpace(m.panel.Get(n))
			if len(dd) < 2 {
				dd = "0" + dd
			}
			ymd := m.panel.Get(n[:5] + "_YYYYMM")[:7] + "/" + dd
			if appt.CheckYMD(ymd) {
				return ymd
			}

			continue
		}

		if k == tcell.KeyCtrlP {
			for i := 0; i < 6; i++ {
				common.currentTime = prevMonth(common.currentTime) //godate.go
			}
			continue
		}

		if k == tcell.KeyCtrlN {
			for i := 0; i < 6; i++ {
				common.currentTime = nextMonth(common.currentTime) //godate.go
			}
			continue
		}

		if n == "-" || k == tcell.KeyF3 {
			common.currentTime = prevMonth(common.currentTime) //godate.go
			continue
		}

		if n == "+" || k == tcell.KeyF4 {
			common.currentTime = nextMonth(common.currentTime) //godate.go
			continue
		}

		if n == "H" || k == tcell.KeyF1 {
			//break
		}

		if n == "G" || k == tcell.KeyF5 {
			rs := godate.Run(common)
			if rs == "Q" {
				//break
			}
			continue
		}

		if n == "M" || k == tcell.KeyF7 {
			rs := mlist.Run(common)
			if rs == "Q" {
				break
			}
			continue
		}

		if n == "W" || k == tcell.KeyF8 {
			rs := wlist.Run(common)
			if rs == "Q" {
				break
			}
			continue
		}

		if n == "A" || k == tcell.KeyF9 {
			break
		}

		if n == "T" || k == tcell.KeyF10 {
			todolist.Run(common)
			continue
		}

		if n == "Q" || k == tcell.KeyF12 {
			return "Q"
		}

	}
	return ""
}
