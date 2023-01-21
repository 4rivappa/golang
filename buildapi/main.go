package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware or helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Building API in GoLang")

	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "4", CourseName: "GoLang", CoursePrice: 200, Author: &Author{FullName: "hitesh chowdary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "6", CourseName: "React", CoursePrice: 100, Author: &Author{FullName: "hitesh chowdary", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "8", CourseName: "Python", CoursePrice: 300, Author: &Author{FullName: "hitesh chowdary", Website: "go.dev"}})
	courses = append(courses, Course{CourseId: "2", CourseName: "Rust", CoursePrice: 200, Author: &Author{FullName: "hitesh chowdary", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", ServeHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// serving
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1> Home Page </h1><h2> -from func ServeHome()</h2>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting id from response
	params := mux.Vars(r)
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	returnString := "No course found with given id: " + string(params["id"])
	json.NewEncoder(w).Encode(returnString)
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// checking if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Error: Body is empty")
		return
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	// checking for this case {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Error: Passed an empty object")
		return
	}

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// checking if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Error: Body is empty")
		return
	}

	var course Course
	json.NewDecoder(r.Body).Decode(&course)

	// checking for this case {}
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Error: Passed an empty object")
		return
	}

	// MissUnderstanding id is sent through url, not through the json object (put)
	// if course.CourseId == "" {
	// 	json.NewEncoder(w).Encode("Error: Id not passed")
	// 	return
	// }
	params := mux.Vars(r)

	for index, iter_course := range courses {
		if iter_course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			course.CourseId = params["id"]
			courses = append(courses, course)
			// json.NewEncoder(w).Encode("Success: course updated")
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("Error: Given Id is not found")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// checking if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Error: Body is empty")
		return
	}

	// This is not at all required - user doesn't give any json input
	// var course Course
	// json.NewDecoder(r.Body).Decode(&course)

	// // checking for this case {}
	// if course.IsEmpty() {
	// 	json.NewEncoder(w).Encode("Error: Passed an empty object")
	// 	return
	// }

	// MissUnderstanding id is sent through url, not through the json object (put)
	// if course.CourseId == "" {
	// 	json.NewEncoder(w).Encode("Error: Id not passed")
	// 	return
	// }

	params := mux.Vars(r)

	for index, iter_course := range courses {
		if iter_course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Success: course deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("Error: Given Id is not found")
	return
}
