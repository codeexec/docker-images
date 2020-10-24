
#!/bin/bash
set -eux

# https://golang.org/doc/install

VER=1.15.3

echo "installing go ${VER}"

wget -q https://dl.google.com/go/go${VER}.linux-amd64.tar.gz
tar -C /usr/local -xzf go${VER}.linux-amd64.tar.gz

#ln -s /usr/local/go/bin/go /usr/bin/go
#ln -s /usr/local/go/bin/gofmt /usr/bin/gofmt

rm go${VER}.linux-amd64.tar.gz
