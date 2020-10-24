#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/deno.toml
# https://github.com/denoland/deno_install
echo "installing deno"

export DENO_INSTALL=/usr/local
curl -fsSL https://deno.land/x/install/install.sh | sh
#curl -fsSL https://deno.land/x/install/install.sh | DENO_INSTALL=/usr sh -s v1.4
npm i -g typescript-language-server@0.4.0 typescript@3.9.5
