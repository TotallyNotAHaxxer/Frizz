package Var

import (
	"encoding/json"
	"io/ioutil"
	FrizzDb "main/Modules/Gsrc/Data"
)

var (
	DatabaseVariable FrizzDb.Frizz
)

func GenerateField() {
	file, _ := json.MarshalIndent(DatabaseVariable, "", " ")
	_ = ioutil.WriteFile("Modules/Gsrc/Database/FrizzDatabase.json", file, 0600)
}
