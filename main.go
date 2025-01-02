package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpfel example example.c

import (
	"fmt"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func main() {
	// Open the eBPF object file (compiled .o file)
	bpfFile := "example_bpfel.o"
	prog, err := ebpf.LoadCollection(bpfFile)
	if err != nil {
		log.Fatalf("Failed to load BPF program: %v", err)
	}
	defer prog.Close()

	// Find the tracepoint program inside the loaded collection
	tracepointProg, ok := prog.Programs["tracepoint_program"]
	if !ok {
		log.Fatalf("tracepoint program not found in the collection")
	}

	// Attach the BPF program to the tracepoint
	// Example tracepoint: syscalls/sys_enter_execve
	tracepoint, err := link.Tracepoint("syscalls", "sys_enter_execve", tracepointProg, nil)
	if err != nil {
		log.Fatalf("Failed to attach to tracepoint: %v", err)
	}
	defer tracepoint.Close()

	fmt.Println("eBPF program attached to tracepoint. Press Ctrl+C to exit.")

	// Run the program and handle events
	// For tracepoints, you may need to capture output or process events
	// Here we simulate a running program by sleeping indefinitely
	select {}
}
