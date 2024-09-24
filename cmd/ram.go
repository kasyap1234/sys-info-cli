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
    fmt.Printf("Total: %v, Used: %v, Free: %v, Used Percentage: %f%%\n", v.Total, v.Used, v.Free, v.UsedPercent)
}
func init() {
    rootCmd.AddCommand(ramCmd)
}
