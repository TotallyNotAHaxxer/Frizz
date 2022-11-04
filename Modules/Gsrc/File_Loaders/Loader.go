package Frizz_Loader

import (
	"fmt"
	"os"

	FrizzCred "main/Modules/Gsrc/Credentials"
	Email "main/Modules/Gsrc/Email"
	Engine "main/Modules/Gsrc/Engine"
	FrizzNetw "main/Modules/Gsrc/Network"
	FrizzSession "main/Modules/Gsrc/Sessions"
	FrizzDB "main/Modules/Gsrc/TypeVar"
	FrizzWifu "main/Modules/Gsrc/Wifi"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

func Converter() {
	FrizzDB.GenerateField()
	Engine.Decode_Json_And_Generate()
}

func Reader(pcapf string) {
	handler, x := pcap.OpenOffline(pcapf)
	if x != nil {
		fmt.Println("[-] ERROR: Could not process file -> ", x)
		os.Exit(0)
	}
	defer handler.Close()
	packetsrc := gopacket.NewPacketSource(handler, handler.LinkType())
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
	}
	Converter()
}
