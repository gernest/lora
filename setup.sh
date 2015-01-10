#!/bin/bash
# Copyright 2015 Geofrey Ernest a.k.a gernest, All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License"): you may
#  not use this file except in compliance with the License. You may obtain
#  a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.


#if dumb terminal (file browser) run xterm so basicaly the script must run on almost any distro
if [ $TERM == "dumb" ]; then xterm -hold -e $0; fi

clear

echo  "Golang Programming Environment Installer"

#get last version of go compiler (e.g. go1.3.3.)
v=`echo $(wget -qO- golang.org) | awk '{ if (match($0,/go([0-9]+.)+/)) print substr($0,RSTART,RLENGTH) }'`

#B0003
if [ -z "$v" ]; then
   echo "No network connection"
   exit
fi

#get host computer arch (e.g. i686|amd64)
#if [[ $(uname -i) == "i386" ]]; then a="386"; else a="amd64"; fi
#B0002,B0007
case $(uname -m) in
i686 ) a="386";;
   * ) a="amd64"
esac

#get kernel name (e.g. linux|freebsd)
k=$(uname -s | tr '[:upper:]' '[:lower:]')

#B0005
test -f ${XDG_CONFIG_HOME:-~/.config}/user-dirs.dirs && source ${XDG_CONFIG_HOME:-~/.config}/user-dirs.dirs

#build compiler name (e.g. go1.3.3.linux-386.tar.gz)
n=${v}${k}-${a}.tar.gz

echo "Download last compiler $n..."

wget --no-check-certificate -Nq -P ${XDG_DOWNLOAD_DIR} https://storage.googleapis.com/golang/$n
#rd -r go
echo "Unpack..."
tar -xf ${XDG_DOWNLOAD_DIR}/$n -C $HOME


#get host computer LONG_BIT (e.g 32|64)
a=$(getconf LONG_BIT)



echo "Unpack..."
tar -xf ${XDG_DOWNLOAD_DIR}/$n -C $HOME

bashrc=
echo "Creatting \$GOPATH"
GOPATH=$HOME/gosrc
mkdir -p $GOPATH/src
mkdir -p $GOPATH/bin
mkdir -p $GOPATH/pkg

echo "Finished installing"
exit 0