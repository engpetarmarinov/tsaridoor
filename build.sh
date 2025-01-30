#!/bin/bash
env GOOS=linux GOARCH=arm64 go build -v -o tsaridoor-hotrelease
echo "Done. Press any key to exit."
read -r
