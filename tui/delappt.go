package tui

import (
	"apb/appt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
)

// ---------------------------------------
// Panel
// ---------------------------------------
func DelApptPanel() *taps.Panel {

	var styleMatrix = [][]string{
		{"label", "aqua", "default"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"edit", "white, underline", "black"},
		{"edit_focus", "yellow", "black"},
	}

	var doc = `
StartX = 10
StartY = 5
EndX = 45
EndY = 13
Rect = true

[[Field]]
Name = "L_CONFIRM"
X = 4
Y = 0
FieldLen = 20
Style = "select"
FieldType = "label"

[[Field]]
Name = "S_ONE"
Data = "X"
X = 2
Y = 2
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "One Occurence"
X = 4
Y = 2
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_SDATE"
X = 20
Y = 2
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]
Name = "S_ALL"
Data = "-"
X = 2
Y = 3
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "All Occurences in the Range"
X = 4
Y = 3
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_RSDATE"
X = 4
Y = 4
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]	
Data = "to"
X = 16
Y = 4
Style = "label"
FieldType = "label"

[[Field]]	
Name = "E_REDATE"
X = 20
Y = 4
FieldLen = 10
Style = "edit, edit_focus"
FieldType = "edit"

[[Field]]
Name = "S_FULL"
Data = "-"
X = 2
Y = 5
FieldLen = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Data = "Fully Delete Occurences"
X = 4
Y = 5
Style = "label"
FieldType = "label"


[[Field]]
Name = "OK"
Data = "OK"
X = 6
Y = 7
Style = "select, select_focus"
FieldType = "select"

[[Field]]
Name = "Cancel"
Data = "Cancel"
X = 15
Y = 7
Style = "select, select_focus"
FieldType = "select"
`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type DelAppt struct {
	panel *taps.Panel
}

func (m *DelAppt) errCheck(common *Common) (string, int) {
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

func (m *DelAppt) delAppt(common *Common) {
	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)
	u := manager.GetApptByID(common.apptID)
	udel := appt.DeletedEntry{}
	if m.panel.Get("S_ONE") == "X" {
		//@@@
		appt.UpdateDeletedRepeatSdate(&udel, m.panel.Get("E_SDATE"))
		appt.UpdateDeletedRepeatEdate(&udel, m.panel.Get("E_SDATE"))
		u.DeletedEntry = append(u.DeletedEntry, udel)
		manager.UpdateAppt(u, common.apptID)
	}

	if m.panel.Get("S_ALL") == "X" {
		appt.UpdateDeletedRepeatSdate(&udel, m.panel.Get("E_RSDATE"))
		appt.UpdateDeletedRepeatEdate(&udel, m.panel.Get("E_REDATE"))
		u.DeletedEntry = append(u.DeletedEntry, udel)
		manager.UpdateAppt(u, common.apptID)
	}

	if m.panel.Get("S_FULL") == "X" {
		manager.DeleteAppt(common.apptID)
		common.apptID = -1
	}
	manager.Close()
}

func (m *DelAppt) doFormat(common *Common, msg string) {
	t := common.getCurrentTime()

	manager := appt.GetManager(common.databaseName)
	manager.Connect(common.databaseName, common.connectString)
	u := manager.GetApptByID(common.apptID)
	manager.Close()

	if u.RepeatType == appt.GetRepeatType("NONE") {
		m.panel.GetDataField("S_ONE").Disabled()
		m.panel.GetDataField("S_ALL").Disabled()
		m.panel.GetDataField("E_SDATE").Disabled()
		m.panel.GetDataField("E_RSDATE").Disabled()
		m.panel.GetDataField("E_REDATE").Disabled()
		m.panel.Store("-", "S_ONE")
		m.panel.Store("-", "S_ALL")
		m.panel.Store("X", "S_FULL")
	} else {
		m.panel.GetDataField("S_ONE").Enabled()
		m.panel.GetDataField("S_ALL").Enabled()
		m.panel.GetDataField("E_SDATE").Enabled()
		m.panel.GetDataField("E_RSDATE").Enabled()
		m.panel.GetDataField("E_REDATE").Enabled()
	}

	m.panel.Store(t.Format(appt.DATE_FORMAT), "E_SDATE")
	//@@@
	m.panel.Store(appt.GetRepeatSdate(u), "E_RSDATE")
	m.panel.Store(appt.GetRepeatEdate(u), "E_REDATE")
	m.panel.Store(msg, "L_CONFIRM")
}

func (m *DelAppt) Run(common *Common, msg string) string {
	if m.panel == nil {
		m.panel = DelApptPanel()
	}

	for {
		m.doFormat(common, "  "+msg)
		m.panel.Say()
		k, n := m.panel.Read()
		if k == tcell.KeyEscape || k == tcell.KeyF12 {
			break
		}

		if n == "S_ONE" {
			m.panel.Store("X", "S_ONE")
			m.panel.Store("-", "S_ALL")
			m.panel.Store("-", "S_FULL")
			continue
		}

		if n == "S_ALL" {
			m.panel.Store("-", "S_ONE")
			m.panel.Store("X", "S_ALL")
			m.panel.Store("-", "S_FULL")
			continue
		}

		if n == "S_FULL" {
			m.panel.Store("-", "S_ONE")
			m.panel.Store("-", "S_ALL")
			m.panel.Store("X", "S_FULL")
			continue
		}

		if k == tcell.KeyEnter {
			if n == "OK" {
				if m.panel.Get("S_FULL") != "X" {
					errMsg, num := m.errCheck(common)

					if num > NO_ERROR {
						m.panel.Store(errMsg, "ERR_MSG")
						m.panel.SelectFocus = num
						continue
					}
				}

				m.delAppt(common)
				//m.panel.Store("D.MG deleted .", "ERR_MSG")
				return "OK"
			}
			if n == "Cancel" {
				return ""
			}
		}
	}

	return ""
}
