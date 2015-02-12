## 样例环境
测试环境一共5台机器分别是三台consul server两台consul agent，innosql安装在两台consul agent上。  
consul server 的ip地址分别为192.168.2.71, 192.168.2.72, 192.168.2.72  
consul agent 的ip地址为192.168.2.61, 192.168.2.62

### 1. consul
consul的二进制文件，版本号：Consul v0.5.0rc1-66-g954aec6.rc1 (954aec66231b79c161a4122b023fbcad13047f79)  
consul server和consul agent 使用相同的二进制文件, 使用不同的配置文件

### 2. 启动consul server集群
1.根据配置文件将consul启动为第一台consul server  
编写consul server配置文件，文件路径为/etc/consul.d/config.json   
文件内容如下：
	
```bash
	{
		"bootstrap": true,
		"server": true,
		"datacenter": "dc1",
		"data_dir": "/tmp/consul",
		"node_name": <hostname>,	# 根据hostname填写变量
		"addresses": {
			"http": <ip>,	# 根据主机ip地址填写变量
			 "rpc": <ip>	# 根据主机ip地址填写变量
		}
	}
```
将consul二进制文件添加到/usr/bin/  
启动consul server 命令如下：

```bash
consul agent -config-dir /etc/consul.d/
```

2.根据配置文件将consul启动为第二台consul server  
编写consul server配置文件，文件路径为/etc/consul.d/config.json  
文件内容如下:
	
```bash
	{
		"bootstrap": flase,
		"server": true,
		"datacenter": "dc1",
		"data_dir": "/tmp/consul",
		"node_name": <hostname>,	# 根据hostname填写变量
		"addresses": {
		  	"http": <ip>,	# 根据主机ip地址填写变量
		 	 "rpc": <ip>	# 根据主机ip地址填写变量
		},
		"start_join": [
			<ip> 	# 根据第一台consul server的ip地址填写变量
		 ]
	}
```
将consul二进制文件添加到/usr/bin/  
启动consul server 命令如下：

```bash
consul agent -config-dir /etc/consul.d/
```

3.根据配置文件将consul启动为第三台consul server  
编写consul server配置文件，文件路径为/etc/consul.d/config.json   
文件内容如下:
	
```bash
	{
		"bootstrap": flase,
		"server": true,
		"datacenter": "dc1",
		"data_dir": "/tmp/consul",
		"node_name": <hostname>,	# 根据hostname填写变量
		"addresses": {
		  	"http": <ip>,	# 根据主机ip地址填写变量
		 	 "rpc": <ip>	# 根据主机ip地址填写变量
		},
		"start_join": [
			<ip>,	 # 根据第一台consul server的ip地址填写变量
			<ip>	 # 根据第二台consul server的ip地址填写变量
		 ]
	}
```
将consul二进制文件添加到/usr/bin/  
启动consul server 命令如下：

```bash
consul agent -config-dir /etc/consul.d/
```

4.切换第一台consul server的启动模式  
将第一台consul server kill 掉  
执行以下命令

```bash
pkill consul
```

编写consul server配置文件，文件路径为/etc/consul.d/config.json  
文件内容如下:

```bash
	{
		"bootstrap": flase,
		"server": true,
		"datacenter": "dc1",
		"data_dir": "/tmp/consul",
		"node_name": <hostname>,	# 根据hostname填写变量
		"addresses": {
		  	"http": <ip>,	# 根据主机ip地址填写变量
		 	 "rpc": <ip>	# 根据主机ip地址填写变量
		},
		"start_join": [
			<ip>,	 # 根据第二台consul server的ip地址填写变量
			<ip>	 # 根据第三台consul server的ip地址填写变量
		 ]
	}
```
重新启动第一台consul server  
启动consul server 命令如下:

```bash
consul agent -config-dir /etc/consul.d/
```

### 3. 启动consul agent 并加入集群
1.根据配置文件将consul启动为第一台consul agent  
编写consul agent配置文件，文件路径为/etc/consul.d/config.json  
文件内容如下:

```bash
{
	"bootstrap": flase,
	"server": flase,
	"datacenter": "dc1",
	"data_dir": "/tmp/consul",
	"node_name": <hostname>,	# 根据hostname填写变量
	"addresses": {
	  	"http": <ip>,	# 根据主机ip地址填写变量
	 	 "rpc": <ip>	# 根据主机ip地址填写变量
	},
	"start_join": [
		<ip>, 	 # 根据第一台consul server的ip地址填写变量
		<ip>,	 # 根据第二台consul server的ip地址填写变量
		<ip>	 # 根据第三台consul server的ip地址填写变量
	 ]
}
```
将consul二进制文件添加到/usr/bin/  
启动consul agent 命令如下：

```bash
consul agent -config-dir /etc/consul.d/
```

2.根据配置文件将consul启动为第二台consul agent  
编写consul agent配置文件，文件路径为/etc/consul.d/config.json  
文件内容如下:

