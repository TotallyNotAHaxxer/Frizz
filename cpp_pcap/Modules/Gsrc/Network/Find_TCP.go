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

/*
The Website and server will be hosting Useragents, HTTP data, bodies and requests very differently.


So we have the following table layour

|---------------------------------------------------------------------------------------------------------------------------------------|
| Frizz | <button> Dashboard																											|
|       | ------------------------------------------------------------------------------------------------------------------------------|
| t     |
| o     |
| p     | 	  ++++++
| i     |     +    + ( Pie chart showing uagent details)   [  |====|
| c     |     ++++++                                       [  |====| |====|
| s     |      											   [  |====| |====| |====|
| of    |                                                  [  |====| |====| |====| ( bar chart showing all user agent operating systems)
| d     |												   [___________________________
| a     |....................................................................................................................................... ( HR )
| s     |
| h     |   { Table }
| b     |		| Hostname information | Useragent full | Useragent type and OS |
| o     |       |----------------------|----------------|-----------------------|
| a     |       | ...                  | .....          | ......                |
| r     |       |::::::::::::::::::::::|::::::::::::::::|:::::::::::::::::::::::|
| d     |
|-------|===============================================================================================================================================


This design is going to be for most of the HTTP information such as host to NTLM, BASIC, DIGEST and other authentication forms. It would be smart to store them in their own

individual seperators or structures

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
						if l := line.Host; l != "" {
							DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames, l)
							DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames)
							if line.UserAgent() != "" { // if host AND useragent is not empty then append them to the host and user agent list var
								// what we need to do here is a bit weird, we can not assume that the length of the array is going to be the amount of hosts.
								// So solution: Count all the possible hosts and useragents in an array. Then index each array by the number of hosts in the length appended
								// to it, this makes it a bit more simpler upon generation. So lets get to it
								// The reason we need a host is because no request should be made without some form of user agent as it is a massive part of a request.
								/*

								 Keep in mind there is a reason we are using different type structures for this, in this function we are not just looking for a host
								 but looking for a host with a valid user agent within the same line and request feed. So this means we need to create seperate functions to parse
								*/
								DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Uagent = append(DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Uagent, line.UserAgent())
								DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Uagent = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Uagent) // Remove and re append data
								DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Host = append(DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Host, l)
								DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Host = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.Uagent_Host_Host) // remove duplicates
								// then we can use regex to parse the rest of the data if we wanted to, for testing just indexing
								// the code below is a perfect example of how we will use the length to calculate the verified useragents.
								// This code takes the length of the HTTP host names and uses its own length under a for loop to index the user agents with possible hosts that have
								// the same exact user agent. This makes it better because keep in mind we are just parsing valid and used user agents not just all random and possible
								// user agents possible. This logic can be improved so much but for beta it works.
								/*
									var sizeof int
									sizeof = len(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames)
									if sizeof > 0 {
										for i := 0; i < sizeof; i++ {
											if i == len(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents) {
												fmt.Println("{+] Breaking because I is the same length as arrayx -> ", len(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents))
												fmt.Println("sizeof = ", sizeof)
												fmt.Println("uage = ", len(DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents))
												os.Exit(0)
											} else {
												fmt.Println("[+] ====== \033[31m" + DatabaseVar.DatabaseVariable.HTTPD.HTTP_Hostnames[i] + " ---- \033[32m" + DatabaseVar.DatabaseVariable.HTTPD.HTTP_Useragents[i])
											}
										}
									}
								*/
							}
						}
						// url
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
								DatabaseVar.DatabaseVariable.CredentialsHTTPNTLM += 1
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
									DatabaseVar.DatabaseVariable.CredentialsHTTPBASIC += 1
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
								DatabaseVar.DatabaseVariable.CredentialsHTTPDIGEST += 1

							}
						}
						// Detect HTTP Negotiate Authentication
						regf3 := regexp.MustCompile("(?i)Authorization: Negotiate (.*)")
						dtf3 := regf3.FindAllStringSubmatch(string(tcp.Payload), 1)
						if dtf3 != nil {
							for p := range dtf3 {
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE = append(DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE, string(strings.Trim(dtf3[p][0], "Authorization: Negotiate ")))
								DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE = Frizz_Helper.ValueRemover(DatabaseVar.DatabaseVariable.HTTPD.HTTP_NEGOTIATE)
								DatabaseVar.DatabaseVariable.CredentialsHTTPNEG += 1
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
