#!/usr/bin/env bash
kill -9 $(ps -el | grep sm_server | awk '{print $4" "$5}')
cd /home/hevlfreis/projects/Gloria/src
pwd
go build -o sm_server
nohup ./sm_server &
