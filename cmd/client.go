package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/dubbkins/glow/reporting"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client_command = &cobra.Command{

	Use:   "client",
	Short: "grpc client",
	Long:  `grpc client`,
	Run: func(cmd *cobra.Command, args []string) {
		var conn *grpc.ClientConn
		var err error
		conn, err = grpc.Dial(fmt.Sprintf(":%d", server_port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()

		client := reporting.NewReportingClient(conn)
		response, err := client.Run(context.Background(), &reporting.Report{Name: "test"})
		if err != nil {
			log.Fatalf("could not run: %v", err)
		}
		fmt.Println(response)
	},
}
