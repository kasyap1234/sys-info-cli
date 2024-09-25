package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/process"
	"github.com/spf13/cobra"
)

var processCmd = &cobra.Command{
	Use:   "processList",
	Short: "Get all process list currently running on the system ",
	Run: func(cmd *cobra.Command, args []string) {
		processStats()
	},
}

func processStats() {
	processesList, err := process.Processes()
	if err != nil {
		color.Red("Error getting process list:", err)
		return
	}
	color.Yellow("PID | Name | Status | CPU Usage | Memory Usage")
	for _, process := range processesList {
		processName, err := process.Name()
		if err != nil {
			color.Red("Error getting process name:", err)
			continue
		}
		processStatus, err := process.Status()
		if err != nil {
			color.Red("Error getting process status:", err)
			continue
		}
		cpuUsage, err := process.CPUPercent()
		if err != nil {
			color.Red("Error getting CPU usage:", err)
			continue
		}
		memoryUsage, err := process.MemoryInfo()
		if err != nil {
			color.Red("Error getting memory usage:", err)
			continue
		}
		fmt.Printf("%d | %s | %s | %.2f%% | %d bytes\n", process.Pid, processName, processStatus, cpuUsage, memoryUsage.RSS)
	}

}

func init() {
	rootCmd.AddCommand(processCmd)

}
