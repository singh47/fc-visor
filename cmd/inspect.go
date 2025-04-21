package cmd

import (
	"fmt"
	"time"
	"github.com/spf13/cobra"
	"github.com/rivo/tview" 
	"github.com/singh47/fc-visor/firecracker"
)

var socketPath string

var inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect a specific Firecracker VM",
	Run: func(cmd *cobra.Command, args []string) {
		if socketPath == "" {
			fmt.Println("Please provide --socket path to the Firecracker API socket")
			return
		}
		info, err := firecracker.GetVMInfo(socketPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("VM Info for %s:\n%+v\n", socketPath, info)
	},
}

var metricsCmd = &cobra.Command{
	Use:   "metrics",
	Short: "Get Firecracker VM metrics",
	Run: func(cmd *cobra.Command, args []string) {
		if socketPath == "" {
			fmt.Println("Please provide --socket path to the Firecracker API socket")
			return
		}
		metrics, err := firecracker.GetVMMetrics(socketPath)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Printf("Metrics for %s:\n%+v\n", socketPath, metrics)
	},
}

var topCmd = &cobra.Command{
	Use:   "top",
	Short: "Live monitor Firecracker VM metrics",
	Run: func(cmd *cobra.Command, args []string) {
		if socketPath == "" {
			fmt.Println("Please provide --socket path to the Firecracker API socket")
			return
		}
		app := tview.NewApplication()
		text := tview.NewTextView().SetDynamicColors(true).SetRegions(true).SetChangedFunc(func() {
			app.Draw()
		})
		go func() {
			for {
				metrics, err := firecracker.GetVMMetrics(socketPath)
				if err != nil {
					text.SetText(fmt.Sprintf("Error: %v", err))
					return
				}
				output := fmt.Sprintf("[yellow]CPU:[white] %dus  MemRSS: %dKB  NetRx: %dB  NetTx: %dB\n",
					metrics.CPUUsageUs,
					metrics.MemoryRSS,
					metrics.NetworkRxBytes,
					metrics.NetworkTxBytes)
				text.SetText(output)
				time.Sleep(1 * time.Second)
			}
		}()
		if err := app.SetRoot(text, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}
	},
}

func init() {
	inspectCmd.Flags().StringVar(&socketPath, "socket", "", "Path to Firecracker socket")
	metricsCmd.Flags().StringVar(&socketPath, "socket", "", "Path to Firecracker socket")
	topCmd.Flags().StringVar(&socketPath, "socket", "", "Path to Firecracker socket")
	rootCmd.AddCommand(inspectCmd)
	rootCmd.AddCommand(metricsCmd)
	rootCmd.AddCommand(topCmd)
}