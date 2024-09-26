package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/spf13/cobra"
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
	Args:  cobra.MaximumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		directory := ""
		if len(args) > 0 {
			directory = args[0]

		}
		if err := ramCSV(directory); err != nil {
			color.Red("Error writing to CSV file %s", err)
		}
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
func ramCSV(directory string) error {
	// file, err := os.Create("ram_stats.csv")
	if directory == "" {
		directory = "."
	}
	filename := "ram_stats.csv"
	filePath := path.Join(directory, filename)
	file, err := os.Create(filePath)

	if err != nil {
		color.Red("Error creating CSV file:", err)
		return err

	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	writer.Write([]string{"Timestamp", "Total RAM (GB)", "Used RAM (GB)", "Free RAM (GB)", "Used RAM Percentage"})
	ticker := time.NewTicker(3 * time.Second)
	log.Println("CSV file created successfully.")
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
		writer.Flush()
		color.Green("CSV file updated successfully.")

	}
	
	return nil 
}


func init() {
	rootCmd.AddCommand(ramCmd)
	rootCmd.AddCommand(ramCSVCmd)

}
