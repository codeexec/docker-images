#!/bin/bash
set -eux

echo "installing deno"

npm i -g typescript-language-server@0.4.0 typescript@3.9.5
curl -fsSL https://deno.land/x/install/install.sh | DENO_INSTALL=/usr sh -s v1.1.0
