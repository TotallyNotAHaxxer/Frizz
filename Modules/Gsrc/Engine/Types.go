package Frizz_Engine

import "time"

type Frizz struct {
	EmailInfo struct {
		EmailMailbox  []string `json:"Email_Mailbox"`
		EmailCC       []string `json:"Email_CC"`
		EmailFROM     []string `json:"Email_FROM"`
		EmailRecieved []string `json:"Email_Recieved"`
		EmailSession  []string `json:"Email_Session"`
	} `json:"Email_Info"`
	SMTPSessionInf struct {
		Body []string `json:"Body"`
	} `json:"SMTP_SessionInf"`
	Imapcreds struct {
		ImapPlaintext        []string `json:"IMAP_PLAINTEXT"`
		IMAPBASE64Decoded    []string `json:"IMAP_BASE64_Decoded"`
		IMAPBASE64Encoded    []string `json:"IMAP_BASE64_Encoded"`
		IMAPDIGESTMD5Encoded []string `json:"IMAP_DIGEST_MD5_Encoded"`
		IMAPDIGESTMD5Decoded []string `json:"IMAP_DIGEST_MD5_Decoded"`
	} `json:"IMAPCREDS"`
	Creds struct {
		Ftp struct {
			FTPUsername []string `json:"FTP_Username"`
			FTPPassword []string `json:"FTP_Password"`
		} `json:"Ftp"`
		SMTP struct {
			SMTPUsername         []string `json:"SMTP_Username"`
			SMTPPassword         []string `json:"SMTP_Password"`
			SMTPPlainauth        []string `json:"SMTP_Plainauth"`
			SMTPPlainauthDecodec []string `json:"SMTP_Plainauth_Decodec"`
		} `json:"Smtp"`
		SMTPCram struct {
			Encoded []string `json:"Encoded"`
			Decoded []string `json:"Decoded"`
		} `json:"Smtp_Cram"`
	} `json:"Creds"`
	Sshs struct {
		Body []string `json:"Body"`
	} `json:"SSHS"`
	Ftps struct {
		Data []string `json:"Data"`
	} `json:"FTPS"`
	Tels struct {
		Body []string `json:"Body"`
	} `json:"TELS"`
	Httpd struct {
		HTTPUseragents      []string `json:"HTTP_Useragents"`
		HTTPHostnames       []string `json:"HTTP_Hostnames"`
		HTTPUrls            []string `json:"HTTP_URLS"`
		HTTPNTLMEncoded     []string `json:"HTTP_NTLM_Encoded"`
		HTTPBasicEncoded    []string `json:"HTTP_BASIC_ENCODED"`
		HTTPBasicDecoded    []string `json:"HTTP_BASIC_DECODED"`
		HTTPDigest          []string `json:"HTTP_DIGEST"`
		HTTPNegotiate       []string `json:"HTTP_NEGOTIATE"`
		HTTPFullSessionData []string `json:"HTTP_FULL_SESSION_DATA"`
	} `json:"HTTPD"`
	WifiProbe struct {
		ProbeSSID      string    `json:"Probe_SSID"`
		ProbeBSSID     string    `json:"Probe_BSSID"`
		ProbeMAC       string    `json:"Probe_MAC"`
		ProbeChannel   string    `json:"Probe_Channel"`
		ProbeRate      string    `json:"Probe_Rate"`
		ProbeVendor    []string  `json:"Probe_Vendor"`
		ProbeFrequency int       `json:"Probe_Frequency"`
		ProbeTime      time.Time `json:"Probe_Time"`
	} `json:"WifiProbe"`
}

type HostSysInformation struct {
	ServerOperatingSystem                       string `json:"Server_OperatingSystem"`
	ServerOperatingSystemFileSystem             string `json:"Server_OperatingSystemFileSystem"`
	ServerOperatingSystemVersion                string `json:"Server_OperatingSystemVersion"`
	ServerOperatingArchitecture                 string `json:"Server_OperatingArchitecture"`
	ServerHardwareInfoCPUVendorID               string `json:"Server_HardwareInfo_CPU_VendorID"`
	ServerHardwareInfoCPUIndexNum               string `json:"Server_HardwareInfo_CPU_IndexNum"`
	ServerHardwareInfoCPUFamily                 string `json:"Server_HardwareInfo_CPU_Family"`
	ServerHardwareInfoCPUNumberOfCores          string `json:"Server_HardwareInfo_CPU_NumberOfCores"`
	ServerHardwareInfoCPUModelName              string `json:"Server_HardwareInfo_CPU_ModelName"`
	ServerHardwareInfoCPUSpeed                  string `json:"Server_HardwareInfo_CPU_Speed"`
	ServerHardwareInfoCPUCacheSize              string `json:"Server_HardwareInfo_CPU_CacheSize"`
	ServerHardwareInfoCPUMicronode              string `json:"Server_HardwareInfo_CPU_Micronode"`
	ServerHardwareInfoCPUModel                  string `json:"Server_HardwareInfo_CPU_Model"`
	ServerHardwareInfoCPUPhysID                 string `json:"Server_HardwareInfo_CPU_PhysID"`
	ServerHardwareInfoCPUStep                   string `json:"Server_HardwareInfo_CPU_Step"`
	ServerHardwareInfoMEMFree                   string `json:"Server_HardwareInfo_MEM_Free"`
	ServerHardwareInfoMEMTotal                  string `json:"Server_HardwareInfo_MEM_Total"`
	ServerHardwareInfoMEMUsed                   string `json:"Server_HardwareInfo_MEM_Used"`
	ServerHardwareInfoOSPHostname               string `json:"Server_HardwareInfo_OSP_Hostname"`
	ServerHardwareInfoOSPUptime                 string `json:"Server_HardwareInfo_OSP_Uptime"`
	ServerHardwareInfoOSPProcRunning            string `json:"Server_HardwareInfo_OSP_ProcRunning"`
	ServerHardwareInfoOSPHOSTID                 string `json:"Server_HardwareInfo_OSP_HOSTID"`
	ServerHardwareInfoOSPHOSTOS                 string `json:"Server_HardwareInfo_OSP_HOSTOS"`
	ServerHardwareInfoOSPHOSTPLAT               string `json:"Server_HardwareInfo_OSP_HOSTPLAT"`
	ServerHardwareInfoOSPHOSTKERNELVERSION      string `json:"Server_HardwareInfo_OSP_HOST_KERNEL_VERSION"`
	ServerHardwareInfoOSPHOSTKERNELARCHITECTURE string `json:"Server_HardwareInfo_OSP_HOST_KERNEL_ARCHITECTURE"`
	ServerHardwareInfoOSPHOSTPLATFORMVERSION    string `json:"Server_HardwareInfo_OSP_HOST_PLATFORM_VERSION"`
	ServerHardwareInfoOSPHOSTPLATFORMFAMILY     string `json:"Server_HardwareInfo_OSP_HOST_PLATFORM_FAMILY"`
}
