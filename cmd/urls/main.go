package main

import (
	"fmt"
	"work/URLs/internal/config"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
}
