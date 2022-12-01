package Frizz_Generation

import (
	GenEngine "main/Modules/Gsrc/Generator/Engine"
)

var (
	oc           = make(chan string)
	DatabaseChan = make(chan string)
)

func Run() {
	go GenEngine.Find__Send(DatabaseChan) // Loads database files
	GenEngine.Read__Store(DatabaseChan)   // Stores Data from files in type
	go GenEngine.Producer(oc)
	GenEngine.Handler(oc)
	GenEngine.Read() //
}
