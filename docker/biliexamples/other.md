# Shortcut for future

```dockerfile
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:7.6.2

docker run -d --name mysql01 -p 3310:3306 -v /home/mysql/conf:/etc/mysql/conf.d -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -e TZ=Asia/Kuala_Lumpur mysql

docker run -d --name mysql01 -p 8806:3306 -e MYSQL_ROOT_PASSWORD=123456 -e TZ=Asia/Kuala_Lumpur mysql
```

```
=======
tomcat
=======
docker pull tomcat
docker run -it -p 8080:8080 --name t1 tomcat
// tips: 404 is normal, maybe not expose firewall
// maybe inside /usr/local/tomcat/webapps is nothing
// --> rename: mv webapss.dist webapps (for new version tomcat) --> ok

=======
mysql
=======
docker pull mysql:5.7
docker run -p 3306:3306 --privileged=true -v /abc/mysql/log:/var/log/mysql -v /abc/mysql/data:/var/lib/mysql -v /abc/mysql/conf:/etc/mysql/conf.d -e MYSQL_ROOT_PASSWORD=345678 -d --name mysql mysql:5.7

/abc/mysql/conf: vim my.cnf // to solve latin1 problem
[client]
default_character_set=utf8
[mysqld]
collation_server = utf8_general_ci
character_set_server = utf8

docker exec -it containerID /bin/bash
mysql -u root -p

=======
redis
=======
docker pull redis:6.0.8

cp redis.conf /app/redis/ // my config file
-- update --
-- enable redis external connection: 注释掉#bind 127.0.0.1
-- enable demonize : daemonize yes
-- optional: appendonly yes
-- optional: requirepass 123
-- optional: protected-mode no

docker run -d -p 6379:6379 --privileged=true -v /app/redis/redis.conf:/etc/redis/redis.conf -v /app/redis/data:/data --name redis redis:6.0.8 redis-server /etc/redis/redis.conf //连在一起, 让redis-server去读 /etc/redis/redis.conf

docker exec -it .... /bin/bash
redis-cli
```

```bash
==========
MYSQL -主从
==========
主
==
docker run -p 3307:3306 --name mysql-master \
-v /mydata/mysql-master/log:/var/log/mysql \
-v /mydata/mysql-master/data:/var/lib/mysql \
-v /mydata/mysql-master/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root  \
-d mysql:5.7

// 进入/mydata/mysql-master/conf目录下新建my.cnf
vim my.cnf

[mysqld]
## 设置server_id，同一局域网中需要唯一
server_id=101 
## 指定不需要同步的数据库名称
binlog-ignore-db=mysql  
## 开启二进制日志功能
log-bin=mall-mysql-bin  
## 设置二进制日志使用内存大小（事务）
binlog_cache_size=1M  
## 设置使用的二进制日志格式（mixed,statement,row）
binlog_format=mixed  
## 二进制日志过期清理时间。默认值为0，表示不自动清理。
expire_logs_days=7  
## 跳过主从复制中遇到的所有错误或指定类型的错误，避免slave端复制中断。
## 如：1062错误是指一些主键重复，1032错误是因为主从数据库数据不一致
slave_skip_errors=1062

docker restart mysql-master
docker exec -it mysql-master /bin/bash
mysql -uroot -proot
CREATE USER 'slave'@'%' IDENTIFIED BY '123456';
GRANT REPLICATION SLAVE, REPLICATION CLIENT ON *.* TO 'slave'@'%';

==
从
==
docker run -p 3308:3306 --name mysql-slave \
-v /mydata/mysql-slave/log:/var/log/mysql \
-v /mydata/mysql-slave/data:/var/lib/mysql \
-v /mydata/mysql-slave/conf:/etc/mysql \
-e MYSQL_ROOT_PASSWORD=root  \
-d mysql:5.7

# 进入/mydata/mysql-slave/conf目录下新建my.cnf
vim my.cnf

[mysqld]
## 设置server_id，同一局域网中需要唯一
server_id=102
## 指定不需要同步的数据库名称
binlog-ignore-db=mysql  
## 开启二进制日志功能，以备Slave作为其它数据库实例的Master时使用
log-bin=mall-mysql-slave1-bin  
## 设置二进制日志使用内存大小（事务）
binlog_cache_size=1M  
## 设置使用的二进制日志格式（mixed,statement,row）
binlog_format=mixed  
## 二进制日志过期清理时间。默认值为0，表示不自动清理。
expire_logs_days=7  
## 跳过主从复制中遇到的所有错误或指定类型的错误，避免slave端复制中断。
## 如：1062错误是指一些主键重复，1032错误是因为主从数据库数据不一致
slave_skip_errors=1062  
## relay_log配置中继日志
relay_log=mall-mysql-relay-bin  
## log_slave_updates表示slav
· show master status;e将复制事件写进自己的二进制日志
log_slave_updates=1  
## slave设置为只读（具有super权限的用户除外）
read_only=1

docker restart mysql-slave
#在主:show master status; 注意看master_log_file. position
docker exec -it mysql-slave /bin/bash
mysql -uroot -proot

change master to master_host='192.168.149.130', master_user='slave', master_password='123456', master_port=3307, master_log_file='mall-mysql-bin.000001', master_log_pos=617, master_connect_retry=30;

show slave status \G; # \G show key -- value instead of table
#看到 Slave_IO_Running 和 Slave_SQL_Running 还没开始
start slave;
show slave status \G;

```

