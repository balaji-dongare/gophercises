package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/balaji-dongare/gophercises/CLI/task/cmd"
	"github.com/balaji-dongare/gophercises/CLI/task/dbrepository"
)

func main() {
	dir, _ := os.Getwd()
	databasepath := filepath.Join(dir, "tasks.db")
	if err := dbrepository.InitDatabase(databasepath); err != nil {
		fmt.Println(err)
	}
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
