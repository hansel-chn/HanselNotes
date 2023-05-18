# SUMMARY

## note

网络层：IP数据报（Datagram)
传输层：UDP数据报（Datagram)
传输层：报文段（Segment）TCP

IP层：数据包(Packet)， 数据链路层：数据帧（Frame)，

分组/包(packet)
分组是在网络中传输的二进制格式的单元，为了提供通信性能和可靠性，每个用户发送的数据会被分成多个更小的部分。在每个部分的前面加上一些必要地控制信息组成的首部，有时也会加上尾部，就构成了一个分组。它的起始和目的地是网络层。

数据报(datagram)
面向无连接的数据传输，其工作过程类似于报文交换。采用数据报方式传输时，被传输的分组称为数据报。通常是指起始点和目的地都使用无连接网络服务的网络层的信息单元。(指IP数据报)
数据报:这在两层中使用。如果网络协议是IP，则数据单位称为数据报。在传输层，如果协议是UDP，我们也使用数据报。因此，我们将其区分为UDP数据报、IP数据报。

## MAC地址

MAC地址由12位16进制数组成，分成6组，每组2位，组与组之间用-相连。如00-16-EA-AE-3C-40，就是一个MAC地址。

MAC地址，全称Media Access Control
Address，直译为媒体存取控制地址。MAC地址又名物理地址、硬件地址、以太网地址，是每个网卡出厂时其厂家赋予其的一个全球独一无二的序列号。所以说，MAC地址并不是对电脑而言的，而是针对网卡的，计算机上有几个网卡，就有几个MAC地址。因此，我们在描述MAC地址时，不能说一台电脑的MAC地址，而应该说一个网卡的MAC地址。一台计算机上可能有网卡也可能没有网卡，可能有一个网卡也可能有多个网卡。一般来说，一台电脑至少有一个网卡，也就存在着至少一个MAC地址。拿笔记本电脑来说，联网方式有有线和无线两种，网卡也就有有线网卡和无线网卡两个，所以笔记本一般存在两个MAC地址。台式机抛开自行安装的无线网卡不谈，一般只有有线联网这一种联网方式，也就只有一张有线网卡，所以台式机一般只有一个MAC地址。

## ARP

ARP包就是一个广播帧（注：经过数据链路成封装后的Mac帧，帧是在数据链路层的一种叫法） ARP地址解析协议：在主机发送帧前将目的IP地址转换为目的MAC地址、 ARP缓存表里存有对应的IP地址和MAC地址
基于功能来考虑，ARP是链路层协议；基于分层/包封装来考虑，ARP是网络层协议。(此方法对于ICMP协议同样管用)

## 广播和泛洪

在泛洪时，交换机将帧发送给所有，因为它不知道如何到达目的地。 在广播中，创建帧的主机本身向每个人发送帧。

当交换机接收到单播帧(具有特定mac地址的帧，用于特定设备)时，它在mac地址表中查找帧的目的mac。如果在它的表中没有目标mac的条目，它将简单地将帧发送给连接到它的端口的每个人。这个帧仍然是一个<mark>**单播帧**</mark>
，因为它的报头中有一个特定的目的mac地址。该开关不会改变帧中的任何报头数据。所有接收到帧的设备都会比较帧中的目的mac地址和自己的mac地址，如果不匹配就丢弃帧。

广播:广播帧由主机自己创建。如果目的mac地址为`ffffffffffff`，则该帧变为<mark>**广播帧**</mark>。通常情况下，主机在ARP进程中创建这种类型的帧。当交换机接收到广播帧时，它将它发送给连接到它的每个人。

## ICMP协议

TTL是什么意思?TTL是 Time To Live的缩写，该字段指定IP包被路由器丢弃之前允许通过的最大网段数量。
ICMP（Internet控制消息协议）是IP协议的辅助协议。ICMP协议用来传递网络设备之间的查错和控制信息，起到收集各种网络信息、诊断和排除网络故障的作用，大大提升了IP数据报文交互成功的机会。
ICMP功能大致分为两类：差错通知和信息查询
[ICMP](https://zhuanlan.zhihu.com/p/387469317)
ICMP中的tracert命令 tracert命令用于调查与目的主机通信的所有经过的地址，跟ping命令一样，也是ICMP中的典型代表之一。

## 状态码

> 400 错误请求  
> 401 访问未授权  
> 403 访问禁止，拒绝请求  
> 找不到请求的网页

## MAC地址

6字节，用于在网络中唯一标识网卡

## CA证书

首先 CA 会把持有者的公钥、用途、颁发者、有效时间等信息打成一个包，然后对这些信息进行 Hash 计算，得到一个 Hash 值；

然后 CA 会使用自己的私钥将该 Hash 值加密，生成 Certificate Signature，也就是 CA 对证书做了签名；

最后将 Certificate Signature 添加在文件证书上，形成数字证书；

## tcp和udp连接问题

## 交换机和路由器

[https://www.163.com/dy/article/EOSTK9SD0531A5J3.html](https://www.163.com/dy/article/EOSTK9SD0531A5J3.html)

* 在同一网段中，直接二层，通过arp协议等发送信息
* 不在同一个网段中，通过路由器

> 路由模式和nat模式的区别

* 若内网采用了私网地址，则需要NAT模式进行地址转换；路由模式内网采用了合法公网地址
* NAT地址转换会在路由器出口改变数据包的目的ip和端口

### 路由器-不同网段

DHCP获取ip的时候可以获得一些其他的信息如dns，网关，掩码

* 如何找网关mac地址：根据网关ip进行ARP请求。

路由表作用原理：[https://zhuanlan.zhihu.com/p/145946764](https://zhuanlan.zhihu.com/p/145946764) <br>
[https://mp.weixin.qq.com/s/ktahxXMDtDVufyigU49bXg](https://mp.weixin.qq.com/s/ktahxXMDtDVufyigU49bXg)

### 二三四层交换机

二层交换机-第二层网桥 三层交换机-第三层路由 四层交换机-第四层应用端口号

### MTU为什么存在

[https://developer.aliyun.com/article/222535](https://developer.aliyun.com/article/222535)
MTU的设置基于并发和传输效率的考量
> Maximum transmission unit
> Maximum segment size

* 建立连接时双方协商MSS值

### IP数据包的总长度最大65535

由于IP数据报头部表示总长度的字段为16位

### TLS 建立在 TCP 协议基础上，在 TCP 连接成功建立后，执行握手协议验证证书并协商主加密密钥。

### ACK丢失不会重传，会重传FIN报文等

重要[https://www.cnblogs.com/rexcheny/p/11143128.html](https://www.cnblogs.com/rexcheny/p/11143128.html)
[https://xiaolincoding.com/network/3_tcp/tcp_optimize.html#%E4%B8%BB%E5%8A%A8%E6%96%B9%E7%9A%84%E4%BC%98%E5%8C%96](https://xiaolincoding.com/network/3_tcp/tcp_optimize.html#%E4%B8%BB%E5%8A%A8%E6%96%B9%E7%9A%84%E4%BC%98%E5%8C%96)