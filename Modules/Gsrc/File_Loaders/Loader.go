package Frizz_Loader

import (
	"fmt"
	"os"

	FrizzCred "main/Modules/Gsrc/Credentials"
	FrizzConvJ "main/Modules/Gsrc/Database"
	FrizzPOP3 "main/Modules/Gsrc/Email"
	FrizzNetw "main/Modules/Gsrc/Network"
	FrizzWifu "main/Modules/Gsrc/Wifi"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	FRIZZNETWORK FrizzNetw.HTTP_DATA
	EMAILSTRUCT  FrizzPOP3.Email_Information
	CREDENTIALS  FrizzCred.Credentials
	IMAPCREDS    FrizzCred.IMAP_Credential
	FRIZZWIFI    FrizzWifu.Probe
)

func Reader(pcapf string) {
	handler, x := pcap.OpenOffline(pcapf)
	if x != nil {
		fmt.Println("[-] ERROR: Could not process file -> ", x)
		os.Exit(0)
	}
	defer handler.Close()
	packetsrc := gopacket.NewPacketSource(handler, handler.LinkType())
	for packetsrc := range packetsrc.Packets() {
		EMAILSTRUCT.Match_Email_Information(packetsrc)
		CREDENTIALS.Decoder_Credentials(packetsrc)
		IMAPCREDS.Decoder_IMAP_CREDS(packetsrc)
		FRIZZNETWORK.Read_TCP_TO_HTTP(packetsrc)
		FRIZZWIFI.Processor(packetsrc)
	}
	FrizzConvJ.Conver_To_JSON("Modules/Gsrc/Database/Database_Email.json", EMAILSTRUCT, 0600)
	FrizzConvJ.Conver_To_JSON("Modules/Gsrc/Database/Database_Credentials.json", CREDENTIALS, 0600)
	FrizzConvJ.Conver_To_JSON("Modules/Gsrc/Database/Database_IMAP.json", IMAPCREDS, 0600)
	FrizzConvJ.Conver_To_JSON("Modules/Gsrc/Database/Database_HTTP_Creds.json", FRIZZNETWORK, 0600)
	FrizzConvJ.Conver_To_JSON("Modules/Gsrc/Database/Database_WIFI.json", FRIZZWIFI, 0600)
	// once done load file
}
