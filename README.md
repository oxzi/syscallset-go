<!--
SPDX-FileCopyrightText: Alvar Penning

SPDX-License-Identifier: BSD-3-Clause
-->

# syscallset-go

[![Go Reference](https://pkg.go.dev/badge/github.com/oxzi/syscallset-go.svg)](https://pkg.go.dev/github.com/oxzi/syscallset-go)
[![CI](https://github.com/oxzi/syscallset-go/actions/workflows/ci.yml/badge.svg)](https://github.com/oxzi/syscallset-go/actions/workflows/ci.yml)
[![REUSE status](https://api.reuse.software/badge/github.com/oxzi/syscallset-go)](https://api.reuse.software/info/github.com/oxzi/syscallset-go)

`syscallset-go` is an easy to use Go library allowing any Go program to restrict its own capabilities on Linux through [`seccomp-bpf`](https://man7.org/linux/man-pages/man2/seccomp.2.html).
This makes it possible, for example, to prevent starting other malicious programs or establishing network connections, which drastically reduces the attack surface, also of third libraries.

Since `secomp-bpf` must be given an exact list of syscalls, which are even architecture-dependent, the creation and maintenance of an allow list is extremely tedious.
An easier way is shown by OpenBSD with its [`pledge`](https://man.openbsd.org/pledge.2) command, which expects groups instead of single syscalls.
Back in the Linux world, systemd allows grouped allowing or disallowing of syscalls via [`SystemCallFilter`](https://www.freedesktop.org/software/systemd/man/systemd.exec.html#System%20Call%20Filtering).

This library brings systemd's `SystemCallFilter`s to Go, allowing both systemd's platform-independent syscall sets as well as specific syscalls.
Internally, a given _allow list filter_ is converted to a `seccomp-bpf` profile, which is applied using the [go-seccomp-bpf](https://github.com/elastic/go-seccomp-bpf) library.
All code is Go, no cgo required.


## syscallset-go, the Go Library

The only relevant function is `LimitTo(syscallFilter string) error`, expecting a filter string which will be applied.
This filter will be always an allow list of syscall sets and/or single syscalls.
However, it is possible to restrict a previously used syscall set by subtracting syscalls.

The syntax follows systemd's `SystemCallFilter`, with some small differences.
To avoid duplicate documentation, please refer to the [library's documentation](https://pkg.go.dev/github.com/oxzi/syscallset-go).

The following example would be a simple Go program, restricted to only using `@basic-io` syscalls.
It cannot open network connections, files or even list directories.

```go
package main

import (
  "fmt"

  syscallset "github.com/oxzi/syscallset-go"
)

func main() {
  if err := syscallset.LimitTo("@basic-io"); err != nil {
    panic(err)
  }

  fmt.Println("hello restricted world")
}
```

For more information about the available syscall sets, please refer to
- [systemd's documentation](https://www.freedesktop.org/software/systemd/man/systemd.exec.html#System%20Call%20Filtering),
- [its source code (sorry)](https://github.com/systemd/systemd/blob/main/src/shared/seccomp-util.c), and
- the generated `syscalls.go` file within this repository.

To debug which syscalls are missing, exchange `LimitTo` with `LimitAndLog`.
This function will not terminate for non allowed syscalls, but log them to the audit log.
One can read this log with `auditd -f`.


## jail, Proof-of-Concept Jail

The simple `jail` program within `cmd/jail` allows starting another program with an applied syscallset filter.

```
$ go build ./cmd/jail
$ ./jail '@system-service ~@network-io ~@privileged' cowsay 'a somewhat more secure cow'
 ____________________________
< a somewhat more secure cow >
 ----------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

In this example, the `cowsay` program is allowed all syscalls within systemd's `@system-service` set except those within `@network-io` and `@privileged`.
Thus, cowsay will not be able to, e.g., upload your SSH private key or adjust your computer's clock.

Please only take this tool as an example and do not use it for something serious.
A better tool for exactly this job is already out there, it's called [Firejail](https://firejail.wordpress.com/).


## Updating syscalls.go

To update the `syscalls.go` file within this repository, one needs systemd running on the machine.
Then, just use the generator tool.

```
$ go run generator/gen.go generator/go.tmpl > syscalls.go
```

As systemd's userland applications are sufficient, one can use Docker to update the list to a recent version:

```
user@host $ docker pull archlinux:latest
user@host $ docker run -it --rm -v "$PWD":/app archlinux
root@container # pacman -Syu go
root@container # cd /app
root@container # go run generator/gen.go generator/go.tmpl > syscalls.go
root@container # go run generator/gen.go generator/go-apidoc.tmpl > syscallset.go
root@container # ^D
```


## Security Implications

The whole point of this library is to increase security by reducing privileges.
However, this code was not audited and might blow up in your face.
So please do with caution and feel free to report bugs or even structural errors.
