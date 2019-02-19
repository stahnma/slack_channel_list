#!/bin/bash

. environment
num=`slack_channel_list | wc -l | awk '{print $1}'`
echo "`date +%F`,$num" >> channel_list.log
