package dbrepository

import (
	"database/sql"
	"fmt"

	// driver package
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var initDB = sql.Open

// InitDatabase function used to initialized the sqlite3 db
func InitDatabase(dbPath string) error {
	var err error

	db, err = initDB("sqlite3", dbPath)
	if err != nil {
		fmt.Printf("Error While Connection:%v", err)
		return err
	}
	_, err = db.Exec("create table if not exists tasks(task text primary key)")
	if err != nil {
		fmt.Printf("Error while creating task table:%v", err)
		return err
	}
	return nil
}

// InsertTaskIntoDB function used to insert the task into db
// and return status and error
func InsertTaskIntoDB(task string) (bool, error) {
	stmt, _ := db.Prepare("insert into tasks(task) values(?)")
	defer stmt.Close()
	_, err := stmt.Exec(task)
	if err != nil {
		return false, err
	}
	return true, nil
}

// ReadTodosTaskFromDB function return todos task from db
// return list to string and error
func ReadTodosTaskFromDB() ([]string, error) {
	var tasks []string
	var task string
	stmt, err := db.Prepare("select task from tasks")
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		rows, err := stmt.Query()
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				rows.Scan(&task)
				tasks = append(tasks, task)
			}
		}
	}
	return tasks, err
}

// MarkTaskAsDone function remove the task from db
// return type status and error
func MarkTaskAsDone(ids []int) ([]int, []string, error) {
	var notExist []int
	var taskDone []string
	tasks, err := ReadTodosTaskFromDB()
	if err != nil {
		fmt.Printf("%v", err)
	} else {
		var deleteTask []string
		for _, id := range ids {
			if id-1 < len(tasks) {
				deleteTask = append(deleteTask, tasks[id-1])
			} else {
				notExist = append(notExist, id-1)
			}
		}
		for _, task := range deleteTask {
			stmt, err := db.Prepare("delete from tasks where task=?")
			if err == nil {
				taskDone = append(taskDone, task)
				_, err = stmt.Exec(task)
			}
		}
	}
	return notExist, taskDone, err
}
