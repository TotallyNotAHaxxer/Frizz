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
	SRCMACOUI    []string
	DSTIP        []string
	Network      []string
	UnknownOUIs  int
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
	if addr != "" {
		return true
	} else {
		return false
	}
}

func Parse_ARP(packet gopacket.Packet) {
	ARPL := packet.Layer(layers.LayerTypeARP)
	if ARPL != nil {
		var network string
		arppack, _ := ARPL.(*layers.ARP)
		MACSRC := fmt.Sprintf("%v", net.HardwareAddr(arppack.SourceHwAddress))
		MACDST := fmt.Sprintf("%v", net.HardwareAddr(arppack.DstHwAddress))
		IPSRC := fmt.Sprintf("%v", net.IP(arppack.SourceProtAddress))
		IPDST := fmt.Sprintf("%v", net.IP(arppack.DstProtAddress))
		if CheckAddr(IPDST) {
			Adata.DSTIP = append(Adata.DSTIP, IPDST)
		}
		if CheckAddr(MACDST) {
			Adata.DSTMAC = append(Adata.DSTMAC, MACDST)
		}
		if CheckAddr(MACSRC) {
			Adata.SRCMAC = append(Adata.SRCMAC, MACSRC)
		}
		if CheckAddr(IPSRC) {
			Adata.SRCIP = append(Adata.SRCIP, IPSRC)
		}
		var dstoui, srcoui string
		dstoui = fmt.Sprint(OUI(MACDST))
		srcoui = fmt.Sprint(OUI(MACSRC))
		Adata.SRCMACOUI = append(Adata.SRCMACOUI, srcoui)
		Adata.DSTMACOUI = append(Adata.DSTMACOUI, dstoui)
		network += string(IPSRC) + "@" +
			string(MACSRC) + "@" +
			string(IPDST) + "@" +
			string(MACDST) + "@" +
			dstoui + "@" +
			srcoui + "@"
		Adata.Network = append(Adata.Network, network)
	}
}
