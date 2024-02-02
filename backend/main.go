package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"log"
	"strconv"
	"encoding/json"
)

// allowCORS() - need to handle CORS Middleware to allow requests from my Svelte app
func allowCORS(next httprouter.Handle) httprouter.Handle {
    return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173") // Port could be different!
        w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS") // Only ones needed for this app
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

        // Pre-flight test to make sure the request is safe to send
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next(w, r, ps) // Continues the chain, but in this app there's only one middleware function
    }
}

// fibonacciHandler() handles the API requests
func fibonacciHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    n, err := checkNumber(ps.ByName("length"))
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    fibSequence := fibonacci(n) // The fibonacci sequence for 'n'

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(fibSequence) // Return as JSON
}

// fibonacci() returns 'n' elements of the Fibonacci sequence
func fibonacci(n int) []int {
    // Base cases
	if n <= 0 {
        return []int{}
    }
    if n == 1 {
        return []int{0}
    }

    s := make([]int, n) // initializing sequence...
	s[0], s[1] = 0, 1 // base cases
    for i := 2; i < n; i++ { // generating rest of sequence up to 'n'
        s[i] = s[i-1] + s[i-2]
    }

    return s // (fibonacci sequence)
}

// Ensures the request number is an integer and between 0-100 
func checkNumber(number string) (int, error) {
    n, err := strconv.Atoi(number)
    if err != nil {
        return 0, fmt.Errorf("Must be an integer, try again.") // 0 as a return value is arbitrary
    }
    if n < 0 || n > 100 {
        return 0, fmt.Errorf("The number must be between 0-100, try again.")
    }
    return n, nil
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome to Jack's Fibonacci Sequence API Server!\n")
}

func main() {
    router := httprouter.New()
    
	router.GET("/api/fibonacci/", Index) // No argument
	router.GET("/api/fibonacci/:length", allowCORS(fibonacciHandler)) // With argument
    
	log.Fatal(http.ListenAndServe(":8080", router))
}