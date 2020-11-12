package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// ===================== get

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"time"
// )

// func main() {
// 	resp, err := http.Get("https://www.google.com")
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	defer resp.Body.Close()

// 	b, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	log.Println(string(b))
// }

// func main() {
// 	urls := []string{
// 		"https://www.google.com",
// 		"https://www.yahoo.com",
// 		"https://www.bing.com",
// 		"https://www.ecosia.com",
// 	}

// 	for i := 0; i < len(urls); i++ {
// 		resp, err := http.Get(urls[i])
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}
// 		defer resp.Body.Close()

// 		b, err := ioutil.ReadAll(resp.Body)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(1)
// 		}

// 		fmt.Println(len(b))
// 	}
// }

// ===================== commandLineArgs

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintln(os.Stderr, "Cần ít nhất một input")
// 		os.Exit(1)
// 	}

// 	for i, arg := range os.Args[1:] {
// 		fmt.Printf("[%d] - %q\n", i, arg)
// 	}
// }

// ===================== fetchURLConcurrent

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"os"
// 	"sync"
// 	"time"
// )

// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Fprintln(os.Stderr, "Cần ít nhất một input")
// 		os.Exit(1)
// 	}

// 	var wg sync.WaitGroup
// 	wg.Add(len(os.Args[1:]))

// 	fmt.Println("Đang tải dữ liệu (xử lí đồng thời)")
// 	concurrentStart := time.Now()
// 	for _, arg := range os.Args[1:] {
// 		go fetchURL(arg, &wg)
// 	}

// 	wg.Wait()
// 	fmt.Printf("Hoàn thành trong %v\n", time.Since(concurrentStart))
// }

// func fetchURL(url string, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	response, err := http.Get(url)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "--- Không yêu cầu được thông tin từ %q: %s\n", url, err.Error())
// 		return
// 	}
// 	defer response.Body.Close()

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "--- Không đọc được thông tin từ %q: %s\n", url, err.Error())
// 		return
// 	}

// 	fmt.Printf("--- Nhận được %dB từ %q (%q)\n", len(body), url, response.Header.Get("Content-Type"))
// }

// ===================== fetchURLSequential

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Cần ít nhất một input")
		os.Exit(1)
	}

	fmt.Println("Đang tải dữ liệu (xử lí theo trình tự)")
	sequentialStart := time.Now()
	for _, arg := range os.Args[1:] {
		fetchURL(arg)
	}
	fmt.Printf("Hoàn thành trong %v\n\n\n", time.Since(sequentialStart))

}

func fetchURL(url string) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "--- Không yêu cầu được thông tin từ %q: %s\n", url, err.Error())
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "--- Không đọc được thông tin từ %q: %s\n", url, err.Error())
		return
	}

	fmt.Printf("--- Nhận được %dB từ %q (%q)\n", len(body), url, response.Header.Get("Content-Type"))
}