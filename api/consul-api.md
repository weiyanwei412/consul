page_title: Remote API v0.4.1
page_description: API Documentation for Consul

# Consul Remote API v0.4.1

# 1. Endpoints

##	1.1	KV	

###1.2

获得所有机器的KEY/VALUE DATA

	GET /v1/kv/<key>?recurse

**Example request**:

	curl http://192.168.114.67:8500/v1/kv/web?recurse

**Example response**:

	[
    	{
       		"CreateIndex": 195, 
        	"ModifyIndex": 196, 
        	"LockIndex": 0, 
        	"Key": "web/key1", 
        	"Flags": 0, 
        	"Value": "dGVzdA=="
    	}, 
    	{
        	"CreateIndex": 197, 
        	"ModifyIndex": 197, 
        	"LockIndex": 0, 
        	"Key": "web/key2", 
       		"Flags": 0, 
        	"Value": "dGVzdA=="
    	}
	]

###1.3
插入一个KEY/VALUE DATA

	PUT /v1/kv/<key>

**Example request**:

	curl -X PUT -d 'test' http://192.168.114.67:8500/v1/kv/web/key1

###1.4
查询web/key1的信息

	GET /v1/kv/<key>

**Example request**:

	curl http://192.168.114.67:8500/v1/kv/web/key1

**Example response**:
	
	[
    	{
       		"CreateIndex": 195, 
        	"ModifyIndex": 211, 
        	"LockIndex": 0, 
        	"Key": "web/key1", 
        	"Flags": 0, 
        	"Value": "dGVzdA=="
    	}
	]

###1.5
删除web/sub及所有子串

	DELETE /v1/kv/web?recurse

**Example request**:
	curl -X DELETE http://localhost:8500/v1/kv/web?recurse

###1.6
修改web/key的值

	PUT /v1/kv/<key>?cas=231

**Example request**:

	curl -X PUT -d 'newval' http://192.168.114.67:8500/v1/kv/web/key1?cas=231

###1.7
查询所有的键，以list的方式显示
	GET /v1/kv/?keys

**Example request**:

	curl http://192.168.114.67:8500/v1/kv/?keys

**Example response**:

	[
    	"web/key1", 
    	"web/key2"
	]

###1.8	Parameters:
-   **?dc=**	-	提供数据中心.
-   **?recurse** - 将返回给定前缀的所有信息
-   **?keys** - 以list的方式显示获得所有键
-   **?separator=** - 可用于仅列出给定的分隔符
-   **?raw** - 使用递归的方式得到指定键的值，没有任何编码
-   **?cas** - 此标志可用于检查和设置操作
-   **?acquire=<session>** - 放入一个锁操作
-	**?release=<session>** - 释放一个锁操作


Status Codes:

-   **true** – no error
-   **false** - error，则没有进行更新



##2.1	Agent

###2.2

返回本地所有的检查

	GET /v1/agent/checks

**Example request**:

	http://192.168.114.64:8500/v1/agent/checks

**Example response**:

	{
    	"mysql-1": {
        	"Node": "node5",
        	"CheckID": "mysql-1",
        	"Name": "mysql-1",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "",
        	"ServiceID": "",
        	"ServiceName": ""
    	}
	}

###2.3

返回本地所有服务注册

	GET /v1/agent/services

**Example request**:

	http://192.168.114.64:8500/v1/agent/services

**Example response**:
	
	{
    	"mysql-1": {
        	"ID": "mysql-1",
        	"Service": "mysql-1",
        	"Tags": [
            	"master"
        	],
        	"Port": 3306
    	}
	}

###2.4

返回集群中的所有成员

	GET /v1/agent/members

**Example request**:

	http://192.168.114.64:8500/v1/agent/members

**Example response**:

	[
    	{
        	"Name": "node5",
        	"Addr": "192.168.114.64",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"role": "node",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	},
    	{
        	"Name": "node7",
        	"Addr": "192.168.114.66",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"port": "8300",
            	"role": "consul",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	},
    	{
        	"Name": "node6",
        	"Addr": "192.168.114.65",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"port": "8300",
            	"role": "consul",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	},
    	{
        	"Name": "node8",
        	"Addr": "192.168.114.67",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"port": "8300",
            	"role": "consul",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	}
	]

###2.5

返回本地代理和成员信息的配置	

	GET /v1/agent/self

**Example request**:

	http://192.168.114.64:8500/v1/agent/serf

