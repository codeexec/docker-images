#!/bin/bash
set -eux

# https://bellard.org/quickjs/
VER=2020-09-06
echo "installing quickjs ${VER}"

wget -q https://bellard.org/quickjs/quickjs-${VER}.tar.xz
tar -xf quickjs-${VER}.tar.xz
cd quickjs-${VER}
make
make install
cd ..

rm -rf quickjs-${VER}*
