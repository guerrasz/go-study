#!/bin/bash

# This script creates a new package and .go file in the specified directory
# Usage: ./create_package.sh <directory> <package>

if [ $# -ne 2 ]; then
    echo "Error -> correct usage: $0 <directory> <package>"
    exit 1
fi

directory=$1
package=$2

mkdir -p $directory/$package
touch $directory/$package/$package.go