#!/bin/sh

go run template.go
git add .
git cm -m "草"
git push origin master
