# Path to the btfhub-archive repository (needs to be present)
BTFHUB_ARCHIVE ?= $(HOME)/btfhub-archive

# Determine the architecture for the current machine.
# Normalize common architecture names to match BTFHub naming convention.
ARCH ?= $(shell uname -m | sed 's/x86_64/x86/' | sed 's/aarch64/arm64/' | sed 's/ppc64le/powerpc/' | sed 's/mips.*/mips/')

# Output directory or filename prefix
OUTPUT ?= btfs

# Find all eBPF C source files in the current directory
EBPF_SRC ?= $(wildcard *.c)  # eBPF source files (modify pattern as needed)

# Replace `.c` with `.o` to get the list of eBPF object files
EBPF_OBJ ?= $(EBPF_SRC:.c=.o)  # eBPF object files

# Default target that runs both generate and build
.PHONY: all
all: generate btfgen build

# Compile eBPF programs using `go generate`
.PHONY: generate
generate: 
	@echo "Compiling eBPF programs"
	@go generate

# Generate BTF files if your program is expected to run on non-BTF-enabled kernels.
# This is optional and not invoked by default.
.PHONY: btfgen
btfgen:
	@echo "Invoking Makefile.btfgen with BTFHUB_ARCHIVE=$(BTFHUB_ARCHIVE), ARCH=$(ARCH), OUTPUT=$(OUTPUT)"
	@$(MAKE) -f Makefile.btfgen BTFHUB_ARCHIVE=$(BTFHUB_ARCHIVE) ARCH=$(ARCH) OUTPUT=$(OUTPUT)

# Build the Go programs in the current directory
.PHONY: build
build: 
	@echo "Building Go programs"
	@go build

# Clean up generated output and eBPF object files
.PHONY: clean
clean:
	@echo "Cleaning output directory: $(OUTPUT)"
	@rm -rf $(OUTPUT) $(EBPF_OBJ)
