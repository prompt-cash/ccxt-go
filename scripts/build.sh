#!/bin/bash
set -e

if [ ! -d "$PWD/scripts" ]; then
  echo "Please run this shell script from project's root folder."
  exit 0
fi

mkdir -p "$PWD/bin"
go build -o "$PWD/bin/transpiler"