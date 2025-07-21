#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>
#include "SystemInfo.h"

// OS Detection
#ifdef _WIN32
    #define OS_WINDOWS
    #include <windows.h>
    #include <pdh.h>
    #include <pdhmsg.h>
#elif __APPLE__
    #define OS_MACOS
    #include <sys/types.h>
    #include <sys/sysctl.h>
    #include <mach/mach.h>
#elif __linux__
    #define OS_LINUX
    #include <sched.h>
    #include <unistd.h>
#else
    #define OS_UNKNOWN
#endif


OperatingSystem SystemInfo::detectOS() {
	#ifdef OS_WINDOWS
				return OperatingSystem::Windows;
			#elif OS_MACOS
				return OperatingSystem::MacOS;
			#elif OS_LINUX
				return OperatingSystem::Linux;
			#else
				return OperatingSystem::Unknown;
			#endif
}

std::string SystemInfo::getOSName() {
	switch(detectOS()) {
				case OperatingSystem::Windows: return "Windows";
				case OperatingSystem::MacOS: return "macOS";
				case OperatingSystem::Linux: return "Linux";
				case OperatingSystem::Unknown: return "Unknown";
			}

			return "Unknown";
}

bool SystemInfo::fileExists(const std::string& path) {
	std::ifstream file(path);
	return file.good();
}

int SystemInfo::getCoreCount() {
	#ifdef OS_WINDOWS
		SYSTEM_INFO sysinfo;
		GetSystemInfo(&sysinfo);
		return sysinfo.dwNumberOfProcessors;
	#elif OS_MACOS
		int count;
		size_t size = sizeof(count);
		if (sysctlbyname("hw.ncpu", &count, &size, NULL, 0) == 0) {
			return count;
		}
		return -1;
	#elif OS_LINUX
		return sysconf(_SC_NPROCESSORS_ONLN);
	#else
		return -1;
	#endif
}
