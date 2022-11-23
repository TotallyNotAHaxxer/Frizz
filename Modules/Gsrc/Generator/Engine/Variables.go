package Engine

import (
	"sync"
	"time"
)

var (
	StructureFrizzPointer Frizz
	StructurePreProcessor HostSysInformation
	StructurePcapPlusPlus PcppCaptureInfo
	StructureAppInfo      APPINFO
	StructureServerInfo   ServerInf
)
var Files = []string{
	"Modules/Server/HTML/Analytics.html",
	"Modules/Server/HTML/PacketStats.html",
	"Modules/Server/HTML/Useragents.html",
	"Modules/Server/HTML/Hostnames.html",
	"Modules/Server/HTML/URLs.html",
	"Modules/Server/HTML/DNS.html",
	"Modules/Server/HTML/OpenPorts.html",
	"Modules/Server/HTML/ARP.html",
	"Modules/Server/HTML/Ethernet.html",
	"Modules/Server/HTML/Servers.html",
	"Modules/Server/HTML/Wifi.html",
	"Modules/Server/HTML/WifiOspf.html",
	"Modules/Server/HTML/FTP.html",
	"Modules/Server/HTML/SSH.html",
	"Modules/Server/HTML/Telnet.html",
	"Modules/Server/HTML/SMTP.html",
	"Modules/Server/HTML/SIP.html",
	"Modules/Server/HTML/AuthFTPCreds.html",
	"Modules/Server/HTML/AuthSSHCreds.html",
	"Modules/Server/HTML/AuthIMAP.html",
	"Modules/Server/HTML/AuthDigest.html",
	"Modules/Server/HTML/AuthNTLM.html",
	"Modules/Server/HTML/AuthBASIC.html",
	"Modules/Server/HTML/AuthNegotiation.html",
	"Modules/Server/HTML/AuthSMTP.html",
	"Modules/Server/HTML/Emails.html",
	"Modules/Server/HTML/Cc.html",
	"Modules/Server/HTML/From.html",
	"Modules/Server/HTML/Recv.html",
	"Modules/Server/HTML/Convos.html",
	"Modules/Server/HTML/Masher.html",
	"Modules/Server/HTML/Raw.html",
	"Modules/Server/HTML/ServerRequirements.html",
	"Modules/Server/HTML/JSONDB.html",
	"Modules/Server/HTML/AppInfo.html",
	"Modules/Server/HTML/ServerInfo.html",
	"Modules/Server/HTML/Documentation.html",
}

var Database = []string{
	"Modules/Gsrc/Database/FrizzDatabase.json",
	"Modules/Gsrc/Database/PCPP.json",
	"Modules/Gsrc/Database/PreProcessor.json",
	"Modules/Gsrc/Database/ServerInfo.json",
	"Modules/Gsrc/Database/ApplicationInformationStorage.json",
}

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
		SSH struct {
			SSHUsername []string `json:"SSHUsername"`
			SSHPassword []string `json:"SSHPassword"`
		} `json:"Ssh"`
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
		HTTPUseragentHost   []string `json:"HTTP_USERAGENT_HOST"`
		UagentHostHost      []string `json:"Uagent_Host_Host"`
		UagentHostUagent    []string `json:"Uagent_Host_Uagent"`
	} `json:"HTTPD"`
	WifiProbe struct {
		ProbeSSID      []string  `json:"Probe_SSID"`
		ProbeBSSID     []string  `json:"Probe_BSSID"`
		ProbeMAC       []string  `json:"Probe_MAC"`
		ProbeChannel   []string  `json:"Probe_Channel"`
		ProbeRate      []string  `json:"Probe_Rate"`
		ProbeVendor    []string  `json:"Probe_Vendor"`
		ProbeFrequency []string  `json:"Probe_Frequency"`
		ProbeTime      time.Time `json:"Probe_Time"`
	} `json:"WifiProbe"`
	CredentialsAll        int `json:"CredentialsAll"`
	CredentialsFTP        int `json:"CredentialsFTP"`
	CredentialsIMAP       int `json:"CredentialsIMAP"`
	CredentialsSSH        int `json:"CredentialsSSH"`
	CredentialsSMTP       int `json:"CredentialsSMTP"`
	CredentialsHTTPNTLM   int `json:"CredentialsHTTPNTLM"`
	CredentialsHTTPBASIC  int `json:"CredentialsHTTPBASIC"`
	CredentialsHTTPNEG    int `json:"CredentialsHTTPNEG"`
	CredentialsHTTPDIGEST int `json:"CredentialsHTTPDIGEST"`
}

