require "./src/Server_Router.cr"


def main 
    application = Server::Base.new
    application_settings = JParse::J.new
    puts "Host - > http://localhost:8080"
    application.process "/" do "" end
    application.process "/example" do  "" end
    application.process "/HTML" do "" end
    application.process "/Future" do "" end
    application.run(application_settings.preproc, true)
end


main
