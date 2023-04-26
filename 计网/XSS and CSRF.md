# XSS and CSRF

[https://juejin.cn/post/6844903814311444493](https://juejin.cn/post/6844903814311444493)

## XSS

## 同源策略下，为什么还会存在csrf
* 同源策略不阻止发送请求 如前所述，SOP仅适用于xmlhttprequest。这意味着根据规范，浏览器必须发送ORIGIN报头以及通过xmlhttprequest发出的请求。
* 同源策略(Same Origin Policy, SOP)是一种浏览器级别的安全控制，它规定了由一个源提供服务的文档或脚本如何与来自其他源的资源进行交互。基本上，它可以防止在一个源下运行的脚本从另一个源读取数据。仍然允许跨域请求和表单提交，但不允许从其他来源读取数据。这意味着，如果你在一个易受攻击的网站上执行CSRF攻击，导致一些服务器端状态改变(例如，用户创建，文档删除等)，攻击将成功，但你将无法读取响应。简而言之，SOP只是防止读取来自不同来源的数据。它不包括用于执行CSRF攻击的跨域表单提交。
[https://stackoverflow.com/questions/33261244/why-same-origin-policy-isnt-enough-to-prevent-csrf-attacks](https://stackoverflow.com/questions/33261244/why-same-origin-policy-isnt-enough-to-prevent-csrf-attacks)
[https://security.stackexchange.com/questions/157061/how-does-csrf-correlate-with-same-origin-policy](https://security.stackexchange.com/questions/157061/how-does-csrf-correlate-with-same-origin-policy)

csrf 特征：

攻击⼀般来源于第三方域名 ccsrf 不能获取到 cookie，但是可以利用浏览器的特性去使用。 接口的所有参数都是可以预测的（攻击网站清楚要伪造请求接口的请求参数）

从前面提到的跨域知识点中，能了解到浏览器对于 cookie 也是存在同源限制的，也就是与 cookie（domain）处于不同源的网站，浏览器是不会让该网站获取到这个 cookie。那为什么csrf攻击还会成功呢？其实这个与浏览器使用
cookie 的方式有关。 浏览器使用 cookie 情况主要包括以下几点：

除了跨域 XHR 请求情况下，浏览器在发起请求的时候会把符合要求的 cookie 自动带上。(域名，有效期，路径，secure 属性)

跨域 XHR 的请求的情况下，也可以携带 Cookie。

浏览器允许跨域提交表单

也就是说，浏览器中有页面或网站向某个域名发送请求时，其请求都会自动带上该域名下的所有 cookie。

* E11同源策略： IE 11 不会在跨站CORS请求上添加Origin标头，Referer头将仍然是唯一的标识。最根本原因是因为IE 11对同源的定义和其他浏览器有不同，有两个主要的区别，可以参考MDN
  Same-origin_policy#IE_Exceptions

* 302重定向： 在302重定向之后Origin不包含在重定向的请求中，因为Origin可能会被认为是其他来源的敏感信息。对于302重定向的情况来说都是定向到新的服务器上的URL，因此浏览器不想将Origin泄漏到新的服务器上。

[https://juejin.cn/post/6844903934310498312](https://juejin.cn/post/6844903934310498312)
[https://tech.meituan.com/2018/10/11/fe-security-csrf.html](https://tech.meituan.com/2018/10/11/fe-security-csrf.html)

### referer and origin
[https://cloud.tencent.com/developer/article/1467299](https://cloud.tencent.com/developer/article/1467299)

### 同源策略与跨域以及怎么解决跨域，原因。同源本质（数据到了浏览器后被丢弃了）
[https://zhuanlan.zhihu.com/p/104984869](https://zhuanlan.zhihu.com/p/104984869)