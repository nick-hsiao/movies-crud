package main

import (
	_ "encoding/json"
	"fmt"
	_ "log"
	_ "math/rand"
	_ "net/http"
	_ "strconv"

	"rsc.io/quote"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
}

func main() {
	fmt.Println(quote.Go())
}
