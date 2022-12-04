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
require "./src/Server_Router.cr"
require "./src/Server_Loader.cr"
require "./src/Server_ECR_Loader.cr"


puts "\x1b[H\x1b[2J\x1b[3J"

def banner 
    puts "\e[38;5;242m"
    puts <<-'EOF'
         ______   ______     __     ______     ______    
        /\  ___\ /\  == \   /\ \   /\___  \   /\___  \   
        \ \  __\ \ \  __<   \ \ \  \/_/  /__  \/_/  /__  
         \ \_\    \ \_\ \_\  \ \_\   /\_____\   /\_____\ 
          \/_/     \/_/ /_/   \/_/   \/_____/   \/_____/                     
    EOF
end

puts banner

def main 
    runner = ShellCMD::Loader.new
    runner.exec("./parser #{ARGV[0]} f f f f f") # leaving empty arguments for false use of no flags this is bad....really bad
    application = Server::Base.new
    application_settings = JParse::J.new
    application.process "/" do "" end
    application.process "/example" do  "" end
    application.process "/HTML" do "" end
    application.process "/Future" do "" end
    application.process "/data" do "" end
    application.process "/Useragents.html" do "" end
    application.process "/JSONDB.html" do "" end
    application.process "/ServerRequirements.html" do "" end
    application.process "/ServerInfo.html" do "" end
    application.process "/AppInfo.html" do "" end
    application.process "/Masher.html" do "" end
    application.process "/Extractor.html" do "" end
    application.process "/ParseNew.html" do "" end
    application.process "/Documentation.html" do "" end
    application.process "/AuthIMAP.html"  do "" end
    application.process "/AuthFTPCreds.html"  do "" end
    application.process "/AuthSSHCreds.html"  do "" end
    application.process "/AuthSMTP.html"  do "" end
    application.process "/AuthDigest.html"  do "" end
    application.process "/AuthBASIC.html"  do "" end
    application.process "/AuthNTLM.html"  do "" end
    application.process "/AuthNegotiation.html"  do "" end
    application.process "/Telnet.html" do "" end
    application.process "/FTP.html" do "" end
    application.process "/SMTP.html" do "" end
    application.process "/SSH.html" do "" end
    application.process "/AuthTelnet.html" do "" end
    application.process "/Cc.html" do "" end
    application.process "/From.html" do "" end
    application.process "/Recv.html" do "" end
    application.process "/POP3" do "" end
    application.process "/SIP.html" do "" end
    application.process "/Emails.html" do "" end
    application.process "/Hostnames.html" do "" end
    application.process "/Home.html" do "" end
    application.process "/URLs.html" do "" end
    application.process "/Wifi.html" do "" end
    application.process "/WifiOspf.html" do "" end
    application.process "/ARP.html" do "" end
    application.process "/OpenPorts.html" do "" end
    application.process "/Ethernet.html" do "" end
    application.process "/Raw.html" do "" end
    application.process "/HTTPSESSION.html" do "" end    
    application.process "/Servers.html" do "" end
    application.run(application_settings.preproc, true)
end


main
