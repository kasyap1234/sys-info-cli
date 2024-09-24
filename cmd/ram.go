package cmd
import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/shirou/gopsutil/v3/mem"
)

var ramCmd = &cobra.Command{
    Use:   "ram",
    Short: "Get RAM statistics",
    Run: func(cmd *cobra.Command, args []string) {
       ramStats();
    },
}

func ramStats() {
    v, _ := mem.VirtualMemory()
    totalGB := float64(v.Total) / (1024 * 1024 * 1024)
    usedGB := float64(v.Used) / (1024 * 1024 * 1024)
    freeGB := float64(v.Free) / (1024 * 1024 * 1024)
    fmt.Printf("Total: %.2f GB, Used: %.2f GB, Free: %.2f GB, Used Percentage: %.2f%%\n", totalGB, usedGB, freeGB, v.UsedPercent)
}

func init() {
    rootCmd.AddCommand(ramCmd)
}
