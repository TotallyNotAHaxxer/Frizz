include Modules/MakeConf/PCPP.mk

all:
	g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o capture.o C++/main.cpp
	g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o capture capture.o $(PCAPPP_LIBS)
	go build -o parser main.go 
	crystal build  Modules/Server/router.cr
	rm capture.o

# As of right now do not erase all source code files, when the user installs then erase all source code files and just keep plain directories

clean:
	rm main.o
	rm main