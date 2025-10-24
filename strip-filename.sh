#!/usr/bin/env bash
set -euo pipefail

full_path="${1}"
stripped_path="$(basename "${full_path}")"
removed_extension="${stripped_path%.*}"

echo "${removed_extension}"
