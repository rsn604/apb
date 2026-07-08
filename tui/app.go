package tui

import (
	"apb/appt"
	"fmt"
	"github.com/rsn604/taps"
	//"os"
	"flag"
	"strings"
)

func App() {
	common := NewCommon()
	common.setCurrentTime()

	lang := flag.String("lang", "", "Language")
	flag.Parse()

	common.lang = strings.ToUpper(*lang)

	if flag.NArg() == 2 {
		common.databaseName = flag.Arg(0)
		common.connectString = flag.Arg(1)
	} else if flag.NArg() == 1 {
		common.databaseName = "BOLT"
		common.connectString = flag.Arg(0)
	} else {
		taps.Quit()
		fmt.Println("Usage: apb <DB name>")
		return
	}

	//@@@@@

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
