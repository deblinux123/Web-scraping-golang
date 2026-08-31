package main

import (
	"fmt"
	"net/url"
)

func main() {
	u, err := url.Parse("https://dummyjson.com/products")

	if err != nil {
		panic(err)
	}

	q := u.Query()

	q.Set("limit", "20")
	q.Set("skip", "40")

	u.RawQuery = q.Encode()

	fmt.Println(u.String())
}
