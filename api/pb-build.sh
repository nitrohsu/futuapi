#!/usr/bin/env bash

for file in $(ls)$(pwd)$(); do
  newFile=$(echo $file | cut -d . -f1)
  mkdir $newFile
  #  protoc --go_out=plugins=grpc:../api/$newFile $file
  mv $file $newFile/
done
