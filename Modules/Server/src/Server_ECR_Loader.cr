require "ecr"


class Parser404
    @filepaths : Array(String)
    def initialize(*filepaths, @message : String)
        @filepaths = filepaths.to_a
    end
  
    ECR.def_to_s "Modules/Server/ErrorBased/404.ecr"
end


