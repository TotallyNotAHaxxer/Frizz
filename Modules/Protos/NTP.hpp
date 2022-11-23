#pragma once
#include <iostream>
#include <pcapplusplus/Packet.h>
#include <pcapplusplus/Layer.h>
#include <pcapplusplus/ProtocolType.h>
#include <pcapplusplus/NtpLayer.h>

// Modules

#include "../../Modules/Protos/Extras/const.hpp"


// Simply just push the key | EX: 00000001
void NTP(pcpp::Packet RP) {
    pcpp::NtpLayer* ntp = RP.getLayerOfType<pcpp::NtpLayer>();
    if (ntp != NULL) {
       pcpps.NTP_KET.push_back(ntp->getKeyID());
    }
}