package Frizz_WIFI

import (
	"bytes"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	ll "github.com/google/gopacket/layers"
)

var WPAS = []byte{
	0, 0x50,
	0xf2, 1,
}

func Parse_ENC(p gopacket.Packet, d *ll.Dot11) (bool, string) {
	switch d.Flags {
	case ll.Dot11FlagsWEP:
		F = true
		ENCRYPTION = "WEP"
	}
	for _, l := range p.Layers() {
		if l.LayerType() == ll.LayerTypeDot11InformationElement {
			i, k := l.(*ll.Dot11InformationElement)
			if k {
				F = true
				if i.ID == layers.Dot11InformationElementIDRSNInfo {
					ENCRYPTION = "WPA2"
				} else {
					if ENCRYPTION == "" && i.ID == ll.Dot11InformationElementIDVendor && i.Length >= 8 && bytes.Equal(i.OUI, WPAS) && bytes.HasPrefix(i.Info, []byte{1, 0}) {
						ENCRYPTION = "WPA"
					}
				}

			}
		}
	}
	if ENCRYPTION == "" && F {
		ENCRYPTION = "OPEN"
	}
	return F, ENCRYPTION
}
