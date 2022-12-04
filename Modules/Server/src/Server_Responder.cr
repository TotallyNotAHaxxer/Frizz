#
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
# : This code belongs to the scare security organization. This code you may use and may re distribute as this is a open source project as a base for modern day applications which need simple local HTTP routing
# :
# : Please make sure before distributing code is not directly pasted and this code is taken into affect, modified, remade, re named, re created, re modified and other forms of manipulation are required
# : 
# : If you do not modify make sure to give credits to the contributors and the security organization that developed this code
# : 
# :                 https://github.com/Scare-Security/Safe-And-Secure-Router
# : 
# :...............................................................................................................................................................................................................................................................................................................................................................................................................................................................................
#
#
#
#
# Server concept: All HTLM code and requsts must be to a direct file in the HTML that is handled logically to the server
#
# if the HTML sidebar has Modules/Server/HTML/example.html as a link it will not process it and it will give you a fake page 
#
# this is because the server is set to directly handle direct filepaths after the port number not direct module paths.
#
require "./Template.cr"

filename = "example.html"

module Responder 
    class Switcher
        def return_base_path(path_of_file : String)
            server_Locations = {
                "/example" => "Modules/Server/Static/example.html",
                "/"        => "Modules/Server/HTML/Home.html",
                "/data" => "Modules/Server/Static/Data.html",
                "/Useragents.html" => "Modules/Server/HTML/Useragents.html",
                "/JSONDB.html" => "Modules/Server/HTML/JSONDB.html",
                "/ServerRequirements.html" => "Modules/Server/HTML/ServerRequirements.html",
                "/ServerInfo.html" => "Modules/Server/HTML/ServerInfo.html",
                "/AppInfo.html" => "Modules/Server/HTML/AppInfo.html",
                "/Masher.html" => "Modules/Server/HTML/Masher.html",
                "/Extractor.html" => "Modules/Server/Static/Extractor.html",
                "/ParseNew.html" => "Modules/Server/Static/ParseNew.html",
                "/Documentation.html" => "Modules/Server/Static/Documentation.html",
                "/AuthIMAP.html" => "Modules/Server/HTML/AuthIMAP.html",
                "/AuthFTPCreds.html" => "Modules/Server/HTML/AuthFTPCreds.html",
                "/AuthSSHCreds.html" => "Modules/Server/HTML/AuthSSHCreds.html",
                "/AuthSMTP.html" => "Modules/Server/HTML/AuthSMTP.html",
                "/AuthDigest.html" => "Modules/Server/HTML/AuthDigest.html",
                "/AuthBASIC.html" => "Modules/Server/HTML/AuthBASIC.html",
                "/AuthNTLM.html" => "Modules/Server/HTML/AuthNTLM.html",
                "/AuthNegotiation.html" => "Modules/Server/HTML/AuthNegotiation.html",
                "/Telnet.html" => "Modules/Server/HTML/Telnet.html",
                "/FTP.html" => "Modules/Server/HTML/FTP.html",
                "/SSH.html" => "Modules/Server/HTML/SSH.html",
                "/SMTP.html" => "Modules/Server/HTML/SMTP.html",
                "/AuthTelnet.html" => "Modules/Server/HTML/AuthTelnet.html",
                "/Cc.html" => "Modules/Server/HTML/Cc.html",
                "/Recv.html" => "Modules/Server/HTML/Recv.html",
                "/From.html" => "Modules/Server/HTML/From.html",
                "/POP3" => "Modules/Server/HTML/Convos.html",
                "/SIP.html" => "EMPTY PATH",
                "/Emails.html" => "Modules/Server/HTML/Emails.html",
                "/Hostnames.html" => "Modules/Server/HTML/Hostnames.html",
                "/URLs.html" => "Modules/Server/HTML/URLs.html",
                "/Wifi.html" => "Modules/Server/HTML/Wifi.html",
                "/WifiOspf.html" => "Modules/Server/HTML/WifiOspf.html",
                "/ARP.html" => "Modules/Server/HTML/ARP.html",
                "/OpenPorts.html" => "Modules/Server/HTML/OpenPorts.html",
                "/Ethernet.html" => "Modules/Server/HTML/Ethernet.html",
                "/Raw.html" => "EMPTY PATH",
                "/HTTPSESSION.html" => "Modules/Server/HTML/HTTPSESSION.html",
                "/Servers.html" => "Modules/Server/HTML/Servers.html"
            }
            server_Locations[path_of_file]
        end
    end
end