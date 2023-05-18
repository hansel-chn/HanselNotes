# Socket

## poll epoll select

[https://mp.weixin.qq.com/s?__biz=MjM5Njg5NDgwNA==&mid=2247484905&idx=1&sn=a74ed5d7551c4fb80a8abe057405ea5e&scene=21#wechat_redirect](https://mp.weixin.qq.com/s?__biz=MjM5Njg5NDgwNA==&mid=2247484905&idx=1&sn=a74ed5d7551c4fb80a8abe057405ea5e&scene=21#wechat_redirect)

1. tcp_v4_rcv 中首先根据收到的网络包的 header 里的 source 和 dest 信息来在本机上查询对应的 socket。
2. 调用 tcp_queue_rcv 函数中完成了将接收数据放到 socket 的接收队列上。
3. 接收完成，调用sock_def_readable，sock_def_readable判断等待队列不为空，执行等待队列上的函数ep_poll_callback，ep_poll_callback 根据等待任务队列项上的额外的 base
   指针可以找到 epitem， 进而也可以找到 eventpoll对象。
4. 把自己的 epitem 添加到 epoll 的就绪队列rdllist中。
5. 接着它又会查看 eventpoll 对象上的等待队列里是否有等待项（epoll_wait 执行的时候会设置）。
6. 如果没执行软中断的事情就做完了。如果有等待项，那就查找到等待项里设置的回调函数。

[https://imageslr.com/2020/02/27/select-poll-epoll.html](https://imageslr.com/2020/02/27/select-poll-epoll.html)
[https://juejin.cn/post/6881596144963551245](https://juejin.cn/post/6881596144963551245)

> epoll怎么解决性能开销大的问题

通过回调函数的方式在，

> epoll原理
[https://developer.aliyun.com/article/1097552](https://developer.aliyun.com/article/1097552)

* epoll_create先创建一个epoll实例，返回一个指向该实例的文件描述符。这个epoll实例内部存储了一个红黑树和双向链表，红黑树包括了所有待监听的文件描述符，双向链表存储了事件就绪的文件描述符。
* epoll_ctl（control缩写ctl，功能有点像select中 fdset的几个api（添加要监听的文件描述符））将待监听的文件描述符加入epoll实例的监听列表中，当事件准备就绪，中断程序将文件描述符添加到就绪链表中。
* epoll_wait功能相当于select，检查就绪队列，为空阻塞，不为空返回

[最全面epoll](https://cloud.tencent.com/developer/news/787829)
在 epoll_ctl 中首先根据传入 fd 找到 eventpoll、socket相关的内核对象 。对于 EPOLL_CTL_ADD 操作来说，会然后执行到 ep_insert 函数。所有的注册都是在这个函数中完成的。 对于每一个
socket，调用 epoll_ctl 的时候，都会为之分配一个 epitem。

epoll_ctl过程中， 在 ep_ptable_queue_proc 函数中，新建了一个等待队列项，并注册其回调函数为 ep_poll_callback 函数。然后再将这个等待项添加到 socket
的等待队列中。<font color=LightCoral>在socket等待队列注册了回调函数，使得数据就绪时可以将文件描述符加入epoll实例的就绪链表</font>

epoll 到底用没用到 mmap？ 没有

> 边缘触发和水平触发

指的是epoll的方式边缘触发和水平触发

*

个人觉得问题在for循环上，因为边缘io只有在文件描述符状态变化时才发生通知，所以需要使用循环来处理（不管阻塞式io还是非阻塞io都需要使用循环处理读完缓冲区，这个过程需要多次，因为每次读的内容定长。但是最后一次读取，非阻塞io可以返回没有数据可读的错误信息，不会阻塞；而阻塞式io回阻塞直到新的数据）

* 水平触发可以使用阻塞i/o，不需要使用循环结构，所以不会发生上述情况。

但是，存在另一种情况，当某个socket接收缓冲区有新数据分节到达，然后select报告这个socket描述符可读，但随后，协议栈检查到这个新分节检验和错误，然后丢弃这个分节，这时候调用read则无数据可读，如果socket没有被设置nonblocking，此read将阻塞当前线程。      
所以I/O多路复用不管采用什么触发模式，为了减少出错，绝大部分时候是和非阻塞的socket联合使用。

## 底层中断

对于外部中断，CPU在执行当前指令的最后一个时钟周期去查询INTR引脚，若查询到中断请求信号有效，同时在系统开中断（即IF=1）的情
况下，CPU向发出中断请求的外设回送一个低电平有效的中断应答信号，作为对中断请求INTR的应答，系统自动进入中断响应周期。

## 为什么linux使用epoll多路复用，而不是用异步io？

linux目前异步io支持的不好，多路复用能解决大部分问题，折中方案

## socket对应唯一的四元组

操作系统针对不同的传输方式（TCP，UDP）会在内核中各自维护一个Socket双向链表，当数据包到达网卡时，会根据数据包的源端口，源ip，目的端口从对应的链表中找到其对应的Socket，并会将数据拷贝到Socket的缓冲区，等待应用程序读取。

socket 的本质是一种资源，它包含了端到端的四元组信息，用来标识数据包的归属。因此，尽管 tcp 协议的端口号只有 65535 个，但是进程可拥有的 socket 数据却不限于此（受限于进程最大文件描述符数据）；

## 一致性哈希