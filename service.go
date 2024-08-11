package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

type Habit struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

func InitDatabase() {
	var err error
	DB, err = sql.Open("sqlite3", "./habits.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS habits (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status TEXT
	);`)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateHabit(title string, status string) (int64, error) {
	result, err := DB.Exec("INSERT INTO habits (title, status) VALUES (?, ?)", title, status)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteHabit(id int64) error {
	_, err := DB.Exec("DELETE FROM habits WHERE id = ?", id)

	return err
}

func ReadHabits() []Habit {
	rows, _ := DB.Query("SELECT id, title, status FROM habits")
	defer rows.Close()

	habits := make([]Habit, 0)
	for rows.Next() {
		var habit Habit
		rows.Scan(&habit.Id, &habit.Title, &habit.Status)
		habits = append(habits, habit)
	}

	return habits
}