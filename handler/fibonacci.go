package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func GetFibonacci() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nStr := r.URL.Query().Get("n")
		if nStr == "" {
			http.Error(w, "Missing 'n' parameter", http.StatusBadRequest)
			return
		}

		n, err := strconv.Atoi(nStr)
		if err != nil || n < 0 {
			http.Error(w, "'n' must be a positive integer", http.StatusBadRequest)
			return
		}

		var wg sync.WaitGroup
		results := make(chan int, 10)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				results <- fibonacci(n)
			}()
		}

		go func() {
			wg.Wait()
			close(results)
		}()

		total := 0
		for result := range results {
			total += result
		}

		fmt.Fprintf(w, "Fibonacci(%d) = %d", n, total)
	}
}
