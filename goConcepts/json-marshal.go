package main

import (
	"encoding/json"
	"fmt"
)

// ServicePortName struct
type ServicePortName struct {
	Name string
	Port string
	x    int
}

// ServiceEndpoint struct
type ServiceEndpoint struct {
	Endpoint        string
	x               int
	ServicePortName ServicePortName
}

func main() {
	svc := ServiceEndpoint{
		"/api/users", 1, ServicePortName{
			"app.user", "9000", 2,
		},
	}

	s, _ := json.MarshalIndent(svc, "", "\t")
	fmt.Println(string(s))
	fmt.Printf("%+v", svc)
}
