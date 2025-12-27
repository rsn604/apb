package main

import (
	"apb/appt"
	"encoding/json"
	"fmt"
	"os"
)

var databaseName = "BOLT"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: apbunload <DB name>")
		return
	}

	manager := appt.GetManager(databaseName)
	err := manager.Connect(databaseName, os.Args[1])
	if err != nil {
		panic(err)
	}
	manager.Define()
	appts := manager.GetAppts()
	manager.Close()
	for _, ap := range appts {
		buf, _ := json.Marshal(ap)
		fmt.Println(string(buf))
	}
}
