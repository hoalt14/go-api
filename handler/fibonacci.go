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
		result := 0
		mu := sync.Mutex{}

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				fib := fibonacci(n)
				mu.Lock()
				result = fib
				mu.Unlock()
			}()
		}

		wg.Wait()

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Fibonacci(%d) = %d", n, result)
	}
}
