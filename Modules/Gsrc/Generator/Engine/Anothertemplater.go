package Engine

import (
	"fmt"
	"log"
	Frizzlog "main/Modules/Gsrc/File_Loaders"
	Frizz_Helpers "main/Modules/Gsrc/Helpers"
	FrizzNet "main/Modules/Gsrc/Network"
	FrizzWifu "main/Modules/Gsrc/Wifi"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var datah HTMLData

func LoadBasedOnType(cred string) string {
	var templater string
	switch cred {
	case "ftp":
		templater += `
			<thead>
			<tr>
				<th>
					<h1>Username</h1>
				</th>
				<th>
					<h1>Password</h1>
				</th>
			</tr>
			</thead>
		`
		for k := 0; k < len(StructureFrizzPointer.Creds.Ftp.FTPPassword); k++ {
			templater += "<tr>"
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.Ftp.FTPPassword[k]) + Generate_TDE()
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.Ftp.FTPUsername[k]) + Generate_TDE()
			templater += "</tr>"
		}
	case "smtp":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>Encoded | Login</h1>
			</th>
			<th>
				<h1>Decoded | Password</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for k := 0; k < len(StructureFrizzPointer.Creds.SMTPCram.Decoded); k++ {
			templater += "<tr>"
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.SMTPCram.Encoded[k]) + Generate_TDE()
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.SMTPCram.Decoded[k]) + Generate_TDE()
			templater += Generate_TD() + "CRAM-MD5" + Generate_TDE()
			templater += "</tr>"
		}
		for l := 0; l < len(StructureFrizzPointer.Creds.SMTP.SMTPPlainauth); l++ {
			templater += "<tr>"
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.SMTP.SMTPPlainauth[l]) + Generate_TDE()
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.SMTP.SMTPPlainauthDecodec[l]) + Generate_TDE()
			templater += Generate_TD() + "SMTP PLAIN AUTHENTICATION" + Generate_TDE()
			templater += "</tr>"
		}
		for q := 0; q < len(StructureFrizzPointer.Creds.SMTP.SMTPUsername); q++ {
			templater += "<tr>"
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.SMTP.SMTPUsername[q]) + Generate_TDE()
			templater += Generate_TD() + datah.StripHTML(StructureFrizzPointer.Creds.SMTP.SMTPPassword[q]) + Generate_TDE()
			templater += Generate_TD() + "SMTP PlainText authentication" + Generate_TDE()
			templater += "</tr>"
		}
	case "imap":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>Encoded | Login</h1>
			</th>
			<th>
				<h1>Decoded | Password</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for l := 0; l < len(StructureFrizzPointer.Imapcreds.IMAPBASE64Decoded); l++ {
			templater += "<tr>"
			if StructureFrizzPointer.Imapcreds.IMAPBASE64Encoded[l] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Imapcreds.IMAPBASE64Encoded[l] + Generate_TDE()
			}
			if StructureFrizzPointer.Imapcreds.IMAPBASE64Decoded[l] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Imapcreds.IMAPBASE64Decoded[l] + Generate_TDE()
			}
			templater += Generate_TD() + "IMAP BASE64 ENCODED" + Generate_TDE()
			templater += "</tr>"
		}
		for i := 0; i < len(StructureFrizzPointer.Imapcreds.ImapPlaintext); i++ {
			templater += "<tr>"

			if StructureFrizzPointer.Imapcreds.ImapPlaintext[i] != "" {
				splitter := strings.Split(StructureFrizzPointer.Imapcreds.ImapPlaintext[i], " ")
				for k := range splitter {
					templater += Generate_TD() + splitter[k] + Generate_TDE()
				}
				templater += Generate_TD() + "IMAP PLAIN TEXT" + Generate_TDE()

			} else {
				templater += Generate_TD() + "empty" + Generate_TDE()
				templater += Generate_TD() + "empty" + Generate_TDE()
			}
			templater += "</tr>"
		}
	case "ssh":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>Encoded | Login</h1>
			</th>
			<th>
				<h1>Decoded | Password</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for l := 0; l < len(StructureFrizzPointer.Creds.SSH.SSHPassword); l++ {
			templater += "<tr>"
			if StructureFrizzPointer.Creds.SSH.SSHPassword[l] != "" && StructureFrizzPointer.Creds.SSH.SSHUsername[l] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Creds.SSH.SSHPassword[l] + Generate_TDE()
				templater += Generate_TD() + StructureFrizzPointer.Creds.SSH.SSHUsername[l] + Generate_TDE()
				templater += Generate_TD() + "SSH PLAIN TEXT" + Generate_TDE()
			}
			templater += "</tr>"
		}
	case "httpdigest":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>Decoded Digest</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for k := 0; k < len(StructureFrizzPointer.Httpd.HTTPDigest); k++ {
			templater += "<tr>"
			if StructureFrizzPointer.Httpd.HTTPDigest[k] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Httpd.HTTPDigest[k] + Generate_TDE()
				templater += Generate_TD() + "HTTP MD5-Digest" + Generate_TDE()
			}
			templater += "</tr>"
		}
	case "httpbasic":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>HTTP BASIC | Encoded | Base64</h1>
			</th>
			<th>
				<h1>HTTP BASIC | Decoded | Base64</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for q := 0; q < len(StructureFrizzPointer.Httpd.HTTPBasicDecoded); q++ {
			templater += "<tr>"
			if StructureFrizzPointer.Httpd.HTTPBasicDecoded[q] != "" && StructureFrizzPointer.Httpd.HTTPBasicEncoded[q] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Httpd.HTTPBasicEncoded[q] + Generate_TDE()
				templater += Generate_TD() + StructureFrizzPointer.Httpd.HTTPBasicDecoded[q] + Generate_TDE()
				templater += Generate_TD() + "HTTP BASIC AUTHENTICATION" + Generate_TDE()
			}
			templater += "</tr>"
		}
	case "httpntlm":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>HTTP NTLM ENCODED</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for l := 0; l < len(StructureFrizzPointer.Httpd.HTTPNTLMEncoded); l++ {
			templater += "<tr>"
			if StructureFrizzPointer.Httpd.HTTPNTLMEncoded[l] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Httpd.HTTPNTLMEncoded[l] + Generate_TDE()
				templater += Generate_TD() + "HTTP NTLM" + Generate_TDE()
			}
			templater += "</tr>"

		}
	case "httpnegotiate":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>HTTP Negotiate | Encoded</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		for p := 0; p < len(StructureFrizzPointer.Httpd.HTTPNegotiate); p++ {
			templater += "<tr>"
			if StructureFrizzPointer.Httpd.HTTPNegotiate[p] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Httpd.HTTPNegotiate[p] + Generate_TDE()
				templater += Generate_TD() + "HTTP Negotiate" + Generate_TDE()
			}
			templater += "</tr>"

		}
	case "telnet":
		templater += `
		<thead>
		<tr>
			<th>
				<h1>Username</h1>
			</th>
			<th>
				<h1>Password</h1>
			</th>
			<th>
				<h1>Type</h1>
			</th>
		</tr>
		</thead>
	`
		var body string
		for p := 0; p < len(StructureFrizzPointer.Tels.Body); p++ {
			body += StructureFrizzPointer.Tels.Body[p]
		}
		re_password := regexp.MustCompile("(?i)Password: (.*)")
		re_username := regexp.MustCompile("(?i)login: (.*)")
		Use_c := re_password.FindAllStringSubmatch(string(body), 1)
		Use_c0 := re_username.FindAllStringSubmatch(body, -1)
		var credarr, userarr []string
		for i := range Use_c {
			credarr = append(credarr, Use_c[i][1])
		}
		for k := range Use_c0 {
			userarr = append(userarr, Use_c0[k][1])
		}
		for i := 0; i < len(credarr); i++ {
			templater += "<tr>"
			if credarr[i] != "" {
				templater += Generate_TD() + credarr[i] + Generate_TDE()
				templater += Generate_TD() + userarr[i] + Generate_TDE()
				templater += Generate_TD() + "Telnet" + Generate_TDE()
			}
			templater += "</tr>"
		}
	}

	return templater
}

