FROM ubuntu:18.04

WORKDIR /home/runner
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update
# https://github.com/replit/polygott/blob/master/packages.txt + more
# some packages are needed for later scripts
RUN apt-get install -y ca-certificates locales xz-utils unzip curl wget git subversion mercurial sqlite3 libsqlite3-dev gnupg build-essential libc6-dbg libssl-dev libreadline-dev libicu-dev libxml2 libedit2  lsb-release software-properties-common apt-utils apt-transport-https libpython2.7 python-pip python-wheel python-dev libevent-dev cmake ninja-build tzdata php-cli php-pear -y

# versions in 18.04:
# gcc 7.5
# python 2.7.17
# python3 3.6.9

# chromium-chromedriver valgrind 

ADD lang lang
RUN chmod ug+x lang/*.sh

# RUN lang/python2.sh
# RUN lang/python3.sh

# php 7.2
# RUN lang/php.sh

# quickjs
RUN lang/quickjs.sh

# nodejs 14.x
RUN lang/nodejs.sh

# deno 1.1
RUN lang/deno.sh

# ruby 2.7.1
RUN lang/ruby.sh

# swift 5.2.4
RUN lang/swift.sh

# dart 2.9.0-18.0.dev
RUN lang/dart.sh

# lua 5.3.5
RUN lang/lua.sh

# clang 10
RUN lang/clang-llvm.sh

# go 1.14.2
RUN lang/go.sh

# java 11.0.7
RUN lang/java.sh

RUN rm -rf lang