**Example response**:


	[
    	{
        	"Name": "node5",
        	"Addr": "192.168.114.64",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"role": "node",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	},
    	{
        	"Name": "node7",
        	"Addr": "192.168.114.66",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"port": "8300",
            	"role": "consul",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	},
    	{
        	"Name": "node6",
        	"Addr": "192.168.114.65",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"port": "8300",
            	"role": "consul",
            	"vsn": "2",
            	"vsn_max": "2",
           		"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	},
    	{
        	"Name": "node8",
        	"Addr": "192.168.114.67",
        	"Port": 8301,
        	"Tags": {
            	"build": "0.4.1:",
            	"dc": "dc1",
            	"port": "8300",
            	"role": "consul",
            	"vsn": "2",
            	"vsn_max": "2",
            	"vsn_min": "1"
        	},
        	"Status": 1,
        	"ProtocolMin": 1,
        	"ProtocolMax": 2,
        	"ProtocolCur": 2,
        	"DelegateMin": 2,
        	"DelegateMax": 4,
        	"DelegateCur": 4
    	}
	]

###2.6

将本机加入集群
	GET /v1/agent/join/<address>

**Example request**:

	http://192.168.114.65:8500/v1/agent/join/192.168.114.64

###2.7

使节点为离开状态

	GET /v1/agent/force-leave/<node>

**Example request**:

	http://192.168.114.64:8500/v1/agent/force-leave/node5

###2.8

动态注册检查配置

	PUT /v1/agent/check/register

**Example request**:
	
	http://192.168.114.64:8500/v1/agent/check/register
	{
  		"ID": "mysql-1",
  		"Name": "mysql-1",
  		"Script": "/root/mysqlcheck.sh orz1bsg",
  		"Interval": "10s"
	}

###2.9

动态注销检查配置

	POST /v1/agent/check/deregister/<checkID>

**Example request**:

	http://192.168.114.64:8500/v1/agent/check/deregister/mysql-1

###2.10

将检查状态设置成passing   注册检查时需要设置TTL

	GET /v1/agent/check/pass/<checkID>

**Example request**:

	http://192.168.114.64:8500/v1/agent/check/pass/mysql-1

###2.11

将服务状态设置成warning

	GET /v1/agent/check/warn/<checkID>

**Example request**:

	http://192.168.114.64:8500/v1/agent/check/warn/mysql-1

###2.12

将服务状态设置成critical

	GET /v1/agent/check/fail/<checkID>

**Example request**:

	http://192.168.114.64:8500/v1/agent/check/fail/mysql-1

###2.13

动态注册服务

	PUT /v1/agent/service/register

**Example request**:

	http://192.168.114.64:8500/v1/agent/service/register

	{
  		"Name": "mysql-1",
  		"Tags": [
    		"master"
  		],
  		"Port": 3306,
  		"Check": {
    		"Script": "/root/mysqlcheck.sh orz1bsg",
    		"Interval": "10s"
  		}
	}

###2.14

动态注销服务
	GET v1/agent/service/deregister/<serviceID>

**Example request**:

	http://192.168.114.64:8500/v1/agent/service/deregister/mysql-1


##3.1	Catalog

###3.2

动态注册服务和检查

	PUT /v1/catalog/register

**Example request**:

	http://192.168.114.64:8500/v1/catalog/register

	{
  		"Datacenter": "dc1",
  		"Node": "node5",
  		"Address": "192.168.114.64",
  		"Service": {
    		"ID": "mysql-1",
    		"Service": "mysql-1",
    		"Tags": [
      		"master",
      		"v1"
    		]
 		},
  		"port": 3306,
  		"Check": {
    		"Node": "node5",
    		"CheckID": "service:mysql-1",
    		"Name": "mysql-1 health check",
    		"Notes": "script based health check",
    		"Script": "/root/mysqlcheck.sh orz1bsg",
    		"Interval": "3s",
    		"ServiceID": "mysql-1"
  		}
	}

###3.3

动态注销服务和检查

	PUT /v1/catalog/deregister

**Example request**:
	
	http://192.168.114.64:8500/v1/catalog/deregister	

	{
  		"Datacenter": "dc1",
  		"Node": "node5",
  		"ServiceID": "mysql-1"
	}

###3.4

返回所有数据中心

	GET /v1/catalog/datacenters


**Example request**:

	http://192.168.114.64:8500/v1/catalog/datacenters

**Example response**:

	[
    	"dc1"
	]

###3.5

返回集群中的节点

	/v1/catalog/nodes


**Example request**:

	http://192.168.114.64:8500/v1/catalog/nodes

**Example response**:

	[
    	{
        	"Node": "node5", 
        	"Address": "192.168.114.64"
    	}, 
    	{
        	"Node": "node6", 
        	"Address": "192.168.114.65"
    	}, 
    	{
        	"Node": "node7", 
        	"Address": "192.168.114.66"
    	}, 
    	{
        	"Node": "node8", 
        	"Address": "192.168.114.67"
    	}
	]

###3.6

查询注册了的服务

	GET /v1/catalog/services

**Example request**:

	http://192.168.114.64:8500/v1/catalog/services

**Example response**:

	{
    	"consul": [],
    	"mysql-1": [
        	"master",
        	"v1"
    	]
	}

###3.7

返回指定服务的信息

	GET /v1/catalog/service/<service>

**Example request**:

	http://192.168.114.64:8500/v1/catalog/service/mysql-1

**Example response**:

	[
    	{
        	"Node": "node5",
        	"Address": "192.168.114.64",
        	"ServiceID": "mysql-1",
        	"ServiceName": "mysql-1",
        	"ServiceTags": [
            	"master",
            	"v1"
        	],
        	"ServicePort": 0
    	}
	]

###3.8

返回指定节点的具体信息

	GET /v1/catalog/node/<node>


**Example request**:

	http://192.168.114.64:8500/v1/catalog/node/node5

**Example response**:

	{
    	"Node": {
        	"Node": "node5",
        	"Address": "192.168.114.64"
    	},
    	"Services": {
        	"mysql-1": {
            	"ID": "mysql-1",
            	"Service": "mysql-1",
            	"Tags": [
                	"master",
                	"v1"
            	],
            	"Port": 0
        	}
    	}
	}

##4.1	Health

###4.2

查询节点的健康信息

	GET /v1/health/node/<node>


**Example request**:

	http://192.168.114.64:8500/v1/health/node/node5

**Example response**:

	[
    	{
        	"Node": "node5",
        	"CheckID": "mysql-1",
        	"Name": "mysql-1",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "",
        	"ServiceID": "",
        	"ServiceName": ""
    	},
    	{
        	"Node": "node5",
        	"CheckID": "serfHealth",
        	"Name": "Serf Health Status",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "Agent alive and reachable",
        	"ServiceID": "",
        	"ServiceName": ""
    	}
	]

###4.3

返回指定服务的检查信息

	GET /v1/health/checks/<service>


**Example request**:

	http://192.168.114.64:8500/v1/health/checks/mysql-1

**Example response**:

	[
    	{
        	"Node": "node5",
        	"CheckID": "service:mysql-1",
        	"Name": "mysql-1 health check",
        	"Status": "critical",
        	"Notes": "script based health check",
        	"Output": "",
        	"ServiceID": "mysql-1",
        	"ServiceName": "mysql-1"
    	}
	]

###4.4

返回服务的健康信息

	GET /v1/health/service/<service>


**Example request**:

	http://192.168.114.64:8500/v1/health/service/mysql-1

**Example response**:

	[
    	{
        	"Node": {
            	"Node": "node5",
            	"Address": "192.168.114.64"
        	},
        	"Service": {
            	"ID": "mysql-1",
            	"Service": "mysql-1",
            	"Tags": [
                	"master",
                	"v1"
            	],
            	"Port": 0
        	},
        	"Checks": [
            	{
                	"Node": "node5",
                	"CheckID": "service:mysql-1",
                	"Name": "mysql-1 health check",
                	"Status": "critical",
                	"Notes": "script based health check",
                	"Output": "",
                	"ServiceID": "mysql-1",
                	"ServiceName": "mysql-1"
            	},
            	{
                	"Node": "node5",
                	"CheckID": "serfHealth",
                	"Name": "Serf Health Status",
                	"Status": "passing",
                	"Notes": "",
                	"Output": "Agent alive and reachable",
                	"ServiceID": "",
                	"ServiceName": ""
            	},
            	{
               		"Node": "node5",
                	"CheckID": "mysql-1",
                	"Name": "mysql-1",
                	"Status": "passing",
                	"Notes": "",
                	"Output": "",
                	"ServiceID": "",
                	"ServiceName": ""
            	}
        	]
    	}
	]

###4.5

返回给定状态的检查

	GET /v1/health/state/<state>


**Example request**:

	http://192.168.114.64:8500/v1/health/state/passing

**Example response**:

	[
    	{
        	"Node": "node6",
        	"CheckID": "serfHealth",
        	"Name": "Serf Health Status",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "Agent alive and reachable",
        	"ServiceID": "",
        	"ServiceName": ""
    	},
    	{
        	"Node": "node7",
        	"CheckID": "serfHealth",
        	"Name": "Serf Health Status",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "Agent alive and reachable",
        	"ServiceID": "",
        	"ServiceName": ""
    	},
    	{
        	"Node": "node8",
        	"CheckID": "serfHealth",
        	"Name": "Serf Health Status",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "Agent alive and reachable",
        	"ServiceID": "",
       		"ServiceName": ""
    	},
    	{
        	"Node": "node5",
        	"CheckID": "serfHealth",
        	"Name": "Serf Health Status",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "Agent alive and reachable",
        	"ServiceID": "",
        	"ServiceName": ""
    	},
    	{
        	"Node": "node5",
        	"CheckID": "mysql-1",
        	"Name": "mysql-1",
        	"Status": "passing",
        	"Notes": "",
        	"Output": "",
        	"ServiceID": "",
       		"ServiceName": ""
    	}
	]

##5.1	Session

###5.2

创建一个新的会话

	PUT /v1/session/create


**Example request**:

	http://192.168.114.64:8500/v1/session/create

	{
  		"LockDelay": "15s",
  		"Name": "mysql",
  		"Node": "node5",
  		"Checks": ["mysql-1","mysql-2"]
	}

**Example response**:

	{
    	"ID": "d7d177c6-5c36-59f9-410d-d9e78a53256d"
	}

###5.3

注销一个给定的会话

	PUT /v1/session/destroy/<session>


**Example request**:

	http://192.168.114.64:8500/v1/session/destroy/4d2c59bf-c381-38ae-5b01-d843efe85307

##5.4

查询某个session信息

	GET /v1/session/info/<session>


**Example request**:

	http://192.168.114.64:8500/v1/session/info/4d2c59bf-c381-38ae-5b01-d843efe85307

**Example response**:

	[
    	{
        	"CreateIndex": 781,
        	"ID": "ae186218-00e1-8f0b-1fa4-b56175d3f3c8",
        	"Name": "mysql",
        	"Node": "node5",
        	"Checks": [
            	"mysql-1"
        	],
        	"LockDelay": 15000000000
    	}
	]	

##5.5

返回某个节点的session

	GET /v1/session/node/<node>


**Example request**:

	http://192.168.114.64:8500/v1/session/node/node5


**Example response**:

	[
    	{
       		"CreateIndex": 781,
        	"ID": "ae186218-00e1-8f0b-1fa4-b56175d3f3c8",
        	"Name": "mysql",
        	"Node": "node5",
        	"Checks": [
            	"mysql-1"
        	],
        	"LockDelay": 15000000000
    	}
	]

##5.6

得到所有的session
	GET /v1/session/list

**Example request**:

	http://192.168.114.64:8500/v1/session/list

**Example response**:

	[
    	{
        	"CreateIndex": 781,
        	"ID": "ae186218-00e1-8f0b-1fa4-b56175d3f3c8",
        	"Name": "mysql",
        	"Node": "node5",
        	"Checks": [
            	"mysql-1"
        	],
        	"LockDelay": 15000000000
    	}
	]

##6.1	ACL

###6.2

创建一个新令牌

	PUT /v1/acl/create


**Example request**:

	http://192.168.114.64:8500/v1/acl/create

	{
  		"Name": "test_token",
  		"Type": "client",
  		"Rules": " key \"testread/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "
	}


**Example response**:

	{
    	"ID": "907f9f17-4814-77a3-f4d8-41ed26c33ca6"
	}

###6.3

更新一个列表

	PUT	/v1/acl/update


**Example request**:

	http://192.168.114.64:8500/v1/acl/update

	{
  		"ID": "567b1eb3-2e7d-ede5-9ad1-29619321fd76",
  		"Name": "test_token1",
  		"Type": "client",
  		"Rules": " key \"test1/\" { policy = \"write\" } "
	}

**Example response**:

	{
    	"ID": "567b1eb3-2e7d-ede5-9ad1-29619321fd76"
	}

###6.4

销毁一个列表
	
	PUT /v1/acl/destroy/<id>


**Example request**:
		
	http://192.168.114.64:8500/v1/acl/destroy/43629a35-8c7c-b490-001a-2eeb13f85e75


**Example response**:
	
	true

###6.5

列出当前所有的ACL

	GET /v1/acl/list

**Example request**:

	http://192.168.114.64:8500/v1/acl/list

**Example response**:

	[
    	{
        	"CreateIndex": 162,
        	"ModifyIndex": 448,
        	"ID": "567b1eb3-2e7d-ede5-9ad1-29619321fd76",
        	"Name": "test_token1",
        	"Type": "client",
        	"Rules": " key \"test1/\" { policy = \"write\" } "
    	},
    	{
        	"CreateIndex": 172,
        	"ModifyIndex": 172,
        	"ID": "58755da3-3b6b-b624-7b17-1690b5223f0a",
        	"Name": "test_token1",
        	"Type": "client",
        	"Rules": " key \"test1/\" { policy = \"write\" } "
    	},
    	{
        	"CreateIndex": 450,
        	"ModifyIndex": 450,
        	"ID": "907f9f17-4814-77a3-f4d8-41ed26c33ca6",
        	"Name": "test_token",
        	"Type": "client",
        	"Rules": " key \"testread/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "
    	},
    	{
        	"CreateIndex": 2,
        	"ModifyIndex": 2,
        	"ID": "anonymous",
        	"Name": "Anonymous Token",
        	"Type": "client",
        	"Rules": ""
    	},
    	{
        	"CreateIndex": 3,
        	"ModifyIndex": 3,
        	"ID": "master_token_test",
        	"Name": "Master Token",
        	"Type": "management",
        	"Rules": ""
    	}
	]

###6.6

查看某条ACL信息

	GET /v1/acl/info/<id>

**Example request**:

	http://192.168.114.64:8500/v1/acl/info/567b1eb3-2e7d-ede5-9ad1-29619321fd76

**Example response**:

	[
    	{
        	"CreateIndex": 162,
        	"ModifyIndex": 448,
        	"ID": "567b1eb3-2e7d-ede5-9ad1-29619321fd76",
        	"Name": "test_token1",
        	"Type": "client",
        	"Rules": " key \"test1/\" { policy = \"write\" } "
    	}
	]

###6.7

克隆ACL信息

	PUT /v1/acl/clone/<id>

**Example request**:

	http://192.168.114.64:8500/v1/acl/clone/567b1eb3-2e7d-ede5-9ad1-29619321fd76

**Example response**:

	{
    	"ID": "0e23123a-c90e-229f-8538-ea58273c4a36"
	}

##7.1	event

###7.2

创建一个新的用户事件

	PUT	/v1/event/fire/<name>

**Example request**:
	
	http://192.168.114.64:8500/v1/event/fire/test_token1

**Example response**:

	{
    	"ID": "c36e2d97-94ab-e3e3-8187-f74c3724f94e",
    	"Name": "test_token1",
    	"Payload": "ewogICJJRCI6ICI1NjdiMWViMy0yZTdkLWVkZTUtOWFkMS0yOSIsCiAgIk5hbWUiOiAidGVzdF90b2tlbjEiLAogICJQYXlsb2FkIjogbnVsbCwKICAiTm9kZUZpbHRlciI6ICIiLAogICJTZXJ2aWNlRmlsdGVyIjogIiIsCiAgIlRhZ0ZpbHRlciI6ICIiLAogICJWZXJzaW9uIjogMSwKICAiTFRpbWUiOiAwCn0=",
    	"NodeFilter": "",
    	"ServiceFilter": "",
    	"TagFilter": "",
    	"Version": 1,
    	"LTime": 0
	}

###7.3

得到最近发生的事件列表

	GET	/v1/event/list

**Example request**:

	http://192.168.114.64:8500/v1/event/list

**Example response**:

	[
    	{
        	"ID": "4f806790-f0c2-de9c-10a2-0bd222e6f822",
        	"Name": "test_token1",
        	"Payload": "ewogICJJRCI6ICI1NjdiMWViMy0yZTdkLWVkZTUtOWFkMS0yOTYxOTMyMWZkNzYiLAogICJOYW1lIjogInRlc3RfdG9rZW4xIiwKICAiUGF5bG9hZCI6IG51bGwsCiAgIk5vZGVGaWx0ZXIiOiAiIiwKICAiU2VydmljZUZpbHRlciI6ICIiLAogICJUYWdGaWx0ZXIiOiAiIiwKICAiVmVyc2lvbiI6IDEsCiAgIkxUaW1lIjogMAp9",
        	"NodeFilter": "",
        	"ServiceFilter": "",
        	"TagFilter": "",
        	"Version": 1,
        	"LTime": 3
    	},
    	{
        	"ID": "b5457d55-6b7e-d351-0851-43d2350b3668",
        	"Name": "test_token1",
        	"Payload": "ewogICJJRCI6ICI1NjdiMWViMy0yZTdkLWVkZTUtOWFkMS0yOTYxOTMyMWZkNzYiLAogICJOYW1lIjogInRlc3RfdG9rZW4xIiwKICAiUGF5bG9hZCI6IG51bGwsCiAgIk5vZGVGaWx0ZXIiOiAiIiwKICAiU2VydmljZUZpbHRlciI6ICIiLAogICJUYWdGaWx0ZXIiOiAiIiwKICAiVmVyc2lvbiI6IDEsCiAgIkxUaW1lIjogMAp9",
        	"NodeFilter": "",
        	"ServiceFilter": "",
        	"TagFilter": "",
        	"Version": 1,
        	"LTime": 4
    	},
    	{
        	"ID": "525752e1-50df-7b55-284a-ea6ce4b50bd6",
        	"Name": "test_token1",
        	"Payload": "ewogICJJRCI6ICI1NjdiMWViMy0yZTdkLWVkZTUtOWFkMS0yOSIsCiAgIk5hbWUiOiAidGVzdF90b2tlbjEiLAogICJQYXlsb2FkIjogbnVsbCwKICAiTm9kZUZpbHRlciI6ICIiLAogICJTZXJ2aWNlRmlsdGVyIjogIiIsCiAgIlRhZ0ZpbHRlciI6ICIiLAogICJWZXJzaW9uIjogMSwKICAiTFRpbWUiOiAwCn0=",
        	"NodeFilter": "",
        	"ServiceFilter": "",
        	"TagFilter": "",
        	"Version": 1,
        	"LTime": 5
    	},
    	{
        	"ID": "c5eaefbd-c1a9-c079-4be2-a7c25d8b164b",
        	"Name": "test_token1",
        	"Payload": "ewogICJJRCI6ICI1NjdiMWViMy0yZTdkLWVkZTUtOWFkMS0yOSIsCiAgIk5hbWUiOiAidGVzdF90b2tlbjEiLAogICJQYXlsb2FkIjogbnVsbCwKICAiTm9kZUZpbHRlciI6ICIiLAogICJTZXJ2aWNlRmlsdGVyIjogIiIsCiAgIlRhZ0ZpbHRlciI6ICIiLAogICJWZXJzaW9uIjogMSwKICAiTFRpbWUiOiAwCn0=",
        	"NodeFilter": "",
        	"ServiceFilter": "",
        	"TagFilter": "",
        	"Version": 1,
        	"LTime": 6
    	},
    	{
        	"ID": "c36e2d97-94ab-e3e3-8187-f74c3724f94e",
        	"Name": "test_token1",
        	"Payload": "ewogICJJRCI6ICI1NjdiMWViMy0yZTdkLWVkZTUtOWFkMS0yOSIsCiAgIk5hbWUiOiAidGVzdF90b2tlbjEiLAogICJQYXlsb2FkIjogbnVsbCwKICAiTm9kZUZpbHRlciI6ICIiLAogICJTZXJ2aWNlRmlsdGVyIjogIiIsCiAgIlRhZ0ZpbHRlciI6ICIiLAogICJWZXJzaW9uIjogMSwKICAiTFRpbWUiOiAwCn0=",
        	"NodeFilter": "",
        	"ServiceFilter": "",
        	"TagFilter": "",
        	"Version": 1,
        	"LTime": 7
    	}
	]
	

##	8.1	Status

###8.2

返回当前的leader节点

	GET	/v1/status/leader

**Example request**:

	http://192.168.114.64:8500/v1/status/leader

**Example response**:

	"192.168.114.67:8300"

###8.3

返回同级的地址和端口号

	GET /v1/status/peers


**Example request**:

	http://192.168.114.64:8500/v1/status/peers

**Example response**:

	[
    	"192.168.114.67:8300",
    	"192.168.114.65:8300",
    	"192.168.114.66:8300"
	]