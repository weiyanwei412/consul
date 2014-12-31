##consul安全认证 
###ACL使用
####使用acl启动consul时，需要以下配置在/etc/consul.d/目录下:
    acl_datacenter:dc1    #####ACL数据中心，用于服务器
    acl_default_policy:deny    ####默认的ACL policy
    acl_master_token:master_token_test    #####ACL令牌
    acl_token:anonymous              ######默认的ACL令牌
    
####以acl方式启动在使用kv,agent,catalog,health等端点的时候需要带上token否则无法操作
###1.1:ACL
插入一个KV键值对,插入成功返回true

    PUT /v1/kv/<key>?token=<token>
**Example request**

    curl -X PUT -d 'aaaa' http://192.168.2.61:8500/v1/kv/key1?token=master_token_test
    
**Example response**

    true
    
查询键值对,如果有该键查询成功返回键值对,查询失败返回空

    GET /v1/kv/<key>?token=<token>

**Example request**

    curl http://192.168.2.61:8500/v1/kv/key1?token=master_token _test

**Example response**

    [
        {
            "CreateIndex":138,
            "ModifyIndex":142,
            "LockIndex":1,
            "key":"key1",
            "Flags":0,
            "value":"YWFhYQ"
        }
    ]

创建一个acl返回一个ID号/v1/acl/create

    PUT /v1/acl/create?token=<token>

**Example request**

    curl -X PUT -d '{"Name":"test_token","Type":"client","Rules":" key \"testread/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "}' http://192.168.2.61:8500/v1/acl/create?token=master_token_test
        {
            "Name": "test_token",
            "Type": "client",
            "Rules": " key \"testread/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "
        }
        
**Example response**

    "ID": "795cb1a7-bda9-dd3d-b266-a7c49de1100d"

查看存在的acl列表

    GET /v1/acl/list?token=<token>
    
**Example request**

    curl http://192.168.2.61:8500/v1/acl/list?token=master_token_test
    
