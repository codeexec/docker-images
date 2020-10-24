#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/nodejs.toml

echo "installing nodejs 15.x"

# nodejs 15.x

curl -sL https://deb.nodesource.com/setup_15.x | bash -
apt-get install -y nodejs

# npm install -g jest@23.1.0 prettier@1.13.4 babylon@6.15 babel-traverse@6.21 walker@1.0.7

# TODO: use just yarn?
# https://www.npmjs.com/package/prettier
# https://www.npmjs.com/package/yarn
# https://github.com/npm/cli
npm install -g prettier@2.1.2 yarn@1.22.10 npm@7.0.5
