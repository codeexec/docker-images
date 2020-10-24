#!/bin/bash
set -eux

# https://dartdoc.takyam.com/downloads/linux.html
# https://github.com/replit/polygott/blob/master/languages/dart.toml

echo "installing dart dev"

curl -s https://dl-ssl.google.com/linux/linux_signing_key.pub | apt-key add -
curl -s https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_stable.list > /etc/apt/sources.list.d/dart_stable.list
curl -s https://storage.googleapis.com/download.dartlang.org/linux/debian/dart_unstable.list > /etc/apt/sources.list.d/dart_unstable.list
apt-get update
apt-get install dart

# TODO: maybe /usr/local/bin/pub ?
ln -s /usr/lib/dart/bin/pub /usr/bin/pub
