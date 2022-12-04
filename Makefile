include Modules/MakeConf/PCPP.mk
#g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o capture.o C++/main.cpp
#g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o capture capture.o $(PCAPPP_LIBS)
all:
	go build -o parser main.go 
	crystal build  Modules/Server/router.cr

# As of right now do not erase all source code files, when the user installs then erase all source code files and just keep plain directories

libs:
	sudo apt install libssl-dev      
	sudo apt install libxml2-dev     
	sudo apt install libyaml-dev   
	sudo apt install libgmp-dev      
	sudo apt install libz-dev
	sudo apt-get install libpcap-dev
go:
	go build -o parser main.go 

#cpp:
#g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o capture.o C++/main.cpp
#g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o capture capture.o $(PCAPPP_LIBS)

crystal:
	crystal build  Modules/Server/router.cr

clean:
	rm main.o
	rm main
