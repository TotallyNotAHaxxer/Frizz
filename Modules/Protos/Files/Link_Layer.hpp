#pragma once
#include <pcapplusplus/EthLayer.h>
#include <pcapplusplus/RawPacket.h>
#include <string>

class Detect{
    public:
        string Res = "";
        string LL(pcpp::LinkLayerType Link) {
            switch(Link) {
                case pcpp::LINKTYPE_ETHERNET :
                    Res = "Ethernet";
                case pcpp::LINKTYPE_NULL : 
                    Res = "127.0.0.1 / ::1 / Loopback";
                case pcpp::LINKTYPE_DLT_RAW1 : 
                    Res = "RAW IP";
                case pcpp::LINKTYPE_DLT_RAW2 :
                    Res = "RAW IP";
            }
            return Res
        }
};
