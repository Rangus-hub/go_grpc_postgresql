package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	pb "main/proto"

	"github.com/jackc/pgx/v4"
	"github.com/lib/pq"
	"google.golang.org/grpc"
)

func init() {
	fmt.Println("Initializing PostgreSQL driver...")
	// Register the PostgreSQL driver
	sql.Register("postgres", &pq.Driver{})
	fmt.Println("PostgreSQL driver initialized.")
}

type server struct {
	conn *pgx.Conn

	pb.EmployeeServiceServer
}

func (s *server) GetAllEmployees(ctx context.Context, req *pb.GetAllEmployeesRequest) (*pb.GetAllEmployeesResponse, error) {

	create_sql := `
	create table if not exists Employee(
		id SERIAL PRIMARY KEEY, 
		name text, 
		mobile int,
		email text
	);
	`

	_, err := s.conn.Exec(context.Background(), create_sql)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Table creation failed: %v\n", err)
		os.Exit(1)
	}

	log.Printf("Received: %v", req.GetName())

	// Fetch all employees from the database
	employees, err := fetchAllEmployeesFromDatabase()
	if err != nil {
		return nil, err
	}

	// Convert database records to protobuf messages
	var employeeMessages []*pb.Employee
	for _, employee := range employees {
		employeeMessages = append(employeeMessages, &pb.Employee{
			Id:     int32(employee.Id),
			Name:   employee.Name,
			Mobile: employee.Mobile,
			Email:  employee.Email,
		})
	}

	return &pb.GetAllEmployeesResponse{Employees: employeeMessages}, nil
}

func fetchAllEmployeesFromDatabase() ([]*pb.Employee, error) {
	// Establish a connection to the PostgreSQL database
	db, err := sql.Open("postgres", "user=postgres password=PostGrePsd dbname=Company sslmode=disable")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Prepare the SQL query to fetch all employees
	rows, err := db.Query("SELECT id, name, mobile, email FROM employee")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the result set and build a slice of Employee structs
	var employees []*pb.Employee
	for rows.Next() {
		var employee pb.Employee
		if err := rows.Scan(&employee.Id, &employee.Name, &employee.Mobile, &employee.Email); err != nil {
			return nil, err
		}
		employees = append(employees, &employee)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterEmployeeServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	fmt.Println("Server Running on port: 50051")
}
