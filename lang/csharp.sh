#!/bin/bash
set -eux

# https://github.com/replit/polygott/blob/master/languages/csharp.toml
# https://docs.microsoft.com/en-us/dotnet/core/install/linux-ubuntu#1804-

echo "installing .NET Core 3.1"

wget -q https://packages.microsoft.com/config/ubuntu/18.04/packages-microsoft-prod.deb -O packages-microsoft-prod.deb
dpkg -i packages-microsoft-prod.deb
apt-get update
apt-get install -y dotnet-sdk-3.1 aspnetcore-runtime-3.1
