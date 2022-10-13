package Frizz_Email

import (
	"encoding/base64"
	"regexp"
	"strings"

	"github.com/google/gopacket"
	ll "github.com/google/gopacket/layers"
)

var (
	Email_OK   = "(?i)OK mailbox (.*)"
	Email_CC   = "(?i)Cc:(.*)"
	Email_FROM = "(?i)From: (.*)"
	Email_RECV = "(?i)Received: (.*)"
)

type Email_Information struct {
	Email_Mailbox  []string // Emails found in mailbox
	Email_CC       []string // Emails found in CC
	Email_AUTH_RAW []string // Raw Base 64 authentication
	Email_AUTH_DEC []string // Decoded Base 64 authentication
	Email_FROM     []string // Parses any emails and data within the from response
	Email_Recieved []string // Parses any emails or data that matched the Email_RECV payload
	Email_Session  []string // Any email conversations or entire payloads
}

/*
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

func ValueRemover(cl []string) []string {
	k := make(map[string]bool)
	l := []string{}
	for _, y := range cl {
		if _, u := k[y]; !u {
			k[y] = true
			l = append(l, y)
		}
	}
	return l
}

func VALB64(s string) bool {
	_, x := base64.StdEncoding.DecodeString(s)
	return x == nil
}

func DECB64(s string) string {
	f, _ := base64.StdEncoding.DecodeString(s)
	return string(f)
}

func (P3 *Email_Information) Match_Email_Information(pkt gopacket.Packet) {
	if lay := pkt.Layer(ll.LayerTypeTCP); lay != nil {
		APPLICATION := pkt.ApplicationLayer()
		if APPLICATION != nil {
			// Check for email or mailbox messages
			MailBoxReg := regexp.MustCompile(Email_OK)
			CcMessageReg := regexp.MustCompile(Email_CC)
			FrMessageReg := regexp.MustCompile(Email_FROM)
			RecMessageReg := regexp.MustCompile(Email_RECV)
			CcMessageReg_c := CcMessageReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
			FrMessageReg_c := FrMessageReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
			RecMessageReg_c := RecMessageReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
			MailBoxReg_c := MailBoxReg.FindAllStringSubmatch(string(APPLICATION.Payload()), 1)
			for o := range MailBoxReg_c {
				P3.Email_Mailbox = append(P3.Email_Mailbox, strings.Trim(MailBoxReg_c[o][0], `["OK mailbox`))
			}
			P3.Email_Mailbox = ValueRemover(P3.Email_Mailbox)
			for l := range CcMessageReg_c {
				P3.Email_CC = append(P3.Email_CC, strings.Trim(CcMessageReg_c[l][0], `CC:x`))
				P3.Email_CC = ValueRemover(P3.Email_CC)
			}
			if strings.Contains(string(pkt.ApplicationLayer().Payload()), "==") {
				if VALB64(string(pkt.ApplicationLayer().Payload())) {
					P3.Email_AUTH_RAW = append(P3.Email_AUTH_RAW, string(pkt.ApplicationLayer().Payload()))
					P3.Email_AUTH_DEC = append(P3.Email_AUTH_DEC, DECB64(string(pkt.ApplicationLayer().Payload())))
				}
				P3.Email_AUTH_DEC = ValueRemover(P3.Email_AUTH_DEC)
				P3.Email_AUTH_RAW = ValueRemover(P3.Email_AUTH_RAW)
			}
			for p := range FrMessageReg_c {
				P3.Email_FROM = append(P3.Email_FROM, strings.Trim(FrMessageReg_c[p][0], "From:[="))
				P3.Email_FROM = ValueRemover(P3.Email_FROM)
			}
			for o := range RecMessageReg_c {
				P3.Email_Recieved = append(P3.Email_Recieved, strings.Trim(RecMessageReg_c[o][0], "Received: from"))
				P3.Email_Recieved = ValueRemover(P3.Email_Recieved)
			}

			for _, k := range List {
				if strings.Contains(string(APPLICATION.Payload()), k) {
					P3.Email_Session = append(P3.Email_Session, string(APPLICATION.Payload()))
					P3.Email_Session = ValueRemover(P3.Email_Session)
				}
			}
		}
	}
}
