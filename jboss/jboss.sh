#/bin/bash
CS_addr=""
CS_port="8500"
SVC_name=$1
SVC_service_port=$2
SVC_check=""
CA_addr="127.0.0.1"
CA_port="8500"
curl -X PUT -d "{"Name":"${SVC_name}","Port":${SVC_service_port},"Check":{"Script":"${SVC_check}","Interval":"10s"}}" http://${CA_addr}:${CA_port}/v1/agent/service/register
pass=`curl -X GET http://${CA_addr}:${CA_port}/v1/health/service/${SVC_name} | grep -o -w "passing" | wc -l` > /dev/null 2>&1

if [ $pass -ne 2 ]; then
	echo "注册服务失败!"
	exit
fi

while read line
do
	
	key=`echo ${line} | cut -d "=" -f1`
	values=`echo ${line} | cut -d "=" -f2`
	curl -X PUT -d '${values}' http://${CS_addr}:${CS_port}/v1/kv/service/${SVC_name}/${key} > /dev/null 2>&1
done <config.cfg
