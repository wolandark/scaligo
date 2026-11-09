#!/bin/sh

mkdir -p build
for os_arch in "linux amd64" "windows amd64" "darwin arm64" "darwin amd64"; do
	set -- $os_arch
	ext=""; [ "$1" = "windows" ] && ext=".exe"
	GOOS=$1 GOARCH=$2 go build -o build/scaligo_${1}_${2}${ext} .
done
