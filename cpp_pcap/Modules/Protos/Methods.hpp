#pragma once
#include <pcapplusplus/RawPacket.h>

// Modules
#include "../../Modules/Protos/Ethernet.hpp"
#include "../../Modules/Protos/HTTP.hpp"
#include "../../Modules/Protos/SDP.hpp"
#include "/usr/local/include/pcapplusplus/PcapFileDevice.h"
#include "../../Modules/Protos/Namespace/NameSpaceFuncs.hpp"
#include "../../Modules/Protos/PTPETH.hpp"
#include "../../Modules/Protos/SIP.hpp"
#include "../../Modules/Protos/ARP.hpp"
#include "ETH-TALK.hpp"

int Parse_Layers(pcpp::RawPacket PKT, pcpp::IFileReaderDevice* reader) {
    while (reader->getNextPacket(PKT)) {
    	pcpp::Packet parsedPacket(&PKT);
        //DATA_HTTP(parsedPacket);
        //Data_ETHERNET(parsedPacket);
        //Talk_SRC_TO_DST(parsedPacket);
       // ARP(parsedPacket);
    }
    return 0;
}