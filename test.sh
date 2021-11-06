#!/bin/bash
if [ -f build/aptly ]; then
  rm -v build/aptly
fi
GOCACHE=/tmp/gocache go build -o build/aptly
if [ -f build/aptly ]; then
  ./build/aptly repo copy1 test-main-repo from package 'Name (=qemu-utils), Version (=1:3.1+dfsg.1-1+dde)'
  ./build/aptly repo copy1 test-main-repo from package 'Name (=qemu-utils)'
fi
