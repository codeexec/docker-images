#!/bin/bash
set -eux

# https://docs.microsoft.com/en-us/dotnet/core/install/linux-ubuntu#1804-

echo "installing .NET Core 3.1"

wget -q https://packages.microsoft.com/config/ubuntu/20.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb
dpkg -i packages-microsoft-prod.deb
apt-get update
apt-get install -y dotnet-sdk-3.1 aspnetcore-runtime-3.1

# run dotnet once to get rid of the "first time running" output
# https://github.com/dotnet/sdk/issues/10411
dotnet new

rm -rf packages-microsoft-prod.deb
