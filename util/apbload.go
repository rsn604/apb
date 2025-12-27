package main

import (
	"apb/appt"
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

var databaseName = "BOLT"

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: apbload <load file> <DB name>")
		return
	}

	fp, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)

	manager := appt.GetManager(databaseName)
	err = manager.Connect(databaseName, os.Args[2])
	if err != nil {
		panic(err)
	}
	manager.Define()

	var ap appt.Appointment
	for scanner.Scan() {
		err := json.Unmarshal([]byte(scanner.Text()), &ap)
		if err != nil {
			fmt.Println(err)
			return
		}
		manager.AddAppt(ap)
	}
	/*
		appts := manager.GetAppts()
		for _, ap:= range(appts){
			buf, _ := json.Marshal(ap)
			fmt.Println(string(buf))
		}
	*/
	manager.Close()
}
