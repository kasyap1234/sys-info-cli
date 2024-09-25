package cmd

import (
	"encoding/csv"
	"fmt"
	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
	"os"
	"time"
)

var ramCmd = &cobra.Command{
	Use:   "ram",
	Short: "Get RAM statistics",
	Run: func(cmd *cobra.Command, args []string) {
		ramStats()
	},
}
var ramCSVCmd = &cobra.Command{
	Use:   "ram-csv",
	Short: "Write Ram Statistics to CSV file",
	Run: func(cmd *cobra.Command, args []string) {
		ramCSV()
	},
}

func ramStats() {
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		color.Yellow("x----------------------------------------------------------------x")
		color.Cyan("RAM Statistics")

		v, _ := mem.VirtualMemory()
		totalGB := float64(v.Total) / (1024 * 1024 * 1024)
		usedGB := float64(v.Used) / (1024 * 1024 * 1024)
		freeGB := float64(v.Free) / (1024 * 1024 * 1024)
		fmt.Printf("Total: %.2f GB, Used: %.2f GB, Free: %.2f GB, Used Percentage: %.2f%%\n", totalGB, usedGB, freeGB, v.UsedPercent)

		fmt.Println("x----------------------------------------------------------------x")
	}
}
func ramCSV() {
	file, err := os.Create("ram_stats.csv")
	if err != nil {
		color.Red("Error creating CSV file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Timestamp", "Total RAM (GB)", "Used RAM (GB)", "Free RAM (GB)", "Used RAM Percentage"})
	ticker := time.NewTicker(3 * time.Second)
	defer ticker.Stop()
	for range ticker.C {
		v, _ := mem.VirtualMemory()
		totalGB := float64(v.Total) / (1024 * 1024 * 1024)
		usedGB := float64(v.Used) / (1024 * 1024 * 1024)
		freeGB := float64(v.Free) / (1024 * 1024 * 1024)

		timestamp := time.Now().Format(time.RFC3339)
		writer.Write([]string{
			timestamp,
			fmt.Sprintf("%.2f", totalGB),
			fmt.Sprintf("%.2f", usedGB),
			fmt.Sprintf("%.2f", freeGB),
			fmt.Sprintf("%.2f", v.UsedPercent),
		})
	}
}

func init() {
	rootCmd.AddCommand(ramCmd)
	rootCmd.AddCommand(ramCSVCmd)

}
