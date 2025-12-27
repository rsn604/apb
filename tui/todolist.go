package tui

import (
	"apb/appt"
	"fmt"
	"time"

	"github.com/rsn604/taps"
	//"github.com/mattn/go-runewidth"
	"github.com/gdamore/tcell/v2"
	"strings"
)

func ToDoListPanel() *taps.Panel {
	var help = `ToDolist

Display ToDo .

(1) Add ToDo
 Move cursor over blank line
  and press Enter.( or use the F2 key,
  or mouse to select the row.)
"ToDo Detail" screen will appear.

(2) Change Status
 Press <F9> to change status. 
"F" : Finished
"!" : Delayed

The following operations are the same
  as "Apptlist".
(2) Delete ToDo
(3) Key Operations
(4) View and Change Date
(5) ToDO View and Search
`
	var styleMatrix = [][]string{
		{"APPT_H01", "yellow", "default"},
		{"CAL_MMYYYY", "yellow", "default"},
		{"PFKEY", "white", "default"},
		{"label", "aqua", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"list", "white", "default"},
		{"list_focus", "black", "aqua"},
		{"CAL", "white", "default"},
		{"CAL_HOLYDAY", "red", "default"},
		{"CAL_TODAY", "black", "green"},
		{"CAL_FOCUS", "black", "aqua"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 9999
EndY = 9999
#Rect = true
ExitKey = ["F1", "F2", "F3", "F4", "F5", "F6", "F7", "F8", "F9", "F12", "Ctrl-T", "Ctrl-P", "Ctrl-N"]

# -------------------------------------------------
[[Field]]	
Name = "TODO_H01"
X = 1
Y = 1
FieldLen = 40
Style = "TODO_H01"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
Name = "S_TODO_LIST"
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

[[Field]]	
Name = "W"
Data = "<W>"
X = 29
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "#"
Data = "Fin"
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
Data = " F1  F2  F3  F4  F5  F6  F7  F8  F9  F12"
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
Name = "CAL_MMYYYY"
X = 9980
Y = 2
FieldLen = 18
Style = "CAL_MMYYYY"
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
#Name = "T01"
X = 9980
Y = 11
Data = "Next Appointment"
Style = "label"
FieldType = "label"

[[Field]]	
Name = "L_NEXTAPPT"
X = 9979
Y = 12
#FieldLen = 19
Style = "list"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
X = 9977
Y = 9991
FieldLen=0
COLS = 9996
Rect = true
Style = "list"
FieldType = "label"

[[Field]]	
#Name = "T01"
X = 9980
Y = 9992
Data = "Appointments"
Style = "label"
FieldType = "label"

[[Field]]
Name = "APPT_LIST"
X = 9978
Y = 9993
FieldLen = 21
Rows = 4
Style = "list, list_focus"
FieldType = "select"
#FieldType = "label"
`
	return taps.NewPanel(doc, styleMatrix, help)
}

// -------------------------------------------------
type ToDoList struct {
	panel *taps.Panel
}

func getTodoList(common *Common) []string {
	var listData []string
	m := make(map[string]int)

	appts := appt.GetTodoToday(common.databaseName, common.connectString, common.getCurrentTime())

	ymd := common.currentTime.Format(appt.DATE_FORMAT)
	ymd_today := time.Now().Format(appt.DATE_FORMAT)
	x := ""

	for _, ap := range appts {
		if ymd > ymd_today && ap.RepeatType == appt.GetRepeatType("NONE") {
			continue
		}

		priority := ap.Priority
		if ap.RepeatType == appt.GetRepeatType("NONE") {
			dueDate := appt.GetDueDate(ap)
			if ymd == dueDate {
				x = "!"
			} else if ymd > dueDate {
				x = "#"
			} else {
				x = " "
			}
		}
		if ap.TodoStatus == "F" {
			x = "F"
			priority = 99
		}
		//@@@@
		s := setApptID(fmt.Sprintf("%2d %s %s @%s", ap.Priority, x, ap.Description, appt.GetDueDate(ap)), 9975, ap.Id)
		m[s] = priority
	}
	listData = appt.SortMap(m)
	return listData
}

// ---------------------------------------
func (m *ToDoList) update(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)
	u := manager.GetApptByID(common.apptID)
	if u.TodoStatus == "F" {
		u.TodoStatus = ""
	} else {
		u.TodoStatus = "F"
	}
	manager.UpdateAppt(u, common.apptID)
	manager.Close()
}

func (m *ToDoList) doFormat(common *Common) {
	t := common.getCurrentTime()
	_, week := t.ISOWeek()
	d1, d2 := getDaysOfYear(common)
	m.panel.Store(fmt.Sprintf("Week %02d  %4d/%02d/%02d %s  %03d/%03d", week, t.Year(), int(t.Month()), t.Day(), t.Weekday(), d1, d2), "TODO_H01")

	listData, focusday := GetMonthCalendar(t)
	m.panel.StoreList(listData, "CAL")
	m.panel.Store(fmt.Sprintf("%04d/%02d  %s", t.Year(), int(t.Month()), t.Month()), "CAL_MMYYYY")

	todoList := getTodoList(common)
	todoList = append(todoList, setApptID("     ", 9975, -1))
	m.panel.StoreList(todoList, "S_TODO_LIST")

	//@@@@
	// find.go
	keys, _ := findAllAppts(common, "", C_NEXT, appt.STATE_APPT)
	if len(keys) > 0 {
		m.panel.StoreList(keys[:4], "APPT_LIST")
	}
	// apptlist.go
	dd, ap := getNextAppoint(common)
	if dd != "" {
		m.panel.Store(dd[5:16]+" "+ap.Description, "L_NEXTAPPT")
	}

	ts := time.Now()
	_, today := GetMonthCalendar(ts)
	if t.Year() == ts.Year() && t.Month() == ts.Month() {
		setTodayMark(m.panel, today) // apptlist.go
	} else {
		resetTodayMark(m.panel, today)
	}

	m.panel.Say()
	setDayFocus(m.panel, focusday) // apptlist.go

}

func (m *ToDoList) Run(common *Common) {
	if m.panel == nil {
		m.panel = ToDoListPanel()
	}
	calendar := &Calendar{}
	godate := &GoDate{}
	todo := &ToDo{}
	wlist := &Wlist{}
	mlist := &Mlist{}
	delappt := &DelAppt{}
	find := &Find{}
	help := &Help{}

	m.panel.ModifyFieldLen("TODO_H01", taps.GetFieldX(9975))
	m.panel.ModifyFieldLen("S_TODO_LIST", taps.GetFieldX(9975))
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

		if (k == tcell.KeyF2 || k == tcell.KeyEnter) && strings.HasPrefix(n, "S_TODO_LIST") {
			s := strings.Split(m.panel.Field[m.panel.SelectFocus].Data, " ")
			common.startTime = s[len(s)-1]
			common.apptID = extractApptID(m.panel.Field[m.panel.SelectFocus].Data)
			rs := todo.Run(common)
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
			rs := find.Run(common, FIND_TODO)
			if len(rs) > 0 {
				kk := strings.Split(rs, " ")
				//@@@@@
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

		//if n == "#" || k == tcell.KeyF9 {
		if k == tcell.KeyF9 {
			common.apptID = extractApptID(m.panel.Field[m.panel.SelectFocus].Data)
			if common.apptID >= 0 {
				m.update(common)
			}
			continue
		}

		if n == "Q" || k == tcell.KeyF12 {
			break
		}
	}
}
