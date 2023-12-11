#!/bin/bash

Platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386" "darwin/amd64")
length=${#Platforms[@]}

for Platform in "${Platforms[@]}"; do
    GOOS=${Platform%%/*}
    Arch=${Platform##*/}

    export GOOS=$GOOS
    export GOARCH=$Arch

    Output="bin/$GOOS/$Arch/iggy"

    echo "Building iggy: Building for $GOOS/$Arch."

    if [ $GOOS = 'windows' ]; then
        Output+='.exe'
    fi

    go build -o $Output

    if [ $? -ne 0 ]; then
        echo "Build failed"
        exit 1
    fi

done