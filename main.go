package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
)

func init() {

	infoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func fibonacci(n int) []int {
	seq := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 || i == 1 {
			seq[i] = 1
		} else {
			seq[i] = seq[i-1] + seq[i-2]
		}
	}
	return seq
}

func fibonacciHandler(w http.ResponseWriter, r *http.Request) {
	nStr := r.URL.Query().Get("n")
	if nStr == "" {
		http.Error(w, "Missing 'n' query parameter", http.StatusBadRequest)
		errorLogger.Println("Missing 'n' parameter")
		return
	}

	n, err := strconv.Atoi(nStr)
	if err != nil || n <= 0 {
		http.Error(w, "Invalid 'n' parameter", http.StatusBadRequest)
		errorLogger.Printf("Invalid 'n' value: %v", nStr)
		return
	}

	infoLogger.Printf("Generating Fibonacci sequence of %d numbers", n)
	result := fibonacci(n)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func main() {
	http.HandleFunc("/fibonacci", fibonacciHandler)
	infoLogger.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		errorLogger.Fatalf("Server failed to start: %v", err)
	}
}
