package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

type Student struct {
	Name       string
	Id         int
	Age        int
	SignUpTime time.Time
}

var students = map[int]Student{}

func memUsage() {
	fmt.Println("--------------------------- Memory usage -------------------------------")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc : %.2f MiB\n", float64(m.Alloc)/1024/1024)
	fmt.Printf("Total Alloc : %.2f MiB\n", float64(m.TotalAlloc)/1024/1024)
	fmt.Printf("Sys : %.2f MiB\n", float64(m.Sys)/1024/1024)
	fmt.Printf("NumGC : %v\n", m.NumGC)
	fmt.Println("-----------------------------------------------------------------------------------")
}

func registerStudent() {
	fmt.Println("-------------------- Student sign up ------------------")
	rand.Seed(time.Now().UnixNano())
	var s Student
	fmt.Print("Enter your name: ")
	fmt.Scan(&s.Name)

	for {
		var ageInput string
		fmt.Print("Enter your age: ")
		fmt.Scan(&ageInput)
		age, err := strconv.Atoi(ageInput)
		if err != nil {
			fmt.Println("Invalid age. Please enter a number.")
			continue
		}
		s.Age = age
		break
	}

	for {
		s.Id = rand.Intn(90000) + 10000
		if _, exists := students[s.Id]; !exists {
			break
		}
	}

	s.SignUpTime = time.Now()
	students[s.Id] = s

	fmt.Println("--------------------------- You're successfully registered ------------------------")
	fmt.Println("Name          : ", s.Name)
	fmt.Println("Age           : ", s.Age)
	fmt.Println("Student ID    : ", s.Id)
	fmt.Println("Registration  : ", s.SignUpTime.Format("2006-01-02 15:04:05"))
	fmt.Println("-----------------------------------------------------------------------------------")
}

func searchStudent() {
	fmt.Print("Enter student ID to search: ")
	var input string
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID. Please enter a number.")
		return
	}

	student, ok := students[id]
	if ok {
		fmt.Println("------------------------- Student Found -------------------------------")
		fmt.Printf("Name       : %s\n", student.Name)
		fmt.Printf("Age        : %d\n", student.Age)
		fmt.Printf("ID         : %d\n", student.Id)
		fmt.Printf("SignUpTime : %s\n", student.SignUpTime.Format("2006-01-02 15:04:05"))
		fmt.Println("-------------------------------------------------------------------")
	} else {
		fmt.Println("Student not found. Make sure you've entered the correct ID.")
	}
}

func changeNameByID() {
	fmt.Print("Enter student ID to change name: ")
	var input string
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID.")
		return
	}

	student, exists := students[id]
	if !exists {
		fmt.Println("Student not found.")
		return
	}

	fmt.Print("Enter the new name: ")
	fmt.Scan(&student.Name)

	students[id] = student
	fmt.Println("Name successfully changed to", student.Name)
}

func deleteStudent() {
	var input string
	fmt.Printf("Enter the student id : ")
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid ID, please enter a number")
		return
	}

	if _, ok := students[id]; ok {
		delete(students, id)
		fmt.Println("Student successfuly deleted")

	} else {
		fmt.Println("Student not found")

	}

}

func main() {
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1 - Register Student")
		fmt.Println("2 - Search Student")
		fmt.Println("3 - change Name")
		fmt.Println("4 - Delete Student")
		fmt.Println("5 - Memory Usage")
		fmt.Println("6 - Exit")
		fmt.Print("Enter choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			registerStudent()
		case 2:
			searchStudent()
		case 3:
			changeNameByID()
		case 4:
			deleteStudent()
		case 5:
			memUsage()
		case 6:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
