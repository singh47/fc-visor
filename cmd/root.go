package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fc-visor",
	Short: "Firecracker VM metrics & info tool",
	Long:  `fc-visor lets you inspect and monitor Firecracker microVMs easily.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}