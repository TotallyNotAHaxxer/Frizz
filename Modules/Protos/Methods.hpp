#pragma once
#include <pcapplusplus/RawPacket.h>

// Modules
#include "../../Modules/Protos/Ethernet.hpp"
#include "/usr/local/include/pcapplusplus/PcapFileDevice.h"
#include "../../Modules/Protos/PTPETH.hpp"
#include "../../Modules/Protos/ARP.hpp"
#include "ETH-TALK.hpp"
#include "../../Modules/Loader.cpp"




int Parse_Layers(pcpp::RawPacket PKT, pcpp::IFileReaderDevice* reader) {
    while (reader->getNextPacket(PKT)) {
    	pcpp::Packet parsedPacket(&PKT);
        Data_ETHERNET(parsedPacket);
        Talk_SRC_TO_DST(parsedPacket);
        ARP(parsedPacket);
    }
    return 0;
}