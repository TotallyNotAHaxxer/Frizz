package Frizz_WIFI

import (
	"bytes"
	"time"

	Frizz_Helpers "main/Modules/Gsrc/Helpers"
	DatabaseVar "main/Modules/Gsrc/TypeVar"

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

type WifiInformation struct {
	SSID   []string
	VENDOR []string
	ENC    []string
	FREQ   []string
	RATE   []string
	ADDR2  []string
}

func Decoder(l *ll.Dot11MgmtProbeReq) {
	PB = l.LayerContents()
	for k := uint64(0); k < uint64(len(PB)); {
		id := ll.Dot11InformationElementID(PB[k])
		k++
		switch id {
		case ll.Dot11InformationElementIDSSID:
			e := uint64(PB[k])
			k++
			if e > 0 {
				DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID = append(DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID, string(PB[k:k+e]))
				DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.WifiProbe.Probe_SSID)
				k += e
			}
		case ll.Dot11InformationElementIDVendor:
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_Vendor = PB[k+1:]
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
	if layer := pack.Layer(ll.LayerTypeDot11); layer != nil {

		question, _ := layer.(*ll.Dot11)
		dt, enc := Parse_ENC(pack, question)
		if dt {
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_Encryption = append(DatabaseVar.DatabaseVariable.WifiProbe.Probe_Encryption, enc)
		}
		if question.Address2.String() != "" && question.Address2.String() != "ff:ff:ff:ff:ff:ff" && question.Address2.String() != "00:00:00:00:00:00" {
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_MAC = append(DatabaseVar.DatabaseVariable.WifiProbe.Probe_MAC, question.Address2.String())
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_MAC = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.WifiProbe.Probe_MAC)
		}
		if j := pack.Layer(ll.LayerTypeDot11MgmtProbeReq); j != nil {
			dec, _ := j.(*ll.Dot11MgmtProbeReq)
			Decoder(dec)
		}
		if is := pack.Layer(ll.LayerTypeRadioTap); is != nil {
			rad, _ := is.(*ll.RadioTap)
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_Frequency = append(DatabaseVar.DatabaseVariable.WifiProbe.Probe_Frequency, rad.DBMAntennaSignal)
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_Rate = append(DatabaseVar.DatabaseVariable.WifiProbe.Probe_Rate, rad.Rate.String())
			DatabaseVar.DatabaseVariable.WifiProbe.Probe_Rate = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.WifiProbe.Probe_Rate)
		}
	}
	CheckData()
}
