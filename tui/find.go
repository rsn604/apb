package tui

import (
	"apb/appt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"sort"
	"strings"
)

// ---------------------------------------
// Panel
// ---------------------------------------
func FindPanel() *taps.Panel {

	var styleMatrix = [][]string{
		{"errmsg", "red", "default"},
		{"label", "aqua", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"edit", "white, underline", "black"},
		{"edit_focus", "yellow", "black"},
	}

	var doc = `
StartX = 10
StartY = 4
EndX = 45
EndY = 11
Rect = true

[[Field]]
Name = "L_CONFIRM"
X = 4
Y = 0
FieldLen = 20
Style = "select"
FieldType = "label"

[[Field]]	
Data = "Find"
X = 2
Y = 2
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_FIND"
X = 7
Y = 2
FieldLen = 25
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]
Name = "S_NEXT"
Data = "Next"
X = 4
Y = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]
Name = "S_PREV"
Data = "Previous"
X = 12
Y = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]
Name = "S_CANCEL"
Data = "Cancel"
X = 25
Y = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]
Name = "ERR_MSG"
X = 3
Y = 6
FieldLen = 30
Style = "errmsg"
FieldType = "label"

`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type Find struct {
	panel *taps.Panel
}

func isContains(field string, findStr string) bool {
	if findStr == "" {
		//return false
		return true
	}
	return strings.Contains(strings.ToLower(field), strings.ToLower(findStr))
}

func searchStr(findStr string, ap appt.Appointment) bool {
	if isContains(ap.Description, findStr) {
		return true
	}
	//@@@
	if isContains(appt.GetSdate(ap), findStr) {
		return true
	}
	if isContains(ap.Stime, findStr) {
		return true
	}
	if isContains(ap.Etime, findStr) {
		return true
	}
	if isContains(ap.Location, findStr) {
		return true
	}
	for _, n := range ap.Note {
		if isContains(n, findStr) {
			return true
		}
	}
	return false
}

func findAllAppts(common *Common, findStr, direction string, state int) ([]string, map[string]appt.Appointment) {
	apptMap := map[string]appt.Appointment{}
	if state == appt.STATE_APPT {
		if direction == C_NEXT {
			apptMap = appt.GetNextAppts(common.databaseName, common.connectString, common.getCurrentTime())

		} else {
			apptMap = appt.GetPreviousAppts(common.databaseName, common.connectString, common.getCurrentTime())

		}
	} else {
		if direction == C_NEXT {
			apptMap = appt.GetNextTodos(common.databaseName, common.connectString, common.getCurrentTime())

		} else {
			apptMap = appt.GetPreviousTodos(common.databaseName, common.connectString, common.getCurrentTime())

		}

	}

	keys := []string{}
	for k := range apptMap {
		keys = append(keys, k)
	}
	if direction == C_NEXT {
		sort.Strings(keys)
	} else {
		sort.Sort(sort.Reverse(sort.StringSlice(keys)))
	}
	return keys, apptMap
}

func findAppt(common *Common, findStr, direction string, state int) string {
	keys, apptMap := findAllAppts(common, findStr, direction, state)
	for _, k := range keys {
		if searchStr(findStr, apptMap[k]) {
			return k
		}
	}
	return ""
}

func (m *Find) doFormat(common *Common, msg string) {
	m.panel.Store(msg, "L_CONFIRM")
}

func (m *Find) Run(common *Common, msg string) string {
	state := appt.STATE_APPT
	if msg == FIND_TODO {
		state = appt.STATE_TODO
	}
	if m.panel == nil {
		m.panel = FindPanel()
	}
	m.doFormat(common, "  "+msg)
	for {
		m.panel.Say()
		k, n := m.panel.Read()
		if k == tcell.KeyEscape || k == tcell.KeyF12 {
			break
		}

		if k == tcell.KeyEnter {
			if n == "S_CANCEL" {
				return ""
			}

			findStr := m.panel.Get("E_FIND")
			direction := C_NEXT
			if n == "S_PREV" {
				direction = C_PREV
			}

			//if len(findStr) > 0 {
			k := findAppt(common, findStr, direction, state)
			if len(k) > 0 {
				kk := strings.Split(k, " ")
				if appt.CheckYMD(kk[0]) {
					return k
				}
			}
			//}
			m.panel.SelectFocus = m.panel.GetFieldNumber("E_FIND")
			m.panel.Store("D.ER No Record found.", "ERR_MSG")
		}
	}

	return ""
}
