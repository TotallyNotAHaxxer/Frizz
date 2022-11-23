<p align="center">
  <img src="logo.png" width="200" title="frizz logo">
</p>

# What is frizz 

frizz is a web interface based NFAT ( Network Forensics Analyitics Tool ) which was designed to capture and parse the most valuable information to someone looking through packet captures. This web interface was designed to be modern and sleek with some decent backend work to it. There is many forms of information frizz will give you such as 

| Type and or title | Data element or list of data | 
| ----------------- | ---------------------------- |
| Authentication    | HTTP Digest/NTLM/BASIC/Negotiate, SMTP, FTP, SSH, IMAP | 
| User Information  | Email Cc/From/To/Inbox/recv/chats |
| Session Information |  FTP sessions, SSH sessions, Telnet sessions, SMTP sessions | 
| Server and App info | Server information, user hardware, databases used | 
| Documentation | Information about tools | 
| HTTP info | HTTP Hostnames, Useragents, URL's, GET, HEAD, POST, PUT | 
| Host info | Ports, Servers and DNS information |
| Proto info | Ethernet and ARP | 
| Dot11 Elements | OSPF Authentication, Radio tap, Probe requests |

# installing pcapplusplus [Required]

```
	git clone https://github.com/seladb/PcapPlusPlus.git ; cd PcapPlusPlus ; ./PcapPlusPlus/configure-linux.sh --default ; make ; make install ; rm -rf PcapPlusPlus
```

# installing crystal [Required]

crystal is a programming language that is used to run the server, without crystal there is no server so there is no tool. install crystal with 

`snap install crystal`

or 

https://crystal-lang.org/install/on_debian/

# install frizz [Required]

`git clone https://github.com/ArkAngeL43/Frizz.git ; cd Frizz ; make all `

> Note: Frizz relies on the PcapPlusPlus and Gopacket packet capturing and parsing libraries. This means you must have libssl installed, libpcap, libc etc libs. Below is a command to install all of the required libs. you can run command `make libs` if you do not have the libs installed.

once done making all files run 

`./router pcapfile.pcap`

replace pcapfile.pcap with you're desired pcap or pcapng file

# How does frizz work 

frizz uses many useful algorithms and techniques for parsing and figuring out what packets have inside of them. When you tell frizz to load a new packet capture or pcap file it will use a series of regular expressions, manual conditionals and a prediction technique to assume certain data. Frizz will take certain data such as credentials and load them through individal parsers to test the data string by payload to see what exactly it matches to. If it matches a Base64/32 encoding it will decode it and parse it off as a certain string such as HTTP MD5-Dige

# Developer notes about the framework 

> 1: Frizz was designed to be a modern day NFAT to beat a very specific group of people in compeition who were selling nothing close in fact worse for more money. This means that frizz is still in testing and this release is not the most satisfying framework or supportive framework as it is still missing a few things. However when frizz gets updated it will be sold off with more deeper features. See the list below for more information on what frizz will hold 

| Frizz Professional license feature | feature description / tool description | 
| ---------------------------------- | -------------------------------------- |
| Support for more credentials       | More credential types such as SMB NTLM, IMAP XYMPKI, MB - NTLMSSP, Postgres, MSSQL (TDS 7.0+), LDAP etc | 
| Support for more information       | This will include more user information such as phone numbers which will use internal databases to grab geography information on the phone number, this means frizz later on will also be able to parse geopgraphical maps, generate images, generate cell information, parse information of certain computers and devices, give more valuable MAC information etc. This type of information was not given in frizz beta because of how valuable and how private the databases are to the developer 
| Support for more protocols | Support for CAN, USB, Wifi, UDP, RTP/SIP etc will be implimented | 
| Support for a better user interface and command line interface | Right now frizz has a decent to moderate web interface but the CLI is nothing other than a single argument. Later on frizz will have more Intrusion Detection Systems embedded in the web interface and options  to listen for connections and to parse those connections | 
| Support for a better more offensive and defensive interface | This means that frizz will be designed to fit needs such as service scanning, exploit checking, port service information, service and host information, DNS geographical locations and whois information, tools to login and to re create and assemble TCP/Telnet/SSH/SMTP sessions | 
| Support with bigger databases | Right now frizz does not have a database for geographical information, MAC information, Wifi or key information etc it purely assumes what the type is based on length, type of letters and numbers, sequences and bits in those payloads. Later on frizz will have offline databases that can load that kind of data | 
| Closed source | The paid version of frizz will be closed source but much more faster, responsive, nicer, cleaner and way more accurate with data information | 

> Frizz is weird: Frizz has a weird code base and is mostly a mash or remake of other parsers such as the user agent parser. Frizz relies on external sources such as CDN's for charts and other things which does not make this a fully offline interface however you can still run it offline. A majority of the frizz project is decently designed but weirdly implimented but after all this is beta and beta desinges are wack. IF YOU PLAN TO READ SOURCE CODE READ THE NOTES DO NOT JUST IGNORE THEM.
