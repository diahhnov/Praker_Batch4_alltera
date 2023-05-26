package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Score int
}

type Students []Student

func (s Students) calculateAverageScore() float64 {
	jumlah := 0

	for _, student := range s {
		jumlah += student.Score
	}
	rataRata := float64(jumlah) / float64(len(s))
	return rataRata
}

func (s Students) findMinScore() Student {
	minScore := math.MaxInt32
	var minScoreStudent Student

	for _, student := range s {
		if student.Score < minScore {
			minScore = student.Score
			minScoreStudent = student
		}
	}
	return minScoreStudent
}

func (s Students) findMaxScore() Student {
	maxScore := math.MinInt32
	var maxScoreStudent Student

	for _, student := range s {
		if student.Score > maxScore {
			maxScore = student.Score
			maxScoreStudent = student
		}
	}
	return maxScoreStudent
}

func main() {
	students := Students{}

	for i := 0; i < 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name: ", i+1)
		fmt.Scanln(&name)

		fmt.Printf("Input %d Student's Score: ", i+1)
		fmt.Scanln(&score)

		student := Student{
			Name:  name,
			Score: score,
		}

		students = append(students, student)
	}

	rataSkor := students.calculateAverageScore()
	skorTerendah := students.findMinScore()
	skorTertinggi := students.findMaxScore()

	fmt.Printf("rata-rata Skor: %.2f\n", rataSkor)
	fmt.Printf("Min Score of Students: %s (%d)\n", skorTerendah.Name, skorTerendah.Score)
	fmt.Printf("Max Score of Students: %s (%d)\n", skorTertinggi.Name, skorTertinggi.Score)
}
