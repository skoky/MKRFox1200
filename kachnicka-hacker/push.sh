#!/bin/bash

n=$1

[ -z $n ] && {
    echo "argument repeat not provided\nusage: $(basename $0) <:repeat>"
    exit 0
}

{
    for f in $(seq 1 $n); do
        sh push_one.sh &
    done
} & wait