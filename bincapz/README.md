# Bincapz

This is an implemenation of the [Bincapz](https://github.com/chainguard-dev/bincapz) tool.

## Usage

### Inspect binary

```sh
dagger call inspect-binary --binary "samples/macOS/2023.3CX/libffmpeg.dylib"
```

### Diff

```sh
dagger call diff --old-binary "samples/macOS/2023.3CX/libffmpeg.dirty.dylib" --new-binary "samples/macOS/2023.3CX/libffmpeg.dylib"
```
