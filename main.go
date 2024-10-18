package main

import (
	"os"

	"golog/loger"
)

func main() {
	lg := cmd.New_OneLogger(os.Stdout)
	_, err := os.Open("ksdgsdgsgdsgd")

	if err != nil {
		lg.Error(err)
	}
}
