//go:build linux

package main

import (
	"encoding/json"
	"io/ioutil"
	FrizzRead "main/Modules/File_Loaders"
	FrizzPreproc "main/Modules/Server_Info"
)

/*
You may be wondering why exactly this project also uses go, given the circumstances with PCAPPLUSPLUS returning some wrongly formated data
which results in segmentation faults, we are going to be using go for the mean time. There is a firm reason go will be used not just for pcap parsing
but its main appearence is templating and pcap parsing.
*/

var PreProcess FrizzPreproc.Server

func init() {
	PreProcess.CPU()
	PreProcess.HOST()
	PreProcess.OS()
	PreProcess.MEM()
	file, _ := json.MarshalIndent(PreProcess, "", " ")
	_ = ioutil.WriteFile("test.json", file, 0644)
}

func main() {
	///home/xea43p3x/Desktop/Projects/frizz/src/PCAP/Pcap_Examples/Ftp.pcap
	///home/xea43p3x/Desktop/Projects/frizz/src/PCAP/Pcap_Examples/IMAP - Authenticate Plain (Base64).pcap
	FrizzRead.Reader("Modules/EXAMPLE_PCAP/MASHED.pcap")
}
