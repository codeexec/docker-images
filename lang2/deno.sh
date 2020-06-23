#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/deno.toml
echo "installing deno"

curl -fsSL https://deno.land/x/install/install.sh | DENO_INSTALL=/usr sh -s v1.1.0
npm i -g typescript-language-server@0.4.0 typescript@3.9.5
