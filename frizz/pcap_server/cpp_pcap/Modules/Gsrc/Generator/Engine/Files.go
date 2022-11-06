package Engine

import (
	"fmt"
	"os"
)

func ExistQ(filename string) bool {
	if _, x := os.Stat(filename); x != nil {
		return false // not found
	} else {
		return true // found
	}
}

func Create(filename string) {
	f, x := os.Create(filename)
	if x != nil {
		fmt.Println("[Engine] Error: Could not create file: ", x)
	} else {
		defer f.Close()
		fmt.Println("[Engine] Stat: Successfully created file: ", filename)
	}
}

func Reload(filename string) {
	x := os.Remove(filename)
	if x != nil {
		fmt.Println("[Engine] Error: Could not remove file: ", x)
	} else {
		fmt.Println("[Engine] Stat: Successfully removed file: ", filename)
	}
}
