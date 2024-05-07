/*package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"main/proto"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := proto.NewEmployeeServiceClient(conn)

	// Call the GetAllEmployees RPC
	response, err := client.GetAllEmployees(context.Background(), &proto.GetAllEmployeesRequest{})
	if err != nil {
		log.Fatalf("Error calling GetAllEmployees RPC: %v", err)
	}

	// Process the response
	for _, employee := range response.Employees {
		log.Printf("Employee ID: %d, Name: %s, Mobile: %d, Email: %s", employee.Id, employee.Name, employee.Mobile, employee.Email)
	}
}
*/

package main

import (
	"context"
	"log"
	"main/proto"

	"google.golang.org/grpc"
)

func main() {

	/*
		ERROR: panic: sql: Register called twice for driver postgres
		goroutine 1 [running]:
		database/sql.Register({0x1385b43, 0x8}, {0x144bde0, 0x18283a0})

		CAUSE
		The error you encountered, "sql: Register called twice for driver postgres,
		" typically occurs when the sql.Register function is called more than once
		for the same driver. In your code, the sql.Register function is called
		explicitly in the client code and implicitly within the
		fetchAllEmployeesFromDatabase function. This implicit call happens because
		fetchAllEmployeesFromDatabase attempts to open a connection to the PostgreSQL
		 database without explicitly registering the driver


	*/
	// Set up a connection to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client instance
	client := proto.NewEmployeeServiceClient(conn)

	// Call the GetAllEmployees RPC
	response, err := client.GetAllEmployees(context.Background(), &proto.GetAllEmployeesRequest{})
	if err != nil {
		log.Fatalf("Error calling GetAllEmployees RPC: %v", err)
	}

	// Process the response
	for _, employee := range response.Employees {
		log.Printf("Employee ID: %d, Name: %s, Mobile: %d, Email: %s", employee.Id, employee.Name, employee.Mobile, employee.Email)
	}
}
