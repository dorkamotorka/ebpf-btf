package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpfel example example.c

import (
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
		log.Fatalf("Loading eBPF objects:", err)
	}
	defer objs.Close()

	// Attach Tracepoint
	tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.TracepointProgram, nil)
	if err != nil {
		log.Fatalf("Attaching Tracepoint: %s", err)
	}
	defer tp.Close()

	log.Println("eBPF program attached to tracepoint. Press Ctrl+C to exit.")

	select {}
}
