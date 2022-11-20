#pragma once
#include <stdlib.h>
#include <iostream>
#include <fstream>
#include <sstream>

std::string LLSTR(pcpp::LinkLayerType ll) {
	if (ll == pcpp::LINKTYPE_ETHERNET) {
		return "Ethernet";
	}
	if (ll == pcpp::LINKTYPE_IEEE802_5) {
		return "IEEE 802.5 Token Ring";
	}
	else if (ll == pcpp::LINKTYPE_LINUX_SLL) {
		return "Linux cooked capture";
	}
	else if (ll == pcpp::LINKTYPE_NULL) {
		return "Null/Loopback";
	}
	else if (ll == pcpp::LINKTYPE_RAW || ll == pcpp::LINKTYPE_DLT_RAW1 || ll == pcpp::LINKTYPE_DLT_RAW2) {
		std::ostringstream stream;
		stream << "Raw IP (" << ll << ")";
		return stream.str();
	}
	std::ostringstream stream;
	stream << (int)ll;
	return stream.str();
}