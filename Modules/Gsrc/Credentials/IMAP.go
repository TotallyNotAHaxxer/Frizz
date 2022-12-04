package Frizz_Credential

import (
	Frizz_Helper "main/Modules/Gsrc/Helpers"
	DatabaseVar "main/Modules/Gsrc/TypeVar"

	"regexp"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

//DIGEST MD5 EXAMPLE {MD5} -> dXNlcm5hbWU9ImhlbW1pbmd3YXkiLHJlYWxtPSIiLG5vbmNlPSJoWmE1eml1UzBFUU9uZlp2QThDNjBnPT0iLGNub25jZT0iNGhnb0ZpZEk5RHc9IixuYz0wMDAwMDAwMSxxb3A9YXV0aCxkaWdlc3QtdXJpPSJpbWFwLzEwLjAuMS4xMDIiLHJlc3BvbnNlPTRhMGRhMTYzY2U1NjcyODdmZmE3NzdiYTJlZmJjOTdlLGNoYXJzZXQ9dXRmLTg=
//IMAP PLAIN AUTH    {B64} -> AUTH PLAIN AGRpZ2l0YWxpbnZlc3RpZ2F0b3JAbmV0d29ya3NpbXMuY29tAG5hcGllcjEyMw==
//IMAP PLAIN TEXT    {TXT} -> LOGIN neulingern XXXXXX
// I seriously... hate the linear programming here................................................................	..

/*

Like most layers in this list

given the fact that we are parsing the raw payload, we want to avoid loading the data types with junk

so we detect the layer port by map, if the map returns a IMAP message then we are good to go and parse the rest
*/
var IMAP_PORTS = map[string]string{
	"143(imap)": "IMAP",
	"993(imap)": "IMAP",
}

// REGEXP
var (
	msg                  string
	msg2                 string
	PLAIN_LOGIN_IMAP_MSG = "(?i)LOGIN(.*)"
	USE_PLAIN_LOGIN      = regexp.MustCompile(PLAIN_LOGIN_IMAP_MSG)
	// Plain authentication B64
	PLAIN_LOGIN_IMAP_BASE64 = "(?i)2 authenticate plain(.*)"
	USE_PLAIN_LOGIN_BASE64  = regexp.MustCompile(PLAIN_LOGIN_IMAP_BASE64)
	// DIGEST-MD5 	authentication
	DIGEST_LOGIN_IMAP_DIGEST_MD5 = "(?i)AUTHENTICATE DIGEST-MD5(.*)"
	USE_DIGEST_LOGIN_DIGEST_MD5  = regexp.MustCompile(DIGEST_LOGIN_IMAP_DIGEST_MD5)
)

func Decoder_IMAP_CREDS(pkt gopacket.Packet) {
	if lay0 := pkt.Layer(layers.LayerTypeEthernet); lay0 != nil {
		if lay1 := pkt.Layer(layers.LayerTypeIPv4); lay1 != nil {
			if lay2 := pkt.Layer(layers.LayerTypeTCP); lay2 != nil {
				lay2_c := lay2.(*layers.TCP)
				if pkt.ApplicationLayer() != nil {
					if IMAP_PORTS[lay2_c.DstPort.String()] == "IMAP" {
						result_plain := USE_PLAIN_LOGIN.FindAllStringSubmatch(string(pkt.ApplicationLayer().Payload()), 1)
						if result_plain != nil {
							// Because this is plain text authentication we do NOT need base64 encoding testing
							for i := range result_plain {
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_PLAINTEXT = append(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_PLAINTEXT, strings.Trim(result_plain[i][0], `" LOGIN `))
								DatabaseVar.DatabaseVariable.CredentialsIMAP += 1
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_PLAINTEXT = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_PLAINTEXT)
							}
						}

						// This means that it is plain Base64 authentication
						// now we need to implimnet logic to predict the next packet
						result_plain_b64 := USE_PLAIN_LOGIN_BASE64.FindAllStringSubmatch(string(pkt.ApplicationLayer().Payload()), 1)
						if msg == "next" {
							a := Frizz_Helper.VALB64(string(pkt.ApplicationLayer().Payload()))
							if a {
								dec := Frizz_Helper.DECB64(string(pkt.ApplicationLayer().Payload()))
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Decoded = append(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Decoded, dec)
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Decoded = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Decoded)
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Encoded = append(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Encoded, string(pkt.ApplicationLayer().Payload()))
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Encoded = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_BASE64_Encoded)
								msg = "mode"
								DatabaseVar.DatabaseVariable.CredentialsIMAP += 1
							}
						}
						if result_plain_b64 != nil {
							msg = "next"

						}
						// This next one means that it is DIGEST MD5
						// More prediction is required for this to work
						/*

							since we are using the if statements and variables to assume the next IMAP packet will be an authentication key
							we need to put the if statements above their checks, if you put the if statement that checks the variable below
							after the variable is set, you will not get the intended result

						*/
						result_digest_md5 := USE_DIGEST_LOGIN_DIGEST_MD5.FindAllStringSubmatch(string(pkt.ApplicationLayer().Payload()), 1)
						if msg2 == "next" {
							if Frizz_Helper.VALB64(string(pkt.ApplicationLayer().Payload())) {
								dec := Frizz_Helper.DECB64(string(pkt.ApplicationLayer().Payload()))
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Decoded = append(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Decoded, dec)
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Encoded = append(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Encoded, string(pkt.ApplicationLayer().Payload()))
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Encoded = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Encoded)
								DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Decoded = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.IMAPCREDS.IMAP_DIGEST_MD5_Decoded)
								DatabaseVar.DatabaseVariable.CredentialsIMAP += 1
							}
							msg2 = "0"
						}
						if result_digest_md5 != nil {
							msg2 = "next"
						}
					}

				}
			}
		}
	}
}

var a string
