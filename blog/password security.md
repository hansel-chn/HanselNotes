# Abstract

* Don't store passwords plain-text -> hashing(easy to calculate but hard to reverse)
* Weak password -> hash -> salted hash(due to rainbow table, etc.)
* hashing on front end or back end
# Security
## prerequisite
错误认识：hashing password与数据传输相关

* Transport security
 > 数据传输阶段的安全性有TLS保证，密码明文传输在HTTPS下无问题(hash password 目的并不是为了解决密码明文传输的问题，明文在HTTPS下安全)--延申[[#man-in-the-middle attack (MITM)]]。
* Database security 
> hashing password目的是为了保证数据库安全性，当数据库泄露时，即使攻击者可以访问数据库，也不能直接获取到明文的密码。

[Transport security and Database security](https://security.stackexchange.com/questions/110948/password-hashing-on-frontend-or-backend/#110949)
##  Hashing password
Hashing password保证了数据库泄露时，攻击者不能直接获取用户的明文password，然而，weak password 和某些hash函数减弱了破解难度。
### Vulnerability
* [Password entropy](https://advancedweb.hu/how-to-hash-passwords-and-when-not-to/)
> low-entropy password 带来的不安全性，使密码容易被破解(比如利用 [rainbow table](https://en.wikipedia.org/wiki/Rainbow_table#:~:text=A%20rainbow%20table%20is%20a,form%2C%20but%20as%20hash%20values.),low-entropy password经常会出现在常见的rainbow table中)

* [Speed of  hash computation](https://advancedweb.hu/how-to-hash-passwords-and-when-not-to/)
> Fast hash functions such as SHA1, MD5.快速hash使得攻击者可以更快的尝试不同的password。所以建议故意一些更慢的hash增加攻击者的代价，如bcrypt。一些更为现代的hash算法不仅设计的较慢而且会在其他地方如内存特殊设计，降低攻击者的破解效率。
### Improvement
*  High-entropy passwords
*  Hash functions that increase the cost of the attacker
* salted hash
>  Salted hash增加安全性。通过加盐迫使攻击者单独的攻击每一个密码。
>  
>  讨论salted hash中salt是否有必要随机核心要理解加盐的目的。加盐的通常做法是为每一个password生成不同的盐值。要注意的是，加盐的目的并不是为了阻碍攻击者获取salt，通常情况下，攻击者获取了数据库的访问权限即可查看对应hash的salt。更为重要是防止轻易生成包含其password的rainbow table。比如，有些开发者通常会将username作为salt，username+password的组合并不是high-entropy password(比如经常出现的admin，root等用户名)，这种形式的加盐并未实现加盐的目的。
>  
>  In addition, high-entropy password本身就没有必要进行salted hash，high-entropy password得到的hash也是high-entropy hash。
* Password hashing on frontend or backend?
> 讨论应该在前端对password做hash还是在后端做？ 一般情况下，在后端对password做hash并存入数据库。
> 
> 网上出现的在前端做hash的原因
> 1. 由于慢hash等原因，希望减少服务器端的压力；
> 2. 用户不信任客户端，不想展示给客户端自己保存的真实密码，由于自己的密码同时用在其他地方；
> 
>针对第一点，如果只在前端做hash，前端被hash处理的password等同于原先的明文“密码”，明文“密码”无任何处理存入数据库，失去了hash处理密码的意义。当数据库被攻破，攻击者可以直接利用密码的hash值登录用户账户；亦或是做一些部分准备工作，hash的前置处理，但是JS处理这些工作性能很差，无法给予有效的帮助
> 
> 针对第二点，用户不信任服务端是一个伪命题，如果服务端不受信任，服务端发送的处理hash的JS代码也不可被信任。替代方法是用户使用一个master password，针对不同的网站，利用网站域名做为salt为某个网站生成特定的密码，或者为每个网站生成一个high-entropy密码。
### Other threat
与hash不相关的其他威胁
* 担心客户端信息被攻击者拦截，发起replay attack。
> Use nonce to prevent from replay attack.
> 加入timestamp从而只允许在时间戳浮动范围几秒内的请求的正常访问；使用序列码，每一条消息被赋予唯一的序列数值，下次请求服务端确认请求是否重复，是否为服务端期待的下一次请求。核心就是标识出不可重复出现的唯一请求，并且希望该参数不可轻易被攻击者获取。
* TLS中的明文密码传输
> password在TLS中明文传输没有什么问题，传输保证了安全性。可能会出现的意外包括中间人攻击？通过mutual authentication解决CA证书不可靠的问题。
### Other technologies
* SRP(Secure Remote Password protocol)
* Man-in-the-middle attack (MITM) 在客户端不可信证书的情况下会出现，公钥不可靠的问题
* HMAC，可以使用共享密钥提供身份验证，而不是使用非对称加密的数字签名。
* SHA-1 collision, etc. why does collisions cause security threat?

# reference
[1]  [should password be hashed server-side or client-side?](https://security.stackexchange.com/questions/8596/https-security-should-password-be-hashed-server-side-or-client-side/23285#23285)
[2]  [How to hash passwords and when not to](https://advancedweb.hu/how-to-hash-passwords-and-when-not-to/)
[3]  [rainbow table](https://en.wikipedia.org/wiki/Rainbow_table#:~:text=A%20rainbow%20table%20is%20a,form%2C%20but%20as%20hash%20values.)
[4]  [Does it make security sense to hash password on client end](https://stackoverflow.com/questions/1380168/does-it-make-security-sense-to-hash-password-on-client-end?rq=1/#1380199)
[5]  [mutual authentication](https://security.stackexchange.com/questions/88805/does-mutual-authentication-have-any-impact-on-mitm-possibilities)
[6]  [HMAC](https://security.stackexchange.com/questions/20129/how-and-when-do-i-use-hmac)
[7]  [Collision]()(https://www.youtube.com/watch?v=Zl1TZJGfvPo)