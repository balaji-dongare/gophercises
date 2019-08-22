package main

import (
	"os"
	"path/filepath"

	"github.com/Bala-G/gophercises/CLI/task/cmd"
	"github.com/Bala-G/gophercises/CLI/task/dbrepository"
)

func main() {
	dir, _ := os.Getwd()
	databasepath := filepath.Join(dir, "tasks.db")
	if err := dbrepository.InitDatabase(databasepath); err != nil {
		panic(err)
	}
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
