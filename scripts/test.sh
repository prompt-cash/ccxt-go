#!/usr/bin/env bash
set -e

if [ ! -d "$PWD/scripts" ]; then
  echo "Please run this shell script from project's root folder."
  exit 0
fi

mkdir -p "$PWD/testlog"

PKGS=(
  "app"
  "pkg/parser"
)

if [ -z "$1" ]; then
  echo "running test for all packages"
else
  PKGS=()
  PKGS=$@
fi

go clean -testcache
for PKG in "${PKGS[@]}"; do
  COVERFILENAME=$(echo $PKG | sed 's/\//_/g')
  go test -v "$PWD/$PKG" -coverprofile "$PWD/testlog/$COVERFILENAME.out"
done
cd ..

# run the main test after (depends on data being created before)
#if [ -z "$1" ]; then
  #go test -v "$PWD/tests" -coverprofile "$PWD/testlog/main.out"
#fi

echo "To view code coverage in your browser: go tool cover -html=testlog/{PACKAGE}.out"