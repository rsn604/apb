package tui

import (
	"apb/appt"

	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strconv"
	"strings"
)

func CustomPanel() *taps.Panel {
	var help = `Repeat Custom

Configure "Repeat Custom" settings.

(1) Custom Repeat Type
Select either "By Date Number"
 or "By Day Position."
(Press the "ENTER" key to display an "X,"
 or select by mouse.)

-  "By DateNumber"
Enter the execution day(DD).

- "By Day Position"
 Week : Select "Day Position" from "1ST" to "Last".
 Day  : Select Weekday from "Sunday" to "Saturday".
Press "ENTER" key and select one.

(2) Month
Select "Month" from "JAN" to "Dec".

(3) Duration
- Starting : Specify the start date.
- Ending   : Specify the end date.

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
X = 0
Y = 0
FieldLen = 49
Rows = 6
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Custom Repeat Type "
X = 2
Y = 0
Style = "title"
FieldType = "label"

[[Field]]
Name = "S_DAYNUMBER"
Data = "X"
X = 2
Y = 1
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "By Day Number"
X = 4
Y = 1
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_DAYNUMBER"
X = 20
Y = 1
FieldLen = 2
DataLen = 2
Attr = "N"
Style = "edit, edit_focus"
FieldType = "edit"
ExitKey = ["Up"]

# -------------------------------------------------
[[Field]]
Name = "S_DAYPOSITION"
Data = "-"
X = 2
Y = 2
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "By Day Position"
X = 4
Y = 2
Style = "label"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
Data = "Week"
X = 4
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_POS_1ST"
Data = "-"
X = 10
Y = 3
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "1st"
X = 12
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_POS_2ND"
Data = "-"
X = 17
Y = 3
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "2nd"
X = 19
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_POS_3RD"
Data = "-"
X = 24
Y = 3
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "3rd"
X = 26
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_POS_4TH"
Data = "-"
X = 31
Y = 3
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "4th"
X = 33
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_POS_LAST"
Data = "-"
X = 38
Y = 3
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Last"
X = 40
Y = 3
Style = "label"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
Data = "Day"
X = 4
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_MONDAY"
Data = "-"
X = 10
Y = 4
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Mon"
X = 12
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_TUESDAY"
Data = "-"
X = 17
Y = 4
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Tue"
X = 19
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_WEDNESDAY"
Data = "-"
X = 24
Y = 4
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Wed"
X = 26
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_THURSDAY"
Data = "-"
X = 31
Y = 4
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Thu"
X = 33
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_FRIDAY"
Data = "-"
X = 38
Y = 4
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Fri"
X = 40
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_SATURDAY"
Data = "-"
X = 17
Y = 5
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Sat"
X = 19
Y = 5
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_DAY_SUNDAY"
Data = "-"
X = 24
Y = 5
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Sun"
X = 26
Y = 5
Style = "label"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
X = 0
Y = 6
FieldLen = 49
Rows = 3
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Months "
X = 2
Y = 6
Style = "title"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
Name = "S_MONTH_JANUARY"
Data = "X"
X = 3
Y = 7
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Jan"
X = 5
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_FEBRUARY"
Data = "X"
X = 10
Y = 7
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Feb"
X = 12
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_MARCH"
Data = "X"
X = 17
Y = 7
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Mar"
X = 19
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_APRIL"
Data = "X"
X = 24
Y = 7
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Apr"
X = 26
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_MAY"
Data = "X"
X = 31
Y = 7
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "May"
X = 33
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_JUNE"
Data = "X"
X = 38
Y = 7
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Jun"
X = 40
Y = 7
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_JULY"
Data = "X"
X = 3
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Jul"
X = 5
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_AUGUST"
Data = "X"
X = 10
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Aug"
X = 12
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_SEPTEMBER"
Data = "X"
X = 17
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Sep"
X = 19
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_OCTOBER"
Data = "X"
X = 24
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Oct"
X = 26
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_NOVEMBER"
Data = "X"
X = 31
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Nov"
X = 33
Y = 8
Style = "label"
FieldType = "label"

[[Field]]	
Name = "S_MONTH_DECEMBER"
Data = "X"
X = 38
Y = 8
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Dec"
X = 40
Y = 8
Style = "label"
FieldType = "label"

# -------------------------------------------------
[[Field]]	
X = 0
Y = 9
FieldLen = 30
Rows = 3
Rect=true
Style = "linerect"
FieldType = "label"

[[Field]]	
Data = " Duration "
X = 2
Y = 9
Style = "title"
FieldType = "label"

[[Field]]	
Data = "Starting:"
X = 3
Y = 10
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_RSDATE"
X = 17
Y = 10
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]	
Data = "Ending:"
X = 3
Y = 11
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_REDATE"
X = 17
Y = 11
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
type Custom struct {
	panel *taps.Panel
}

func (m *Custom) errCheck(common *Common) (string, int) {
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

func (m *Custom) getDayPosition() int {
	x := 0
	for _, f := range m.panel.Field {
		if strings.HasPrefix(f.Name, "S_POS_") {
			if m.panel.Get(f.Name) == "X" {
				x = x | appt.PositionTBL[f.Name[len("S_POS_"):]]
			}
		}
	}
	return x
}

func (m *Custom) getDayOfWeek() int {
	x := 0
	for _, f := range m.panel.Field {
		if strings.HasPrefix(f.Name, "S_DAY_") {
			if m.panel.Get(f.Name) == "X" {
				x = x | appt.WeekTBL[f.Name[len("S_DAY_"):]]
			}
		}
	}
	return x
}

func (m *Custom) GetMonth() int {
	x := 0
	for _, f := range m.panel.Field {
		if strings.HasPrefix(f.Name, "S_MONTH_") {
			if m.panel.Get(f.Name) == "X" {
				x = x | appt.MonthTBL[f.Name[len("S_MONTH_"):]]
			}
		}
	}
	return x
}

func (m *Custom) update(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)

	id := common.apptID
	n := 0

	if id >= 0 {
		u := manager.GetApptByID(id)
		u.RepeatType = appt.GetRepeatType("CUSTOM")

		if m.panel.Get("S_DAYNUMBER") == "X" {
			n, _ = strconv.Atoi(strings.TrimSpace(m.panel.Get("E_DAYNUMBER")))
			u.DayNumber = n

		} else {
			u.DayNumber = 0
		}

		if m.panel.Get("S_DAYPOSITION") == "X" {
			u.DayPosition = m.getDayPosition()
			u.DayOfWeek = m.getDayOfWeek()
		} else {
			u.DayPosition = 0
			u.DayOfWeek = 0
		}

		u.Month = m.GetMonth()
		//@@@
		appt.UpdateRepeatSdate(&u, m.panel.Get("E_RSDATE"))
		appt.UpdateRepeatEdate(&u, m.panel.Get("E_REDATE"))

		manager.UpdateAppt(u, id)
	}
	manager.Close()
}

func (m *Custom) markDayPosition(x int) {
	for k, v := range appt.PositionTBL {
		if (v & x) != 0 {
			m.panel.Store("X", "S_POS_"+k)
		} else {
			m.panel.Store("-", "S_POS_"+k)
		}
	}
}

func (m *Custom) markDayOfWeek(x int) {
	for k, v := range appt.WeekTBL {
		if (v & x) != 0 {
			m.panel.Store("X", "S_DAY_"+k)
		} else {
			m.panel.Store("-", "S_DAY_"+k)
		}
	}
}

func (m *Custom) markMonth(x int) {
	for k, v := range appt.MonthTBL {
		if (v & x) != 0 {
			m.panel.Store("X", "S_MONTH_"+k)
		} else {
			m.panel.Store("-", "S_MONTH_"+k)
		}
	}
}

func (m *Custom) enableDayPosition(enabled bool) {
	for _, f := range m.panel.Field {
		if strings.HasPrefix(f.Name, "S_POS_") || strings.HasPrefix(f.Name, "S_DAY_") {
			if enabled {
				f.Enabled()
			} else {
				f.Disabled()
			}
		}
	}
}

func (m *Custom) doFormat(common *Common) {
	t := common.getCurrentTime()

	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)
	id := common.apptID
	u := manager.GetApptByID(id)
	manager.Close()

	if u.DayNumber > 0 {
		m.panel.Store("X", "S_DAYNUMBER")
		m.panel.Store("-", "S_DAYPOSITION")
		m.panel.Store(fmt.Sprintf("%02d", u.DayNumber), "E_DAYNUMBER")
		m.enableDayPosition(false)
	} else {
		m.panel.Store(fmt.Sprintf("%02d", t.Day()), "E_DAYNUMBER")
	}
	if u.DayPosition > 0 {
		m.panel.Store("X", "S_DAYPOSITION")
		m.panel.Store("-", "S_DAYNUMBER")
		m.enableDayPosition(true)
		m.markDayPosition(u.DayPosition)
		m.markDayOfWeek(u.DayOfWeek)
	} else {
		dayOfWeek := 1 << int(t.Weekday())
		m.panel.Store("X", "S_DAY_"+appt.GetWeekString(dayOfWeek))

		weekNumber := 1 << (appt.GetWeekNumber(t) - 1)

		//m.panel.Store("X", "S_POS_"+appt.GetPositionString(appt.GetWeekNumber(t)))
		m.panel.Store("X", "S_POS_"+appt.GetPositionString(weekNumber))

	}

	if u.Month > 0 {
		m.markMonth(u.Month)
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

func (m *Custom) Run(common *Common) string {
	if m.panel == nil {
		m.panel = CustomPanel()
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

		if n == "S_DAYNUMBER" {
			m.panel.Store("X", "S_DAYNUMBER")
			m.panel.Store("-", "S_DAYPOSITION")
			m.enableDayPosition(false)
			continue
		}
		if n == "S_DAYPOSITION" {
			m.panel.Store("-", "S_DAYNUMBER")
			m.panel.Store("X", "S_DAYPOSITION")
			m.enableDayPosition(true)
			continue
		}

		if k == tcell.KeyEnter && len(n) > 2 {
			if n[:2] == "S_" {
				if m.panel.Get(n) == "X" {
					m.panel.Store("-", n)
				} else {
					m.panel.Store("X", n)
				}
			}
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
			if n == "E_DAYNUMBER" || k == tcell.KeyUp {
				return "Q"
			}
		*/
	}
	return ""
}
