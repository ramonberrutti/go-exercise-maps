package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	split := strings.Fields(s)

	for _,v := range split {
		m[v] += 1;
	}

	return m;
}

func main() {
	wc.Test(WordCount)
}
