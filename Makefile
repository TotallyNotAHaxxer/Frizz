include Modules/MakeConf/PCPP.mk

all:
	g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o main.o C++/main.cpp
	g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o main main.o $(PCAPPP_LIBS)
	rm main.o

clean:
	rm main.o
	rm main