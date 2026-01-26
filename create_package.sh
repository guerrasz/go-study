#!/bin/bash

# This script creates a new package and .go file in the specified directory
# Usage: ./create_package.sh <directory> <package>

if [ $# -ne 2 ]; then
    echo "Error -> correct usage: $0 <directory> <package>"
    exit 1
fi

directory=$1
package=$2

# find dir that starts with $1
parent=$(find . -maxdepth 1 -name "${directory}_*" | head -n 1)

if [ -z "$parent" ]; then
    echo "Error: Directory prefix '${directory}' not found."
    exit 1
fi

# find highest number in directory
highest=$(ls $parent | sort | tail -1 | cut -c 1 -c 2)

# add 1 to highest
current=$((10#$highest + 1))

# create new dir with name $2
mkdir -p ${parent}/${current}_${package}

# create .go file with name $2
touch ${parent}/${current}_${package}/${package}.go

# add main function to .go file
printf "package main\n\nfunc main() {\n\n}" > ${parent}/${current}_${package}/${package}.go