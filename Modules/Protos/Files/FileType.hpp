#pragma once
#include <pcapplusplus/PcapFileDevice.h>
#include "Layer_Str.hpp"


std::string File_Type(pcpp::IFileReaderDevice* read) {
    std::string A;
    if (dynamic_cast<pcpp::IFileReaderDevice*>(read) != NULL) {
        pcpp::PcapFileReaderDevice* FT = dynamic_cast<pcpp::PcapFileReaderDevice*>(read);
		pcpp::LinkLayerType ll = FT->getLinkLayerType();
	    A = LLSTR(ll); // convert the link layer type to a string to be able to be read in proper format
    }  
    if (dynamic_cast<pcpp::SnoopFileReaderDevice*>(read) != NULL) {
        pcpp::SnoopFileReaderDevice* FT = dynamic_cast<pcpp::SnoopFileReaderDevice*>(read);
		pcpp::LinkLayerType ll = FT->getLinkLayerType();
	    A = LLSTR(ll);
    }
    return A;
}