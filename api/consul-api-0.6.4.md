page_title: v0.6.4 Add API
page_description: API Documentation for Consul

# Consul Remote API v0.6.4

# 1. Endpoints

## 1.1 Query

### 1.2

创建一个准备查询

	POST	/v1/query

**Example request**:

	{
  		"Name": "my-query", 
  		"Session": "adf4238a-882b-9ddc-4a9d-5b6758e4159e",
  		"Token": "",
  		"Near": "node1",
 		"Service": {
   			"Service": "redis",
   			"Failover": {
      			"NearestN": 3,
     			"Datacenters": [
        			"dc1",
        			"dc2"
      			]
    	 	},
    	"OnlyPassing": false,
   		"Tags": [
      		"master",
      		"!experimental"
    	]
  		},
  		"DNS": {
    		"TTL": "10s"
  		}
	}
	
Name:可选字段,可以用来查询

Session:可选字段,在会话无效时自动删除这个查询，如果没有提供需要手动删除，默认值为空

Token:可选字段,如果指定,是捕获的ACL令牌,即每次执行查询时再次用作ACL令牌,这使得客户查询较少,甚至没有ACL令牌被执行,所以这应该小心使用,令牌本身只能由客户机与管理令牌一起使用,如果字段为空或省略,客户端将使用ACL标记,以确定是否具有访问该服务的查询,如果客户端未提供一个ACL令牌,将使用匿名令牌,默认值为空

Near:可选字段,允许指定一个特定得节点基于网络坐标得距离排序,指定最近的实例节点将第一个被返回,随后的节点将按响应的往返时间做升序,如果给定的节点不存在,响应的节点将无序的

Service:查询服务的名称,这是必须指定的.

Failover:包含两个字段(NearestN,Datacenters),这两个都是可选的



