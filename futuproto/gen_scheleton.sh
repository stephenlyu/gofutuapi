#!/usr/bin/env bash

ROOT=../libcpp/Include/Proto/

for f in $ROOT/*.proto;
do

  protoc -I $ROOT $f --go_out=import:../../../../.
done
