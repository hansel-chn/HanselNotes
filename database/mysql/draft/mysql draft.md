# Effects of Large and Small Tables as Driving Tables in Nested-Loop Join 

* 只考虑关联表都不存在索引的情况，疑惑这种情况下，小表驱动大表或者大表驱动小表真的有性能差异吗？
> [mysql doc](https://dev.mysql.com/doc/refman/8.0/en/nested-loop-joins.html)，首先，在无索引的情况下，在MYSQL 8.0.18之前，Block Nested-Loop Join Algorithm被用于优化无索引多表联查的情形中，MYSQL 8.0.20后，采用Hash Join替代。
>  [zhihu forum](https://www.zhihu.com/question/35906621) 讨论Block Nested-Loop Join Algorithm下扫描表的行数 N + λ * N * M，N越小扫描表的行数越少，但感觉N相对 λ * N * M来说，差别并不会很大。
>  官方doc没有找到join buffer如何具体存储数据的，只找到这句话，
>  - Only columns of interest to a join are stored in its join buffer, not whole rows.
>不同表columns of interest指的到底是什么不太清楚（只是指相关联的列还是最后select的列也包含），如果关联表的不同表之间的columns of interest相差很大，这个可能更会影响扫描表的性能？
>又犯毛病了，其实这种情况不用过多讨论，实际情况下还是需要设置合适的索引优化性能。
> 

# Explain
## Explain Output
* key_len 
>- n * m bytes, where n is the limit given (varchar(n)) and m is the potential number of bytes per character for the given character set (1 for latin1, 3 for utf8, 4 for utf8mb4)
>- plus 2 if VAR (varchar/varbinary)
>- plus 1 for NULL (even though NULLness might be stored using a single bit for the Engine and ROW_FORMAT in use.)

# Mysql Data Types
## CHAR and VACHAR
[mysql doc](https://dev.mysql.com/doc/refman/5.7/en/char.html)
In contrast to `CHAR`, `VARCHAR` values are stored as a 1-byte or 2-byte length prefix plus data. The length prefix indicates the number of bytes in the value. A column uses one length byte if values require no more than 255 bytes, two length bytes if values may require more than 255 bytes.

CHAR: When `CHAR` values are stored, they are right-padded with spaces to the specified length. When `CHAR` values are retrieved, trailing spaces are removed unless the [`PAD_CHAR_TO_FULL_LENGTH`](https://dev.mysql.com/doc/refman/8.0/en/sql-mode.html#sqlmode_pad_char_to_full_length) SQL mode is enabled.

VARCHAR: `VARCHAR` values are not padded when they are stored. Trailing spaces are retained when values are stored and retrieved, in conformance with standard SQL.

#  How to ensuring data integrity
> 断电等情况造成的数据库文件损坏

[https://www.51cto.com/article/696677.html](https://www.51cto.com/article/696677.html)

redolog 和 bin log提交顺序

1. 什么时候，哪个时间点之后可以被认为提交成功？为什么？和主从复制相关吗