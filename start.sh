#!/usr/bin/env bash
kill -9 $(ps -el | grep api_server | awk '{print $4" "$5}')
cd /home/hevlfreis/projects/Gloria/src
pwd
go build -o api_server
cd ..
nohup ./src/api_server -port=8080 &
