- 了解版本和安装
- 了解MySQL逻辑分层
- 了解InnoDB和MyISAM
- SQL解析过程
- 索引
- B树，B+树，数据结构【树】
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
服务层：
引擎层：
存储层：
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
​

行锁：一次锁一行，锁的多，性能必然降低，但是不容易出错
表锁：一个表
​

# 3.SQL解析过程、索引、B树
SQL性能低，执行时间长，等待时间长，SQL语句太差（join、子查询），索引失效，服务器参数设置（线程数）
## 编写过程
> select .. from .. join .. on .. where .. group by .. having .. order by .. limit ..

## 解析过程
[步步深入：MySQL架构总览->查询执行流程->SQL解析顺序](https://www.cnblogs.com/annsshadow/p/5037667.html)
```sql
FROM <left_table>
ON <join_condition>
<join_type> JOIN <right_table>
WHERE <where_condition>
GROUP BY <group_by_list>
HAVING <having_condition>
SELECT
DISTINCT <select_list>
ORDER BY <order_by_condition>
LIMIT <limit_number>
```
从哪里获取，不断的过滤条件，要选择一样或不一样的，排好序，那才知道要取前几条呢。
![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642667408528-de14a913-6be0-4f3d-9898-2f569052a580.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=380&id=u5db48ef6&margin=%5Bobject%20Object%5D&name=image.png&originHeight=760&originWidth=966&originalType=binary&ratio=1&rotation=0&showTitle=false&size=330679&status=done&style=none&taskId=u89e25a53-d2bb-43b9-a54b-53d986da60b&title=&width=483)
## B树和B+树
[MySQL索引背后的数据结构及算法原理](http://blog.codinglabs.org/articles/theory-of-mysql-index.html)
[B树、B+树详解](https://www.cnblogs.com/lianzhilei/p/11250589.html)​
《MySQL技术内幕  InnoDB存储引擎》，第五章 索引与算法
​

B树，B-tree，B=Balance，平衡多路查找树
B+树，在B树上做**改进**


### B树
![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642748120120-d8337f9c-b69e-4b66-b21d-a3a60005f078.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=445&id=u37c73843&margin=%5Bobject%20Object%5D&name=image.png&originHeight=890&originWidth=1868&originalType=binary&ratio=1&rotation=0&showTitle=false&size=169469&status=done&style=none&taskId=u7811a979-9534-47c3-82f9-64af2463000&title=&width=934)
也可以是三叉树，17、35黑色块是数据块；P1、P2、P3红色块是指针；白色是磁盘。数据块就是指针的分界线。
如果要找28，只需要找三次

- 在17和35之间，通过P2指针找到下一层
- 在26和30之间，通过P2指针找到下一层
- 找到28

三层B-tree，可以存放上百万数据，可以拥有多个子节点。
​

### B+树
![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642748973875-df3b286a-f328-484f-884f-f8ba082b36e7.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=641&id=u2fdecdd0&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1282&originWidth=1910&originalType=binary&ratio=1&rotation=0&showTitle=false&size=328425&status=done&style=none&taskId=u51e0d707-7a9d-4b97-9f51-96af373577d&title=&width=955)

- 中间节点不保存数据，只用来索引，数据都在叶子节点上
- 叶子节点，是自小到大的顺序链接

​

## 索引
### 原理
优化SQL就是优化索引，索引就是目录
indexs，帮助MySQL高效获取数据的**数据结构**（B+树），就是**树**
将索引的字段，进行树状化
​

> 为啥要用B+树？
```sql
INSERT INTO user(name,age) 
VALUES
('niuniu','30'),
('bainiangzi','600'),
('zhangshuoshi','29'),
('chegong','28');
```
> 插入数据，给 age 字段加上普通索引

![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642699630054-915a1fce-b83d-41a2-8b5e-c702e61811b8.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=178&id=u8f687d93&margin=%5Bobject%20Object%5D&name=image.png&originHeight=356&originWidth=346&originalType=binary&ratio=1&rotation=0&showTitle=false&size=54111&status=done&style=none&taskId=ub69a9414-9e59-4486-a6ae-c1498e13742&title=&width=173)

> 小的放左，大的放右

![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642746361265-8ccf04af-fd0f-48be-a4eb-f6a37dba677c.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=376&id=u07f4659b&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1504&originWidth=908&originalType=binary&ratio=1&rotation=0&showTitle=false&size=1965879&status=done&style=none&taskId=u84a9ed6d-7ca9-4a51-bc4c-4e03b1f6779&title=&width=227)
按照索引字段排好B树，31对应数据硬件地址（16进制）
> select * from user where age = 30;

不加索引，一条条查，要查5次；
加索引，查3次。
​

### 缺点
索引本身很大，很占空间
索引不是所有情况都适用

- 少量数据不需要
- 频繁更新的字段不适合-更新一次，需要重新排序
- 很少使用的字段

索引会降低增删改效率，不敢要操作数据列，还要操作索引树
​

### 优点
提高查询效率（降低IO使用率）
降低CPU使用率
​

​

索引只能加给数字字段吗？非数字字段怎么排序。
​

### 分类

- 单值索引：单列，一个表可以有多个单值索引
- 主键索引：字段值不能重复，该字段值不能是**NULL**
- 唯一索引：字段值不能重复（age=23，很多人都有23岁），一般为ID
- 复合索引：多个列构成的索引（相当于书的二级目录，z:zhao）(name:age，找张三，有两个张三，再找23岁的张三；如果只有一个张三，就不用找年龄了)
### 使用
添加索引
> create 索引类型 索引名 on 表名(字段)
> create index dept_index on tb(dept); // 单值
> create unique index name_index on tb(name); // 唯一索引
> create index name_dept_index on tb(name,dept); // 复合索引
> ​

> alter table tb add index dept_index(dept);
> alter table tb add unique index name_index(name);
> alter table tb add index name_dept_index(name,dept);

​

删除索引
> drop table tb; // 暴力删除，删除表，对应的索引就没有了
> ​

> drop index 索引名 on 表名
> drop index name_index on tb; // 删除索引，不用加索引类型



查询索引
> show index from tb;

![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642770141256-188c8e30-ad12-47a5-8056-dc58f12c620d.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=126&id=uf47f8a8d&margin=%5Bobject%20Object%5D&name=image.png&originHeight=252&originWidth=1536&originalType=binary&ratio=1&rotation=0&showTitle=false&size=132300&status=done&style=none&taskId=ue913fc47-abf8-4d37-9ae6-7fdac7db947&title=&width=768)


如果字段是主键，该字段默认为主键索引
​

# 4.SQL性能问题
[MySQL8.0优化文档](https://dev.mysql.com/doc/refman/8.0/en/optimization.html)
## 执行计划
explain，可以模拟SQL优化器执行SQL语句（手动优化）
explain SQL语句
> explain select * from goods;

![image.png](https://cdn.nlark.com/yuque/0/2022/png/1927971/1642772493085-8017476d-4fa9-449e-8468-c66405d779b9.png#clientId=u53d83fd1-b386-4&crop=0&crop=0&crop=1&crop=1&from=paste&height=53&id=u6a6edad4&margin=%5Bobject%20Object%5D&name=image.png&originHeight=106&originWidth=1868&originalType=binary&ratio=1&rotation=0&showTitle=false&size=53899&status=done&style=none&taskId=uf5c04b31-9ab9-4e23-9e6b-3e90a2aa4fc&title=&width=934)
> id：查询编号
> select_type：查询类型
> table：表名
> partitions：
> type：
> possible_keys：预测用到的索引
> key：实际用到的索引
> key_len：实际使用索引长度
> ref：表之间引用关系
> rows：通过索引查到的数据量
> filtered：
> Extra：额外信息

​

### 测试数据
user

| id | name |
| --- | --- |

goods

| id | user_id | category_id | title | desc | is_on | is_recommend |
| --- | --- | --- | --- | --- | --- | --- |

category

| id | name |
| --- | --- |



```sql
# 添加单值索引
create index goods_user_id_index on goods(user_id);
create index goods_category_id_index on goods(category_id);
create index goods_is_on_index on goods(is_on);
create index goods_title_index on goods(title);
create index goods_is_recommend_index on goods(is_recommend);

# 删除索引
drop index goods_user_id_index on goods;
drop index goods_category_id_index on goods;
drop index goods_is_on_index on goods;
drop index goods_title_index on goods;
drop index goods_is_recommend_index on goods;

# 查询user_id为3，或者category_id为15的商品信息
select g.* from goods as g, users as u, categories as c 
where u.id=g.user_id 
and c.id=g.category_id 
and (user_id=3 or category_id=16);

explain
select g.* from goods as g, users as u, categories as c 
where u.id=g.user_id 
and c.id=g.category_id 
and (user_id=3 or category_id=16);
```
### id
id值相同，顺序从上向下执行，先执行u再c然后g表，u8->c57->g1020
表的执行顺序，因数量的改变而改变，原因是笛卡尔积

| a | b | c |  | 笛卡尔积 |
| --- | --- | --- | --- | --- |
| 3 | 4 | 5 |  | 3*4=12*5=60 |
| 5 | 4 | 3 |  | 5*4=20*3=60 |

明显第二种方法比较**占**内存，因为计算过程大（20），数据小的表优先查询
​

id值不同，越大越优先，子查询的时候有优先级
```sql
# 查询email='root@163.com'，添加了哪些的字段的商品
select distinct(c.name) 
from goods as g, users as u, categories as c
where u.id=g.user_id 
and c.id=g.category_id 
and u.email='root@163.com';

explain
select distinct(c.name) 
from goods as g, users as u, categories as c
where u.id=g.user_id 
and c.id=g.category_id 
and u.email='root@163.com';

# 多表查询可以改为子查询
explain
select distinct(c.name)
from categories as c
where c.id in
(select g.category_id from goods as g where g.user_id=
	(select u.id from users as u where u.email='root@163.com')
);
```








## 查询优化器
查询优化器会干扰优化（自动优化）
​

