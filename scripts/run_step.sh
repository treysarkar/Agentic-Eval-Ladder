#!/usr/bin/env bash
# Usage: scripts/run_step.sh 4   -> runs TestStep01 .. TestStep04
set -euo pipefail
n=${1:?usage: run_step.sh N}
pattern="^TestStep($(seq -f '%02g' -s '|' 1 "$n"))_"
go test ./tests/ -count=1 -run "$pattern" -v