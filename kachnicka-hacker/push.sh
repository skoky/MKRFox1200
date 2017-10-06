#!/bin/bash

n=$1

[ -z $n ] && {
    echo "argument repeat not provided\nusage: $(basename $0) <:repeat>"
    exit 0
}

epoch() {
	awk 'BEGIN {srand(); print srand()}'
}
calc() {
  input="$*"
  sed -e 's/,/./g' <<< "scale=35; ${input//[[:blank:]]/}" | bc -l | sed 's/^\./0./;s/0*$//;s/\.$//'
}

now=$(epoch)

{
    for f in $(seq 1 $n); do
        sh push_one.sh -ts=$(calc "$now + (720 * $f)") &
    done
} & wait