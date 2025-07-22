mod system_info;

use system_info::SystemInfo;

fn main() {
    
	let os_name = SystemInfo::get_os_name();
	println!("Operating System: {}", os_name);

	let core_count = SystemInfo::get_core_count();
	println!("Cpu Core Count: {}", core_count);

}
