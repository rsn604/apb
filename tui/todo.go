package tui

import (
	"apb/appt"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strconv"
	"strings"
)

func ToDoPanel() *taps.Panel {
	var help = `ToDo Detail

Add or update ToDo.

(1) Add new ToDo
-"Start Date", "Due Date"
Default value is displayed. You can
 overwrite and change it.

-"Description", "Priority", "Note"
Enter Description, Priority and Notes for ToDo.

(2) Update Appt
"<X>" F9: Update with on-screen data

(3) "Repeat" settings
"<R>" F8: Go to the "Repeat" settings,
select one of the following options:

NoRepeat : Cancel "Repeat" function.
Daily    : Go to "Daily Repeat" settings.
Monthly  : Go to "Monthly Repeat" settings.
Yearly   : Go to "Yearly Repeat" settings.
Custom   : Go to "Custom Repeat" settings.

(4) Other
"<G>" F6: Calendar View -> Assign to
                           "Start Date".
"ESC" "<Q>" F12 : Go to previous screen

`
	var styleMatrix = [][]string{
		{"label", "aqua", "default"},
		{"status", "yellow", "default"},
		{"linerect", "white", "default"},
		{"edit", "white, underline", "black"},
		{"edit_focus", "yellow", "black"},
		{"PFKEY", "white", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"note", "white", "black"},
		{"note_focus", "yellow,underline", "black"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 9999
EndY = 9999
ExitKey = ["F1", "F2", "F3", "F4", "F5", "F6", "F8", "F9", "F10", "F12"]

# -------------------------------------------------
[[Field]]	
Data = "Description"
X = 0
Y = 1
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_DESCRIPTION"
X = 13
Y = 1
FieldLen = 30
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]	
Data = "Start Date"
X = 0
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_SDATE"
X = 13
Y = 3
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]	
Data = "Due Date"
X = 0
Y = 5
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_DUEDATE"
X = 13
Y = 5
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

#[[Field]]	
#Data = "Carry Forward"
#X = 25
#Y = 5
#Style = "label"
#FieldType = "label"

# -------------------------------------------------
[[Field]]	
Data = "Priority"
X = 0
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_PRIORITY"
X = 13
Y = 7
FieldLen = 2
DataLen = 2
Attr = "N"
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]	
Data = "Repeat Status : "
X = 0
Y = 9
Style = "label"
FieldType = "label"

[[Field]]	
Name = "L_REPEATTYPE"
X = 17
Y = 9
Style = "status"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
#Name = "LINERECT"
X = 6
Y = 11
FieldLen=9993
Rows = 6
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = "Note"
X = 0
Y = 12
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_NOTE"
X = 7
Y = 12
FieldLen=9990
Rows = 3
Style = "note, note_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]
Name = "ERR_MSG"
X = 1
Y = 9997
FieldLen = 40
Style = "errmsg"
FieldType = "label"

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
Data = "<A>"
X = 5
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

#[[Field]]	
#Name = "F"
#Data = "<F>"
#X = 13
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

[[Field]]	
Name = "G"
Data = "<G>"
X = 21
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "R"
Data = "<R>"
X = 29
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "X"
Data = "<X>"
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
type ToDo struct {
	panel *taps.Panel
}

func (m *ToDo) errCheck(common *Common) (string, int) {
	sDate := m.panel.Get("E_SDATE")
	n0 := m.panel.GetFieldNumber("E_SDATE")
	dDate := m.panel.Get("E_DUEDATE")
	n1 := m.panel.GetFieldNumber("E_DUEDATE")
	errMsg := ""

	if !appt.CheckYMD(sDate) {
		errMsg = "D.ER Start Date format error ."
		return errMsg, n0
	}

	if !appt.CheckYMD(dDate) {
		errMsg = "D.ER Due Date format error ."
		return errMsg, n1
	}

	if sDate >= dDate {
		errMsg = "D.ER Start/Due Date error ."
		return errMsg, n0
	}
	return errMsg, NO_ERROR
}

