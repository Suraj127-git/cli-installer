#!/bin/bash

# Build for macOS (Intel)
export GOOS=darwin
export GOARCH=amd64
go build -o cli-installer-macos
if [ $? -eq 0 ]; then
    echo "Build for macOS (Intel) successful"
    zip cli-installer-macos.zip cli-installer-macos
else
    echo "Build for macOS (Intel) failed"
fi

# Build for Windows 11 (64-bit)
export GOOS=windows
export GOARCH=amd64
go build -o cli-installer-windows.exe
if [ $? -eq 0 ]; then
    echo "Build for Windows 11 (64-bit) successful"
    zip cli-installer-windows.zip cli-installer-windows.exe
else
    echo "Build for Windows 11 (64-bit) failed"
fi

