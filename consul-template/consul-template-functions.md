## consul-template API Functions

### file

	{{file “/root/test.txt”}} # 引用本地磁盘文件中的内容
	
### key

	{{key "service/N2PC4F/leader"}} # 查询consul kv里面的键

	{{key "service/N2PC4F/leader@dc1"}} # 根据数据中心来查询consul kv里面的键

### ls

	{{range ls "service/N2PC4F@dc1"}} #查询以service/N2PC4F为前缀的键值队，如果忽略数据中心的指定，就是在本地的数据中心查询
	{{.key}} {{.Value}} {{end}} #打印键的缺省部分和值	
	
### nodes

	{{nodes}} # 查询consul里面的所有节点
	
	{{nodes "@dc1"}}
	
### service

	{{service "master.N2PC4F@dc1:3306"}} # 查询N2PC4F服务，tags为master，并且是在dc1的数据中心，端口号为3006的服务
	
	{{service "N2PC4F"}} # 所查询有N2PC4F服务，标签，数据中心，端口号都是可选项
	
	{{range service "N2PC4F@dc1"}} #  查询N2PC4F服务并且在dc1数据中心的所有机器
	server {{.Name}} {{.Address}}:{{.Port}} {{end}} # 打印出节点名，ip地址，和端口号
	
	{{service "N2PC4F" "any"}} # 返回所有状态的服务
	{{service "N2PC4F" "passing"}} # 返回passing状态的服务
	{{service "N2PC4F" "passing,warning,critical"}} # 返回passing，warning，critical状态的服务
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
