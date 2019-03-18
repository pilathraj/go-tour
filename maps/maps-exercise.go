package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	s1 := strings.Split(s, " ")
	m := make(map[string]int)
	for _, v := range s1 {
		cnt, ok := m[v]
		if ok {
			cnt = cnt + 1
		} else {
			cnt = 1
		}
		m[v] = cnt
	}
	return m
}

func wordCount2(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

func wordCountinfo() {
	//wc.Test(WordCount)
	fmt.Println(wordCount("I am learning Go"))
	fmt.Println(wordCount2("I am learning Go"))
}
