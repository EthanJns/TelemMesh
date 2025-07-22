use std::fs;
use std::path::Path;
use sysinfo::System;

#[derive(Debug, Clone, PartialEq)]
pub enum OperatingSystem {
	Windows,
	MacOS,
	Linux,
	Unkown,
}

pub struct SystemInfo;

impl SystemInfo {
	pub fn detect_os() -> OperatingSystem {
		if cfg!(target_os = "windows") {
			return OperatingSystem::Windows;
		}

		if cfg!(target_os = "macos") {
			return OperatingSystem::MacOS;
		}

		if cfg!(target_os = "linux") {
			return OperatingSystem::Linux;
		}

		return OperatingSystem::Unkown;
	}

	pub fn get_os_name() -> String {
		match Self::detect_os() {
			OperatingSystem::Windows => "Windows".to_string(),
			OperatingSystem::MacOS => "macOS".to_string(),
			OperatingSystem::Linux => "Linux".to_string(),
			OperatingSystem::Unkown => "Unkown".to_string(),
		}
	}

	pub fn file_exists(path: &str) -> bool {
		Path::new(path).exists()
	}

	pub fn get_core_count() -> i32 {
		std::thread::available_parallelism()
            .map(|n| n.get() as i32)
            .unwrap_or(1)
	}
}