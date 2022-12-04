/*
┌──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┐
│    ___o .--.               ____                                          ____                                      __                      ___o .--.    │
│   /___| |--|              /\  _`\                                       /\  _`\                                 __/\ \__                  /___| |OO|    │
│  /'   |_|  |_             \ \,\L\_\    ___     __     _ __    __        \ \,\L\_\     __    ___   __  __  _ __ /\_\ \ ,_\  __  __        /'   |_|  |_   │
│       (_    _)             \/_\__ \   /'___\ /'__`\  /\`'__\/'__`\  __o__\/_\__ \   /'__`\ /'___\/\ \/\ \/\`'__\/\ \ \ \/ /\ \/\ \           (_     _)  │
│       | |   \                /\ \L\ \/\ \__//\ \L\.\_\ \ \//\  __/    |    /\ \L\ \/\  __//\ \__/\ \ \_\ \ \ \/ \ \ \ \ \_\ \ \_\ \           | |   \   │
│       | |___/                \ `\____\ \____\ \__/.\_\\ \_\\ \____\  / \   \ `\____\ \____\ \____\\ \____/\ \_\  \ \_\ \__\\/`____ \          | |___/   │
│                               \/_____/\/____/\/__/\/_/ \/_/ \/____/  _______\/_____/\/____/\/____/ \/___/  \/_/   \/_/\/__/ `/___/> \                   │
│                                                                     /\______\                                                  /\___/                   │
│                                                                     \/______/                                                  \/__/                    │
│                                                                                                                                                         │
│           Professional Digital forensics, Network hacking, Stegonography, Recon, OSINT, Bluetooth, CAN and Web Exploitation Secruity Team               │
│																																						  │
│																																						  │
│																																						  │
│																																						  │
│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│
│																Package: Engine															             	  │
│																																						  │
│	This file holds all of the file based functions such as creation, deletion and checking if the file exists even stripping file names and titles       │
│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│
│																																						  │
│Package status		     -> OK | Working																												  │
│Security status         -> OK | Secure																													  │
│Performance Status      -> OK | Performant 																											  │
│Bug Status              -> OK | NONE																													  │
│Error status            -> OK | NONE																													  │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘
*/
package Engine

import (
	"io/ioutil"
	"os"
)

type OSCounter struct {
	Windows int
	Linux   int
	IOS     int
}

var OSC OSCounter

func OS_Counter(vars string) {
	if vars == "Windows" {
		OSC.Windows++
	} else if vars == "Linux" {
		OSC.Linux++
	} else if vars == "IOS" || vars == "Mac OS X" || vars == "iOS" || vars == "iPad" || vars == "iPhone" {
		OSC.IOS++
	}
}

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
		//("[Engine] Error: Could not create file: ", x)
	} else {
		defer f.Close()
		//("[Engine] Stat: Successfully created file: ", filename)
	}
}

func Reload(filename string) {
	x := os.Remove(filename)
	if x != nil {
		//("[Engine] Error: Could not remove file: ", x)
	} else {
		//("[Engine] Stat: Successfully removed file: ", filename)
	}
}

// Write data to the file
func Write(filename, data string) error {
	if ExistQ(filename) {
		x := ioutil.WriteFile(filename, []byte(data), 0)

		if x != nil {
			return x
		}
	}
	return nil
}
