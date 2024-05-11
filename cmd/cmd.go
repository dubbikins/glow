package cmd

import (
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:   "glow",
	Short: "",
	Long:  `.`,
}

var (
	server_port int
	client_port int
)

func Execute() {
	cmd.Execute()
}

func init() {
	server_command.PersistentFlags().IntVarP(&server_port, "port", "p", 8080, "The port to listen on")
	client_command.PersistentFlags().IntVarP(&client_port, "port", "p", 8081, "The port to listen on")
	// cmd.PersistentFlags().StringVarP(&cert_file, "cert", "c", "/etc/letsencrypt/live/haus-of-kins.com/fullchain.pem", "The certificate file to use for TLS")
	// cmd.PersistentFlags().StringVarP(&key_file, "key", "k", "/etc/letsencrypt/live/haus-of-kins.com/privkey.pem", "The key file to use for TLS")
	cmd.AddCommand(
		server_command,
		client_command,
	)

}
