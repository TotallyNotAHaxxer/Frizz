package Frizz_Net

import (
	"bufio"
	"bytes"
	"net/http"
	"regexp"
	"strings"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"

	Frizz_Helper "main/Modules/Gsrc/Helpers"
)

type HTTP_DATA struct {
	HTTP_Useragents    []string
	HTTP_Hostnames     []string
	HTTP_URLS          []string
	HTTP_NTLM_Encoded  []string
	HTTP_BASIC_ENCODED []string
	HTTP_BASIC_DECODED []string
	HTTP_DIGEST        []string
	HTTP_NEGOTIATE     []string
}

/*

- HTTP Auth forms
	* NTLM
	* Negotiate
	* Digest
	* Basic
*/

func (Netw *HTTP_DATA) Read_TCP_TO_HTTP(packet gopacket.Packet) {
	if lay := packet.Layer(layers.LayerTypeTCP); lay != nil {
		if tcp := lay.(*layers.TCP); tcp != nil {
			if len(tcp.Payload) != 0 {
				r := bufio.NewReader(bytes.NewReader(tcp.Payload))
				line, x := http.ReadRequest(r)
				if x == nil {
					switch line.Proto {
					case "HTTP/1.0", "HTTP/1.1", "HTTP/2", "HTTP/2.0", "HTTP/3", "HTTP/3.0", "HTTP/4.0": // Supporting all HTTP protocol types
						if line.UserAgent() != "" {
							Netw.HTTP_Useragents = append(Netw.HTTP_Useragents, line.UserAgent())
							Netw.HTTP_Useragents = Frizz_Helper.ValueRemover(Netw.HTTP_Useragents) // Remove and re append data
						}
						if l := line.Host; l != "" {
							Netw.HTTP_Hostnames = append(Netw.HTTP_Hostnames, l)
							Netw.HTTP_Hostnames = Frizz_Helper.ValueRemover(Netw.HTTP_Hostnames)
						}
						if u := line.URL; u != nil {
							if h := line.Host; h != "" {
								Netw.HTTP_URLS = append(Netw.HTTP_URLS, "http://"+h+u.String())
								Netw.HTTP_URLS = Frizz_Helper.ValueRemover(Netw.HTTP_URLS)
							}
						}
						// Detect NTLM hash
						reg := regexp.MustCompile("(?i)Authorization: NTLM (.*)")
						dt := reg.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dt != nil {
							for p := range dt {
								Netw.HTTP_NTLM_Encoded = append(Netw.HTTP_NTLM_Encoded, strings.Trim(dt[p][0], "Authorization: NTLM"))
								Netw.HTTP_NTLM_Encoded = Frizz_Helper.ValueRemover(Netw.HTTP_NTLM_Encoded)
							}
						}
						// Detect BASIC authentication | Decode and encoded
						regf := regexp.MustCompile("(?i)Authorization: Basic (.*)")
						dtf := regf.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf != nil {
							for p := range dtf {
								if Frizz_Helper.VALB64(string(strings.Trim(dtf[p][0], "Authorization: Basic "))) {
									Frizz_Helper.DECB64((strings.Trim(dtf[p][0], "Authorization: Basic ")))
									Netw.HTTP_BASIC_DECODED = append(Netw.HTTP_BASIC_DECODED, Frizz_Helper.DECB64((strings.Trim(dtf[p][0], "Authorization: Basic "))))
									Netw.HTTP_BASIC_DECODED = Frizz_Helper.ValueRemover(Netw.HTTP_BASIC_DECODED)
									Netw.HTTP_BASIC_ENCODED = append(Netw.HTTP_BASIC_ENCODED, strings.Trim(dtf[p][0], "Authorization: Basic "))
									Netw.HTTP_BASIC_ENCODED = Frizz_Helper.ValueRemover(Netw.HTTP_BASIC_ENCODED)
								}
							}
						}

						// Detect HTTP Digest Authentication
						regf2 := regexp.MustCompile("(?i)Authorization: Digest (.*)")
						dtf2 := regf2.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf2 != nil {
							for p := range dtf2 {
								Netw.HTTP_DIGEST = append(Netw.HTTP_DIGEST, string(strings.Trim(dtf2[p][0], "Authorization: Digest ")))
								Netw.HTTP_DIGEST = Frizz_Helper.ValueRemover(Netw.HTTP_DIGEST)
							}
						}
						// Detect HTTP Negotiate Authentication
						regf3 := regexp.MustCompile("(?i)Authorization: Negotiate (.*)")
						dtf3 := regf3.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf3 != nil {
							for p := range dtf3 {
								Netw.HTTP_NEGOTIATE = append(Netw.HTTP_NEGOTIATE, string(strings.Trim(dtf3[p][0], "Authorization: Negotiate ")))
								Netw.HTTP_NEGOTIATE = Frizz_Helper.ValueRemover(Netw.HTTP_NEGOTIATE)
							}
						}
					}
				}
			}
		}
	}
}
