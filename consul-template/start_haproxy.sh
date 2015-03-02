#!/bin/bash

ServiceName=$1

CTDir=/etc/consul-template/

CTTemplateFile=/tmp/haproxy.ctmpl

HaproxyTemplateFile=/etc/haproxy/haproxy.cfg

ConsulServerIP=$2

if [ ! -d "${CTDir}" ]; then
	mkdir -p "${CTDir}"
	
	if [ $? -ne 0 ]; then
		echo "Create a consul-template template directory failed"
		exit 102
	fi
fi
cp "${CTTemplateFile}" "${CTDir}"
if [ $? -ne 0 ]; then
	echo "Copy the template file to the consul-template directory failed"
	exit 100
fi
sed -i "s/N2PC4F/${ServiceName}/g" ${CTDir}/haproxy.ctmpl
if [ $? -ne 0 ]; then
	echo "Replace the service name failed"
	exit 101
fi
gocode/bin/consul-template -consul ${ConsulServerIP} -template "${CTDir}/haproxy.ctmpl:${HaproxyTemplateFile}:service haproxy restart" &
echo "haproxy successful start"
exit 200
