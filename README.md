# eBPF BTF (BPF Type Format) Demo Repository

This repository provides an example for building CO-RE (Compile-Once-Run-Everywhere) eBPF programs.

![image-tmp (13)](https://github.com/user-attachments/assets/37404be6-f504-4421-90d8-c2bec9fc114d)

In other words - if you're seeing this error:
```
Loading eBPF objects: ...: ...: apply CO-RE relocations: load kernel spec: no BTF found for kernel version x.xx.x-xxx-generic: not supported
```
you came to the right place.

It includes `Makefile` scripts to:
- generate skeleton code for eBPF kernel program
- generate minimal BTF files based on the skeleton code
- builds the final program binary that can run also on kernel versions that aren't compiled with BTF support

## Requirements

- [bpftool](https://github.com/libbpf/bpftool)
- Access to the [btfhub-archive repository](https://github.com/aquasecurity/btfhub-archive), which should be cloned locally.
- [eBPF Dependencies](https://ebpf-go.dev/guides/getting-started/#ebpf-c-program)
- Go installed on your system

## Features

- **Minimal BTF Generation:** Automated creation of minimal BTF files tailored for specific architectures.
- **Architecture Support:** Embeds BTF information for different kernel versions across both `x86` and `arm64` architectures and Operating System like `fedora`, `ubuntu`, `centos` etc..
- **Example eBPF Program:** Includes a sample eBPF program (example.c) demonstrating BTF integration.

## Repository Structure

- `Makefile`: Main build script handling generation and compilation.
- `Makefile.btfgen`: Handles the generation of minimal BTF files.
- `example.c`: Sample eBPF program demonstrating BTF usage.
- `main.go`: eBPF user space program in Go for interfacing with the kernel space eBPF program.
- `btfgen.go`, `btfgen_amd64.go`, `btfgen_arm64.go`: Go files for embedding BTF specs for different architectures into the output binary.
- `vmlinux.h`: Header file required for BTF and eBPF program compilation.
- `vmlinux_missing.h` (optional): Since `vmlinux.h` contains only kernel BTF types you can also define the potential macros you might need here.

## How to use

First clone [btfhub-archive repository](https://github.com/aquasecurity/btfhub-archive) - includes BTF files for existing published kernels that don't support embedded BTF.

Then set path to it using:
```
export BTFHUB_ARCHIVE=~/btfhub-archive
```

You are now ready to generate minimal BTF files and build the output binary using:
```
make
```

## License

Some parts of the code have been inspired by the concepts utilized in the [Inspektor Gadget](https://github.com/inspektor-gadget/inspektor-gadget) which is licensed under the LGPL-2.1 OR BSD-2-Clause license.
