# md语法

文字引用，表示层级更深，使用`>`，可以有多个`>`
> 第一层
>> 第二层
> 111111

# 变量类型

1. > `string`转`int`，`int`转`string`  
   > string转成int：   
   `int, err := strconv.Aoti(string)`  
   > string转成int64：  
   `int64, err := strconv.ParseInt(string, 10, 64)`  
   > int转成string：  
   `string := strconv.Itoa(int)`  
   > int64转成string：  
   `string := strconv.FormatInt(int64, 10)`

3. > byte 是 uint8的别名  
   > 一个byte长度为8，即八位一个字节  
   > 一个byte等于八个bit  
   > 一个bit表示一位  
   > 所以八位一个字节 8bit == byte  
   > byte        alias for uint8  
   > rune        alias for int32  
   > uint8 可以表示一个字符,所以[]unit8切片可以表示一个字符串  
   > 数据库不同  
   > int  11是表示宽度  
   > varchar 255 表示255个字符长度  