func (m *ToDo) doFormat(common *Common) {
	t := common.getCurrentTime()
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)

	id := common.apptID
	if id >= 0 {
		u := manager.GetApptByID(id)
		m.panel.Store(u.Description, "E_DESCRIPTION")
		//@@@@
		m.panel.Store(appt.GetSdate(u), "E_SDATE")
		//@@
		m.panel.Store(appt.GetDueDate(u), "E_DUEDATE")
		//m.panel.Store(u.Location, "E_LOCATION")
		m.panel.Store(fmt.Sprintf("%2d", u.Priority), "E_PRIORITY")

		if len(u.Note) > 0 {
			m.panel.StoreList(u.Note, "E_NOTE")
		} else {
			m.panel.StoreList([]string{""}, "E_NOTE")
		}
		m.panel.Store(appt.GetRepeatString(u.RepeatType), "L_REPEATTYPE")
	} else {
		m.panel.Store("", "E_DESCRIPTION")
		m.panel.Store(fmt.Sprintf(t.Format(appt.DATE_FORMAT)), "E_SDATE")
		m.panel.Store(fmt.Sprintf(t.AddDate(0, 0, 7).Format(appt.DATE_FORMAT)), "E_DUEDATE")
		m.panel.Store(fmt.Sprintf("%2d", 0), "E_PRIORITY")

		m.panel.StoreList([]string{""}, "E_NOTE")
		m.panel.Store("NONE", "L_REPEATTYPE")
	}

	manager.Close()
}

func (m *ToDo) Run(common *Common) string {
	if m.panel == nil {
		m.panel = ToDoPanel()
	}
	repeat := &RepeatOption{}
	godate := &GoDate{}
	help := &Help{}

	m.doFormat(common)
	for {

		m.panel.Say()
		k, n := m.panel.Read()
		if k == tcell.KeyEscape {
			break
		}

		if n == "H" || k == tcell.KeyF1 {
			help.Run(m.panel.GetHelp())
			continue
		}

		if n == "G" || k == tcell.KeyF6 {
			rs := godate.Run(common)
			if appt.CheckYMD(rs) {
				m.panel.Store(rs, "E_SDATE")
			}
			continue
		}

		if n == "R" || k == tcell.KeyF8 {
			id := common.apptID
			if id >= 0 {
				rs := repeat.Run(common)
				if rs == "Q" {
					//break
				}
				manager := appt.GetManager(common.databaseName)
				manager.Connect(common.databaseName, common.connectString)
				u := manager.GetApptByID(id)
				m.panel.Store(appt.GetRepeatString(u.RepeatType), "L_REPEATTYPE")
				manager.Close()

			} else {
				m.panel.Store("D.ER Record not added .", "ERR_MSG")
			}
			continue
		}

		if n == "A" || n == "X" || k == tcell.KeyF2 || k == tcell.KeyF10 {
			errMsg, num := m.errCheck(common)
			if num > NO_ERROR {
				m.panel.Store(errMsg, "ERR_MSG")
				m.panel.SelectFocus = num
				continue
			}
			manager := appt.GetManager(common.databaseName)
			manager.Connect(common.databaseName, common.connectString)

			id := common.apptID
			var u appt.Appointment
			if id >= 0 && (n == "X" || k == tcell.KeyF10) {
				u = manager.GetApptByID(id)
			} else {
				u = appt.Appointment{}
				u.State = appt.STATE_TODO
				u.RepeatType = 0
			}
			u.Description = m.panel.Get("E_DESCRIPTION")

			appt.UpdateSdate(&u, m.panel.Get("E_SDATE"))
			//@@@@
			appt.UpdateDueDate(&u, m.panel.Get("E_DUEDATE"))

			u.Priority, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_PRIORITY")))
			u.Note = m.panel.GetList("E_NOTE")

			msg := ""
			if id >= 0 && (n == "X" || k == tcell.KeyF10) {
				manager.UpdateAppt(u, id)
				msg = "D.MG Record updated ."
			} else {
				common.apptID, _ = manager.AddAppt(u)
				msg = "D.MG Record added ."

			}
			manager.Close()
			/*
				if n == "X" || k == tcell.KeyF10 {
					break
				}
			*/
			m.panel.Store(msg, "ERR_MSG")
			m.doFormat(common)
			continue
		}

		if n == "Q" || k == tcell.KeyF12 {
			return "Q"
		}

	}
	return ""
}
