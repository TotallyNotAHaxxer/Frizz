#pragma once
#include <cstdint>
#include <iostream>

using namespace std; 

struct PCPP_Packet_Information {
    int Packets_Parsed;
    int Application_Layers;
    int HTTP_Cookies;
    int HTTP_GET;
    int HTTP_PUT;
    int HTTP_POST;
    int HTTP_DELETE; 
    int HTTP_CONNECT;
    int HTTP_HEAD;
    int HTTP_OPTIONS;                                       // HTTP total OPTIONS requests   | example - 4 HTTP ( OPTION ) REQUESTS
    int HTTP_PATCH;                                        // HTTP total PATH requests       | example - 9 HTTP ( PATCH ) Requests
    int HTTP_TRACE;                                       // HTTP total trace requests       | example - 2 HTTP ( TRACE ) Requests
    int HTTP_UNKNOWN;                                    // HTTP total unknown               | example - 12 unknown requests
    int HTTP_TOTAL_REQUESTS;                            // HTTP total requests               | example - 90 GET, 12 PUT, 13 POST
    int HTTP_TOTAL_RESPONSES;                          // HTTP Total responses               | example - 14 total response bodies
    vector< string > HTTP_URL_PATHS;             // HTTP URL PATHS                      | example - /images/layout/logo.png
    vector< string > HTTP_REQUESTS;             // HTTP Requests                        | example - Wget/1.12 (linux-gnu)
    vector< string > HTTP_HOSTS;               // HTTP Hostnames                        | example - packetlife.net
    vector< string > HTTP_HP;                 // HTTP Host path                         | example - packetlife.net/images/layout/logo.png
    vector< string > HTTP_REQ_URI;           // HTTP request + the URI or host request  | example - GET packetlife.net/images/layout/logo.png
    vector< string > HTTP_USERAGENTS;       //  HTTP UserAgents                         | example - Wget/1.12 (linux-gnu)
    vector< string > HTTP_FLAGS;           //  HTTP UserAgents                         | example - Wget/1.12 (linux-gnu)
    vector< string > FTP_Recreation;      // FTP Session recreation                    | example - 61 Request: CWD /  
    vector< string > Telnet_Recreation;  // Telnet Session recreation                  | example - TELNET: DATA: \r\n
    vector< string > Ethernet_SRC;       // Ether net source 
    vector< string > Ethernet_DST;       // Ether net Destination
    vector< string > ARP_SRC_IP;         // ARP sender IP
    vector< string > ARP_SRC_MAC;        // ARP sender mac
    vector< string > ARP_TARGET_IP;      // ARP target IP
    vector< string > ARP_TARGET_MAC;     // ARP target mac
    vector< string > ARP_FULL_TALK; // ARP all data mashed
    vector< string > ETH_TALK; // src to dst example | 00:00:00:00:00:00 -> ff:ff:ff:ff:ff:ff
    vector< int32_t> NTP_KET;            // NTP Key ID's
    string File_Link_Layer_Type;        // File link layer                             | example - Pcapfile / PCAPNG / SNOOPFILE
};

