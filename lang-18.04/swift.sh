#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/swift.toml
echo "installing swift"

# https://swift.org/download/#releases

VER=5.2.4
wget -q https://swift.org/builds/swift-${VER}-release/ubuntu1804/swift-${VER}-RELEASE/swift-${VER}-RELEASE-ubuntu18.04.tar.gz

# TODO: verify sig
#wget -q https://swift.org/builds/swift-${VER}-release/ubuntu1804/swift-${VER}-RELEASE/swift-${VER}-RELEASE-ubuntu18.04.tar.gz.sig

tar -zxf swift-${VER}-RELEASE-ubuntu18.04.tar.gz --strip-components=1 -C /
rm swift-*
chmod -R go+r /usr/lib/swift
