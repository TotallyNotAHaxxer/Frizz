package Frizz_Net

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

type Servers struct {
	LDAP ldap
}

type ldap struct {
	ServerAddresses []string
}

var (
	Server Servers
)

func FindServerLDAP(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp, ok := lay.(*layers.TCP); ok {
			if tcp.SrcPort.String() == "389(ldap)" {
				if eth := packet.Layer(layers.LayerTypeIPv4); eth != nil {
					if eth := eth.(*layers.IPv4); eth != nil {
						Server.LDAP.ServerAddresses = append(Server.LDAP.ServerAddresses, eth.SrcIP.String())
					}
				}
			}
			if tcp.DstPort.String() == "389(ldap)" {
				if eth := packet.Layer(layers.LayerTypeIPv4); eth != nil {
					if eth := eth.(*layers.IPv4); eth != nil {
						Server.LDAP.ServerAddresses = append(Server.LDAP.ServerAddresses, eth.DstIP.String())
					}
				}
			}
		}
	}
}
