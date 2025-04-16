package main

//go:generate go run github.com/cilium/ebpf/cmd/bpf2go -target bpfel example example.c

import (
        "log"
        "os"
        "os/signal"
        "syscall"

        "github.com/cilium/ebpf"
        "github.com/cilium/ebpf/link"
        "github.com/cilium/ebpf/rlimit"
)

func main() {
        if err := rlimit.RemoveMemlock(); err != nil {
                log.Fatalf("Failed to remove rlimit memlock: %v", err)
        }

        opts := ebpf.CollectionOptions{
                Programs: ebpf.ProgramOptions{
			// This is where the BTF Spec is loaded
                        KernelTypes: GetBTFSpec(),
                },
        }

        var objs exampleObjects
        if err := loadExampleObjects(&objs, &opts); err != nil {
                log.Fatalf("Loading eBPF objects: %v", err)
        }
        defer objs.Close()

        tp, err := link.Tracepoint("syscalls", "sys_enter_execve", objs.TracepointProgram, nil)
        if err != nil {
                log.Fatalf("Attaching Tracepoint: %s", err)
        }
        defer tp.Close()

        log.Println("eBPF program attached to tracepoint. Press Ctrl+C to exit.")

        // Set up signal catching
        sig := make(chan os.Signal, 1)
        signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

        // Wait for a signal
        <-sig

        log.Println("Received signal, exiting gracefully.")
}
