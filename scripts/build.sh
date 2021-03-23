#!/bin/bash
# vim: ai:ts=8:sw=8:noet
# Build the image
# Intended to be run from CI or local
set -eufo pipefail
export SHELLOPTS	# propagate set to children by default
IFS=$'\t\n'

# Check required commands are in place
command -v docker-compose >/dev/null 2>&1 || { echo 'Please install docker or use image that has it'; exit 1; }

docker-compose build