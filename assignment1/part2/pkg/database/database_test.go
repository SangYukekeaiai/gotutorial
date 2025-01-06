package database

import (
	// "math"
	"testing"
)

func TestNewDatabase (t *testing.T) {
	db := NewDatabase()

	if len(db.students) != 0 {
		t.Errorf("expected students list to be empty")
	}
}

func TestAddStudent(t *testing.T) {
	db := NewDatabase()
	db.AddStudent("YY", "173346")
	if len(db.students) != 1{
		t.Errorf("expected 1 student, got '%d'", len(db.students))
	}
	studentName := db.students[0].Name
	if studentName != "YY" {
		t.Errorf("expected student name is 'YY', got '%s'", studentName)
	}
	studentID := db.students[0].ID
	if studentID != "173346" {
		t.Errorf("expected student name is 'YY', got '%s'", "173346")
	}
}

func TestFindStudentByID(t *testing.T) {
	db := NewDatabase()
	db.AddStudent("YY", "173346")
	db.AddStudent("KK", "173350")

	stu := db.find_student_by_id("173346")
	if stu == nil {
		t.Errorf("expected student to be found, got nil")
	}
	if stu.Name != "YY" {
		t.Errorf("expected student name to be 'YY', got '%s'", stu.Name)
	}

	nonexistentStudent := db.find_student_by_id("999999")
	if nonexistentStudent != nil {
		t.Errorf("expected nil for nonexistent student, got '%v'", nonexistentStudent)
	}
}
func TestFindStudentByName(t *testing.T) {
	db := NewDatabase()
	db.AddStudent("YY", "173346")
	db.AddStudent("KK", "173350")

	stu := db.find_student_by_name("YY")
	if stu == nil {
		t.Errorf("expected student to be found, got nil")
	}
	if stu.ID != "173346" {
		t.Errorf("expected student ID to be '173346', got '%s'", stu.ID)
	}

	nonexistentStudent := db.find_student_by_name("Unknown")
	if nonexistentStudent != nil {
		t.Errorf("expected nil for nonexistent student, got '%v'", nonexistentStudent)
	}
}

func TestFindStudentsByCourse(t *testing.T) {
	db := NewDatabase()
	stu1 := db.AddStudent("YY", "173346")
	stu2 := db.AddStudent("KK", "173350")

	stu1.AddCourse("Math", 3.0, 4.0)
	stu2.AddCourse("Math", 3.0, 3.0)
	stu2.AddCourse("Science", 2.0, 3.5)

	mathStudents := db.FindStudentsByCourse("Math")
	if len(mathStudents) != 2 {
		t.Errorf("expected 2 students taking Math, got '%d'", len(mathStudents))
	}

	scienceStudents := db.FindStudentsByCourse("Science")
	if len(scienceStudents) != 1 {
		t.Errorf("expected 1 student taking Science, got '%d'", len(scienceStudents))
	}
	if scienceStudents[0].Name != "KK" {
		t.Errorf("expected student 'KK' to be taking Science, got '%s'", scienceStudents[0].Name)
	}

	noStudents := db.FindStudentsByCourse("History")
	if len(noStudents) != 0 {
		t.Errorf("expected 0 students taking History, got '%d'", len(noStudents))
	}
}