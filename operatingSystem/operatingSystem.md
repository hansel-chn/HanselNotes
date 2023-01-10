## 地址总线，数据总线

1. 计算机字长取决于计算机字长取决于哪种总线宽度?数据总线/地址总线?为什么？

   答1:取决于数据总线，宽度！代表一次可以取多少的数据！地址总线，相当于取数据的地址范围！ 可以这样比方，地址总线，相当于旅馆里的房号。 而数据总线，相当于一间房的大小！（不完全对，见下方）
2. 
    [例子](https://blog.csdn.net/weixin_39923623/article/details/119069555?utm_medium=distribute.pc_relevant.none-task-blog-2~default~baidujs_baidulandingword~default-4-119069555-blog-80732236.pc_relevant_recovery_v2&spm=1001.2101.3001.4242.3&utm_relevant_index=7)
3. 计算机的“字长”---俗称是多少位的CPU 

   这里上下文特指计算机的字长，它就是CPU里寄存器的宽度，学了计算机组成应该知道，CPU里除了控制器、运算器之外，还有很多寄存器（当然现在CPU还有Cache等），机器的字长就是其中通用寄存器（GPR）的位宽。
    
   CPU在设计的时候会让运算器和通用寄存器的位宽保持一致，在硬件上就是CPU一次能进行的多少位数据运算，所以字长反应了CPU的数据运算能力（一次可以进行几位的数据运算，得到的结果是几位的）。
4. 1. 32位和64位指的是cpu一次能处理的数据的长度（也就是寄存器的位数），源头的定义和数据总线和地址总线都没有关系2. 如果数据总线的长度小于字长的话，那么会浪费cpu的处理能力，大于字长的话，传动过来的数据cpu一次处理不完，所以一般数据总线的长度等于字长3. 指针也是数据，所以cpu一次处理的数据长度和指针长度最好是相等的，而指针的长度和地址总线又是对应的，所以地址总线的长度一般情况也会等于字长综上所述：字长==数据总线长度==地址总线长度（但是实际中地址总线很多高位用不到，所以默认置0，省去了一些地址总线，总的长度是小于64的，只是64位的机器理论上可寻址的范围在2的64次方B)
5. [总结1](https://www.bilibili.com/read/cv19325810)
   [总结2](https://blog.csdn.net/weixin_29440125/article/details/118988698)
   [题目](https://blog.csdn.net/jingjingshizhu/article/details/117289697?spm=1001.2101.3001.6650.1&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-117289697-blog-80732236.pc_relevant_recovery_v2&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-1-117289697-blog-80732236.pc_relevant_recovery_v2&utm_relevant_index=2)
6. 2.数据总线的宽度与字长及CPU位数

字长指CPU同一时间内可以处理的二进制数的位数，数据总线传输的数据或指令的位数要与字长一致。否则，如果数据总线宽度大于字长则一条数据或指令要分多次传输，则分开传输的几组数据也就没有意义了；如果数据总线宽度小于字长，则CPU的利用率要降低，对资源是种浪费。

另外，如果字长为n位，一般称CPU是n位的。所以说数据总线的宽度与字长及CPU的位数是一致的。