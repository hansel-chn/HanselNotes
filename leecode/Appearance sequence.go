package main

import (
	"strconv"
	"strings"
)

/*
这样使用
cur.WriteString(strconv.Itoa(end - start))
cur.WriteByte(prev[start])
重复了两遍，因为边界最后虽然判断结束但是没有执行最后一次加入字符串的结果

官方题解使用了三次循环，但是其实是一样的，不额外重复语句是因为通过最内层循环进行了长度判断，出循环仍可以进行添加字符。通过两个for把语句和在一起了，其
实就是用while解决
for j < len(prev) && prev[j] == prev[start]

？错误方法，只判断start=end，在最后一个数没有判断长度，会越界且不停止，括号里写的方法没有意义（还有一种是把for当while，start和end不相等则继续
向前拓展字符串。相等的话跳出循环（其实就是当end遇到不同字符串的时候，cur添加字符，start=end，跳出循环）。外层end+1，for继续找下一个相同字符）
*/
func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := strings.Builder{}
		start, end := 0, 0
		for end < len(prev) {
			if prev[start] == prev[end] { // 这里区别导致还得再写一遍，在一个条件的基础下做另一个事情。自己写的就是两个条件出来的地方不一样
				end++
			} else {
				cur.WriteString(strconv.Itoa(end - start))
				cur.WriteByte(prev[start])
				start = end
			}
		}
		cur.WriteString(strconv.Itoa(end - start))
		cur.WriteByte(prev[start])
		prev = cur.String()
	}
	return prev
}

// 官方题解
func countAndSayLeeCode(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] { // here 满足这两个条件了做这个事
				j++
			}
			cur.WriteString(strconv.Itoa(j - start))
			cur.WriteByte(prev[start])
		}
		prev = cur.String()
	}
	return prev
}

func main() {

}
