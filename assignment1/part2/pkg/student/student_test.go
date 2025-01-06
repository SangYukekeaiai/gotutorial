package student

import (
	"math"
	"testing"
)

func TestNewStudent(t *testing.T) {
	s := NewStudent("John Doe", "12345")

	if s.Name != "John Doe" {
		t.Errorf("expected Name to be 'John Doe', got '%s'", s.Name)
	}
	if s.ID != "12345" {
		t.Errorf("expected ID to be '12345', got '%s'", s.ID)
	}
	if s.Major != "" {
		t.Errorf("expected Major to be empty, got '%s'", s.Major)
	}
	if s.age != 0 {
		t.Errorf("expected Age to be 0, got '%d'", s.age)
	}
	if len(s.Courses) != 0 {
		t.Errorf("expected Courses to be empty, got '%d'", len(s.Courses))
	}
}

func TestAddCourse(t *testing.T) {
	s := NewStudent("John Doe", "12345")
	s.AddCourse("Math 101", 3.0, 4.0)

	if len(s.Courses) != 1 {
		t.Errorf("expected 1 course, got '%d'", len(s.Courses))
	}

	course := s.Courses[0]
	if course.Name != "Math 101" {
		t.Errorf("expected Course Name to be 'Math 101', got '%s'", course.Name)
	}
	if course.Credit_hrs != 3.0 {
		t.Errorf("expected CreditHours to be 3.0, got '%f'", course.Credit_hrs)
	}
	if course.Grade != 4.0 {
		t.Errorf("expected Grade to be 4.0, got '%f'", course.Grade)
	}
}

func TestCalculateGPA(t *testing.T) {
	s := NewStudent("John Doe", "12345")

	// Test GPA for no courses
	gpa := s.CalculateGPA()
	if gpa != 0.0 {
		t.Errorf("expected GPA to be 0.0, got '%f'", gpa)
	}

	// Add courses
	s.AddCourse("Math 101", 3.0, 4.0) // A grade
	s.AddCourse("History 101", 2.0, 3.0) // B grade
	s.AddCourse("Physics 101", 1.0, 2.0) // C grade

	// Expected GPA: (4*3 + 3*2 + 2*1) / (3 + 2 + 1) = 3.33...
	expectedGPA := (4.0*3.0 + 3.0*2.0 + 2.0*1.0) / (3.0 + 2.0 + 1.0)

	gpa = s.CalculateGPA()
	if math.Abs(gpa-expectedGPA) > 1e-9 {
		t.Errorf("expected GPA to be '%f', got '%f'", expectedGPA, gpa)
	}
}
