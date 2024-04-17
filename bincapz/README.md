# Bincapz

This is an implemenation of the [Bincapz](https://github.com/chainguard-dev/bincapz) tool.

## Usage

```sh
dagger functions
```

### Local

#### Inspect binary

```sh
dagger call inspect-binary --binary "samples/macOS/2023.3CX/libffmpeg.dylib"
```

#### Diff

```sh
dagger call diff --old-binary "samples/macOS/2023.3CX/libffmpeg.dirty.dylib" --new-binary "samples/macOS/2023.3CX/libffmpeg.dylib"
```

### Remote

```sh
dagger -m github.com/denhamparry/daggerverse/bincapz@v0.1 functions
```

#### Remote - Inspect binary

```sh
dagger -m github.com/denhamparry/daggerverse/bincapz@v0.1 call inspect-binary --binary ""
```

#### Remote - Diff

```sh
dagger -m github.com/denhamparry/daggerverse/bincapz@v0.1 call diff --old-binary "samples/macOS/2023.3CX/libffmpeg.dirty.dylib" --new-binary "samples/macOS/2023.3CX/libffmpeg.dylib"
```
