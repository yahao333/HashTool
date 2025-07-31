package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/yahao333/hashtool/internal/web"
)

var port int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start web server",
	Long:  `Start web server to provide hash generation and verification services via REST API.`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf(":%d", port)
		fmt.Printf("Starting HashTool server on port %d...\n", port)
		fmt.Printf("Server running at http://localhost%s\n", addr)
		fmt.Println("API Endpoints:")
		fmt.Println("  POST /api/hash   - Generate hash")
		fmt.Println("  POST /api/verify - Verify hash")
		
		router := web.SetupRoutes()
		log.Fatal(http.ListenAndServe(addr, router))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run server on")
}