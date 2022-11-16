package main

import (
	"encoding/json"
	"io/ioutil"
	FrizzRead "main/Modules/Gsrc/File_Loaders"
	FrizzEngine "main/Modules/Gsrc/Generator"
	FrizzPreproc "main/Modules/Gsrc/Server_Info"
	"os"
)

/*
You may be wondering why exactly this project also uses go, given the circumstances with PCAPPLUSPLUS returning some wrongly formated data
which results in segmentation faults, we are going to be using go for the mean time. There is a firm reason go will be used not just for pcap parsing
but its main appearence is templating and pcap parsing.

this file gets run by loader.cpp securely and well
*/

var PreProcess FrizzPreproc.Server

func init() {
	PreProcess.CPU()
	PreProcess.HOST()
	PreProcess.OS()
	PreProcess.MEM()
	file, _ := json.MarshalIndent(PreProcess, "", " ")
	_ = ioutil.WriteFile("Modules/Gsrc/Database/PreProcessor.json", file, 0644)
}

func main() {
	a := os.Args[1]
	if a != "" {
		FrizzRead.Reader(a)
	}
	FrizzEngine.Run()
}
