#!/bin/bash
set -euf -o pipefail

export AOC_SESSION=53616c7465645f5f040426528978a5d0b08fde3fdc9ef4205e55508a1b81e56bbb1e65904d8315a4a81d6a3b73988510fbe434c158086ee6a80e9cff876a89b7

# functions
function echogrey() {
	echo -e "\033[0;90m$1\033[0m"
}

function template() {
	cat <<EOF
package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// solve part 1 here
	return 42
}
EOF
}

# two args YEAR and DAY
YEAR="${1:-}"
DAY="${2:-}"
if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
	echo "Usage: $0 <YEAR> <DAY>"
	exit 1
fi
# pad DAY to 2 digits
DAY=$(printf "%02d" $DAY)
DIR="./$YEAR/$DAY"
# create missing files as needed
if [ ! -d "$DIR" ]; then
	mkdir -p "$DIR"
	echogrey "Created directory $DIR"
fi
if [ ! -f "$DIR/code.go" ]; then
	template >"$DIR/code.go"
	echogrey "Created file code.go"
fi
# go run
cd "$DIR" && go run code.go
