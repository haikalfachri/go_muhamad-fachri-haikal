package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      struct { 
		Rate  float64 `json:"rate"`
		Count int     `json:"count"`
	} `json:"rating"`
}

func main() {
	var url string = "https://fakestoreapi.com/products"
	var productCh = make(chan Product)
	var products []Product
	var resp *http.Response
	var err error

	go func(){
		resp, err = http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		
		err = json.NewDecoder(resp.Body).Decode(&products)
		if err != nil {
			panic(err)
		}
		
		for _, p := range products {
			productCh <- p
		}
		defer close(productCh)
	}()
	
	fmt.Println("products data")
	for p := range productCh {
		fmt.Println("====")
		fmt.Println("title: ", p.Title)
		fmt.Println("price: ", p.Price)
		fmt.Println("category: ", p.Category)
		fmt.Println("====")
	}
}
