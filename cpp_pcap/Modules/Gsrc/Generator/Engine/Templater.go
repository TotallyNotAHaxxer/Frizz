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
│	This file is the main templating engine, the templater that holds diagrams, for loops itterating over the frizz J structure and the loader for the    │
│   HTTP server. This means this file is completely dedicated to parsing and filling the pre generated HTML templates in TemplaterVars.go. The reason that│
│	we generate the data and fill it into currently existing files is just for templating reasons, everytime this code is ran the channels will be forced │
│	to delete the current HTML files and re generate new ones thus this file and module code will be called to re fill those templates up. Of course this │
│	will not happen everytime you make a GET request but everytime the files or data inside of the structure or JSON file is changed via POST.            │
│   Frizz has a built in packet masher and re uploader to re upload and parse new files, once the POST request is made or found with a specific header    │
│	the go code will be ran by the C++ inline which is run by the server. This makes it a bit more performant rather than keeping it with RPCT or re mod- │
│	ing the files with new templates that are constantly generated one by one by the server. Using threads is a much better way to do this same with      │
│	loading the data from the channels and instead of using Golangs templating language to just manually create the bodies with the engine                │
│																																						  │
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
	"fmt"
)

func Itterator(table bool, tabletype string, Template string, Element ...[]string) string {
	sizeof := len(Element[0]) // First array
	// Range and detect
	for i := 0; i < sizeof; i++ {
		if i > len(Element[1]) || i == len(Element[1]) {
			break
		}
		switch table {
		case true:
			if tabletype == "uagent" {
				Template += "<tr>"
				Template += fmt.Sprintf("<td>%s</td>", Element[0][i])
				Template += fmt.Sprintf("<td>%s</td>", Element[1][i])
				udata := AgentLoader(Element[1][i]) // Load host and agent information
				if udata.OS == "" {
					Template += fmt.Sprintf("<td>%s</td>", "Operating system was not found or supported")
				} else {
					Template += fmt.Sprintf("<td>%s</td>", udata.OS)
				}
				Template += "<tr>"
			}
		}
	}
	Template += Template_EndTags
	Template += Template_CSS
	Template += fmt.Sprintf(TemplatebarvarJS, OSC.Linux, OSC.Windows, OSC.IOS, BC.Apple, BC.Google, BC.XiaoMi, BC.Microsoft)
	Template += TemplatebarJS
	// generate bar values
	if Template != "" {
		return Template
	} else {
		return ""
	}
}

func Read() { // Read go simple
	LoadOntoType()
	var HTMLTemplate string
	HTMLTemplate += Template_Top
	Write("Modules/Server/HTML/Useragents.html", Itterator(true, "uagent", HTMLTemplate, StructureFrizzPointer.Httpd.UagentHostHost, StructureFrizzPointer.Httpd.UagentHostUagent))
}
