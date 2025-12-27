package tui

import (
	"fmt"
	"time"

	"apb/appt"
	"github.com/gdamore/tcell/v2"
	"github.com/rsn604/taps"
	"strings"
)

func MlistPanel() *taps.Panel {
	var help = `Monthly List

Display Month's Appts.

(1) View Appt
Change Month(next) : CTRL-N
            (prev) : CTRL-P
Change Day  (next) : Right Arrow
            (prev) : Left Arrow
Return to Today    : CTRL-T

(2) Change Appt
Add New Appt       : F2 -> Go to "ApptList"
                   : Select Day
                     -> Go to "Apptlist"
Find Appt          : "<F>" F4
View Calendar      : "<G>" F5
View 6-Month       : "<6>" F6
Weekly Appt        : F8
                     -> Go to "Weekly List"

"ESC" "<Q>" F12 -> Quit.

`
	var styleMatrix = [][]string{
		{"PFKEY", "white", "default"},
		{"ULINE", "white, underline", "defaut"},
		{"NOLINE", "white", "defaut"},
		{"L_YM", "lightcyan, underline", "defaut"},
		{"label", "aqua", "default"},
		{"S_DAY", "white, underline", "default"},
		{"S_DAY_FOCUS", "black", "yellow"},
		{"select", "yellow", "default"},
		{"select_focus", "black", "yellow"},
		{"list", "white", "default"},
		{"list_focus", "white", "aqua"},
	}

	var doc = `
StartX = 0
StartY = 0
EndX = 9999
EndY = 9999
ExitKey = ["F1", "F4", "F5", "F6", "F10", "F12", "Ctrl-T", "Ctrl-P", "Ctrl-N"]

[[Field]]	
Data = " Sunday   Monday  Tuesday  Wednesday Tursday  Friday  Saturday"
X = 0
Y = 0
Style = "label"
FieldType = "label"

# -------------------------------------------------
[[Field]]
GridFields =[
{Name="L_APPT_LIST", X=0, Y=1, FieldLen=8, Rows=2, Style="list", FieldType="label"},
{Name="L_ULINE", X=0, Y=3, FieldLen=5, Style="ULINE", FieldType="label"},
{Name="S_DAY", X=5, Y=3, FieldLen=3, Style="S_DAY, S_DAY_FOCUS", FieldType="select", ExitKey = ["F2", "F8"]},
{Name="L_VERLINE", X=8, Y=1, FieldLen=1, Rows=3, Style="list", FieldType="label"},
]
Cols = 7
FieldLen = 9
Rows = 6

# -------------------------------------------------
[[Field]]
Name = "H"
Data = "<H>"
X = 1
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]
Name = "A"
Data = "Add"
X = 5
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "label"

[[Field]]	
Name = "F"
Data = "<F>"
X = 13
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "G"
Data = "<G>"
X = 17
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "6"
Data = "<6>"
X = 21
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "W"
Data = "WL"
X = 29
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "label"

#[[Field]]	
#Name = "T"
#Data = "<T>"
#X = 33
#Y = 9998
#FieldLen = 4
#Style = "select, select_focus"
#FieldType = "select"

[[Field]]	
Name = "Q"
Data = "<Q>"
X = 37
Y = 9998
FieldLen = 4
Style = "select, select_focus"
FieldType = "select"

[[Field]]	
Name = "L_YM"
X = 43
Y = 9998
FieldLen = 7
Style = "L_YM"
FieldType = "label"

[[Field]]	
Name = "L01"
Data = " F1  F2  F3  F4  F5  F6  F7  F8  F10 F12"
X = 0
Y = 9999
FieldLen = 41
Style = "PFKEY"
FieldType = "label"
`
	return taps.NewPanel(doc, styleMatrix, help)
}

// -------------------------------------------------
type Mlist struct {
	panel *taps.Panel
}

// ---------------------------------------
func (m *Mlist) getApptDescription(appts []appt.Appointment) []string {
	var listData []string
	for _, ap := range appts {
		listData = append(listData, ap.Description)
	}
	return listData
}

