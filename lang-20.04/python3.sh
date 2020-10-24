#!/bin/bash
set -eux

echo "installing python3"

# https://github.com/replit/polygott/blob/master/languages/python3.toml


apt-get install -y python3-pip python3-wheel python3-dev
