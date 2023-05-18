# Transport Layer

## TCP

[https://xiaolincoding.com/network/3_tcp/tcp_feature.html#%E6%8B%A5%E5%A1%9E%E6%8E%A7%E5%88%B6](https://xiaolincoding.com/network/3_tcp/tcp_feature.html#%E6%8B%A5%E5%A1%9E%E6%8E%A7%E5%88%B6)

### 超时重传

### 为什么端口号最大只能是65535?

其中Source Port和Destination Port都占2字节，也就是16位，所以最大能表示的端口号是2^16=65536，由于0号端口有特殊用途不能使用，实际最大端口号是65536-1 =
65535。所以，当别人问到为什么端口号最大只能是65535？你就知道了，这个并不是操作系统的限制，而是TCP协议的限制。
