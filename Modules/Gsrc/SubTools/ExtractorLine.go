package SubUtils

import (
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

type Data struct { // Injector data
	Filename string
	Output   string
	Line_Num int
	Snaplen  uint32
	Packet   gopacket.Packet
}

var counter = 0
var handle *pcap.Handle
var x error

func Empty(object interface{}) bool {
	if object != nil {
		return false
	} else {
		return true
	}
}

func (Struct *Data) Inject() {
	f, x := os.Open("output.pcap")
	if x != nil {
		f, x = os.Create("output.pcap")
	}
	w := pcapgo.NewWriter(f)
	w.WriteFileHeader(Struct.Snaplen, layers.LinkTypeEthernet)
	w.WritePacket(Struct.Packet.Metadata().CaptureInfo, Struct.Packet.Data())
	defer f.Close()
}

func (Struct *Data) Loader() {
	handle, x = pcap.OpenOffline(Struct.Filename)
	if x != nil {
		log.Fatal(x)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		counter++
		if counter == Struct.Line_Num {
			if packet != nil {
				Struct.Packet = packet
				break
			}
		}
	}
	Struct.Inject()
}
