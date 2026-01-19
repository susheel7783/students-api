package storage

import "github.com/susheel7783/students-api/internal/types"

type Storage interface {
	// Define storage interface methods here
	CreateStudent(name string, email string, age int) (int64, error)
	GetStudentById(id int64) (types.Student, error)
	GetStudents() ([]types.Student, error)
}
