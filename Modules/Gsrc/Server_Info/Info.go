package Frizz_Server

import (
	"os"
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
)

/*
This file will grab all of the elements the server needs such as the following

Hardware Information
File System information
Operating System Information
Network Information
Access to the file system
Reader information
Kernel Information
Graphics information
Browser information such as the name
Internet addresses
*/

var X error

type Server struct {
	Server_OperatingSystem                           string
	Server_OperatingSystemFileSystem                 string
	Server_OperatingSystemVersion                    string
	Server_OperatingArchitecture                     string
	Server_HardwareInfo_CPU_VendorID                 string
	Server_HardwareInfo_CPU_IndexNum                 string
	Server_HardwareInfo_CPU_Family                   string
	Server_HardwareInfo_CPU_NumberOfCores            string
	Server_HardwareInfo_CPU_ModelName                string
	Server_HardwareInfo_CPU_Speed                    string
	Server_HardwareInfo_CPU_CacheSize                string
	Server_HardwareInfo_CPU_Micronode                string
	Server_HardwareInfo_CPU_Model                    string
	Server_HardwareInfo_CPU_PhysID                   string
	Server_HardwareInfo_CPU_Step                     string
	Server_HardwareInfo_MEM_Free                     string
	Server_HardwareInfo_MEM_Total                    string
	Server_HardwareInfo_MEM_Used                     string
	Server_HardwareInfo_OSP_Hostname                 string
	Server_HardwareInfo_OSP_Uptime                   string
	Server_HardwareInfo_OSP_ProcRunning              string
	Server_HardwareInfo_OSP_HOSTID                   string
	Server_HardwareInfo_OSP_HOSTOS                   string
	Server_HardwareInfo_OSP_HOSTPLAT                 string
	Server_HardwareInfo_OSP_HOST_KERNEL_VERSION      string
	Server_HardwareInfo_OSP_HOST_KERNEL_ARCHITECTURE string
	Server_HardwareInfo_OSP_HOST_PLATFORM_VERSION    string
	Server_HardwareInfo_OSP_HOST_PLATFORM_FAMILY     string
}

func Sav(x error) {
	if x != nil {
		os.Exit(0) // Should i exit? Or pass it to the debugger, i think i might move this later
	}
}

// Init OS Info
func (Inf *Server) OS() {
	Inf.Server_OperatingSystem = runtime.GOOS
	Inf.Server_OperatingArchitecture = runtime.GOARCH
	switch Inf.Server_OperatingSystem {
	case "linux":
		Inf.Server_OperatingSystemFileSystem = "/"
	case "windows":
		Inf.Server_OperatingSystemFileSystem = `\`
	}
}

// Init CPU Info
func (Inf *Server) CPU() {
	cs, x := cpu.Info()
	Sav(x)
	Inf.Server_HardwareInfo_CPU_IndexNum = strconv.FormatInt(int64(cs[0].CPU), 10)
	Inf.Server_HardwareInfo_CPU_VendorID = cs[0].VendorID
	Inf.Server_HardwareInfo_CPU_Family = cs[0].Family
	Inf.Server_HardwareInfo_CPU_NumberOfCores = strconv.FormatInt(int64(cs[0].Cores), 10)
	Inf.Server_HardwareInfo_CPU_ModelName = cs[0].ModelName
	Inf.Server_HardwareInfo_CPU_Speed = strconv.FormatFloat(cs[0].Mhz, 'f', 2, 64) + " Mhz"
	Inf.Server_HardwareInfo_CPU_CacheSize = strconv.FormatInt(int64(cs[0].CacheSize), 10)
	Inf.Server_HardwareInfo_CPU_Micronode = cs[0].Microcode
	Inf.Server_HardwareInfo_CPU_Model = cs[0].Model
	Inf.Server_HardwareInfo_CPU_PhysID = cs[0].PhysicalID
}

// Init Host information
func (Inf *Server) HOST() {
	hs, x := host.Info()
	Sav(x)
	Inf.Server_HardwareInfo_OSP_Hostname = hs.Hostname
	Inf.Server_HardwareInfo_OSP_Uptime = strconv.FormatUint(hs.Uptime, 10)
	Inf.Server_HardwareInfo_OSP_HOSTID = hs.HostID
	Inf.Server_HardwareInfo_OSP_ProcRunning = strconv.FormatUint(hs.Procs, 10)
	Inf.Server_HardwareInfo_OSP_HOST_KERNEL_ARCHITECTURE = hs.KernelArch
	Inf.Server_HardwareInfo_OSP_HOST_KERNEL_VERSION = hs.KernelVersion
	Inf.Server_HardwareInfo_OSP_HOST_PLATFORM_VERSION = hs.PlatformVersion
	Inf.Server_HardwareInfo_OSP_HOST_PLATFORM_FAMILY = hs.PlatformFamily
	Inf.Server_HardwareInfo_OSP_HOSTPLAT = hs.Platform
	Inf.Server_HardwareInfo_OSP_HOSTOS = hs.OS
}

// Init Memory information
func (Inf *Server) MEM() {
	mem, x := mem.VirtualMemory()
	Sav(x)
	Inf.Server_HardwareInfo_MEM_Free = strconv.FormatUint(mem.Free, 10) + " Bytes "
	Inf.Server_HardwareInfo_MEM_Total = strconv.FormatUint(mem.Total, 10) + " Bytes "
	Inf.Server_HardwareInfo_MEM_Used = strconv.FormatFloat(mem.UsedPercent, 'f', 2, 64) + "%"
}
