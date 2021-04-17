# Dummy files generator

```sh
$ dummy --help
 Dummy file generator, creates files with random names

Usage:
  dummy [flags]

Flags:
  -d, --destinationFolder string   Destination folder
  -x, --fileExtension string       Files will be generated with this file ext
  -h, --help                       help for dummy
  -m, --megas int                  Megas for each file (default 10)
  -n, --numFiles int               Number of files to generate (default 1)
```

## build
---
Linux with make and cross compilation
```sh
$ make
go build -o bin/linux/dummy main.go
GOOS=windows GOARCH=386 go build -o bin/windows/dummy.exe main.go
```
Windows
```cmd
$ go build -o bin/windows/dummy.exe main.go
```