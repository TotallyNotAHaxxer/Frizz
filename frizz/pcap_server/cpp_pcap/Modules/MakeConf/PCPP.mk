### COMMON ###

# includes
PCAPPP_INCLUDES := -I/usr/local/include/pcapplusplus


# libs
PCAPPP_LIBS := -lPcap++ -lPacket++ -lCommon++

# post build
PCAPPP_POST_BUILD :=

# build flags
PCAPPP_BUILD_FLAGS := -fPIC -std=c++20

ifdef PCAPPP_ENABLE_CPP_FEATURE_DETECTION
    PCAPPP_BUILD_FLAGS += -DPCAPPP_CPP_FEATURE_DETECTION
endif

ifndef CXXFLAGS
CXXFLAGS := -O2 -g -Wall
endif

PCAPPP_BUILD_FLAGS += $(CXXFLAGS)
### LINUX ###

# libs
PCAPPP_LIBS += -lpcap -lpthread

# allow user to add custom LDFLAGS
PCAPPP_BUILD_FLAGS += $(LDFLAGS)
