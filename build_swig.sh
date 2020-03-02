#!/usr/bin/env bash

swig -go -cgo -intgosize 64 -c++ gofutuapi.i
