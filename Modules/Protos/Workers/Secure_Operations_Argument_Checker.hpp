/*
# ................................................................................................................................................................................................................................................................................................................................................................................................................................................................................
# : 
# : 
# : 
# : _____/\\\\\\\\\\\___________________________________________________________________________/\\\\\\\\\\\_________________________________________________________________________________________________________/\\\\\\\\\\\\_________________________________________________/\\\\\\_______________________________________________________________________________________________        
# : ___/\\\/////////\\\_______________________________________________________________________/\\\/////////\\\______________________________________________________________________________________________________\/\\\////////\\\______________________________________________\////\\\_______________________________________________________________________________________________       
# : __\//\\\______\///_______________________________________________________________________\//\\\______\///____________________________________________________________/\\\_____/\\\_________/\\\__/\\\___________\/\\\______\//\\\________________________________________________\/\\\____________________/\\\\\\\\\_______________________________________________________/\\\______      
# :   ___\////\\\_____________/\\\\\\\\__/\\\\\\\\\_____/\\/\\\\\\\______/\\\\\\\\______________\////\\\_____________/\\\\\\\\______/\\\\\\\\__/\\\____/\\\__/\\/\\\\\\\__\///___/\\\\\\\\\\\___\//\\\/\\\____________\/\\\_______\/\\\_____/\\\\\\\\___/\\\____/\\\_____/\\\\\\\\_____\/\\\________/\\\\\_____/\\\/////\\\____/\\\\\__/\\\\\_______/\\\\\\\\___/\\/\\\\\\____/\\\\\\\\\\\_     
# :    ______\////\\\________/\\\//////__\////////\\\___\/\\\/////\\\___/\\\/////\\\________________\////\\\________/\\\/////\\\___/\\\//////__\/\\\___\/\\\_\/\\\/////\\\__/\\\_\////\\\////_____\//\\\\\_____________\/\\\_______\/\\\___/\\\/////\\\_\//\\\__/\\\____/\\\/////\\\____\/\\\______/\\\///\\\__\/\\\\\\\\\\___/\\\///\\\\\///\\\___/\\\/////\\\_\/\\\////\\\__\////\\\////__    
# :     _________\////\\\____/\\\___________/\\\\\\\\\\__\/\\\___\///___/\\\\\\\\\\\____________________\////\\\____/\\\\\\\\\\\___/\\\_________\/\\\___\/\\\_\/\\\___\///__\/\\\____\/\\\__________\//\\\______________\/\\\_______\/\\\__/\\\\\\\\\\\___\//\\\/\\\____/\\\\\\\\\\\_____\/\\\_____/\\\__\//\\\_\/\\\//////___\/\\\_\//\\\__\/\\\__/\\\\\\\\\\\__\/\\\__\//\\\____\/\\\______   
# :      __/\\\______\//\\\__\//\\\_________/\\\/////\\\__\/\\\_________\//\\///////______________/\\\______\//\\\__\//\\///////___\//\\\________\/\\\___\/\\\_\/\\\_________\/\\\____\/\\\_/\\___/\\_/\\\_______________\/\\\_______/\\\__\//\\///////_____\//\\\\\____\//\\///////______\/\\\____\//\\\__/\\\__\/\\\_________\/\\\__\/\\\__\/\\\_\//\\///////___\/\\\___\/\\\____\/\\\_/\\__  
# :      _\///\\\\\\\\\\\/____\///\\\\\\\\_\//\\\\\\\\/\\_\/\\\__________\//\\\\\\\\\\___________\///\\\\\\\\\\\/____\//\\\\\\\\\\__\///\\\\\\\\_\//\\\\\\\\\__\/\\\_________\/\\\____\//\\\\\___\//\\\\/________________\/\\\\\\\\\\\\/____\//\\\\\\\\\\____\//\\\______\//\\\\\\\\\\__/\\\\\\\\\__\///\\\\\/___\/\\\_________\/\\\__\/\\\__\/\\\__\//\\\\\\\\\\_\/\\\___\/\\\____\//\\\\\___ 
# :       ___\///////////________\////////___\////////\//__\///____________\//////////______________\///////////_______\//////////_____\////////___\/////////___\///__________\///______\/////_____\////__________________\////////////_______\//////////______\///________\//////////__\/////////_____\/////_____\///__________\///___\///___\///____\//////////__\///____\///______\/////____
# :
# : 
# :  This code was developed by the scare security development team and cyber security group. This code is licensed under the MIT license which means you are free to distribute this code
# :  
# :  As long as you modify the data within this file. It would be nice to also let you know that if you distribute this code without any changes or modifications you should give credits to the organization seen below
# : 
# :                                                 https://github.com/Scare-Security
# : 
# :...............................................................................................................................................................................................................................................................................................................................................................................................................................................................................
*/
#pragma once
#include <iostream>
#include <vector>
#include <iomanip>
bool statement_Stat;

