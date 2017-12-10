package main

import (
	"os"

	"github.com/beeceej/decompress"
)

func main() {
	new(decompress.TGZDecompress).
		Decompress(os.Args[1], os.Args[2])
}
