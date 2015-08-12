#!/bin/bash
set -x
set -e

PORT="0"

if [ -e /home/ubuntu/webapp/citycloud.cf-deploy-ui ]; then
	while [ $PORT -ne 1 ]
	do
		sleep 5s
		PORT=`netstat -lnt|grep 3306|wc -l` ## word count 3306 for test mysql is running
	done
	
	nohup /home/ubuntu/webapp/citycloud.cf-deploy-ui &
fi