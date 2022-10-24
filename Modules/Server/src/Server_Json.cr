require "json"

module JParse 
    class J 
        def preproc
            json = File.open("Settings/Server.json") do |file|
                a = JSON.parse(file)
                dt = a["Port"].to_s
                puts "[+] Server on: #{a["URL"]}"
                dt.to_i
            end
        end
    end
end