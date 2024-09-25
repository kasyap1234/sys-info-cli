package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
    Use:   "disk",
    Short: "Get disk statistics",
    Run: func(cmd *cobra.Command, args []string) {
       diskStats();
    },
}

func diskStats() {
	partitionStats, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Error getting disk partitions:", err)
		return
	}
	color.Yellow("Device | Mountpoint | Fstype | Opts")
	for _, partition := range partitionStats {
		fmt.Printf("%s | %s | %s | %s\n", partition.Device, partition.Mountpoint, partition.Fstype, partition.Opts)
	}
}

func init() {
    rootCmd.AddCommand(diskCmd)
}


