package Frizz_Net

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	Frizz_Helper "main/Modules/Gsrc/Helpers"
	DatabaseVar "main/Modules/Gsrc/TypeVar"
)

/*

- HTTP Auth forms
	* NTLM
	* Negotiate
	* Digest
	* Basic
*/

func Read_TCP_TO_HTTP(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp := lay.(*layers.TCP); tcp != nil {
			if len(tcp.Payload) != 0 {
				r := bufio.NewReader(bytes.NewReader(tcp.Payload))
				line, x := http.ReadRequest(r)
				if x == nil {
					switch line.Proto {
					case "HTTP/1.0", "HTTP/1.1", "HTTP/2", "HTTP/2.0", "HTTP/3", "HTTP/3.0", "HTTP/4.0": // Supporting all HTTP protocol types
						if line.UserAgent() != "" {
							DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents, line.UserAgent())
							DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents) // Remove and re append data
						}
						if l := line.Host; l != "" {
							DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames, l)
							DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames)
						}
						if u := line.URL; u != nil {
							if h := line.Host; h != "" {
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_URLS = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_URLS, "http://"+h+u.String())
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_URLS = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_URLS)
							}
						}
						// Detect NTLM hash
						reg := regexp.MustCompile("(?i)Authorization: NTLM (.*)")
						dt := reg.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dt != nil {
							for p := range dt {
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_NTLM_Encoded = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_NTLM_Encoded, strings.Trim(dt[p][0], "Authorization: NTLM"))
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_NTLM_Encoded = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_NTLM_Encoded)
							}
						}
						// Detect BASIC authentication | Decode and encoded
						regf := regexp.MustCompile("(?i)Authorization: Basic (.*)")
						dtf := regf.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf != nil {
							for p := range dtf {
								if Frizz_Helper.VALB64(string(strings.Trim(dtf[p][0], "Authorization: Basic "))) {
									Frizz_Helper.DECB64((strings.Trim(dtf[p][0], "Authorization: Basic ")))
									DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_DECODED = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_DECODED, Frizz_Helper.DECB64((strings.Trim(dtf[p][0], "Authorization: Basic "))))
									DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_DECODED = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_DECODED)
									DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_ENCODED = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_ENCODED, strings.Trim(dtf[p][0], "Authorization: Basic "))
									DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_ENCODED = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_BASIC_ENCODED)
								}
							}
						}

						// Detect HTTP Digest Authentication
						regf2 := regexp.MustCompile("(?i)Authorization: Digest (.*)")
						dtf2 := regf2.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf2 != nil {
							for p := range dtf2 {
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_DIGEST = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_DIGEST, string(strings.Trim(dtf2[p][0], "Authorization: Digest ")))
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_DIGEST = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_DIGEST)
							}
						}
						// Detect HTTP Negotiate Authentication
						regf3 := regexp.MustCompile("(?i)Authorization: Negotiate (.*)")
						dtf3 := regf3.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf3 != nil {
							for p := range dtf3 {
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE, string(strings.Trim(dtf3[p][0], "Authorization: Negotiate ")))
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE)
							}
						}

						// Append HTTP layer data
						var data string

						if line.Header != nil {
							data += string(fmt.Sprint(line.Header))
						}

						if line.ContentLength != 0 {
							data += string(fmt.Sprint(line.ContentLength))
						}
						if line.Method != "" {
							data += string(fmt.Sprint(line.Method))
						}

						if line.URL != nil {
							data += string(fmt.Sprint(line.URL))
						}

						if line.ProtoMajor != 0 {
							data += string(fmt.Sprint(line.ProtoMajor))
						}

						if line.RequestURI != "" {
							data += string(fmt.Sprint(line.RequestURI))
						}

						DatabaseVar.DatabaseVariable.HTTPD.HTTP_FULL_SESSION_DATA = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_FULL_SESSION_DATA, data)
					}
				}
			}
		}
	}
}
