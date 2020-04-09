package main

import (
	"archive/zip"
	"io"
	"log"
	"os"
	"path/filepath"
)

// ARŞİVLENMİŞ DOSYALARI DIŞARI ÇIKARMAK
func main() {

	zr, err := zip.OpenReader("deneme.zip")
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()

	for _, f := range zr.Reader.File {
		zippedFile, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedFile.Close()

		targetDir := "./"
		extractedFilePath := filepath.Join(
			targetDir,
			f.Name,
		)
		if f.FileInfo().IsDir() {
			log.Println("Klasör Oluşturuluyor!", extractedFilePath)
			os.MkdirAll(extractedFilePath, f.Mode())
		} else {
			log.Println("Dosya Çıkarılıyor", f.Name)
			outFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				f.Mode(),
			)
			if err != nil {
				println(err)
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, zippedFile)
			if err != nil {
				println(err)
			}

		}

	}
}
