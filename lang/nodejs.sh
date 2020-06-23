#!/bin/bash
set -eux

# nodejs 14.x

curl -sL https://deb.nodesource.com/setup_14.x | bash -
apt-get install -y nodejs

# npm install -g jest@23.1.0 prettier@1.13.4 babylon@6.15 babel-traverse@6.21 walker@1.0.7

# TODO: use just yarn?
npm install -g prettier@2.0.5 yarn@1.22.4
