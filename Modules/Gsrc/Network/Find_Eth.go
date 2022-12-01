package Frizz_Net

import (
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type EthernetData struct {
	Network  []string
	TotalIPS int
	TotalMAC int
	Ouis     int
}

var Ethdata EthernetData

func LocateEthernet(packet gopacket.Packet) {
	var netnet string
	pkt := packet.Layer(layers.LayerTypeEthernet)
	if pkt != nil {
		PKTDT := pkt.(*layers.Ethernet)
		if PKTDT != nil {
			// now test for TCP
			PKTIP := packet.Layer(layers.LayerTypeIPv4)
			if PKTIP != nil {
				pktIP := PKTIP.(*layers.IPv4)
				if pktIP != nil {
					// Now add data for the network
					// Network string will look like
					// SrcAddr->DstAddr->SrcMac->DstMac->DstMacOUI->SrcMacOUI
					netnet += pktIP.SrcIP.String() + "$" + pktIP.DstIP.String() + "$" + PKTDT.SrcMAC.String() + "$" + PKTDT.DstMAC.String() + "$"
					if PKTDT.DstMAC.String() != "" {
						Ethdata.TotalMAC++
					}
					if PKTDT.SrcMAC.String() != "" {
						Ethdata.TotalMAC++
					}
					if pktIP.SrcIP.String() != "" {
						Ethdata.TotalIPS++
					}
					if pktIP.DstIP.String() != "" {
						Ethdata.TotalIPS++
					}
					// Get OUI
					dstmacoui := OUI(string(PKTDT.DstMAC.String()))
					if fmt.Sprint(dstmacoui) != "Unknown" {
						Ethdata.Ouis++
					}
					srcmacoui := OUI(string(PKTDT.SrcMAC.String()))
					if fmt.Sprint(srcmacoui) != "Unknown" {
						Ethdata.Ouis++
					}
					var foui, soui string
					foui = fmt.Sprint(dstmacoui)
					soui = fmt.Sprint(srcmacoui)
					netnet += foui + "$" + soui
				}
			}
		}
	}
	Ethdata.Network = append(Ethdata.Network, netnet)
}
