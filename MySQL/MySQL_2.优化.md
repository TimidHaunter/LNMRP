- 了解版本和安装
- 了解MySQL逻辑分层
- 了解InnoDB和MyISAM
- SQL解析过程
- 索引
- B树，数据结构【树】
# 1.MySQL版本
8.0
5.x：5.4-5.x（**5.5** 5.7）
## 安装
yum
rpm
## 登录
> mysql -u root -p

报错，没有输入密码
ERROR 1045 (28000): Access denied for user 'root'@'localhost'
## 开机自启
> chkconfig mysql on // 打开开机自启
> chkconfig mysql off // 关闭开机自启
> ntsysv // 检查是否开机自启

## 安装路径
> ps -ef | grep mysql
> --datadir=/var/lib/mysql // 数据存放目录
> --pid-file=/var/lib/mysql/bigdata01.pid // pid文件目录，唯一标示符

# 2.原理
## MySQL逻辑分层
[MySQL体系架构简介](https://zhuanlan.zhihu.com/p/43736857)


![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642572121565-db32d75c-6e4a-4ef1-a5c7-9d97b9f7c474.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=308&id=uf02682fb&margin=%5Bobject%20Object%5D&name=image.png&originHeight=616&originWidth=508&originalType=binary&ratio=1&rotation=0&showTitle=false&size=226672&status=done&style=none&taskId=u4c96fd7e-f2b8-4f9d-8572-53322ac0e8f&title=&width=254)
应用层：
服务层
引擎层
存储层
## 支持那些引擎
```shell
show engines;
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| Engine             | Support | Comment                                                        | Transactions | XA   | Savepoints |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
| FEDERATED          | NO      | Federated MySQL storage engine                                 | NULL         | NULL | NULL       |
| MEMORY             | YES     | Hash based, stored in memory, useful for temporary tables      | NO           | NO   | NO         |
| InnoDB             | DEFAULT | Supports transactions, row-level locking, and foreign keys     | YES          | YES  | YES        |
| PERFORMANCE_SCHEMA | YES     | Performance Schema                                             | NO           | NO   | NO         |
| MyISAM             | YES     | MyISAM storage engine                                          | NO           | NO   | NO         |
| MRG_MYISAM         | YES     | Collection of identical MyISAM tables                          | NO           | NO   | NO         |
| BLACKHOLE          | YES     | /dev/null storage engine (anything you write to it disappears) | NO           | NO   | NO         |
| CSV                | YES     | CSV storage engine                                             | NO           | NO   | NO         |
| ARCHIVE            | YES     | Archive storage engine                                         | NO           | NO   | NO         |
+--------------------+---------+----------------------------------------------------------------+--------------+------+------------+
show engines \G; // 便于查看
```
默认支持 InnoDB
```shell
show variables like '%storage_engine%';
+----------------------------------+-----------+
| Variable_name                    | Value     |
+----------------------------------+-----------+
| default_storage_engine           | InnoDB    |
| default_tmp_storage_engine       | InnoDB    |
| disabled_storage_engines         |           |
| internal_tmp_disk_storage_engine | InnoDB    |
| internal_tmp_mem_storage_engine  | TempTable |
+----------------------------------+-----------+
```
​

## InnoDB 和 MyISAM
事务优先：高并发操作
性能优先

行锁：一次锁一行，锁的多，性能必然降低，但是不容易出错
表锁：一个表

# 3.SQL解析过程、索引、B树


