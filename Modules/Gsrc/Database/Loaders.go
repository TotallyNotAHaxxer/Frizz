package Frizz_Database_Loader

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
)

func Conver_To_JSON(filename string, structure interface{}, mode fs.FileMode) {
	file, _ := json.MarshalIndent(structure, "", " ")
	_ = ioutil.WriteFile(filename, file, mode)
}
