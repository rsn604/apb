package tui

import (
	"apb/appt"
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strconv"
	"strings"
)

func YearlyPanel() *taps.Panel {
	var help = `Repeat Yearly

Configure "Repeat Yearly" settings.

(1) Frequency
- Repeat Every x Year(s):
  Specify the number of months between each interval.
  To run every three Years, set this to "3."

(2) Duration
- Starting : Specify the start date.
- Ending   : Specify the end date.

(3) Yearly Repeat Type
Select either "By Date"
 or "By Day Position."
(Press the "ENTER" key to display an "X,"
 or select by mouse.)

-  "By Date"
Enter the execution date(MM/DD).

- "By Day Position"
Specify day of  the week
(from "Sunday" to "Saturday") for execution.
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
Data = "Year(s)"
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
FieldLen = 51
Rows = 3
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Yearly Repeat Type "
X = 3
Y = 7
Style = "title"
FieldType = "label"

[[Field]]
Name = "S_DATEOFYEAR"
Data = "X"
X = 4
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "By Date"
X = 7
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_DATEOFYEAR"
X = 24
Y = 8
FieldLen = 5
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

[[Field]]	
Data = "in"
X = 39
Y = 9
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH"
X = 42
Y = 9
FieldLen = 10 
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
type Yearly struct {
	panel *taps.Panel
}

func (m *Yearly) errCheck(common *Common) (string, int) {
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

	_, flag := appt.CheckMMDD(m.panel.Get("E_DATEOFYEAR"))
	if !flag {
		errMsg = "D.ER Date format error ."
		return errMsg, m.panel.GetFieldNumber("E_DATEOFYEAR")
	}

	return errMsg, NO_ERROR
}

func (m *Yearly) update(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)

	id := common.apptID
	n := 0

	if id >= 0 {
		u := manager.GetApptByID(id)
		u.RepeatType = appt.GetRepeatType("YEARLY")

		n, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_FREQUENCY")))
		u.Frequency = n
		//@@@
		appt.UpdateRepeatSdate(&u, m.panel.Get("E_RSDATE"))
		appt.UpdateRepeatEdate(&u, m.panel.Get("E_REDATE"))

		if m.panel.Get("S_DATEOFYEAR") == "X" {
			u.DateOfYear = m.panel.Get("E_DATEOFYEAR")

		} else {
			u.DateOfYear = ""
		}

		if m.panel.Get("S_DAYPOSITION") == "X" {
			u.DayPosition = appt.GetPosition(m.panel.Get("S_POSITION"))
			u.DayOfWeek = appt.GetWeek(m.panel.Get("S_DAYOFWEEK"))
			u.Month = appt.GetMonth(m.panel.Get("S_MONTH"))
		} else {
			u.DayPosition = 0
			u.DayOfWeek = 0
			u.Month = 0
		}

		manager.UpdateAppt(u, id)
	}
	manager.Close()

}

func (m *Yearly) doFormat(common *Common) {
	t := common.getCurrentTime()

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

	if len(u.DateOfYear) > 0 {
		m.panel.Store("X", "S_DATEOFYEAR")
		m.panel.Store("-", "S_DAYPOSITION")
		m.panel.Store(u.DateOfYear, "E_DATEOFYEAR")
		m.panel.GetDataField("E_DATEOFYEAR").Enabled()
		m.panel.GetDataField("S_POSITION").Disabled()
		m.panel.GetDataField("S_DAYOFWEEK").Disabled()
		m.panel.GetDataField("S_MONTH").Disabled()
	} else {
		m.panel.Store(fmt.Sprintf("%02d/%02d", t.Month(), t.Day()), "E_DATEOFYEAR")

	}

	if u.DayPosition > 0 {
		m.panel.Store("-", "S_DATEOFYEAR")
		m.panel.Store("X", "S_DAYPOSITION")
		m.panel.GetDataField("E_DATEOFYEAR").Disabled()
		m.panel.GetDataField("S_POSITION").Enabled()
		m.panel.GetDataField("S_DAYOFWEEK").Enabled()
		m.panel.GetDataField("S_MONTH").Enabled()
		m.panel.Store(appt.GetPositionString(u.DayPosition), "S_POSITION")
		if u.DayOfWeek > 0 {
			m.panel.Store(appt.GetWeekString(u.DayOfWeek), "S_DAYOFWEEK")
		}
		if u.Month > 0 {
			m.panel.Store(appt.GetMonthString(u.Month), "S_MONTH")
		}

	} else {
		dayOfWeek := 1 << int(t.Weekday())
		m.panel.Store(appt.GetWeekString(dayOfWeek), "S_DAYOFWEEK")

		weekNumber := 1 << (appt.GetWeekNumber(t) - 1)
		m.panel.Store(appt.GetPositionString(weekNumber), "S_POSITION")

		month := 1 << (int(t.Month()) - 1)
		m.panel.Store(appt.GetMonthString(month), "S_MONTH")
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

func (m *Yearly) Run(common *Common) string {
	if m.panel == nil {
		m.panel = YearlyPanel()
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

		if n == "S_MONTH" {
			selector.panel = taps.ModifyPanelPosition(SelectorPanel(), 43, 11)
			rc := selector.Run(common, appt.SortMap(appt.MonthTBL))
			if rc != "" {
				m.panel.Store(rc, "S_MONTH")
			}
		}

		if n == "S_DATEOFYEAR" {
			m.panel.Store("X", "S_DATEOFYEAR")
			m.panel.Store("-", "S_DAYPOSITION")
			m.panel.GetDataField("E_DATEOFYEAR").Enabled()
			m.panel.GetDataField("S_POSITION").Disabled()
			m.panel.GetDataField("S_DAYOFWEEK").Disabled()
			m.panel.GetDataField("S_MONTH").Disabled()
		}

		if n == "S_DAYPOSITION" {
			m.panel.Store("-", "S_DATEOFYEAR")
			m.panel.Store("X", "S_DAYPOSITION")
			m.panel.GetDataField("E_DATEOFYEAR").Disabled()
			m.panel.GetDataField("S_POSITION").Enabled()
			m.panel.GetDataField("S_DAYOFWEEK").Enabled()
			m.panel.GetDataField("S_MONTH").Enabled()
		}

		if n == "X" || k == tcell.KeyF10 {
			errMsg, num := m.errCheck(common)

			if num > NO_ERROR {
				m.panel.Store(errMsg, "ERR_MSG")
				m.panel.SelectFocus = num
			} else {
				m.panel.Store(errMsg, "")
				//m.panel.GetDataField("ERR_MSG").Say()
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
