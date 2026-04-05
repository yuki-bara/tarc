// SPDX-License-Identifier: 0BSD
// Author: Makkhawan Sardlah
package tarc

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func createtar(dir string, tw *tar.Writer) error {

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(dir, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(tw, file)
			return err
		}

		return nil
	})
	return err
}

func createtargz(dir string, out *os.File) error {

	gw := gzip.NewWriter(out)
	defer gw.Close()

	tw := tar.NewWriter(gw)
	defer tw.Close()

	return createtar(dir, tw)
}

func Compressfile(dir string, destination string, tarball string) error {
	out, err := os.Create(destination)
	if err != nil {
		return err
	}
	defer out.Close()
	switch tarball {
	case "*":
		tw := tar.NewWriter(out)
		defer tw.Close()
		return createtar(dir, tw)
	case "GZ":
		return createtargz(dir, out)
	}
	return nil
}
