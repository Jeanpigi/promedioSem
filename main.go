package main

import "fmt"

type Subject struct {
	id int
	name  string
	academicCredits int
	grade float32
}

type Student struct {
	name string
	subjects []Subject
}
// TotalLending is a function which calculate total sum academic credits of student
func TotalLending(e *Student) int {
	total := 0
	for _, value := range e.subjects {
		total += value.academicCredits
	}
	return total
}
// ParcialSum is a function which firstly multiply grades and academic credits, then sum total of this result
func ParcialSum(e *Student) int {
	total := 0
	for _, value := range e.subjects {
		intGrade := int(value.grade)
		total += value.academicCredits * intGrade
	}

	return total
}
// Average is a function which is calculated by dividing the sum of the values in the set by their number.
func Average(a, b int) float32 {
	total := float32(a) / float32(b)
	return total
}

func main() {

	jeanpierre := Student{
		name: "Jean Pierre",
		subjects:  []Subject {
			{
				id: 1101,
				name:  "Programming",
				academicCredits: 5,
				grade: 3.0,
			},
			{
				id: 1010,
				name:  "Biology",
				academicCredits: 1,
				grade: 5.0,
			},
			{
				id: 2088,
				name:  "Chemistry",
				academicCredits: 3,
				grade: 4.0,
			},
			{
				id: 1120,
				name:  "Mathematics",
				academicCredits: 3,
				grade: 3.0,
			},
			{
				id: 3020,
				name:  "English",
				academicCredits: 3,
				grade: 5.0,
			},
		},
	}

	fmt.Println(Average(ParcialSum(&jeanpierre), TotalLending(&jeanpierre)))

}
