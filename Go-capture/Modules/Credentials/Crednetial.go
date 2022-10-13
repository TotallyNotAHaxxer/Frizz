package Frizz_Credential

import (
	"regexp"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

/*
Like most of the files that use go to parse packets most of this is payload based

all credentials will be found using regex of USER or PASS especially in SMTP and FTP

the current issue is getting to identify what lists exactly they are and where they were found

because FTP and SMTP have similar sets for recongnizing authentication ( plaintext ) with

USER and PASS parameters, it will be hard to determine it, however the table will be as follows



| 		String 		| 			String 		 |  String  |  String  |  String |     String     |
| Source IP Address | Destination IP Address | Username | Password | Service | Regex verified |
| ----------------- | ---------------------- | -------- | -------- | ------- | -------------- |

This table will help the user identify exactly where the password came from instead of just outputting

exactly where the password came from rather than basing it on the service.

The way we will go about decoding is making sadly enough a few conditionals, the reason conditionals

in this sense are not unsafe but rather imperformant is because of how slow conditionals can be, later on

it will be implimented to use waitgroup and sync channels.

for example

if packet has ethernet
	- If packeet has IPv4
		- If Packet has TCP
				- Parse data and start other functions
*/

// While you can configure services to connect to whatever, the standard will be 21 for SMTP and 25 for SMTP
// We will be comparing the destination port in this case from the decoded TCP packet.
// The way we will decode is decode by each layer and finmd payloads if certain conditions are met
var PortService = map[string]string{
	"21": "FTP",
	"25": "SMTP",
}

var (
	Regex_USER = "(?i)USER (.*)"
	Regex_PASS = "(?i)PASS (.*)"
	Use_USER   = regexp.MustCompile(Regex_USER)
)

type Credentials struct {
	Credentials_Source   string // Source IP
	Credentials_Destin   string // Destination IP
	Credentials_Username string // Username
	Credentials_Password string // Password
	Credentials_Service  string // We should base this on a port
}

func (Cred *Credentials) Decoder_Credentials(packet gopacket.Packet) {
	APPLET := packet.ApplicationLayer()
	if L0 := packet.Layer(layers.LayerTypeEthernet); L0 != nil {
		//		L0_c, _ := L0.(*layers.Ethernet)
		if L1 := packet.Layer(layers.LayerTypeIPv4); L1 != nil {
			//, _ := L1.(*layers.IPv4)
			if L2 := packet.Layer(layers.LayerTypeTCP); L2 != nil {
				L2_c, _ := L2.(*layers.TCP)
				switch L2_c.NextLayerType().String() {
				case "Payload":
					if APPLET != nil {
						Use_c := Use_USER.FindAllStringSubmatch(string(APPLET.Payload()), 1)
						for k := range Use_c {
							Cred.Credentials_Password = string(Use_c[k][0])
						}
					}
				}
			}
		}
	}
}
