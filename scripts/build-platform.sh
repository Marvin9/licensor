#!/usr/bin/env bash

rm -rf ./build
mkdir build

# os=$(go tool dist list)
os=(
    linux/386
    linux/amd64
    linux/arm
    linux/arm64
    windows/386
    windows/amd64
)

function log() {
    echo $1
}

SECONDS=0
printf "STARTED BUILDING BINARIES:\n\n"

for os_arch in "${os[@]}"
do
    GOOS=${os_arch%/*}
    if [[ "$GOOS" == "windows" || "$GOOS" == "linux" ]]; then
        GOARCH=${os_arch#*/}
        FILE_EXT=""
        if [[ "$GOOS" == "windows" ]];then
            FILE_EXT=".exe"
        fi

        FILE_NAME="licensor_${GOOS}_${GOARCH}${FILE_EXT}"
        log "BUILDING $FILE_NAME..."

        built=$(GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./build/${FILE_NAME})
        
        is=$([ $? -eq 0 ] && echo "✔️" || echo "❌")
        printf '\e[A\e[K'
        log "$FILE_NAME $is"
    fi
done

printf "\n\n"
duration=$SECONDS

min=$(($duration/60))
sec=$(($duration%60))

function _time() {
    val=$([ $1 -gt 9 ] && echo "$1" || echo "0$1")
    printf "$val $2"
}

_time $min "Minutes and "
_time $sec "Seconds elapsed"

printf "\n"