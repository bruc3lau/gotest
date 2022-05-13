package main

import "fmt"

type student struct {
	id   int
	name string
}

func main() {
	var students []student
	students = append(students, student{1, "john"})
	students = append(students, student{2, "adam"})
	s1 := student{3, "bruce"}
	fmt.Printf("%v, %+v\n", &s1, &s1)
	students = append(students, s1)
	fmt.Println(students)

	for idx, s := range students {
		if s.id == 2 {
			students = append(students[:idx], students[idx+1:]...)
		}
	}
	fmt.Println(students)
	fmt.Println(666)
}
