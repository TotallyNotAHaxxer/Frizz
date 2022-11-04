package Frizz_Engine

import (
	"encoding/json"
	"fmt"
	FrizzHelp "main/Modules/Gsrc/Helpers"
	"os"
)

var (
	File_Database     = "Modules/Gsrc/Database/FrizzDatabase.json"
	File_PreProcessor = "Modules/Gsrc/Database/PreProcessor.json"
)

func Loader(rangevalue []string) {
	for k := 0; k < len(rangevalue); k++ {
		fmt.Println(rangevalue[k])
	}
}

func Decode_Json_And_Generate() {
	data, x := FrizzHelp.Fjloader(File_Database)
	if x != nil {
		fmt.Println("Could not load database file")
		os.Exit(1)
	}
	var vals Frizz
	json.Unmarshal(data, &vals)
	Loader(vals.Imapcreds.IMAPBASE64Decoded)
	Loader(vals.SMTPSessionInf.Body)
}
