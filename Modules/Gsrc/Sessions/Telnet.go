package Frizz_Sessions

import (
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func GetBody(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp, ok := lay.(*layers.TCP); ok {
			if tcp.DstPort.String() == "23(telnet)" || tcp.SrcPort.String() == "23(telnet)" {
				if string(tcp.Payload) != "" {
					DatabaseVar.DatabaseVariable.TELS.Body = append(DatabaseVar.DatabaseVariable.TELS.Body, string(tcp.Payload))
				}
			}
		}
	}
}
