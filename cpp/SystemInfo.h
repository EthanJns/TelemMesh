#ifndef SYSTEMINFO_H
#define SYSTEMINFO_H

#include <string>

enum class OperatingSystem {
    Windows,
    MacOS,
    Linux,
    Unknown
};

class SystemInfo {
public:
    static OperatingSystem detectOS();
    static std::string getOSName();
    static bool fileExists(const std::string& path);
    static int getCoreCount();
};

#endif