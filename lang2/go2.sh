
#!/bin/bash
set -eux

# go 1.14.2
# https://golang.org/doc/install

echo "installing go 1.14.4"

wget -q https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.14.4.linux-amd64.tar.gz

ln -s /usr/local/go/bin/go /usr/bin/go
ln -s /usr/local/go/bin/gofmt /usr/bin/gofmt

rm go1.14.4.linux-amd64.tar.gz
