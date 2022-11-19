#pragma once
#include <iostream>
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/Layer.h>
#include <pcapplusplus/ProtocolType.h>
#include <pcapplusplus/BgpLayer.h>

// Modules

#include "../../Modules/Protos/Extras/const.hpp"


void BGP(pcpp::Packet RP) {
    pcpp::BgpLayer* bgp = RP.getLayerOfType<pcpp::BgpLayer>();
    if (bgp != NULL) {
       // Null - Do not run
    }
}