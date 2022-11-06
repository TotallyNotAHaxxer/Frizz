package Engine

import (
	"fmt"
	FrizzHelp "main/Modules/Gsrc/Helpers"
	"os"
)

// Function allows for constant ranges, instead of constant for loops
func Loader(rangevalue []string) {
	for k := 0; k < len(rangevalue); k++ {
		fmt.Println(rangevalue[k])
	}
}

func ReturnJsonInterface(filename string) []byte {
	data, x := FrizzHelp.Fjloader(filename)
	if x != nil {
		fmt.Println("Could not load database file")
		os.Exit(1)
		return nil
	} else {
		return data
	}
}
