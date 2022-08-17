package cmd

import (
	"context"
	"fmt"
	"go_client/app"
	"go_client/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all key-value pairs",
	Long:  "Lists all key-value pairs",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get grpc connection
		conn, connErr := utils.DialGrpc()
		if connErr != nil {
			return connErr
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
