// Bincapz - Enumerates program capabilities and malicious behaviors using fragment analysis.
//
// Features
// - Analyzes binaries from any architecture - arm64, amd64, riscv, ppc64, sparc64
// - Supports scripting languages such as bash, PHP, Perl, Ruby, NodeJS, and Python
// - Integrates YARA forge for rules by Avast, Elastic, FireEye, Google, Nextron, and others.
// - 12,000+ rules that detect everything from ioctl's to malware
// - Tuned for especially excellent performance with Linux programs
// - Diff-friendly output in Markdown, JSON, YAML outputs
// - CI/CD friendly
//
// Shortcomings
// - Does not attempt to process archive files (jar, zip, apk)
// - Minimal rule support for Windows and Java (help wanted!)
// - Early in development; output is subject to change

package main

import (
	"context"
)

type Bincapz struct{}

// To inspect a binary, pass it as an argument to dump a list of predicted capabilities
func (m *Bincapz) InspectBinary(ctx context.Context, binary *File) (string, error) {
	return dag.Container().
		From("cgr.dev/chainguard/bincapz").
		WithMountedFile("/tmp/binary", binary).
		WithExec([]string{"/tmp/binary"}).
		Stdout(ctx)
}

// Make sure an update doesn't introduce unexpected capability changes
func (m *Bincapz) Diff(ctx context.Context, old_binary *File, new_binary *File) (string, error) {
	return dag.Container().
		From("cgr.dev/chainguard/bincapz").
		WithMountedFile("/tmp/old_binary", old_binary).
		WithMountedFile("/tmp/new_binary", new_binary).
		WithExec([]string{"--diff", "/tmp/old_binary", "/tmp/new_binary"}).
		Stdout(ctx)
}
