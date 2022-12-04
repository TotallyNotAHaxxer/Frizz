package Frizz_Net

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type FrizzPortsSrc struct {
	Hostname            []string
	PortsOfHost         []string
	Network             []string
	TotalEthernetandTCP int
}

var FrizP FrizzPortsSrc

func FindPortsSrc(pkt gopacket.Packet) {
	var net string
	Tcp := pkt.Layer(layers.LayerTypeTCP)
	if Tcp != nil {
		FrizP.TotalEthernetandTCP++
		tcpsrc := Tcp.(*layers.TCP)
		Ipv4 := pkt.Layer(layers.LayerTypeIPv4)
		if Ipv4 != nil {
			ipv4 := Ipv4.(*layers.IPv4)
			FrizP.Hostname = append(FrizP.Hostname, ipv4.SrcIP.String())
			net += ipv4.SrcIP.String() + "$"
			FrizP.TotalEthernetandTCP++
		}
		FrizP.PortsOfHost = append(FrizP.PortsOfHost, tcpsrc.SrcPort.String())
		net += tcpsrc.SrcPort.String()
	}
	FrizP.Network = append(FrizP.Network, net)
}
