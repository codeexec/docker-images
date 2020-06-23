#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/lua.toml
echo "installing lua"

# http://www.lua.org/download.html
wget -q http://www.lua.org/ftp/lua-5.3.5.tar.gz
tar zxf lua-5.3.5.tar.gz
cd lua-5.3.5
make linux install
cd ..

# https://luarocks.org/
wget -q https://luarocks.org/releases/luarocks-3.3.1.tar.gz
tar zxpf luarocks-3.3.1.tar.gz
cd luarocks-3.3.1
./configure
make
make install
cd ..

rm -rf lua*
