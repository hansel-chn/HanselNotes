package main

import "fmt"

/*
Given a string s, remove duplicate letters so that every letter appears once and only once. You must make sure your result is
the smallest in lexicographical order
 among all possible results.
*/

func removeDuplicateLetters(s string) string {
	sMap := make(map[byte][]int)
	for i := 0; i < len(s); i++ {
		if v, ok := sMap[s[i]]; ok {
			v[0]++
		} else {
			sMap[s[i]] = []int{1, 0}
		}
	}
	fmt.Println(sMap)
	monotonicStack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if sMap[s[i]][1] == 1 {
			sMap[s[i]][0]--
			continue
		}
		if 0 == i || monotonicStack[len(monotonicStack)-1] < s[i] {
			monotonicStack = append(monotonicStack, s[i])
			sMap[s[i]][0]--
			sMap[s[i]][1] = 1
		} else {
			for len(monotonicStack) > 0 && monotonicStack[len(monotonicStack)-1] >= s[i] && sMap[monotonicStack[len(monotonicStack)-1]][0] > 0 {
				sMap[monotonicStack[len(monotonicStack)-1]][1] = 0
				monotonicStack = monotonicStack[:len(monotonicStack)-1]
			}
			monotonicStack = append(monotonicStack, s[i])
			sMap[s[i]][0]--
			sMap[s[i]][1] = 1
		}
	}
	return string(monotonicStack)
}

func main() {
	//s := "cbacdcbc"
	//fmt.Println(removeDuplicateLetters(s))
	s1 := "bbcaac"
	fmt.Println(removeDuplicateLetters(s1))
}
