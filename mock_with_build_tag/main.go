package main

import (
	"fmt"
	"my/go-my-sandbox/mock_with_build_tag/client"
)

func main() {
	c := client.NewClient()
	fmt.Println(c.Tag())
}
