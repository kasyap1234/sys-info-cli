package cmd 
import (
    "fmt"
    "github.com/shirou/gopsutil/v3/cpu"
    "github.com/spf13/cobra"
)
var cpuCmd = &cobra.Command{
    Use:   "cpu",
    Short: "Get CPU statistics",
    Run: func(cmd *cobra.Command, args []string) {
       cpuStats();
    },
}

func cpuStats() {
    percent, _ := cpu.Percent(0, false)
	cpuInfo, _ := cpu.Info()
	fmt.Printf("CPU Model: %s\n", cpuInfo[0].ModelName)
	fmt.Printf("CPU Cores: %d\n", cpuInfo[0].Cores)
	fmt.Printf("CPU Frequency: %f GHz\n", cpuInfo[0].Mhz/1000)
	fmt.Printf("Cpu Family: %s\n", cpuInfo[0].Family)
    fmt.Printf("CPU Usage: %f%%\n", percent[0])
}