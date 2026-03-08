#!/usr/bin/env bash
# Integration test: verify every command in README.md returns exit code 0.
# Run from the repo root. Builds the binary first.
set -euo pipefail

BINARY="./raku-cli"
PASS=0
FAIL=0
FAILED_CMDS=()

# Checkin/checkout dates always 30/31 days from today
CHECKIN=$(date -d "+30 days" +%Y-%m-%d 2>/dev/null || date -v+30d +%Y-%m-%d)
CHECKOUT=$(date -d "+31 days" +%Y-%m-%d 2>/dev/null || date -v+31d +%Y-%m-%d)

build() {
    echo "Building $BINARY ..."
    local go_bin
    if command -v go &>/dev/null; then
        go_bin=go
    elif [ -x "/home/$USER/go/bin/go" ]; then
        go_bin="/home/$USER/go/bin/go"
    elif [ -x /usr/local/go/bin/go ]; then
        go_bin=/usr/local/go/bin/go
    else
        echo "ERROR: go not found. Add it to PATH." >&2
        return 1
    fi
    "$go_bin" build -o "$BINARY" . 2>&1
}

run() {
    local desc="$1"; shift
    # Rate-limit: short pause between API calls
    sleep 1
    if "$BINARY" "$@" --quiet > /dev/null 2>&1; then
        echo "  PASS  $desc"
        ((PASS++)) || true
    else
        echo "  FAIL  $desc"
        ((FAIL++)) || true
        FAILED_CMDS+=("$desc")
    fi
}

build

echo ""
echo "Running API integration tests ..."
echo ""

# Ichiba
run "ichiba items --keyword ノートパソコン"        ichiba items --keyword "ノートパソコン"
run "ichiba items --min-price/max-price/hits"      ichiba items --keyword "本" --min-price 1000 --max-price 5000 --hits 10
run "ichiba genre --genre-id 0"                    ichiba genre --genre-id 0
run "ichiba ranking --genre-id 555086"             ichiba ranking --genre-id 555086

# Books
run "books search --keyword golang"                books search --keyword "golang"
run "books book --keyword 村上春樹"                books book --keyword "村上春樹"
run "books cd --keyword 米津玄師"                  books cd --keyword "米津玄師"
run "books dvd --keyword ジブリ"                   books dvd --keyword "ジブリ"
run "books magazine --keyword 週刊"                books magazine --keyword "週刊"
run "books game --keyword ポケモン"                books game --keyword "ポケモン"
run "books genre --genre-id 001"                   books genre --genre-id 001

# Travel
run "travel hotels (hokkaido/sapporo)"             travel hotels --large-area japan --middle-area hokkaido --small-area sapporo --detail-area A
run "travel hotel --hotel-no 901"                  travel hotel --hotel-no 901
run "travel vacant (tokyo)"                        travel vacant --large-area japan --middle-area tokyo --small-area tokyo --detail-area A --checkin-date "$CHECKIN" --checkout-date "$CHECKOUT" --adult-num 2
run "travel area"                                  travel area
run "travel ranking --genre onsen"                 travel ranking --genre onsen

# Misc
run "misc recipe"                                  misc recipe
run "misc kobo --keyword 漫画"                     misc kobo --keyword "漫画"
run "misc gora --keyword 箱根"                     misc gora --keyword "箱根"

echo ""
echo "Results: $PASS passed, $FAIL failed"

if [ "$FAIL" -gt 0 ]; then
    echo ""
    echo "Failed commands:"
    for cmd in "${FAILED_CMDS[@]}"; do
        echo "  - $cmd"
    done
    exit 1
fi
