#!/usr/bin/env bash
set -ueo pipefail
rundir=$(cd -P -- "$(dirname -- "$0")" && printf '%s\n' "$(pwd -P)")
cd "$rundir"

# Custom run:

CONTAINER_NAME="shares"
IMAGE_NAME="${CONTAINER_NAME}"
IMAGE_REPO="me/${IMAGE_NAME}"
IMAGE_TAG="latest"

run:nodemon () {
  nodemon -e go,sum,json -x "go run shares.go -k 6 -t 4 || exit 1"
}

run:build () {
  GOOS=darwin GOARCH=amd64 go build -o shares-darwin-amd64
  GOOS=darwin GOARCH=arm64 go build -o shares-darwin-arm64 
  GOOS=linux GOARCH=amd64 go build -o shares-linux-amd64
  GOOS=windows GOARCH=amd64 go build -o shares-windows-amd64.exe
}

# Common run:

run:help(){
  set +x
  printf "Commands:\n"
  declare -F | awk '/^declare -f run:/ { printf("  %s\n", substr($0,16)); }'
  exit 0
}

[ -z "${1:-}" ] && run:help

cmd=$1
shift

set -x
run:$cmd "$@"

