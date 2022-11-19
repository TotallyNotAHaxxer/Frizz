#pragma once
#include <iostream>
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/Layer.h>
#include <pcapplusplus/ProtocolType.h>
#include <pcapplusplus/PPPoELayer.h>
// Modules
#include "../../Modules/Protos/Extras/const.hpp"

void PPP(pcpp::Packet RP) {
    pcpp::PPPoELayer* ppoe = RP.getLayerOfType<pcpp::PPPoELayer>();
    if (ppoe != NULL) {
        // dont use
    }
}