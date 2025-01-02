BTFHUB_ARCHIVE ?= $(HOME)/btfhub
ARCH ?= $(shell uname -m | sed 's/x86_64/x86/' | sed 's/aarch64/arm64/' | sed 's/ppc64le/powerpc/' | sed 's/mips.*/mips/')
OUTPUT ?= btfs
EBPF_SRC ?= $(wildcard *.c)  # eBPF source files (modify pattern as needed)
EBPF_OBJ ?= $(EBPF_SRC:.c=.o)  # eBPF object files

# Default target
.PHONY: all
all: generate btfgen build

# Compile eBPF programs
.PHONY: generate
generate: 
	@echo "Compiling eBPF programs"
	@go generate

# Invoke Makefile.btfgen
.PHONY: btfgen
btfgen:
	@echo "Invoking Makefile.btfgen with BTFHUB_ARCHIVE=$(BTFHUB_ARCHIVE), ARCH=$(ARCH), OUTPUT=$(OUTPUT)"
	@$(MAKE) -f Makefile.btfgen BTFHUB_ARCHIVE=$(BTFHUB_ARCHIVE) ARCH=$(ARCH) OUTPUT=$(OUTPUT)

# Build Go programs
.PHONY: build
build: 
	@echo "Building Go programs"
	@go build

# Clean target
.PHONY: clean
clean:
	@echo "Cleaning output directory: $(OUTPUT)"
	@rm -rf $(OUTPUT) $(EBPF_OBJ)
