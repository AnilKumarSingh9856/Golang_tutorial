package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"math/rand"

	"github.com/gorilla/mux"
)

// Model for course - file
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
var coursesDB []Course

// middleware, helper - file

func (c *Course) IsEmpty() bool {
	return  c.CourseId == "" && c.CourseName == ""
}

func main() {

}

// Sending a API json response for all courses in golang
// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by LearnCodeOnline </h1>"))	
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Gel all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(coursesDB)
}

// get one course based on request id in golang
func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through courses, find matching id and return the response
	for _, course := range coursesDB{
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("Did not find any course with given id")
	return
}

//  add a course controller in golang
func addOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Adding one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty

	if r.Body == nil{
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	// generate unique id, string
	// append course into courseDB

	course.CourseId = strconv.Itoa(rand.Intn(100))
	coursesDB = append(coursesDB, course)
	json.NewEncoder(w).Encode(course)
	return
}