type PcppCaptureInfo []struct {
	AddressResolutinProtocolTalkFullConvo []string `json:"Address Resolutin Protocol Talk ( Full convo ) "`
	AddressResolutionProtocolTalkDSTIP    []string `json:"Address Resolution Protocol Talk ( DST IP ) "`
	AddressResolutionProtocolTalkDSTMAC   []string `json:"Address Resolution Protocol Talk ( DST MAC ) "`
	AddressResolutionProtocolTalkSRCIP    []string `json:"Address Resolution Protocol Talk ( SRC IP )"`
	AddressResolutionProtocolTalkSRCMAC   []string `json:"Address Resolution Protocol Talk ( SRC MAC ) "`
	ETHERNETDestinationMACS               []string `json:"ETHERNET (destination) MAC's "`
	ETHERNETSourceMACS                    []string `json:"ETHERNET (source) MAC's "`
	ETHERNETTALKSRCTODST                  []string `json:"ETHERNET TALK SRC TO DST"`
	FileLinkLayerType                     string   `json:"File Link Layer Type"`
	TotalApplicationPackets               int      `json:"Total_Application_Packets"`
	TotalParsedPackets                    int      `json:"Total_Parsed_Packets"`
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

type JSONSERVERINFO struct {
	ServerUrls         []string `json:"Server_Urls"`
	ServerFiles        []string `json:"Server_Files"`
	ServerDatabase     []string `json:"Server_Database"`
	ServerPreProcessor []string `json:"Server_PreProcessor"`
	ServerSupport      []string `json:"Server_Support"`
	ServerLanguages    []string `json:"Server_Languages"`
	ServerImports      []string `json:"Server_Imports"`
	ServerSuggests     []string `json:"Server_Suggests"`
	ServerVersion      string   `json:"Server_Version"`
	ServerMainPort     int      `json:"Server_Main_Port"`
	ServerMainURL      string   `json:"Server_Main_URL"`
	ServerPorts        []int    `json:"Server_Ports"`
}

type APPINFO struct {
	Version                         string `json:"Version"`
	Name                            string `json:"Name"`
	ServerName                      string `json:"Server Name"`
	ApplicationLanguages            string `json:"Application_Languages"`
	ApplicationFileFormats          string `json:"Application_FileFormats"`
	ApplicationSecurity             string `json:"Application_Security"`
	ApplicationBugs                 string `json:"Application_Bugs"`
	ApplicationTabs                 string `json:"Application_Tabs"`
	ApplicationIndexes              string `json:"Application_Indexes"`
	ApplicationProtocols            string `json:"Application_Protocols"`
	ApplicationSupport              string `json:"Application_Support"`
	ApplicationBinaries             string `json:"Application_Binaries"`
	ApplicationRunBinary            string `json:"Application_Run_Binary"`
	ApplicationRunTime              string `json:"Application_Run_Time"`
	ApplicationSupportedFileFormats string `json:"Application_Supported_File_Formats"`
}

type EngineViewer struct {
	File string
	WG   sync.Mutex
}

var importlist = []string{
	"https://unpkg.com/boxicons@2.1.2/css/boxicons.min.css",
	"https://cdnjs.cloudflare.com/ajax/libs/Chart.js/2.9.4/Chart.js",
	"https://cdn.plot.ly/plotly-latest.min.js",
}
