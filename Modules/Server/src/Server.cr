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
require "./Server_Responder.cr"
require "./Server_Router.cr"
require "./Server_Debug.cr"
require "./Template.cr"
require "./Server_File_Info.cr"
require "./Server_Loader.cr"

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
                if @router.has_key?(req.path.to_s)
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
                                end
                            end
                            if pcapfout != "" && listpcapf != "" 
                                # Run the pcap file
                                runner = ShellCMD::Loader.new
                                runner.exec("./parser ... masher #{listpcapf} #{pcapfout}") # mash the packet and return status
                                ctx.response.print "PCAP file #{pcapfout} has been saved, you can now go back to the server ( This file was not parsed, you must either upload a new file or restart the program to load the data of this file )"
                            else if pcapfileinputlineforsearch != "" && pcapfileinputforsearch != ""
                                # run the parser
                                runner = ShellCMD::Loader.new
                                cmd = "./parser ... lineextract #{pcapfileinputforsearch} #{pcapfileinputlineforsearch}"
                                puts cmd
                                runner.exec(cmd)
                                ctx.response.print "Generated a pcaop file with name [output.pcap] in the current filepath ( sorry right now this file is static :( the name of the pcap file in later versions will definitely be changed )"
                            end
                        end
                        end
                    end
                else
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