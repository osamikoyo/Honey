package search

import (
	"io"
	"os"
	"strings"
)

func Search(value string, path string) (string, error){
	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	var result []byte
	if _, err := io.ReadFull(file, result); err != nil{
		return "", err
	}

	logs := strings.Split(string(result), "]]]")

	for _, log := range logs{
		themes := strings.Fields(log)
		for _, them := range themes{
			if them == value{
				return log, nil
			}
		}
	}

	return "", nil
}