func LoadCredentials(filename string, credential string) {
	var anotherftemplate string
	anotherftemplate += StandardTop
	anotherftemplate += Standard_LinksList
	anotherftemplate += StaticSection
	anotherftemplate += Generate_Div("home-content")
	anotherftemplate += Generate_Div("overview-boxes")
	anotherftemplate += RetCredType(credential)
	anotherftemplate += "</div><hr><br>"
	anotherftemplate += `<table class="container_Overview"><br><br>`
	anotherftemplate += "<tbody>"
	anotherftemplate += LoadBasedOnType(credential)
	anotherftemplate += `</tbody></table>`
	anotherftemplate += StaticStyle
	anotherftemplate += StaticJS
	Write(filename, anotherftemplate)
}

func LoadSessions(filename, session string) {
	var sessiontemplate string

	Write(filename, sessiontemplate)
}

func RetCredType(credential string) string {
	switch credential {
	case "ftp":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsFTP))
	case "smtp":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsSMTP))
	case "imap":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsIMAP))
	case "ssh":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsSSH))
	case "digest":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsHTTPDIGEST))
	case "basic":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsHTTPBASIC))
	case "ntlm":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsHTTPNTLM))
	case "negotiate":
		return Generate_Box("Total credentials found", fmt.Sprint(StructureFrizzPointer.CredentialsHTTPNEG))
	case "telnet":
		return Generate_Box("Total credentials found", "Unknown | Engine unsupported type")
	default:
		return "empty or error"
	}
}

// SESSION GENERATION | SSH, TELNET, SMTP, FTP

