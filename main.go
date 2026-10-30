package main

import (
	"fmt"
	"log"
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
		log.Println("Слайс не может быть пустым")
		return nil
	}

	elements := make([]int, size)
	for i := range elements {
		elements[i] = rand.Int()
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		log.Println("Слайс не может быть пустым")
		return 0
	}

	maxInt := data[0]
	for _, num := range data {
		if num > maxInt {
			maxInt = num
		}
	}

	return maxInt
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	dataSize := len(data)
	if dataSize == 0 {
		log.Println("Слайс не может быть пустым")
		return 0
	}

	chunks := CHUNKS
	if dataSize < chunks {
		chunks = dataSize
	}

	maxInts := make([]int, chunks)
	chunkSize := dataSize / chunks
	remainingElements := dataSize % chunks
	var wg sync.WaitGroup
	start := 0

	for i := 0; i < chunks; i++ {
		end := start + chunkSize
		if i < remainingElements {
			end++
		}
		chunk := data[start:end]
		start = end

		wg.Add(1)
		go func(i int, chunk []int) {
			defer wg.Done()
			maxInts[i] = maximum(chunk)
		}(i, chunk)
	}

	wg.Wait()

	return maximum(maxInts)
}

func main() {
	// generating data slice
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	timeStart := time.Now()
	data := generateRandomElements(SIZE)
	elapsed := time.Since(timeStart)
	fmt.Printf("Время генерации: %d ms\n", elapsed.Milliseconds())

	//finding max value with 1 thread
	fmt.Println("Ищем максимальное значение в один поток")
	timeStart = time.Now()
	max := maximum(data)
	elapsed = time.Since(timeStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())

	//finding max value with 8 threads
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	timeStart = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(timeStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed.Milliseconds())
}
