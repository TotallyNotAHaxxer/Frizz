package Frizz_Loader

import (
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
		os.Exit(0)
	}
	defer handler.Close()
	packetsrc := gopacket.NewPacketSource(handler, handler.LinkType())
	var total int
	STT := time.Now()
	var packarr []gopacket.Packet
	for packetsrc := range packetsrc.Packets() {
		Email.Match_Email_Information(packetsrc)
		FrizzCred.Decoder_Credentials(packetsrc)
		FrizzCred.Decoder_IMAP_CREDS(packetsrc)
		FrizzWifu.Processor(packetsrc)
		FrizzSession.GetObject(packetsrc)
		FrizzSession.GetSession(packetsrc)
		FrizzSession.LoadSMTPSession(packetsrc)
		FrizzSession.GetBody(packetsrc)
		FrizzNetw.Parse_ARP(packetsrc)
		FrizzNetw.FindPortsSrc(packetsrc)
		FrizzNetw.LocateEthernet(packetsrc)
		FrizzNetw.Aw.LoadRaw(packetsrc)
		FrizzNetw.FindServerLDAP(packetsrc)

		total++
		packarr = append(packarr, packetsrc)
	}
	FrizzNetw.StartDraw()               // Start drawing HTTP document, I know this is a bit of a mess and a tad bit wacky but due to current support and issues as well as bugs this will draw here
	FrizzNetw.Read_TCP_TO_HTTP(packarr) // itterate through every bit of packet array and store in template raw

	FrizzNetw.EndDocument() // End the document
	ENT := time.Now()
	Analytics.Filename = pcapf
	Analytics.TotalPackets = total
	Analytics.TimeToParse = ENT.Sub(STT).Seconds()
	Converter()
}