func GenBoxBSessionN(sessionname string) (string, string, string) {
	switch sessionname {
	case "ftp":
		return Generate_Box("Session type", "FTP"), Generate_Box("Session Lists", "1"), Generate_Box("Session protocols", "FTP")
	case "smtp":
		return Generate_Box("Session type", "SMTP"), Generate_Box("Session Lists", "1"), Generate_Box("Session protocols", "SMTP")
	case "ssh":
		return Generate_Box("Session type", "SSH"), Generate_Box("Session Lists", "1"), Generate_Box("Session protocols", "SSH")
	case "telnet":
		return Generate_Box("Session type", "Telnet"), Generate_Box("Session Lists", "1"), Generate_Box("Session protocols", "Telnet")
	default:
		return "empty", "error", "during make"
	}
}

func GenerateTableBasedOnSessionN(sessionname, sessionbody string) string {
	var tmpl string
	tmpl += `<table class="container_Overview"><br><br>`
	tmpl +=
		`
<thead>
	<tr>
		<th>
			<h1>Session Variable</h1>
		</th>
		<th>
			<h1>Session Value</h1>
		</th>
	</tr>
</thead>
`
	tmpl += "<tbody>"
	regpat := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	submatchall := regpat.FindAllString(sessionbody, -1)
	tmpl += "<tr>"
	tmpl += Generate_TD() + "Found Hosts" + Generate_TDE()
	tmpl += Generate_TD()
	submatchall = Frizz_Helpers.ValueRemover(submatchall)
	for _, element := range submatchall {
		if element != "" {
			tmpl += element + "<br>"
		}
	}
	tmpl += Generate_TDE()
	tmpl += "</tr>"
	tmpl += "</tbody></table>"
	tmpl += "<br>"
	tmpl += `<pre class="syntax">`
	switch sessionname {
	case "ssh":
		for o := 0; o < len(StructureFrizzPointer.Sshs.Body); o++ {
			tmpl += StructureFrizzPointer.Sshs.Body[o] // ssh body
		}
	case "ftp":
		for o := 0; o < len(StructureFrizzPointer.Ftps.Data); o++ {
			tmpl += StructureFrizzPointer.Ftps.Data[o] // ftp body
		}
	case "smtp":
		for o := 0; o < len(StructureFrizzPointer.SMTPSessionInf.Body); o++ {
			tmpl += datah.StripHTML(StructureFrizzPointer.SMTPSessionInf.Body[o]) // smtp body
		}
	case "telnet":
		for o := 0; o < len(StructureFrizzPointer.Tels.Body); o++ {
			tmpl += StructureFrizzPointer.Tels.Body[o] // telnet body
		}
	default:
		tmpl += "shit man got an issue"
	}
	tmpl += `</pre>`
	return tmpl

}

func CompareChars(word string) bool {
	var vari bool
	for i, c := range word {
		if i < len(word)-1 {
			if string(word[i+1]) == string(c) {
				vari = true
			} else {
				vari = false
			}
		}
	}
	return vari
}

func LoadSessionTemplates(filename, sessionname string) {
	var sessionbasedtemplate string
	sessionbasedtemplate += StandardTopSessions            // adding top
	sessionbasedtemplate += Standard_LinksList             // Adding links
	sessionbasedtemplate += StaticSection                  // Adding static sections
	sessionbasedtemplate += Generate_Div("home-content")   // Generating custom div using the engine
	sessionbasedtemplate += Generate_Div("overview-boxes") // Generating tag and box preperation using the engine
	b1, b2, b3 := GenBoxBSessionN(sessionname)             // Getting the session and data boxes based on the session type
	sessionbasedtemplate += b1                             // Adding box 1 to the template
	sessionbasedtemplate += b2                             // Adding box 2 to the template
	sessionbasedtemplate += b3                             // Adding box 3 to the template
	sessionbasedtemplate += "</div><hr><br>"               // Adding closing tags to the templates first top section
	var body string
	if sessionname == "ftp" {
		for i := 0; i < len(StructureFrizzPointer.Ftps.Data); i++ {
			body += StructureFrizzPointer.Ftps.Data[i]
		}
	} else if sessionname == "ssh" {
		for i := 0; i < len(StructureFrizzPointer.Sshs.Body); i++ {
			body += StructureFrizzPointer.Sshs.Body[i]
		}
	} else if sessionname == "telnet" {
		for i := 0; i < len(StructureFrizzPointer.Tels.Body); i++ {
			if CompareChars(StructureFrizzPointer.Tels.Body[i]) {
				body += strings.Trim(StructureFrizzPointer.Tels.Body[i], "�")
			}
		}
	} else if sessionname == "smtp" {
		for i := 0; i < len(StructureFrizzPointer.SMTPSessionInf.Body); i++ {
			body += StructureFrizzPointer.SMTPSessionInf.Body[i]
		}
	} else {
		log.Fatal("Un named or based session name")
	}
	sessionbasedtemplate += GenerateTableBasedOnSessionN(sessionname, body)
	//
	sessionbasedtemplate += StaticJS           // loading js
	sessionbasedtemplate += StandardStyleSheet // generating stylesheet

	sessionbasedtemplate += ` 
	<style>
	.syntax {
		background: black;
		color: white;
		margin: 10px;
		margin-top: -135px;
		border: solid thin #333;
		max-width: 1000px;
		width: 100%;
		display: inline-block;
		white-space: pre-wrap;
		
	}

	.syntax span {
		counter-increment: linecounter;
		white-space: pre-wrap;
		
	}

	.syntax span:before {
		content: counter(linecounter);
		width: 1.2em;
		text-align: center;
		display: inline-block;
		border-right: 1px solid #444;
		margin-right: 10px;
		font-style: normal !important;
		color: #444 !important;
		white-space: pre-wrap;
		
	}</style>`
	Write(filename, sessionbasedtemplate)
	/*


	 */
}

