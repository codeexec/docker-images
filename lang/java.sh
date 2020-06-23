#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/java.toml

echo "installing java"

apt-get install -y openjdk-11-jdk maven
