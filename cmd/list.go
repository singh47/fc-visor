package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/singh47/fc-visor/firecracker"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List running Firecracker VMs",
	Run: func(cmd *cobra.Command, args []string) {
		sockets := firecracker.DiscoverSockets()
		for _, sock := range sockets {
			fmt.Println("Found Firecracker socket:", sock)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
