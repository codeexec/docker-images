FROM ubuntu:20.04

WORKDIR /home/runner
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update
# https://github.com/replit/polygott/blob/master/packages.txt + more
# some packages are needed for later scripts
RUN apt-get install -y locales curl wget xz-utils unzip git subversion mercurial bzr build-essential sqlite3 libsqlite3-dev gnupg libc6-dbg libssl-dev libreadline-dev libicu-dev libevent-dev  libedit2 libxml2 lsb-release software-properties-common apt-utils apt-transport-https pkg-config cmake ninja-build

# versions in 20.04:
# gcc 9.3
# lua 5.3
# luajit 2.1.0-beta3
# python 2.7.18rc1
# python3 3.8.2
# perl 5.30
# php 7.4.3
# nodejs 10.19.0
# go 1.13.8
# ruby 2.7.0

ADD lang lang
RUN chmod ug+x lang/*.sh

RUN lang/clang-llvm.sh
ENV DOTNET_CLI_TELEMETRY_OPTOUT=1
RUN lang/csharp.sh
RUN lang/dart.sh
# must install before deno.sh
RUN lang/nodejs.sh
RUN lang/deno.sh
RUN lang/go2.sh
RUN lang/java.sh
RUN lang/lua.sh
RUN lang/luajit.sh
RUN lang/php.sh
RUN lang/python3.sh
RUN lang/quickjs.sh
RUN lang/ruby.sh
RUN lang/rust.sh
RUN lang/swift.sh

RUN rm -rf lang
