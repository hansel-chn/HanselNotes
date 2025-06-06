# HTTP

## 面试题

[面试题](https://xiaolincoding.com/network/2_http/http_interview.html#%E4%BB%80%E4%B9%88%E6%98%AF%E5%8D%8F%E5%95%86%E7%BC%93%E5%AD%98)

## GET 和 POST 方法都是安全和幂等的吗？

先说明下安全和幂等的概念：

在 HTTP 协议里，所谓的「安全」是指请求方法不会「破坏」服务器上的资源。 所谓的「幂等」，意思是多次执行相同的操作，结果都是「相同」的。 如果从 RFC 规范定义的语义来看：

GET 方法就是安全且幂等的，因为它是「只读」操作，无论操作多少次，服务器上的数据都是安全的，且每次的结果都是相同的。所以，可以对 GET
请求的数据做缓存，这个缓存可以做到浏览器本身上（彻底避免浏览器发请求），也可以做到代理上（如nginx），而且在浏览器中 GET 请求可以保存为书签。 POST
因为是「新增或提交数据」的操作，会修改服务器上的资源，所以是不安全的，且多次提交数据就会创建多个资源，所以不是幂等的。所以，浏览器一般不会缓存 POST 请求，也不能把 POST 请求保存为书签。 做个简要的小结。

GET 的语义是请求获取指定的资源。GET 方法是安全、幂等、可被缓存的。

POST 的语义是根据请求负荷（报文主体）对指定的资源做出处理，具体的处理方式视资源类型而不同。POST 不安全，不幂等，（大部分实现）不可缓存。

注意， 上面是从 RFC 规范定义的语义来分析的。

但是实际过程中，开发者不一定会按照 RFC 规范定义的语义来实现 GET 和 POST 方法。比如：

可以用 GET 方法实现新增或删除数据的请求，这样实现的 GET 方法自然就不是安全和幂等。 可以用 POST 方法实现查询数据的请求，这样实现的 POST 方法自然就是安全和幂等。 曾经有个笑话，有人写了个博客，删除博客用的是 GET
请求，他觉得没人访问就连鉴权都没做。然后 Google 服务器爬虫爬了一遍，他所有博文就没了。。。

如果「安全」放入概念是指信息是否会被泄漏的话，虽然 POST 用 body 传输数据，而 GET 用 URL 传输，这样数据会在浏览器地址拦容易看到，但是并不能说 GET 不如 POST 安全的。

因为 HTTP 传输的内容都是明文的，虽然在浏览器地址拦看不到 POST 提交的 body 数据，但是只要抓个包就都能看到了。

所以，要避免传输过程中数据被窃取，就要使用 HTTPS 协议，这样所有 HTTP 的数据都会被加密传输。

## 浏览器缓存

[浏览器缓存](https://juejin.cn/post/6844903593275817998)
[状态码](https://www.jianshu.com/p/faae1830d8b5)

## 同一端口下的应用层的不同协议报文，应用层怎么区分

当在同一端口下存在多个应用层协议时，应用层可以通过以下方法来区分不同协议的报文：

协议标识字段：不同的应用层协议通常会在报文中包含特定的标识字段或头部信息，用于标识协议类型。应用程序可以检查报文中的标识字段，根据其值来确定所使用的协议。例如，HTTP 协议的请求报文中包含 "GET"、"POST"
等方法标识字段，SMTP 协议的报文中包含以 "HELO"、"MAIL FROM" 开头的命令等。

报文格式和结构：每个应用层协议都有自己的报文格式和结构规范。应用程序可以根据报文的格式和结构来判断协议类型。例如，HTTP 协议的报文包含请求行、请求头和请求体，而 DNS 协议的报文包含查询类型和查询数据等。

连接建立过程：在应用层协议的连接建立过程中，可能存在特定的握手阶段或协商过程。应用程序可以根据连接建立时的握手或协商过程来判断协议类型。例如，在 WebSocket
协议中，客户端和服务器之间会进行握手，应用程序可以检查握手阶段的报文来判断是否为 WebSocket 协议。

上下文或会话状态：应用程序可能会维护一些上下文信息或会话状态，用于追踪连接或会话的属性。根据上下文或会话状态的信息，应用程序可以判断当前连接或会话所使用的协议类型。

通过结合以上的方法，应用层可以区分同一端口下的不同应用层协议报文。应用程序可以根据协议类型选择相应的解析和处理逻辑，以确保正确处理接收到的报文。

## websocket为什么采用http连接

* 因为防火墙可以限制进出网络的特定协议和端口，采用80端口来协商服务器是否支持websocket协议。若不支持仍可以使用http协议；如果直接采用某种协议连接可能直接被防火墙丢弃了没有回复。

* 基于https安全性

* 方便可以直接基于http路由去做

> 但是什么场景客户端会不知道服务端支持什么协议

* 可能突然由于一些原因服务器不支持协议了？但是还需要通信，所以要http？？，所以要存在两套逻辑？

```text
兼容性：通过使用 HTTP 协议进行初始握手连接，WebSocket 可以利用现有的 HTTP 基础设施和网络设备，例如代理服务器、防火墙等。大多数网络环境都允许通过常见的 HTTP 端口（如 80 和 443）进行通信。这样就可以确保 WebSocket 连接能够成功穿透网络设备，实现更好的兼容性。

简化部署：使用 HTTP 进行初始握手连接使得部署 WebSocket 应用程序变得更加简单。因为绝大多数 Web 服务器都支持 HTTP 协议，并且已经配置好相应的处理机制。这样一来，开发人员可以将 WebSocket 应用程序部署到现有的 Web 服务器上，而无需额外的配置或修改。

逐步升级：WebSocket 的握手过程允许逐步升级协议。当客户端发起 HTTP 请求时，服务器可以检查请求头中的特定字段来判断客户端是否支持 WebSocket。如果服务器支持 WebSocket，它会返回一个特定的握手响应，指示客户端可以升级到 WebSocket 协议进行后续通信。如果服务器不支持 WebSocket，它可以以标准的 HTTP 响应继续处理。这样的逐步升级机制可以保证与不同的客户端进行兼容，并允许服务器在支持和不支持 WebSocket 的客户端之间进行区分。
```

## 服务器端出现大量time_tait和close_wait原因

[https://zhuanlan.zhihu.com/p/591724475](https://zhuanlan.zhihu.com/p/591724475)

* 未开启`keepAlive`，不管是哪一方禁用`keepAlive`，定义`connection: close`，都是在服务器端关闭

> 1. “客户端禁用了 HTTP Keep-Alive，服务端开启 HTTP Keep-Alive”。 客户端请求的 `header` 定义了 `connection：close` 信息，连接不再重用，服务器端关闭
> 2. “客户端开启了 HTTP Keep-Alive，服务端禁用了 HTTP Keep-Alive” 在服务端主动关闭连接的情况下，只要调用一次 close() 就可以释放连接，剩下的工作由内核 TCP 栈直接进行了处理，整个过程只有一次 syscall；如果是要求 客户端关闭，则服务端在写完最后一个 response 之后需要把这个 socket 放入 readable 队列，调用 select / epoll 去等待事件；然后调用一次 read() 才能知道连接已经被关闭，这其中是两次 syscall，多一次用户态程序被激活执行，而且 socket 保持时间也会更长。

* 长连接超时，服务器主动断开连接

* HTTP 长连接的请求数量达到上限

## 队头阻塞及应对方法，http2和http3如何解决该问题的

[https://zhuanlan.zhihu.com/p/330300133](https://zhuanlan.zhihu.com/p/330300133)