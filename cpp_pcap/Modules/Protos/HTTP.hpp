#pragma once

#include "string"
#include <cstdio>
#include <pcapplusplus/HttpLayer.h>
#include <pcapplusplus/Layer.h>
#include <pcapplusplus/TcpLayer.h>
#include <iostream>
#include "stdlib.h"
#include "SystemUtils.h"
#include "../../Modules/Protos/Extras/const.hpp"
#include "../../Modules/Protos/Extras/Structure.hpp"
#include "../../Modules/Protos/Extras/Value_Checker.hpp"




std::string Header_HTTP(pcpp::HttpRequestLayer::HttpMethod layer_method) {
    switch (layer_method) {
        case pcpp::HttpRequestLayer::HttpGET:
            pcpps.HTTP_GET++;
            return "GET";
        case pcpp::HttpRequestLayer::HttpHEAD:
            pcpps.HTTP_HEAD++;
            return "HEAD";
        case pcpp::HttpRequestLayer::HttpCONNECT:
            pcpps.HTTP_CONNECT++;
            return "CONNECT";
        case pcpp::HttpRequestLayer::HttpPOST:
            pcpps.HTTP_POST++;
            return "POST";
        case pcpp::HttpRequestLayer::HttpPUT:
            pcpps.HTTP_PUT++;
            return "PUT";
        case pcpp::HttpRequestLayer::HttpTRACE:
            pcpps.HTTP_TRACE++;
            return "TRACE";
        case pcpp::HttpRequestLayer::HttpOPTIONS:
            pcpps.HTTP_OPTIONS++;
            return "OPTIONS";
        case pcpp::HttpRequestLayer::HttpDELETE:
            pcpps.HTTP_DELETE++;
            return "DELETE";
        case pcpp::HttpRequestLayer::HttpPATCH:
            pcpps.HTTP_PATCH++;
            return "PATCH";
        default:
            return "Not found";
    }
}


bool Check_Value_If_FILL(std::string data) {
    if (data == "") {
        return false;
    } else {
        return true;
    }
}

/*
void HTTP_RESPONSE(pcpp::Packet RP) {
    pcpp::HttpResponseLayer* HL = RP.getLayerOfType<pcpp::HttpResponseLayer>();
    if (HL != NULL) {
        HL
    }
}
*/

void DATA_HTTP(pcpp::Packet RP) {
    	pcpp::HttpRequestLayer* RL = RP.getLayerOfType<pcpp::HttpRequestLayer>();
        if (RL == NULL) {
		    return;
        } else {
            std::string s1 = Header_HTTP(RL->getFirstLine()->getMethod());
            std::string s2 = RL->getFirstLine()->getUri();
            std::string s3 = RL->getFieldByName(PCPP_HTTP_HOST_FIELD)->getFieldValue();
            std::string s4 = RL->getFieldByName(PCPP_HTTP_USER_AGENT_FIELD)->getFieldValue();
            std::string s = RL->getUrl();
            bool s1b = Check_Value_If_FILL(s1);
            bool s2b = Check_Value_If_FILL(s2);
            bool s3b = Check_Value_If_FILL(s3);
            bool s4b = Check_Value_If_FILL(s4);
            bool sb  = Check_Value_If_FILL(s);
            if (s1b) {
                if (sb) {
                    std::string vals2 = s1 + "  " + s;
                    if (!Value_Container::Checker(pcpps.HTTP_REQ_URI, vals2)) {
                        pcpps.HTTP_REQ_URI.push_back(vals2);
                    }
                }
            }
            if (s2b) {
                if (!Value_Container::Checker(pcpps.HTTP_URL_PATHS, s2)) {
                    pcpps.HTTP_URL_PATHS.push_back(s2);
                }
            }
            if (s3b) {
                if (!Value_Container::Checker(pcpps.HTTP_HOSTS, s3)) {
                    pcpps.HTTP_HOSTS.push_back(s3);
                }
            }
            if (s4b) {
                if (!Value_Container::Checker(pcpps.HTTP_USERAGENTS, s4)) {
                 pcpps.HTTP_USERAGENTS.push_back(s4);

                }
            }
            if (sb) {
                if (!Value_Container::Checker(pcpps.HTTP_HP, s)) {
                    pcpps.HTTP_HP.push_back(s);
                }
            }
        }
}