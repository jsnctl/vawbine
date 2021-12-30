#!/bin/zsh

go test ./...
go build
./vawbine
ffplay -f f32le -ar 44100 -showmode 1 sound.bin