package tui

import (
	"time"
)

type Common struct {
	currentTime time.Time
	//selectedData  string
	startTime     string
	apptID        int
	databaseName  string
	connectString string
	cols          int
	rows          int
}

func NewCommon() *Common {
	common := &Common{}
	common.apptID = -1
	return common
}

func (self *Common) setCurrentTime() {
	self.currentTime = time.Now()
}

func (self *Common) getCurrentTime() time.Time {
	return self.currentTime
}

func (self *Common) getApptID() int {
	return self.apptID
}
