package Frizz_Sessions

import (
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}

func GetBody(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp, ok := lay.(*layers.TCP); ok {
			if tcp.DstPort.String() == "23(telnet)" || tcp.SrcPort.String() == "23(telnet)" {
				APPLET := packet.ApplicationLayer()
				if APPLET != nil {
					DatabaseVar.DatabaseVariable.TELS.Body = append(DatabaseVar.DatabaseVariable.TELS.Body, string(APPLET.Payload()))
				}
			}
		}
	}
}
