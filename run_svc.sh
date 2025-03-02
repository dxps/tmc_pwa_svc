#!/bin/sh

## Using `entr` (from https://eradman.com/entrproject/)
find . -name "*.go" | entr -r make dev
