#!/bin/bash
ps aux | grep api| grep -v grep | awk '{print $2}'| xargs kill

nohup ./api  > api.log  2>&1  &

