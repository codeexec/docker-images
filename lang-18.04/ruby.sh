#!/bin/bash

# https://github.com/replit/polygott/blob/master/languages/ruby.toml

# /usr/share/rvm/scripts/rvm fails when this is set
#set -eux

echo "installing ruby"

# https://phoenixnap.com/kb/how-to-install-ruby-on-ubuntu-18-04
# https://linuxize.com/post/how-to-install-ruby-on-ubuntu-18-04/

apt-add-repository -y ppa:rael-gc/rvm
apt update
apt install rvm -y
source /usr/share/rvm/scripts/rvm
rvm install 2.7.1
rvm use 2.7.1 --default
