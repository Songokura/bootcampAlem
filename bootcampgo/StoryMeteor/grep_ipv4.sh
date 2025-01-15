#!/bin/bash


if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

filename="$1"


if [ ! -f "$filename" ]; then
    echo "File not found: $filename"
    exit 1
fi


ipv4_regex='([0-9]{1,3}\.){3}[0-9]{1,3}'


grep -oE "$ipv4_regex" "$filename" | while read -r ip; do
    valid=true
    IFS='.' read -r -a octets <<< "$ip"
    for octet in "${octets[@]}"; do
        if ((octet < 0 || octet > 255)); then
            valid=false
            break
        fi
    done
    if $valid; then
        echo "$ip"
    fi
done