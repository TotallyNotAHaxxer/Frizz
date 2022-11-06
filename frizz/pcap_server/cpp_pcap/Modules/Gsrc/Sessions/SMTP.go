package Frizz_Sessions

import (
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func LoadSMTPSession(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp, ok := lay.(*layers.TCP); ok {
			if tcp.SrcPort.String() == "25(smtp)" || tcp.DstPort.String() == "25(smtp)" {
				if string(tcp.Payload) != "" {
					DatabaseVar.DatabaseVariable.SMTP_SessionInf.Body = append(DatabaseVar.DatabaseVariable.SMTP_SessionInf.Body, string(tcp.Payload))
				}
			}
		}
	}
}
