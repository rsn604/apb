package tui

import (
	"apb/appt"

	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
)

// ---------------------------------------
// Panel
// ---------------------------------------
func RepeatPanel() *taps.Panel {

	var styleMatrix = [][]string{
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 49
EndY = 2
Rect = true
ExitKey = ["F12"]

[[Field]]	
Name = "N"
Data = "NoRepeat"
X = 2
Y = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "D"
Data = "Daily"
X = 11
Y = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "W"
Data = "Weekly"
X = 17
Y = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "M"
Data = "Monthly"
X = 24
Y = 1
FieldLen = 8
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "Y"
Data = "Yearly"
X = 32
Y = 1
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "C"
Data = "Custom"
X = 40
Y = 1
Style = "select, select_focus"
FieldType = "select"

`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type RepeatOption struct {
	panel *taps.Panel
}

// ---------------------------------------
// Main
// ---------------------------------------
func (t *RepeatOption) Run(common *Common) string {
	if t.panel == nil {
		t.panel = RepeatPanel()
	}

	daily := &Daily{}
	weekly := &Weekly{}
	monthly := &Monthly{}
	yearly := &Yearly{}
	custom := &Custom{}
	confirm := &Confirm{}

	id := common.apptID
	if id >= 0 {
		manager := appt.GetManager(common.databaseName)
		manager.Connect(common.databaseName, common.connectString)
		u := manager.GetApptByID(id)
		for i := 0; i < 6; i++ {
			if ((u.RepeatType >> i) & 0x01) != 0 {
				t.panel.SelectFocus = i
				break
			}
		}
		manager.Close()
	}

	for {
		taps.Clear()
		t.panel.Say()
		k, n := t.panel.Read()
		if k == tcell.KeyEscape || k == tcell.KeyF12 {
			break
		}

		if n == "N" {
			rs := confirm.Run("Set 'NO_REPEAT' ?")
			if rs == "OK" {
				manager := appt.GetManager(common.databaseName)
				manager.Connect(common.databaseName, common.connectString)
				u := manager.GetApptByID(id)
				u.RepeatType = 0
				manager.UpdateAppt(u, id)
				manager.Close()
				break
			}
		}

		if n == "D" {
			rs := daily.Run(common)
			if rs == "Q" {
				break
			}
		}

		if n == "W" {
			rs := weekly.Run(common)
			if rs == "Q" {
				break
			}
		}

		if n == "M" {
			rs := monthly.Run(common)
			if rs == "Q" {
				break
			}
		}

		if n == "Y" {
			rs := yearly.Run(common)
			if rs == "Q" {
				break
			}
		}

		if n == "C" {
			rs := custom.Run(common)
			if rs == "Q" {
				break
			}
		}

	}
	return "Q"
}
