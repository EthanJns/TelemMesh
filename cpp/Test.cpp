#include<iostream>
#include "SystemInfo.h"
using namespace std;

class SystemInfo;
int main() {
	
	cout<<SystemInfo::getOSName()<<endl;
	cout<<SystemInfo::getCoreCount()<<endl;

	return 1;
}