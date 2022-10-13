#pragma once 
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/Layer.h>
#include <pcapplusplus/ProtocolType.h>
#include <pcapplusplus/SdpLayer.h>
#include "../../Modules/Protos/Extras/const.hpp"
#include "../../Modules/Protos/Namespace/NameSpaceFuncs.hpp"

void SDP(pcpp::Packet RP) {
    pcpp::SdpLayer* sdp = RP.getLayerOfType<pcpp::SdpLayer>();
    // right now nothing to test, function is valid and alive but not filled
    // USE_WARN: Do not use this function, there will be no value assigned
}