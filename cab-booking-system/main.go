package main 

import (
	"github.com/google/uuid"
	"fmt"
	"math"
)

type CabBookingSystem struct {
	Drivers map[string]*Driver
	Riders map[string]*Rider
	Trips map[string]*Trip
}

type Driver struct {
	ID string 
	Name string 
	Location [2]int 
	Available bool 
}

type Rider struct {
	ID string 
	Name string 
	Location [2]int
	TripHistory []*Trip
}

type Trip struct {
	ID string 
	Rider *Rider 
	Driver *Driver
	StartLocation [2]int
	EndLocation [2]int
	Status string
}

const (
	DriverStatusAvailable = "Available"
	DriverStatusOnTrip = "On Trip"
	TripStatusNotStarted = "Not Started"
	TripStatusInProgress = "In Progress"
	TripStatusCompleted = "Completed"
)

func (CBS *CabBookingSystem) RegisterDriver(name string, location [2]int) *Driver {
	for _, d := range CBS.Drivers {
		if d.Name == name {
			fmt.Println("A driver with same name has already been registered:", name)
			return nil 
		}
	}

	driver := &Driver{
		ID: uuid.New().String(),
		Name: name,
		Location: location,
		Available: true,
	}
	CBS.Drivers[driver.ID] = driver
	fmt.Printf("Driver %s registered with ID %s\n", driver.Name, driver.ID)
	return driver
} 

func (CBS *CabBookingSystem) RegisterRider(name string, location [2]int) *Rider {
	for _, r := range CBS.Riders {
		if r.Name == name {
			fmt.Println("A rider with same name has already been registered:", name)
			return nil
		}
	}

	rider := &Rider{
		ID: uuid.New().String(),
		Name: name,
		Location: location,
	}
	CBS.Riders[rider.ID] = rider 
	fmt.Printf("Rider %s registered with ID %s\n", rider.Name, rider.ID)
	return rider
}

func (CBS *CabBookingSystem) BookCab(riderID string, destinationLocation [2]int) *Trip {
	rider, exists := CBS.Riders[riderID]
	if !exists {
		fmt.Println("Rider does not exists")
		return nil
	}

	var nearestDriver *Driver
	shortestDistance := math.MaxFloat64

	for _, driver := range CBS.Drivers {
		if driver.Available {
			distance := geteuclideanDistance(rider.Location, driver.Location)
			if distance < shortestDistance {
				shortestDistance = distance
				nearestDriver = driver 
			}
		}
	}

	if nearestDriver == nil {
		fmt.Println("No available drivers at this moment")
		return nil 
	}

	nearestDriver.Available = false 

	trip := &Trip {
		ID: uuid.New().String(),
		Rider: rider,
		Driver: nearestDriver,
		StartLocation: rider.Location,
		EndLocation: destinationLocation,
		Status: TripStatusInProgress,
	}

	CBS.Trips[trip.ID] = trip 
	fmt.Printf("\nTrip Booked: Rider %s with Driver %s\n", rider.Name, nearestDriver.Name)
	return trip 
}

func geteuclideanDistance(a, b [2]int) float64 {
	x := float64(a[0] - b[0])
	y := float64(a[1] - b[1])
	return math.Sqrt(x*x + y*y)
}


func (CBS *CabBookingSystem) StartTrip(tripID string) string {
	trip, found := CBS.Trips[tripID]
	if !found {
		fmt.Println("Trip not found")
		return "Trip not found"
	}

	trip.Status = TripStatusInProgress
	fmt.Printf("Trip %s started\n", tripID)
	return "Trip Started"
}

func (CBS *CabBookingSystem) EndTrip(tripID string) string {
	trip, found := CBS.Trips[tripID]
	if !found {
		fmt.Println("Trip not found")
		return "Trip not found"
	}

	trip.Status = TripStatusCompleted
	trip.Driver.Available = true
	trip.Driver.Location = trip.EndLocation
	trip.Rider.Location = trip.EndLocation
	trip.Rider.TripHistory = append(trip.Rider.TripHistory, trip)
	fmt.Printf("Trip %s completed\n", tripID)
	return "Trip Completed"
}

func (CBS *CabBookingSystem) GetDriverStatus(driverID string) string{
	driver, exists := CBS.Drivers[driverID]
	if !exists {
		return "Driver Not Found"
	}

	if driver.Available {
		return DriverStatusAvailable
	}
	return DriverStatusOnTrip
}

func (CBS *CabBookingSystem) GetRiderHistory(riderID string) []*Trip {
	rider, found := CBS.Riders[riderID]
	if !found {
		fmt.Println("Rider not found")
		return nil
	}

	if len(rider.TripHistory) ==  0 {
		fmt.Println("Rider trip history is empty")
		return nil
	}

	return rider.TripHistory
}

func main(){
	fmt.Println("Welcome to Samuel's Cab Booking System")

	system := &CabBookingSystem{
		Drivers: make(map[string]*Driver),
		Riders: make(map[string]*Rider),
		Trips: make(map[string]*Trip),
	}

	// To Register Drivers
	driver1 := system.RegisterDriver("Paulson", [2]int{2, 3})
	driver2 := system.RegisterDriver("Daniel", [2]int{5, 5})
	driver3 := system.RegisterDriver("Sammy", [2]int{6, 6})
	
	// To Register Rider
	rider1 := system.RegisterRider("Swag", [2]int{7, 2})
	rider2 := system.RegisterRider("Samuel", [2]int{9, 2})
	rider3 := system.RegisterRider("Treasure", [2]int{6, 6})

	// To Book a trip, start trip and end the trip
	trip1 := system.BookCab(rider1.ID, [2]int{3, 8})
	if trip1 != nil {
		system.StartTrip(trip1.ID)
		system.EndTrip(trip1.ID)
	}

	trip2 := system.BookCab(rider2.ID, [2]int{9, 8})
	if trip2 != nil {
		system.StartTrip(trip2.ID)
		system.EndTrip(trip2.ID)
	}
	
	trip3 := system.BookCab(rider3.ID, [2]int{9, 9})
	if trip3 != nil {
		system.StartTrip(trip3.ID)
		system.EndTrip(trip3.ID)
	}

	// To get the driaver status
	fmt.Printf("\nDriver %s is available: %v\n", driver1.Name, system.GetDriverStatus(driver1.ID))
	fmt.Printf("\nDriver %s is available: %v\n", driver2.Name, system.GetDriverStatus(driver2.ID))
	fmt.Printf("\nDriver %s is available: %v\n", driver3.Name, system.GetDriverStatus(driver3.ID))

	for _, rider := range []*Rider{rider1, rider2, rider3} {
		fmt.Printf("\nTrip History for Rider %s:\n", rider.Name)
		trips := system.GetRiderHistory(rider.ID)
		if trips == nil {
			continue
		}

		for _, trip := range trips {
			fmt.Printf("Trip ID: %s | From %v -> To: %v | Driver: %s | Status: %s\n", trip.ID, trip.StartLocation, trip.EndLocation, trip.Driver.Name, trip.Status)
		}
	}
}
