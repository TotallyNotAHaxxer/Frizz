package Frizz_WIFI

import (
	"bytes"
	"fmt"
	"time"

	"github.com/google/gopacket"
	ll "github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
)

var (
	X              error
	Monitor        bool
	Timeout        time.Duration
	Controller     *pcap.Handle
	PB             []byte
	PB2            []byte
	PROBE_DATA     []string
	CIPHER         = ""
	AUTHENTICATION = ""
	F              = false
	ENCRYPTION     = ""
)

type Probe struct {
	Probe_SSID      string
	Probe_BSSID     string
	Probe_MAC       string
	Probe_Channel   string
	Probe_Rate      string
	Probe_Vendor    []byte
	Probe_Frequency int8
	Probe_Time      time.Time
}

func (PROBE *Probe) Decoder(l *ll.Dot11MgmtProbeReq) {
	PB = l.LayerContents()
	for k := uint64(0); k < uint64(len(PB)); {
		id := ll.Dot11InformationElementID(PB[k])
		k++
		switch id {
		case ll.Dot11InformationElementIDSSID:
			e := uint64(PB[k])
			k++
			if e > 0 {
				PROBE.Probe_SSID = string(PB[k : k+e])
				k += e
			}
		case ll.Dot11InformationElementIDVendor:
			PROBE.Probe_Vendor = PB[k+1:]
			return

		default:
			e := uint64(PB[k])
			k += 1 + e
		}
	}
}

func Vals(m string, s int) []string {
	by := ""
	byy := []string{}
	ru := bytes.Runes([]byte(m))
	le := len(ru)
	for i, l := range ru {
		by = by + string(l)
		if (i+1)%s == 0 {
			byy = append(byy, by)
		} else if (i + 1) == le {
			byy = append(byy, by)
		}
	}
	return byy
}

func Processor(pack gopacket.Packet) {
	typer := Probe{
		Probe_Time: time.Now(),
	}
	if layer := pack.Layer(ll.LayerTypeDot11); layer != nil {

		question, _ := layer.(*ll.Dot11)
		dt, enc := Parse_ENC(pack, question)
		if dt {
			fmt.Println("ENCRYPTION -> ", enc)
		}

		if question.Address2.String() != "" && question.Address2.String() != "ff:ff:ff:ff:ff:ff" && question.Address2.String() != "00:00:00:00:00:00" {
			typer.Probe_MAC = question.Address2.String()
		}
		if j := pack.Layer(ll.LayerTypeDot11MgmtProbeReq); j != nil {
			dec, _ := j.(*ll.Dot11MgmtProbeReq)
			typer.Decoder(dec)
		}
		if is := pack.Layer(ll.LayerTypeRadioTap); is != nil {
			rad, _ := is.(*ll.RadioTap)
			typer.Probe_Frequency = rad.DBMAntennaSignal
			typer.Probe_Rate = rad.Rate.String()
		}
	}
	addr4 := OUI(typer.Probe_MAC)
	if addr4 != nil {
		if typer.Probe_MAC != "" && typer.Probe_Frequency != 0x00 && typer.Probe_SSID != "" {
			fmt.Println("PROBE MAC 		\t- ", typer.Probe_MAC)
			fmt.Println("PROBE MAC OUI 	\t- ", addr4)
			fmt.Println("PROBE FREQ    	\t- ", typer.Probe_Frequency)
			fmt.Println("PROBE SSID    	\t- ", typer.Probe_SSID)
		}
	}
}