> master_host：主数据库的IP地址；
>
> master_port：主数据库的运行端口；
>
> master_user：在主数据库创建的用于同步数据的用户账号；
>
> master_password：在主数据库创建的用于同步数据的用户密码；
>
> master_log_file：指定从数据库要复制数据的日志文件，通过查看主数据的状态，获取File参数；
>
> master_log_pos：指定从数据库从哪个位置开始复制数据，通过查看主数据的状态，获取Position参数；
>
> master_connect_retry：连接失败重试的时间间隔，单位为秒。

> **Redis 3主3从**

```bash
关闭防火墙+启动docker后台服务

docker run -d --name redis-node-1 --net host --privileged=true -v /data/redis/share/redis-node-1:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6381
 
docker run -d --name redis-node-2 --net host --privileged=true -v /data/redis/share/redis-node-2:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6382
 
docker run -d --name redis-node-3 --net host --privileged=true -v /data/redis/share/redis-node-3:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6383
 
docker run -d --name redis-node-4 --net host --privileged=true -v /data/redis/share/redis-node-4:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6384
 
docker run -d --name redis-node-5 --net host --privileged=true -v /data/redis/share/redis-node-5:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6385
 
docker run -d --name redis-node-6 --net host --privileged=true -v /data/redis/share/redis-node-6:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6386

docker exec -it redis-node-1 /bin/bash

构建主从关系
-----------
# --cluster-replicas 1 表示为每个master创建一个slave节点
redis-cli --cluster create 192.168.111.147:6381 192.168.111.147:6382 192.168.111.147:6383 192.168.111.147:6384 192.168.111.147:6385 192.168.111.147:6386 --cluster-replicas 1

# · 链接进入6381作为切入点，查看集群状态, 查看节点状态
redis-cli -p 6381 -c # -c (cluster)防止路由失效，which cannot save the data not in his slot
cluster info
cluster nodes

# 查看集群状态
redis-cli --cluster check 自己IP:6381

扩容 ==> 4主4从
--------------
docker run -d --name redis-node-7 --net host --privileged=true -v /data/redis/share/redis-node-7:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6387

docker run -d --name redis-node-8 --net host --privileged=true -v /data/redis/share/redis-node-8:/data redis:6.0.8 --cluster-enabled yes --appendonly yes --port 6388

redis-cli --cluster add-node 自己实际IP地址:6387 自己实际IP地址:6381
# 6387 就是将要作为master新增节点
# 6381 就是原来集群节点里面的领路人，相当于6387拜拜6381的码头从而找到组织加入集群

redis-cli --cluster check 真实ip地址:6381
# Result: 0 slots | 0 slaves。并没有槽号

# 重新分派槽号
redis-cli --cluster reshard IP地址:端口号
# redis-cli --cluster reshard 192.168.111.147:6381
##-- how many slots do you want to move (from 1 to 16384)
4096 #16384/4
##-- what is the receiving node ID?
asdfbh34hbadsf8y123hbakdf #新加入的node的ID
##-- Please enter the source node IDs
all

# 分配从
redis-cli --cluster add-node ip:新slave端口 ip:新master端口 --cluster-slave --cluster-master-id 新主机节点ID
#redis-cli --cluster add-node 192.168.111.147:6388 192.168.111.147:6387 --cluster-slave --cluster-master-id e4781f644d4a4e4d4b4d107157b9ba8144631451-------这个是6387的编号，按照自己实际情况

缩容 ==> 3主3从
--------------
#先删从
redis-cli --cluster del-node ip:从机端口 从机6388节点ID 
# redis-cli --cluster del-node 192.168.111.147:6388 5d149074b7e57b802287d1797a874ed7a1a284a8

#将6387的槽号清空，重新分配，本例将清出来的槽号都给6381
# if want to seperate balance, need to split three times for 6381, 6382, 6383
redis-cli --cluster reshard 192.168.111.147:6381
##--How many slots do you want to move (from 1 to 16384)?
4096
##--What is the receiving node ID? 
qewryhuiqweruu1239h # this is 6381 id, 让他接手空出来的slot
##-- Please enter the source node IDs
##-- Source node #1
basdhfkdasfihaifdh123 # this is 6387 id which to be deleted
##-- Source node #2
done

redis-cli --cluster del-node ip:端口 6387节点ID
# redis-cli --cluster del-node 192.168.111.147:6387 e4781f644d4a4e4d4b4d107157b9ba8144631451

```

**Tips:**

- 如果一个master死了，再回来不会变成master，会变成slave先
- 如果扩展了, slot不是平均分。 eg 3machine: 0-5460 , 5461-10922, 10923-16383 ===> 1365-5460, 6827-10922, 12288-16383, {0-1364, 5461-6826...}

