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
	var elements []int

	if size < 2 {
		log.Println("Размер слайса для генерации случайных чисел не может быть меньше 2")
		return nil
	}

	for i := 0; i < size; i++ {
		randInt := rand.Intn(1000)
		elements = append(elements, randInt)
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	var maxInt int

	if len(data) < 2 {
		log.Println("Размер слайса для получения элемента с максимальным значением не может быть меньше 2")
		return 0
	}

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
	maxInts := make([]int, CHUNKS)
	var wg sync.WaitGroup

	if dataSize < 2 {
		log.Println("Размер слайса для получения элемента с максимальным значением не может быть меньше 2")
		return 0
	}

	for i := 0; i < CHUNKS; i++ {
		start := i * dataSize / CHUNKS
		end := (i + 1) * dataSize / CHUNKS

		chunk := data[start:end]

		wg.Add(1)
		go func(i int, chunk []int) {
			defer wg.Done()

			max := chunk[0]
			for _, n := range chunk {
				if n > max {
					max = n
				}
			}

			maxInts[i] = max
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
	fmt.Printf("Время генерации: %d ms\n", elapsed)

	//finding max value with 1 thread
	fmt.Println("Ищем максимальное значение в один поток")
	timeStart = time.Now()
	max := maximum(data)
	elapsed = time.Since(timeStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)

	//finding max value with 8 threads
	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	timeStart = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(timeStart)
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max, elapsed)
}
