#pragma once


#include <pcapplusplus/Layer.h>
#include <pcapplusplus/EthLayer.h>
#include <pcapplusplus/SystemUtils.h>
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/PcapFileDevice.h>
#include "Extras/Value_Checker.hpp"
#include <pcapplusplus/Packet.h>
#include "../../Modules/Protos/Extras/const.hpp"


// literally same exact thing as the Data_Ethernet function but just concentrates strings into a conversation
void Talk_SRC_TO_DST(pcpp::Packet RP) {
       pcpp::EthLayer* ETHL = RP.getLayerOfType<pcpp::EthLayer>();
       if (ETHL != NULL) {
        std::string talk1_src = ETHL->getSourceMac().toString();
        std::string talk2_dst = ETHL->getDestMac().toString();
        if (
            talk1_src != "ff:ff:ff:ff:ff:ff" 
                    &&
                        talk1_src != "00:00:00:00:00:00"
                                && 
                                    talk1_src != ""
                                    && 
                                        talk2_dst != "ff:ff:ff:ff:ff:ff" 
                                                &&  //            ff:ff:ff:ff:ff:ff
                                                    talk2_dst != "00:00:00:00:00:00") {
                                                        std::string concentrater = talk1_src + " -> " + talk2_dst;
                                                        if (!Value_Container::Checker(pcpps.ETH_TALK, concentrater)) {
                                                            pcpps.ETH_TALK.push_back(concentrater);
                                                        }
                                                    }
       }
}