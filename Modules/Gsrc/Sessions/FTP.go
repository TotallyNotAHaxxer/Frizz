package Frizz_Sessions

/*

This method of finding session data is actually a little to simple, given this is the beta version of frizz it does not need to be the best

however later on it may be nice to upgrade this to be more "smarter" and efficient than it currently is

this falls the same for all systems
		| -> Find the TCP packet
				|-> Check port
						|-> If port number is standard then continue
								|-> If body is not empty then append to list and store in DB

*/

import (
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

func GetObject(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp, ok := lay.(*layers.TCP); ok {
			if tcp.DstPort.String() == "21(ftp)" || tcp.SrcPort.String() == "21(ftp)" {
				if string(tcp.Payload) != "" {
					DatabaseVar.DatabaseVariable.FTPS.Data = append(DatabaseVar.DatabaseVariable.FTPS.Data, string(tcp.Payload))
				}
			}
		}
	}

}
