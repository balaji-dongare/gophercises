package dbrepository

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// initdb initalize db for test environment
func initdb() {
	dir, _ := os.Getwd()
	databasepath := filepath.Join(dir, "tasks.db")
	InitDatabase(databasepath)
}

// TestInitDatabase testcase for initalize database function
func TestInitDatabase(t *testing.T) {
	databasepath := filepath.Join("github.com\balaji-dongare\\gophercises\\CLI\task\\dbrepository\tasks.db")
	fmt.Println(databasepath)
	if err := InitDatabase(databasepath); err == nil {
		t.Errorf("error:%v", err)
	}
}

// TestInsertTaskIntoDB testcase for insert task
func TestInsertTaskIntoDB(t *testing.T) {
	initdb()
	_, err := InsertTaskIntoDB("Testinggophercises")
	if err != nil {
		t.Errorf("Unable to Insert Task into todo list")
	}

	Status, _ := InsertTaskIntoDB("Testing  CLI gophercises")
	if Status {
		t.Logf("Inserted Task into todo list")
	}

}

// TestReadTodosTaskFromDB testcase for read todo tasks from db
func TestReadTodosTaskFromDB(t *testing.T) {
	initdb()
	_, err := ReadTodosTaskFromDB()
	if err != nil {
		t.Errorf("Unable to read the Task from DB")
	}
}

// MarkTaskAsDone testcase for mark todo tasks as done
func TestMarkTaskAsDone(t *testing.T) {
	initdb()
	var ids = []int{1, 2}
	_, tasks, err := MarkTaskAsDone(ids)
	if err != nil {
		t.Error("error in doing task")
	}
	if len(tasks) < 0 {
		t.Logf("Not a single task mark as done")
	} else {
		fmt.Printf("Following Task Completed:\n")
		for i, task := range tasks {
			t.Logf("%d. %v\n", i+1, task)
		}
	}

	ids = []int{1, 2, 3}
	_, tasks, err = MarkTaskAsDone(ids)
	if err != nil {
		t.Error("error in doing task")
	}
}

func TestInitDatabaseError(t *testing.T) {
	initdb := initDB
	defer func() {
		initDB = initdb
	}()
	initDB = func(driverName, dataSourceName string) (*sql.DB, error) {
		return nil, errors.New("Error go in DB Open")
	}
	if err := InitDatabase("databasepath"); err == nil {
		t.Errorf("error:%v", err)
	}
}
