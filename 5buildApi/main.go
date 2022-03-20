package main

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
	return c.CourseId == "" && c.CourseName == ""
}