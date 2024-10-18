package zipping

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func GoZip(source string, tagret string) error {

	zipFile, err := os.Create(tagret)
    if err != nil {
        return err
    }
    defer zipFile.Close()

    writer := zip.NewWriter(zipFile)
    defer writer.Close()

    err = filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }

        if info.IsDir() {
            return nil 
        }

        relPath, err := filepath.Rel(source, path)
        if err != nil {
            return err
        }

        zipEntry, err := writer.Create(relPath)
        if err != nil {
            return err
        }

        file, err := os.Open(path)
        if err != nil {
            return err
        }
        defer file.Close()

        _, err = io.Copy(zipEntry, file)
        return err
    })

    return err

}
