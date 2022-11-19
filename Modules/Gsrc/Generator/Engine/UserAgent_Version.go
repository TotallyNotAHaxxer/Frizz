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

import "strings"

var CheckTokens = []string{
	"Android", "iPhone", "iPad", "Windows NT", "Windows Phone OS",
	"Macintosh", "Macintosh", "Macintosh", "Googlebot", "Applebot",
	"Opera Mini", "OPR", "OPT", "OPiOS", "CriOS",
	"MSIE", "AdsBot-Google-Mobil", "Mediapartners-Google", "AdsBot-Google",
	"XiaoMi", "Chrome", "Brave Chrome", "Safari"}

var GetTokens = []string{
	"FxiOS", "Firefox", "vivaldi",
	"EdgiOS", "Edge", "Edg", "Edga",
	"bingbot", "YandexBot", "SamsungBrowser",
	"HeadlessChrome",
}

var OperatingSystemNmaes = map[string]string{
	"Android":          And,
	"iPhone":           IOS,
	"iPad":             IOS,
	"Windows NT":       Win,
	"Windows Phone OS": WinP,
	"Macintosh":        Mac,
	"Linux":            Linux,
	"CrOS":             ChromeOS,
	"linux-gnu":        Linux,
	"linux":            Linux,
}

/*

	"Opera":               Opera,
	"Opera Mini":          OperaMini,
	"Opera Touch":         OperaTouch,
	"Chrome":              Chrome,
	"HeadlessChrome":      HeadlessChrome,
	"Firefox":             Firefox,
	"Internet Explorer":   InternetExplorer,
	"Safari":              Safari,
	"Edge":                Edge,
	"Vivaldi":             Vivaldi,
	"Google Ads Bot":      GoogleAdsBot,
	"Googlebot":           Googlebot,
	"Twitterbot":          Twitterbot,
	"facebookexternalhit": FacebookExternalHit,
	"Applebot":            Applebot,
	"Bingbot":             Bingbot,
*/

// This function gets the version or returns the verison of the user agent and its name.
func GetVersion(agent string) (string, string) {
	indexer := strings.LastIndex(agent, " ")
	if indexer == -1 {
		return agent, ""
	}
	switch agent[:indexer] {
	case "Linux", "Windows NT", "Windows Phone OS", "MSIE", "Android":
		return agent[:indexer], agent[indexer+1:]
	case "CrOS x86_64", "CrOS aarch64":
		jaded := strings.LastIndex(agent[:indexer], " ")
		return agent[:jaded], agent[jaded+1 : indexer]
	default:
		return agent, ""
	}

}