struct FileExtensions  {
    std::vector<std::string> Lextensions; // List of extensions
};
class Security_Of_Arguments {
    private:
        FileExtensions exten;
    public:
        bool Check_Command_Before_Constructor(int typeof, std::string Constructor) { // bool will be the verification of the parameter of the command and builder, if the command was safe then it is good to go on and will return true
            auto LT = std::time(nullptr);
            switch (typeof) {
                case 1:
                   // Mode: - FileInput | - FileOutput
                   std::ifstream streamer;
                   streamer.open(Constructor.c_str());
                            std::cout << "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m][" << std::put_time(std::gmtime(&LT), "\e[38;5;55m%F %T  %z\e[38;5;49m]") << "[\e[38;5;39mSecurity\e[38;5;49m] \e[38;5;55m| \e[38;5;20mChecking filename " << Constructor << std::endl;

                   if (streamer) {
                        streamer.close();
                            std::cout << "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m][" << std::put_time(std::gmtime(&LT), "\e[38;5;55m%F %T  %z\e[38;5;49m]") << "[\e[38;5;39mSecurity\e[38;5;49m] \e[38;5;55m| \e[38;5;20mFile found and opened \e[38;5;39m(\e[38;5;55m1/2\e[38;5;39m)"  << std::endl;
                        exten.Lextensions.push_back(".pcap");
                        exten.Lextensions.push_back(".cap");
                        exten.Lextensions.push_back(".pcapng");
                        if (fex(Constructor.c_str(), Constructor) == "PASSED") {
                            std::cout << "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m][" << std::put_time(std::gmtime(&LT), "\e[38;5;55m%F %T  %z\e[38;5;49m]") << "[\e[38;5;39mSecurity\e[38;5;49m] \e[38;5;55m| \e[38;5;20mFile verification passed \e[38;5;39m(\e[38;5;55m2/2\e[38;5;39m)"  << std::endl;
                            statement_Stat = true;
                        }
                   } else {
                        std::cout << "File does not exist !!!!!!!!!!" << std::endl;
                        statement_Stat = false;
                        streamer.close();
                   }
            }
            return statement_Stat;
        }

        // File and other system security functions 
        std::string fex(const char* fname, const std::string &File) {
            std::string TestMsg; // Resulting message
            for ( // Check size
                std::size_t k = 0; *(fname + k); k++
            ) {
                if (*(fname + k) == '\\') { 
                    TestMsg = "FAIL";
                    std::cout << "[-] Error: This is a windows filepath, windows is not a supported operating system in frizz .01 BETA. Please make sure the path does includes '/' instaead of '\\' " << std::endl;
                } 
                if (*(fname + k) == '/') { // Verify path is linux
                    // continue checks
                    unsigned int sz = exten.Lextensions.size();
                    std::string Expected_Extension;
                    for(unsigned int i = 0; i < sz; i++) {
                        Expected_Extension = exten.Lextensions[i];
                        if (
                            Expected_Extension.length() <= File.length() && std::equal(Expected_Extension.rbegin(), Expected_Extension.rend(), File.rbegin())
                        ) {
                            TestMsg="PASSED";
                            break;
                        } else {
                            TestMsg="FAILED";
                        }
                    }
                }
            }
            return TestMsg;
        }
}; 