// Writer for email / POP information | Cc, From, To, Subject, Message, RECV
func DrawByType(payloadt string) (string, string, string) {
	switch payloadt {
	case "cc":
		return Generate_Box("Payload", "Cc"), Generate_Box("Payload Regex", "1"), Generate_Box("Session protocols", "POP3")
	case "from":
		return Generate_Box("Payload", "From"), Generate_Box("Payload Regex", "1"), Generate_Box("Session protocols", "POP3")
	case "recv":
		return Generate_Box("Payload", "Recieve"), Generate_Box("Payload Regex", "1"), Generate_Box("Session protocols", "POP3")
	case "conversation":
		return Generate_Box("Payload", "None | *"), Generate_Box("Payload Regex", "1"), Generate_Box("Session protocols", "POP3")
	case "*em":
		return Generate_Box("Payload", "&&*&&"), Generate_Box("Payload Regex", "5"), Generate_Box("Session protocols", "POP, SMTP, IMAP")
	default:
		return Generate_Box("Payload", "Unknown"), Generate_Box("Payload Regex", "Unknown"), Generate_Box("Session protocols", "POP3")
	}
}

func DrawTableByPayload(payload string) (bodytempl string) {
	switch payload {
	case "cc":
		bodytempl +=
			`
		<thead>
			<tr>
				<th>
					<h1>Email</h1>
				</th>
				<th>
					<h1>Payload type</h1>
				</th>
			</tr>
		</thead>
		`
		bodytempl += "<tbody>"
		for i := 0; i < len(StructureFrizzPointer.EmailInfo.EmailCC); i++ {
			bodytempl += "<tr>"
			bodytempl += Generate_TD() + StructureFrizzPointer.EmailInfo.EmailCC[i] + Generate_TDE()
			bodytempl += Generate_TD() + "Found using 'Cc' within regex engine" + Generate_TDE()
			bodytempl += "</tr>"
		}
	case "from":
		bodytempl += `
			<thead>
			<tr>
				<th>
					<h1>Email</h1>
				</th>
				<th>
					<h1>Payload type</h1>
				</th>
			</tr>
		</thead>
		`
		bodytempl += "<tbody>"
		for k := 0; k < len(StructureFrizzPointer.EmailInfo.EmailFROM); k++ {
			bodytempl += "<tr>"
			bodytempl += Generate_TD() + StructureFrizzPointer.EmailInfo.EmailFROM[k] + Generate_TDE()
			bodytempl += Generate_TD() + "Found using the 'FROM' payload in regex engine" + Generate_TDE()
			bodytempl += "</tr>"
		}
	case "recv":
		bodytempl += `
		<thead>
		<tr>
			<th>
				<h1>Email</h1>
			</th>
			<th>
				<h1>Payload type</h1>
			</th>
		</tr>
	</thead>
	`
		bodytempl += "<tbody>"
		for l := 0; l < len(StructureFrizzPointer.EmailInfo.EmailRecieved); l++ {
			bodytempl += "<tr>"
			bodytempl += Generate_TD() + StructureFrizzPointer.EmailInfo.EmailRecieved[l] + Generate_TDE()
			bodytempl += Generate_TD() + "Found using 'RECV' within the engine" + Generate_TDE()
			bodytempl += "</tr>"
		}
	case "conversation":
		bodytempl += `
		<thead>
		<tr>
			<th>
				<h1>Value</h1>
			</th>
			<th>
				<h1>Payload</h1>
			</th>
		</tr>
	</thead>
		`
		bodytempl += "<tbody>"
		var body string
		for k := 0; k < len(StructureFrizzPointer.EmailInfo.EmailSession); k++ {
			body += StructureFrizzPointer.EmailInfo.EmailSession[k]
		}
		// use regex to find data eg hosts and emails
		regemail := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
		emails := regemail.FindAllString(body, -1)
		bodytempl += "<tr>"
		bodytempl += Generate_TD() + "Found Emails" + Generate_TDE()
		bodytempl += Generate_TD()
		emails = Frizz_Helpers.ValueRemover(emails)
		for _, element := range emails {
			if element != "" {
				bodytempl += element + "<br>"
			}
		}
		bodytempl += Generate_TDE()
		bodytempl += "</tr>"
		// use regex to find hosts
		regpat := regexp.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
		submatchall := regpat.FindAllString(body, -1)
		bodytempl += "<tr>"
		bodytempl += Generate_TD() + "Found hosts" + Generate_TDE()
		bodytempl += Generate_TD()
		submatchall = Frizz_Helpers.ValueRemover(submatchall)
		for _, element2 := range submatchall {
			if element2 != "" {
				bodytempl += element2 + "<br>"
			}
		}
		bodytempl += Generate_TDE()
		bodytempl += "</tr>"
		bodytempl += "</tbody></table>"
		bodytempl += "<br>"
		bodytempl += `<pre class="syntax">`
		for o := 0; o < len(StructureFrizzPointer.EmailInfo.EmailSession); o++ {
			reg := regexp.MustCompile(`<.*?>`)
			bodytempl += reg.ReplaceAllString(StructureFrizzPointer.EmailInfo.EmailSession[o], "")
		}
		bodytempl += `</pre>`
	case "*em":
		bodytempl += `
		<thead>
		<tr>
			<th>
				<h1>Section</h1>
			</th>
			<th>
				<h1>Email</h1>
			</th>
		</tr>
	</thead>
		`
		bodytempl += "<tbody>"
		bodytempl += "<tr>"
		bodytempl += Generate_TD() + "Found Emails" + Generate_TDE()
		bodytempl += Generate_TD()
		var bodytosearch string
		// Basically load anything email related or SMTP, FTP, SSH, TELNET, Session, POP related onto one string to be searched with regex
		for k := 0; k < len(StructureFrizzPointer.EmailInfo.EmailSession); k++ {
			bodytosearch += StructureFrizzPointer.EmailInfo.EmailSession[k]
		}
		for l := 0; l < len(StructureFrizzPointer.SMTPSessionInf.Body); l++ {
			bodytosearch += StructureFrizzPointer.SMTPSessionInf.Body[l]
		}
		for a := 0; a < len(StructureFrizzPointer.Tels.Body); a++ {
			bodytosearch += StructureFrizzPointer.Tels.Body[a]
		}
		for o := 0; o < len(StructureFrizzPointer.Ftps.Data); o++ {
			bodytosearch += StructureFrizzPointer.Ftps.Data[o]
		}
		regemail := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9_-]+)`)
		emails := regemail.FindAllString(bodytosearch, -1)
		emails = Frizz_Helpers.ValueRemover(emails)
		for _, elem := range emails {
			if elem != "" {
				bodytempl += elem + "<br>"
			}
		}
		bodytempl += "</tr>"
		bodytempl += Generate_TDE()
		bodytempl += "</tr>"
		bodytempl += "</tbody></table>"
		bodytempl += "<br>"
	}
	return bodytempl
}

func DrawDocumentPOP(docfile, payload string) {
	var templatepop3 string
	templatepop3 += StandardTopSessions // adding top
	templatepop3 += Standard_LinksList  // Adding links
	templatepop3 += StaticSection       // Adding static sections
	templatepop3 += Generate_Div("home-content")
	templatepop3 += Generate_Div("overview-boxes")
	tb1, tb2, tb3 := DrawByType(payload)
	templatepop3 += tb1
	templatepop3 += tb2
	templatepop3 += tb3
	templatepop3 += "</div><hr><br>"
	templatepop3 += `<table class="container_Overview"><br><br>`
	templatepop3 += "<tbody>"
	templatepop3 += DrawTableByPayload(payload)
	// load table generation
	templatepop3 += `</tbody></table>`
	templatepop3 += StaticJS           // loading js
	templatepop3 += StandardStyleSheet // generating stylesheet
	Write(docfile, templatepop3)
}

// HTTP DATA SECTION

func LoadTableByMethod(method string) string {
	var bod string
	switch method {
	case "hosts":
		bod += `
		<thead>
		<tr>
			<th>
				<h1>Hostname</h1>
			</th>
		</tr>
	</thead>`
		for i := 0; i < len(StructureFrizzPointer.Httpd.HTTPHostnames); i++ {
			bod += "<tr>"
			bod += Generate_TD() + StructureFrizzPointer.Httpd.HTTPHostnames[i] + Generate_TDE()
			bod += "</tr>"
		}

	case "urls":
		bod += `
		<thead>
		<tr>
			<th>
				<h1>Hostname</h1>
			</th>
		</tr>
	</thead>`
		for i := 0; i < len(StructureFrizzPointer.Httpd.HTTPUrls); i++ {
			bod += "<tr>"
			bod += Generate_TD() + StructureFrizzPointer.Httpd.HTTPUrls[i] + Generate_TDE()
			bod += "</tr>"
		}
	default:
		bod += `
		<thead>
		<tr>
			<th>
				<h1>Engine Error</h1>
			</th>
		</tr>
	</thead>`
		bod += "<tr>"
		bod += Generate_TDE() + "ERROR: WARNING: FATAL: GENERATION ENGINE -> For some reason when decoding data this felt weird | Message " + method + " Does not exist as a valid body message" + Generate_TDE()
		bod += "</tr>"
	}
	return bod
}

func DrawBoxByMethod(method string) (string, string) {
	switch method {
	case "urls":
		return Generate_Box("Total URLS", fmt.Sprint(len(StructureFrizzPointer.Httpd.HTTPUrls))), Generate_Box("Topic", "URL search")
	case "hosts":
		return Generate_Box("Total Hostnames", fmt.Sprint(len(StructureFrizzPointer.Httpd.HTTPHostnames))), Generate_Box("Topic", "Hostnames")
	default:
		return Generate_Box("Error", "Message decoding error| Unknown message"), Generate_Box("Error", " Message decoding error| Unknown message")
	}
}

func DrawDocumentHTTP(docfile, method string) {
	var HTTPTEMPLATE string
	HTTPTEMPLATE += StandardTopSessions
	HTTPTEMPLATE += Standard_LinksList
	HTTPTEMPLATE += StaticSection
	HTTPTEMPLATE += Generate_Div("home-content")
	HTTPTEMPLATE += Generate_Div("overview-boxes")
	tb1, tb2 := DrawBoxByMethod(method)
	HTTPTEMPLATE += tb1
	HTTPTEMPLATE += tb2
	HTTPTEMPLATE += "</div><hr><br>"
	HTTPTEMPLATE += `<table class="container_Overview"><br><br>`
	HTTPTEMPLATE += "<tbody>"
	HTTPTEMPLATE += LoadTableByMethod(method)
	HTTPTEMPLATE += `</tbody></table>`
	HTTPTEMPLATE += StandardStyleSheet
	HTTPTEMPLATE += StaticJS
	Write(docfile, HTTPTEMPLATE)
}

// Draw and generate HTTP general section

func DrawDocHTTPGENERAL(df string) {
	var HTTPGEN string
	HTTPGEN += StandardTopSessions
	HTTPGEN += Standard_LinksList
	HTTPGEN += StaticSection
	// Generate
	HTTPGEN += `<pre class="syntax">`
	for i := 0; i < len(StructureFrizzPointer.Httpd.HTTPFullSessionData); i++ {
		HTTPGEN += StructureFrizzPointer.Httpd.HTTPFullSessionData[i]
	}
	HTTPGEN += `</pre>`

	// Data

	HTTPGEN += StandardStyleSheet
	HTTPGEN += StaticJS
	Write(df, HTTPGEN)
}

// Draw and generate home page
func LoadBox(title, value string) string {
	return Generate_Box(title, value)
}

var DescMap = map[string]string{
	"Modules/Server/HTML/Useragents.html":         "Shows all collected useragents from HTTP requests packets",
	"Modules/Server/HTML/Hostnames.html":          "Shows all collected hostnames from HTTP request packets",
	"Modules/Server/HTML/URLs.html":               "Shows all collected URL's from HTTP request packets",
	"Modules/Server/HTML/DNS.html":                "Shows all collected DNS data from DNS based packets",
	"Modules/Server/HTML/OpenPorts.html":          "Shows all hostnames and their open ports at the time",
	"Modules/Server/HTML/ARP.html":                "Shows all ARP requests, responses etc",
	"Modules/Server/HTML/Ethernet.html":           "Shows all ethernet packets",
	"Modules/Server/HTML/Servers.html":            "Shows all server information gathered",
	"Modules/Server/HTML/Wifi.html":               "Shows all wifi information collected such as Probe request and response",
	"Modules/Server/HTML/WifiOspf.html":           "Shows all Wifi Warningsentication requests and responses",
	"Modules/Server/HTML/FTP.html":                "Shows FTP session data",
	"Modules/Server/HTML/SSH.html":                "Shows SSH session data",
	"Modules/Server/HTML/Telnet.html":             "Shows Telnet session data",
	"Modules/Server/HTML/SMTP.html":               "Shows SMTP session data",
	"Modules/Server/HTML/SIP.html":                "Shows SIP session data",
	"Modules/Server/HTML/AuthFTPCreds.html":       "Shows AuthFTPCreds session data",
	"Modules/Server/HTML/AuthSSHCreds.html":       "Shows SSH credentials",
	"Modules/Server/HTML/AuthIMAP.html":           "Shows IMAP credentials",
	"Modules/Server/HTML/AuthDigest.html":         "Shows all HTTP DIGEST based authentication and credentials",
	"Modules/Server/HTML/AuthNTLM.html":           "Shows all HTTP NTLM based authentication and credentials",
	"Modules/Server/HTML/AuthBASIC.html":          "Shows all HTTP BASIC based authentication and credentials",
	"Modules/Server/HTML/AuthNegotiation.html":    "Shows all Negotiate HTTO based authentication and credentials",
	"Modules/Server/HTML/AuthSMTP.html":           "Shows all SMTP credentials Plain/B64/B32...etc",
	"Modules/Server/HTML/Emails.html":             "Shows found emails",
	"Modules/Server/HTML/Cc.html":                 "Shows found emails by Cc payload",
	"Modules/Server/HTML/From.html":               "Shows found data in SMTP packets using the From payload",
	"Modules/Server/HTML/Recv.html":               "Shows found data in SMTP packets using the Recv payload",
	"Modules/Server/HTML/Convos.html":             "Shows all conversations found in a SMTP session",
	"Modules/Server/HTML/Masher.html":             "Opens a file of directories to pcap file and mashes them all into one pcap file [ read docs for more info ]",
	"Modules/Server/HTML/Raw.html":                "Displays packets in raw form",
	"Modules/Server/HTML/ServerRequirements.html": "Shows server requirements or data the server collects",
	"Modules/Server/HTML/JSONDB.html":             "Shows the JSON files and their contents of the data stored in the local server",
	"Modules/Server/HTML/AppInfo.html":            "Shows application data such as versions, types, files, paths, imports, bugs etc ",
	"Modules/Server/HTML/ServerInfo.html":         "Shows server information such as about the server, types, files and importing CDN's",
	"Modules/Server/HTML/Documentation.html":      "Documentation of frizz",
	"Modules/Server/HTML/AuthTelnet.html":         "Telnet authentication",
	"Modules/Server/HTML/HTTPSESSION.html":        "Shows all HTTP GET/POST/FORM REQUESTS found in the packet data",
	"Modules/Server/HTML/Home.html":               "Current page you are at",
}

func GenerateHomeTable() string {
	var Table string
	Table += `
		<thead>
		<tr>
			<th>
				<h1>File Index Number</h1>
			</th>
			<th>
				<h1>URL</h1>
			</th>
			<th>
				<h1>Description</h1>
			</th>
		</tr>
	</thead>
	`
	amountindexes, _ := strconv.Atoi(StructureAppInfo.ApplicationIndexes)
	for l := 0; l < amountindexes; l++ {
		Table += "<tr>"
		Table += Generate_TD() + fmt.Sprint(l) + Generate_TDE()
		p, err := filepath.Abs(StructureServerInfo.ServerFiles[l])

		if err != nil {

			log.Fatal(err)
		}
		Table += Generate_TD() + "http://localhost:5674/" + p + Generate_TDE()
		// Finall generate description
		Table += Generate_TD() + DescMap[StructureServerInfo.ServerFiles[l]] + Generate_TDE()
	}

	return Table
}

func DrawHome(df string) {
	var Home string
	Home += StandardTop
	Home += Standard_LinksList
	Home += StaticSection
	Home += Generate_Div("home-content")
	Home += Generate_Div("overview-boxes")
	Home += LoadBox("Total Packets", fmt.Sprint(Frizzlog.Analytics.TotalPackets))
	Home += LoadBox("File Extension", fmt.Sprint(filepath.Ext(Frizzlog.Analytics.Filename)))
	Home += LoadBox("Seconds to parse file", fmt.Sprint(Frizzlog.Analytics.TimeToParse))
	Home += "</div><hr><br>"
	Home += `<table class="container_Overview"><br><br>`
	Home += `<tbody>`
	Home += GenerateHomeTable()
	Home += `</tbody></table>`
	Home += StandardStyleSheet
	Home += StaticJS
	Write(df, Home)
}

// Generate wifi data

func GenerateDangerouWifi(fname string) {
	var template string
	template += WIFIHTMLTOP
	template += Generate_Div("home-content")
	template += Generate_Div("overview-boxes")
	template += LoadBox("Dangerous by string", fmt.Sprint(FrizzWifu.DangerousStructure.NumStringBased))
	template += LoadBox("Dangerous by length", fmt.Sprint(FrizzWifu.DangerousStructure.NumLengthBased))
	template += "</div><hr><br><br><br><br>"
	template += `<div class="codeheader" id="Topic">Possible Dangerous SSID's </div>`
	template += `<table class="container_Overview"><br><br>`
	template += `
	<thead>
		<tr>
			<th>
				<h1>SSID</h1>
			</th>
			<th>
				<h1>Length</h1>
			</th>
			<th>
				<h1>Valid by string?</h1>
			</th>
			<th>
				<h1>Valid by length?</h1>
			</th>
		</tr>
	</thead>
	`
	template += "</tr></thead><tbody>"
	for i := 0; i < len(FrizzWifu.DangerousStructure.DangerouSSID); i++ {
		template += "<tr>"
		template += Generate_TD() + FrizzWifu.DangerousStructure.DangerouSSID[i] + Generate_TDE()
		template += Generate_TD() + fmt.Sprint(len(FrizzWifu.DangerousStructure.DangerouSSID[i])) + Generate_TDE()
		template += Generate_TD() + "NOT ASSUMED" + Generate_TDE()
		if len(FrizzWifu.DangerousStructure.DangerouSSID[i]) >= FrizzWifu.DangerousWarnLen {
			template += Generate_TD() + "True" + Generate_TDE()
		} else {
			template += Generate_TD() + "False" + Generate_TDE()
		}
		template += "</tr>"
	}
	template += `
	</tbody>
	</table>
	`
	template += StaticJS
	template += WIFIHTMLSTYLE

	Write(fname, template)
}

