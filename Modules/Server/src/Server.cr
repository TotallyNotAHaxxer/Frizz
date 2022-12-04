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
require "http/server"
require "ecr"
require "./Server_Responder.cr"
require "./Server_Router.cr"
require "./Server_Debug.cr"
require "./Template.cr"
require "./Server_File_Info.cr"
require "./Server_Loader.cr"
require "./Server_System.cr"


module Server 
    class Base 
        def initialize
            @router = {} of String => ( -> String )
        end
        def run(int port, bool debug)
            pcapfout = ""
            listpcapf = ""
            pcapfileinputlineforsearch = ""
            pcapfileinputforsearch = ""
            newpcapfiletocheck = ""
            date_post = Time.local
            server_Debug_Finf = Debugger_Utils_File_Info::Debugger_Utils_File_Info_Finf.new
            server = HTTP::Server.new do |ctx|
                if debug 
                    server_Debug = Debugger::Debug.new
                    server_Debug.debug(ctx)
                end
                ctx.response.content_type = "multipart/form-data"
                ctx.response.content_type = "boundary=aA40"
                req = ctx.request
                val = req.path.to_s
                if @router.has_key?(req.path.to_s)
                    if req.path.to_s == "/Raw.html"
                        iod = ParserMessage.new("This file may have been too big to process on the server. This must mean the data stored in the file is way too large or the packet data is too long. This may be a bug for larger pcap files")
                        fmodule = Files::FileWritersAndReaders.new
                        indev = "Modules/Server/ErrorBased/Error.html"
                        File.touch(indev)
                        fmodule.fwrite(indev, iod)
                        File.open(indev) do |f|
                            IO.copy(f, ctx.response)
                        end
                        File.delete(indev)
                    end
                    if req.path.to_s == "/SIP.html"
                        iod =  ParserMessage.new("Sorry this file is out of order or not generated | This may be due to the fact that this file was not ready by release")
                        fmodule = Files::FileWritersAndReaders.new
                        indev = "Modules/Server/ErrorBased/InDev.html"
                        File.touch(indev)
                        fmodule.fwrite(indev, iod)
                        File.open(indev) do |f|
                            IO.copy(f, ctx.response)
                        end
                        File.delete(indev)
                    end
                    if req.method == "GET"
                        aoo = Responder::Switcher.new 
                        fo = aoo.return_base_path(req.path.to_s)
                        server_Debug_Finf.get_file_inf(req.path.to_s) 
                        if File.exists?(fo.to_s)
                            ctx.response.content_length = File.size(fo.to_s)
                            File.open(fo.to_s) do |f|
                                IO.copy(f, ctx.response)
                            end
                    end 
                    else if req.method == "POST"
                            HTTP::FormData.parse(ctx.request) do |part|
                                case part.name
                                    when "list" # first agrument for packet mashing
                                        listpcapf = part.body.gets_to_end
                                    when "pcapout" # second argument for packet mashing
                                        pcapfout = part.body.gets_to_end
                                    when "inputpcapfile"
                                        pcapfileinputforsearch = part.body.gets_to_end
                                    when "inputlinenum"
                                        pcapfileinputlineforsearch = part.body.gets_to_end
                                    when "checkinpcap"
                                        newpcapfiletocheck = part.body.gets_to_end
                                end
                            end
                            if newpcapfiletocheck != ""
                                puts "[TOOL] Using: New loader | Load new pcap files onto the interface"
                                runner = ShellCMD::Loader.new
                                runner.exec("././parser #{newpcapfiletocheck} f f f f f")
                                iod = ParserMessage.new("Pcap file #{newpcapfiletocheck} has been sucessfully loaded into the database, you can now re load the web server or go back to the root directory to see you're results")
                                msgfile = "Modules/Server/ErrorBased/Message.html"
                                fmodule = Files::FileWritersAndReaders.new
                                File.touch(msgfile)
                                fmodule.fwrite(msgfile, iod)
                                File.open(msgfile) do |f|
                                    IO.copy(f, ctx.response)
                                end
                                File.delete(msgfile)
                            end
                            if pcapfout != "" && listpcapf != "" 
                                # Run the pcap file
                                puts "[TOOL] Using -> Masher - Mashing file"
                                runner = ShellCMD::Loader.new
                                runner.exec("./parser ... masher #{listpcapf} #{pcapfout}") # mash the packet and return status
                                iod = ParserMessage.new("PCAP file #{pcapfout} has been saved, you can now go back to the server ( This file was not parsed, you must either upload a new file or restart the program to load the data of this file )")
                                msgfile = "Modules/Server/ErrorBased/Message.html"
                                fmodule = Files::FileWritersAndReaders.new
                                File.touch(msgfile)
                                fmodule.fwrite(msgfile, iod)
                                File.open(msgfile) do |f|
                                    IO.copy(f, ctx.response)
                                end
                                File.delete(msgfile)
                            else if pcapfileinputlineforsearch != "" && pcapfileinputforsearch != ""
                                # run the parser
                                puts "[TOOL] EXTRACTOR | Extract a certain line from a file into another"
                                runner = ShellCMD::Loader.new
                                cmd = "./parser ... lineextract #{pcapfileinputforsearch} #{pcapfileinputlineforsearch}"
                                puts cmd
                                runner.exec(cmd)
                                iod = ParserMessage.new("Generated a pcaop file with name [output.pcap] in the current filepath ( sorry right now this file is static :( the name of the pcap file in later versions will definitely be changed )")
                                msgfile = "Modules/Server/ErrorBased/Message.html"
                                fmodule = Files::FileWritersAndReaders.new
                                File.touch(msgfile)
                                fmodule.fwrite(msgfile, iod)
                                File.open(msgfile) do |f|
                                    IO.copy(f, ctx.response)
                                end
                                File.delete(msgfile)
                            end
                        end
                        end
                    end
                else
                    fourofourfile = "Modules/Server/ErrorBased/404.html"
                    fmodule = Files::FileWritersAndReaders.new
                    
                    
                    iod = Parser404.new(
                        "/Home.html",
                        "/PacketStats.html",
                        "/Useragents.html",
                        "/Hostnames.html",
                        "/URLs.html",
                        "/DNS.html",
                        "/OpenPorts.html",
                        "/ARP.html",
                        "/Ethernet.html",
                        "/Servers.html",
                        "/Wifi.html",
                        "/WifiOspf.html",
                        "/FTP.html",
                        "/SSH.html",
                        "/Telnet.html",
                        "/SMTP.html",
                        "/AuthFTPCreds.html",
                        "/AuthSSHCreds.html",
                        "/AuthIMAP.html",
                        "/AuthDigest.html",
                        "/AuthNTLM.html",
                        "/AuthBASIC.html",
                        "/AuthNegotiation.html",
                        "/AuthSMTP.html",
                        "/Emails.html",
                        "/Cc.html",
                        "/From.html",
                        "/Recv.html",
                        "/Convos.html",
                        "/Masher.html",
                        "/Raw.html",
                        "/ServerRequirements.html",
                        "/JSONDB.html",
                        "/AppInfo.html",
                        "/ServerInfo.html",
                        "/Documentation.html",
                        message: "404 Page or filepath not found"
                    ).to_s
                    File.touch(fourofourfile)
                    fmodule.fwrite(fourofourfile, iod)
                    File.open(fourofourfile) do |f|
                        IO.copy(f, ctx.response)
                    end
                    File.delete(fourofourfile)
                    puts "\e[38;5;49m[\e[38;5;124mERROR\e[38;5;49m]\e[38;5;49m[\e[38;5;124m#{date_post}\e[38;5;49m] \e[38;6;124m Router does not support the filepath #{req.path.to_s}"
                end
            end
            server.listen(port)
        end
        def process(router, &block : (-> String))
            @router[router.to_s] = block
        end
    end
end