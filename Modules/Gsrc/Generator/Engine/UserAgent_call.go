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
│           Professional Digital forensics, Network hacking, Stegonography, Recon, OSINT, Bluetooth, CAN and Web Exploitation Expert Secruity Team        │
│																																						  │
│																																						  │
│																																						  │
│																																						  │
│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│
│																Package: Engine															             	  │
│																																						  │
│	This series titles UserAgents is a remake and better or optimized to fit the Frizz project, the members of the scare security development team do     │
│	Not claim full credits for the files as this software is re modified and a re written version og gouseragent which is a very small library to parse   │
│	Useragents and gain information off of them such as their OS, Type, verison, URL etc. Please make sure you understand this before making any claims   │
│								against or for the scare security development team or any contributors to the frizz NFAT project.                         │
│━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━│
│																																						  │
│Package status		     -> OK | Working																												  │
│Security status         -> OK | Secure																													  │
│Performance Status      -> OK | Performant 																											  │
│Bug Status              -> OK | NONE																													  │
│Error status            -> OK | NONE																													  │
└──────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────┘

In later versions of frizz more information for the user agent will be able to be parsed, this was a better rewrite which made more sense to frizz. Taking in the fact

that we want to give users as much information as possible but not too much. the OS name, agent and host is decent enough information for a beta release and the time being
*/
package Engine

import (
	"fmt"
	"strings"
)

type BrandCounter struct {
	Apple     int
	XiaoMi    int
	Google    int
	Microsoft int
}

var BC BrandCounter

func AgentLoader(agent string) UserAgentData {
	useragent := UserAgentData{
		String: agent,
	}
	tokens := Parser(agent)
	for agentindex := range CheckTokens {
		switch {
		case tokens.Exists(fmt.Sprint(CheckTokens[agentindex])):
			if OperatingSystemNmaes[strings.Trim(CheckTokens[agentindex], " ")] != "" {
				useragent.OS = OperatingSystemNmaes[strings.Trim(CheckTokens[agentindex], " ")]
				OS_Counter(useragent.OS)
				tok := fmt.Sprint(CheckTokens[agentindex])
				if tok == "Windows" ||
					tok == "Windows NT" ||
					tok == "Windows XP" ||
					tok == "Windows - Phone" {
					BC.Microsoft += 1
				}

				if tok == "Chrome" || tok == "Brave Chrome" || tok == "CriOS" || tok == "Googlebot" {
					BC.Google += 1
				}

				if tok == "XiaoMi" {
					BC.XiaoMi += 1
				}
				if tok == "iPhone" || tok == "iPad" || tok == "Macintosh" || tok == "OPiOS" {
					BC.Apple += 1
				}
			}
		}
	}
	return useragent
}
