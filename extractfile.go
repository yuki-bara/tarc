package tarc

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func extracttar(tr *tar.Reader, destination string) error {
	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		target := filepath.Join(destination, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(target, 0755); err != nil {
				return err
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(target), 0755); err != nil {
				return err
			}

			outFile, err := os.Create(target)
			if err != nil {
				return err
			}

			if _, err := io.Copy(outFile, tr); err != nil {
				outFile.Close()
				return err
			}
			outFile.Close()
		}
	}
	return nil
}

func extracttargz(f *os.File, destination string) error {
	gr, err := gzip.NewReader(f)
	if err != nil {
		return err
	}
	defer gr.Close()
	tr := tar.NewReader(gr)
	return extracttar(tr, destination)
}

func Extractfile(file string, destination string, tarball string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	switch tarball {
	case "*":
		tr := tar.NewReader(f)
		return extracttar(tr, destination)
	case "GZ":
		return extracttargz(f, destination)
	}
	return err
}
