#!/usr/bin/env bash

set -euo pipefail
IFS=$'\n\t'

(trap 'kill 0' SIGINT; \
bash -c './run_css.sh' & \
bash -c './run_app.sh'
)
