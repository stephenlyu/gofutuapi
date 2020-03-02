#!/usr/bin/env bash

for f in ../libcpp/Include/Proto/*.proto;
do
  sed -i '/go_package/d' $f
  file_name=`basename $f`
  main_name=${file_name%%\.proto}
  sed -i "/java_package/aoption go_package = \"github.com/stephenlyu/gofutuapi/futuproto/${main_name}\";\n" $f
done