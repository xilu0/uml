package main

import (
	"fmt"
	"strings"
	"sync"
)

// Map函数
func Map(word string, ch chan<- map[string]int) {
	frequency := make(map[string]int)
	frequency[word] = 1
	ch <- frequency
}

// Reduce函数
func Reduce(frequencies []map[string]int) map[string]int {
	result := make(map[string]int)
	for _, freq := range frequencies {
		for word, count := range freq {
			result[word] += count
		}
	}
	return result
}

func main() {
	documents := []string{"apple banana", "apple orange", "banana orange", "banana"}
	ch := make(chan map[string]int, len(documents))

	var wg sync.WaitGroup
	wg.Add(len(documents))

	// 并发执行Map
	for _, doc := range documents {
		go func(doc string) {
			defer wg.Done()
			words := strings.Fields(doc)
			for _, word := range words {
				Map(word, ch)
			}
		}(doc)
	}

	// 收集Map结果
	var frequencies []map[string]int
	for freq := range ch {
		frequencies = append(frequencies, freq)
	}
	// 等待Map阶段完成
	wg.Wait()
	close(ch)
	// 执行Reduce
	result := Reduce(frequencies)
	fmt.Println(result)
}
