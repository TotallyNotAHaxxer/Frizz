#pragma once
#include <iostream>
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/ProtocolType.h>
#include <pcapplusplus/Layer.h>
#include <pcapplusplus/SipLayer.h>
#include "../../Modules/Protos/Extras/const.hpp"
#include "../../Modules/Protos/Namespace/NameSpaceFuncs.hpp"

using namespace std;

// /home/xea43p3x/Desktop/Projects/frizz/src/PCAP/Pcap_Examples/SIP_CALL_RTP_G711.pcap
void SIP(pcpp::Packet RP) {
    pcpp::SipLayer* PSI = RP.getLayerOfType<pcpp::SipLayer>();
    pcpp::SipResponseLayer* RPSI = RP.getLayerOfType<pcpp::SipResponseLayer>();
    if (PSI != NULL) {
        cout << "[+] Parsing SIP....." << endl;
        //cout << RPSI->getFirstLine() << endl; SEGFAULT
         // Nothing here yet, do not run function
    }
}

