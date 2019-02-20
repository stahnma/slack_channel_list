#!/bin/bash


cd /home/stahnma/go/src/github.com/stahnma/slack_channel_list
. environment
num=`/usr/local/bin/slack_channel_list | wc -l | awk '{print $1}'`
echo "`date +%F`,$num" >>  /home/stahnma/go/src/github.com/stahnma/slack_channel_list/channel_list.log
