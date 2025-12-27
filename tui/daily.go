package tui

import (
	"fmt"
	//"time"
	"apb/appt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strconv"
	"strings"
)

func DailyPanel() *taps.Panel {
	var help = `Repeat Daily

Configure "Repeat Daily" settings.

(1) Frequency
- Repeat Every x Day(s):
  Specify the number of days between each interval.
  To run every three days, set this to "3."

(2) Duration
- Starting : Specify the start date.
- Ending   : Specify the end date.

(3) Key Operation
Commit "Repeat Daily" settings : "<X>" F10

"ESC"     -> Return to the previous screen.
"<Q>" F12 -> Go to the "Details" screen.

`
	var styleMatrix = [][]string{
		{"errmsg", "red", "default"},
		{"label", "aqua", "default"},
		{"title", "yellow", "default"},
		{"linerect", "white", "default"},
		{"edit", "white, underline", "black"},
		{"edit_focus", "yellow", "black"},
		{"PFKEY", "white", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
	}

	var doc = `
StartX = 0
StartY = 3
EndX = 9999
EndY = 9999
ExitKey = ["F1", "F2", "F3", "F4", "F5", "F6", "F8", "F9", "F10", "F12"]

# -------------------------------------------------
[[Field]]	
X = 1
Y = 0
FieldLen = 30
Rows = 2
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Frequency "
X = 3
Y = 0
Style = "title"
FieldType = "label"

[[Field]]	
Data = "Repeat Every"
X = 3
Y = 1
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_FREQUENCY"
X = 17
Y = 1
FieldLen = 2
DataLen = 2
Attr = "N"
Style = "edit, edit_focus"
FieldType = "edit"
ExitKey = ["Up"]

[[Field]]	
Data = "Day(s)"
X = 20
Y = 1
Style = "label"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
X = 1
Y = 3
FieldLen = 30
Rows = 3
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Duration "
X = 3
Y = 3
Style = "title"
FieldType = "label"

[[Field]]	
Data = "Starting:"
X = 3
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_RSDATE"
X = 17
Y = 4
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]	
Data = "Ending:"
X = 3
Y = 5
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_REDATE"
X = 17
Y = 5
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

# -------------------------------------------------
[[Field]]
Name = "ERR_MSG"
X = 1
Y = 9994
FieldLen = 40
Style = "errmsg"
FieldType = "label"

# -------------------------------------------------
[[Field]]
Name = "H"
Data = "<H>"
X = 1
Y = 9995
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

#[[Field]]	
#Name = "C"
#Data = "<C>"
#X = 29
#Y = 9995
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

[[Field]]	
Name = "X"
Data = "<X>"
X = 33
Y = 9995
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "Q"
Data = "<Q>"
X = 37
Y = 9995
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "L01"
Data = " F1  F2  F3  F4  F5  F6  F7  F8  F10 F12"
X = 0
Y = 9996
FieldLen = 41
Style = "PFKEY"
FieldType = "label"
`
	return taps.NewPanel(doc, styleMatrix, help)
}

// -------------------------------------------------
type Daily struct {
	panel *taps.Panel
}

func (m *Daily) errCheck(common *Common) (string, int) {
	startDate := m.panel.Get("E_RSDATE")
	endDate := m.panel.Get("E_REDATE")
	errMsg := ""
	if !appt.CheckYMD(startDate) {
		errMsg = "D.ER StartDate format error ."
		return errMsg, m.panel.GetFieldNumber("E_RSDATE")
	}
	if !appt.CheckYMD(endDate) {
		errMsg = "D.ER EndDate format error ."
		return errMsg, m.panel.GetFieldNumber("E_REDATE")
	}
	if startDate >= endDate {
		errMsg = "D.ER Duration error ."
		return errMsg, m.panel.GetFieldNumber("E_RSDATE")
	}
	return errMsg, -1
}

func (m *Daily) update(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)
	id := common.apptID
	n := 0
	if id >= 0 {
		u := manager.GetApptByID(id)
		u.RepeatType = appt.GetRepeatType("DAILY")
		n, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_FREQUENCY")))
		u.Frequency = n
		//@@@
		appt.UpdateRepeatSdate(&u, m.panel.Get("E_RSDATE"))
		appt.UpdateRepeatEdate(&u, m.panel.Get("E_REDATE"))
		manager.UpdateAppt(u, id)
	}
	manager.Close()
}

func (m *Daily) doFormat(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)
	id := common.apptID
	u := manager.GetApptByID(id)
	manager.Close()

	if u.Frequency == 0 {
		m.panel.Store(" 1", "E_FREQUENCY")
	} else {
		m.panel.Store(fmt.Sprintf("%2d", u.Frequency), "E_FREQUENCY")
	}

	//@@@@
	if appt.GetRepeatSdate(u) == "" {
		m.panel.Store(common.currentTime.Format(appt.DATE_FORMAT), "E_RSDATE")
		m.panel.Store(common.currentTime.AddDate(5, 0, 0).Format(appt.DATE_FORMAT), "E_REDATE")

	} else {
		m.panel.Store(appt.GetRepeatSdate(u), "E_RSDATE")
		m.panel.Store(appt.GetRepeatEdate(u), "E_REDATE")
	}
}

func (m *Daily) Run(common *Common) string {
	if m.panel == nil {
		m.panel = DailyPanel()
	}
	m.doFormat(common)
	help := &Help{}
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

		if n == "X" || k == tcell.KeyF10 {
			errMsg, num := m.errCheck(common)

			if num >= 0 {
				m.panel.Store(errMsg, "ERR_MSG")
				m.panel.SelectFocus = num
			} else {
				m.update(common)
				m.panel.Store("D.MG updated .", "ERR_MSG")
				//return "Q"
			}
		}

		if n == "Q" || k == tcell.KeyF12 {
			return "Q"
		}
		/*
			if n == "E_FREQUENCY" || k == tcell.KeyUp {
				return "Q"
			}
		*/
	}
	return ""
}
