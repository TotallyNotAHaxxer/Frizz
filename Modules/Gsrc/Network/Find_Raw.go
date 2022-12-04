package Frizz_Net

import "github.com/google/gopacket"

// literally a raw packet

type Raw struct {
	Packet []gopacket.Packet
}

var Aw Raw

func (r *Raw) LoadRaw(p gopacket.Packet) { r.Packet = append(r.Packet, p) }
