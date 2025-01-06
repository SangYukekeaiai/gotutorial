package database

import (
	"csci780example/assignment1/part2/pkg/student"
)

type Database struct {
	students []*student.Student
}

func NewDatabase() *Database {
	return &Database{
		students: []*student.Student{},
	}
}

func (db *Database) AddStudent(name, id string) *student.Student {
	newStudent := student.NewStudent(name, id)
	db.students = append(db.students, newStudent)
	return newStudent
}

func (db *Database) find_student_by_id(id string) *student.Student {
	for _, student := range db.students {
		if student.ID == id {
			return student
		}
	}
	return nil
}

func (db *Database) find_student_by_name(name string) *student.Student {
	for _, student := range db.students {
		if student.Name == name {
			return student
		}
	}
	return nil
}

func (db *Database) FindStudentsByCourse(name string) []*student.Student {
	var result []*student.Student
	for _, student := range db.students {
		for _, course := range student.Courses {
			if course.Name == name {
				result = append(result, student)
				break
			}
		}
	}
	return result
}
