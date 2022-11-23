package Frizz_DatabaseStructure

import "time"

type Frizz struct {
	Email_Info            Email_Information
	SMTP_SessionInf       SMTP_Session
	IMAPCREDS             IMAP_Credential
	Creds                 Credentials
	SSHS                  SSH_Session
	FTPS                  FTP_Session
	TELS                  TELNET_session
	HTTPD                 HTTP_DATA
	WifiProbe             Probe
	CredentialsAll        int
	CredentialsFTP        int
	CredentialsIMAP       int
	CredentialsSSH        int
	CredentialsSMTP       int
	CredentialsHTTPNTLM   int
	CredentialsHTTPBASIC  int
	CredentialsHTTPNEG    int
	CredentialsHTTPDIGEST int
}

type Email_Information struct {
	Email_Mailbox []string // Emails found in mailbox
	Email_CC      []string // Emails found in CC
	//Email_AUTH_RAW []string // Raw Base 64 authentication
	//Email_AUTH_DEC []string // Decoded Base 64 authentication
	Email_FROM     []string // Parses any emails and data within the from response
	Email_Recieved []string // Parses any emails or data that matched the Email_RECV payload
	Email_Session  []string // Any email conversations or entire payloads
}

type SMTP_Session struct {
	Body []string
}

type IMAP_Credential struct {
	IMAP_PLAINTEXT          []string
	IMAP_BASE64_Decoded     []string
	IMAP_BASE64_Encoded     []string
	IMAP_DIGEST_MD5_Encoded []string
	IMAP_DIGEST_MD5_Decoded []string
}

type SSH struct {
	SSHUsername []string
	SSHPassword []string
}

type Credentials struct {
	Ftp       FTP_Creds
	Smtp      SMTP_Creds
	Smtp_Cram CramMD5
	Ssh       SSH
}

type FTP_Creds struct {
	FTP_Username []string
	FTP_Password []string
}

type CramMD5 struct {
	Encoded []string
	Decoded []string
}

type SMTP_Creds struct {
	SMTP_Username          []string
	SMTP_Password          []string
	SMTP_Plainauth         []string
	SMTP_Plainauth_Decodec []string
}

type SSH_Session struct {
	Body []string
}

type FTP_Session struct {
	Data []string
}

type TELNET_session struct {
	Body []string
}

type HTTP_DATA struct {
	HTTP_Useragents        []string
	HTTP_Hostnames         []string
	HTTP_URLS              []string
	HTTP_NTLM_Encoded      []string
	HTTP_BASIC_ENCODED     []string
	HTTP_BASIC_DECODED     []string
	HTTP_DIGEST            []string
	HTTP_NEGOTIATE         []string
	HTTP_FULL_SESSION_DATA []string
	HTTP_USERAGENT_HOST    []string
	Uagent_Host_Host       []string
	Uagent_Host_Uagent     []string
}

type Probe struct {
	Probe_SSID      []string
	Probe_BSSID     []string
	Probe_MAC       []string
	Probe_Channel   []string
	Probe_Rate      []string
	Probe_Vendor    []byte
	Probe_Frequency []int8
	Probe_Time      time.Time
}
