![](Logo.png)

# About 

Firzz is an open sourced packet parsing tool for penetration testers and digital forensics experts alike. Currently frizz uses the crystal programming language to back its servers while having a neat and extremely performant C++ and Go backend. It can parse many protocols from HTTP ( GET, POST, HOST, USERAGENTS ) to ETHERNET to even POP3. Frizz mainly focuses on finding sensative information such as emails, email conversations, ftp sessions, session logins, get requests and authorization requests, device and information while also providing you with options to view firzz's entire json database. Along with the amazing web interface there is also a neat little feature which allows you to view server information, current stats, hardware information etc.

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
