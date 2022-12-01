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
*/
package Engine

import (
	"bytes"
	"strings"
)

var (
	slash, url    bool
	buffer, value bytes.Buffer
)

type UINF struct {
	Key string
	Val string
}

type UINFS struct {
	lists []UINF
}

func Parser(Agent string) UINFS {
	cl := UINFS{
		lists: make([]UINF, 0, 8),
	} // client data
	addToken := func() {
		if buffer.Len() != 0 {
			trimmer := strings.TrimSpace(buffer.String())
			if _, sign := ignore[trimmer]; !sign {
				if url {
					trimmer = strings.TrimPrefix(trimmer, "+")
				}
				if value.Len() == 0 {
					var variable string
					// Check the version of the agent
					trimmer, variable = GetVersion(trimmer)
					// Add token
					cl.add(trimmer, variable)
				} else {
					cl.add(trimmer, strings.TrimSpace(value.String()))
				}
			}
		}
		buffer.Reset()
		value.Reset()
		slash = false
		url = false
	}
	part := false
	bua := []byte(Agent)
	for i, c := range bua {
		switch {
		case c == 41:
			addToken()
			part = false
		case part && c == 59:
			addToken()
		case c == 40:
			addToken()
			part = true
		case slash && c == 32:
			addToken()
		case slash:
			value.WriteByte(c)
		case c == 47 && !url:
			if i != len(bua)-1 && bua[i+1] == 47 && (bytes.HasSuffix(buffer.Bytes(), []byte("http:")) || bytes.HasSuffix(buffer.Bytes(), []byte("https:"))) {
				buffer.WriteByte(c)
				url = true
			} else {
				slash = true
			}
		default:
			buffer.WriteByte(c)
		}
	}
	addToken()
	return cl
}
