// SPDX-FileCopyrightText: Alvar Penning
//
// SPDX-License-Identifier: BSD-3-Clause

// DO NOT EDIT DIRECTLY!

// Package syscallset is an easy to use Go library allowing any Go program to
// restrict its own capabilities on Linux through seccomp-bpf and predefined
// syscall sets from systemd.
//
// To self-limit a Go application, use this package's LimitTo or LimitAndLog
// functions.
//
// The following syscall groups from systemd are available:
//
//	- aio: Asynchronous IO
//	- basic-io: Basic IO
//	- chown: Change ownership of files and directories
//	- clock: Change the system time
//	- cpu-emulation: System calls for CPU emulation functionality
//	- debug: Debugging, performance monitoring and tracing functionality
//	- default: System calls that are always permitted
//	- file-system: File system operations
//	- io-event: Event loop system calls
//	- ipc: SysV IPC, POSIX Message Queues or other IPC
//	- keyring: Kernel keyring access
//	- known: All known syscalls declared in the kernel
//	- memlock: Memory locking control
//	- module: Loading and unloading of kernel modules
//	- mount: Mounting and unmounting of file systems
//	- network-io: Network or Unix socket IO, should not be needed if not network facing
//	- obsolete: Unusual, obsolete or unimplemented system calls
//	- pkey: System calls used for memory protection keys
//	- privileged: All system calls which need super-user capabilities
//	- process: Process control, execution, namespacing operations
//	- raw-io: Raw I/O port access
//	- reboot: Reboot and reboot preparation/kexec
//	- resources: Alter resource settings
//	- sandbox: Sandbox functionality
//	- setuid: Operations for changing user/group credentials
//	- signal: Process signal handling
//	- swap: Enable/disable swap devices
//	- sync: Synchronize files and memory to storage
//	- system-service: General system service operations
//	- timer: Schedule operations by time
//
package syscallset

