#!/bin/sh

gf version || exit 1

rm internal/packed/build-pack-data.go

gf build -n oh-my-duo-bin

cp ./temp/v1.0.0/linux_amd64/oh-my-duo-bin oh-my-duo-linux-amd64
cp ./temp/v1.0.0/windows_amd64/oh-my-duo-bin.exe oh-my-duo-windows-amd64.exe
