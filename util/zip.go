package util

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func ZipDir(dir, zipFile string) {

	fz, err := os.Create(zipFile)
	if err != nil {
		panic("Create zip file failed: " + err.Error())
	}
	defer fz.Close()

	w := zip.NewWriter(fz)
	defer w.Close()

	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			// FIX: 存在 ./ 情况 先-2 以后换成其他方式
			fDest, err := w.Create(path[len(dir)-2:])
			if err != nil {
				panic("Create failed: " + err.Error())
				return nil
			}
			fSrc, err := os.Open(path)
			if err != nil {
				panic("Open failed: " + err.Error())
				return nil
			}
			defer fSrc.Close()
			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				panic("Copy failed: " + err.Error())
				return nil
			}
		}
		return nil
	})
}

func UnzipDir(zipFile, dir string) {

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		panic("UnzipDir zip file failed: " + err.Error())
	}
	defer r.Close()

	for _, f := range r.File {
		func() {
			path := dir + string(filepath.Separator) + f.Name
			os.MkdirAll(filepath.Dir(path), 0755)
			fDest, err := os.Create(path)
			if err != nil {
				panic("Create failed: " + err.Error())
				return
			}
			defer fDest.Close()

			fSrc, err := f.Open()
			if err != nil {
				panic("Open failed: " + err.Error())
				return
			}
			defer fSrc.Close()

			_, err = io.Copy(fDest, fSrc)
			if err != nil {
				panic("Copy failed: " + err.Error())
				return
			}
		}()
	}
}
