package cmd

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/dubbkins/glow/reporting"
	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedReportingServer
}

func (s *server) Run(ctx context.Context, in *pb.Report) (*pb.ReportResult, error) {
	if in.GetIsAsync() {
		go func() {
			log.Printf("Received async: %v", in.GetName())
		}()
		return &pb.ReportResult{Id: (uuid.New().String())}, nil
	}
	log.Printf("Received: %v", in.GetName())
	return &pb.ReportResult{Id: (uuid.New().String())}, nil
}

var server_command = &cobra.Command{

	Use:   "server",
	Short: "grpc server",
	Long:  `grpc server`,
	Run: func(cmd *cobra.Command, args []string) {

		listener, err := net.Listen("tcp", fmt.Sprintf(":%d", server_port))
		if err != nil {
			log.Fatalf("failed to start listener on port %d: %v", server_port, err)
		}

		s := grpc.NewServer()
		pb.RegisterReportingServer(s, &server{})
		log.Printf("server listening at %v", listener.Addr())
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}
