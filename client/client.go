package main

import (
	"context"
	"fmt"
	"go_client/app"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// Parse command flag
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Expected args. Usage:\n list -> list all pairs\n add <key> <value> -> adds pair to db")
		os.Exit(0)
	}

	command := args[0]
	if command != "list" && command != "add" {
		fmt.Println("Command can be list or add")
		os.Exit(0)
	}

	// Create gRPC connection
	conn, connErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if connErr != nil {
		log.Fatalf("Could not init connection: %v", connErr)
	}
	defer conn.Close()

	client := app.NewDictServiceClient(conn)

	// Call gRPC
	if command == "list" {
		list, err := client.GetPairs(context.Background(), &app.Empty{})
		if err != nil {
			fmt.Printf("Error occured: %v\n", err)
			os.Exit(1)
		}

		for _, pair := range list.GetPairs() {
			fmt.Printf("%v=%v\n", pair.GetKey(), pair.GetValue())
		}
	} else if command == "add" {
		if len(args) != 3 {
			fmt.Println("Expected args. Usage:\n list -> list all pairs\n add <key> <value> -> adds pair to db")
			os.Exit(0)
		}

		res, err := client.AddPair(context.Background(), &app.KeyValuePair{Key: args[1], Value: args[2]})
		if err != nil {
			fmt.Printf("Error occured: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(res.GetOk())
	}
}
