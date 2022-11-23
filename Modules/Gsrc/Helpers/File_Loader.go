package Frizz_Helper

import (
	"io/ioutil"
	"os"
)

func Fjloader(filename string) ([]byte, error) {
	f, x := os.Open(filename)
	if x != nil {
		return nil, x
	}
	defer f.Close()
	bv, _ := ioutil.ReadAll(f)
	return bv, nil
}
