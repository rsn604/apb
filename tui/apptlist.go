package tui

import (
	"apb/appt"
	"fmt"
	"github.com/rsn604/taps"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
	"sort"
	"strconv"
	"strings"
)

const (
	NO_ERROR  = -1
	ID_SEP    = "@ID@:"
	C_DAYS    = 7
	C_WEEKS   = 5
	C_HOURS   = 24
	C_PREV    = "PREV"
	C_NEXT    = "NEXT"
	FIND_TODO = "Find Todo"
)

func ApptListPanel() *taps.Panel {
	var help = `Apptlist

Display today's Appt.

(1) Add Appt
 Move cursor over the time you want to add
  and press Enter.(or use the F2 key,
  or mouse to select the row.)
"Detail" screen will appear.

(2) Delete Appt
 Move the cursor over the app and press DEL (F3).
Follow "Confirmation" prompts.

(3) Key Operations
 The bottom two lines provide guidance.
String like "<A>" indicate that you can
access the function by pressing the "A" key
directly.
The function keys below perform the same
function.
You can also select the area enclosed in <> by mouse.

The following applies to all screens.
 "ESC"         : return to previous screen.
 "<Q>" ,F12    : exit or return to previous screen.
 "<TAB>"       : go next field.
 "<SHIFT-TAB>" : go previous field.

(4) View and Change Date
Change Day Units(next): CTRL-N (Right Arrow)
                (prev): CTRL-P (Left Arrow)
Return to Today       : CTRL-T
View Calendar         : "<G>" F5
View 6-Month Calendar : "<6>" F6

(5) Appt View and Search
View Monthly Appt : "<M>" F7
View Weekly Appt  : "<W>" F8
View ToDoList     : "<T>" F10
Find Appt         : "<F>" F4
`
	var styleMatrix = [][]string{
		{"APPT_H01", "yellow", "default"},
		{"CAL_YYYYMM", "yellow", "default"},
		{"PFKEY", "white", "default"},
		{"label", "aqua", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"list", "white", "default"},
		{"list_focus", "black", "aqua"},
		{"CAL", "white", "black"},
		{"CAL_HOLYDAY", "red", "black"},
		{"CAL_TODAY", "black", "green"},
		{"CAL_FOCUS", "black", "aqua"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 9999
EndY = 9999
ExitKey = ["F1", "F4", "F5", "F6", "F7", "F8", "F10", "F12", "Ctrl-T", "Ctrl-P", "Ctrl-N"]
# -------------------------------------------------
[[Field]]	
Name = "L_APPT_H01"
X = 1
Y = 1
FieldLen = 40
Style = "L_APPT_H01"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
Name = "S_APPT_LIST"
X = 1
Y = 2
FieldLen = 40
Rows = 9996
Style = "list, list_focus"
FieldType = "select"
ExitKey = ["Right", "Left", "Delete", "F2", "F3"]

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
#Data = "<A>"
Data = "Add"
X = 5
Y = 9998
FieldLen = 4
Style = "select, select_focus"
#FieldType = "select"
#FieldType = "label"

[[Field]]	
Name = "D"
#Data = "<D>"
Data = "Del"
X = 9
Y = 9998
FieldLen = 4
Style = "select, select_focus"
#FieldType = "select"
FieldType = "label"

[[Field]]
Name = "F"
Data = "<F>"
X = 13
Y = 9998
FieldLen = 4
Style ="select, select_focus"
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

[[Field]]	
Name = "W"
Data = "<W>"
X = 29
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "T"
Data = "<T>"
X = 33
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "Q"
Data = "<Q>"
X = 37
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
#Name = "L01"
Data = " F1  F2  F3  F4  F5  F6  F7  F8  F10 F12"
X = 0
Y = 9999
FieldLen = 41
Style = "PFKEY"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
X = 9976
Y = 1
FieldLen=0
Rows = 9996
Rect = true
Style = "list"
FieldType = "label"

[[Field]]	
X = 9977
Y = 10
FieldLen=0
COLS = 9996
Rect = true
Style = "list"
FieldType = "label"

[[Field]]	
Name = "CAL_YYYYMM"
X = 9980
Y = 2
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 9978
Y = 3
Data = "Su Mo Tu We Th Fr Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL"
X = 9978
Y = 4
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 1
Style = "CAL, CAL_FOCUS"
FieldType = "label"

[[Field]]	
X = 9980
Y = 11
Data = "Next Appointment"
Style = "label"
FieldType = "label"

[[Field]]	
Name = "L_NEXTAPPT"
X = 9978
Y = 12
#FieldLen = 19
Style = "list"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
X = 9977
Y = 13
FieldLen=0
COLS = 9996
Rect = true
Style = "list"
FieldType = "label"

[[Field]]	
#Name = "T01"
X = 9980
Y = 14
Data = "ToDo List"
Style = "label"
FieldType = "label"

[[Field]]
Name = "TODO_LIST"
X = 9978
Y = 15
FieldLen = 21
Rows = 5
Style = "list, list_focus"
#FieldType = "select"
FieldType = "label"
`
	return taps.NewPanel(doc, styleMatrix, help)
}

// -------------------------------------------------
type ApptList struct {
	panel *taps.Panel
}

/*
func PrintAppts(appts []appt.Appointment) {
	for _, ap := range appts {
		log.Printf("id:%d repeat:%d desc:%s sdate:%s stime:%s etime:%s location:%s ap.ConDays:%d frequency:%d DayNumber:%d DayPosition:%d DayOfWeek:%d Month:%d DateOfYear:%s RepeatSdate:%s RepeatEdate:%s\n", ap.Id, ap.RepeatType, ap.Description, ap.Sdate, ap.Stime, ap.Etime, ap.Location, ap.ConDays, ap.Frequency, ap.DayNumber, ap.DayPosition, ap.DayOfWeek, ap.Month, ap.DateOfYear, ap.RepeatSdate, ap.RepeatEdate)
	}
	log.Printf("---------------------------\n")
}
*/
// -------------------------------------------------
func getDaysOfYear(common *Common) (int, int) {
	t := common.getCurrentTime()
	y1 := time.Date(t.Year(), 1, 1, 0, 0, 0, 0, time.Local)
	d1 := int(t.Sub(y1).Hours()/24) + 1

	y2 := time.Date(t.Year(), 12, 31, 0, 0, 0, 0, time.Local)
	d2 := int(y2.Sub(t).Hours()/24) + 1
	return d1, d2
}

func setDayFocus(p *taps.Panel, day int) int {
	pos := p.GetFieldNumber(p.GetFirstListName("CAL"))
	sel := 0
	for i := pos; i < len(p.Field); i++ {
		if strings.HasPrefix(p.Field[i].Name, "CAL") {
			if ((i - pos) % C_DAYS) == 0 {
				p.ResetFieldStyle(p.Field[i].Name, "CAL_HOLYDAY, CAL_FOCUS")
				p.Field[i].Say()
			}
			if i == day+pos {
				sel = i
				taps.SetFocusedStyle(p.Field[i])
				p.Field[i].Say()
			}
		} else {
			break
		}
	}
	return sel
}

func setTodayMark(p *taps.Panel, day int) {
	pos := p.GetFieldNumber(p.GetFirstListName("CAL")) + day
	p.ResetFieldStyle(p.Field[pos].Name, "CAL_TODAY, CAL_FOCUS")
	//p.Field[pos].Say()
}

func resetTodayMark(p *taps.Panel, day int) {
	pos := p.GetFieldNumber(p.GetFirstListName("CAL")) + day
	p.ResetFieldStyle(p.Field[pos].Name, "CAL, CAL_FOCUS")
	//p.Field[pos].Say()
}

func getApptHour(appts []appt.Appointment, i int) []appt.Appointment {
	var s []appt.Appointment
	for _, ap := range appts {
		if ap.Stime[:2] == fmt.Sprintf("%02d", i) {
			s = append(s, ap)
		}
	}
	return s
}

func extractApptID(selectedData string) int {
	id := -1
	pos := strings.Index(selectedData, ID_SEP)
	if pos > 0 {
		x := selectedData[pos+len(ID_SEP):]
		id, _ = strconv.Atoi(x)
	}
	return id
}

func setApptID(s string, fieldLen, id int) string {
	x := 0
	ss := []rune(s)

	for i := 0; i < len(ss); i++ {
		x += runewidth.RuneWidth(ss[i])
	}

	for x < taps.GetFieldX(fieldLen) {
		s = s + " "
		x++
	}
	s = s + fmt.Sprintf(" @ID@:%05d", id)
	return s
}

// -------------------------------------------------
func getApptList(common *Common) []string {
	var listData []string

	appts := appt.GetApptToday(common.databaseName, common.connectString, common.getCurrentTime())
	//PrintAppts(appts)
	emark := " "
	etime2 := ""
	for i := 0; i < C_HOURS; i++ {
		//hourAppts := m.getApptHour(appts, i)
		hourAppts := getApptHour(appts, i)
		if emark != " " && etime2[:2] < fmt.Sprintf("%02d", i) {
			emark = " "
			etime2 = ""
		}
		if hourAppts == nil {
			listData = append(listData, fmt.Sprintf(" %s %02d:00", emark, i))
			continue
		}

		for _, hourAppt := range hourAppts {
			if hourAppt.Stime[3:5] != "00" {
				listData = append(listData, fmt.Sprintf("   %02d:00  ", i))
			}
			emark = "|"
			etime2 = hourAppt.Etime

			//s := m.setApptID(fmt.Sprintf(" %s %s %s @%s ", emark, hourAppt.Stime, hourAppt.Description, hourAppt.Location), hourAppt.Id)
			s := setApptID(fmt.Sprintf(" %s %s %s @%s ", emark, hourAppt.Stime, hourAppt.Description, hourAppt.Location), 9975, hourAppt.Id)
			listData = append(listData, s)
		}

	}

	return listData
}

// -------------------------------------------------
func getNextAppoint(common *Common) (string, appt.Appointment) {
	//apptMap := appt.GetNextAppts(common.databaseName, common.connectString, common.getCurrentTime())   // apptchecker.go
	apptMap := appt.GetNextAppts(common.databaseName, common.connectString, time.Now()) // apptchecker.go

	keys := []string{}
	for k := range apptMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	//@@@@
	if len(keys) == 0 {
		return "", appt.Appointment{}
	} else {
		return keys[0], apptMap[keys[0]]
	}
}

// ---------------------------------------
// Main
// ---------------------------------------
func (m *ApptList) doFormat(common *Common) {
	t := common.getCurrentTime()

	_, week := t.ISOWeek()
	d1, d2 := getDaysOfYear(common)
	m.panel.Store(fmt.Sprintf("Week %02d  %4d/%02d/%02d %s  %03d/%03d", week, t.Year(), int(t.Month()), t.Day(), t.Weekday(), d1, d2), "L_APPT_H01")

	listData, focusday := GetMonthCalendar(t)
	m.panel.StoreList(listData, "CAL")
	m.panel.Store(fmt.Sprintf("%04d/%02d  %s", t.Year(), int(t.Month()), t.Month()), "CAL_YYYYMM")
	m.panel.StoreList(getTodoList(common), "TODO_LIST") //todolist.go
	m.panel.StoreList(getApptList(common), "S_APPT_LIST")

	dd, ap := getNextAppoint(common)
	if dd != "" {
		m.panel.Store(dd[5:16]+" "+ap.Description, "L_NEXTAPPT")
	}

	ts := time.Now()
	_, today := GetMonthCalendar(ts)
	if t.Year() == ts.Year() && t.Month() == ts.Month() {
		setTodayMark(m.panel, today)
	} else {
		resetTodayMark(m.panel, today)
	}

	m.panel.Say()
	setDayFocus(m.panel, focusday)
}

func (m *ApptList) Run(common *Common) {
	if m.panel == nil {
		m.panel = ApptListPanel()
	}
	calendar := &Calendar{}
	godate := &GoDate{}
	detail := &Detail{}
	todolist := &ToDoList{}
	wlist := &Wlist{}
	mlist := &Mlist{}
	delappt := &DelAppt{}
	find := &Find{}
	help := &Help{}

	m.panel.ModifyFieldLen("L_APPT_H01", taps.GetFieldX(9975))
	m.panel.ModifyFieldLen("S_APPT_LIST", taps.GetFieldX(9975))
	for {
		m.doFormat(common)
		k, n := m.panel.Read()
		if k == tcell.KeyEscape {
			break
		}
		if k == tcell.KeyRight || k == tcell.KeyCtrlN {
			common.currentTime = common.currentTime.AddDate(0, 0, 1)
			continue
		}
		if k == tcell.KeyLeft || k == tcell.KeyCtrlP {
			common.currentTime = common.currentTime.AddDate(0, 0, -1)
			continue
		}

		if k == tcell.KeyCtrlT {
			common.setCurrentTime()
			continue
		}

		if n == "H" || k == tcell.KeyF1 {
			help.Run(m.panel.GetHelp())
			continue
		}

		if (k == tcell.KeyF2 || k == tcell.KeyEnter) && strings.HasPrefix(n, "S_APPT_LIST") {
			s := strings.Split(m.panel.Field[m.panel.SelectFocus].Data, " ")
			common.startTime = s[len(s)-1]
			common.apptID = extractApptID(m.panel.Field[m.panel.SelectFocus].Data)

			rs := detail.Run(common)
			if rs == "Q" {
				//break
			}
			continue
		}

		if k == tcell.KeyF3 || k == tcell.KeyDelete {
			common.apptID = extractApptID(m.panel.Field[m.panel.SelectFocus].Data)
			if common.apptID >= 0 {
				rs := delappt.Run(common, "Delete item ?")
				if rs == "OK" {
				}
			}
			continue
		}

		if n == "F" || k == tcell.KeyF4 {
			rs := find.Run(common, "Find Appointment")
			if len(rs) > 0 {
				kk := strings.Split(rs, " ")
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, kk[0])

			}
			continue
		}

		if n == "G" || k == tcell.KeyF5 {
			rs := godate.Run(common)
			if appt.CheckYMD(rs) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, rs)
			}
			continue
		}

		if n == "6" || k == tcell.KeyF6 {
			rs := calendar.Run(common)
			if appt.CheckYMD(rs) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, rs)
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

		if n == "T" || k == tcell.KeyF10 {
			todolist.Run(common)
			continue
		}

		if n == "Q" || k == tcell.KeyF12 {
			break
		}

	}
}
