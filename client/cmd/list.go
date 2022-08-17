package cmd

import (
	"context"
	"fmt"
	"go_client/app"
	"log"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all key-value pairs",
	Long:  "Lists all key-value pairs",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create gRPC connection
		conn, connErr := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if connErr != nil {
			log.Fatalf("Could not init connection: %v", connErr)
		}
		defer conn.Close()

		client := app.NewDictServiceClient(conn)

		list, err := client.GetPairs(context.Background(), &app.Empty{})
		if err != nil {
			return err
		}

		for _, pair := range list.GetPairs() {
			fmt.Printf("%v=%v\n", pair.GetKey(), pair.GetValue())
		}

		return nil
	},
}
