package main

import (
	"fmt"

	"github.com/dennnniska/shortUrl/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Print(cfg)
}
