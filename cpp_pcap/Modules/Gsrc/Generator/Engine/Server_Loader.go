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
│ This file is apart of the Engine package which is the file and database loader for the server json files THIS DOES NOT PARSE THE DATABASE               │
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

import "strings"

type ServerInfo struct {
	Server_Urls         []string
	Server_Files        []string
	Server_Database     []string
	Server_PreProcessor []string
	Server_Support      []string
	Server_Languages    []string
	Server_Imports      []string
	Server_Suggests     []string
	Server_Version      string
	Server_Main_Port    int
	Server_Main_URL     string
	Server_Ports        []int
}

var (
	Languages         = []string{"JavaScript", "Jquery", "HTML", "CSS"}
	Server_Support    = []string{"Chrome", "Firefox", "Brave Chrome", "Edge"}
	SuggestedBrowsers = []string{"Firefox", "Chrome"}
	Version           = "5.0.7 [ ALPHA ]"
	MainURL           = "https://localhost:5674"
	MainPort          = 5674
	SI                ServerInfo
)

func LoadOntoType() {
	for i := range Files {
		SI.Server_Files = append(SI.Server_Files, Files[i])
	}
	for k := range Database {
		SI.Server_Database = append(SI.Server_Database, Database[k])
	}
	var Possible_URLS = []string{
		"http://localhost:5674",
		"http://0.0.0.0:5674",
		"http://127.0.0.1:5674",
	}
	for l := range Possible_URLS {
		SI.Server_Urls = append(SI.Server_Urls, Possible_URLS[l])
		if strings.Contains(Possible_URLS[l], "PreProcessor") {
			SI.Server_PreProcessor = append(SI.Server_PreProcessor, Possible_URLS[l])
		}
	}
	SI.Server_Support = Server_Support
	SI.Server_Suggests = SuggestedBrowsers
	SI.Server_Languages = Languages
	SI.Server_Version = Version
	SI.Server_Main_URL = MainURL
	SI.Server_Main_Port = MainPort
	SI.Server_Imports = importlist
	SI.Server_Ports = []int{8080, 5674}

	Marshal(SI, "Modules/Gsrc/Database/ServerInfo.json")
}