func (m *Mlist) doFormat(common *Common) {
	t := common.getCurrentTime()
	tf := time.Date(t.Year(), t.Month(), 1, 0, 0, 0, 0, time.Local)
	days := tf.AddDate(0, 1, -1).Day()
	startCol := int(tf.Weekday())

	m.panel.ClearGridData("S_DAY")
	m.panel.ClearGridList("L_VERLINE")
	m.panel.ClearGridList("L_APPT_LIST")

	m.panel.ResetFieldStyle("L_ULINE", "NOLINE")
	m.panel.ResetFieldStyle("S_DAY", "NOLINE")

	col := 0
	row := 0
	vlines := []string{"|", "|", "|", "|"}

	tf = tf.AddDate(0, 0, -startCol)
	i := 0

	for {
		//if i > days+startCol && col == C_DAYS-1{
		if i >= days+startCol && col == 0 {
			break
		}
		appts := appt.GetApptToday(common.databaseName, common.connectString, tf)
		m.panel.StoreGridData(fmt.Sprintf("%02d", int(tf.Day())), "S_DAY", col, row)
		m.panel.StoreGridList(vlines, "L_VERLINE", col, row)
		m.panel.StoreGridList(m.getApptDescription(appts), "L_APPT_LIST", col, row)
		if t.Format(appt.DATE_FORMAT) == tf.Format(appt.DATE_FORMAT) {
			m.panel.SelectFocus = m.panel.GetGridFieldNumber("S_DAY", col, row)
		}

		m.panel.ResetFieldStyle(m.panel.GetGridFieldName("L_ULINE", col, row), "ULINE")
		m.panel.ResetFieldStyle(m.panel.GetGridFieldName("S_DAY", col, row), "S_DAY, S_DAY_FOCUS")

		col++
		if col == C_DAYS {
			col = 0
			row++
		}
		tf = tf.AddDate(0, 0, 1)
		i++
	}

	m.panel.Store(t.Format(appt.YYYYMM_FORMAT), "L_YM")
	m.panel.Say()
}

func (m *Mlist) Run(common *Common) string {

	if m.panel == nil {
		m.panel = MlistPanel()
	}
	calendar := &Calendar{}
	godate := &GoDate{}
	wlist := &Wlist{}
	find := &Find{}
	help := &Help{}

	for {
		m.doFormat(common)
		k, n := m.panel.Read()
		if k == tcell.KeyEscape {
			break
		}

		if k == tcell.KeyCtrlT {
			common.setCurrentTime()
			continue
		}

		if k == tcell.KeyCtrlP {
			common.currentTime = prevMonth(common.currentTime) //godate.go
			continue
		}

		if k == tcell.KeyCtrlN {
			common.currentTime = nextMonth(common.currentTime) //godate.go
			continue
		}

		if n == "H" || k == tcell.KeyF1 {
			help.Run(m.panel.GetHelp())
			continue
		}

		if (k == tcell.KeyF2 || k == tcell.KeyEnter) && strings.HasPrefix(n, "S_DAY") {
			ymd := common.currentTime.Format(appt.YYYYMM_FORMAT) + "/" + strings.TrimSpace(m.panel.Get(n))
			if appt.CheckYMD(ymd) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, ymd)
				break
			}
			continue
		}

		if n == "F" || k == tcell.KeyF4 {
			rs := find.Run(common, "Find Appointment")
			if len(rs) > 0 {
				kk := strings.Split(rs, " ")
				//@@@@@
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, kk[0])
			}

			continue
		}

		if n == "G" || k == tcell.KeyF5 {
			rs := godate.Run(common)
			if appt.CheckYMD(rs) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, rs)
			}
			continue
		}

		if n == "6" || k == tcell.KeyF6 {
			rs := calendar.Run(common)
			if appt.CheckYMD(rs) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, rs)
			}
			continue
		}

		if (k == tcell.KeyF8 || k == tcell.KeyEnter) && strings.HasPrefix(n, "S_DAY") {
			ymd := common.currentTime.Format(appt.YYYYMM_FORMAT) + "/" + strings.TrimSpace(m.panel.Get(n))
			if appt.CheckYMD(ymd) {
				common.currentTime, _ = time.Parse(appt.DATE_FORMAT, ymd)

				rs := wlist.Run(common)
				if rs == "Q" {
					break
				}
			}
			continue
		}

		if n == "Q" || k == tcell.KeyF12 {
			//break
			return "Q"
		}

	}
	return ""
}
