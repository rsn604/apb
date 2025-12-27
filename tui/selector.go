package tui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
)

// ---------------------------------------
// Panel
// ---------------------------------------
func SelectorPanel() *taps.Panel {
	var styleMatrix = [][]string{
		{"list", "white", "default"},
		{"list_focus", "white", "aqua"},
	}

	var doc = `
StartX = 20
StartY = 9
EndX = 35
EndY = 17
Rect = true

[[Field]]	
Name = "SELECTOR"
X = 2
Y = 1
Rows = 7
Style = "list, list_focus"
FieldType = "select"
`
	return taps.NewPanel(doc, styleMatrix, "")
}

// -------------------------------------------------
type Selector struct {
	panel *taps.Panel
}

// ---------------------------------------
// Main
// ---------------------------------------
func (t *Selector) Run(common *Common, selectList []string) string {
	if t.panel == nil {
		t.panel = SelectorPanel()
	}
	t.panel.StoreList(selectList, "SELECTOR")
	t.panel.Say()
	k, n := t.panel.Read()
	if k == tcell.KeyEnter {
		return t.panel.Get(n)
	}
	return ""
}