func GenerateWifiTable(topicid string, rangval string, HeadValues ...string) string {
	var WifiTable string
	WifiTable += fmt.Sprintf(`<div class="codeheader" id="Topic">%s</div>`, topicid)
	WifiTable += `<table class="container_Overview"><br><br>`
	WifiTable += `<thead><tr>`
	for o := 0; o < len(HeadValues); o++ {
		WifiTable += `<th>`
		WifiTable += Generate_H1START() + HeadValues[o] + Generate_H1END()
		WifiTable += `</th>`
	}
	WifiTable += `</tr></thead>`
	WifiTable += `<tbody>`
	switch rangval {
	case "OUI":
		for i := 0; i < len(StructureFrizzPointer.WifiProbe.ProbeMAC); i++ {
			WifiTable += "<tr>"
			WifiTable += Generate_TD() + StructureFrizzPointer.WifiProbe.ProbeMAC[i] + Generate_TDE()
			var tmpl string
			for _, ouis := range FrizzWifu.OUI(StructureFrizzPointer.WifiProbe.ProbeMAC[i]) {
				tmpl += ouis
			}
			WifiTable += Generate_TD() + tmpl + Generate_TDE()
			WifiTable += "</tr>"
		}
	case "SSID":
		for l := 0; l < len(StructureFrizzPointer.WifiProbe.ProbeSSID); l++ {
			WifiTable += "<tr>"
			WifiTable += Generate_TD() + "SSID" + Generate_TDE()
			WifiTable += Generate_TD() + StructureFrizzPointer.WifiProbe.ProbeSSID[l] + Generate_TDE()
			WifiTable += "</tr>"
		}
	case "MAC":
		for o := 0; o < len(StructureFrizzPointer.WifiProbe.ProbeMAC); o++ {
			WifiTable += "<tr>"
			WifiTable += Generate_TD() + StructureFrizzPointer.WifiProbe.ProbeMAC[o] + Generate_TD()
			WifiTable += "</tr>"
		}
	}
	WifiTable += "</tbody>"
	WifiTable += `</table>`
	return WifiTable
}

