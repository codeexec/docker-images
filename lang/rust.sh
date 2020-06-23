#!/bin/bash
set -eux

# https://forge.rust-lang.org/infra/other-installation-methods.html
# https://github.com/replit/polygott/blob/master/languages/rust.toml

curl https://sh.rustup.rs -sSf | sh -s -- -y
