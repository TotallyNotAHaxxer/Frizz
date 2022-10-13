#include "Modules/Loader.cpp"
#include "Modules/Protos/Extras/const.hpp"
#include <iostream>



using namespace std;


int main(int argc, char* argv[]) {
    Return_F PCAP_Reader;
    PCAP_Reader.PcapFile = argv[1];
    pcpp::IFileReaderDevice* read = PCAP_Reader.Ret_File();
    if (read == NULL) {
        std::cerr << "[ERROR] Could not determine File type - " << std::endl;
        return 1;
    } else {
        std::string FT;
        FT = File_Type(read);
        pcpps.File_Link_Layer_Type = FT;
        Packet_INIT(read);
        PCAP_Reader.Call_Json(); 
    }
}