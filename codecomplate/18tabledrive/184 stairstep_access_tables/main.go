package main

func main() {
	studentScore := float32(65.0)
	studentGrade := GetGreade(studentScore)
	println(studentGrade)
}

//阶梯访问表
func GetGreade(score float32) string {
	var rangeLimit = []float32{50.0, 65.0, 75.0, 90.0, 100.0}
	var grades = []string{"F", "D", "C", "B", "A"}

	maxGradeLevel := len(grades) - 1

	gradeLevel := 0
	studentGrade := "A"

	for studentGrade == "A" && gradeLevel < maxGradeLevel {
		if score < rangeLimit[gradeLevel] {
			studentGrade = grades[gradeLevel]
		}
		gradeLevel++
	}
	return studentGrade
}
