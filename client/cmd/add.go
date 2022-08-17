package cmd

import (
	"context"
	"errors"
	"fmt"
	"go_client/app"
	"go_client/utils"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add <key> <value>",
	Short: "Adds a key-value pair in the database",
	Long:  "Adds a key-value pair in the database",
	RunE: func(cmd *cobra.Command, args []string) error {
		// Get grpc connection
		conn, connErr := utils.DialGrpc()
		if connErr != nil {
			return connErr
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
