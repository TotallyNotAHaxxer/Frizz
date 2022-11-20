#pragma once
#include <cstdint>
#include <iostream>


struct PCPP_Packet_Information {
    int Packets_Parsed;
    int Application_Layers;
    std::vector< std::string > FTP_Recreation;      // FTP Session recreation                    | example - 61 Request: CWD /  
    std::vector< std::string > Telnet_Recreation;  // Telnet Session recreation                  | example - TELNET: DATA: \r\n
    std::vector< std::string > Ethernet_SRC;       // Ether net source 
    std::vector< std::string > Ethernet_DST;       // Ether net Destination
    std::vector< std::string > ARP_SRC_IP;         // ARP sender IP
    std::vector< std::string > ARP_SRC_MAC;        // ARP sender mac
    std::vector< std::string > ARP_TARGET_IP;      // ARP target IP
    std::vector< std::string > ARP_TARGET_MAC;     // ARP target mac
    std::vector< std::string > ARP_FULL_TALK; // ARP all data mashed
    std::vector< std::string > ETH_TALK; // src to dst example | 00:00:00:00:00:00 -> ff:ff:ff:ff:ff:ff
    std::vector< std::int32_t> NTP_KET;            // NTP Key ID's
    std::string File_Link_Layer_Type;        // File link layer                             | example - Pcapfile / PCAPNG / SNOOPFILE
};

