#pragma once
#include <iostream>
#include <pcapplusplus/PcapFileDevice.h>
#include <pcapplusplus/RawPacket.h>
#include <string>
#include <vector>
#include <fstream>
#include <iomanip>
#include "../../../Modules/Protos/JSON_HEAD/json.hpp"
#include "../../Protos/Methods.hpp"
#include "../../Protos/Files/FileType.hpp"
#include "../../Protos/Extras/Structure.hpp"
#include "../../Protos/Extras/const.hpp"



using nlohmann::json;




void to_json(json & j,
    const PCPP_Packet_Information & t) {
    j = json {
        {
                "Total_Parsed_Packets",
                t.Packets_Parsed
        }, {
                "Total_Application_Packets",
                t.Application_Layers
        }, {
                "Total_Application_Cookies_HTTP",
                t.HTTP_Cookies
        }, {
                "Total_Application_GET_HTTP",
                t.HTTP_GET
        }, {
                "Total_Application_PUT_HTTP",
                t.HTTP_PUT
        }, {
                "Total_Application_POST_HTTP",
                t.HTTP_POST
        }, {
                "Total_Application_DELETE_HTTP",
                t.HTTP_DELETE
        }, {
                "Total_Application_CONNECT_HTTP",
                t.HTTP_CONNECT
        }, {
                "Total_Application_HEAD_HTTP",
                t.HTTP_HEAD
        }, {
                "Total_Application_OPTIONS_HTTP",
                t.HTTP_OPTIONS
        }, {
                "Total_Application_PATCH_HTTP",
                t.HTTP_PATCH
        }, {
          "     Total_Application_TRACE_HTTP",
                t.HTTP_TRACE
        }, {
                "Total_Application_REQUESTS_HTTP",
                t.HTTP_TOTAL_REQUESTS
        }, {
                "Total_Application_RESPONSES_HTTP",
                t.HTTP_TOTAL_RESPONSES
        }, {
                "File Link Layer Type",
                t.File_Link_Layer_Type
        }, {
                "URI(s) ",
                t.HTTP_URL_PATHS
        }, {
                "Request(s) ",
                t.HTTP_REQ_URI
        }, {
                "Useragent(s) ",
                t.HTTP_USERAGENTS
        }, {
                "URL(s) ",
                t.HTTP_HP
        }, {
                "ETHERNET (source) MAC's ",
                t.Ethernet_SRC
        }, {
                "ETHERNET (destination) MAC's ",
                t.Ethernet_DST
        }, {
                "ETHERNET TALK SRC TO DST",
                t.ETH_TALK
        }, {
                "Address Resolution Protocol Talk ( SRC IP )",
                t.ARP_SRC_IP
        }, {
                "Address Resolution Protocol Talk ( DST IP ) ",
                t.ARP_TARGET_IP
        },  {
                "Address Resolution Protocol Talk ( SRC MAC ) ",
                t.ARP_SRC_MAC
        },  {
                "Address Resolution Protocol Talk ( DST MAC ) ",
                t.ARP_TARGET_MAC
        },  {
                "Address Resolutin Protocol Talk ( Full convo ) ",
                t.ARP_FULL_TALK
        }
    };
}

int Run(){
    json outout;
    outout.push_back(pcpps); // push back the PCPP information HTTP structure
    std::ofstream o("Info.json");
    o << std::setw(10) << outout << std::endl;
    return 0;
}

void Packet_INIT(pcpp::IFileReaderDevice* READ) {
    pcpp::RawPacket Raw;
    Parse_Layers(Raw, READ);
}

