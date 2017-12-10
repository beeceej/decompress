package main

import (
	"os"

	"github.com/beeceej/decompress"
)

func main() {
	d := decompress.TGZDecompress{}

	d.Decompress(os.Args[1], os.Args[2])
}
