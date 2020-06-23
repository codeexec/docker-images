#!/bin/bash
set -eux

echo "installing clang 10"

apt-get install -y clang clang-format lldb lld clangd
