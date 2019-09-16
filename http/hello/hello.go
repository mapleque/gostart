package main

import (
	"fmt"
	"log"
	"net/http"
)

// RunHelloService run hello service as an example of gostart/http
//
// Command:
//  `go run github.com/mapleque/gostart/http/hello`
//
// Run in console:
//     curl http://localhost/hello?name=cookie
//     #> Hello, cookie，see you 1 times.
//     curl http://localhost/hello?name=cookie
//     #> Hello, cookie，see you 2 times.
//     curl http://localhost/hello?name=cookie
//     #> Hello, cookie，see you 3 times.
func RunHelloService(addr string) {
	storage := map[string]int{}
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if _, exist := storage[name]; !exist {
			storage[name] = 0
		}
		storage[name]++
		fmt.Fprintf(w, "Hello, %s，see you %d times.", name, storage[name])
	})
	fmt.Printf("service is running on %s.\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}

func main() {
	RunHelloService(":80")
}
