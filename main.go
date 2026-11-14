package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	randomSlice := make([]int, size)
	for i := range randomSlice {
		randomSlice[i] = rand.Int()
	}
	return randomSlice
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maximum := data[0]
	for _, v := range data {
		if v > maximum {
			maximum = v
		}
	}
	return maximum
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	maxSlice := make([]int, CHUNKS)
	chunkSize := len(data) / CHUNKS
	var wg sync.WaitGroup
	for i := range CHUNKS {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS {
			end = len(data)
		}
		wg.Add(1)
		go func(index int, start int, end int) {
			defer wg.Done()
			maxSlice[index] = maximum(data[start:end])
		}(i, start, end)
	}

	wg.Wait()
	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел", SIZE)
	randomSlice := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(randomSlice)
	elapsed := time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(randomSlice)
	elapsed = time.Since(start)

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
