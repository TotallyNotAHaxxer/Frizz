require "./Template.cr"

filename = "example.html"

module Responder 
    class Switcher
        def return_base_path(path_of_file : String)
            server_Locations = {
                "/example" => "example_File_Input/example.html",
                "/"        => "HTML/example.html",
                "/HTML" => "HTML/LobbyMisc/Lobby_Music",
                "/Future" => "HTML/Future/Future_Lobby"
            }
            server_Locations[path_of_file]
        end
    end
end