#!/bin/bash

round() {
	printf %.2f $1
}
rand_range() {
	printf $(($1 + RANDOM%(1+$2-$1)))
}
epoch() {
	awk 'BEGIN {srand(); print srand()}'
}

target=http://localhost:8080
device=xxx
now=$(epoch)
gps=$(./gps)
lat=$(round ${gps%% *})
lng=$(round ${gps#* })
temperature=$(rand_range 1 200 | xxd -pu)
snr=1
station=2

queryString="x=$temperature&time=$now&snr=$snr&station=$station&lat=$lat&lng=$lng&device=$device"

curl -s $target/push?$queryString \
	-H 'Accept-Encoding: gzip, deflate, br' \
	-H 'Accept-Language: en-US,en;q=0.8,cs;q=0.6,ru;q=0.4,it;q=0.2,de;q=0.2' \
	-H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/60.0.3112.113 Safari/537.36' \
	-H 'Accept: text/plain,*/*;q=0.8' \
	-H 'Cache-Control: max-age=0' \
	-H 'Connection: keep-alive' \
	--compressed > /dev/null