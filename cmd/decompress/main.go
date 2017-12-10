package main

import "github.com/beeceej/easy-tar-gz"
import "os"

func main() {
	d := decompress.TGZDecompress{}

	d.Decompress(os.Args[1], os.Args[2])
}
