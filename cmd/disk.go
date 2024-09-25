package cmd

import (
	"encoding/csv"
	"fmt"
	 "os"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/disk"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "diskStats",
	Short: "Get disk statistics",
	Run: func(cmd *cobra.Command, args []string) {
		writeToCSV, _ := cmd.Flags().GetBool("csv")
		diskStats(writeToCSV)
	},
}

func diskStats(writeToCSV bool) {
	partitionStats, err := disk.Partitions(true)
	if err != nil {
		fmt.Println("Error getting disk partitions:", err)
		return
	}
	if writeToCSV {
		file, err := os.Create("disk_stats.csv")
		if err != nil {
			color.Red("Error creating CSV file:", err)
			return
		}
		defer file.Close()
		writer := csv.NewWriter(file)
		defer writer.Flush()
		writer.Write([]string{"Device", "Mountpoint", "Fstype", "Opts"})
		for _, partition := range partitionStats {
			writer.Write([]string{
				partition.Device,
				partition.Mountpoint,
				partition.Fstype,
				partition.Opts,
			})
		}
	}
}

func init() {
	rootCmd.AddCommand(diskCmd)
}
