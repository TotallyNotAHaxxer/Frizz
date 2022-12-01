package Frizz_Net

import (
	"fmt"
	"net"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/macs"
)

// Here is how this will work
/*

The HTML document needs to count how many times a given hardware element has been caught

for example

hardware element -> samsungG11 phone was caught 15 times

basically we take an array that has that name in it such as

{
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
	"samsungG11",
}

count how many times it repeated, generate the HTML then erase the elements.

in other words

Parse data -> Load struct during HTML generation time -> load bar chart values -> erase repeating elements -> generate table

*/
// repeat functions ew
func OUI(m string) []string {
	var alist []string
	if mac, x := net.ParseMAC(m); x == nil {
		prefix := [3]byte{
			mac[0],
			mac[1],
			mac[2],
		}
		manufacturer, e := macs.ValidMACPrefixMap[prefix]
		if e {
			alist = append(alist, manufacturer)
		}
		return alist
	}
	return []string{"Unknown"}
}

type ARPData struct {
	SRCMAC       []string
	DSTMACOUI    []string
	DSTMAC       []string
	SRCIP        []string
	DSTIP        []string
	AmountSrcIP  int
	AmountDstIP  int
	AmountSrcMAC int
	AmountDstMAC int
}

func CountData(arz []string) map[string]int {
	lp := make(map[string]int)
	for _, l := range arz {
		lp[l] = lp[l] + 1
	}
	return lp
}

var Adata ARPData

func CheckAddr(addr string) bool {
	if addr != "" && addr != "ff:ff:ff:ff:ff:ff" && addr != "00:00:00:00:00:00" && addr != "0.0.0.0" {
		return true
	} else {
		return false
	}
}

func Parse_ARP(packet gopacket.Packet) {
	ARPL := packet.Layer(layers.LayerTypeARP)
	if ARPL != nil {
		arppack, _ := ARPL.(*layers.ARP)
		if CheckAddr(string(arppack.DstProtAddress)) {
			Adata.DSTIP = append(Adata.DSTIP, string(arppack.DstProtAddress))
			fmt.Println(string(arppack.DstHwAddress))
		}
		if CheckAddr(string(arppack.DstHwAddress)) {
			Adata.DSTMAC = append(Adata.DSTMAC, string(arppack.DstHwAddress))
		}
		if CheckAddr(string(arppack.SourceHwAddress)) {
			Adata.SRCMAC = append(Adata.SRCMAC, string(arppack.SourceHwAddress))
		}
		if CheckAddr(string(arppack.SourceProtAddress)) {
			Adata.SRCIP = append(Adata.SRCIP, string(arppack.SourceProtAddress))
		}
		if OUI(string(arppack.DstHwAddress)) != nil {
			getOUI := OUI(string(arppack.DstHwAddress))
			for _, l := range getOUI {
				Adata.DSTMACOUI = append(Adata.DSTMACOUI, l)
				break // we only want one OUI
			}
		}
	}
}
