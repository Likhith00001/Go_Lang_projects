package main

import (
	"fmt"
)

type Marks struct {
	S_name string
	mark   float64
}
type Student struct {
	name    string
	s_id    int
	marks   Marks //slice or map
	total   float64
	average float64
	grade   string
}

func CraeteStudent() {
	s1 := Student{}
	fmt.Println("Enter Student name")
	fmt.Scanln(&s1.name)
	fmt.Println("Enter Student ID")
	fmt.Scanln(&s1.s_id)
	fmt.Println("Enter number of subjects")
	num := 0
	fmt.Scanln(&num)
	CreateMarks(&s1, num)
	TotalMarks(&s1, num)
	AverageMarks(&s1, num)
	Grade(&s1)
	Display(&s1, num)
}
func CreateMarks(s1 *Student, num int) {

	for i := 0; i < num; i++ {
		fmt.Println("enter subject name and marks")
		fmt.Scanln(&s1.marks.S_name, &s1.marks.mark)
		if s1.marks.mark > 100 || s1.marks.mark < 0 {
			fmt.Println("wrong marks entered")
			//s1.marks = 0
			CreateMarks(s1, num)
		}
	}

}
func TotalMarks(s1 *Student, num int) {
	Total := 0.0
	for i := 0; i < num; i++ {
		Total += s1.marks.mark
	}
	s1.total = Total

}

func AverageMarks(s1 *Student, num int) {
	s1.average = s1.total / float64(num)

}

func Grade(s1 *Student) {

	switch {
	case s1.average <= 100 && s1.average >= 90:
		s1.grade = "A"
	case s1.average < 90 && s1.average >= 80:
		s1.grade = "B"
	case s1.average < 80 && s1.average >= 70:
		s1.grade = "C"
	case s1.average < 50 && s1.average >= 60:
		s1.grade = "D"
	case s1.average < 60 && s1.average >= 50:
		s1.grade = "E"
	default:
		s1.grade = "F"
	}
}

func Display(s1 *Student, num int) {
	fmt.Println("Student Details: \n", "Student Name: ", s1.name, "\n Student Roll number: ", s1.s_id)
	for i := 0; i < num; i++ {
		fmt.Println("Subjectc: ", s1.marks.S_name, " Marks: ", s1.marks.mark)
	}
	fmt.Println("Total Marks: ", s1.total, "Average Marks: ", s1.average, "Grade: ", s1.grade)
}

func StartStudent() {
	fmt.Println("Welcome to the student information data: ")
	CraeteStudent()

}

func main() {
	StartStudent()

}
