#!/bin/bash

curl http://192.168.2.71:8500/v1/session/list?pretty=1 >a.txt

cat a.txt | grep -w "ID" | awk '{print $2}' | sed 's/\"//g'|sed 's/\,//g'>b.txt

while read line
do
	curl -X PUT http://192.168.2.71:8500/v1/session/destroy/${line}
	echo ${line}

done < b.txt

