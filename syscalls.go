// SPDX-FileCopyrightText: systemd contributors
// SPDX-FileCopyrightText: Alvar Penning
//
// SPDX-License-Identifier: LGPL-2.1-or-later
// SPDX-License-Identifier: BSD-3-Clause

// DO NOT EDIT DIRECTLY!

// This file was automatically generated by generator/gen.go based on the output
// of `systemd-analyze syscall-filter`. For more information, please refer to:
// - https://www.freedesktop.org/software/systemd/man/systemd.exec.html
// - https://www.freedesktop.org/software/systemd/man/systemd-analyze.html
// - https://github.com/systemd/systemd/blob/main/src/shared/seccomp-util.c
// - https://github.com/systemd/systemd/blob/main/src/analyze/analyze-syscall-filter.c

// Generated based on systemd 254 (254.3-1-arch)

package syscallset

var syscallSets = map[string][]string{
	"aio": []string{
		"io_cancel",
		"io_destroy",
		"io_getevents",
		"io_pgetevents",
		"io_pgetevents_time64",
		"io_setup",
		"io_submit",
		"io_uring_enter",
		"io_uring_register",
		"io_uring_setup",
	},
	"basic-io": []string{
		"_llseek",
		"close",
		"close_range",
		"dup",
		"dup2",
		"dup3",
		"lseek",
		"pread64",
		"preadv",
		"preadv2",
		"pwrite64",
		"pwritev",
		"pwritev2",
		"read",
		"readv",
		"write",
		"writev",
	},
	"chown": []string{
		"chown",
		"chown32",
		"fchown",
		"fchown32",
		"fchownat",
		"lchown",
		"lchown32",
	},
	"clock": []string{
		"adjtimex",
		"clock_adjtime",
		"clock_adjtime64",
		"clock_settime",
		"clock_settime64",
		"settimeofday",
	},
	"cpu-emulation": []string{
		"modify_ldt",
		"subpage_prot",
		"switch_endian",
		"vm86",
		"vm86old",
	},
	"debug": []string{
		"lookup_dcookie",
		"perf_event_open",
		"pidfd_getfd",
		"ptrace",
		"rtas",
		"s390_runtime_instr",
		"sys_debug_setcontext",
	},
	"default": []string{
		"arch_prctl",
		"brk",
		"cacheflush",
		"clock_getres",
		"clock_getres_time64",
		"clock_gettime",
		"clock_gettime64",
		"clock_nanosleep",
		"clock_nanosleep_time64",
		"exit",
		"exit_group",
		"futex",
		"futex_time64",
		"futex_waitv",
		"get_robust_list",
		"get_thread_area",
		"getegid",
		"getegid32",
		"geteuid",
		"geteuid32",
		"getgid",
		"getgid32",
		"getgroups",
		"getgroups32",
		"getpgid",
		"getpgrp",
		"getpid",
		"getppid",
		"getrandom",
		"getresgid",
		"getresgid32",
		"getresuid",
		"getresuid32",
		"getrlimit",
		"getsid",
		"gettid",
		"gettimeofday",
		"getuid",
		"getuid32",
		"membarrier",
		"mmap",
		"mmap2",
		"mprotect",
		"munmap",
		"nanosleep",
		"pause",
		"prlimit64",
		"restart_syscall",
		"riscv_flush_icache",
		"riscv_hwprobe",
		"rseq",
		"rt_sigreturn",
		"sched_getaffinity",
		"sched_yield",
		"set_robust_list",
		"set_thread_area",
		"set_tid_address",
		"set_tls",
		"sigreturn",
		"time",
		"ugetrlimit",
	},
	"file-system": []string{
		"access",
		"chdir",
		"chmod",
		"close",
		"creat",
		"faccessat",
		"faccessat2",
		"fallocate",
		"fchdir",
		"fchmod",
		"fchmodat",
		"fcntl",
		"fcntl64",
		"fgetxattr",
		"flistxattr",
		"fremovexattr",
		"fsetxattr",
		"fstat",
		"fstat64",
		"fstatat64",
		"fstatfs",
		"fstatfs64",
		"ftruncate",
		"ftruncate64",
		"futimesat",
		"getcwd",
		"getdents",
		"getdents64",
		"getxattr",
		"inotify_add_watch",
		"inotify_init",
		"inotify_init1",
		"inotify_rm_watch",
		"lgetxattr",
		"link",
		"linkat",
		"listxattr",
		"llistxattr",
		"lremovexattr",
		"lsetxattr",
		"lstat",
		"lstat64",
		"mkdir",
		"mkdirat",
		"mknod",
		"mknodat",
		"newfstatat",
		"oldfstat",
		"oldlstat",
		"oldstat",
		"open",
		"openat",
		"openat2",
		"readlink",
		"readlinkat",
		"removexattr",
		"rename",
		"renameat",
		"renameat2",
		"rmdir",
		"setxattr",
		"stat",
		"stat64",
		"statfs",
		"statfs64",
		"statx",
		"symlink",
		"symlinkat",
		"truncate",
		"truncate64",
		"unlink",
		"unlinkat",
		"utime",
		"utimensat",
		"utimensat_time64",
		"utimes",
	},
	"io-event": []string{
		"_newselect",
		"epoll_create",
		"epoll_create1",
		"epoll_ctl",
		"epoll_ctl_old",
		"epoll_pwait",
		"epoll_pwait2",
		"epoll_wait",
		"epoll_wait_old",
		"eventfd",
		"eventfd2",
		"poll",
		"ppoll",
		"ppoll_time64",
		"pselect6",
		"pselect6_time64",
		"select",
	},
	"ipc": []string{
		"ipc",
		"memfd_create",
		"mq_getsetattr",
		"mq_notify",
		"mq_open",
		"mq_timedreceive",
		"mq_timedreceive_time64",
		"mq_timedsend",
		"mq_timedsend_time64",
		"mq_unlink",
		"msgctl",
		"msgget",
		"msgrcv",
		"msgsnd",
		"pipe",
		"pipe2",
		"process_madvise",
		"process_vm_readv",
		"process_vm_writev",
		"semctl",
		"semget",
		"semop",
		"semtimedop",
		"semtimedop_time64",
		"shmat",
		"shmctl",
		"shmdt",
		"shmget",
	},
	"keyring": []string{
		"add_key",
		"keyctl",
		"request_key",
	},
	"known": []string{
		"_llseek",
		"_newselect",
		"_sysctl",
		"accept",
		"accept4",
		"access",
		"acct",
		"add_key",
		"adjtimex",
		"afs_syscall",
		"alarm",
		"arc_gettls",
		"arc_settls",
		"arc_usr_cmpxchg",
		"arch_prctl",
		"atomic_barrier",
		"atomic_cmpxchg_32",
		"bdflush",
		"bind",
		"bpf",
		"break",
		"brk",
		"cachectl",
		"cacheflush",
		"cachestat",
		"capget",
		"capset",
		"chdir",
		"chmod",
		"chown",
		"chown32",
		"chroot",
		"clock_adjtime",
		"clock_adjtime64",
		"clock_getres",
		"clock_getres_time64",
		"clock_gettime",
		"clock_gettime64",
		"clock_nanosleep",
		"clock_nanosleep_time64",
		"clock_settime",
		"clock_settime64",
		"clone",
		"clone2",
		"clone3",
		"close",
		"close_range",
		"connect",
		"copy_file_range",
		"creat",
		"create_module",
		"delete_module",
		"dipc",
		"dup",
		"dup2",
		"dup3",
		"epoll_create",
		"epoll_create1",
		"epoll_ctl",
		"epoll_ctl_old",
		"epoll_pwait",
		"epoll_pwait2",
		"epoll_wait",
		"epoll_wait_old",
		"eventfd",
		"eventfd2",
		"exec_with_loader",
		"execv",
		"execve",
		"execveat",
		"exit",
		"exit_group",
		"faccessat",
		"faccessat2",
		"fadvise64",
		"fadvise64_64",
		"fallocate",
		"fanotify_init",
		"fanotify_mark",
		"fchdir",
		"fchmod",
		"fchmodat",
		"fchown",
		"fchown32",
		"fchownat",
		"fcntl",
		"fcntl64",
		"fdatasync",
		"fgetxattr",
		"finit_module",
		"flistxattr",
		"flock",
		"fork",
		"fp_udfiex_crtl",
		"fremovexattr",
		"fsconfig",
		"fsetxattr",
		"fsmount",
		"fsopen",
		"fspick",
		"fstat",
		"fstat64",
		"fstatat64",
		"fstatfs",
		"fstatfs64",
		"fsync",
		"ftime",
		"ftruncate",
		"ftruncate64",
		"futex",
		"futex_time64",
		"futex_waitv",
		"futimesat",
		"get_kernel_syms",
		"get_mempolicy",
		"get_robust_list",
		"get_thread_area",
		"getcpu",
		"getcwd",
		"getdents",
		"getdents64",
		"getdomainname",
		"getdtablesize",
		"getegid",
		"getegid32",
		"geteuid",
		"geteuid32",
		"getgid",
		"getgid32",
		"getgroups",
		"getgroups32",
		"gethostname",
		"getitimer",
		"getpagesize",
		"getpeername",
		"getpgid",
		"getpgrp",
		"getpid",
		"getpmsg",
		"getppid",
		"getpriority",
		"getrandom",
		"getresgid",
		"getresgid32",
		"getresuid",
		"getresuid32",
		"getrlimit",
		"getrusage",
		"getsid",
		"getsockname",
		"getsockopt",
		"gettid",
		"gettimeofday",
		"getuid",
		"getuid32",
		"getunwind",
		"getxattr",
		"getxgid",
		"getxpid",
		"getxuid",
		"gtty",
		"idle",
		"init_module",
		"inotify_add_watch",
		"inotify_init",
		"inotify_init1",
		"inotify_rm_watch",
		"io_cancel",
		"io_destroy",
		"io_getevents",
		"io_pgetevents",
		"io_pgetevents_time64",
		"io_setup",
		"io_submit",
		"io_uring_enter",
		"io_uring_register",
		"io_uring_setup",
		"ioctl",
		"ioperm",
		"iopl",
		"ioprio_get",
		"ioprio_set",
		"ipc",
		"kcmp",
		"kern_features",
		"kexec_file_load",
		"kexec_load",
		"keyctl",
		"kill",
		"landlock_add_rule",
		"landlock_create_ruleset",
		"landlock_restrict_self",
		"lchown",
		"lchown32",
		"lgetxattr",
		"link",
		"linkat",
		"listen",
		"listxattr",
		"llistxattr",
		"lock",
		"lookup_dcookie",
		"lremovexattr",
		"lseek",
		"lsetxattr",
		"lstat",
		"lstat64",
		"madvise",
		"mbind",
		"membarrier",
		"memfd_create",
		"memfd_secret",
		"memory_ordering",
		"migrate_pages",
		"mincore",
		"mkdir",
		"mkdirat",
		"mknod",
		"mknodat",
		"mlock",
		"mlock2",
		"mlockall",
		"mmap",
		"mmap2",
		"modify_ldt",
		"mount",
		"mount_setattr",
		"move_mount",
		"move_pages",
		"mprotect",
		"mpx",
		"mq_getsetattr",
		"mq_notify",
		"mq_open",
		"mq_timedreceive",
		"mq_timedreceive_time64",
		"mq_timedsend",
		"mq_timedsend_time64",
		"mq_unlink",
		"mremap",
		"msgctl",
		"msgget",
		"msgrcv",
		"msgsnd",
		"msync",
		"multiplexer",
		"munlock",
		"munlockall",
		"munmap",
		"name_to_handle_at",
		"nanosleep",
		"newfstatat",
		"nfsservctl",
		"nice",
		"old_adjtimex",
		"old_getpagesize",
		"oldfstat",
		"oldlstat",
		"oldolduname",
		"oldstat",
		"oldumount",
		"olduname",
		"open",
		"open_by_handle_at",
		"open_tree",
		"openat",
		"openat2",
		"or1k_atomic",
		"osf_adjtime",
		"osf_afs_syscall",
		"osf_alt_plock",
		"osf_alt_setsid",
		"osf_alt_sigpending",
		"osf_asynch_daemon",
		"osf_audcntl",
		"osf_audgen",
		"osf_chflags",
		"osf_execve",
		"osf_exportfs",
		"osf_fchflags",
		"osf_fdatasync",
		"osf_fpathconf",
		"osf_fstat",
		"osf_fstatfs",
		"osf_fstatfs64",
		"osf_fuser",
		"osf_getaddressconf",
		"osf_getdirentries",
		"osf_getdomainname",
		"osf_getfh",
		"osf_getfsstat",
		"osf_gethostid",
		"osf_getitimer",
		"osf_getlogin",
		"osf_getmnt",
		"osf_getrusage",
		"osf_getsysinfo",
		"osf_gettimeofday",
		"osf_kloadcall",
		"osf_kmodcall",
		"osf_lstat",
		"osf_memcntl",
		"osf_mincore",
		"osf_mount",
		"osf_mremap",
		"osf_msfs_syscall",
		"osf_msleep",
		"osf_mvalid",
		"osf_mwakeup",
		"osf_naccept",
		"osf_nfssvc",
		"osf_ngetpeername",
		"osf_ngetsockname",
		"osf_nrecvfrom",
		"osf_nrecvmsg",
		"osf_nsendmsg",
		"osf_ntp_adjtime",
		"osf_ntp_gettime",
		"osf_old_creat",
		"osf_old_fstat",
		"osf_old_getpgrp",
		"osf_old_killpg",
		"osf_old_lstat",
		"osf_old_open",
		"osf_old_sigaction",
		"osf_old_sigblock",
		"osf_old_sigreturn",
		"osf_old_sigsetmask",
		"osf_old_sigvec",
		"osf_old_stat",
		"osf_old_vadvise",
		"osf_old_vtrace",
		"osf_old_wait",
		"osf_oldquota",
		"osf_pathconf",
		"osf_pid_block",
		"osf_pid_unblock",
		"osf_plock",
		"osf_priocntlset",
		"osf_profil",
		"osf_proplist_syscall",
		"osf_reboot",
		"osf_revoke",
		"osf_sbrk",
		"osf_security",
		"osf_select",
		"osf_set_program_attributes",
		"osf_set_speculative",
		"osf_sethostid",
		"osf_setitimer",
		"osf_setlogin",
		"osf_setsysinfo",
		"osf_settimeofday",
		"osf_shmat",
		"osf_signal",
		"osf_sigprocmask",
		"osf_sigsendset",
		"osf_sigstack",
		"osf_sigwaitprim",
		"osf_sstk",
		"osf_stat",
		"osf_statfs",
		"osf_statfs64",
		"osf_subsys_info",
		"osf_swapctl",
		"osf_swapon",
		"osf_syscall",
		"osf_sysinfo",
		"osf_table",
		"osf_uadmin",
		"osf_usleep_thread",
		"osf_uswitch",
		"osf_utc_adjtime",
		"osf_utc_gettime",
		"osf_utimes",
		"osf_utsname",
		"osf_wait4",
		"osf_waitid",
		"pause",
		"pciconfig_iobase",
		"pciconfig_read",
		"pciconfig_write",
		"perf_event_open",
		"perfctr",
		"personality",
		"pidfd_getfd",
		"pidfd_open",
		"pidfd_send_signal",
		"pipe",
		"pipe2",
		"pivot_root",
		"pkey_alloc",
		"pkey_free",
		"pkey_mprotect",
		"poll",
		"ppoll",
		"ppoll_time64",
		"prctl",
		"pread64",
		"preadv",
		"preadv2",
		"prlimit64",
		"process_madvise",
		"process_mrelease",
		"process_vm_readv",
		"process_vm_writev",
		"prof",
		"profil",
		"pselect6",
		"pselect6_time64",
		"ptrace",
		"putpmsg",
		"pwrite64",
		"pwritev",
		"pwritev2",
		"query_module",
		"quotactl",
		"quotactl_fd",
		"read",
		"readahead",
		"readdir",
		"readlink",
		"readlinkat",
		"readv",
		"reboot",
		"recv",
		"recvfrom",
		"recvmmsg",
		"recvmmsg_time64",
		"recvmsg",
		"remap_file_pages",
		"removexattr",
		"rename",
		"renameat",
		"renameat2",
		"request_key",
		"restart_syscall",
		"riscv_flush_icache",
		"riscv_hwprobe",
		"rmdir",
		"rseq",
		"rt_sigaction",
		"rt_sigpending",
		"rt_sigprocmask",
		"rt_sigqueueinfo",
		"rt_sigreturn",
		"rt_sigsuspend",
		"rt_sigtimedwait",
		"rt_sigtimedwait_time64",
		"rt_tgsigqueueinfo",
		"rtas",
		"sched_get_affinity",
		"sched_get_priority_max",
		"sched_get_priority_min",
		"sched_getaffinity",
		"sched_getattr",
		"sched_getparam",
		"sched_getscheduler",
		"sched_rr_get_interval",
		"sched_rr_get_interval_time64",
		"sched_set_affinity",
		"sched_setaffinity",
		"sched_setattr",
		"sched_setparam",
		"sched_setscheduler",
		"sched_yield",
		"seccomp",
		"security",
		"select",
		"semctl",
		"semget",
		"semop",
		"semtimedop",
		"semtimedop_time64",
		"send",
		"sendfile",
		"sendfile64",
		"sendmmsg",
		"sendmsg",
		"sendto",
		"set_mempolicy",
		"set_mempolicy_home_node",
		"set_robust_list",
		"set_thread_area",
		"set_tid_address",
		"setdomainname",
		"setfsgid",
		"setfsgid32",
		"setfsuid",
		"setfsuid32",
		"setgid",
		"setgid32",
		"setgroups",
		"setgroups32",
		"sethae",
		"sethostname",
		"setitimer",
		"setns",
		"setpgid",
		"setpgrp",
		"setpriority",
		"setregid",
		"setregid32",
		"setresgid",
		"setresgid32",
		"setresuid",
		"setresuid32",
		"setreuid",
		"setreuid32",
		"setrlimit",
		"setsid",
		"setsockopt",
		"settimeofday",
		"setuid",
		"setuid32",
		"setxattr",
		"sgetmask",
		"shmat",
		"shmctl",
		"shmdt",
		"shmget",
		"shutdown",
		"sigaction",
		"sigaltstack",
		"signal",
		"signalfd",
		"signalfd4",
		"sigpending",
		"sigprocmask",
		"sigreturn",
		"sigsuspend",
		"socket",
		"socketcall",
		"socketpair",
		"splice",
		"spu_create",
		"spu_run",
		"ssetmask",
		"stat",
		"stat64",
		"statfs",
		"statfs64",
		"statx",
		"stime",
		"stty",
		"subpage_prot",
		"swapcontext",
		"swapoff",
		"swapon",
		"switch_endian",
		"symlink",
		"symlinkat",
		"sync",
		"sync_file_range",
		"sync_file_range2",
		"syncfs",
		"sys_debug_setcontext",
		"syscall",
		"sysfs",
		"sysinfo",
		"syslog",
		"sysmips",
		"tee",
		"tgkill",
		"time",
		"timer_create",
		"timer_delete",
		"timer_getoverrun",
		"timer_gettime",
		"timer_gettime64",
		"timer_settime",
		"timer_settime64",
		"timerfd",
		"timerfd_create",
		"timerfd_gettime",
		"timerfd_gettime64",
		"timerfd_settime",
		"timerfd_settime64",
		"times",
		"tkill",
		"truncate",
		"truncate64",
		"tuxcall",
		"ugetrlimit",
		"ulimit",
		"umask",
		"umount",
		"umount2",
		"uname",
		"unlink",
		"unlinkat",
		"unshare",
		"uselib",
		"userfaultfd",
		"ustat",
		"utime",
		"utimensat",
		"utimensat_time64",
		"utimes",
		"utrap_install",
		"vfork",
		"vhangup",
		"vm86",
		"vm86old",
		"vmsplice",
		"vserver",
		"wait4",
		"waitid",
		"waitpid",
		"write",
		"writev",
	},
	"memlock": []string{
		"mlock",
		"mlock2",
		"mlockall",
		"munlock",
		"munlockall",
	},
	"module": []string{
		"delete_module",
		"finit_module",
		"init_module",
	},
	"mount": []string{
		"chroot",
		"fsconfig",
		"fsmount",
		"fsopen",
		"fspick",
		"mount",
		"mount_setattr",
		"move_mount",
		"open_tree",
		"pivot_root",
		"umount",
		"umount2",
	},
	"network-io": []string{
		"accept",
		"accept4",
		"bind",
		"connect",
		"getpeername",
		"getsockname",
		"getsockopt",
		"listen",
		"recv",
		"recvfrom",
		"recvmmsg",
		"recvmmsg_time64",
		"recvmsg",
		"send",
		"sendmmsg",
		"sendmsg",
		"sendto",
		"setsockopt",
		"shutdown",
		"socket",
		"socketcall",
		"socketpair",
	},
	"obsolete": []string{
		"_sysctl",
		"afs_syscall",
		"bdflush",
		"break",
		"create_module",
		"ftime",
		"get_kernel_syms",
		"getpmsg",
		"gtty",
		"idle",
		"lock",
		"mpx",
		"prof",
		"profil",
		"putpmsg",
		"query_module",
		"security",
		"sgetmask",
		"ssetmask",
		"stime",
		"stty",
		"sysfs",
		"tuxcall",
		"ulimit",
		"uselib",
		"ustat",
		"vserver",
	},
	"pkey": []string{
		"pkey_alloc",
		"pkey_free",
		"pkey_mprotect",
	},
	"privileged": []string{
		"_sysctl",
		"acct",
		"adjtimex",
		"bpf",
		"capset",
		"chown",
		"chown32",
		"chroot",
		"clock_adjtime",
		"clock_adjtime64",
		"clock_settime",
		"clock_settime64",
		"delete_module",
		"fanotify_init",
		"fanotify_mark",
		"fchown",
		"fchown32",
		"fchownat",
		"finit_module",
		"init_module",
		"ioperm",
		"iopl",
		"kexec_file_load",
		"kexec_load",
		"lchown",
		"lchown32",
		"nfsservctl",
		"open_by_handle_at",
		"pciconfig_iobase",
		"pciconfig_read",
		"pciconfig_write",
		"pivot_root",
		"quotactl",
		"quotactl_fd",
		"reboot",
		"s390_pci_mmio_read",
		"s390_pci_mmio_write",
		"setdomainname",
		"setfsuid",
		"setfsuid32",
		"setgroups",
		"setgroups32",
		"sethostname",
		"setresuid",
		"setresuid32",
		"setreuid",
		"setreuid32",
		"settimeofday",
		"setuid",
		"setuid32",
		"swapoff",
		"swapon",
		"vhangup",
	},
	"process": []string{
		"capget",
		"clone",
		"clone3",
		"execve",
		"execveat",
		"fork",
		"getrusage",
		"kill",
		"pidfd_open",
		"pidfd_send_signal",
		"prctl",
		"rt_sigqueueinfo",
		"rt_tgsigqueueinfo",
		"setns",
		"swapcontext",
		"tgkill",
		"times",
		"tkill",
		"unshare",
		"vfork",
		"wait4",
		"waitid",
		"waitpid",
	},
	"raw-io": []string{
		"ioperm",
		"iopl",
		"pciconfig_iobase",
		"pciconfig_read",
		"pciconfig_write",
		"s390_pci_mmio_read",
		"s390_pci_mmio_write",
	},
	"reboot": []string{
		"kexec_file_load",
		"kexec_load",
		"reboot",
	},
	"resources": []string{
		"ioprio_set",
		"mbind",
		"migrate_pages",
		"move_pages",
		"nice",
		"sched_setaffinity",
		"sched_setattr",
		"sched_setparam",
		"sched_setscheduler",
		"set_mempolicy",
		"set_mempolicy_home_node",
		"setpriority",
		"setrlimit",
	},
	"sandbox": []string{
		"landlock_add_rule",
		"landlock_create_ruleset",
		"landlock_restrict_self",
		"seccomp",
	},
	"setuid": []string{
		"setgid",
		"setgid32",
		"setgroups",
		"setgroups32",
		"setregid",
		"setregid32",
		"setresgid",
		"setresgid32",
		"setresuid",
		"setresuid32",
		"setreuid",
		"setreuid32",
		"setuid",
		"setuid32",
	},
	"signal": []string{
		"rt_sigaction",
		"rt_sigpending",
		"rt_sigprocmask",
		"rt_sigsuspend",
		"rt_sigtimedwait",
		"rt_sigtimedwait_time64",
		"sigaction",
		"sigaltstack",
		"signal",
		"signalfd",
		"signalfd4",
		"sigpending",
		"sigprocmask",
		"sigsuspend",
	},
	"swap": []string{
		"swapoff",
		"swapon",
	},
	"sync": []string{
		"fdatasync",
		"fsync",
		"msync",
		"sync",
		"sync_file_range",
		"sync_file_range2",
		"syncfs",
	},
	"system-service": []string{
		"_llseek",
		"_newselect",
		"accept",
		"accept4",
		"access",
		"add_key",
		"alarm",
		"arch_prctl",
		"arm_fadvise64_64",
		"bind",
		"brk",
		"cacheflush",
		"capget",
		"capset",
		"chdir",
		"chmod",
		"chown",
		"chown32",
		"clock_getres",
		"clock_getres_time64",
		"clock_gettime",
		"clock_gettime64",
		"clock_nanosleep",
		"clock_nanosleep_time64",
		"clone",
		"clone3",
		"close",
		"close_range",
		"connect",
		"copy_file_range",
		"creat",
		"dup",
		"dup2",
		"dup3",
		"epoll_create",
		"epoll_create1",
		"epoll_ctl",
		"epoll_ctl_old",
		"epoll_pwait",
		"epoll_pwait2",
		"epoll_wait",
		"epoll_wait_old",
		"eventfd",
		"eventfd2",
		"execve",
		"execveat",
		"exit",
		"exit_group",
		"faccessat",
		"faccessat2",
		"fadvise64",
		"fadvise64_64",
		"fallocate",
		"fchdir",
		"fchmod",
		"fchmodat",
		"fchown",
		"fchown32",
		"fchownat",
		"fcntl",
		"fcntl64",
		"fdatasync",
		"fgetxattr",
		"flistxattr",
		"flock",
		"fork",
		"fremovexattr",
		"fsetxattr",
		"fstat",
		"fstat64",
		"fstatat64",
		"fstatfs",
		"fstatfs64",
		"fsync",
		"ftruncate",
		"ftruncate64",
		"futex",
		"futex_time64",
		"futex_waitv",
		"futimesat",
		"get_mempolicy",
		"get_robust_list",
		"get_thread_area",
		"getcpu",
		"getcwd",
		"getdents",
		"getdents64",
		"getegid",
		"getegid32",
		"geteuid",
		"geteuid32",
		"getgid",
		"getgid32",
		"getgroups",
		"getgroups32",
		"getitimer",
		"getpeername",
		"getpgid",
		"getpgrp",
		"getpid",
		"getppid",
		"getpriority",
		"getrandom",
		"getresgid",
		"getresgid32",
		"getresuid",
		"getresuid32",
		"getrlimit",
		"getrusage",
		"getsid",
		"getsockname",
		"getsockopt",
		"gettid",
		"gettimeofday",
		"getuid",
		"getuid32",
		"getxattr",
		"inotify_add_watch",
		"inotify_init",
		"inotify_init1",
		"inotify_rm_watch",
		"io_cancel",
		"io_destroy",
		"io_getevents",
		"io_pgetevents",
		"io_pgetevents_time64",
		"io_setup",
		"io_submit",
		"io_uring_enter",
		"io_uring_register",
		"io_uring_setup",
		"ioctl",
		"ioprio_get",
		"ioprio_set",
		"ipc",
		"kcmp",
		"keyctl",
		"kill",
		"lchown",
		"lchown32",
		"lgetxattr",
		"link",
		"linkat",
		"listen",
		"listxattr",
		"llistxattr",
		"lremovexattr",
		"lseek",
		"lsetxattr",
		"lstat",
		"lstat64",
		"madvise",
		"mbind",
		"membarrier",
		"memfd_create",
		"migrate_pages",
		"mkdir",
		"mkdirat",
		"mknod",
		"mknodat",
		"mlock",
		"mlock2",
		"mlockall",
		"mmap",
		"mmap2",
		"move_pages",
		"mprotect",
		"mq_getsetattr",
		"mq_notify",
		"mq_open",
		"mq_timedreceive",
		"mq_timedreceive_time64",
		"mq_timedsend",
		"mq_timedsend_time64",
		"mq_unlink",
		"mremap",
		"msgctl",
		"msgget",
		"msgrcv",
		"msgsnd",
		"msync",
		"munlock",
		"munlockall",
		"munmap",
		"name_to_handle_at",
		"nanosleep",
		"newfstatat",
		"nice",
		"oldfstat",
		"oldlstat",
		"oldolduname",
		"oldstat",
		"olduname",
		"open",
		"openat",
		"openat2",
		"pause",
		"personality",
		"pidfd_open",
		"pidfd_send_signal",
		"pipe",
		"pipe2",
		"poll",
		"ppoll",
		"ppoll_time64",
		"prctl",
		"pread64",
		"preadv",
		"preadv2",
		"prlimit64",
		"process_madvise",
		"process_vm_readv",
		"process_vm_writev",
		"pselect6",
		"pselect6_time64",
		"pwrite64",
		"pwritev",
		"pwritev2",
		"read",
		"readahead",
		"readdir",
		"readlink",
		"readlinkat",
		"readv",
		"recv",
		"recvfrom",
		"recvmmsg",
		"recvmmsg_time64",
		"recvmsg",
		"remap_file_pages",
		"removexattr",
		"rename",
		"renameat",
		"renameat2",
		"request_key",
		"restart_syscall",
		"riscv_flush_icache",
		"riscv_hwprobe",
		"rmdir",
		"rseq",
		"rt_sigaction",
		"rt_sigpending",
		"rt_sigprocmask",
		"rt_sigqueueinfo",
		"rt_sigreturn",
		"rt_sigsuspend",
		"rt_sigtimedwait",
		"rt_sigtimedwait_time64",
		"rt_tgsigqueueinfo",
		"sched_get_priority_max",
		"sched_get_priority_min",
		"sched_getaffinity",
		"sched_getattr",
		"sched_getparam",
		"sched_getscheduler",
		"sched_rr_get_interval",
		"sched_rr_get_interval_time64",
		"sched_setaffinity",
		"sched_setattr",
		"sched_setparam",
		"sched_setscheduler",
		"sched_yield",
		"select",
		"semctl",
		"semget",
		"semop",
		"semtimedop",
		"semtimedop_time64",
		"send",
		"sendfile",
		"sendfile64",
		"sendmmsg",
		"sendmsg",
		"sendto",
		"set_mempolicy",
		"set_mempolicy_home_node",
		"set_robust_list",
		"set_thread_area",
		"set_tid_address",
		"set_tls",
		"setfsgid",
		"setfsgid32",
		"setfsuid",
		"setfsuid32",
		"setgid",
		"setgid32",
		"setgroups",
		"setgroups32",
		"setitimer",
		"setns",
		"setpgid",
		"setpriority",
		"setregid",
		"setregid32",
		"setresgid",
		"setresgid32",
		"setresuid",
		"setresuid32",
		"setreuid",
		"setreuid32",
		"setrlimit",
		"setsid",
		"setsockopt",
		"setuid",
		"setuid32",
		"setxattr",
		"shmat",
		"shmctl",
		"shmdt",
		"shmget",
		"shutdown",
		"sigaction",
		"sigaltstack",
		"signal",
		"signalfd",
		"signalfd4",
		"sigpending",
		"sigprocmask",
		"sigreturn",
		"sigsuspend",
		"socket",
		"socketcall",
		"socketpair",
		"splice",
		"stat",
		"stat64",
		"statfs",
		"statfs64",
		"statx",
		"swapcontext",
		"symlink",
		"symlinkat",
		"sync",
		"sync_file_range",
		"sync_file_range2",
		"syncfs",
		"sysinfo",
		"tee",
		"tgkill",
		"time",
		"timer_create",
		"timer_delete",
		"timer_getoverrun",
		"timer_gettime",
		"timer_gettime64",
		"timer_settime",
		"timer_settime64",
		"timerfd_create",
		"timerfd_gettime",
		"timerfd_gettime64",
		"timerfd_settime",
		"timerfd_settime64",
		"times",
		"tkill",
		"truncate",
		"truncate64",
		"ugetrlimit",
		"umask",
		"uname",
		"unlink",
		"unlinkat",
		"unshare",
		"userfaultfd",
		"utime",
		"utimensat",
		"utimensat_time64",
		"utimes",
		"vfork",
		"vmsplice",
		"wait4",
		"waitid",
		"waitpid",
		"write",
		"writev",
	},
	"timer": []string{
		"alarm",
		"getitimer",
		"setitimer",
		"timer_create",
		"timer_delete",
		"timer_getoverrun",
		"timer_gettime",
		"timer_gettime64",
		"timer_settime",
		"timer_settime64",
		"timerfd_create",
		"timerfd_gettime",
		"timerfd_gettime64",
		"timerfd_settime",
		"timerfd_settime64",
		"times",
	},
}
