package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "fc-visor",
	Short:   "Firecracker VM metrics & info tool",
	Long:    `fc-visor lets you inspect and monitor Firecracker microVMs easily.`,
	Version: "0.1.0",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: "Show help for fc-visor and its subcommands",
		Run: func(cmd *cobra.Command, args []string) {
			_ = rootCmd.Usage()
		},
	})
}
