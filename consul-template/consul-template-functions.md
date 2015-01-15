## consul-template API Functions

### file

	{{file “/root/test.txt”}} # 引用本地磁盘文件中的内容
	
### key

	{{key "service/N2PC4F/leader"}} # 查询consul kv里面的键

	{{key "service/N2PC4F/leader@dc1"}} # 根据数据中心来查询consul kv里面的键

### ls

	{{range ls "service/N2PC4F@dc1"}} #查询以service/N2PC4F为前缀的一级键值队，比如能查询到/service/N2PC4F/name，但是查询不到/service/N2PC4F/name/aa，如果忽略数据中心的指定，就是在本地的数据中心查询
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
	
### services

	{{services}} # 查询catalog端点里面的所有服务
	
	{{services "@dc1"}} # 查询catalog端点的服务，可以指定要在那个数据中心查询，是个可选项
	
	{{range services}} #遍历catalog端点的服务
	{{.Name}} # 打印服务名
	{{range .Tags}} # 遍历catalog端点服务的tags
		{{.}}{{end}} # 打印tags
	{{end}}
	
### tree

	{{range tree "service/N2PC4F@dc1"}} # 查询以service/N2PC4F为前缀的键值队，比如能查询到/service/N2PC4F/name，也能查询到/service/N2PC4F/name/aa，如果忽略数据中心的指定，就是在本地的数据中心查询
	{{.Key}} {{.Value}}{{end}} # 打印出键和值
	
## consul-template Helper Functions	
	
### byTag

	{{range $tag,$services := service "N2PC4F" | byTag}} {{$tag}} # 打印出有N2PC4F服务的节点，并且根据Tags分类
	{{range $services}} server {{.Name}} {{.Address}}:{{.Port}}
	{{end}}{{end}}
	
### env

	{{env "PATH"}} # 读取给定环境变量的值
	
	{{env "PATH" | toLower}} # 这个函数可以链接到操作输出
	
### parseJSON

	{{with $d := key "service/mysql-1/leader" | parseJSON}}{{$d.Node}}{{end}} # 解析consul里面的键值对，值必须为json，consul 里面的键值对解析出值里面的某个键的值
	
	{{with $d := file "/root/test.json" | parseJSON}}{{$d.Node}}{{end}} #读取本地json文件，解析文件中指定键值对，并打印出相应的值。
	
### regexReplaceAll
	
	
	
	
	
	
	
	
	
	
	
	
	
	
	
