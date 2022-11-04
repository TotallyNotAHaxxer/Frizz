package Frizz_Credential

import (
	"regexp"
	"strings"

	Frizz_Helpers "main/Modules/Gsrc/Helpers"
	Frizz_Data "main/Modules/Gsrc/TypeVar"

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

Frizz currently supports only the following credential types for SMTP and FTP:
	- Plain Auth
	- CRAM MD5 (Will decode some of the base64 but not differenciate the MD5 hash)

This method of logic is very very long and very slow, I plan to implement something more sophisticated here and will make this
way more efficient and faster than current

*/

// While you can configure services to connect to whatever, the standard will be 21 for SMTP and 25 for SMTP
// We will be comparing the destination port in this case from the decoded TCP packet.
// The way we will decode is decode by each layer and finmd payloads if certain conditions are met
var PortService = map[string]string{
	"21(ftp)":  "FTP",
	"25(smtp)": "SMTP",
}

/*

In a modern SMTP connection, some settings are where the server will base 64 encode the

responses, we need to be able to verify which one is the input and which one is the output

since we know it did not match the regex payload we can go to verify the string and check

if it is base64, if it is base64 then decode it, if the decoded text is equal to the map value

which we can count as a login form value then it will be set as the title



*/

var InputMapsSMTP = map[string]string{
	"Password:": "Login Form Pass",
	"Username:": "Login Form User",
	"�":         "nah",
}

var (
	Regex_USER_FTP                                       = "(?i)USER (.*)"
	Regex_PASS_FTP                                       = "(?i)PASS (.*)"
	Regex_PLAIN_SMTP                                     = "(?i)AUTH PLAIN (.*)"
	Regex_CRAM5_SMTP                                     = "(?i)AUTH CRAM-MD5(.*)"
	Regex_AUTH2                                          = "(?i)AUTH LOGIN\r\n(.*)"
	Use_USER_FTP                                         = regexp.MustCompile(Regex_USER_FTP)
	Use_PASS_FTP                                         = regexp.MustCompile(Regex_PASS_FTP)
	Use_PLAIN_SMTP                                       = regexp.MustCompile(Regex_PLAIN_SMTP)
	Use_AUTH2                                            = regexp.MustCompile(Regex_AUTH2)
	Use_CRAM5_SMTP                                       = regexp.MustCompile(Regex_CRAM5_SMTP)
	Table_Str_SMTP                                       string
	Table_Str_FTP                                        string
	NextPacket                                           string
	Prediction_Username_Plain_SMTP                       bool
	Prediction_Password_Plain_SMTP                       bool
	Prediction_Detector_Is_Next_Packet_Username_Password string
)

func Decoder_Credentials(packet gopacket.Packet) {
	APPLET := packet.ApplicationLayer()
	if L0 := packet.Layer(layers.LayerTypeEthernet); L0 != nil {
		if L1 := packet.Layer(layers.LayerTypeIPv4); L1 != nil {
			//L1_c, _ := L1.(*layers.IPv4)
			if L2 := packet.Layer(layers.LayerTypeTCP); L2 != nil {
				L2_c, _ := L2.(*layers.TCP)
				switch L2_c.NextLayerType().String() {
				case "Payload":
					if PortService[L2_c.DstPort.String()] == "FTP" {
						if APPLET != nil {
							Use_c := Use_USER_FTP.FindAllStringSubmatch(string(APPLET.Payload()), 1)
							Use_c0 := Use_PASS_FTP.FindAllStringSubmatch(string(APPLET.Payload()), 1)
							for i := range Use_c {
								Frizz_Data.DatabaseVariable.Creds.Ftp.FTP_Username = append(Frizz_Data.DatabaseVariable.Creds.Ftp.FTP_Username, strings.Trim(Use_c[i][1], "USER"))
							}
							for i := range Use_c0 {
								Frizz_Data.DatabaseVariable.Creds.Ftp.FTP_Password = append(Frizz_Data.DatabaseVariable.Creds.Ftp.FTP_Password, strings.Trim(Use_c0[i][1], "PASS"))
							}
						}
					}

					if PortService[L2_c.DstPort.String()] == "SMTP" {
						if APPLET != nil {
							Use_c1 := Use_PLAIN_SMTP.FindAllStringSubmatch(string(APPLET.Payload()), 1)
							Use_c2 := Use_CRAM5_SMTP.FindAllStringSubmatch(string(APPLET.Payload()), 1)
							if Use_c1 != nil {
								Decoded_Value := Frizz_Helpers.DECB64(strings.Trim(string(APPLET.Payload()), "AUTH PLAIN"))
								Frizz_Data.DatabaseVariable.Creds.Smtp.SMTP_Plainauth = Frizz_Helpers.Appender(Use_c1, Frizz_Data.DatabaseVariable.Creds.Smtp.SMTP_Plainauth)
								Frizz_Data.DatabaseVariable.Creds.Smtp.SMTP_Plainauth_Decodec = append(Frizz_Data.DatabaseVariable.Creds.Smtp.SMTP_Plainauth_Decodec, Decoded_Value)
							}
							if Use_c2 != nil {
								NextPacket = "Finished here!"
							} else {
								if NextPacket == "Finished here!" {
									Decoded_Value := Frizz_Helpers.DECB64(string(APPLET.Payload())[0:])
									Frizz_Data.DatabaseVariable.Creds.Smtp_Cram.Decoded = append(Frizz_Data.DatabaseVariable.Creds.Smtp_Cram.Decoded, Decoded_Value)
									Frizz_Data.DatabaseVariable.Creds.Smtp_Cram.Encoded = append(Frizz_Data.DatabaseVariable.Creds.Smtp_Cram.Encoded, string(APPLET.Payload())[0:])
								}
							}
						}
					}
				}
			}
		}
	}
}
