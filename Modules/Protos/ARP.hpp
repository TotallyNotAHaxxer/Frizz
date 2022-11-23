#pragma once 
#include <iostream>
#include <pcapplusplus/ArpLayer.h>
#include <pcapplusplus/Packet.h>
#include "../../Modules/Protos/Extras/const.hpp"
#include "../../Modules/Protos/Extras/Value_Checker.hpp"

using namespace std;

void ARP(pcpp::Packet RP) {
    pcpp::ArpLayer* PRA = RP.getLayerOfType<pcpp::ArpLayer>();
    if (PRA != NULL) {
        if (
            PRA -> getTargetMacAddress().toString() != "00:00:00:00:00:00"
                    &&
                        PRA -> getTargetMacAddress().toString() != "ff:ff:ff:ff:ff:ff"
                                    && 
                                        PRA->getSenderMacAddress().toString() != "ff:ff:ff:ff:ff:ff"
                                                && 
                                                    PRA -> getSenderMacAddress().toString() != "00:00:00:00:00:00"
        ) {
                                                    if (
                                                         !Value_Container::Checker(pcpps.ARP_TARGET_MAC, PRA->getTargetMacAddress().toString()) 
                                                    ) {
                                                                pcpps.ARP_TARGET_MAC.push_back(PRA->getTargetMacAddress().toString());
                                                    }
                                                                if (
                                                                        !Value_Container::Checker(pcpps.ARP_SRC_MAC, PRA->getSenderMacAddress().toString())
                                                                ) {
                                                                        pcpps.ARP_SRC_MAC.push_back(PRA->getSenderMacAddress().toString());
                                                                }
                                                                            if (
                                                                                !Value_Container::Checker(pcpps.ARP_SRC_IP, PRA->getSenderIpAddr().toString())
                                                                            ) {
                                                                                pcpps.ARP_SRC_IP.push_back(PRA->getSenderIpAddr().toString());

                                                                            }
                                                                                        if (
                                                                                            !Value_Container::Checker(pcpps.ARP_TARGET_IP, PRA->getTargetIpAddr().toString())
                                                                                        ) {
                                                                                             pcpps.ARP_TARGET_IP.push_back(PRA->getTargetIpAddr().toString());
                                                                                        }
        }
    }
}

