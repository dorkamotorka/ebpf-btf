//go:build ignore
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_core_read.h>

#define ARGSIZE 256

SEC("tracepoint/syscalls/sys_enter_execve")
int tracepoint_program(struct trace_event_raw_sys_enter *ctx) {
    char *filename_ptr = (char *)BPF_CORE_READ(ctx, args[0]);

    u8 filename[ARGSIZE];
    bpf_core_read_user_str(&filename, sizeof(filename), filename_ptr);

    bpf_printk("Tracepoint (CO-RE) triggered for execve syscall with parameter filename: %s\n", filename);
    return 0;
}

char _license[] SEC("license") = "GPL";

