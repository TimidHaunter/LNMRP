## SQL分类
DDL：数据`定义`语言，这些语句定义了不同的数据段，数据库、表、列、索引等数据库对象。常用的语句关键字主要包括`create,drop,alter`等。

DML：数据`操纵`语句，用于添加，删除，更新和查询数据库记录，并检查数据完整性。常用的语句关键字主要包括`insert,delete,update,select`等。

DCL：数据 `控制`语句，用于控制不同数据段直接的许可和访问级别的语句。这些语句定义了数据库、表、字段，用户的访问权限和安全级别。主要的语句关键字包括`grant,revoke`等。



编码格式：指的是某个字符范围的编码规则。

utf8mb4 



排序规则：是针对某个字符集中的字符比较大小的一种规则。

utf8mb4_general_ci



存储引擎

不同的存储引擎，数据的存储格式不一样。

例如：Memory，不用磁盘来存储，关闭数据库服务，数据就消失了。



页格式

Q：MySQL默认存储引擎是InnoDB，数据放在磁盘上，但是处理数据在内存。怎么把磁盘数据加载到内存？

A：将数据划分为若干个页，以页作为磁盘和内存之间交互的基本单位，InnoDB中页的大小一般为 16 KB，最小存储单元。



行格式

以记录为单位来向表中插入数据的，这些记录在磁盘上的存放方式也被称为`行格式`或者`记录格式`。

Compact

Redundant

Dynamic

Compressed



指定行格式语法

```mysql
ALTER TABLE `integer` ROW_FORMAT=compact;

CREATE TABLE `integer` (
  `tinyint_a` tinyint unsigned NOT NULL COMMENT '无符号tinyint显示长度为10',
  `tinyint_b` tinyint NOT NULL COMMENT '有符号tinyint',
  `tinyint_c` tinyint(3) unsigned zerofill NOT NULL COMMENT '无符号tinyint，显示长度3',
  `tinyint_d` tinyint(8) unsigned zerofill NOT NULL COMMENT '无符号tinyint，显示长度3'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=COMPACT COMMENT='整数';
```



InnoDB数据页

![show_engines.png](G:\data\LNMRP\$Image\MySQL\innodb_data_page.png)

行记录

![show_engines.png](G:\data\LNMRP\$Image\MySQL\row_data.png)



delete_mask

min_rec_mask

next_record：从当前记录的真实数据到下一条记录的真实数据的地址偏移量。



多条记录通过`next_record`指定顺序，其实就是一个链表（记录在页中按照主键值由小到大顺序串联成一个单链表）

![show_engines.png](G:\data\LNMRP\$Image\MySQL\many_row_data.png)



页目录

查找，最笨的方法从infimum（最小记录）开始，一直找。遍历查找效率很慢。

二分法找到对应的槽。















