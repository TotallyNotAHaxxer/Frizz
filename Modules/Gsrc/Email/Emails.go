package Frizz_Email

import (
	"regexp"
	"strings"

	Frizz_Helpers "main/Modules/Gsrc/Helpers"
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"github.com/google/gopacket"
	ll "github.com/google/gopacket/layers"
)

var (
	Email_OK   = "(?i)OK mailbox (.*)"
	Email_CC   = "(?i)Cc:(.*)"
	Email_FROM = "(?i)From: (.*)"
	Email_RECV = "(?i)Received: (.*)"
)

/*

Frizz_Helpers.

HTML TABLE LAYOUT OF EMAIL INFORMATION EXAMPLE


|  String	   |     String      |     Boolean    |
| ............ | ............... | .............. |
| Email addr   | Payload matched | Regex verified |
| ------------ | --------------- | -------------- |


Email from will be ALL emails found in fields
		- Email_Mailbox
		- Email_CC
		- Email_FROM

Payload matched will be the payload or regex that was used to find the value which will be included
in the general text.

Regex verified is a boolean statement, this will tell the user if this is a verified email using
regular expressions, NOT THE ONE USED TO MATCH THE PAYLOAD but the one to VERIFY the email


Note while frizz does use some real patterns and uses raw decoding, for protocols that are

NOT supported by the third party packages used it uses pattern scanning and matching to determine

what value is what, using regex we can find things such as emails and credentials. We will also use

regex to verify the end string. Emails such as email conversations which are beta features will not be

fully implimented and will be payload based, which are held in maps.
*/

// The most imperformant thing a person could do
var List = []string{
	"hey",
	"hello",
	"email",
	"to",
	"from",
	"data",
	"thank",
	"thank you",
	"thanks",
	"email",
	"cc",
	"Cc:"}

var Possible_Ports = map[string]string{
	"110(pop3)": "POP",
	"995(pop4)": "POP"}

func Match_Email_Information(pkt gopacket.Packet) {
	if lay := pkt.Layer(ll.LayerTypeTCP); lay != nil {
		layerdst := lay.(*ll.TCP)
		APPLICATION := pkt.ApplicationLayer()
		if APPLICATION != nil {
			if Possible_Ports[layerdst.DstPort.String()] == "POP" || Possible_Ports[layerdst.SrcPort.String()] == "POP" {
				MailBoxReg := regexp.MustCompile(Email_OK)
				CcMessageReg := regexp.MustCompile(Email_CC)
				FrMessageReg := regexp.MustCompile(Email_FROM)
				RecMessageReg := regexp.MustCompile(Email_RECV)
				CcMessageReg_c := CcMessageReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
				FrMessageReg_c := FrMessageReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
				RecMessageReg_c := RecMessageReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
				MailBoxReg_c := MailBoxReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
				for o := range MailBoxReg_c {

					DatabaseVar.DatabaseVariable.Email_Info.Email_Mailbox = append(DatabaseVar.DatabaseVariable.Email_Info.Email_Mailbox, strings.Trim(MailBoxReg_c[o][0], `["OK mailbox`))
				}
				DatabaseVar.DatabaseVariable.Email_Info.Email_Mailbox = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_Mailbox)
				for l := range CcMessageReg_c {
					DatabaseVar.DatabaseVariable.Email_Info.Email_CC = append(DatabaseVar.DatabaseVariable.Email_Info.Email_CC, strings.Trim(CcMessageReg_c[l][0], `CC:x`))
					DatabaseVar.DatabaseVariable.Email_Info.Email_CC = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_CC)
				}
				/*

					KEEP OUT FOR DEVELOPMENT


						if strings.Contains(string(pkt.ApplicationLayer().Payload()), "==") {
							if Frizz_Helpers.VALB64(string(pkt.ApplicationLayer().Payload())) {
								DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_RAW = append(DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_RAW, string(pkt.ApplicationLayer().Payload()))
								DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_DEC = append(DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_DEC, Frizz_Helpers.DECB64(string(pkt.ApplicationLayer().Payload())))
							}
							DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_DEC = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_DEC)
							DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_RAW = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_AUTH_RAW)
						}
				*/
				for p := range FrMessageReg_c {
					DatabaseVar.DatabaseVariable.Email_Info.Email_FROM = append(DatabaseVar.DatabaseVariable.Email_Info.Email_FROM, strings.Trim(FrMessageReg_c[p][0], "From:[="))
					DatabaseVar.DatabaseVariable.Email_Info.Email_FROM = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_FROM)
				}
				for o := range RecMessageReg_c {
					DatabaseVar.DatabaseVariable.Email_Info.Email_Recieved = append(DatabaseVar.DatabaseVariable.Email_Info.Email_Recieved, strings.Trim(RecMessageReg_c[o][0], "Received: from"))
					DatabaseVar.DatabaseVariable.Email_Info.Email_Recieved = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_Recieved)
				}

				for _, k := range List {
					if strings.Contains(string(APPLICATION.Payload()), k) {
						DatabaseVar.DatabaseVariable.Email_Info.Email_Session = append(DatabaseVar.DatabaseVariable.Email_Info.Email_Session, string(APPLICATION.Payload()))
						DatabaseVar.DatabaseVariable.Email_Info.Email_Session = Frizz_Helpers.ValueRemover(DatabaseVar.DatabaseVariable.Email_Info.Email_Session)
					}
				}
			}
		}
	}
}
