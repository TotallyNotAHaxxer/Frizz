package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	FrizzRead "main/Modules/Gsrc/File_Loaders"
	FrizzEngine "main/Modules/Gsrc/Generator"
	FrizzEngineWrite "main/Modules/Gsrc/Generator/Engine"
	FrizzNetwork "main/Modules/Gsrc/Network"
	FrizzPreproc "main/Modules/Gsrc/Server_Info"
	FrizzExtra "main/Modules/Gsrc/SubTools"
	"math/rand"
	"os"
	"strconv"
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

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func RandomFile(lent int) string {
	lengthlength := len(chars)
	bytemap := make([]byte, lengthlength)
	rand.Read(bytemap)
	for p := 0; p < lengthlength; p++ {
		bytemap[p] = chars[int(bytemap[p])%lengthlength]
	}
	return string(bytemap) + ".pcap"
}

func main() {
	// multiple arguments
	a := os.Args[1]
	if os.Args[2] != "" {
		if a == "..." {
			// This is a string value of commands
			if os.Args[2] == "masher" {
				listoffiles := os.Args[3]
				outputfile := os.Args[4]
				FrizzExtra.FLOADER(listoffiles, outputfile)
			} else if os.Args[2] == "lineextract" {
				var D FrizzExtra.Data
				D.Filename = os.Args[3]
				D.Snaplen = 1024
				statment, x := strconv.Atoi(os.Args[4])
				if x != nil {
					log.Fatal(x)
				} else {
					D.Line_Num = statment
				}
				D.Output = RandomFile(16)
				fmt.Println("Output file - ", D.Output)
				D.Loader()
			}
		} else {
			if a != "" {
				FrizzRead.Reader(a)
			}
			FrizzEngine.Run()
		}
	}
	FrizzEngineWrite.Write("Modules/Server/HTML/HTTPSESSION.html", FrizzNetwork.Doc.DocDoc)
	// For some reason this file was getting over written
}
