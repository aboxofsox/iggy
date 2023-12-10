#!/bin/bash

URL='https://api.github.com/repos/aboxofsox/iggy/releases/latest'
RootDir="$HOME/iggy"

if [ ! -d "$RootDir" ]; then
    mkdir -p "$RootDir"
fi

LatestRelease=$(curl -s $URL)

Assets=$(echo $LatestRelease | jq -r '.assets[]')

for Asset in $Assets; do
    Name=$(echo $Asset | jq -r '.name')
    URL=$(echo $Asset | jq -r '.browser_download_url')

    Output="$RootDir/$Name"

    echo "Downloading $Name"

    curl -L $URL -o $Output
done

Path="$PATH:$RootDir"
export PATH="$Path"