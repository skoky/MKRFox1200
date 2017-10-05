#!/bin/bash

n=$1

[ -z $n ] && {
    echo "argument repeat-each-sec not provided\nusage: $(basename $0) <:repeat-each-sec>"
    exit 0
}

watch -n$n sh push_one.sh > /dev/null
