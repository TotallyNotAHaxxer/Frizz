package Frizz_WIFI

import (
	Frizz_Helpers "main/Modules/Gsrc/Helpers"
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"strings"
)

// Dangerous.go is a file that contains an IDS in a way that is able to detect dangerous or possibly
// suspicous SSID's such as ones that are too long, named pwned or something along those lines

var PwnedCovers = []string{
	"pwned",
	"pwn",
	"lolz",
	"anonymous",
	"PwNeD",
	"pawn",
	"Pwner",
	"pwner",
	"PWNER",
	"PAWN",
	"PWNED",
	"own",
	"never gonna give you up",
	"never gonna let you down",
	"never gonna run around and desert you",
	"never gonna make you cry",
	"never gonna say goodbye",
	"never gonna tell a lie and hurt you",
	"fucked",
	"radiojam",
	"LULZ",
	"lulz",
}

const (
	DangerousWarnLen = 29 // If the SSID is more than 29 characters long it could be assumed that this is a deauther or weird SSID
)

type DangerousData struct {
	DangerouSSID   []string
	Dangerousstr   []string
	Dangerousnum   []string
	NumLengthBased int
	NumStringBased int
}

var DangerousStructure DangerousData

func CheckData() {
	for i := 0; i < len(DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID); i++ {
		if len(DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID[i]) >= DangerousWarnLen {
			DangerousStructure.DangerouSSID = append(DangerousStructure.DangerouSSID, DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID[i])
			DangerousStructure.Dangerousnum = append(DangerousStructure.Dangerousnum, DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID[i])
			DangerousStructure.DangerouSSID = Frizz_Helpers.ValueRemover(DangerousStructure.DangerouSSID)
			DangerousStructure.Dangerousnum = Frizz_Helpers.ValueRemover(DangerousStructure.Dangerousnum)
		}
		for k := 0; k < len(PwnedCovers); k++ {
			if strings.Compare(DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID[i], PwnedCovers[k]) == 0 {
				DangerousStructure.DangerouSSID = append(DangerousStructure.DangerouSSID, DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID[i])
				DangerousStructure.DangerouSSID = Frizz_Helpers.ValueRemover(DangerousStructure.DangerouSSID)
				DangerousStructure.Dangerousstr = append(DangerousStructure.Dangerousstr, DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID[i])
				DangerousStructure.Dangerousstr = Frizz_Helpers.ValueRemover(DangerousStructure.Dangerousstr)
			}
		}
	}
	DangerousStructure.NumLengthBased = len(DangerousStructure.Dangerousnum)
	DangerousStructure.NumStringBased = len(DangerousStructure.Dangerousstr)
}
