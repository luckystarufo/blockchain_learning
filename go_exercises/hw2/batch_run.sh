#!/bin/bash

for f in *.go; do
    echo "=== $f ==="
    go run "$f"
done