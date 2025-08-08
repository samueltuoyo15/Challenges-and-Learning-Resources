package main

import "fmt"

// Repository pattern in Go, three key concepts for me to understand
// 1. Interface
// 2. Implementation
// 3. Constructor

type Rider struct {
	name   string
	status string
}

// interface
type RiderRepository interface {
	registerRider(name string) string
	getRiderStatus(name string) string
}

// implementation
type RiderInterfaceImplementationStruct struct {
	riders map[string]*Rider
}

func (r *RiderInterfaceImplementationStruct) registerRider(name string) string {
	if r.riders == nil {
		r.riders = make(map[string]*Rider)
	}
	r.riders[name] = &Rider{
		name:   name,
		status: "Available",
	}

	return name
}

func (r *RiderInterfaceImplementationStruct) getRiderStatus(name string) string {
	if rider, exists := r.riders[name]; exists {
		return rider.status
	}
	return "Rider not found"
}

// constructor
func NewRiderRepositoryInterface() *RiderInterfaceImplementationStruct {
	return &RiderInterfaceImplementationStruct{
		riders: make(map[string]*Rider),
	}
}

func main() {
	fmt.Println("Rider Repository Pattern Example")
	registerRiderRepo := NewRiderRepositoryInterface()

	fmt.Println("Rider registered successfully:", registerRiderRepo.registerRider("Sam"))
	fmt.Println("Rider registered successfully:", registerRiderRepo.registerRider("Mark"))

	fmt.Println("Sam Status", registerRiderRepo.getRiderStatus("Sam"))
	fmt.Println("Mark Status", registerRiderRepo.getRiderStatus("Mark"))
}
