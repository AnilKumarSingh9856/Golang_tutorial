package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

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
	return c.CourseName == ""
}

func main() {

	fmt.Println("API - LearnCodeOnline.in")
	r := mux.NewRouter()

	// seeding
	coursesDB = append(coursesDB, Course{CourseId: "2" , CourseName: "Golang", CoursePrice: 244, 
	Author: &Author{FullName: "Anil Kumar Singh", Website: "hitech.com"}})	

	coursesDB = append(coursesDB, Course{CourseId: "3" , CourseName: "Rust", CoursePrice: 300, 
	Author: &Author{FullName: "Anil Kumar Singh", Website: "hitech.com"}})	

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", addOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	




	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

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

	// TODO: LOOP, REMOVE DUPLICATE COURSE
	for _, old_course := range coursesDB{
		if old_course.CourseName == course.CourseName{
			json.NewEncoder(w).Encode("Cannot add same course twice")
			return
		}
	}

	// generate unique id, string
	// append course into courseDB

	course.CourseId = strconv.Itoa(rand.Intn(100))
	coursesDB = append(coursesDB, course)
	json.NewEncoder(w).Encode(course)
	return
}


// update a course controller in golang
func updateOneCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("Updating one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove, add new content with same id

	for index, course := range coursesDB{
		if course.CourseId == params["id"]{
			coursesDB = append(coursesDB[:index], coursesDB[index+1:]...)
			var newCourse Course
			_ = json.NewDecoder(r.Body).Decode(&newCourse)
			newCourse.CourseId = params["id"]
			coursesDB = append(coursesDB, newCourse)
			json.NewEncoder(w).Encode(newCourse)
			return
		}
	}
	// TODO: send a response when id is not found
}


// delete a course controller in golang
func deleteOneCourse (w http.ResponseWriter, r *http.Request){
	fmt.Println("deleting one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop, id, remove (index, index+1)

	for index, course := range coursesDB{
		if course.CourseId == params["id"]{
			coursesDB = append(coursesDB[:index], coursesDB[index+1:]...)
			return
		}
	}
}