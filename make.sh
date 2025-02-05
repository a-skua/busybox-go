#!/usr/bin/env sh

command=$1

for dir in $(ls -d */); do
	make $command -C $dir
done