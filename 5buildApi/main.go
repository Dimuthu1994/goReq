package main

import (
	"encoding/json"
	"fmt"
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

//fake DB
var courses []Course

//middleware
func (c *Course) IsEmpty() bool {
	return  c.CourseName == ""
}

//handle get req
func serveHome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Welcome to API by Dimuthu</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type","application/json")
	
	//grab id from req
	params := mux.Vars(r)
	//loop through courses, find matching id and return response
	for _,course := range courses{
		if course.CourseId == params["id"]{
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found on given id")
	return
}

func addCourse(w http.ResponseWriter, r *http.Request){
	fmt.Println("add a Course")
	w.Header().Set("Content-Type","application/json")
	
	if r.Body == nil{
		json.NewEncoder(w).Encode("please send some data")
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty(){
		json.NewEncoder(w).Encode("No data in JSON")
		return
	}

	//generate unique id
	//id to string
	//append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	
	params := mux.Vars(r)
	for index,course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	//TODO: send a respond when id is not valid
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for index,course := range courses{
		if course.CourseId == params["id"]{
			courses = append(courses[:index],courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(courses)
}