package tui

import (
	"fmt"
	"time"

	"apb/appt"
	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strings"
)

func GoDatePanel() *taps.Panel {

	var styleMatrix = [][]string{
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
		{"edit", "white, underline", "black"},
		{"edit_focus", "yellow", "black"},
	}

	var doc = `
StartX = 9961
StartY = 1
EndX = 9999
EndY = 14
Rect = true
ExitKey = ["F2", "F3", "F4", "F5", "F6", "F7", "F8", "F10", "F12"]

# -------------------------------------------------
[[Field]]	
Name = "CAL_YYYYMM"
X = 8
Y = 1
FieldLen = 18
Style = "CAL_YYYYMM"
FieldType = "label"

[[Field]]	
X = 6
Y = 2
Data = "Su  Mo  Tu  We  Th  Fr  Sa"
Style = "label"
FieldType = "label"

[[Field]]
Name = "CAL"
X = 6
Y = 3
FieldLen = 2
Cols = 7
Rows = 6
ColSpaces = 2
Style = "CAL, CAL_FOCUS"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Data = "Goto "
X = 2
Y = 9
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_YMD"
X = 8
Y = 9
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]	
Name = "G"
Data = "<Go>"
X = 25
Y = 9
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

# -------------------------------------------------
[[Field]]	
Name = "d"
Data = "<d-"
X = 2
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "D"
Data = "<D+"
X = 6
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "T"
Data = "<T>"
X = 10
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "m"
Data = "<m-"
X = 14
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "M"
Data = "<M+"
X = 18
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "y"
Data = "<y-"
X = 22
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "Y"
Data = "<Y+"
X = 26
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "Q"
Data = "<Q>"
X = 34
Y = 11
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "L01"
Data = "F2  F3  F4  F5  F6  F7  F8  F10 F12"
X = 2
Y = 12
#FieldLen = 30
Style = "PFKEY"
FieldType = "label"

`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type GoDate struct {
	panel *taps.Panel
}

func prevMonth(t time.Time) time.Time {
	year, month, day := t.Date()
	_, _, lastDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1).Date()
	if day > lastDay {
		return time.Date(year, month-1, lastDay, 0, 0, 0, 0, time.Local)
	}
	return t.AddDate(0, -1, 0)
}

func nextMonth(t time.Time) time.Time {
	year, month, day := t.Date()
	_, _, lastDay := time.Date(year, month+2, 1, 0, 0, 0, 0, time.Local).AddDate(0, 0, -1).Date()
	if day > lastDay {
		return time.Date(year, month+1, lastDay, 0, 0, 0, 0, time.Local)
	}
	return t.AddDate(0, 1, 0)
}

func (m *GoDate) doFormat(t time.Time) {
	listData, today := GetMonthCalendar(t)
	m.panel.StoreList(listData, "CAL")
	m.panel.Store(fmt.Sprintf("%04d/%02d  %s", t.Year(), int(t.Month()), t.Month()), "CAL_YYYYMM")
	m.panel.Store(t.Format(appt.DATE_FORMAT), "E_YMD")

	m.panel.SelectFocus = setDayFocus(m.panel, today) // apptlist.go
	m.panel.Say()
}

func (m *GoDate) Run(common *Common) string {
	if m.panel == nil {
		m.panel = GoDatePanel()
	}
	for {
		m.doFormat(common.currentTime)

		k, n := m.panel.Read()
		if k == tcell.KeyEscape {
			break
		}

		if n == "d" || k == tcell.KeyF2 {
			common.currentTime = common.currentTime.AddDate(0, 0, -1)
			continue
		}
		if n == "D" || k == tcell.KeyF3 {
			common.currentTime = common.currentTime.AddDate(0, 0, 1)
			continue
		}

		if n == "T" || k == tcell.KeyF4 {
			common.setCurrentTime()
			continue
		}

		if n == "m" || k == tcell.KeyF5 {
			//common.currentTime = common.currentTime.AddDate(0, -1, 0)
			common.currentTime = prevMonth(common.currentTime)
			continue
		}
		if n == "M" || k == tcell.KeyF6 {
			//common.currentTime = common.currentTime.AddDate(0, 1, 0)
			common.currentTime = nextMonth(common.currentTime)
			continue
		}

		if n == "y" || k == tcell.KeyF7 {
			common.currentTime = common.currentTime.AddDate(-1, 0, 0)
			continue
		}
		if n == "Y" || k == tcell.KeyF8 {
			common.currentTime = common.currentTime.AddDate(1, 0, 0)
			continue
		}

		if k == tcell.KeyEnter && len(n) > 3 && n[:3] == "CAL" {
			dd := strings.TrimSpace(m.panel.Get(n))
			if len(dd) < 2 {
				dd = "0" + dd
			}
			return fmt.Sprintf("%04d/%02d/%s", common.currentTime.Year(), int(common.currentTime.Month()), dd)
		}

		if n == "G" {
			goDate := m.panel.Get("E_YMD")
			if appt.CheckYMD(goDate) {
				return goDate
			}
		}

		if n == "Q" || k == tcell.KeyF12 {
			return "Q"
		}

	}
	return ""
}
