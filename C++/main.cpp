#include "../Modules/Loader.cpp"
#include "../Modules/Protos/Extras/const.hpp"
#include "../Modules/Protos/Workers/Secure_Operations_System_Runner.hpp"
#include "../Modules/Protos/Workers/Secure_Operations_Argument_Checker.hpp"
#include <iostream>



int main(int argc, char* argv[]) {
    Return_F PCAP_Reader;
    Security_Of_Arguments Validation;
    if(argv[1] != NULL) {
        PCAP_Reader.PcapFile = argv[1];
        if(Validation.Check_Command_Before_Constructor(1, PCAP_Reader.PcapFile)) {
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
    } 
}