```bash
{
	"bootstrap": flase,
	"server": flase,
	"datacenter": "dc1",
	"data_dir": "/tmp/consul",
	"node_name": <hostname>,	# 根据hostname填写变量
	"addresses": {
	  	"http": <ip>,	# 根据主机ip地址填写变量
	 	"rpc": <ip> 	# 根据主机ip地址填写变量
	},
	"start_join": [
		<ip>, 	 # 根据第一台consul server的ip地址填写变量
		<ip>,	 # 根据第二台consul server的ip地址填写变量
		<ip>	 # 根据第三台consul server的ip地址填写变量
	 ]
}
```
将consul二进制文件添加到/usr/bin/  
启动consul agent 命令如下:

```bash
consul agent -config-dir /etc/consul.d/
```

### 4. 在所有consul agent 上启动innosql服务，配置主从关系

### 5. 在所有consul agent 上注册服务
1.准备服务状态检查脚本,检查脚本内容如下：

```bash
[root@consul-agent1 ~]# cat mysqlcheck.sh       
#!/bin/bash	
name=$1
dbname=`tr '[A-Z]' '[a-z]' <<<"$name"`
#basedir="/bsgchina/$dbname/MYBASE"
datadir="/var/lib/mysql"
bindir="/bin"

MYSQL=$bindir/mysql
MYSQL_USER=<username>	#填写innosql登陆用户名
MYSQL_PASSWORD=<password>	#填写innosql登陆密码
CHECK_TIME=3
MYSQL_OK=1
function check_mysql_health () {

	$MYSQL  -u $MYSQL_USER -p$MYSQL_PASSWORD -S $datadir/mysql.sock -e "show status;" > /dev/null 2>&1

	if [ $? -eq 0 ]; then 
		MYSQL_OK=0;
	else
		MYSQL_OK=1;
	fi

		return $MYSQL_OK;
}

while [ $CHECK_TIME -ne 0 ] 
do
	let "CHECK_TIME -= 1";
	check_mysql_health;
	if [ $MYSQL_OK = 0 ]; then
		CHECK_TIME=0;
		exit 0
	fi
	if [ $MYSQL_OK -eq 1 ] && [ $CHECK_TIME -eq 0 ]; then
		exit 2;
	fi
	sleep 1;
done
```

2.使用consul api接口向consul agent注册innosql服务  
将以下命令中的CONSUL_AGENT_IP替换为consul agent 所在主机的IP。

```bash
curl -X PUT -d '{"Name":"innosql","Tags":["master"],"Port":3306,"Check":{"Script":"/root/mysqlcheck.sh","Interval":"10s"}}' http://CONSUL_AGENT_IP:8500/v1/agent/service/register
```

3.使用consul api向consul service插入service/innosql/leader键(用于应用客户端查询innosql服务连接串)  
将以下命令中的CONSUL_SERVER_IP替换为任意一台consul server的IP, 如 192.168.2.71。

```bash
curl -X PUT http://CONSUL_SERVER_IP:8500/v1/kv/service/innosql/leader
```

### 6. 启动服务故障自动切换机制
1.修改innosql master上bootstrap/conf目录下面的app.conf, 路径为/root/bootstrap  
文件内容如下:

```bash
	appname = bootstrap
	runmode = dev
	loglevel = 7
	hostname = consul-agent1	
	ip = 192.168.2.61
	port = 3306
	username = root
	password = 111111
	datacenter = 
	token = 
	service_ip = 192.168.2.71
	service_port = 8500
	servicename = innosql
```

将上面配置文件中以下字段修改为相应的配置  
>
hostname为consul agent主机名(也就是innosql master的主机名)  
ip 为innosql 的ip地址  
port为innosql 的端口号  
username为innosql用户名  
password为innosql的密码  
service_ip为consul server集群中任意一台ip

2.在innosql **master**节点运行bootstrap二进制文件  

```
[root@consul-agent1 bootstrap]# ./bootstrap 
```
可以通过以下命令查看该键的value和session是否发生了改变。(将其中的CONSUL_SERVER_IP替换为任意一台consul server的IP)

```
curl -X GET http://CONSUL_SERVER_IP:8500/v1/kv/service/innosql/leader
```

3.修改innosql master 和slave 机器上的 mha-handlers/conf目录下面的app.conf, 路径为/root/mha-handlers  
文件内容如下:

```bash
	appname = mha-handlers
	runmode = dev
	loglevel = 7
	hostname = consul-agent1	
	ip = 192.168.2.61
	port = 3306
	username = root
	password = 111111
	datacenter = 
	token = 
	service_ip = 192.168.2.71
	service_port = 8500
	servicename = innosql
```

将上面配置文件中以下字段修改为相应的配置  
>
hostname为consul agent主机名(也就是innosql master的主机名)  
ip 为innosql 的ip地址  
port为innosql 的端口号  
username为innosql用户名  
password为innosql的密码  
service_ip为consul server集群中任意一台ip  

4.分别在consul agent上执行consul watch命令：

```bash
consul watch -http-addr=192.168.2.61:8500 -type key -key service/innosql/leader /root/mha-handlers/mha-handlers 
```
>可以将该命令放在后台执行。日志可在/root/mha-handlers/logs/mha-handlers.log中查看。

### 7. 至此环境搭建完成

### 8. 验证环境是否搭建成功

1. 手动停止master上的mysql实例，然后可以在watch命令的执行过程中或mha-handlers.log查看slave争填service/innosql/leader这个键的值的过程。
2. 