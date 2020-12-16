#!/bin/sh
# dpw@alameda.local
# 2018.02.25
#

[ $# -eq 0 ] && {
    echo "USE: $0 file"
    exit 1
}

file=$1
target=`echo $file | sed -e 's/.go$//'`

[ -d bin ] || mkdir bin

go build -o bin/$target $file && ./bin/$target

