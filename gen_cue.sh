#!/bin/bash

mv json.go json_go_exclude 
cue get go github.com/dgruber/drmaa2interface -v -i -p drmaa2cue --local
mv json_go_exclude json.go
