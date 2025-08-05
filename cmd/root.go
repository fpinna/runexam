package cmd

import (
	"fmt"
	"os"

	"runexam/server"

	"github.com/spf13/cobra"
)

var (
	port     int
	listen   string
	jsonFile string
)

var rootCmd = &cobra.Command{
	Use:   "runexam [json_file]",
	Short: "RunExam - Certification exam simulator",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		jsonFile = args[0]
		if _, err := os.Stat(jsonFile); os.IsNotExist(err) {
			fmt.Printf("File not found: %s\n", jsonFile)
			os.Exit(1)
		}
		server.StartServer(jsonFile, listen, port)
	},
}

func init() {
	rootCmd.Flags().IntVarP(&port, "port", "p", 9171, "Server port")
	rootCmd.Flags().StringVarP(&listen, "listen", "l", "0.0.0.0", "Address to listen on")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
