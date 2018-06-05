package main

import (
	"currency-quote-api/scraping"
	"fmt"
)

func main() {
	goGetCurrrency()
}

func goGetCurrrency() {
	curr := []scraping.Currency{}
	done := make(chan bool)
	go scraping.GetAllCurrency(&curr, done)
	<-done
	fmt.Println(curr)
}
