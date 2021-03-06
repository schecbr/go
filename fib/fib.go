package main

import (
  "big"
  "fmt"
	"http"
  "strconv"
)

var memoize_fibonnaci = map[int]*big.Int { 0: big.NewInt(0), 1: big.NewInt(1), 2: big.NewInt(1) }

func fibonnaci(n int) *big.Int {
  fib_term := new(big.Int)
  if n <= (len(memoize_fibonnaci) - 1) {
    return memoize_fibonnaci[n]
  }
  fib_term.Add(fibonnaci(n-2), fibonnaci(n-1))
  memoize_fibonnaci[n] = fib_term
  return fib_term;
}

func FibHandler(writer http.ResponseWriter, request *http.Request) {
  nth_fib, _ := strconv.Atoi(request.URL.Path[1:]) 
  fib_value := fibonnaci(nth_fib)
  fmt.Fprintf(writer, "The %vth fibonnaci number is is %v", nth_fib, fib_value.String())
}

func main() {
  http.HandleFunc("/", FibHandler)
  http.ListenAndServe(":8080", nil)
}