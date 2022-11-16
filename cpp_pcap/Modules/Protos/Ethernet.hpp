#pragma once
#include <iostream>

#include <pcapplusplus/Layer.h>
#include <pcapplusplus/EthLayer.h>
#include <pcapplusplus/SystemUtils.h>
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/PcapFileDevice.h>
#include "../../Modules/Protos/Extras/const.hpp"
#include "../../Modules/Protos/Extras/Value_Checker.hpp"
#include "../../Modules/Protos/ETH-TALK.hpp"




void Data_ETHERNET(pcpp::Packet RP) {
    pcpp::EthLayer* ETH = RP.getLayerOfType<pcpp::EthLayer>();
    if (ETH != NULL) {
        std::string dt = ETH->getSourceMac().toString();
        std::string dt1 = ETH->getDestMac().toString();
        if (dt != "" && dt != "ff:ff:ff:ff:ff:ff" && dt != "00:00:00:00:00:00") {
            if (!Value_Container::Checker(pcpps.Ethernet_SRC, ETH->getSourceMac().toString())) {
                pcpps.Ethernet_SRC.push_back(ETH->getSourceMac().toString());
            }
        }
        if (dt1 != "" && dt1 != "ff:ff:ff:ff:ff:ff" && dt != "00:00:00:00:00:00") {
            if (!Value_Container::Checker(pcpps.Ethernet_DST, ETH->getDestMac().toString())) {
                pcpps.Ethernet_DST.push_back(ETH->getDestMac().toString());   
            }
        }
        
        //if (ETH->getSourceMac().toString() != "ff:ff:ff:ff:ff:ff" && ETH->getDestMac() != "ff:ff:ff:ff:ff:ff" && ETH->getDestMac() != "00:00:00:00:00:00" && ETH->getSourceMac() != "00:00:00:00:00:00") {
        //        std::cout << "[-] Found ether mac " << std::endl;
        //        pcpps.Ethernet_SRC.push_back(ETH->getSourceMac().toString());
        //        pcpps.Ethernet_DST.push_back(ETH->getDestMac().toString());   
        //        std::cout << "Data" << std::endl; 
        //}
    }
}