func DrawWifi(docf string) {
	var Wifi string
	var AmountOUIsValid int
	Wifi += WIFIHTMLTOP
	Wifi += Generate_Div("home-content")
	Wifi += Generate_Div("overview-boxes")
	Wifi += LoadBox("Total SSID's", fmt.Sprint(len(StructureFrizzPointer.WifiProbe.ProbeSSID)))
	Wifi += LoadBox("Total MAC's", fmt.Sprint(len(StructureFrizzPointer.WifiProbe.ProbeMAC)))
	for i := 0; i < len(StructureFrizzPointer.WifiProbe.ProbeMAC); i++ {
		for _, val := range FrizzWifu.OUI(StructureFrizzPointer.WifiProbe.ProbeMAC[i]) {
			if val != "" {
				AmountOUIsValid++
			}
		}
	}
	Wifi += LoadBox("Total Valid OUI's", fmt.Sprint(AmountOUIsValid))
	Wifi += "</div><hr><br><br><br><br>"
	// Generate tables
	Wifi += GenerateWifiTable("SSID's found", "SSID", "Type", "SSID")
	Wifi += GenerateWifiTable("MAC's found", "OUI", "MAC", "OUI")
	Wifi += StaticJS
	Wifi += WIFIHTMLSTYLE
	Write(docf, Wifi)
}

// Generate ARP index

func repetition(st []string) map[string]int {
	wc := make(map[string]int)
	for _, word := range st {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}

func GenerateARP(file string) {
	var ARPF string
	ARPF += WIFIHTMLSTYLE
	ARPF += Generate_Div("home-content")
	ARPF += Generate_Div("overview-boxes")
	MACSTOTAL := len(FrizzNet.Adata.DSTMAC) + len(FrizzNet.Adata.SRCMAC)
	PROTOTOTAL := len(FrizzNet.Adata.SRCIP) + len(FrizzNet.Adata.DSTIP)
	ARPF += LoadBox("Total MAC's", fmt.Sprint(MACSTOTAL))
	ARPF += LoadBox("Total Proto Addresses", fmt.Sprint(PROTOTOTAL))
	ARPF += "</div><hr><br><br><br><br>"
	ARPF += fmt.Sprint(FrizzNet.Adata.DSTMAC)
	for i := 0; i < len(FrizzNet.Adata.DSTMACOUI); i++ {
		fmt.Println("Value -> ", FrizzNet.Adata.DSTMACOUI[i], " Exists in element -> ", repetition(FrizzNet.Adata.DSTMACOUI))
	}
	// Generate bar values
	Write(file, ARPF)
}
