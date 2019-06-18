package project

type Context struct {
	Name        string
	Version     string
	Description string
}

var Ctx = Context{
	Name:        "Typical-RESTful-Server",
	Version:     "0.1.0",
	Description: "Example of typical and scalable RESTful API Server for Go",
}