**Example response**

    [
        {
            "CreateIndex": 70,
            "ModifyIndex": 70,
            "ID": "795cb1a7-bda9-dd3d-b266-a7c49de1100d",
            "Name": "test_token",
            "Type": "client",
            "Rules": " key \"testread/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" }                    "
        },
        {
            "CreateIndex": 63,
            "ModifyIndex": 63,
            "ID": "afb6c602-b3c3-86f0-dfce-fe6a98c3da7e",
            "Name": "",
            "Type": "client",
            "Rules": ""
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
            "CreateIndex": 88,
            "ModifyIndex": 88,
            "ID": "e28a0a0b-fd09-bf5f-65f3-a60d76319f2e",
            "Name": "test_token",
            "Type": "client",
            "Rules": " key \"key1/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "
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

插入一个KV键值对，并设置读权限,成功返回true

    PUT /v1/kv/testread/<key>?token=<token>

**Example request**

    curl -X PUT -d '300' http://192.168.2.61:8500/v1/kv/testread/key1?token=master_token_test

**Example response**

    true
    
根据token查询某个键的内容

    GET /v1/kv/<key>?token=<token>

**Example request**

    curl http://192.168.2.61:8500/v1/kv/testread/key1?token=master_token_test
    
**Example response**

    [
        {
            "CreateIndex": 134,
            "ModifyIndex": 134,
            "LockIndex": 0,
            "Key": "testread/key1",
            "Flags": 0,
            "Value": "MjAw"
        }
    ]
    
根据ID号来查询某个键的内容

    GET /v1/kv/<key>?token=<token>
    
**Example request**

    curl http://192.168.2.61:8500/v1/kv/testread/key1?token=795cb1a7-bda9-dd3d-b266-a7c49de1100d
    
**Example response**

    [
        {
            "CreateIndex": 134,
            "ModifyIndex": 134,
            "LockIndex": 0,
            "Key": "testread/key1",
            "Flags": 0,
            "Value": "MjAw"
        }
    ]
    
插入一个kv键值对,并设置可写权限,成功返回true

    PUT /v1/kv/<key>?token=<token>
    
**Example request**

    curl -X PUT http://192.168.2.61:8500/v1/kv/testwrite/key2?token=795cb1a7-bda9-dd3d-b266-a7c49de1100d -d 40
    
**Example response**

    true
    
查询插入的人可些键值对，成功返回键值对内容

    GET /v1/kv/<key>?token=<token>
    
**Example request**

    curl http://192.168.2.61:8500/v1/kv/testwrite/key2?token=795cb1a7-bda9-dd3d-b266-a7c49de1100
    
**Example response**

    [
        {
            "CreateIndex": 178,
            "ModifyIndex": 178,
            "LockIndex": 0,
            "Key": "testwrite/key2",
            "Flags": 0,
            "Value": "NDAw"
        }
    ]
    
查询注册的service.

    GET /v1/catalog/services?token=<token>
    
**Example request**

    curl http://192.168.2.61:8500/v1/catalog/services?token=795cb1a7-bda9-dd3d-b266-a7c49de1100d 
    
**Example response**

    {
        "consul":[]
    }
    
创建一个event,不带token也可以查询

    PUT /v1/event/fire/<name>?token=<token>
    
**Example request**

    curl -X PUT http://192.168.2.61:8500/v1/event/fire/testevent?token=795cb1a7-bda9-dd3d-b266-a7c49de1100d
    
**Example response**

    [
        {
            "ID": "aa2ddd58-551a-3dd7-3019-8f336a208431",
            "Name": "testevent",
            "Payload": null,
            "NodeFilter": "",
            "ServiceFilter": "",
            "TagFilter": "",
            "Version": 1,
            "LTime": 36
        }
    ]
    
查询event列表,不带token也可以查询

    GET /v1/event/list?token=<token>
    
**Example request**

    curl http://192.168.2.61:8500/v1/event/list?token=795cb1a7-bda9-dd3d-b266-a7c49de1100d
    
**Example response**

    [
        {
            "ID": "aa2ddd58-551a-3dd7-3019-8f336a208431",
            "Name": "testevent",
            "Payload": null,
            "NodeFilter": "",
            "ServiceFilter": "",
            "TagFilter": "",
            "Version": 1,
            "LTime": 36
        }
    ]
    
查询acl列表

    GET /v1/acl/list?token=<token>
    
**Example request**

    curl http://192.168.2.61:8500/v1/acl/list?token=master_token_test
    
**Example response**

    [
        {
            "CreateIndex": 70,
            "ModifyIndex": 70,
            "ID": "795cb1a7-bda9-dd3d-b266-a7c49de1100d",
            "Name": "test_token",
            "Type": "client",
            "Rules": " key \"testread/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "
        },
        {
            "CreateIndex": 63,
            "ModifyIndex": 63,
            "ID": "afb6c602-b3c3-86f0-dfce-fe6a98c3da7e",
            "Name": "",
            "Type": "client",
            "Rules": ""
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
            "CreateIndex": 88,
            "ModifyIndex": 88,
            "ID": "e28a0a0b-fd09-bf5f-65f3-a60d76319f2e",
            "Name": "test_token",
            "Type": "client",
            "Rules": " key \"key1/\" { policy = \"read\" } key \"testwrite/\" { policy = \"write\" } "
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

##session使用步骤

###1.注册一个service

    PUT /v1/agent/service/register

**Example request**

    curl -X PUT  http://192.168.2.61:8500/v1/agent/service/register
        {
            "Name": "mysql-1",
            "Tags": [
                "master"
            ],
            "Port": 3306,
            "Check": {
                "Script": "/root/mysqlcheck.sh",
                "Interval": "10s"
            }
        }


###2.注册一个KV，service/mysql-1/leader

    PUT /v1/kv/<key>

**Example request**

    curl -X PUT http://192.168.2.61:8500/v1/kv/service/mysql-1/leader
        { 
            "name": "Temperature Measurement",
            "source": "Thermometer",
            "sourceTime": "2014-01-01T16:56:54+11:00",
            "entityRef": "http://example.org/thermometers/99981",
            "context": {
                "value": "37.7",
                "units": "Celsius"
            }
        }

**Example response**

    true

###3.创建一个session

    PUT /v1/session/create

**Example request**

    curl -X PUT http://192.168.2.61:8500/v1/session/create
        {
            "LockDelay": "15s",
            "Name": "mysql",
            "Node": "consul-agent1",
            "Checks": [
                "serfHealth","service:mysql-1"
            ]
        }

**Example response**

    {
        "ID": "bd3deee0-b83b-a084-832d-a3d28c4bfc85"
    }

###4. 将session acquire KV,body内容为当前节点的信息

    PUT /v1/kv/<key>?acquire=<sessionID>

**Example request**

    curl -X PUT http://192.168.2.61:8500/v1/kv/service/mysql-1/leader?acquire=bd3deee0-b83b-a084-832d-a3d28c4bfc85
        {
            "Node":"consul-agent1",
            "Port":"3306",
            "Tag":"Information"
        }

**Example response**

    true

###5.查询KV是否将session添加上去

    GET /v1/kv/<key>

**Example request**

    curl http://192.168.2.61:8500/v1/kv/service/mysql-1/leader

**Example response**

    [
        {
            "CreateIndex": 32,
            "ModifyIndex": 176,
            "LockIndex": 2,
            "Key": "service/mysql-1/leader",
            "Flags": 0,
            "Value": "ewogIAkiTm9kZSI6ImNvbnN1bC1hZ2VudDEiLAogIAkiUG9ydCI6IjMzMDYiCn0=",
            "Session": "bd3deee0-b83b-a084-832d-a3d28c4bfc85"
        }
    ]

###6.注销服务mysql-1

    PUT /v1/agent/service/deregister/<service>

**Example request**

    curl -X PUT http://192.168.2.61:8500/v1/agent/service/deregister/mysql-1

###7.查看一下KV是否还有session ID

    GET /v1/kv/<key>

**Example request**

    curl http://192.168.2.71:8500/v1/kv/service/mysql-1/leader

**Example response**

    [
        {
            "CreateIndex": 32,
            "ModifyIndex": 185,
            "LockIndex": 2,
            "Key": "service/mysql-1/leader",
            "Flags": 0,
            "Value": "ewogIAkiTm9kZSI6ImNvbnN1bC1hZ2VudDEiLAogIAkiUG9ydCI6IjMzMDYiCn0="
        }
    ]

###8.查看session是否还在

    GET /v1/session/list

**Example request**

    curl http://192.168.2.61:8500/v1/session/list

**Example response**

    [ ]    