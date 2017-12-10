package decompress

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Decompressor provides an inteface to decompressing methods
type Decompressor interface {
	Decompress(in, out string) error
}

// TGZDecompress decompress .tar.gz
type TGZDecompress struct {
	Verbose bool
}

// Decompress performs the decompression
func (tg *TGZDecompress) Decompress(in, out string) (err error) {
	var f *os.File
	if f, err = os.Open(in); err != nil {
		return err
	}
	defer f.Close()
	return tg.untar(out, f)
}

// thanks to:
// https://medium.com/@skdomino/thanks-for-bringing-this-up-696512dc93dc
// Steve Domino
func (tg *TGZDecompress) untar(out string, r io.Reader) (err error) {
	var gzr *gzip.Reader

	if gzr, err = gzip.NewReader(r); err != nil {
		return err
	}
	defer gzr.Close()
	tr := tar.NewReader(gzr)
	return tg.unpack(tr, out)
}

func (tg *TGZDecompress) unpack(tr *tar.Reader, out string) error {
	var (
		header *tar.Header
		err    error
	)
	for {
		header, err = tr.Next()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		case header == nil:
			continue
		}
		target := filepath.Join(out, header.Name)
		if tg.Verbose {
			fmt.Printf("Unpacking %s", target)
		}
		switch header.Typeflag {
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err = os.MkdirAll(target, 0755); err != nil {
					return err
				}
			}
		case tar.TypeReg:
			err = func() (err error) {
				var f *os.File
				if f, err = os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode)); err != nil {
					return err
				}
				defer f.Close()
				_, err = io.Copy(f, tr)
				return err
			}()
			if err != nil {
				return err
			}
		case tar.TypeSymlink:
			// unhandled
		}
	}
}
