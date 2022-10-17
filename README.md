<p align="center">
  <img src="Simple.png" width="600" title="Used">
</p>

### Frizz subinfo ### 

> Version: 0.0.1 (BETA)
> Languages: C++, Crystal, Go 
> Operating systems: Linux
> File support: Cap, Pcap, Pcapng
> Features: 34 

# What exactly is frizz

Well many things can describe frizz, mainly frizz is a NFAT. Network Forensics Analysis Tool for pcap files. Frizz will take any pcap file and will purely generate information, parse through every single packet and generate an entire web interface where information is sorted out. Frizz is not for basic packet capturing more than it is for people who want to look for specific information. Despite frizz parsing layers such as Ethernet, ARP, SIP, HTTP, TCP etc it's main job is to parse information. Below is a table of what frizz trys to find.

| Information, data, topics, etc | 
| ------------------------------ | 
| Emails |
| Phone numbers | 
| Useragents | 
| Email chats | 
| Email Cc lists | 
| Email to | 
| Email From | 
| SSH passwords and authentication | 
| Wifi OSPF authentication | 
| SMTP authentication | 
| POP3 authentication | 
| IMAP authentication | 
| FTP authentication | 
| Telnet sessions | 
| FTP sessions | 
| SSH sessions | 
| POP3 sessions |
| SMTP sessions |
| IMAP sessions | 
| HTTP POST | 
| HTTP GET | 
| TCP | 
| ETHERNET | 
| Wifi | 
| SIP invites |
| SIP codes | 
| HTTP URL's | 
| Base64 encodings | 
| Base32 encodings | 
| File types | 
| Files and source code extensions | 

and much more along that list. Frizz also gives you handy information about the current server database such as the json files, json output, json parsers, lists, text files and even debugging logs for the server. Frizz even has a tab for hardware informations and documentation! 

The cool thing about frizzed is that it even allows you to take a giant list of pcap files, mash them together into one giant file, parse them and load them together. This makes it easier for you to parse thousands and thousands of lines of code without having to manually open them or create a masher yourself. Frizz can even be run offline, as long as there is a possible connection to your localhost you are fine.


# Beta notes 

























































# Comparing 



This project was almost a direct inspiration from apackets. I felt with how closed sourced and over priced apackets were for its minimal features i can do better and charge less while also partly making it open source. Here is a simple compare of apackets to frizz 

| Apacket status and information | Frizz features, status and information | 
| ------------------------------ | -------------------------------------- | 
| Limited packet uploading and parsing limit | No limit to packet parsing or uploading | 
| Upwards to 200$ a month for extra storage space but no features | Free and open source ( for now but will never be 200$ a month ) | 
| Only supports 19 packet parsing tabs | Supports over 35+ Protocols and individual tabs | 
| Unclean and sometimes messy and gltichy dashboard | Cleaner, more modern and performant dashboard | 
| Slower at packet parsing sometimes can take 15 to 20 seconds on average sometimes a full minute | Parses up to 300,000 packets of different layers in under a few seconds | 
| Pre generated templates | New and freshly imported and generated templates | 
| No packet mashing | Allows users to take a list of pcap files, mash them into one and even output the results | 
| Need to be online or have network to run it | Do not need to have access to network connections to run it, runs on localhost | 
| Unsafe and sometimes laggy | Safer with handleing exceptions for networking information | 
| Does not allow users to customize background of dashboard | Allows users to customize the colors and color code sof the dashboard | 
| Does not allow custom user filepath or pre settings with YAML or even JSON | Allows users to pre specify server verbosity options, load server config files and more | 
| Older and less stable | More modern and appealing | 
| No signature checking, older databases of OUI's etc | Constantly updated port service names databases, OUI databases, signature databases for either payloads or regex to verify certain packet signatures | 
| Not so specific for data | Has tabs specified for certain information such as SIP invites, SSH messages and codes, FTP messgaes and code, certain authentication, statistics, HTTP parts, hostnames, servers, ports, POP3 IP's, MACS, Emails, To and FROM | 


# Implimentation 

| Feature to impliment in other versions | Paid or no | Description | 
| -------------------------------------- | ---------- | ----------- | 
| Export data in XML, HTML, JSON and YAML | No | Allows the user to export certain data as XML, JSON, HTML table, YAML or even CSV in the further future 
| Databases such as GEO-IP, GEO-MAC, MAC_OUI, Phone region, Net info, Port | Yes and no | Some of these databases will be implimneted for free but will not have full use ability unless the license is free | 
| File checks | Yes | When a file is downloaded, reconstructed or re wired from certain images and certain points this will need to be paid for, basically this feature will take those same files and scan the files for a few things such as its header, viruses, hidden files such as archive files, hidden messages, base64 encoded strings, base32 encoded strings, AES encoded strings etc | 
| Geo MAP of all connections or geoIPs | yes | this is a feature that will be supper hard to impliment given the JS needed will need to be pre generated so will the phone numbers and logos, however this can still easily be done but will be a paid feature due to the work | 
| WHOIS, Geo etc online and offline tools | yes | this will be a paid for feature because it gives the user access to premium whois tools which will drop very valuable information using certain engines and even custom made libraries like geo location, records, domain records, base records, ip information etc | 
| Email and other various data parsing databases | yes and no | throughout the entire program i want to impliment something that will make people go wow and to do that i will be implimenting some email checking databases but will only leave certain more accurate email, phone number, IP, mac etc databases up to premium users


# Future pricing 

| Price | License grade | Features | 
| ----- | ------------- | -------- | 
| $10 USD per month | basic | has extra features, allows data format exporting, more faster and better designed web interface, better and more accurate databases, more protocol support | 
| $20 USD per month | premium | Full features including geo databases, OUI databases, access to certain whois tools, access to CVE search tools using offline databases, management tools, full customization to the web interface, more access over protocol and data information, more accurate data analytics etc | 
