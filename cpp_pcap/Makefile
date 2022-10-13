include /usr/local/etc/PcapPlusPlus.mk

all:
	g++ $(PCAPPP_BUILD_FLAGS) $(PCAPPP_INCLUDES) -c -o main.o main.cpp
	g++ $(PCAPPP_LIBS_DIR) -static-libstdc++ -o main main.o $(PCAPPP_LIBS)

clean:
	rm main.o
	rm main