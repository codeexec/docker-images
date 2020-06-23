#!/bin/bash
set -eux

# go 1.14.2
# https://github.com/replit/polygott/blob/master/languages/go.toml

echo "installing go"

add-apt-repository -y ppa:longsleep/golang-backports
apt-get install -y pkg-config golang-1.14-go
ln -s /usr/lib/go-1.14/bin/go /usr/bin/go
ln -s /usr/lib/go-1.14/bin/gofmt /usr/bin/gofmt
