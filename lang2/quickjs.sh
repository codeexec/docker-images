#!/bin/bash
set -eux

# https://bellard.org/quickjs/
echo "installing quickjs"

# http://www.lua.org/download.html
wget -q https://bellard.org/quickjs/quickjs-2020-04-12.tar.xz
tar -xf quickjs-2020-04-12.tar.xz
cd quickjs-2020-04-12
make
make install
cd ..

rm -rf quickjs-2020-04-12*
