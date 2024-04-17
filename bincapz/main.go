package main

import (
	"context"
)

type Bincapz struct{}

// Outputs an inspection of a binary using Bincapz
func (m *Bincapz) InspectBinary(ctx context.Context, binary *File) (string, error) {
	return dag.Container().
		From("cgr.dev/chainguard/bincapz").
		WithMountedFile("/tmp/binary", binary).
		WithExec([]string{"/tmp/binary"}).
		Stdout(ctx)
}

// Outputs a diff of two binaries using Bincapz
func (m *Bincapz) Diff(ctx context.Context, old_binary *File, new_binary *File) (string, error) {
	return dag.Container().
		From("cgr.dev/chainguard/bincapz").
		WithMountedFile("/tmp/old_binary", old_binary).
		WithMountedFile("/tmp/new_binary", new_binary).
		WithExec([]string{"--diff", "/tmp/old_binary", "/tmp/new_binary"}).
		Stdout(ctx)
}
