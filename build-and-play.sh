#!/bin/zsh

go build
./gotechre
ffplay -f f32le -ar 44100 -showmode 1 sound.bin