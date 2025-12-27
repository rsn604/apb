package tui

import (
	"apb/appt"
	"fmt"
	"github.com/rsn604/taps"
	"os"
)

func App() {
	common := NewCommon()
	common.setCurrentTime()

	if len(os.Args) == 3 {
		common.databaseName = os.Args[1]
		common.connectString = os.Args[2]
	} else if len(os.Args) == 2 {
		common.databaseName = "BOLT"
		common.connectString = os.Args[1]
	} else {
		taps.Quit()
		fmt.Println("Usage: apb <DB name>")
		return
	}

	common.cols, common.rows = taps.GetWindowSize()
	//if common.cols >= 50 {
	if common.cols >= 0 {

		manager := appt.GetManager(common.databaseName)
		err := manager.Connect(common.databaseName, common.connectString)
		if err != nil {
			panic(err)
		}

		err = manager.Define()
		if err != nil {
			panic(err)
		}
		manager.Close()

		m := &ApptList{}
		m.Run(common)

	} else {
		taps.Quit()
		fmt.Printf("Error:Terminal col size < 50. cols:%d rows:%d\n", common.cols, common.rows)
	}

}
