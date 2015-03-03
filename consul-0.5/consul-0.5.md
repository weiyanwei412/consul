## consul和Atlas的集成，托管consul ui
	
	环境为三台server二台agent

### 1.进入https://atlas.hashicorp.com/account/new

	创建Atlas账号:lindaneye
	
![Registration](./pic/001.jpg)
	
### 2.登陆Atlas账号

![Registration](./pic/002.jpg)
	
### 3.创建一个box

	创建一个新的box:lindan
	
![Registration](./pic/003.jpg)
	
### 4.创建一个Atlas的token

	创建token:xxxxxxxxx
	
![Registration](./pic/004.jpg)

	生成token成功如下图所示
	
![Registration](./pic/005.jpg)
		
### 5.设置环境变量

![Registration](./pic/006.jpg)

### 6.启动consul

![Registration](./pic/007.jpg)

### 7.进入Atlas查看

![Registration](./pic/008.jpg)

### 8.将第二台server加入集群

![Registration](./pic/009.jpg)

### 9.将第三台server加入集群

![Registration](./pic/010.jpg)

### 10.将第一台agent加入集群

![Registration](./pic/011.jpg)

### 11.将第二台agent加入集群

![Registration](./pic/012.jpg)

![Registration](./pic/013.jpg)

### 12.执行consul members显示集群成员

![Registration](./pic/014.jpg)

### 13.查看集群中成员和服务信息

![Registration](./pic/015.jpg)


	


