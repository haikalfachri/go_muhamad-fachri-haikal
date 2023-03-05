package main

import "fmt"

type Student struct {
    name  string
    score float64
}

type arrStudent []Student

func (s arrStudent) getAvgSkor() float64 {
	var total float64 = 0.0
    for _, student := range s {
        total += student.score
    }
    return total / float64(len(s))
}

func (s arrStudent) getMinSkor() Student {
    var minStudent Student
    for _, student := range s {
        if student.score < minStudent.score || minStudent.name == "" {
            minStudent = student
        }
    }
    return minStudent
}

func (s arrStudent) getMaxSkor() Student {
    var maxStudent Student
    for _, student := range s {
        if student.score > maxStudent.score || maxStudent.name == "" {
            maxStudent = student
        }
    }
    return maxStudent
}

func main() {
	var lenStudent int = 5
	var arr arrStudent = make(arrStudent, lenStudent)
	var student Student

	for i := range arr{
		fmt.Print("Input ", i+1, " Student's Name ")
		fmt.Scan(&student.name)
		fmt.Print("Input ", i+1, " Student's Score ")
		fmt.Scan(&student.score)
		arr[i] = student
	}
	var result Student

    fmt.Printf("Rata-rata skor siswa: %.2f\n", arr.getAvgSkor())

    result = arr.getMinSkor()
    fmt.Printf("Siswa dengan skor terendah: %s (%f)\n", result.name, result.score)

    result = arr.getMaxSkor()
    fmt.Printf("Siswa dengan skor tertinggi: %s (%f)\n", result.name, result.score)
}