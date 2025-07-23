package sqlite

import (
	"database/sql"

	"github.com/namsh70747/Rest_API/internal/config"
	"github.com/namsh70747/Rest_API/internal/types"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	Db *sql.DB
}

func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL UNIQUE,
	age INTEGER NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		return nil, err
	}
	return &Sqlite{Db: db}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Sqlite) GetStudentById(id int) (types.Student, error) {
	stmt, err := s.Db.Prepare("SELECT id,name,email,age,created_at,updated_at FROM students WHERE id = ?")
	if err != nil {
		return types.Student{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var student types.Student
	err = row.Scan(&student.Id, &student.Name, &student.Email, &student.Age, &student.CreatedAt, &student.UpdatedAt)
	if err != nil {
		return types.Student{}, err
	}
	return student, nil
}
