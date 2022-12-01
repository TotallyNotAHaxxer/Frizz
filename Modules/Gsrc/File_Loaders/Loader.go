package Frizz_Loader

import (
	"fmt"
	"os"
	"time"

	FrizzCred "main/Modules/Gsrc/Credentials"
	Email "main/Modules/Gsrc/Email"
	FrizzNetw "main/Modules/Gsrc/Network"
	FrizzSession "main/Modules/Gsrc/Sessions"
	FrizzDB "main/Modules/Gsrc/TypeVar"
	FrizzWifu "main/Modules/Gsrc/Wifi"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Converter() {
	FrizzDB.GenerateField()
}

type ServerAnalyticsData struct {
	TotalPackets int
	TimeToParse  float64
	Filename     string
}

var Analytics ServerAnalyticsData

func Reader(pcapf string) {
	handler, x := pcap.OpenOffline(pcapf)
	if x != nil {
		fmt.Println("[-] ERROR: Could not process file -> ", x)
		os.Exit(0)
	}
	defer handler.Close()
	packetsrc := gopacket.NewPacketSource(handler, handler.LinkType())
	var total int
	STT := time.Now()
	for packetsrc := range packetsrc.Packets() {
		Email.Match_Email_Information(packetsrc)
		FrizzCred.Decoder_Credentials(packetsrc)
		FrizzCred.Decoder_IMAP_CREDS(packetsrc)
		FrizzNetw.Read_TCP_TO_HTTP(packetsrc)
		FrizzWifu.Processor(packetsrc)
		FrizzSession.GetObject(packetsrc)
		FrizzSession.GetSession(packetsrc)
		FrizzSession.LoadSMTPSession(packetsrc)
		FrizzSession.GetBody(packetsrc)
		FrizzNetw.Parse_ARP(packetsrc)
		FrizzNetw.FindPortsSrc(packetsrc)
		FrizzNetw.LocateEthernet(packetsrc)
		total++
	}
	ENT := time.Now()
	Analytics.Filename = pcapf
	Analytics.TotalPackets = total
	Analytics.TimeToParse = ENT.Sub(STT).Seconds()
	Converter()
}
