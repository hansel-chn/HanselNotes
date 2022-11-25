# Elasticsearch

## es查询

之前讲过Elasticsearch 的wildcard（通配符查询）、regexp（正则查询）、prefix（前缀查询），他们都是致力于模糊搜索，然后在实际的项目中该如何选择，稍不注意就可能到很大性能问题。

使用方式这里就不再赘述了，他们都是基于词条查询，它们也需要遍历倒排索引中的词条列表来找到所有的匹配词条，然后逐个词条地收集对应的文档ID。

针对Numeric datatypes（long, integer, short, byte, double, float....）

基本上不要使用，那样做意义真的不大，另外要关注下数值类型和Term Query有重大变化的介绍。

针对文本类型（text和keyword）

这一类大概是主流需求，

当搜索字段是text类型时：由于它会分词，在执行wildcard、regexp、prefix时和es会检查字段中的每个词条，而不是整个字段。

当搜索字段是keyword类型时：在执行wildcard、regexp、prefix时和es会检查字段中整个文本

prefix查询

如果满足你的需求，前缀匹配是优于wildcard和regexp。

regexp查询和wildcard查询

避免使用一个以通配符开头的模式(比如，*foo或者正则表达式: .*foo)，运行这类查询是非常消耗资源的。

最后再提醒下，如果你想了解它的执行过程及耗时情况（优化项从这里分析），查询是添加profile语法。

## ElasticSearch 中的 Mapping

ES 中的 [Mapping](https://www.cnblogs.com/codeshell/p/14445420.html) 相当于传统数据库中的表定义，它有以下作用：

* 定义索引中的字段的名字。
* 定义索引中的字段的类型，比如字符串，数字等。
* 定义索引中的字段是否建立倒排索引。 一个 Mapping 是针对一个索引中的 Type 定义的：

ES 中的文档都存储在索引的 Type 中 在 ES 7.0 之前，一个索引可以有多个 Type，所以一个索引可拥有多个 Mapping 在 ES 7.0 之后，一个索引只能有一个 Type`（对应"_type": "_doc"）`，所以一个索引只对应一个 Mapping

字段(属性)
一个 mapping 类型中定义了与文档相关的多个字段（属性)

## ES基础
[es基础](https://learnku.com/articles/40400)

[es](https://www.cnblogs.com/qdhxhz/p/11448451.html)

Index，shard，replication （Index类似于kafka中的topic，shard类似于kafka中的partition）

Doc是Index的单条记录，等同于数据库中的行

## ES选举

ES的选举
[主节点选举，分片选举](http://blog.itpub.net/9399028/viewspace-2666851/)
* Master的选举（选举主节点）
* shard的选举（选举主分片）

>Master的选举（选举主节点） Bully选举算法

## ES分片
[扩增分片](https://blog.51cto.com/lee90/2467377)