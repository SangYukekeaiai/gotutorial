package student

type Course_struct struct {
	Name string
	Credit_hrs float64
	Grade float64
}

type Student struct {
	Name string
	ID string
	Major string
	age int
	Courses []Course_struct
}

func NewStudent(name, id string) *Student {
	return &Student{
		Name: name,
		ID: id,
		Major: "",
		age: 0,
		Courses: []Course_struct{},
	}
}

func (s *Student) AddCourse(name string, creditHours, grade float64)  {
	s.Courses = append(s.Courses, Course_struct{Name: name, Credit_hrs: creditHours, Grade: grade})
}

func (s *Student) CalculateGPA() float64 {
	if len(s.Courses) == 0 {
		return 0.0
	}
	var totalPoints, totalHours float64
	for _, course := range s.Courses{
		totalPoints += course.Grade * course.Credit_hrs
		totalHours += course.Credit_hrs
	}

	return totalPoints / totalHours
}
