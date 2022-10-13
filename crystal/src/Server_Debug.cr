require "json"


module Debugger 
    class Debug 
        need = HTTP::Server::Context 
        def debug(need context)
            date_post = Time.local
            supported_methods = {
                "GET": "HTTP GET",
                "POST": "HTTP POST",
            }
            puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] CONNECTION TO   \t <#{supported_methods[context.request.method]}>--<#{context.request.path}> |"
            puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ===================== HEADERS ======================"
            puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] -- \n"
            puts context.request.headers.map {|k, v| "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] \e[38;5;20m#{k} \t = \e[38;5;164m#{v.join("; ")}" }.join("\n")
            if b = context.request.body 
                ct = context.request.headers["Content-Type"]
                if ct.includes?("json") 
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ====================== JSON REQUEST BODY ==============="
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] #{JSON.parse(b).to_pretty_json}"
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ========================== END OF BODY ============================="
                elsif ct.includes?("text") || ct.includes?("json")
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ====================== TEXT/PLAIN REQUEST BODY ================="
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] #{b}"
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ========================== END OF BODY ============================="
                else
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ====================== SIZE OF REQUEST #{context.request.headers["Content-Length"]}=================== "
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] #{b}  "
                    puts "\e[38;5;49m[\e[38;5;55mDEBUG\e[38;5;49m]\e[38;5;49m[\e[38;5;55m#{date_post}\e[38;5;49m] ========================== END OF BODY ============================="
                end
            end
        end
    end
end