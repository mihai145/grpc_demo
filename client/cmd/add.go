package cmd

import (
	"context"
	"errors"
	"fmt"
	"go_client/app"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add <key> <value>",
	Short: "Adds a key-value pair in the database",
	Long:  "Adds a key-value pair in the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create gRPC connection
		conn, connErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if connErr != nil {
			log.Fatalf("Could not init connection: %v", connErr)
		}
		defer conn.Close()

		client := app.NewDictServiceClient(conn)

		if len(args) != 2 {
			return errors.New("expected 2 args: add <key> <value>")
		}

		res, err := client.AddPair(context.Background(), &app.KeyValuePair{Key: args[0], Value: args[1]})
		if err != nil {
			return err
		}
		fmt.Println(res.GetOk())

		return nil
	},
}
