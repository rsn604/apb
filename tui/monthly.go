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

func MonthlyPanel() *taps.Panel {
	var help = `Repeat Monthly

Configure "Repeat Monthly" settings.

(1) Frequency
- Repeat Every x Month(s):
  Specify the number of months between each interval.
  To run every three Months, set this to "3."

(2) Duration
- Starting : Specify the start date.
- Ending   : Specify the end date.

(3) Monthly Repeat Type
Select either "By Day Number" or "By Day Position."
(Press the "ENTER" key to display an "X,"or select
by mouse.)

-  "By Day Number"
Enter the execution day(DD).

- "By Day Position"
Specify the week (from "1st" to "last") and day of
 the week (from "Sunday" to "Saturday") for execution.
Press "ENTER" key and select one.

(4) Key Operation
Enter "Repeat Daily" Settings : "<X>" F10

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
Data = "Month(s)"
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
X = 1
Y = 7
FieldLen = 49
Rows = 3
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Monthly Repeat Type "
X = 3
Y = 7
Style = "title"
FieldType = "label"

[[Field]]
Name = "S_DAYNUMBER"
Data = "X"
X = 4
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "By Day Number"
X = 7
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_DAYNUMBER"
X = 24
Y = 8
FieldLen = 2
DataLen = 2
Attr = "N"
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]
Name = "S_DAYPOSITION"
Data = "-"
X = 4
Y = 9
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "By Day Position"
X = 7
Y = 9
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_POSITION"
X = 24
Y = 9
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "S_DAYOFWEEK"
X = 29
Y = 9
FieldLen = 9
Style = "select, select_focus"
FieldType = "select"

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
type Monthly struct {
	panel *taps.Panel
}

func (m *Monthly) errCheck(common *Common) (string, int) {
	startDate := m.panel.Get("E_RSDATE")
	endDate := m.panel.Get("E_REDATE")
	errMsg := ""
	if !appt.CheckYMD(startDate) {
		errMsg = "D.ER Start Date format error ."
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
	return errMsg, NO_ERROR
}

func (m *Monthly) update(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)

	id := common.apptID
	if id >= 0 {
		u := manager.GetApptByID(id)
		u.RepeatType = appt.GetRepeatType("MONTHLY")

		u.Frequency, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_FREQUENCY")))
		//@@@
		appt.UpdateRepeatSdate(&u, m.panel.Get("E_RSDATE"))
		appt.UpdateRepeatEdate(&u, m.panel.Get("E_REDATE"))

		if m.panel.Get("S_DAYNUMBER") == "X" {
			u.DayNumber, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_DAYNUMBER")))

		} else {
			u.DayNumber = 0
		}
		if m.panel.Get("S_DAYPOSITION") == "X" {
			u.DayPosition = appt.GetPosition(m.panel.Get("S_POSITION"))
			u.DayOfWeek = appt.GetWeek(m.panel.Get("S_DAYOFWEEK"))
		} else {
			u.DayPosition = 0
			u.DayOfWeek = 0
		}

		manager.UpdateAppt(u, id)
	}
	manager.Close()

}

func (m *Monthly) doFormat(common *Common) {
	t := common.getCurrentTime()

	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)

	//id := appt.GetApptID(common.selectedData)
	id := common.apptID
	u := manager.GetApptByID(id)
	manager.Close()

	if u.Frequency == 0 {
		m.panel.Store(" 1", "E_FREQUENCY")
	} else {
		m.panel.Store(fmt.Sprintf("%2d", u.Frequency), "E_FREQUENCY")
	}

	if u.DayNumber > 0 {
		m.panel.Store("X", "S_DAYNUMBER")
		m.panel.Store("-", "S_DAYPOSITION")
		m.panel.Store(fmt.Sprintf("%2d", u.DayNumber), "E_DAYNUMBER")
		m.panel.GetDataField("E_DAYNUMBER").Enabled()
		m.panel.GetDataField("S_POSITION").Disabled()
		m.panel.GetDataField("S_DAYOFWEEK").Disabled()
	} else {
		m.panel.Store(fmt.Sprintf("%02d", t.Day()), "E_DAYNUMBER")
	}

	if u.DayPosition > 0 {
		m.panel.Store("X", "S_DAYPOSITION")
		m.panel.Store("-", "S_DAYNUMBER")
		m.panel.GetDataField("E_DAYNUMBER").Disabled()
		m.panel.GetDataField("S_POSITION").Enabled()
		m.panel.GetDataField("S_DAYOFWEEK").Enabled()
		if u.DayPosition > 0 {
			m.panel.Store(appt.GetPositionString(u.DayPosition), "S_POSITION")
		}
		if u.DayOfWeek > 0 {
			m.panel.Store(appt.GetWeekString(u.DayOfWeek), "S_DAYOFWEEK")
		}
	} else {
		dayOfWeek := 1 << int(t.Weekday())
		m.panel.Store(appt.GetWeekString(dayOfWeek), "S_DAYOFWEEK")

		weekNumber := 1 << (appt.GetWeekNumber(t) - 1)
		m.panel.Store(appt.GetPositionString(weekNumber), "S_POSITION")
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

func (m *Monthly) Run(common *Common) string {
	if m.panel == nil {
		m.panel = MonthlyPanel()
	}
	m.doFormat(common)
	selector := &Selector{}
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

		if n == "S_POSITION" {
			selector.panel = taps.ModifyPanelPosition(SelectorPanel(), 28, 11)
			rc := selector.Run(common, appt.SortMap(appt.PositionTBL))
			if rc != "" {
				m.panel.Store(rc, "S_POSITION")
			}
		}

		if n == "S_DAYOFWEEK" {
			selector.panel = taps.ModifyPanelPosition(SelectorPanel(), 31, 11)
			rc := selector.Run(common, appt.SortMap(appt.WeekTBL))
			if rc != "" {
				m.panel.Store(rc, "S_DAYOFWEEK")
			}
		}

		if n == "S_DAYNUMBER" {
			m.panel.Store("X", "S_DAYNUMBER")
			m.panel.Store("-", "S_DAYPOSITION")
			m.panel.GetDataField("E_DAYNUMBER").Enabled()
			m.panel.GetDataField("S_POSITION").Disabled()
			m.panel.GetDataField("S_DAYOFWEEK").Disabled()
		}

		if n == "S_DAYPOSITION" {
			m.panel.Store("-", "S_DAYNUMBER")
			m.panel.Store("X", "S_DAYPOSITION")
			m.panel.GetDataField("E_DAYNUMBER").Disabled()
			m.panel.GetDataField("S_POSITION").Enabled()
			m.panel.GetDataField("S_DAYOFWEEK").Enabled()
		}

		if n == "X" || k == tcell.KeyF10 {
			errMsg, num := m.errCheck(common)

			if num > NO_ERROR {
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
