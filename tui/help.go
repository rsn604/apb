package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strings"
)

// ---------------------------------------
// Panel
// ---------------------------------------
func HelpPanel() *taps.Panel {
	var styleMatrix = [][]string{
		{"edit", "white", "black"},
		{"edit_focus", "yellow", "black"},
	}

	var doc = `
StartX = 1
StartY = 2
EndX = 49
EndY = 20
Rect = true
ExitKey = ["F12"]

[[Field]]
Name = "E_HELP"
X = 2
Y = 1
Rows = 17
Style = "edit, edit_focus"
FieldType = "edit"
`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type Help struct {
	panel *taps.Panel
}

func (m *Help) Run(msg string) string {
	if m.panel == nil {
		m.panel = HelpPanel()
	}
	m.panel.StoreList(strings.Split(msg, "\n"), "E_HELP")
	m.panel.SetBrowseMode("E_HELP", true)
	m.panel.Say()
	for {
		k, _ := m.panel.Read()
		if k == tcell.KeyEscape || k == tcell.KeyF12 {
			break
		}
	}
	return "OK"
}
