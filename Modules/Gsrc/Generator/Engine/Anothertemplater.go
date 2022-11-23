package Engine

import (
	"fmt"
	"strings"
)

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
		`
		for k := 0; k < len(StructureFrizzPointer.Creds.Ftp.FTPPassword); k++ {
			templater += "<tr>"
			templater += Generate_TD() + StructureFrizzPointer.Creds.Ftp.FTPPassword[k] + Generate_TDE()
			templater += Generate_TD() + StructureFrizzPointer.Creds.Ftp.FTPUsername[k] + Generate_TDE()
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
	`
		for k := 0; k < len(StructureFrizzPointer.Creds.SMTPCram.Decoded); k++ {
			templater += "<tr>"
			templater += Generate_TD() + StructureFrizzPointer.Creds.SMTPCram.Encoded[k] + Generate_TDE()
			templater += Generate_TD() + StructureFrizzPointer.Creds.SMTPCram.Decoded[k] + Generate_TDE()
			templater += Generate_TD() + "CRAM-MD5" + Generate_TDE()
			templater += "</tr>"
		}
		for l := 0; l < len(StructureFrizzPointer.Creds.SMTP.SMTPPlainauth); l++ {
			templater += "<tr>"
			templater += Generate_TD() + StructureFrizzPointer.Creds.SMTP.SMTPPlainauth[l] + Generate_TDE()
			templater += Generate_TD() + StructureFrizzPointer.Creds.SMTP.SMTPPlainauthDecodec[l] + Generate_TDE()
			templater += Generate_TD() + "SMTP PLAIN AUTHENTICATION" + Generate_TDE()
			templater += "</tr>"
		}
		for q := 0; q < len(StructureFrizzPointer.Creds.SMTP.SMTPUsername); q++ {
			templater += "<tr>"
			templater += Generate_TD() + StructureFrizzPointer.Creds.SMTP.SMTPUsername[q] + Generate_TDE()
			templater += Generate_TD() + StructureFrizzPointer.Creds.SMTP.SMTPPassword[q] + Generate_TDE()
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
	`
		for p := 0; p < len(StructureFrizzPointer.Httpd.HTTPNegotiate); p++ {
			templater += "<tr>"
			if StructureFrizzPointer.Httpd.HTTPNegotiate[p] != "" {
				templater += Generate_TD() + StructureFrizzPointer.Httpd.HTTPNegotiate[p] + Generate_TDE()
				templater += Generate_TD() + "HTTP Negotiate" + Generate_TDE()
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
	anotherftemplate += "</thead><tbody>"
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
	default:
		return "empty or error"
	}
}
