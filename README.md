# eBPF BTF Demo Repository

This repository provides an example for building CO-RE (Compile-Once-Run-Everywhere) eBPF programs.

It includes `Makefile` scripts to generate minimal BTF files that are embedded into the output binary, enabling the deployment of eBPF programs on systems without BTF.

## Features

- **Minimal BTF Generation:** Automated creation of minimal BTF files tailored for specific architectures using `make btfgen`.
- **Architecture Support:** Embeds BTF information, allowing the output binary to be run on different kernel versions across both `x86` and `arm64` architectures.
- **Example eBPF Program:** Includes a sample eBPF program (example.c) demonstrating BTF integration.

## Repository Structure

- **Makefile:** Main build script handling generation and compilation.
- **Makefile.btfgen:** Handles the generation of minimal BTF files.
- **example.c:** Sample eBPF program demonstrating BTF usage.
- **main.go:** Go program interfacing with the eBPF program.
- **btfgen.go, btfgen_amd64.go, btfgen_arm64.go:** Go files managing BTF generation for different architectures.
- **vmlinux.h, vmlinux_missing.h:** Header files required for BTF and eBPF program compilation.

## How to use

First clone [btfhub-archive repository](https://github.com/aquasecurity/btfhub-archive).

Then set path to it using:
```
export BTFHUB_ARCHIVE=~/btfhub-archive
```

You are now ready to generate minimal BTF files and build the output binary using:
```
make
```
