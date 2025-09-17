package main

import "fmt"

func mergeAlternately(word1 string, word2 string) (ret string) {
	p, q := 0, 0
	for p < len(word1) || q < len(word2) {
		if p < len(word1) {
			fmt.Println("word1", ret)
			ret, p = ret+string(word1[p]), p+1
			fmt.Println("word2", ret, "p:", p)

		}

		if q < len(word2) {
			fmt.Println("word2", ret)
			ret, q = ret+string(word2[q]), q+1
			fmt.Println("word2", ret, "p:", p)
		}

	}
	return
}

func main() {
	res := mergeAlternately("abc", "efg")
	fmt.Println(res)
}
