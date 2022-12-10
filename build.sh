#!/bin/bash
go env GOOS=linux GOARCH=arm GOARM=5 
go build -v
echo "Done. Press any key to exit."
read -r
