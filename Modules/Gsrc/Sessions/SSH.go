package Frizz_Sessions

import (
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func GetSession(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp, ok := lay.(*layers.TCP); ok {
			if tcp.SrcPort.String() == "21(ssh)" || tcp.DstPort.String() == "21(ssh)" {
				if string(tcp.Payload) == "" {
					DatabaseVar.DatabaseVariable.SSHS.Body = append(DatabaseVar.DatabaseVariable.SSHS.Body, string(tcp.Payload))
				}
			}
		}
	}
}
