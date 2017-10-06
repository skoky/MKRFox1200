#!/bin/bash

unset {http,https,ftp}_proxy

round() {
	printf %.2f $1
}

rand_range() {
	printf $(bc -l <<< "scale=8; $1+$RANDOM*$2/32767")
}

epoch() {
	awk 'BEGIN {srand(); print srand()}'
}

location() {
	./gps 2> /dev/null || echo "50.0756082884077 14.4183340921249"
}

for i in "$@"; do
case $i in
    -ts=*|--timestamp=*)
    now="${i#*=}"
    shift
    ;;
    *)
    ;;
esac
done

target=http://localhost:8080
device=xxx
[ -z $now ] && now=$(epoch)
gps=$(location)
lat=$(round ${gps%% *})
lng=$(round ${gps#* })
temperature=$(rand_range 1 40 | xxd -pu)
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
	> /dev/null