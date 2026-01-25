package tui

import (
	"apb/appt"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strconv"
	"strings"
)

func DetailPanel() *taps.Panel {
	var help = `Detail

Add or update Appt.

(1) Add new Appt
-"Start Time", "End Time", "Start Date"
The selected field is displayed. You can
overwrite and change it.

-"Description", "Location", "Note"
Enter Description, Location, and Notes for Appt.

-"No. Consecutive Days"
If you want to set consecutive days, specify the number of days.

Enter the information and select "<A>" .
Once the addition is complete, you can
select "<R>" and "<X>".

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
		{"errmsg", "red", "default"},
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
Data = "Start Time"
X = 0
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_STIME"
X = 13
Y = 3
FieldLen = 5
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]	
Data = "Start Date"
X = 25
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_SDATE"
X = 37
Y = 3
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]	
Data = "End Time"
X = 0
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_ETIME"
X = 13
Y = 4
FieldLen = 5
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]	
Data = "No. Consecutive Days"
X = 25
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_CONDAYS"
X = 47
Y = 4
FieldLen = 2
DataLen = 2
Attr = "N"
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]	
Data = "Location"
X = 0
Y = 6
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_LOCATION"
X = 13
Y = 6
FieldLen = 20
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
#[[Field]]	
#Data = "Views"
#X = 0
#Y = 8
#Style = "label"
#FieldType = "label"

[[Field]]	
Data = "ALARM(Not supported)"
X = 25
Y = 8
Style = "label"
FieldType = "label"

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
Rows = 4
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
#Name = "N"
#Data = "<N>"
#X = 9
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

#[[Field]]	
#Name = "F"
#Data = "<F>"
#X = 13
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

#[[Field]]	
#Name = "P"
#Data = "<P>"
#X = 17
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
type Detail struct {
	panel *taps.Panel
}

func (m *Detail) errCheck(common *Common) (string, int) {
	sDate := m.panel.Get("E_SDATE")
	sTime := m.panel.Get("E_STIME")
	eTime := m.panel.Get("E_ETIME")

	errMsg := ""

	if !appt.CheckHHMM(sTime) {
		errMsg = "D.ER Start Time format error ."
		return errMsg, m.panel.GetFieldNumber("E_STIME")
	}

	if !appt.CheckYMD(sDate) {
		errMsg = "D.ER Start Date format error ."
		return errMsg, m.panel.GetFieldNumber("E_SDATE")
	}

	if !appt.CheckHHMM(eTime) {
		errMsg = "D.ER End Time format error ."
		return errMsg, m.panel.GetFieldNumber("E_ETIME")
	}

	if sTime >= eTime {
		errMsg = "D.ER Start/End Time error ."
		return errMsg, m.panel.GetFieldNumber("E_STIME")
	}
	return errMsg, NO_ERROR
}

func (m *Detail) doFormat(common *Common) {
	t := common.getCurrentTime()
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)

	id := common.apptID
	if id >= 0 {
		u := manager.GetApptByID(id)
		m.panel.Store(u.Description, "E_DESCRIPTION")
		//@@@@
		//m.panel.Store(t.Format(appt.DATE_FORMAT), "E_SDATE")
		m.panel.Store(appt.GetSdate(u), "E_SDATE")
		m.panel.Store(u.Stime, "E_STIME")
		m.panel.Store(u.Etime, "E_ETIME")
		m.panel.Store(u.Location, "E_LOCATION")
		m.panel.Store(fmt.Sprintf("%2d", u.ConDays), "E_CONDAYS")
		if len(u.Note) > 0 {
			m.panel.StoreList(u.Note, "E_NOTE")
		} else {
			m.panel.StoreList([]string{""}, "E_NOTE")
		}
		m.panel.Store(appt.GetRepeatString(u.RepeatType), "L_REPEATTYPE")
		m.panel.GetDataField("R").Enabled()
		m.panel.GetDataField("X").Enabled()

	} else {
		m.panel.Store("", "E_DESCRIPTION")
		m.panel.Store("0", "E_CONDAYS")
		m.panel.Store("", "E_LOCATION")
		//m.panel.Store("", "ERR_MSG")

		m.panel.Store(fmt.Sprintf(t.Format(appt.DATE_FORMAT)), "E_SDATE")
		//@@@@@
		hhmm := common.startTime
		m.panel.Store(hhmm, "E_STIME")
		hh, _ := strconv.Atoi(strings.TrimSpace(strings.Split(hhmm, ":")[0]))
		hh++
		m.panel.Store(fmt.Sprintf("%02d:00", hh), "E_ETIME")
		m.panel.StoreList([]string{""}, "E_NOTE")
		m.panel.Store("NONE", "L_REPEATTYPE")
		m.panel.GetDataField("R").Disabled()
		m.panel.GetDataField("X").Disabled()
	}

	manager.Close()
}

func (m *Detail) Run(common *Common) string {
	if m.panel == nil {
		m.panel = DetailPanel()
	}
	repeat := &RepeatOption{}
	godate := &GoDate{}
	help := &Help{}

	m.doFormat(common)
	for {
		m.panel.Say()
		k, n := m.panel.Read()

		m.panel.Store("", "ERR_MSG")
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
				u.RepeatType = 0
			}
			u.Description = m.panel.Get("E_DESCRIPTION")
			u.Stime = m.panel.Get("E_STIME")
			u.Etime = m.panel.Get("E_ETIME")
			u.Location = m.panel.Get("E_LOCATION")
			u.ConDays, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_CONDAYS")))
			u.Note = m.panel.GetList("E_NOTE")

			msg := ""
			if id >= 0 && (n == "X" || k == tcell.KeyF10) {
				appt.UpdateSdate(&u, m.panel.Get("E_SDATE"))
				manager.UpdateAppt(u, id)
				msg = "D.MG Record updated ."
			} else {
				appt.UpdateSdate(&u, m.panel.Get("E_SDATE"))
				common.apptID, _ = manager.AddAppt(u)
				msg = "D.MG Record added ."

			}
			manager.Close()
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
