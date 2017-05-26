#!/bin/bash

# This script tests multiple packages and creates a consolidated cover profile

#function die() {
#  echo $*
#  exit 1
#}

# Initialize coverage.txt
echo "mode: count" > coverage.txt

# Initialize error tracking
ERROR=""

# Test each package and append coverage profile info to coverage.txt
PACKAGES=$(govendor list +l,^p | while read line; do echo $line | awk '{print $2}'; done)

for pkg in $PACKAGES
do
    go test -v -covermode=count -coverprofile=profile_tmp.cov $pkg #|| ERROR="Error testing $pkg"
    tail -n +2 profile_tmp.cov >> coverage.txt #|| die "Unable to append coverage for $pkg"
done

rm profile_tmp.cov

if [ ! -z "$ERROR" ]
then
    die "Encountered error, last error was: $ERROR"
fi

curl -s https://codecov.io/bash | bash -s arg1 arg2
