#!/bin/sh

gf version || exit 1

rm internal/packed/data.go

gf build -ps resource -pd internal/packed/data.go -n oh-my-duo-bin
