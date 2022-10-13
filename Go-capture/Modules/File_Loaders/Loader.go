package Frizz_Loader

import (
	"fmt"
	"os"

	FrizzCred "main/Modules/Credentials"
	FrizzPOP3 "main/Modules/Email"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	EMAILSTRUCT FrizzPOP3.Email_Information
	CREDENTIALS FrizzCred.Credentials
)

func Reader(pcapf string) {
	handler, x := pcap.OpenOffline(pcapf)
	if x != nil {
		fmt.Println("[-] ERROR: Could not process file")
		os.Exit(0)
	}
	defer handler.Close()
	packetsrc := gopacket.NewPacketSource(handler, handler.LinkType())
	for packetsrc := range packetsrc.Packets() {
		EMAILSTRUCT.Match_Email_Information(packetsrc)
		CREDENTIALS.Decoder_Credentials(packetsrc)
	}
}
