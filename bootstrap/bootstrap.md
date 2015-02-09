##　测试环境一共5台机器分别是三台consul server两台consul agent，innosql安装在两台consul agent上
	
	我的三台consul server机器的ip地址分别为192.168.2.71,192.168.2.72,192.168.2.72
	两台consul agent 机器的ip地址为192.168.2.61,192.168.2.62

### 1.安装consul

#### 用go get安装,前提是先安装go

	go get -v github.com/hashicorp/consul
	
### 2.编辑consul配置文件

	mkdir /etc/consul.d
	# 创建config.json添加如下内容,注意第一台机器必须以bootstrap方式启动,将bootstrap的值修改为true即可
	# 当集群server全部启动之后,在将第一台以bootstrap方式启动的机器的配置文件bootstrap的值修改为false在重启即可。
	#配置文件中server为该机器是否以server的方式启动
	# data_dir为数据目录,不用创建,consul会自动创建
	# node_name为主机名
	# addresses一个为http的地址,一个为rpc通信地址
	# start_join为consul server地址
	{
  		"bootstrap": false,
  		"server": false,
  		"datacenter": "dc1",
  		"data_dir": "/tmp/consul",
  		"node_name": "consul-agent1",
  		"addresses": {
  			  "http": "192.168.2.61",
  			  "rpc": "192.168.2.61"
  		},
 		 "start_join": [
    			"192.168.2.71",
    			"192.168.2.72",
   			 "192.168.2.73"
 		 ]
	}
	
### 3.启动consul

	# 先启动consul server,consul server配置文件中server的值为true,consul server节点一般为3,5,7
	 gocode/bin/consul agent -config-dir /etc/consul.d/
	# 启动完consul server之后在启动consul agent,consul agent配置文件中server的值为flase
	gocode/bin/consul agent -config-dir /etc/consul.d/
	
### 4.注册innosql服务

	# 在consul agent上注册innosql服务,在注册服务之前先写一个检查innosql数据库的状态的脚本,innosql数据库正常返回0,不正常返回大于1的数,下面这条curl脚本路径实在/root/下面.192.168.2.61这个ip地址为配置文件中addresses里面的地址
	curl -X PUT -d '{"Name":"innosql","Tags":["master"],"Port":3306,"Check":{"Script":"/root/mysqlcheck.sh","Interval":"10s"}}' http://192.168.2.61:8500/v1/agent/service/register
	
### 5.向consul server插入service/innosql/leader kv键值对

	curl -X PUT http://192.168.2.71:8500/v1/kv/service/innosql/leader
						
### 6.编辑bootstrap的配置文件(配置文件在bootstrap目录下conf目录下)

	# hostname为consul agent主机名(也就是innosql master的主机名)
	# ip为本机ip地址
	# port为innosql的端口号
	# username为innosql用户名
	# password为innosql的密码
	# service_ip为consul server的任意一台机器的ip
	#service_port为consul server的端口号默认为8500
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
	
### 7.执行bootstrap程序(bootstrap程序只在innosql的master上执行)
	
	进去bootstrap目录
	go build
	./bootstrap
	
### 8.在两台consul agent上配置watch.json(watch.json在/etc/consul.d目录下)
	
	# watch.json配置文件如下所示
	{
		"watches": [
 		   {
			 "type": "key",
    			 "key": "service/innosql/leader",
    			 "handler": "/root/mha-handlers/mha-handlers"
   		   }
  		]
	}

### 9.在两台consul agent上将mha-handlers程序放在/root/目录下

	mv mha-handlers /root/
	cd /root/mha-handlers/
	go build 
	
### 10.在两台consul agent上reload配置

	# 我机器上consul的启动程序在gocode/bin/目录下,-rpc-addr的值是/etc/consul.d/config.json里面addresses键值对里面的rpc的值,端口号默认为8400
	gocode/bin/consul reload -rpc-addr=192.168.2.61:8400
	
### 11.按innosql的主从故障切换流程测试

	Agent异常退出
	Agent与Innosql通信异常
	Agent与console server 通信异常
	主机掉电
	主机死机
	主机CPU跑满
	主机网络异常
	Innosql 服务crash
	Innosql 夯死