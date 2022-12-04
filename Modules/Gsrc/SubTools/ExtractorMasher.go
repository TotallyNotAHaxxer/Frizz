package SubUtils

import (
	"bufio"
	"log"
	"os"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/google/gopacket/pcap"
	"github.com/google/gopacket/pcapgo"
)

var (
	Filepaths  []string // All filepaths loaded from a file
	Good_Files []string // Files that were found and validated
	Bad_Files  []string // Files that were not found or located
	Packets    []gopacket.Packet
)

func Check(filename string) bool {
	if _, x := os.Stat(filename); x != nil {
		return false // Not a file or existing file
	} else {
		return true // is a file and valid
	}
}

func Load(Good_File string, OUTPUT string) {
	handle, x := pcap.OpenOffline(Good_File)
	if x != nil {
		log.Fatal(x)
	}
	defer handle.Close()
	pktsrc := gopacket.NewPacketSource(handle, handle.LinkType())
	for pkt := range pktsrc.Packets() {
		Packets = append(Packets, pkt)
	}
	Saver(OUTPUT)
}

func Saver(OUTF string) {
	f, _ := os.Create(OUTF)
	WRITER := pcapgo.NewWriter(f)
	WRITER.WriteFileHeader(1024, layers.LinkTypeEthernet)
	defer f.Close()
	for _, pkt := range Packets {
		WRITER.WritePacket(pkt.Metadata().CaptureInfo, pkt.Data())
	}

}

// Call F loader
// Dstf is the destination file, or the file that will be read
// dst is the output pcap file after all has been mashed
func FLOADER(dstf, dst string) {
	f, x := os.Open(dstf)
	if x != nil {
		log.Fatal(x)
	} else {
		defer f.Close()
		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			Filepaths = append(Filepaths, scanner.Text())
		}
	}
	for _, k := range Filepaths {
		if Check(k) {
			Good_Files = append(Good_Files, k)
		} else {
			Bad_Files = append(Bad_Files, k)
		}
	}
	for _, p := range Good_Files {
		Load(p, dst)
	}
}
