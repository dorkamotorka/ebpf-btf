package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpfel example example.c

import (
	"fmt"
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func main() {
	opts := ebpf.CollectionOptions{
		Programs: ebpf.ProgramOptions{
			KernelTypes: GetBTFSpec(),
		},
	}

	var objs exampleObjects
	if err := loadExampleObjects(&objs, &opts); err != nil {
		log.Fatal("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// Attach Tracepoint
	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.TracepointProgram, nil)
	if err != nil {
		log.Fatalf("Attaching Tracepoint: %s", err)
	}
	defer tp.Close()

	fmt.Println("eBPF program attached to tracepoint. Press Ctrl+C to exit.")

	// Run the program and handle events
	// For tracepoints, you may need to capture output or process events
	// Here we simulate a running program by sleeping indefinitely
	select {}
}
