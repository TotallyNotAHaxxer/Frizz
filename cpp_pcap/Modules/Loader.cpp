#pragma once
#include <getopt.h>
#include <stdlib.h>
#include <iostream>
#include <stdlib.h>
#include "/usr/local/include/pcapplusplus/SystemUtils.h"
#include "/usr/local/include/pcapplusplus/Packet.h"
#include "/usr/local/include/pcapplusplus/EthLayer.h"
#include "/usr/local/include/pcapplusplus/IPv4Layer.h"
#include "/usr/local/include/pcapplusplus/TcpLayer.h"
#include "/usr/local/include/pcapplusplus/HttpLayer.h"
#include "/usr/local/include/pcapplusplus/PcapFileDevice.h"
#include "/usr/local/include/pcapplusplus/RawPacket.h"
#include "Protos/Extras/JSON.hpp"
#include "Protos/Extras/Colors.hpp"


using namespace std;

class Return_F{
    public:
        Color COL;
        std::string PcapFile; 
        pcpp::IFileReaderDevice* Ret_File() {
            	pcpp::IFileReaderDevice* reader = pcpp::IFileReaderDevice::getReader(PcapFile);
                if (!reader->open()) {
                    delete reader;
                    std::cout << "[-] ERROR: Could not open packet capture file, got error when opening" << std::endl;
                    exit(1);
                } else {
                    COL.Message = " FILE HAS BEEN OPENED ! ";
                    COL.Color_Mid = eajvmrqh;
                    COL.Color_First = mnvjvfar;
                    COL.Out();
                }
                return reader;
        }
        void Call_Json() {
            Run();
        }
};